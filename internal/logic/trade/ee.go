package trade

import (
	"context"
	"crypto/ecdsa"
	"eeGame/internal/logic/trade/contract"
	"eeGame/internal/logic/trade/contract_ meta"
	"eeGame/utility/response"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/sirupsen/logrus"
	"math/big"
	"strings"
)

type Ee struct {
	client                *ethclient.Client
	wethAddress           common.Address
	D711Engine            *contract__meta.D711Engine
	D711EngineAddress     common.Address
	D711StableCoin        *contract.D711StableCoin
	D711StableCoinAddress common.Address
	fromAddress           common.Address
	userAddress           common.Address
	chainId               *big.Int
	privateKey            *ecdsa.PrivateKey
}

func NewEe(ctx context.Context) (*Ee, error) {
	cfg := g.Cfg().MustGet(ctx, "ETH").Map()
	client := response.EthClient
	privateKey, err := crypto.HexToECDSA(cfg["privateKey"].(string))
	if err != nil {
		return nil, fmt.Errorf("privateKey error : %v", err)
	}
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	d711EngineAddress := common.HexToAddress(cfg["d711EngineAddress"].(string))
	d711StableCoinAddress := common.HexToAddress(cfg["d711StableCoinAddress"].(string))
	chainId := big.NewInt(11155111)

	d711Engine, err := contract__meta.NewD711Engine(d711EngineAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to load DSC Engine contract_metadata: %v", err)
	}
	d711, err := contract.NewD711StableCoin(d711StableCoinAddress, client)
	if err != nil {
		return nil, fmt.Errorf("failed to load DSC Engine contract_metadata: %v", err)
	}
	weth := common.HexToAddress(cfg["wethToken"].(string))

	return &Ee{
		client:                client,
		D711Engine:            d711Engine,
		D711StableCoin:        d711,
		wethAddress:           weth,
		privateKey:            privateKey,
		fromAddress:           fromAddress,
		D711EngineAddress:     d711EngineAddress,
		chainId:               chainId,
		userAddress:           common.HexToAddress(cfg["userAddress"].(string)),
		D711StableCoinAddress: common.HexToAddress(cfg["d711StableCoinAddress"].(string)),
	}, nil
}

func (t *Ee) DepositCollateralAndMintEe(ctx context.Context, amountD711ToMint *big.Int) error {
	nonce, gasFeeCap, gasTipCap, err := t.prepareTransaction(ctx)
	if err != nil {
		return err
	}
	amountCollateral, err := t.getAmountWethForMintingD711(amountD711ToMint)
	if err != nil {
		return err
	}

	operations := []func(nonce uint64) error{
		func(nonce uint64) error {
			// 授权ETH
			return t.approveWETH(nonce, gasFeeCap, gasTipCap, amountCollateral)
		},
		func(nonce uint64) error {
			// 授权D711
			maxUint256 := new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(1))
			return t.approveD711(nonce, gasFeeCap, gasTipCap, maxUint256)
		},
		func(nonce uint64) error {
			return t.depositCollateralAndMintD711(nonce, gasFeeCap, gasTipCap, err, amountCollateral, amountD711ToMint)
		},
	}
	for _, operation := range operations {
		if err := operation(nonce); err != nil {
			return err
		}
		nonce++
	}
	return nil
}

func (t *Ee) getAmountWethForMintingD711(amountD711 *big.Int) (*big.Int, error) {
	d711, err := t.D711Engine.GetAmountWethForMintingD711(t.creatCallOpts(), t.wethAddress, amountD711)
	if err != nil {
		return nil, err
	}
	return d711, err
}

func (t *Ee) depositCollateralAndMintD711(nonce uint64, gasFeeCap *big.Int, gasTipCap *big.Int, err error, amountCollateral *big.Int, amountDscToMint *big.Int) error {
	// 发送交易，抵押&铸币
	auth := t.createTransactor(nonce, gasFeeCap, gasTipCap)
	err = t.getGasLimit(auth, t.D711EngineAddress, contract__meta.D711EngineMetaData.ABI, "depositCollateralAndMintD711",
		t.wethAddress, amountCollateral, amountDscToMint)
	if err != nil {
		return err
	}
	tx, err := t.D711Engine.DepositCollateralAndMintD711(auth, t.wethAddress, amountCollateral, amountDscToMint)
	if err != nil {
		return fmt.Errorf("failed to execute depositCollateralAndMintDsc: %v", err)
	}
	if err := t.waitForTransaction(tx); err != nil {
		return err
	}
	return nil
}

