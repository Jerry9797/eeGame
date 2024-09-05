package request

import (
	"eeGame/internal/model/entity"
	"eeGame/internal/websocket/game/logic/component"
)

type CreatRoomReq struct {
	GameRuleId string             `json:"gameRuleId"`
	GameRule   component.GameRule `json:"gameRule"`
}

type EnterRoomReq struct {
	RoomId string      `json:"roomId"`
	Uid    int         `json:"uid"`
	User   entity.User `json:"user"`
}
