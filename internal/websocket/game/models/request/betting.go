package request

type UserBettingReq struct {
	RoomId  string `json:"roomId"`
	Uid     int    `json:"uid"`
	Betting int64  `json:"betting"`
}
