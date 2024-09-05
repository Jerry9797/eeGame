package logic

import (
	"context"
	"eeGame/internal/logic/trade"
	"eeGame/internal/model/entity"
	"eeGame/internal/websocket/framework/net"
	"eeGame/internal/websocket/game/logic/component"
	"eeGame/internal/websocket/game/models/request"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"math/rand"
	"sort"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type Room struct {
	RoomId        string
	Round         int
	CurrentBettor int   // 当前下注玩家的 Uid
	LastBet       int64 // 上一次下注的金额
	BettingOrder  []int // 下注顺序
	UserIds       []int
	UsersMap      map[int]*RoomUser
	UserConnMap   map[int]*net.WsConnection
	RoomCreator   *RoomCreator
	GameFrame     *GameFrame
	kickSchedules map[int]*time.Timer
}

type RoomCreator struct {
	Uid int `json:"uid"`
}

func (r *Room) UserEntryRoom(session *net.Session, user *entity.User) error {
	logrus.Infof("玩家【%s】进入房间", user.Username)
	// 进入房间，将玩家加入房间
	_, ok := r.UsersMap[user.Uid]
	if ok {
		return fmt.Errorf("玩家【%s】已在房间内", user.Username)
	}
	r.UsersMap[user.Uid] = &RoomUser{
		UserInfo:   user,
		ChairID:    0,
		UserStatus: None,
	}
	r.UserIds = append(r.UserIds, user.Uid)
	r.UserConnMap[user.Uid] = session.Client
	// 推送房间给客户端
	//r.PushRoom(session, user.Uid)
	session.Put("roomId", r.RoomId)
	// 游戏类型推送给客户端
	//r.PushGameType(session, user.Uid)
	// 通知其他玩家有人进入
	r.NotifyRoomUser(user)

	// 检查是否可以开始游戏
	r.CheckAndStartGame()
	// 超时未准备，踢出房间
	//go r.kickOutOfTimeOutUser(session, user.Uid)
	return nil
}

func (r *Room) PushRoom(session *net.Session, uid int) {
	pushMsg := map[string]any{
		"roomId":     r.RoomId,
		"pushRouter": "UpdateUserInfoPush",
		"userId":     uid,
	}
	client := session.Client
	marshal, _ := json.Marshal(pushMsg)
	client.SendMsg(marshal)
}

func (r *Room) PushGameType(session *net.Session, uid int) {
	pushMsg := map[string]any{
		"gameType":   r.GameFrame,
		"pushRouter": "SelfEntryRoomPush",
		"userId":     uid,
	}
	client := session.Client
	marshal, _ := json.Marshal(pushMsg)
	client.SendMsg(marshal)
}

func (r *Room) PushGameScene(session *net.Session, uid int) {
	roomUserInfoArr := make([]*RoomUser, 0)
	for _, user := range r.UsersMap {
		roomUserInfoArr = append(roomUserInfoArr, user)
	}
	pushMsg := map[string]any{
		"type":       component.GetRoomSceneInfoPush,
		"pushRouter": "RoomMessagePush",
		"data": map[string]any{
			"roomId":          r.RoomId,
			"roomCreatorInfo": r.RoomCreator,
			"gameRule":        r.GameFrame.gameRule,
			"roomUserInfoArr": roomUserInfoArr,
		},
	}
	client := session.Client
	marshal, _ := json.Marshal(pushMsg)
	client.SendMsg(marshal)
}

func (r *Room) RoomMessageHandle(session *net.Session, req request.RoomMessageReq) {

}

func (r *Room) CheckAndStartGame() {
	if len(r.UsersMap) >= r.GameFrame.gameRule.MinPlayerCount {
		r.AssignSeats()
		r.StartGame()
	}
}

func (r *Room) AssignSeats() {
	players := make([]*RoomUser, 0, len(r.UsersMap))
	for _, user := range r.UsersMap {
		players = append(players, user)
	}

	// 随机打乱玩家顺序
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	// 分配座位并设置下注顺序
	r.BettingOrder = make([]int, len(players))
	for i, player := range players {
		player.ChairID = i
		r.BettingOrder[i] = player.UserInfo.Uid
	}
}

func (r *Room) kickOutOfTimeOutUser(session *net.Session, uid int) {
	r.kickSchedules[uid] = time.AfterFunc(30*time.Second, func() {
		logrus.Info("kickOutOfTimeOutUser is running...")
		// 判断用户是否该被踢出
		user, ok := r.UsersMap[uid]
		if !ok {
			logrus.Errorf("user[%v] is kicked", uid)
			return
		}
		if user.UserStatus < Ready {
			delete(r.UsersMap, uid)
			r.kickUser(session, user)
			// 判断房间中有没有人，没人则解散
			if len(r.UsersMap) == 0 {
				r.dismissRoom()
			}
		}
	})
}

func (r *Room) kickUser(session *net.Session, user *RoomUser) {
	// 将user的roomId设置为空，推送用户信息

}

func (r *Room) dismissRoom() {

}

func (r *Room) NotifyRoomUser(user *entity.User) {
	pushMsg := map[string]any{
		"roomId":     r.RoomId,
		"pushRouter": "UpdateUserInfoPush",
		"msg":        user.Username + "进入房间",
	}
	marshal, _ := json.Marshal(pushMsg)
	connMap := r.UserConnMap
	for _, client := range connMap {
		client.SendMsg(marshal)
	}
}

func (r *Room) StartGame() {
	// todo 检查人数是否足够，足够则开始游戏
	le := len(r.UsersMap)
	if le < 2 {
		return
	}
	logrus.Infof("开始游戏")
	r.Round++
	// 洗牌
	deck := New()
	logrus.Infof("deck is: %v", deck)
	cards := deck.DealCard(len(r.UserIds), 3)

	// 发牌,但不立即发送给玩家
	for i, uid := range r.BettingOrder {
		user := r.UsersMap[uid]
		user.Cards = cards[i]
		user.ViewCard = false
	}
	newTrade, err := trade.NewTrade(context.Background())
	if err != nil {
		logrus.Errorf("trade.NewTrade error: %v", err)
		return
	}
	var userAddress []common.Address
	for _, user := range r.UsersMap {
		address := user.Address
		userAddress = append(userAddress, address)
	}
	roomId, err := strconv.ParseInt(r.RoomId, 10, 64)
	if err != nil {
		logrus.Errorf("%v", err)
		return
	}
	err = newTrade.StartGame(context.Background(), big.NewInt(roomId), userAddress, big.NewInt(50))
	if err != nil {
		logrus.Errorf("StartGame error: %v", err)
		return
	}
	// 设置当前下注玩家为第一个玩家
	r.CurrentBettor = r.BettingOrder[0]
	r.LastBet = 0

	// 通知所有玩家游戏开始,包括座位信息和当前下注玩家
	r.NotifyGameStart()

	// 通知第一个玩家下注
	r.NotifyPlayerToBet(r.CurrentBettor)
}

func (r *Room) NotifyGameStart() {
	startInfo := map[string]interface{}{
		"type": "gameStart",
		"data": map[string]interface{}{
			"seats":         r.BettingOrder,
			"currentBettor": r.CurrentBettor,
			"minBet":        r.GameFrame.gameRule.BaseScore,
		},
	}
	marshal, _ := json.Marshal(startInfo)
	for _, client := range r.UserConnMap {
		client.SendMsg(marshal)
	}
}

func (r *Room) NotifyPlayerToBet(uid int) {
	betInfo := map[string]interface{}{
		"type": "yourTurnToBet",
		"data": map[string]interface{}{
			"uid":    uid,
			"minBet": r.LastBet,
		},
	}
	marshal, _ := json.Marshal(betInfo)
	r.UserConnMap[uid].SendMsg(marshal)
}

func (r *Room) ViewCard(uid int) {
	user := r.UsersMap[uid]
	if !user.ViewCard {
		user.ViewCard = true
		cardInfo := map[string]interface{}{
			"type": "viewCard",
			"data": map[string]interface{}{
				"cards": user.Cards,
			},
		}
		marshal, _ := json.Marshal(cardInfo)
		r.UserConnMap[uid].SendMsg(marshal)
	}
}

func (r *Room) Betting(req request.UserBettingReq, rm *RoomUser) {
	values := make([]*RoomUser, 0, len(r.UsersMap))
	for _, value := range r.UsersMap {
		values = append(values, value)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i].ChairID < values[j].ChairID
	})
	chairId := rm.ChairID
	// 下注记录
	rm.Bets = append(rm.Bets, req.Betting)
	// 通知其他玩家，发言人是谁、下注不可低于 req.Betting
	for _, client := range r.UserConnMap {
		msg := map[string]any{
			"spokesman": values[chairId].UserInfo.Uid,
			"betFloor":  req.Betting,
		}
		marshal, _ := json.Marshal(msg)
		client.SendMsg(marshal)
	}
}

func NewRoom(session *net.Session, req request.CreatRoomReq, user *entity.User, roomId string) *Room {
	rc := &RoomCreator{
		Uid: user.Uid,
	}
	r := &Room{
		RoomId:      roomId,
		Round:       0,
		UserIds:     make([]int, 6),
		RoomCreator: rc,
		UserConnMap: make(map[int]*net.WsConnection),
	}
	r.GameFrame = NewGameFrame(component.GameRule{})
	return r
}
