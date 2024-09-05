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

// GameManagementGameRecord is an auto generated low-level Go binding around an user-defined struct.
type GameManagementGameRecord struct {
	RoomId             *big.Int
	IsStarted          bool
	IsEnded            bool
	PlayerAddress      []common.Address
	BetAmounts         []*big.Int
	Winner             common.Address
	PrizePool          *big.Int
	MinQualifiedAmount *big.Int
}

// GameManagementMetaData contains all meta data concerning the GameManagement contract_metadata.
var GameManagementMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"eeAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endGame\",\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_winner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userAddress\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"betAmount\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getGameRecord\",\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"record\",\"type\":\"tuple\",\"internalType\":\"structGameManagement.GameRecord\",\"components\":[{\"name\":\"roomId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"isStarted\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"isEnded\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"playerAddress\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"betAmounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"winner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"prizePool\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"minQualifiedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlayerBalance\",\"inputs\":[{\"name\":\"userAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startGame\",\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"userAddress\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"minQualifiedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"topUpEe\",\"inputs\":[{\"name\":\"amountEe\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawEe\",\"inputs\":[{\"name\":\"amountEe\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EeApproved\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EeTopUp\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EeWithdraw\",\"inputs\":[{\"name\":\"player\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GameEnded\",\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"winner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GameStarted\",\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PrizeDistributed\",\"inputs\":[{\"name\":\"roomId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"winner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"prizeAmount\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"GameManagement__TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// GameManagementABI is the input ABI used to generate the binding from.
// Deprecated: Use GameManagementMetaData.ABI instead.
var GameManagementABI = GameManagementMetaData.ABI

// GameManagement is an auto generated Go binding around an Ethereum contract_metadata.
type GameManagement struct {
	GameManagementCaller     // Read-only binding to the contract_metadata
	GameManagementTransactor // Write-only binding to the contract_metadata
	GameManagementFilterer   // Log filterer for contract_metadata events
}

