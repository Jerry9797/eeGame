package logic

import (
	"eeGame/internal/model/entity"
	"eeGame/internal/websocket/framework/net"
	"eeGame/internal/websocket/game/models/request"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RoomManager struct {
	sync.RWMutex
	roomMap map[string]*Room // key: roomId
}

func (um *RoomManager) CreatRoom(session *net.Session, req request.CreatRoomReq, user *entity.User) error {
	// 创建房间
	roomId := um.creatRoomId()
	newRoom := NewRoom(session, req, user, roomId)
	um.roomMap[roomId] = newRoom
	newRoom.UsersMap = make(map[int]*RoomUser)
	return newRoom.UserEntryRoom(session, user)
}

func (um *RoomManager) GetRoomById(roomId string) *Room {
	r, ok := um.roomMap[roomId]
	if ok {
		return r
	}
	return nil
}

func (um *RoomManager) creatRoomId() string {
	// 生成房间号 6位数
	rand.New(rand.NewSource(time.Now().UnixNano()))
	roomId := rand.Int63n(999999)
	if roomId < 100000 {
		roomId = roomId + 100000
	}
	roomIdStr := fmt.Sprintf("%d", roomId)

	r := um.roomMap[roomIdStr]
	if r != nil {
		// 已存在roomId，重新创建
		um.creatRoomId()
	}
	return roomIdStr
}

func NewRoomManager() *RoomManager {
	return &RoomManager{
		roomMap: make(map[string]*Room),
	}
}
