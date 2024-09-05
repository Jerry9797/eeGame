package route

import (
	"eeGame/internal/websocket/framework/net"
	"eeGame/internal/websocket/game/handler"
	"eeGame/internal/websocket/game/logic"
)

func Register() net.LogicHandle {
	handlers := make(net.LogicHandle)
	entryHandler := handler.NewEntryHandler()
	// 进入游戏，校验用户
	handlers["entryHandler.entry"] = entryHandler.Entry

	um := logic.NewRoomManager()
	roomHandler := handler.NewRoomHandler(um)
	// 创建房间
	handlers["roomHandler.createRoom"] = roomHandler.CreateRoom
	// 加入房间，...开始游戏
	handlers["roomHandler.userEnterRoom"] = roomHandler.UserEnterRoom
	// 下注 Betting，
	handlers["roomHandler.userBetting"] = roomHandler.UserBetting
	// 结算
	handlers["roomHandler.settleAccount"] = roomHandler.SettleAccount

	return handlers
}
