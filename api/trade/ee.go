package trade

import "github.com/gogf/gf/v2/frame/g"

type DepositAndMintReq struct {
	g.Meta       ` tags:"TopUp" method:"post" summary:"抵押铸造Ee"`
	AmountToMint int64
}
type DepositAndMintRes struct {
}

type RedeemCollateralForEeReq struct {
	g.Meta   ` tags:"Withdraw" method:"post" summary:"赎回"`
	EeAmount int64 `v:"required" json:"EeAmount"`
}

type RedeemCollateralForEeRes struct {
}

type GetBalanceReq struct {
	g.Meta ` tags:"Balance" method:"get" summary:"查询余额"`
}
type GetBalanceRes struct {
	Balance string `json:"balance"`
}

type TransferReq struct {
	g.Meta ` tags:"Transfer" method:"post" summary:"交易"`
	Amount int64 `json:"amount"`
}

type TransferRes struct {
}
