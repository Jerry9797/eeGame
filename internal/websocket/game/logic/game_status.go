package logic

import (
	"eeGame/internal/model/entity"
	"github.com/ethereum/go-ethereum/common"
)

type UserInfo struct {
	Uid      int    `json:"uid"`
	UnionID  int64  `json:"unionID"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
}

type RoomUser struct {
	UserInfo   *entity.User   `json:"userInfo"`
	ChairID    int            `json:"chairID"`    //座次
	UserStatus UserStatus     `json:"userStatus"` //房间内玩家状态
	ViewCard   bool           `json:"viewCard"`
	Bets       []int64        `json:"bets"` // 下注记录
	Address    common.Address `json:"address"`
	Cards      []Card         `json:"cards"`
}

type UserStatus int

const (
	None    UserStatus = 0
	Ready              = 1
	Playing            = 2
	Offline            = 4
	Dismiss            = 8
)
