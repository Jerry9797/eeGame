// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract_meta

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// D711EngineMetaData contains all meta data concerning the D711Engine contract_metadata.
var D711EngineMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"tokenAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"priceFeedAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"d711Address\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnD711\",\"inputs\":[{\"name\":\"amountD711ToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositCollateral\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositCollateralAndMintD711\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountD711ToMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"forceLiquidate\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAccountCollateralValueD711\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totCollateralInUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAccountInformation\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"totalD711Minted\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"collateralValueInUsd\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAmountWethForMintingD711\",\"inputs\":[{\"name\":\"wethAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountD711ToMint\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getHealthFactor\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTokenAmountFromUsd\",\"inputs\":[{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtToCover\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"tokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getUsdValueD711\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"liquidate\",\"inputs\":[{\"name\":\"collateralToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"debtToCover\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintD711\",\"inputs\":[{\"name\":\"amountD711\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemCollateral\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"redeemCollateralForD711\",\"inputs\":[{\"name\":\"tokenCollateralAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amountCollateral\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amountD711ToBurn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CollateralDeposited\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CollateralRedeemed\",\"inputs\":[{\"name\":\"redeemedFrom\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"redeemedTo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForceLiquidate\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"D711Engine__AmountMustBeMOreThanZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"D711Engine__D711MintFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"D711Engine__HealthFactorNotImproved\",\"inputs\":[{\"name\":\"healthFactorValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"D711Engine__HealthFactorOk\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"D711Engine__NotAllowToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"D711Engine__TheHealthFactorBroken\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"D711Engine__TokenAddressesAndPriceFeedAddressesMustBeSameLength\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"D711Engine__TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
}

// D711EngineABI is the input ABI used to generate the binding from.
// Deprecated: Use D711EngineMetaData.ABI instead.
var D711EngineABI = D711EngineMetaData.ABI

// D711Engine is an auto generated Go binding around an Ethereum contract_metadata.
type D711Engine struct {
	D711EngineCaller     // Read-only binding to the contract_metadata
	D711EngineTransactor // Write-only binding to the contract_metadata
	D711EngineFilterer   // Log filterer for contract_metadata events
}

