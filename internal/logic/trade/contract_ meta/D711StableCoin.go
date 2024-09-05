// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract__meta

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

// D711StableCoinMetaData contains all meta data concerning the D711StableCoin contract_metadata.
var D711StableCoinMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"_from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"D711StableCoin__BurnAmountExceedsBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"D711StableCoin__CallIsNotD711Engine\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"D711StableCoin__MustBeMoreThanZero\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"D711StableCoin__NotZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// D711StableCoinABI is the input ABI used to generate the binding from.
// Deprecated: Use D711StableCoinMetaData.ABI instead.
var D711StableCoinABI = D711StableCoinMetaData.ABI

// D711StableCoin is an auto generated Go binding around an Ethereum contract_metadata.
type D711StableCoin struct {
	D711StableCoinCaller     // Read-only binding to the contract_metadata
	D711StableCoinTransactor // Write-only binding to the contract_metadata
	D711StableCoinFilterer   // Log filterer for contract_metadata events
}

// D711StableCoinCaller is an auto generated read-only Go binding around an Ethereum contract_metadata.
type D711StableCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D711StableCoinTransactor is an auto generated write-only Go binding around an Ethereum contract_metadata.
type D711StableCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D711StableCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract_metadata events.
type D711StableCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// D711StableCoinSession is an auto generated Go binding around an Ethereum contract_metadata,
// with pre-set call and transact options.
type D711StableCoinSession struct {
	Contract     *D711StableCoin   // Generic contract_metadata binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// D711StableCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract_metadata,
// with pre-set call options.
type D711StableCoinCallerSession struct {
	Contract *D711StableCoinCaller // Generic contract_metadata caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// D711StableCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract_metadata,
// with pre-set transact options.
type D711StableCoinTransactorSession struct {
	Contract     *D711StableCoinTransactor // Generic contract_metadata transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// D711StableCoinRaw is an auto generated low-level Go binding around an Ethereum contract_metadata.
type D711StableCoinRaw struct {
	Contract *D711StableCoin // Generic contract_metadata binding to access the raw methods on
}

// D711StableCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract_metadata.
type D711StableCoinCallerRaw struct {
	Contract *D711StableCoinCaller // Generic read-only contract_metadata binding to access the raw methods on
}

// D711StableCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract_metadata.
type D711StableCoinTransactorRaw struct {
	Contract *D711StableCoinTransactor // Generic write-only contract_metadata binding to access the raw methods on
}

// NewD711StableCoin creates a new instance of D711StableCoin, bound to a specific deployed contract_metadata.
func NewD711StableCoin(address common.Address, backend bind.ContractBackend) (*D711StableCoin, error) {
	contract, err := bindD711StableCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &D711StableCoin{D711StableCoinCaller: D711StableCoinCaller{contract: contract}, D711StableCoinTransactor: D711StableCoinTransactor{contract: contract}, D711StableCoinFilterer: D711StableCoinFilterer{contract: contract}}, nil
}

// NewD711StableCoinCaller creates a new read-only instance of D711StableCoin, bound to a specific deployed contract_metadata.
func NewD711StableCoinCaller(address common.Address, caller bind.ContractCaller) (*D711StableCoinCaller, error) {
	contract, err := bindD711StableCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &D711StableCoinCaller{contract: contract}, nil
}

// NewD711StableCoinTransactor creates a new write-only instance of D711StableCoin, bound to a specific deployed contract_metadata.
func NewD711StableCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*D711StableCoinTransactor, error) {
	contract, err := bindD711StableCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &D711StableCoinTransactor{contract: contract}, nil
}

// NewD711StableCoinFilterer creates a new log filterer instance of D711StableCoin, bound to a specific deployed contract_metadata.
func NewD711StableCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*D711StableCoinFilterer, error) {
	contract, err := bindD711StableCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &D711StableCoinFilterer{contract: contract}, nil
}

// bindD711StableCoin binds a generic wrapper to an already deployed contract_metadata.
func bindD711StableCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := D711StableCoinMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract_metadata method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D711StableCoin *D711StableCoinRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D711StableCoin.Contract.D711StableCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract_metadata, calling
// its default method if one is available.
func (_D711StableCoin *D711StableCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D711StableCoin.Contract.D711StableCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract_metadata method with params as input values.
func (_D711StableCoin *D711StableCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D711StableCoin.Contract.D711StableCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract_metadata method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_D711StableCoin *D711StableCoinCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _D711StableCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract_metadata, calling
// its default method if one is available.
func (_D711StableCoin *D711StableCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D711StableCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract_metadata method with params as input values.
func (_D711StableCoin *D711StableCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _D711StableCoin.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract_metadata method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_D711StableCoin *D711StableCoinCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D711StableCoin.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract_metadata method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_D711StableCoin *D711StableCoinSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _D711StableCoin.Contract.Allowance(&_D711StableCoin.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract_metadata method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_D711StableCoin *D711StableCoinCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _D711StableCoin.Contract.Allowance(&_D711StableCoin.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract_metadata method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_D711StableCoin *D711StableCoinCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _D711StableCoin.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract_metadata method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_D711StableCoin *D711StableCoinSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _D711StableCoin.Contract.BalanceOf(&_D711StableCoin.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract_metadata method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_D711StableCoin *D711StableCoinCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _D711StableCoin.Contract.BalanceOf(&_D711StableCoin.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract_metadata method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_D711StableCoin *D711StableCoinCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _D711StableCoin.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract_metadata method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_D711StableCoin *D711StableCoinSession) Decimals() (uint8, error) {
	return _D711StableCoin.Contract.Decimals(&_D711StableCoin.CallOpts)
}

// Decimals is a free data retrieval call binding the contract_metadata method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_D711StableCoin *D711StableCoinCallerSession) Decimals() (uint8, error) {
	return _D711StableCoin.Contract.Decimals(&_D711StableCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract_metadata method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_D711StableCoin *D711StableCoinCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _D711StableCoin.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract_metadata method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_D711StableCoin *D711StableCoinSession) Name() (string, error) {
	return _D711StableCoin.Contract.Name(&_D711StableCoin.CallOpts)
}

// Name is a free data retrieval call binding the contract_metadata method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_D711StableCoin *D711StableCoinCallerSession) Name() (string, error) {
	return _D711StableCoin.Contract.Name(&_D711StableCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract_metadata method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_D711StableCoin *D711StableCoinCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _D711StableCoin.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract_metadata method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_D711StableCoin *D711StableCoinSession) Owner() (common.Address, error) {
	return _D711StableCoin.Contract.Owner(&_D711StableCoin.CallOpts)
}

// Owner is a free data retrieval call binding the contract_metadata method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_D711StableCoin *D711StableCoinCallerSession) Owner() (common.Address, error) {
	return _D711StableCoin.Contract.Owner(&_D711StableCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract_metadata method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_D711StableCoin *D711StableCoinCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _D711StableCoin.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract_metadata method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_D711StableCoin *D711StableCoinSession) Symbol() (string, error) {
	return _D711StableCoin.Contract.Symbol(&_D711StableCoin.CallOpts)
}

// Symbol is a free data retrieval call binding the contract_metadata method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_D711StableCoin *D711StableCoinCallerSession) Symbol() (string, error) {
	return _D711StableCoin.Contract.Symbol(&_D711StableCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract_metadata method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_D711StableCoin *D711StableCoinCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _D711StableCoin.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract_metadata method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_D711StableCoin *D711StableCoinSession) TotalSupply() (*big.Int, error) {
	return _D711StableCoin.Contract.TotalSupply(&_D711StableCoin.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract_metadata method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_D711StableCoin *D711StableCoinCallerSession) TotalSupply() (*big.Int, error) {
	return _D711StableCoin.Contract.TotalSupply(&_D711StableCoin.CallOpts)
}

// Approve is a paid mutator transaction binding the contract_metadata method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_D711StableCoin *D711StableCoinTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract_metadata method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_D711StableCoin *D711StableCoinSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.Approve(&_D711StableCoin.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract_metadata method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_D711StableCoin *D711StableCoinTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.Approve(&_D711StableCoin.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract_metadata method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_D711StableCoin *D711StableCoinTransactor) Burn(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.contract.Transact(opts, "burn", _amount)
}

// Burn is a paid mutator transaction binding the contract_metadata method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_D711StableCoin *D711StableCoinSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.Burn(&_D711StableCoin.TransactOpts, _amount)
}

// Burn is a paid mutator transaction binding the contract_metadata method 0x42966c68.
//
// Solidity: function burn(uint256 _amount) returns()
func (_D711StableCoin *D711StableCoinTransactorSession) Burn(_amount *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.Burn(&_D711StableCoin.TransactOpts, _amount)
}

// BurnFrom is a paid mutator transaction binding the contract_metadata method 0x79cc6790.
//
// Solidity: function burnFrom(address _from, uint256 _amount) returns()
func (_D711StableCoin *D711StableCoinTransactor) BurnFrom(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.contract.Transact(opts, "burnFrom", _from, _amount)
}

// BurnFrom is a paid mutator transaction binding the contract_metadata method 0x79cc6790.
//
// Solidity: function burnFrom(address _from, uint256 _amount) returns()
func (_D711StableCoin *D711StableCoinSession) BurnFrom(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.BurnFrom(&_D711StableCoin.TransactOpts, _from, _amount)
}

// BurnFrom is a paid mutator transaction binding the contract_metadata method 0x79cc6790.
//
// Solidity: function burnFrom(address _from, uint256 _amount) returns()
func (_D711StableCoin *D711StableCoinTransactorSession) BurnFrom(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.BurnFrom(&_D711StableCoin.TransactOpts, _from, _amount)
}

// Mint is a paid mutator transaction binding the contract_metadata method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_D711StableCoin *D711StableCoinTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract_metadata method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_D711StableCoin *D711StableCoinSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.Mint(&_D711StableCoin.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract_metadata method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns(bool)
func (_D711StableCoin *D711StableCoinTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.Mint(&_D711StableCoin.TransactOpts, _to, _amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract_metadata method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_D711StableCoin *D711StableCoinTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _D711StableCoin.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract_metadata method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_D711StableCoin *D711StableCoinSession) RenounceOwnership() (*types.Transaction, error) {
	return _D711StableCoin.Contract.RenounceOwnership(&_D711StableCoin.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract_metadata method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_D711StableCoin *D711StableCoinTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _D711StableCoin.Contract.RenounceOwnership(&_D711StableCoin.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract_metadata method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_D711StableCoin *D711StableCoinTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract_metadata method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_D711StableCoin *D711StableCoinSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.Transfer(&_D711StableCoin.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract_metadata method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_D711StableCoin *D711StableCoinTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.Transfer(&_D711StableCoin.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract_metadata method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_D711StableCoin *D711StableCoinTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract_metadata method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_D711StableCoin *D711StableCoinSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.TransferFrom(&_D711StableCoin.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract_metadata method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_D711StableCoin *D711StableCoinTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _D711StableCoin.Contract.TransferFrom(&_D711StableCoin.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract_metadata method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_D711StableCoin *D711StableCoinTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _D711StableCoin.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract_metadata method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_D711StableCoin *D711StableCoinSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _D711StableCoin.Contract.TransferOwnership(&_D711StableCoin.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract_metadata method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_D711StableCoin *D711StableCoinTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _D711StableCoin.Contract.TransferOwnership(&_D711StableCoin.TransactOpts, newOwner)
}

// D711StableCoinApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the D711StableCoin contract_metadata.
type D711StableCoinApprovalIterator struct {
	Event *D711StableCoinApproval // Event containing the contract_metadata specifics and raw log

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
func (it *D711StableCoinApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D711StableCoinApproval)
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
		it.Event = new(D711StableCoinApproval)
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
func (it *D711StableCoinApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D711StableCoinApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D711StableCoinApproval represents a Approval event raised by the D711StableCoin contract_metadata.
type D711StableCoinApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract_metadata event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_D711StableCoin *D711StableCoinFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*D711StableCoinApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _D711StableCoin.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &D711StableCoinApprovalIterator{contract: _D711StableCoin.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract_metadata event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_D711StableCoin *D711StableCoinFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *D711StableCoinApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _D711StableCoin.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D711StableCoinApproval)
				if err := _D711StableCoin.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract_metadata event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_D711StableCoin *D711StableCoinFilterer) ParseApproval(log types.Log) (*D711StableCoinApproval, error) {
	event := new(D711StableCoinApproval)
	if err := _D711StableCoin.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D711StableCoinOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the D711StableCoin contract_metadata.
type D711StableCoinOwnershipTransferredIterator struct {
	Event *D711StableCoinOwnershipTransferred // Event containing the contract_metadata specifics and raw log

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
func (it *D711StableCoinOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D711StableCoinOwnershipTransferred)
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
		it.Event = new(D711StableCoinOwnershipTransferred)
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
func (it *D711StableCoinOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D711StableCoinOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D711StableCoinOwnershipTransferred represents a OwnershipTransferred event raised by the D711StableCoin contract_metadata.
type D711StableCoinOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract_metadata event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_D711StableCoin *D711StableCoinFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*D711StableCoinOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _D711StableCoin.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &D711StableCoinOwnershipTransferredIterator{contract: _D711StableCoin.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract_metadata event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_D711StableCoin *D711StableCoinFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *D711StableCoinOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _D711StableCoin.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D711StableCoinOwnershipTransferred)
				if err := _D711StableCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract_metadata event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_D711StableCoin *D711StableCoinFilterer) ParseOwnershipTransferred(log types.Log) (*D711StableCoinOwnershipTransferred, error) {
	event := new(D711StableCoinOwnershipTransferred)
	if err := _D711StableCoin.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// D711StableCoinTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the D711StableCoin contract_metadata.
type D711StableCoinTransferIterator struct {
	Event *D711StableCoinTransfer // Event containing the contract_metadata specifics and raw log

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
func (it *D711StableCoinTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(D711StableCoinTransfer)
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
		it.Event = new(D711StableCoinTransfer)
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
func (it *D711StableCoinTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *D711StableCoinTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// D711StableCoinTransfer represents a Transfer event raised by the D711StableCoin contract_metadata.
type D711StableCoinTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract_metadata event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_D711StableCoin *D711StableCoinFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*D711StableCoinTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _D711StableCoin.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &D711StableCoinTransferIterator{contract: _D711StableCoin.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract_metadata event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_D711StableCoin *D711StableCoinFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *D711StableCoinTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _D711StableCoin.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(D711StableCoinTransfer)
				if err := _D711StableCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract_metadata event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_D711StableCoin *D711StableCoinFilterer) ParseTransfer(log types.Log) (*D711StableCoinTransfer, error) {
	event := new(D711StableCoinTransfer)
	if err := _D711StableCoin.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
