package trade

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TopUpReq struct {
	g.Meta       ` tags:"TopUp" method:"post" summary:"充值"`
	AmountToMint int64
}
type TopUpRes struct {
}

type WithdrawReq struct {
	g.Meta     ` tags:"Withdraw" method:"post" summary:"提现"`
	D711Amount int64 `v:"required" json:"d711Amount"`
}

type WithdrawRes struct {
}
