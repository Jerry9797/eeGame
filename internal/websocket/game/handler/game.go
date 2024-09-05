package handler

import (
	"eeGame/internal/websocket/framework/net"
	"eeGame/internal/websocket/game/logic"
	"eeGame/internal/websocket/game/models/request"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

type GameHandler struct {
	um *logic.RoomManager
}

func (g *GameHandler) RoomMessageNotify(session *net.Session, msg []byte) (any, error) {
	uid := session.Uid
	if uid <= 0 {
		return fmt.Errorf("user doesn't exist"), nil
	}
	var req request.RoomMessageReq
	if err := json.Unmarshal(msg, &req); err != nil {
		return fmt.Errorf("msg Unmarshal error:%v \n", err), err
	}
	roomId, ok := session.Get("roomId")
	if !ok {
		logrus.Errorf("not in room")
	}
	room := g.um.GetRoomById(fmt.Sprintf("%v", roomId))
	if room == nil {
		return fmt.Errorf("GetRoomById is nil"), nil
	}
	room.RoomMessageHandle(session, req)
	return nil, nil
}

func NewGameHandler(um *logic.RoomManager) *GameHandler {
	return &GameHandler{
		um: um,
	}
}