func (t *Ee) prepareTransaction(ctx context.Context) (nonce uint64, gasFeeCap, gasTipCap *big.Int, err error) {
	nonce, err = t.client.PendingNonceAt(ctx, t.fromAddress)
	if err != nil {
		return
	}
	// 获取当前区块Base Feed
	block, err := t.client.BlockByNumber(ctx, nil)
	if err != nil {
		return
	}
	gasTipCap, err = t.client.SuggestGasTipCap(ctx)
	if err != nil {
		return
	}
	// gasFeeCap = baseFee + gasTipCap
	gasFeeCap = new(big.Int).Add(block.BaseFee(), gasTipCap)
	return
}

func (t *Ee) approveWETH(nonce uint64, gasFeeCap, gasTipCap *big.Int, amountCollateral *big.Int) error {
	weth9, err2 := contract__meta.NewWETH9(t.wethAddress, t.client)
	if err2 != nil {
		panic(err2)
	}
	// 如果要返回前端，让用户授权则：auth = &bind.TransactOpts{
	//		From: t.fromAddress,
	//		Signer: nil, // 不签名
	//	}
	auth := t.createTransactor(nonce, gasFeeCap, gasTipCap)
	err := t.getGasLimit(auth, t.wethAddress, contract__meta.WETH9MetaData.ABI, "approve", t.D711EngineAddress, amountCollateral)
	if err != nil {
		return err
	}
	tx, err := weth9.Approve(auth, t.D711EngineAddress, amountCollateral)
	if err != nil {
		return fmt.Errorf("failed to approve collateral: %v", err)
	}
	return t.waitForTransaction(tx)
}

func (t *Ee) createTransactor(nonce uint64, gasFeeCap, gasTipCap *big.Int) *bind.TransactOpts {
	auth, _ := bind.NewKeyedTransactorWithChainID(t.privateKey, t.chainId)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasFeeCap = gasFeeCap
	auth.GasTipCap = gasTipCap
	return auth
}

func (t *Ee) waitForTransaction(tx *types.Transaction) error {
	mined, err := bind.WaitMined(context.Background(), t.client, tx)
	if err != nil {
		return fmt.Errorf("failed to wait for transaction mining: %v", err)
	}
	if mined.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
	}
	fmt.Printf("Transaction succeeded: %s\n", tx.Hash().Hex())
	return nil
}

func (t *Ee) Listen(ctx context.Context) {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{t.D711EngineAddress},
	}
	logChan := make(chan types.Log)
	sub, err := t.client.SubscribeFilterLogs(ctx, query, logChan)
	if err != nil {
		fmt.Printf("SubscribeFilterLogsErr: %v\n", err)
		return
	}
	for {
		select {
		case err := <-sub.Err():
			fmt.Printf("sub.Err: %v\n", err)
		case logs := <-logChan:
			fmt.Printf("logs: %v\n", logs)
		}
	}
}

func (t *Ee) creatCallOpts() *bind.CallOpts {
	callOpts := &bind.CallOpts{
		//From:    common.HexToAddress("0xYourAddress"), // 可选
		Context:     context.Background(),
		BlockNumber: nil, // nil 为最新
	}
	return callOpts
}

func (t *Ee) getGasLimit(auth *bind.TransactOpts, to common.Address, contractAbi string, methodName string, args ...interface{}) error {
	parsed, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		logrus.Errorf("NewReader error")
		return err
	}
	encodedData, err := parsed.Pack(methodName, args...)
	if err != nil {
		logrus.Errorf("parsed pack [%s] error: %v", methodName, err)
		return err
	}
	callMsg := ethereum.CallMsg{
		From:      t.fromAddress,
		To:        &to,
		Data:      encodedData,
		GasFeeCap: auth.GasFeeCap,
		GasTipCap: auth.GasTipCap,
	}
	estimatedGas, err := t.client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		logrus.Errorf("EstimateGas error: %v", err)
		return err
	}
	auth.GasLimit = estimatedGas
	return nil
}

