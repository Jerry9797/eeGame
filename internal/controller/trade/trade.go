package trade

import (
	"context"
	v1 "eeGame/api/trade"
	"eeGame/internal/consts"
	"eeGame/internal/logic/trade"
	"fmt"
	"math/big"
)

type Trade struct {
}

func New() *Trade {
	return &Trade{}
}

func (t *Trade) TopUp(ctx context.Context, req *v1.TopUpReq) (res *v1.TopUpRes, err error) {
	_, ok := ctx.Value(consts.CtxUserId).(int)
	if !ok {
		return res, fmt.Errorf("uid不存在")
	}

	newTrade, err := trade.NewTrade(ctx)
	if err != nil {
		return nil, err
	}
	if err := newTrade.TopUp(ctx, big.NewInt(req.AmountToMint*1e18)); err != nil {
		return nil, err
	}
	return nil, nil
}

func (t *Trade) Withdraw(ctx context.Context, req *v1.WithdrawReq) (res *v1.WithdrawRes, err error) {
	_, ok := ctx.Value(consts.CtxUserId).(int)
	if !ok {
		return res, fmt.Errorf("uid不存在")
	}
	newTrade, err := trade.NewTrade(ctx)
	if err != nil {
		return nil, err
	}
	if err := newTrade.Withdraw(ctx, big.NewInt(req.D711Amount*1e18)); err != nil {
		return nil, err
	}
	return res, nil
}
