package handler

import (
	"context"
	"eeGame/internal/dao"
	"eeGame/internal/logic/trade"
	"eeGame/internal/websocket/framework/net"
	"eeGame/internal/websocket/game/logic"
	"eeGame/internal/websocket/game/models/request"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"math/big"
	"strconv"
)

type RoomHandler struct {
	um *logic.RoomManager
}

func (h *RoomHandler) CreateRoom(session *net.Session, msg []byte) (any, error) {
	uid := session.Uid
	if uid == 0 {
		return fmt.Errorf("user doesn't exist"), nil
	}
	var req request.CreatRoomReq
	if err := json.Unmarshal(msg, &req); err != nil {
		return fmt.Errorf("msg Unmarshal error:%v \n", err), nil
	}
	// 查询用户信息
	user, err2 := dao.User.GetUserById(context.TODO(), uid)
	if err2 != nil {
		logrus.Errorf("GetUserById error: %v\n", err2)
		return fmt.Errorf("GetUserById error: %v\n", err2), nil
	}
	if user == nil {
		logrus.Errorf("user doesn't exist")
		return fmt.Errorf("user doesn't exist"), nil
	}
	// 创建房间
	room := h.um.CreatRoom(session, req, user)
	return room, nil
}

func (h *RoomHandler) UserEnterRoom(session *net.Session, msg []byte) (any, error) {
	var entryR request.EnterRoomReq
	err := json.Unmarshal(msg, &entryR)
	if err != nil {
		return nil, err
	}
	rm := h.um.GetRoomById(entryR.RoomId)
	errU := rm.UserEntryRoom(session, &entryR.User)
	if errU != nil {
		return nil, errU
	}
	return nil, nil
}

func (h *RoomHandler) UserBetting(session *net.Session, msg []byte) (any, error) {
	var bettingR request.UserBettingReq
	err := json.Unmarshal(msg, &bettingR)
	if err != nil {
		return nil, err
	}
	room := h.um.GetRoomById(bettingR.RoomId)
	user, ok := room.UsersMap[bettingR.Uid]
	if !ok {
		return nil, fmt.Errorf("用户【%v】不存在", bettingR.Uid)
	}

	room.Betting(bettingR, user)
	return nil, nil
}

func (h *RoomHandler) SettleAccount(session *net.Session, msg []byte) (any, error) {
	var bettingR request.UserBettingReq
	err := json.Unmarshal(msg, &bettingR)
	if err != nil {
		return nil, err
	}
	room := h.um.GetRoomById(bettingR.RoomId)
	winUser := room.UsersMap[bettingR.Uid]
	fmt.Printf("%v", winUser)
	values := make([]*logic.RoomUser, 0, len(room.UsersMap))
	fmt.Printf("%v", values)

	newTrade, err := trade.NewTrade(context.Background())

	if err != nil {
		return nil, err
	}

	for _, value := range room.UsersMap {
		if value.UserInfo.Uid == bettingR.Uid {
			continue
		}
		user := room.UsersMap[bettingR.Uid]
		bets := user.Bets
		sum := int64(0)
		if len(bets) > 0 {
			for _, bet := range bets {
				sum += bet
			}
		}
		// 转账
		var userAddress []common.Address
		roomId, err := strconv.ParseInt(bettingR.RoomId, 10, 64)
		err = newTrade.StartGame(context.Background(), big.NewInt(roomId), userAddress, big.NewInt(50))
		if err != nil {
			return nil, err
		}
		//err := newTrade.Transfer(value.Address, room.UsersMap[bettingR.Uid].Address, sum)
		if err != nil {
			logrus.Errorf("Transfer from:%v to:%v error: %v", value.Address, room.UsersMap[bettingR.Uid].Address, err)
		}
	}
	return nil, nil
}

func NewRoomHandler(um *logic.RoomManager) *RoomHandler {
	return &RoomHandler{
		um: um,
	}
}