func (t *Ee) getBalance() (*big.Int, *big.Int) {
	opts := t.creatCallOpts()
	information, err := t.D711Engine.GetAccountInformation(opts, t.userAddress)
	if err != nil {
		return nil, nil
	}
	return information.CollateralValueInUsd, information.TotalD711Minted
}

func (t *Ee) GetBalanceERC20(address common.Address) (*big.Int, error) {
	coin, err := contract.NewD711StableCoin(t.D711StableCoinAddress, t.client)
	if err != nil {
		return nil, err
	}
	opts := t.creatCallOpts()
	balance, err := coin.BalanceOf(opts, address)
	return balance, err
}

// 提现
func (t *Ee) RedeemCollateralForEe(ctx context.Context, d711Amount *big.Int) error {
	weth, d711 := t.getBalance()
	logrus.Infof("weth amout: %v", weth)
	logrus.Infof("_711 amout: %v", d711)
	nonce, gasFeeCap, gasTipPrice, err := t.prepareTransaction(ctx)
	if err != nil {
		return err
	}
	if err := t.approveD711(nonce, gasFeeCap, gasTipPrice, d711Amount); err != nil {
		return err
	}
	nonce++
	if err := t.redeemCollateralForDsc(nonce, gasFeeCap, gasTipPrice, d711Amount); err != nil {
		return err
	}
	wethA, d711A := t.getBalance()
	logrus.Infof("wethA amout: %v", wethA)
	logrus.Infof("d711A amout: %v", d711A)
	return nil
}

func (t *Ee) redeemCollateralForDsc(nonce uint64, gasPrice *big.Int, gasTipPrice *big.Int, d711Amount *big.Int) error {
	wethAmount, err := t.getAmountWethForMintingD711(d711Amount)
	if err != nil {
		return err
	}

	auth := t.createTransactor(nonce, gasPrice, gasTipPrice)
	err = t.getGasLimit(auth, t.D711EngineAddress, contract__meta.D711EngineMetaData.ABI, "redeemCollateralForD711",
		t.wethAddress, wethAmount, d711Amount)
	if err != nil {
		return err
	}

	tx, err := t.D711Engine.RedeemCollateralForD711(auth, t.wethAddress, wethAmount, d711Amount)
	if err != nil {
		logrus.Errorf("RedeemCollateralForDsc error:%v", err)
		return err
	} else {
		logrus.Infof("RedeemCollateralForDsc: %v", tx.Hash().String())
	}
	if err := t.waitForTransaction(tx); err != nil {
		logrus.Errorf("waitForTransaction error:%v", err)
		return err
	}
	return nil
}

func (t *Ee) approveD711(nonce uint64, gasFeeCap, gasTipPrice *big.Int, d711Amount *big.Int) error {
	d711Coin, err := contract.NewD711StableCoin(t.D711StableCoinAddress, t.client)
	if err != nil {
		logrus.Errorf("NewDSC error:%v", err)
	}
	auth := t.createTransactor(nonce, gasFeeCap, gasTipPrice)
	err = t.getGasLimit(auth, t.D711StableCoinAddress, contract.D711StableCoinMetaData.ABI, "approve", t.D711EngineAddress, d711Amount)
	if err != nil {
		return err
	}

	tx, err1 := d711Coin.Approve(auth, t.D711EngineAddress, d711Amount)
	if err1 != nil {
		return fmt.Errorf("failed to approveD711 error: %v", err1)
	}
	if err := t.waitForTransaction(tx); err != nil {
		logrus.Errorf("waitForTransaction createTransactor error:%v", err)
		return err
	}
	logrus.Infof("approveD711: %v", tx.Hash().String())
	return nil
}

func (t *Ee) Transfer(from common.Address, to common.Address, sum int64) error {
	nonce, gasFeeCap, gasTipCap, err := t.prepareTransaction(context.Background())
	if err != nil {
		return err
	}
	tx := t.createTransactor(nonce, gasFeeCap, gasTipCap)
	transferFrom, err := t.D711Engine.Transfer(tx, from, to, big.NewInt(sum*1e18))
	if err != nil {
		return err
	}
	err = t.waitForTransaction(transferFrom)
	if err != nil {
		return err
	}
	return nil
}
