package ee

import (
	"context"
	v1 "eeGame/api/trade"
	"eeGame/internal/consts"
	"eeGame/internal/logic/trade"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Ee struct {
}

func NewEe() *Ee {
	return &Ee{}
}

func (t *Ee) DepositCollateralAndMintEe(ctx context.Context, req *v1.DepositAndMintReq) (res *v1.DepositAndMintRes, err error) {
	_, ok := ctx.Value(consts.CtxUserId).(int)
	if !ok {
		return res, fmt.Errorf("uid不存在")
	}
	newEe, err := trade.NewEe(ctx)
	if err != nil {
		return nil, err
	}
	go newEe.Listen(ctx)
	if err := newEe.DepositCollateralAndMintEe(ctx, big.NewInt(req.AmountToMint*1e18)); err != nil {
		return nil, err
	}
	return nil, nil
}

func (t *Ee) RedeemCollateralForEe(ctx context.Context, req *v1.RedeemCollateralForEeReq) (res *v1.RedeemCollateralForEeRes, err error) {
	_, ok := ctx.Value(consts.CtxUserId).(int)
	if !ok {
		return res, fmt.Errorf("uid不存在")
	}
	newTrade, err := trade.NewEe(ctx)
	if err != nil {
		return nil, err
	}
	go newTrade.Listen(ctx)
	if err := newTrade.RedeemCollateralForEe(ctx, big.NewInt(req.EeAmount*1e18)); err != nil {
		return nil, err
	}
	return res, nil
}

func (t *Ee) GetBalance(ctx context.Context, req *v1.GetBalanceReq) (res *v1.GetBalanceRes, err error) {
	newTrade, err := trade.NewEe(ctx)
	if err != nil {
		return nil, err
	}
	address, ok := ctx.Value(consts.CtxAddress).(string)
	if !ok {
		return res, fmt.Errorf("address不存在")
	}

	balance, err := newTrade.GetBalanceERC20(common.HexToAddress(address))
	if err != nil {
		return nil, err
	}
	ba := new(big.Int)

	//response.Json(g.RequestFromCtx(ctx), 200, "获取成功！", nil)
	return &v1.GetBalanceRes{
		Balance: ba.Div(balance, big.NewInt(1e18)).String(),
	}, nil
}

func (t *Ee) TransferFrom(ctx context.Context, req *v1.TransferReq) (res *v1.TransferRes, err error) {
	newTrade, err := trade.NewEe(ctx)
	if err != nil {
		return nil, err
	}
	address, ok := ctx.Value(consts.CtxAddress).(string)
	if !ok {
		return res, fmt.Errorf("address不存在")
	}
	fromAddress := common.HexToAddress(address)
	toAddress := common.HexToAddress("0xA8fB23CCe24cbA08013BEe5Ba45B8f90B730ed9a")
	go newTrade.Listen(ctx)
	err1 := newTrade.Transfer(fromAddress, toAddress, req.Amount)
	if err1 != nil {
		return nil, err1
	}
	return nil, nil
}
