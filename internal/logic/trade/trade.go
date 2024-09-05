package trade

import (
	"context"
	"crypto/ecdsa"
	"eeGame/internal/logic/trade/contract_meta"
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
	"math/big"
	"strings"
)

type Trade struct {
	client                *ethclient.Client
	gameManger            *contract_meta.GameManagement
	chainId               *big.Int
	privateKey            *ecdsa.PrivateKey
	publicKey             common.Address
	gameManagementAddress common.Address
}

func NewTrade(ctx context.Context) (*Trade, error) {
	cfg := g.Cfg().MustGet(ctx, "ETH").Map()
	client := response.EthClient

	privateKey, err := crypto.HexToECDSA(cfg["privateKey"].(string))
	if err != nil {
		return nil, fmt.Errorf("privateKey error : %v", err)
	}
	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	publicKeyAddress := crypto.PubkeyToAddress(*publicKey)

	gameManagementAddressStr := cfg["gameManagement"]
	gameManagementAddress := common.HexToAddress(gameManagementAddressStr.(string))

	eeStableCoinStr := cfg["d711StableCoinAddress"]
	eeStableCoinAddress := common.HexToAddress(eeStableCoinStr.(string))
	gameManger, err := contract_meta.NewGameManagement(eeStableCoinAddress, client)
	if err != nil {
		return nil, err
	}
	return &Trade{
		client:                client,
		gameManger:            gameManger,
		chainId:               big.NewInt(11155111),
		publicKey:             publicKeyAddress,
		privateKey:            privateKey,
		gameManagementAddress: gameManagementAddress,
	}, nil
}

func (t *Trade) StartGame(ctx context.Context, roomId *big.Int, userAddress []common.Address, minQualifiedAmount *big.Int) error {
	nonce, tipCap, feeCap, err := t.prepareTransaction(ctx)
	if err != nil {
		return err
	}
	auth, err := t.createTransactor(nonce, feeCap, tipCap)
	if err != nil {
		return err
	}
	err = t.gatGasLimit(auth, t.gameManagementAddress, contract_meta.GameManagementMetaData.ABI, "startGame", roomId, userAddress, minQualifiedAmount)
	if err != nil {
		return err
	}
	tx, err := t.gameManger.StartGame(auth, roomId, userAddress, minQualifiedAmount)
	err = t.waitForTransaction(tx)
	if err != nil {
		return err
	}
	return nil
}

func (t *Trade) EndGame(ctx context.Context, roomId *big.Int, winner common.Address, userAddress []common.Address, betAmounts []*big.Int) error {
	nonce, tipCap, feeCap, err := t.prepareTransaction(ctx)
	if err != nil {
		return err
	}
	auth, err := t.createTransactor(nonce, feeCap, tipCap)
	if err != nil {
		return err
	}
	err = t.gatGasLimit(auth, t.gameManagementAddress, contract_meta.GameManagementMetaData.ABI, "endGame", roomId, winner, userAddress, betAmounts)
	if err != nil {
		return err
	}
	tx, err := t.gameManger.EndGame(auth, roomId, winner, userAddress, betAmounts)
	if err != nil {
		return err
	}
	err = t.waitForTransaction(tx)
	if err != nil {
		return err
	}
	return nil
}

func (t *Trade) TopUp(ctx context.Context, newInt *big.Int) error {
	return nil
}

func (t *Trade) Withdraw(ctx context.Context, newInt *big.Int) error {
	return nil
}

func (t *Trade) prepareTransaction(ctx context.Context) (nonce uint64, gasTipCap, gasFeeCap *big.Int, err error) {
	nonce, err = t.client.PendingNonceAt(ctx, t.publicKey)
	if err != nil {
		return
	}
	block, err := t.client.BlockByNumber(ctx, nil)
	if err != nil {
		return
	}
	gasTipCap, err = t.client.SuggestGasTipCap(ctx)
	if err != nil {
		return
	}
	gasFeeCap = new(big.Int).Add(block.BaseFee(), gasTipCap)
	return
}

func (t *Trade) createTransactor(nonce uint64, gasFeeCap, gasTipCap *big.Int) (*bind.TransactOpts, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(t.privateKey, t.chainId)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasFeeCap = gasFeeCap
	auth.GasTipCap = gasTipCap
	return auth, nil
}

func (t *Trade) waitForTransaction(tx *types.Transaction) error {
	mined, err := bind.WaitMined(context.Background(), t.client, tx)
	if err != nil {
		return err
	}
	if mined.Status == types.ReceiptStatusFailed {
		return fmt.Errorf("transaction failed: %s", tx.Hash().Hex())
	}
	fmt.Printf("Transaction succeeded: %s\n", tx.Hash().Hex())
	return nil
}

func (t *Trade) gatGasLimit(auth *bind.TransactOpts, to common.Address, contrastAbi string, methodName string, args ...interface{}) error {
	parsed, err := abi.JSON(strings.NewReader(contrastAbi))
	if err != nil {
		return err
	}
	encodedData, err := parsed.Pack(methodName, args...)
	if err != nil {
		return err
	}
	callMsg := ethereum.CallMsg{
		From:      t.publicKey,
		To:        &to,
		Data:      encodedData,
		GasFeeCap: auth.GasFeeCap,
		GasTipCap: auth.GasTipCap,
	}
	estimateGas, err := t.client.EstimateGas(context.Background(), callMsg)
	if err != nil {
		return err
	}
	auth.GasLimit = estimateGas
	return nil
}