// D711EngineCaller is an auto generated read-only Go binding around an Ethereum contract_metadata.
type D711EngineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D711EngineTransactor is an auto generated write-only Go binding around an Ethereum contract_metadata.
type D711EngineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D711EngineFilterer is an auto generated log filtering Go binding around an Ethereum contract_metadata events.
type D711EngineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D711EngineSession is an auto generated Go binding around an Ethereum contract_metadata,
// with pre-set call and transact options.
type D711EngineSession struct {
	Contract     *D711Engine       // Generic contract_metadata binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// D711EngineCallerSession is an auto generated read-only Go binding around an Ethereum contract_metadata,
// with pre-set call options.
type D711EngineCallerSession struct {
	Contract *D711EngineCaller // Generic contract_metadata caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// D711EngineTransactorSession is an auto generated write-only Go binding around an Ethereum contract_metadata,
// with pre-set transact options.
type D711EngineTransactorSession struct {
	Contract     *D711EngineTransactor // Generic contract_metadata transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// D711EngineRaw is an auto generated low-level Go binding around an Ethereum contract_metadata.
type D711EngineRaw struct {
	Contract *D711Engine // Generic contract_metadata binding to access the raw methods on
}

// D711EngineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract_metadata.
type D711EngineCallerRaw struct {
	Contract *D711EngineCaller // Generic read-only contract_metadata binding to access the raw methods on
}

// D711EngineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract_metadata.
type D711EngineTransactorRaw struct {
	Contract *D711EngineTransactor // Generic write-only contract_metadata binding to access the raw methods on
}

// NewD711Engine creates a new instance of D711Engine, bound to a specific deployed contract_metadata.
func NewD711Engine(address common.Address, backend bind.ContractBackend) (*D711Engine, error) {
	contract, err := bindD711Engine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &D711Engine{D711EngineCaller: D711EngineCaller{contract: contract}, D711EngineTransactor: D711EngineTransactor{contract: contract}, D711EngineFilterer: D711EngineFilterer{contract: contract}}, nil
}

// NewD711EngineCaller creates a new read-only instance of D711Engine, bound to a specific deployed contract_metadata.
func NewD711EngineCaller(address common.Address, caller bind.ContractCaller) (*D711EngineCaller, error) {
	contract, err := bindD711Engine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &D711EngineCaller{contract: contract}, nil
}

// NewD711EngineTransactor creates a new write-only instance of D711Engine, bound to a specific deployed contract_metadata.
func NewD711EngineTransactor(address common.Address, transactor bind.ContractTransactor) (*D711EngineTransactor, error) {
	contract, err := bindD711Engine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &D711EngineTransactor{contract: contract}, nil
}

// NewD711EngineFilterer creates a new log filterer instance of D711Engine, bound to a specific deployed contract_metadata.
func NewD711EngineFilterer(address common.Address, filterer bind.ContractFilterer) (*D711EngineFilterer, error) {
	contract, err := bindD711Engine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &D711EngineFilterer{contract: contract}, nil
}

// bindD711Engine binds a generic wrapper to an already deployed contract_metadata.
func bindD711Engine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := D711EngineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract_metadata method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D711Engine *D711EngineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D711Engine.Contract.D711EngineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract_metadata, calling
// its default method if one is available.
func (_D711Engine *D711EngineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D711Engine.Contract.D711EngineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract_metadata method with params as input values.
func (_D711Engine *D711EngineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D711Engine.Contract.D711EngineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract_metadata method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D711Engine *D711EngineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D711Engine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract_metadata, calling
// its default method if one is available.
func (_D711Engine *D711EngineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D711Engine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract_metadata method with params as input values.
func (_D711Engine *D711EngineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D711Engine.Contract.contract.Transact(opts, method, params...)
}

// GetAccountCollateralValueD711 is a free data retrieval call binding the contract_metadata method 0x09906436.
//
// Solidity: function getAccountCollateralValueD711(address user) view returns(uint256 totCollateralInUsd)
func (_D711Engine *D711EngineCaller) GetAccountCollateralValueD711(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D711Engine.contract.Call(opts, &out, "getAccountCollateralValueD711", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAccountCollateralValueD711 is a free data retrieval call binding the contract_metadata method 0x09906436.
//
// Solidity: function getAccountCollateralValueD711(address user) view returns(uint256 totCollateralInUsd)
func (_D711Engine *D711EngineSession) GetAccountCollateralValueD711(user common.Address) (*big.Int, error) {
	return _D711Engine.Contract.GetAccountCollateralValueD711(&_D711Engine.CallOpts, user)
}

// GetAccountCollateralValueD711 is a free data retrieval call binding the contract_metadata method 0x09906436.
//
// Solidity: function getAccountCollateralValueD711(address user) view returns(uint256 totCollateralInUsd)
func (_D711Engine *D711EngineCallerSession) GetAccountCollateralValueD711(user common.Address) (*big.Int, error) {
	return _D711Engine.Contract.GetAccountCollateralValueD711(&_D711Engine.CallOpts, user)
}

// GetAccountInformation is a free data retrieval call binding the contract_metadata method 0x7be564fc.
//
// Solidity: function getAccountInformation(address user) view returns(uint256 totalD711Minted, uint256 collateralValueInUsd)
func (_D711Engine *D711EngineCaller) GetAccountInformation(opts *bind.CallOpts, user common.Address) (struct {
	TotalD711Minted      *big.Int
	CollateralValueInUsd *big.Int
}, error) {
	var out []interface{}
	err := _D711Engine.contract.Call(opts, &out, "getAccountInformation", user)

	outstruct := new(struct {
		TotalD711Minted      *big.Int
		CollateralValueInUsd *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalD711Minted = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CollateralValueInUsd = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAccountInformation is a free data retrieval call binding the contract_metadata method 0x7be564fc.
//
// Solidity: function getAccountInformation(address user) view returns(uint256 totalD711Minted, uint256 collateralValueInUsd)
func (_D711Engine *D711EngineSession) GetAccountInformation(user common.Address) (struct {
	TotalD711Minted      *big.Int
	CollateralValueInUsd *big.Int
}, error) {
	return _D711Engine.Contract.GetAccountInformation(&_D711Engine.CallOpts, user)
}

// GetAccountInformation is a free data retrieval call binding the contract_metadata method 0x7be564fc.
//
// Solidity: function getAccountInformation(address user) view returns(uint256 totalD711Minted, uint256 collateralValueInUsd)
func (_D711Engine *D711EngineCallerSession) GetAccountInformation(user common.Address) (struct {
	TotalD711Minted      *big.Int
	CollateralValueInUsd *big.Int
}, error) {
	return _D711Engine.Contract.GetAccountInformation(&_D711Engine.CallOpts, user)
}

// GetAmountWethForMintingD711 is a free data retrieval call binding the contract_metadata method 0x942eef38.
//
// Solidity: function getAmountWethForMintingD711(address wethAddress, uint256 amountD711ToMint) view returns(uint256)
func (_D711Engine *D711EngineCaller) GetAmountWethForMintingD711(opts *bind.CallOpts, wethAddress common.Address, amountD711ToMint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _D711Engine.contract.Call(opts, &out, "getAmountWethForMintingD711", wethAddress, amountD711ToMint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountWethForMintingD711 is a free data retrieval call binding the contract_metadata method 0x942eef38.
//
// Solidity: function getAmountWethForMintingD711(address wethAddress, uint256 amountD711ToMint) view returns(uint256)
func (_D711Engine *D711EngineSession) GetAmountWethForMintingD711(wethAddress common.Address, amountD711ToMint *big.Int) (*big.Int, error) {
	return _D711Engine.Contract.GetAmountWethForMintingD711(&_D711Engine.CallOpts, wethAddress, amountD711ToMint)
}

// GetAmountWethForMintingD711 is a free data retrieval call binding the contract_metadata method 0x942eef38.
//
// Solidity: function getAmountWethForMintingD711(address wethAddress, uint256 amountD711ToMint) view returns(uint256)
func (_D711Engine *D711EngineCallerSession) GetAmountWethForMintingD711(wethAddress common.Address, amountD711ToMint *big.Int) (*big.Int, error) {
	return _D711Engine.Contract.GetAmountWethForMintingD711(&_D711Engine.CallOpts, wethAddress, amountD711ToMint)
}

// GetHealthFactor is a free data retrieval call binding the contract_metadata method 0xfe6bcd7c.
//
// Solidity: function getHealthFactor(address user) view returns(uint256)
func (_D711Engine *D711EngineCaller) GetHealthFactor(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D711Engine.contract.Call(opts, &out, "getHealthFactor", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetHealthFactor is a free data retrieval call binding the contract_metadata method 0xfe6bcd7c.
//
// Solidity: function getHealthFactor(address user) view returns(uint256)
func (_D711Engine *D711EngineSession) GetHealthFactor(user common.Address) (*big.Int, error) {
	return _D711Engine.Contract.GetHealthFactor(&_D711Engine.CallOpts, user)
}

// GetHealthFactor is a free data retrieval call binding the contract_metadata method 0xfe6bcd7c.
//
// Solidity: function getHealthFactor(address user) view returns(uint256)
func (_D711Engine *D711EngineCallerSession) GetHealthFactor(user common.Address) (*big.Int, error) {
	return _D711Engine.Contract.GetHealthFactor(&_D711Engine.CallOpts, user)
}

// GetTokenAmountFromUsd is a free data retrieval call binding the contract_metadata method 0xafea2e48.
//
// Solidity: function getTokenAmountFromUsd(address collateralToken, uint256 debtToCover) view returns(uint256 tokenAmount)
func (_D711Engine *D711EngineCaller) GetTokenAmountFromUsd(opts *bind.CallOpts, collateralToken common.Address, debtToCover *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _D711Engine.contract.Call(opts, &out, "getTokenAmountFromUsd", collateralToken, debtToCover)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenAmountFromUsd is a free data retrieval call binding the contract_metadata method 0xafea2e48.
//
// Solidity: function getTokenAmountFromUsd(address collateralToken, uint256 debtToCover) view returns(uint256 tokenAmount)
func (_D711Engine *D711EngineSession) GetTokenAmountFromUsd(collateralToken common.Address, debtToCover *big.Int) (*big.Int, error) {
	return _D711Engine.Contract.GetTokenAmountFromUsd(&_D711Engine.CallOpts, collateralToken, debtToCover)
}

// GetTokenAmountFromUsd is a free data retrieval call binding the contract_metadata method 0xafea2e48.
//
// Solidity: function getTokenAmountFromUsd(address collateralToken, uint256 debtToCover) view returns(uint256 tokenAmount)
func (_D711Engine *D711EngineCallerSession) GetTokenAmountFromUsd(collateralToken common.Address, debtToCover *big.Int) (*big.Int, error) {
	return _D711Engine.Contract.GetTokenAmountFromUsd(&_D711Engine.CallOpts, collateralToken, debtToCover)
}

// GetUsdValueD711 is a free data retrieval call binding the contract_metadata method 0x7c066783.
//
// Solidity: function getUsdValueD711(address token, uint256 amount) view returns(uint256)
func (_D711Engine *D711EngineCaller) GetUsdValueD711(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _D711Engine.contract.Call(opts, &out, "getUsdValueD711", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUsdValueD711 is a free data retrieval call binding the contract_metadata method 0x7c066783.
//
// Solidity: function getUsdValueD711(address token, uint256 amount) view returns(uint256)
func (_D711Engine *D711EngineSession) GetUsdValueD711(token common.Address, amount *big.Int) (*big.Int, error) {
	return _D711Engine.Contract.GetUsdValueD711(&_D711Engine.CallOpts, token, amount)
}

// GetUsdValueD711 is a free data retrieval call binding the contract_metadata method 0x7c066783.
//
// Solidity: function getUsdValueD711(address token, uint256 amount) view returns(uint256)
func (_D711Engine *D711EngineCallerSession) GetUsdValueD711(token common.Address, amount *big.Int) (*big.Int, error) {
	return _D711Engine.Contract.GetUsdValueD711(&_D711Engine.CallOpts, token, amount)
}

// BurnD711 is a paid mutator transaction binding the contract_metadata method 0x30ea182c.
//
// Solidity: function burnD711(uint256 amountD711ToBurn) returns()
func (_D711Engine *D711EngineTransactor) BurnD711(opts *bind.TransactOpts, amountD711ToBurn *big.Int) (*types.Transaction, error) {
	return _D711Engine.contract.Transact(opts, "burnD711", amountD711ToBurn)
}

// BurnD711 is a paid mutator transaction binding the contract_metadata method 0x30ea182c.
//
// Solidity: function burnD711(uint256 amountD711ToBurn) returns()
func (_D711Engine *D711EngineSession) BurnD711(amountD711ToBurn *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.BurnD711(&_D711Engine.TransactOpts, amountD711ToBurn)
}

// BurnD711 is a paid mutator transaction binding the contract_metadata method 0x30ea182c.
//
// Solidity: function burnD711(uint256 amountD711ToBurn) returns()
func (_D711Engine *D711EngineTransactorSession) BurnD711(amountD711ToBurn *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.BurnD711(&_D711Engine.TransactOpts, amountD711ToBurn)
}

// DepositCollateral is a paid mutator transaction binding the contract_metadata method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_D711Engine *D711EngineTransactor) DepositCollateral(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _D711Engine.contract.Transact(opts, "depositCollateral", tokenCollateralAddress, amountCollateral)
}

// DepositCollateral is a paid mutator transaction binding the contract_metadata method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_D711Engine *D711EngineSession) DepositCollateral(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.DepositCollateral(&_D711Engine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// DepositCollateral is a paid mutator transaction binding the contract_metadata method 0xa5d5db0c.
//
// Solidity: function depositCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_D711Engine *D711EngineTransactorSession) DepositCollateral(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.DepositCollateral(&_D711Engine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// DepositCollateralAndMintD711 is a paid mutator transaction binding the contract_metadata method 0x66db9faa.
//
// Solidity: function depositCollateralAndMintD711(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountD711ToMint) returns()
func (_D711Engine *D711EngineTransactor) DepositCollateralAndMintD711(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int, amountD711ToMint *big.Int) (*types.Transaction, error) {
	return _D711Engine.contract.Transact(opts, "depositCollateralAndMintD711", tokenCollateralAddress, amountCollateral, amountD711ToMint)
}

// DepositCollateralAndMintD711 is a paid mutator transaction binding the contract_metadata method 0x66db9faa.
//
// Solidity: function depositCollateralAndMintD711(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountD711ToMint) returns()
func (_D711Engine *D711EngineSession) DepositCollateralAndMintD711(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountD711ToMint *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.DepositCollateralAndMintD711(&_D711Engine.TransactOpts, tokenCollateralAddress, amountCollateral, amountD711ToMint)
}

// DepositCollateralAndMintD711 is a paid mutator transaction binding the contract_metadata method 0x66db9faa.
//
// Solidity: function depositCollateralAndMintD711(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountD711ToMint) returns()
func (_D711Engine *D711EngineTransactorSession) DepositCollateralAndMintD711(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountD711ToMint *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.DepositCollateralAndMintD711(&_D711Engine.TransactOpts, tokenCollateralAddress, amountCollateral, amountD711ToMint)
}

// ForceLiquidate is a paid mutator transaction binding the contract_metadata method 0xc85e8e74.
//
// Solidity: function forceLiquidate() returns()
func (_D711Engine *D711EngineTransactor) ForceLiquidate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D711Engine.contract.Transact(opts, "forceLiquidate")
}

// ForceLiquidate is a paid mutator transaction binding the contract_metadata method 0xc85e8e74.
//
// Solidity: function forceLiquidate() returns()
func (_D711Engine *D711EngineSession) ForceLiquidate() (*types.Transaction, error) {
	return _D711Engine.Contract.ForceLiquidate(&_D711Engine.TransactOpts)
}

// ForceLiquidate is a paid mutator transaction binding the contract_metadata method 0xc85e8e74.
//
// Solidity: function forceLiquidate() returns()
func (_D711Engine *D711EngineTransactorSession) ForceLiquidate() (*types.Transaction, error) {
	return _D711Engine.Contract.ForceLiquidate(&_D711Engine.TransactOpts)
}

// Liquidate is a paid mutator transaction binding the contract_metadata method 0x26c01303.
//
// Solidity: function liquidate(address collateralToken, address user, uint256 debtToCover) returns()
func (_D711Engine *D711EngineTransactor) Liquidate(opts *bind.TransactOpts, collateralToken common.Address, user common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _D711Engine.contract.Transact(opts, "liquidate", collateralToken, user, debtToCover)
}

// Liquidate is a paid mutator transaction binding the contract_metadata method 0x26c01303.
//
// Solidity: function liquidate(address collateralToken, address user, uint256 debtToCover) returns()
func (_D711Engine *D711EngineSession) Liquidate(collateralToken common.Address, user common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.Liquidate(&_D711Engine.TransactOpts, collateralToken, user, debtToCover)
}

// Liquidate is a paid mutator transaction binding the contract_metadata method 0x26c01303.
//
// Solidity: function liquidate(address collateralToken, address user, uint256 debtToCover) returns()
func (_D711Engine *D711EngineTransactorSession) Liquidate(collateralToken common.Address, user common.Address, debtToCover *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.Liquidate(&_D711Engine.TransactOpts, collateralToken, user, debtToCover)
}

// MintD711 is a paid mutator transaction binding the contract_metadata method 0xec8b0aa9.
//
// Solidity: function mintD711(uint256 amountD711) returns()
func (_D711Engine *D711EngineTransactor) MintD711(opts *bind.TransactOpts, amountD711 *big.Int) (*types.Transaction, error) {
	return _D711Engine.contract.Transact(opts, "mintD711", amountD711)
}

// MintD711 is a paid mutator transaction binding the contract_metadata method 0xec8b0aa9.
//
// Solidity: function mintD711(uint256 amountD711) returns()
func (_D711Engine *D711EngineSession) MintD711(amountD711 *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.MintD711(&_D711Engine.TransactOpts, amountD711)
}

// MintD711 is a paid mutator transaction binding the contract_metadata method 0xec8b0aa9.
//
// Solidity: function mintD711(uint256 amountD711) returns()
func (_D711Engine *D711EngineTransactorSession) MintD711(amountD711 *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.MintD711(&_D711Engine.TransactOpts, amountD711)
}

// RedeemCollateral is a paid mutator transaction binding the contract_metadata method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_D711Engine *D711EngineTransactor) RedeemCollateral(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _D711Engine.contract.Transact(opts, "redeemCollateral", tokenCollateralAddress, amountCollateral)
}

// RedeemCollateral is a paid mutator transaction binding the contract_metadata method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_D711Engine *D711EngineSession) RedeemCollateral(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.RedeemCollateral(&_D711Engine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// RedeemCollateral is a paid mutator transaction binding the contract_metadata method 0x9acd81b3.
//
// Solidity: function redeemCollateral(address tokenCollateralAddress, uint256 amountCollateral) returns()
func (_D711Engine *D711EngineTransactorSession) RedeemCollateral(tokenCollateralAddress common.Address, amountCollateral *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.RedeemCollateral(&_D711Engine.TransactOpts, tokenCollateralAddress, amountCollateral)
}

// RedeemCollateralForD711 is a paid mutator transaction binding the contract_metadata method 0x117e29c7.
//
// Solidity: function redeemCollateralForD711(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountD711ToBurn) returns()
func (_D711Engine *D711EngineTransactor) RedeemCollateralForD711(opts *bind.TransactOpts, tokenCollateralAddress common.Address, amountCollateral *big.Int, amountD711ToBurn *big.Int) (*types.Transaction, error) {
	return _D711Engine.contract.Transact(opts, "redeemCollateralForD711", tokenCollateralAddress, amountCollateral, amountD711ToBurn)
}

// RedeemCollateralForD711 is a paid mutator transaction binding the contract_metadata method 0x117e29c7.
//
// Solidity: function redeemCollateralForD711(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountD711ToBurn) returns()
func (_D711Engine *D711EngineSession) RedeemCollateralForD711(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountD711ToBurn *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.RedeemCollateralForD711(&_D711Engine.TransactOpts, tokenCollateralAddress, amountCollateral, amountD711ToBurn)
}

// RedeemCollateralForD711 is a paid mutator transaction binding the contract_metadata method 0x117e29c7.
//
// Solidity: function redeemCollateralForD711(address tokenCollateralAddress, uint256 amountCollateral, uint256 amountD711ToBurn) returns()
func (_D711Engine *D711EngineTransactorSession) RedeemCollateralForD711(tokenCollateralAddress common.Address, amountCollateral *big.Int, amountD711ToBurn *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.RedeemCollateralForD711(&_D711Engine.TransactOpts, tokenCollateralAddress, amountCollateral, amountD711ToBurn)
}

// Transfer is a paid mutator transaction binding the contract_metadata method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns()
func (_D711Engine *D711EngineTransactor) Transfer(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D711Engine.contract.Transact(opts, "transfer", from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract_metadata method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns()
func (_D711Engine *D711EngineSession) Transfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.Transfer(&_D711Engine.TransactOpts, from, to, amount)
}

// Transfer is a paid mutator transaction binding the contract_metadata method 0xbeabacc8.
//
// Solidity: function transfer(address from, address to, uint256 amount) returns()
func (_D711Engine *D711EngineTransactorSession) Transfer(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _D711Engine.Contract.Transfer(&_D711Engine.TransactOpts, from, to, amount)
}

// D711EngineCollateralDepositedIterator is returned from FilterCollateralDeposited and is used to iterate over the raw logs and unpacked data for CollateralDeposited events raised by the D711Engine contract_metadata.
type D711EngineCollateralDepositedIterator struct {
	Event *D711EngineCollateralDeposited // Event containing the contract_metadata specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D711EngineCollateralDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D711EngineCollateralDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D711EngineCollateralDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D711EngineCollateralDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D711EngineCollateralDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D711EngineCollateralDeposited represents a CollateralDeposited event raised by the D711Engine contract_metadata.
type D711EngineCollateralDeposited struct {
	User   common.Address
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCollateralDeposited is a free log retrieval operation binding the contract_metadata event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed user, address indexed token, uint256 indexed amount)
func (_D711Engine *D711EngineFilterer) FilterCollateralDeposited(opts *bind.FilterOpts, user []common.Address, token []common.Address, amount []*big.Int) (*D711EngineCollateralDepositedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _D711Engine.contract.FilterLogs(opts, "CollateralDeposited", userRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return &D711EngineCollateralDepositedIterator{contract: _D711Engine.contract, event: "CollateralDeposited", logs: logs, sub: sub}, nil
}

// WatchCollateralDeposited is a free log subscription operation binding the contract_metadata event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed user, address indexed token, uint256 indexed amount)
func (_D711Engine *D711EngineFilterer) WatchCollateralDeposited(opts *bind.WatchOpts, sink chan<- *D711EngineCollateralDeposited, user []common.Address, token []common.Address, amount []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _D711Engine.contract.WatchLogs(opts, "CollateralDeposited", userRule, tokenRule, amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D711EngineCollateralDeposited)
				if err := _D711Engine.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollateralDeposited is a log parse operation binding the contract_metadata event 0xf1c0dd7e9b98bbff859029005ef89b127af049cd18df1a8d79f0b7e019911e56.
//
// Solidity: event CollateralDeposited(address indexed user, address indexed token, uint256 indexed amount)
func (_D711Engine *D711EngineFilterer) ParseCollateralDeposited(log types.Log) (*D711EngineCollateralDeposited, error) {
	event := new(D711EngineCollateralDeposited)
	if err := _D711Engine.contract.UnpackLog(event, "CollateralDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D711EngineCollateralRedeemedIterator is returned from FilterCollateralRedeemed and is used to iterate over the raw logs and unpacked data for CollateralRedeemed events raised by the D711Engine contract_metadata.
type D711EngineCollateralRedeemedIterator struct {
	Event *D711EngineCollateralRedeemed // Event containing the contract_metadata specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D711EngineCollateralRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D711EngineCollateralRedeemed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D711EngineCollateralRedeemed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D711EngineCollateralRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D711EngineCollateralRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D711EngineCollateralRedeemed represents a CollateralRedeemed event raised by the D711Engine contract_metadata.
type D711EngineCollateralRedeemed struct {
	RedeemedFrom common.Address
	RedeemedTo   common.Address
	Token        common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCollateralRedeemed is a free log retrieval operation binding the contract_metadata event 0xb292db12b8dfa23b760ee5e3bd1d3be184cd8f0eeeacd27ce120d4de4e4721f6.
//
// Solidity: event CollateralRedeemed(address indexed redeemedFrom, address indexed redeemedTo, address indexed token, uint256 amount)
func (_D711Engine *D711EngineFilterer) FilterCollateralRedeemed(opts *bind.FilterOpts, redeemedFrom []common.Address, redeemedTo []common.Address, token []common.Address) (*D711EngineCollateralRedeemedIterator, error) {

	var redeemedFromRule []interface{}
	for _, redeemedFromItem := range redeemedFrom {
		redeemedFromRule = append(redeemedFromRule, redeemedFromItem)
	}
	var redeemedToRule []interface{}
	for _, redeemedToItem := range redeemedTo {
		redeemedToRule = append(redeemedToRule, redeemedToItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _D711Engine.contract.FilterLogs(opts, "CollateralRedeemed", redeemedFromRule, redeemedToRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &D711EngineCollateralRedeemedIterator{contract: _D711Engine.contract, event: "CollateralRedeemed", logs: logs, sub: sub}, nil
}

// WatchCollateralRedeemed is a free log subscription operation binding the contract_metadata event 0xb292db12b8dfa23b760ee5e3bd1d3be184cd8f0eeeacd27ce120d4de4e4721f6.
//
// Solidity: event CollateralRedeemed(address indexed redeemedFrom, address indexed redeemedTo, address indexed token, uint256 amount)
func (_D711Engine *D711EngineFilterer) WatchCollateralRedeemed(opts *bind.WatchOpts, sink chan<- *D711EngineCollateralRedeemed, redeemedFrom []common.Address, redeemedTo []common.Address, token []common.Address) (event.Subscription, error) {

	var redeemedFromRule []interface{}
	for _, redeemedFromItem := range redeemedFrom {
		redeemedFromRule = append(redeemedFromRule, redeemedFromItem)
	}
	var redeemedToRule []interface{}
	for _, redeemedToItem := range redeemedTo {
		redeemedToRule = append(redeemedToRule, redeemedToItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _D711Engine.contract.WatchLogs(opts, "CollateralRedeemed", redeemedFromRule, redeemedToRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D711EngineCollateralRedeemed)
				if err := _D711Engine.contract.UnpackLog(event, "CollateralRedeemed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCollateralRedeemed is a log parse operation binding the contract_metadata event 0xb292db12b8dfa23b760ee5e3bd1d3be184cd8f0eeeacd27ce120d4de4e4721f6.
//
// Solidity: event CollateralRedeemed(address indexed redeemedFrom, address indexed redeemedTo, address indexed token, uint256 amount)
func (_D711Engine *D711EngineFilterer) ParseCollateralRedeemed(log types.Log) (*D711EngineCollateralRedeemed, error) {
	event := new(D711EngineCollateralRedeemed)
	if err := _D711Engine.contract.UnpackLog(event, "CollateralRedeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D711EngineForceLiquidateIterator is returned from FilterForceLiquidate and is used to iterate over the raw logs and unpacked data for ForceLiquidate events raised by the D711Engine contract_metadata.
type D711EngineForceLiquidateIterator struct {
	Event *D711EngineForceLiquidate // Event containing the contract_metadata specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *D711EngineForceLiquidateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D711EngineForceLiquidate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(D711EngineForceLiquidate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *D711EngineForceLiquidateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D711EngineForceLiquidateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D711EngineForceLiquidate represents a ForceLiquidate event raised by the D711Engine contract_metadata.
type D711EngineForceLiquidate struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterForceLiquidate is a free log retrieval operation binding the contract_metadata event 0x000da657026ff11be43c632fa17885aa1b35fc0465b724cd4e7ba0151b7a6052.
//
// Solidity: event ForceLiquidate(address indexed user)
func (_D711Engine *D711EngineFilterer) FilterForceLiquidate(opts *bind.FilterOpts, user []common.Address) (*D711EngineForceLiquidateIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _D711Engine.contract.FilterLogs(opts, "ForceLiquidate", userRule)
	if err != nil {
		return nil, err
	}
	return &D711EngineForceLiquidateIterator{contract: _D711Engine.contract, event: "ForceLiquidate", logs: logs, sub: sub}, nil
}

// WatchForceLiquidate is a free log subscription operation binding the contract_metadata event 0x000da657026ff11be43c632fa17885aa1b35fc0465b724cd4e7ba0151b7a6052.
//
// Solidity: event ForceLiquidate(address indexed user)
func (_D711Engine *D711EngineFilterer) WatchForceLiquidate(opts *bind.WatchOpts, sink chan<- *D711EngineForceLiquidate, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _D711Engine.contract.WatchLogs(opts, "ForceLiquidate", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D711EngineForceLiquidate)
				if err := _D711Engine.contract.UnpackLog(event, "ForceLiquidate", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForceLiquidate is a log parse operation binding the contract_metadata event 0x000da657026ff11be43c632fa17885aa1b35fc0465b724cd4e7ba0151b7a6052.
//
// Solidity: event ForceLiquidate(address indexed user)
func (_D711Engine *D711EngineFilterer) ParseForceLiquidate(log types.Log) (*D711EngineForceLiquidate, error) {
	event := new(D711EngineForceLiquidate)
	if err := _D711Engine.contract.UnpackLog(event, "ForceLiquidate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