// GameManagementCaller is an auto generated read-only Go binding around an Ethereum contract_metadata.
type GameManagementCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameManagementTransactor is an auto generated write-only Go binding around an Ethereum contract_metadata.
type GameManagementTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameManagementFilterer is an auto generated log filtering Go binding around an Ethereum contract_metadata events.
type GameManagementFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GameManagementSession is an auto generated Go binding around an Ethereum contract_metadata,
// with pre-set call and transact options.
type GameManagementSession struct {
	Contract     *GameManagement   // Generic contract_metadata binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GameManagementCallerSession is an auto generated read-only Go binding around an Ethereum contract_metadata,
// with pre-set call options.
type GameManagementCallerSession struct {
	Contract *GameManagementCaller // Generic contract_metadata caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GameManagementTransactorSession is an auto generated write-only Go binding around an Ethereum contract_metadata,
// with pre-set transact options.
type GameManagementTransactorSession struct {
	Contract     *GameManagementTransactor // Generic contract_metadata transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GameManagementRaw is an auto generated low-level Go binding around an Ethereum contract_metadata.
type GameManagementRaw struct {
	Contract *GameManagement // Generic contract_metadata binding to access the raw methods on
}

// GameManagementCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract_metadata.
type GameManagementCallerRaw struct {
	Contract *GameManagementCaller // Generic read-only contract_metadata binding to access the raw methods on
}

// GameManagementTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract_metadata.
type GameManagementTransactorRaw struct {
	Contract *GameManagementTransactor // Generic write-only contract_metadata binding to access the raw methods on
}

// NewGameManagement creates a new instance of GameManagement, bound to a specific deployed contract_metadata.
func NewGameManagement(address common.Address, backend bind.ContractBackend) (*GameManagement, error) {
	contract, err := bindGameManagement(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GameManagement{GameManagementCaller: GameManagementCaller{contract: contract}, GameManagementTransactor: GameManagementTransactor{contract: contract}, GameManagementFilterer: GameManagementFilterer{contract: contract}}, nil
}

// NewGameManagementCaller creates a new read-only instance of GameManagement, bound to a specific deployed contract_metadata.
func NewGameManagementCaller(address common.Address, caller bind.ContractCaller) (*GameManagementCaller, error) {
	contract, err := bindGameManagement(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GameManagementCaller{contract: contract}, nil
}

// NewGameManagementTransactor creates a new write-only instance of GameManagement, bound to a specific deployed contract_metadata.
func NewGameManagementTransactor(address common.Address, transactor bind.ContractTransactor) (*GameManagementTransactor, error) {
	contract, err := bindGameManagement(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GameManagementTransactor{contract: contract}, nil
}

// NewGameManagementFilterer creates a new log filterer instance of GameManagement, bound to a specific deployed contract_metadata.
func NewGameManagementFilterer(address common.Address, filterer bind.ContractFilterer) (*GameManagementFilterer, error) {
	contract, err := bindGameManagement(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GameManagementFilterer{contract: contract}, nil
}

// bindGameManagement binds a generic wrapper to an already deployed contract_metadata.
func bindGameManagement(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GameManagementMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract_metadata method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameManagement *GameManagementRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameManagement.Contract.GameManagementCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract_metadata, calling
// its default method if one is available.
func (_GameManagement *GameManagementRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameManagement.Contract.GameManagementTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract_metadata method with params as input values.
func (_GameManagement *GameManagementRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameManagement.Contract.GameManagementTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract_metadata method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GameManagement *GameManagementCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GameManagement.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract_metadata, calling
// its default method if one is available.
func (_GameManagement *GameManagementTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameManagement.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract_metadata method with params as input values.
func (_GameManagement *GameManagementTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GameManagement.Contract.contract.Transact(opts, method, params...)
}

// GetGameRecord is a free data retrieval call binding the contract_metadata method 0xe233d0a5.
//
// Solidity: function getGameRecord(uint256 roomId) view returns((uint256,bool,bool,address[],uint256[],address,uint256,uint256) record)
func (_GameManagement *GameManagementCaller) GetGameRecord(opts *bind.CallOpts, roomId *big.Int) (GameManagementGameRecord, error) {
	var out []interface{}
	err := _GameManagement.contract.Call(opts, &out, "getGameRecord", roomId)

	if err != nil {
		return *new(GameManagementGameRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(GameManagementGameRecord)).(*GameManagementGameRecord)

	return out0, err

}

// GetGameRecord is a free data retrieval call binding the contract_metadata method 0xe233d0a5.
//
// Solidity: function getGameRecord(uint256 roomId) view returns((uint256,bool,bool,address[],uint256[],address,uint256,uint256) record)
func (_GameManagement *GameManagementSession) GetGameRecord(roomId *big.Int) (GameManagementGameRecord, error) {
	return _GameManagement.Contract.GetGameRecord(&_GameManagement.CallOpts, roomId)
}

// GetGameRecord is a free data retrieval call binding the contract_metadata method 0xe233d0a5.
//
// Solidity: function getGameRecord(uint256 roomId) view returns((uint256,bool,bool,address[],uint256[],address,uint256,uint256) record)
func (_GameManagement *GameManagementCallerSession) GetGameRecord(roomId *big.Int) (GameManagementGameRecord, error) {
	return _GameManagement.Contract.GetGameRecord(&_GameManagement.CallOpts, roomId)
}

// GetPlayerBalance is a free data retrieval call binding the contract_metadata method 0xe2734c93.
//
// Solidity: function getPlayerBalance(address userAddr) view returns(uint256)
func (_GameManagement *GameManagementCaller) GetPlayerBalance(opts *bind.CallOpts, userAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GameManagement.contract.Call(opts, &out, "getPlayerBalance", userAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPlayerBalance is a free data retrieval call binding the contract_metadata method 0xe2734c93.
//
// Solidity: function getPlayerBalance(address userAddr) view returns(uint256)
func (_GameManagement *GameManagementSession) GetPlayerBalance(userAddr common.Address) (*big.Int, error) {
	return _GameManagement.Contract.GetPlayerBalance(&_GameManagement.CallOpts, userAddr)
}

// GetPlayerBalance is a free data retrieval call binding the contract_metadata method 0xe2734c93.
//
// Solidity: function getPlayerBalance(address userAddr) view returns(uint256)
func (_GameManagement *GameManagementCallerSession) GetPlayerBalance(userAddr common.Address) (*big.Int, error) {
	return _GameManagement.Contract.GetPlayerBalance(&_GameManagement.CallOpts, userAddr)
}

// Owner is a free data retrieval call binding the contract_metadata method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameManagement *GameManagementCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GameManagement.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract_metadata method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameManagement *GameManagementSession) Owner() (common.Address, error) {
	return _GameManagement.Contract.Owner(&_GameManagement.CallOpts)
}

// Owner is a free data retrieval call binding the contract_metadata method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_GameManagement *GameManagementCallerSession) Owner() (common.Address, error) {
	return _GameManagement.Contract.Owner(&_GameManagement.CallOpts)
}

// EndGame is a paid mutator transaction binding the contract_metadata method 0x732a64f6.
//
// Solidity: function endGame(uint256 roomId, address _winner, address[] userAddress, uint256[] betAmount) returns()
func (_GameManagement *GameManagementTransactor) EndGame(opts *bind.TransactOpts, roomId *big.Int, _winner common.Address, userAddress []common.Address, betAmount []*big.Int) (*types.Transaction, error) {
	return _GameManagement.contract.Transact(opts, "endGame", roomId, _winner, userAddress, betAmount)
}

// EndGame is a paid mutator transaction binding the contract_metadata method 0x732a64f6.
//
// Solidity: function endGame(uint256 roomId, address _winner, address[] userAddress, uint256[] betAmount) returns()
func (_GameManagement *GameManagementSession) EndGame(roomId *big.Int, _winner common.Address, userAddress []common.Address, betAmount []*big.Int) (*types.Transaction, error) {
	return _GameManagement.Contract.EndGame(&_GameManagement.TransactOpts, roomId, _winner, userAddress, betAmount)
}

// EndGame is a paid mutator transaction binding the contract_metadata method 0x732a64f6.
//
// Solidity: function endGame(uint256 roomId, address _winner, address[] userAddress, uint256[] betAmount) returns()
func (_GameManagement *GameManagementTransactorSession) EndGame(roomId *big.Int, _winner common.Address, userAddress []common.Address, betAmount []*big.Int) (*types.Transaction, error) {
	return _GameManagement.Contract.EndGame(&_GameManagement.TransactOpts, roomId, _winner, userAddress, betAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract_metadata method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameManagement *GameManagementTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GameManagement.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract_metadata method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameManagement *GameManagementSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameManagement.Contract.RenounceOwnership(&_GameManagement.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract_metadata method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GameManagement *GameManagementTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GameManagement.Contract.RenounceOwnership(&_GameManagement.TransactOpts)
}

// StartGame is a paid mutator transaction binding the contract_metadata method 0x69dbbc28.
//
// Solidity: function startGame(uint256 roomId, address[] userAddress, uint256 minQualifiedAmount) returns()
func (_GameManagement *GameManagementTransactor) StartGame(opts *bind.TransactOpts, roomId *big.Int, userAddress []common.Address, minQualifiedAmount *big.Int) (*types.Transaction, error) {
	return _GameManagement.contract.Transact(opts, "startGame", roomId, userAddress, minQualifiedAmount)
}

// StartGame is a paid mutator transaction binding the contract_metadata method 0x69dbbc28.
//
// Solidity: function startGame(uint256 roomId, address[] userAddress, uint256 minQualifiedAmount) returns()
func (_GameManagement *GameManagementSession) StartGame(roomId *big.Int, userAddress []common.Address, minQualifiedAmount *big.Int) (*types.Transaction, error) {
	return _GameManagement.Contract.StartGame(&_GameManagement.TransactOpts, roomId, userAddress, minQualifiedAmount)
}

// StartGame is a paid mutator transaction binding the contract_metadata method 0x69dbbc28.
//
// Solidity: function startGame(uint256 roomId, address[] userAddress, uint256 minQualifiedAmount) returns()
func (_GameManagement *GameManagementTransactorSession) StartGame(roomId *big.Int, userAddress []common.Address, minQualifiedAmount *big.Int) (*types.Transaction, error) {
	return _GameManagement.Contract.StartGame(&_GameManagement.TransactOpts, roomId, userAddress, minQualifiedAmount)
}

// TopUpEe is a paid mutator transaction binding the contract_metadata method 0xc0826bf2.
//
// Solidity: function topUpEe(uint256 amountEe) returns()
func (_GameManagement *GameManagementTransactor) TopUpEe(opts *bind.TransactOpts, amountEe *big.Int) (*types.Transaction, error) {
	return _GameManagement.contract.Transact(opts, "topUpEe", amountEe)
}

// TopUpEe is a paid mutator transaction binding the contract_metadata method 0xc0826bf2.
//
// Solidity: function topUpEe(uint256 amountEe) returns()
func (_GameManagement *GameManagementSession) TopUpEe(amountEe *big.Int) (*types.Transaction, error) {
	return _GameManagement.Contract.TopUpEe(&_GameManagement.TransactOpts, amountEe)
}

// TopUpEe is a paid mutator transaction binding the contract_metadata method 0xc0826bf2.
//
// Solidity: function topUpEe(uint256 amountEe) returns()
func (_GameManagement *GameManagementTransactorSession) TopUpEe(amountEe *big.Int) (*types.Transaction, error) {
	return _GameManagement.Contract.TopUpEe(&_GameManagement.TransactOpts, amountEe)
}

// TransferOwnership is a paid mutator transaction binding the contract_metadata method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameManagement *GameManagementTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GameManagement.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract_metadata method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameManagement *GameManagementSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GameManagement.Contract.TransferOwnership(&_GameManagement.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract_metadata method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GameManagement *GameManagementTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GameManagement.Contract.TransferOwnership(&_GameManagement.TransactOpts, newOwner)
}

// WithdrawEe is a paid mutator transaction binding the contract_metadata method 0x61e27142.
//
// Solidity: function withdrawEe(uint256 amountEe) returns()
func (_GameManagement *GameManagementTransactor) WithdrawEe(opts *bind.TransactOpts, amountEe *big.Int) (*types.Transaction, error) {
	return _GameManagement.contract.Transact(opts, "withdrawEe", amountEe)
}

// WithdrawEe is a paid mutator transaction binding the contract_metadata method 0x61e27142.
//
// Solidity: function withdrawEe(uint256 amountEe) returns()
func (_GameManagement *GameManagementSession) WithdrawEe(amountEe *big.Int) (*types.Transaction, error) {
	return _GameManagement.Contract.WithdrawEe(&_GameManagement.TransactOpts, amountEe)
}

// WithdrawEe is a paid mutator transaction binding the contract_metadata method 0x61e27142.
//
// Solidity: function withdrawEe(uint256 amountEe) returns()
func (_GameManagement *GameManagementTransactorSession) WithdrawEe(amountEe *big.Int) (*types.Transaction, error) {
	return _GameManagement.Contract.WithdrawEe(&_GameManagement.TransactOpts, amountEe)
}

// GameManagementEeApprovedIterator is returned from FilterEeApproved and is used to iterate over the raw logs and unpacked data for EeApproved events raised by the GameManagement contract_metadata.
type GameManagementEeApprovedIterator struct {
	Event *GameManagementEeApproved // Event containing the contract_metadata specifics and raw log

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
func (it *GameManagementEeApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameManagementEeApproved)
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
		it.Event = new(GameManagementEeApproved)
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
func (it *GameManagementEeApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameManagementEeApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameManagementEeApproved represents a EeApproved event raised by the GameManagement contract_metadata.
type GameManagementEeApproved struct {
	Player common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEeApproved is a free log retrieval operation binding the contract_metadata event 0x64b7d18df31d06bde999e7972266398b7a7547e298a08fbcb945574ceb0af870.
//
// Solidity: event EeApproved(address indexed player, uint256 amount)
func (_GameManagement *GameManagementFilterer) FilterEeApproved(opts *bind.FilterOpts, player []common.Address) (*GameManagementEeApprovedIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _GameManagement.contract.FilterLogs(opts, "EeApproved", playerRule)
	if err != nil {
		return nil, err
	}
	return &GameManagementEeApprovedIterator{contract: _GameManagement.contract, event: "EeApproved", logs: logs, sub: sub}, nil
}

// WatchEeApproved is a free log subscription operation binding the contract_metadata event 0x64b7d18df31d06bde999e7972266398b7a7547e298a08fbcb945574ceb0af870.
//
// Solidity: event EeApproved(address indexed player, uint256 amount)
func (_GameManagement *GameManagementFilterer) WatchEeApproved(opts *bind.WatchOpts, sink chan<- *GameManagementEeApproved, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _GameManagement.contract.WatchLogs(opts, "EeApproved", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameManagementEeApproved)
				if err := _GameManagement.contract.UnpackLog(event, "EeApproved", log); err != nil {
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

// ParseEeApproved is a log parse operation binding the contract_metadata event 0x64b7d18df31d06bde999e7972266398b7a7547e298a08fbcb945574ceb0af870.
//
// Solidity: event EeApproved(address indexed player, uint256 amount)
func (_GameManagement *GameManagementFilterer) ParseEeApproved(log types.Log) (*GameManagementEeApproved, error) {
	event := new(GameManagementEeApproved)
	if err := _GameManagement.contract.UnpackLog(event, "EeApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameManagementEeTopUpIterator is returned from FilterEeTopUp and is used to iterate over the raw logs and unpacked data for EeTopUp events raised by the GameManagement contract_metadata.
type GameManagementEeTopUpIterator struct {
	Event *GameManagementEeTopUp // Event containing the contract_metadata specifics and raw log

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
func (it *GameManagementEeTopUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameManagementEeTopUp)
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
		it.Event = new(GameManagementEeTopUp)
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
func (it *GameManagementEeTopUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameManagementEeTopUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameManagementEeTopUp represents a EeTopUp event raised by the GameManagement contract_metadata.
type GameManagementEeTopUp struct {
	Player common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEeTopUp is a free log retrieval operation binding the contract_metadata event 0x322e09de7fbbfd229f6290c0ffefb880a1e284d8bef22a2009f8a1c4310cf78e.
//
// Solidity: event EeTopUp(address indexed player, uint256 amount)
func (_GameManagement *GameManagementFilterer) FilterEeTopUp(opts *bind.FilterOpts, player []common.Address) (*GameManagementEeTopUpIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _GameManagement.contract.FilterLogs(opts, "EeTopUp", playerRule)
	if err != nil {
		return nil, err
	}
	return &GameManagementEeTopUpIterator{contract: _GameManagement.contract, event: "EeTopUp", logs: logs, sub: sub}, nil
}

// WatchEeTopUp is a free log subscription operation binding the contract_metadata event 0x322e09de7fbbfd229f6290c0ffefb880a1e284d8bef22a2009f8a1c4310cf78e.
//
// Solidity: event EeTopUp(address indexed player, uint256 amount)
func (_GameManagement *GameManagementFilterer) WatchEeTopUp(opts *bind.WatchOpts, sink chan<- *GameManagementEeTopUp, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _GameManagement.contract.WatchLogs(opts, "EeTopUp", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameManagementEeTopUp)
				if err := _GameManagement.contract.UnpackLog(event, "EeTopUp", log); err != nil {
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

// ParseEeTopUp is a log parse operation binding the contract_metadata event 0x322e09de7fbbfd229f6290c0ffefb880a1e284d8bef22a2009f8a1c4310cf78e.
//
// Solidity: event EeTopUp(address indexed player, uint256 amount)
func (_GameManagement *GameManagementFilterer) ParseEeTopUp(log types.Log) (*GameManagementEeTopUp, error) {
	event := new(GameManagementEeTopUp)
	if err := _GameManagement.contract.UnpackLog(event, "EeTopUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameManagementEeWithdrawIterator is returned from FilterEeWithdraw and is used to iterate over the raw logs and unpacked data for EeWithdraw events raised by the GameManagement contract_metadata.
type GameManagementEeWithdrawIterator struct {
	Event *GameManagementEeWithdraw // Event containing the contract_metadata specifics and raw log

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
func (it *GameManagementEeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameManagementEeWithdraw)
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
		it.Event = new(GameManagementEeWithdraw)
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
func (it *GameManagementEeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameManagementEeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameManagementEeWithdraw represents a EeWithdraw event raised by the GameManagement contract_metadata.
type GameManagementEeWithdraw struct {
	Player common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEeWithdraw is a free log retrieval operation binding the contract_metadata event 0x57e338a7a0f383100f65067e4ee85d34d6d36b08b0560109bff727945d3582fe.
//
// Solidity: event EeWithdraw(address indexed player, uint256 amount)
func (_GameManagement *GameManagementFilterer) FilterEeWithdraw(opts *bind.FilterOpts, player []common.Address) (*GameManagementEeWithdrawIterator, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _GameManagement.contract.FilterLogs(opts, "EeWithdraw", playerRule)
	if err != nil {
		return nil, err
	}
	return &GameManagementEeWithdrawIterator{contract: _GameManagement.contract, event: "EeWithdraw", logs: logs, sub: sub}, nil
}

// WatchEeWithdraw is a free log subscription operation binding the contract_metadata event 0x57e338a7a0f383100f65067e4ee85d34d6d36b08b0560109bff727945d3582fe.
//
// Solidity: event EeWithdraw(address indexed player, uint256 amount)
func (_GameManagement *GameManagementFilterer) WatchEeWithdraw(opts *bind.WatchOpts, sink chan<- *GameManagementEeWithdraw, player []common.Address) (event.Subscription, error) {

	var playerRule []interface{}
	for _, playerItem := range player {
		playerRule = append(playerRule, playerItem)
	}

	logs, sub, err := _GameManagement.contract.WatchLogs(opts, "EeWithdraw", playerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameManagementEeWithdraw)
				if err := _GameManagement.contract.UnpackLog(event, "EeWithdraw", log); err != nil {
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

// ParseEeWithdraw is a log parse operation binding the contract_metadata event 0x57e338a7a0f383100f65067e4ee85d34d6d36b08b0560109bff727945d3582fe.
//
// Solidity: event EeWithdraw(address indexed player, uint256 amount)
func (_GameManagement *GameManagementFilterer) ParseEeWithdraw(log types.Log) (*GameManagementEeWithdraw, error) {
	event := new(GameManagementEeWithdraw)
	if err := _GameManagement.contract.UnpackLog(event, "EeWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameManagementGameEndedIterator is returned from FilterGameEnded and is used to iterate over the raw logs and unpacked data for GameEnded events raised by the GameManagement contract_metadata.
type GameManagementGameEndedIterator struct {
	Event *GameManagementGameEnded // Event containing the contract_metadata specifics and raw log

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
func (it *GameManagementGameEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameManagementGameEnded)
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
		it.Event = new(GameManagementGameEnded)
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
func (it *GameManagementGameEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameManagementGameEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameManagementGameEnded represents a GameEnded event raised by the GameManagement contract_metadata.
type GameManagementGameEnded struct {
	RoomId *big.Int
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGameEnded is a free log retrieval operation binding the contract_metadata event 0x3d5fcd8fa3235aeb3897f25c8ea49989174981afcb01b9d03189d1056a16775b.
//
// Solidity: event GameEnded(uint256 indexed roomId, address indexed winner)
func (_GameManagement *GameManagementFilterer) FilterGameEnded(opts *bind.FilterOpts, roomId []*big.Int, winner []common.Address) (*GameManagementGameEndedIterator, error) {

	var roomIdRule []interface{}
	for _, roomIdItem := range roomId {
		roomIdRule = append(roomIdRule, roomIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _GameManagement.contract.FilterLogs(opts, "GameEnded", roomIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &GameManagementGameEndedIterator{contract: _GameManagement.contract, event: "GameEnded", logs: logs, sub: sub}, nil
}

// WatchGameEnded is a free log subscription operation binding the contract_metadata event 0x3d5fcd8fa3235aeb3897f25c8ea49989174981afcb01b9d03189d1056a16775b.
//
// Solidity: event GameEnded(uint256 indexed roomId, address indexed winner)
func (_GameManagement *GameManagementFilterer) WatchGameEnded(opts *bind.WatchOpts, sink chan<- *GameManagementGameEnded, roomId []*big.Int, winner []common.Address) (event.Subscription, error) {

	var roomIdRule []interface{}
	for _, roomIdItem := range roomId {
		roomIdRule = append(roomIdRule, roomIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _GameManagement.contract.WatchLogs(opts, "GameEnded", roomIdRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameManagementGameEnded)
				if err := _GameManagement.contract.UnpackLog(event, "GameEnded", log); err != nil {
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

// ParseGameEnded is a log parse operation binding the contract_metadata event 0x3d5fcd8fa3235aeb3897f25c8ea49989174981afcb01b9d03189d1056a16775b.
//
// Solidity: event GameEnded(uint256 indexed roomId, address indexed winner)
func (_GameManagement *GameManagementFilterer) ParseGameEnded(log types.Log) (*GameManagementGameEnded, error) {
	event := new(GameManagementGameEnded)
	if err := _GameManagement.contract.UnpackLog(event, "GameEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameManagementGameStartedIterator is returned from FilterGameStarted and is used to iterate over the raw logs and unpacked data for GameStarted events raised by the GameManagement contract_metadata.
type GameManagementGameStartedIterator struct {
	Event *GameManagementGameStarted // Event containing the contract_metadata specifics and raw log

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
func (it *GameManagementGameStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameManagementGameStarted)
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
		it.Event = new(GameManagementGameStarted)
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
func (it *GameManagementGameStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameManagementGameStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameManagementGameStarted represents a GameStarted event raised by the GameManagement contract_metadata.
type GameManagementGameStarted struct {
	RoomId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGameStarted is a free log retrieval operation binding the contract_metadata event 0x50ad08f58a27f2851d7e3a1b3a6a46b290f2ce677e99642d30ff639721e77790.
//
// Solidity: event GameStarted(uint256 indexed roomId)
func (_GameManagement *GameManagementFilterer) FilterGameStarted(opts *bind.FilterOpts, roomId []*big.Int) (*GameManagementGameStartedIterator, error) {

	var roomIdRule []interface{}
	for _, roomIdItem := range roomId {
		roomIdRule = append(roomIdRule, roomIdItem)
	}

	logs, sub, err := _GameManagement.contract.FilterLogs(opts, "GameStarted", roomIdRule)
	if err != nil {
		return nil, err
	}
	return &GameManagementGameStartedIterator{contract: _GameManagement.contract, event: "GameStarted", logs: logs, sub: sub}, nil
}

// WatchGameStarted is a free log subscription operation binding the contract_metadata event 0x50ad08f58a27f2851d7e3a1b3a6a46b290f2ce677e99642d30ff639721e77790.
//
// Solidity: event GameStarted(uint256 indexed roomId)
func (_GameManagement *GameManagementFilterer) WatchGameStarted(opts *bind.WatchOpts, sink chan<- *GameManagementGameStarted, roomId []*big.Int) (event.Subscription, error) {

	var roomIdRule []interface{}
	for _, roomIdItem := range roomId {
		roomIdRule = append(roomIdRule, roomIdItem)
	}

	logs, sub, err := _GameManagement.contract.WatchLogs(opts, "GameStarted", roomIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameManagementGameStarted)
				if err := _GameManagement.contract.UnpackLog(event, "GameStarted", log); err != nil {
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

// ParseGameStarted is a log parse operation binding the contract_metadata event 0x50ad08f58a27f2851d7e3a1b3a6a46b290f2ce677e99642d30ff639721e77790.
//
// Solidity: event GameStarted(uint256 indexed roomId)
func (_GameManagement *GameManagementFilterer) ParseGameStarted(log types.Log) (*GameManagementGameStarted, error) {
	event := new(GameManagementGameStarted)
	if err := _GameManagement.contract.UnpackLog(event, "GameStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameManagementOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GameManagement contract_metadata.
type GameManagementOwnershipTransferredIterator struct {
	Event *GameManagementOwnershipTransferred // Event containing the contract_metadata specifics and raw log

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
func (it *GameManagementOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameManagementOwnershipTransferred)
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
		it.Event = new(GameManagementOwnershipTransferred)
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
func (it *GameManagementOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameManagementOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameManagementOwnershipTransferred represents a OwnershipTransferred event raised by the GameManagement contract_metadata.
type GameManagementOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract_metadata event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameManagement *GameManagementFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GameManagementOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameManagement.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GameManagementOwnershipTransferredIterator{contract: _GameManagement.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract_metadata event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GameManagement *GameManagementFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GameManagementOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GameManagement.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameManagementOwnershipTransferred)
				if err := _GameManagement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GameManagement *GameManagementFilterer) ParseOwnershipTransferred(log types.Log) (*GameManagementOwnershipTransferred, error) {
	event := new(GameManagementOwnershipTransferred)
	if err := _GameManagement.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GameManagementPrizeDistributedIterator is returned from FilterPrizeDistributed and is used to iterate over the raw logs and unpacked data for PrizeDistributed events raised by the GameManagement contract_metadata.
type GameManagementPrizeDistributedIterator struct {
	Event *GameManagementPrizeDistributed // Event containing the contract_metadata specifics and raw log

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
func (it *GameManagementPrizeDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GameManagementPrizeDistributed)
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
		it.Event = new(GameManagementPrizeDistributed)
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
func (it *GameManagementPrizeDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GameManagementPrizeDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GameManagementPrizeDistributed represents a PrizeDistributed event raised by the GameManagement contract_metadata.
type GameManagementPrizeDistributed struct {
	RoomId      *big.Int
	Winner      common.Address
	PrizeAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPrizeDistributed is a free log retrieval operation binding the contract_metadata event 0xdc1d0b6a7cbfa6f3b2d4325bab5a169b265ad0b2b4022fb550364f0dfeeb6a74.
//
// Solidity: event PrizeDistributed(uint256 indexed roomId, address indexed winner, uint256 indexed prizeAmount)
func (_GameManagement *GameManagementFilterer) FilterPrizeDistributed(opts *bind.FilterOpts, roomId []*big.Int, winner []common.Address, prizeAmount []*big.Int) (*GameManagementPrizeDistributedIterator, error) {

	var roomIdRule []interface{}
	for _, roomIdItem := range roomId {
		roomIdRule = append(roomIdRule, roomIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var prizeAmountRule []interface{}
	for _, prizeAmountItem := range prizeAmount {
		prizeAmountRule = append(prizeAmountRule, prizeAmountItem)
	}

	logs, sub, err := _GameManagement.contract.FilterLogs(opts, "PrizeDistributed", roomIdRule, winnerRule, prizeAmountRule)
	if err != nil {
		return nil, err
	}
	return &GameManagementPrizeDistributedIterator{contract: _GameManagement.contract, event: "PrizeDistributed", logs: logs, sub: sub}, nil
}

// WatchPrizeDistributed is a free log subscription operation binding the contract_metadata event 0xdc1d0b6a7cbfa6f3b2d4325bab5a169b265ad0b2b4022fb550364f0dfeeb6a74.
//
// Solidity: event PrizeDistributed(uint256 indexed roomId, address indexed winner, uint256 indexed prizeAmount)
func (_GameManagement *GameManagementFilterer) WatchPrizeDistributed(opts *bind.WatchOpts, sink chan<- *GameManagementPrizeDistributed, roomId []*big.Int, winner []common.Address, prizeAmount []*big.Int) (event.Subscription, error) {

	var roomIdRule []interface{}
	for _, roomIdItem := range roomId {
		roomIdRule = append(roomIdRule, roomIdItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var prizeAmountRule []interface{}
	for _, prizeAmountItem := range prizeAmount {
		prizeAmountRule = append(prizeAmountRule, prizeAmountItem)
	}

	logs, sub, err := _GameManagement.contract.WatchLogs(opts, "PrizeDistributed", roomIdRule, winnerRule, prizeAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GameManagementPrizeDistributed)
				if err := _GameManagement.contract.UnpackLog(event, "PrizeDistributed", log); err != nil {
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

// ParsePrizeDistributed is a log parse operation binding the contract_metadata event 0xdc1d0b6a7cbfa6f3b2d4325bab5a169b265ad0b2b4022fb550364f0dfeeb6a74.
//
// Solidity: event PrizeDistributed(uint256 indexed roomId, address indexed winner, uint256 indexed prizeAmount)
func (_GameManagement *GameManagementFilterer) ParsePrizeDistributed(log types.Log) (*GameManagementPrizeDistributed, error) {
	event := new(GameManagementPrizeDistributed)
	if err := _GameManagement.contract.UnpackLog(event, "PrizeDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
