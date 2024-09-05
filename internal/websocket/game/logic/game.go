package logic

import (
	"eeGame/internal/websocket/game/logic/component"
)

type GameFrame struct {
	gameRule component.GameRule
}

func NewGameFrame(gameRule component.GameRule) *GameFrame {
	return &GameFrame{
		gameRule: gameRule,
	}
}
