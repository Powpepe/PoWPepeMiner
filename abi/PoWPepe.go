// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// PoWPepeMetaData contains all meta data concerning the PoWPepe contract.
var PoWPepeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_initialSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals_\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_difficulty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_miningLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_initialLimitPerMint\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challenge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLimitPerMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRemainingSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"limitPerMint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"mine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minedNonces\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"miningLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"miningTimes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupplyCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoWPepeABI is the input ABI used to generate the binding from.
// Deprecated: Use PoWPepeMetaData.ABI instead.
var PoWPepeABI = PoWPepeMetaData.ABI

// PoWPepe is an auto generated Go binding around an Ethereum contract.
type PoWPepe struct {
	PoWPepeCaller     // Read-only binding to the contract
	PoWPepeTransactor // Write-only binding to the contract
	PoWPepeFilterer   // Log filterer for contract events
}

// PoWPepeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoWPepeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoWPepeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoWPepeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoWPepeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoWPepeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoWPepeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoWPepeSession struct {
	Contract     *PoWPepe         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoWPepeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoWPepeCallerSession struct {
	Contract *PoWPepeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PoWPepeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoWPepeTransactorSession struct {
	Contract     *PoWPepeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PoWPepeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoWPepeRaw struct {
	Contract *PoWPepe // Generic contract binding to access the raw methods on
}

// PoWPepeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoWPepeCallerRaw struct {
	Contract *PoWPepeCaller // Generic read-only contract binding to access the raw methods on
}

// PoWPepeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoWPepeTransactorRaw struct {
	Contract *PoWPepeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoWPepe creates a new instance of PoWPepe, bound to a specific deployed contract.
func NewPoWPepe(address common.Address, backend bind.ContractBackend) (*PoWPepe, error) {
	contract, err := bindPoWPepe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoWPepe{PoWPepeCaller: PoWPepeCaller{contract: contract}, PoWPepeTransactor: PoWPepeTransactor{contract: contract}, PoWPepeFilterer: PoWPepeFilterer{contract: contract}}, nil
}

// NewPoWPepeCaller creates a new read-only instance of PoWPepe, bound to a specific deployed contract.
func NewPoWPepeCaller(address common.Address, caller bind.ContractCaller) (*PoWPepeCaller, error) {
	contract, err := bindPoWPepe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoWPepeCaller{contract: contract}, nil
}

// NewPoWPepeTransactor creates a new write-only instance of PoWPepe, bound to a specific deployed contract.
func NewPoWPepeTransactor(address common.Address, transactor bind.ContractTransactor) (*PoWPepeTransactor, error) {
	contract, err := bindPoWPepe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoWPepeTransactor{contract: contract}, nil
}

// NewPoWPepeFilterer creates a new log filterer instance of PoWPepe, bound to a specific deployed contract.
func NewPoWPepeFilterer(address common.Address, filterer bind.ContractFilterer) (*PoWPepeFilterer, error) {
	contract, err := bindPoWPepe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoWPepeFilterer{contract: contract}, nil
}

// bindPoWPepe binds a generic wrapper to an already deployed contract.
func bindPoWPepe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoWPepeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoWPepe *PoWPepeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoWPepe.Contract.PoWPepeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoWPepe *PoWPepeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoWPepe.Contract.PoWPepeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoWPepe *PoWPepeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoWPepe.Contract.PoWPepeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoWPepe *PoWPepeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoWPepe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoWPepe *PoWPepeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoWPepe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoWPepe *PoWPepeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoWPepe.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PoWPepe *PoWPepeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PoWPepe *PoWPepeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PoWPepe.Contract.Allowance(&_PoWPepe.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_PoWPepe *PoWPepeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PoWPepe.Contract.Allowance(&_PoWPepe.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PoWPepe *PoWPepeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PoWPepe *PoWPepeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PoWPepe.Contract.BalanceOf(&_PoWPepe.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_PoWPepe *PoWPepeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _PoWPepe.Contract.BalanceOf(&_PoWPepe.CallOpts, account)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_PoWPepe *PoWPepeCaller) Challenge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "challenge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_PoWPepe *PoWPepeSession) Challenge() (*big.Int, error) {
	return _PoWPepe.Contract.Challenge(&_PoWPepe.CallOpts)
}

// Challenge is a free data retrieval call binding the contract method 0xd2ef7398.
//
// Solidity: function challenge() view returns(uint256)
func (_PoWPepe *PoWPepeCallerSession) Challenge() (*big.Int, error) {
	return _PoWPepe.Contract.Challenge(&_PoWPepe.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PoWPepe *PoWPepeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PoWPepe *PoWPepeSession) Decimals() (uint8, error) {
	return _PoWPepe.Contract.Decimals(&_PoWPepe.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_PoWPepe *PoWPepeCallerSession) Decimals() (uint8, error) {
	return _PoWPepe.Contract.Decimals(&_PoWPepe.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_PoWPepe *PoWPepeCaller) Difficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_PoWPepe *PoWPepeSession) Difficulty() (*big.Int, error) {
	return _PoWPepe.Contract.Difficulty(&_PoWPepe.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_PoWPepe *PoWPepeCallerSession) Difficulty() (*big.Int, error) {
	return _PoWPepe.Contract.Difficulty(&_PoWPepe.CallOpts)
}

// GetLimitPerMint is a free data retrieval call binding the contract method 0xb32e82c0.
//
// Solidity: function getLimitPerMint() view returns(uint256)
func (_PoWPepe *PoWPepeCaller) GetLimitPerMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "getLimitPerMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLimitPerMint is a free data retrieval call binding the contract method 0xb32e82c0.
//
// Solidity: function getLimitPerMint() view returns(uint256)
func (_PoWPepe *PoWPepeSession) GetLimitPerMint() (*big.Int, error) {
	return _PoWPepe.Contract.GetLimitPerMint(&_PoWPepe.CallOpts)
}

// GetLimitPerMint is a free data retrieval call binding the contract method 0xb32e82c0.
//
// Solidity: function getLimitPerMint() view returns(uint256)
func (_PoWPepe *PoWPepeCallerSession) GetLimitPerMint() (*big.Int, error) {
	return _PoWPepe.Contract.GetLimitPerMint(&_PoWPepe.CallOpts)
}

// GetRemainingSupply is a free data retrieval call binding the contract method 0xe4b7fb73.
//
// Solidity: function getRemainingSupply() view returns(uint256)
func (_PoWPepe *PoWPepeCaller) GetRemainingSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "getRemainingSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRemainingSupply is a free data retrieval call binding the contract method 0xe4b7fb73.
//
// Solidity: function getRemainingSupply() view returns(uint256)
func (_PoWPepe *PoWPepeSession) GetRemainingSupply() (*big.Int, error) {
	return _PoWPepe.Contract.GetRemainingSupply(&_PoWPepe.CallOpts)
}

// GetRemainingSupply is a free data retrieval call binding the contract method 0xe4b7fb73.
//
// Solidity: function getRemainingSupply() view returns(uint256)
func (_PoWPepe *PoWPepeCallerSession) GetRemainingSupply() (*big.Int, error) {
	return _PoWPepe.Contract.GetRemainingSupply(&_PoWPepe.CallOpts)
}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_PoWPepe *PoWPepeCaller) LimitPerMint(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "limitPerMint")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_PoWPepe *PoWPepeSession) LimitPerMint() (*big.Int, error) {
	return _PoWPepe.Contract.LimitPerMint(&_PoWPepe.CallOpts)
}

// LimitPerMint is a free data retrieval call binding the contract method 0xe2ce9f51.
//
// Solidity: function limitPerMint() view returns(uint256)
func (_PoWPepe *PoWPepeCallerSession) LimitPerMint() (*big.Int, error) {
	return _PoWPepe.Contract.LimitPerMint(&_PoWPepe.CallOpts)
}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_PoWPepe *PoWPepeCaller) MinedNonces(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "minedNonces", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_PoWPepe *PoWPepeSession) MinedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _PoWPepe.Contract.MinedNonces(&_PoWPepe.CallOpts, arg0, arg1)
}

// MinedNonces is a free data retrieval call binding the contract method 0x342a252a.
//
// Solidity: function minedNonces(address , uint256 ) view returns(bool)
func (_PoWPepe *PoWPepeCallerSession) MinedNonces(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _PoWPepe.Contract.MinedNonces(&_PoWPepe.CallOpts, arg0, arg1)
}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_PoWPepe *PoWPepeCaller) MiningLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "miningLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_PoWPepe *PoWPepeSession) MiningLimit() (*big.Int, error) {
	return _PoWPepe.Contract.MiningLimit(&_PoWPepe.CallOpts)
}

// MiningLimit is a free data retrieval call binding the contract method 0xc2651503.
//
// Solidity: function miningLimit() view returns(uint256)
func (_PoWPepe *PoWPepeCallerSession) MiningLimit() (*big.Int, error) {
	return _PoWPepe.Contract.MiningLimit(&_PoWPepe.CallOpts)
}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_PoWPepe *PoWPepeCaller) MiningTimes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "miningTimes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_PoWPepe *PoWPepeSession) MiningTimes(arg0 common.Address) (*big.Int, error) {
	return _PoWPepe.Contract.MiningTimes(&_PoWPepe.CallOpts, arg0)
}

// MiningTimes is a free data retrieval call binding the contract method 0x2719881e.
//
// Solidity: function miningTimes(address ) view returns(uint256)
func (_PoWPepe *PoWPepeCallerSession) MiningTimes(arg0 common.Address) (*big.Int, error) {
	return _PoWPepe.Contract.MiningTimes(&_PoWPepe.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PoWPepe *PoWPepeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PoWPepe *PoWPepeSession) Name() (string, error) {
	return _PoWPepe.Contract.Name(&_PoWPepe.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PoWPepe *PoWPepeCallerSession) Name() (string, error) {
	return _PoWPepe.Contract.Name(&_PoWPepe.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PoWPepe *PoWPepeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PoWPepe *PoWPepeSession) Symbol() (string, error) {
	return _PoWPepe.Contract.Symbol(&_PoWPepe.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_PoWPepe *PoWPepeCallerSession) Symbol() (string, error) {
	return _PoWPepe.Contract.Symbol(&_PoWPepe.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PoWPepe *PoWPepeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PoWPepe *PoWPepeSession) TotalSupply() (*big.Int, error) {
	return _PoWPepe.Contract.TotalSupply(&_PoWPepe.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_PoWPepe *PoWPepeCallerSession) TotalSupply() (*big.Int, error) {
	return _PoWPepe.Contract.TotalSupply(&_PoWPepe.CallOpts)
}

// TotalSupplyCap is a free data retrieval call binding the contract method 0xbb102aea.
//
// Solidity: function totalSupplyCap() view returns(uint256)
func (_PoWPepe *PoWPepeCaller) TotalSupplyCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoWPepe.contract.Call(opts, &out, "totalSupplyCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupplyCap is a free data retrieval call binding the contract method 0xbb102aea.
//
// Solidity: function totalSupplyCap() view returns(uint256)
func (_PoWPepe *PoWPepeSession) TotalSupplyCap() (*big.Int, error) {
	return _PoWPepe.Contract.TotalSupplyCap(&_PoWPepe.CallOpts)
}

// TotalSupplyCap is a free data retrieval call binding the contract method 0xbb102aea.
//
// Solidity: function totalSupplyCap() view returns(uint256)
func (_PoWPepe *PoWPepeCallerSession) TotalSupplyCap() (*big.Int, error) {
	return _PoWPepe.Contract.TotalSupplyCap(&_PoWPepe.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PoWPepe *PoWPepeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PoWPepe.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PoWPepe *PoWPepeSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PoWPepe.Contract.Approve(&_PoWPepe.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_PoWPepe *PoWPepeTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PoWPepe.Contract.Approve(&_PoWPepe.TransactOpts, spender, value)
}

// Mine is a paid mutator transaction binding the contract method 0x4d474898.
//
// Solidity: function mine(uint256 nonce) returns()
func (_PoWPepe *PoWPepeTransactor) Mine(opts *bind.TransactOpts, nonce *big.Int) (*types.Transaction, error) {
	return _PoWPepe.contract.Transact(opts, "mine", nonce)
}

// Mine is a paid mutator transaction binding the contract method 0x4d474898.
//
// Solidity: function mine(uint256 nonce) returns()
func (_PoWPepe *PoWPepeSession) Mine(nonce *big.Int) (*types.Transaction, error) {
	return _PoWPepe.Contract.Mine(&_PoWPepe.TransactOpts, nonce)
}

// Mine is a paid mutator transaction binding the contract method 0x4d474898.
//
// Solidity: function mine(uint256 nonce) returns()
func (_PoWPepe *PoWPepeTransactorSession) Mine(nonce *big.Int) (*types.Transaction, error) {
	return _PoWPepe.Contract.Mine(&_PoWPepe.TransactOpts, nonce)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PoWPepe *PoWPepeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoWPepe.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PoWPepe *PoWPepeSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoWPepe.Contract.Transfer(&_PoWPepe.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_PoWPepe *PoWPepeTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PoWPepe.Contract.Transfer(&_PoWPepe.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PoWPepe *PoWPepeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PoWPepe.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PoWPepe *PoWPepeSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PoWPepe.Contract.TransferFrom(&_PoWPepe.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_PoWPepe *PoWPepeTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PoWPepe.Contract.TransferFrom(&_PoWPepe.TransactOpts, from, to, value)
}

// PoWPepeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PoWPepe contract.
type PoWPepeApprovalIterator struct {
	Event *PoWPepeApproval // Event containing the contract specifics and raw log

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
func (it *PoWPepeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoWPepeApproval)
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
		it.Event = new(PoWPepeApproval)
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
func (it *PoWPepeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoWPepeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoWPepeApproval represents a Approval event raised by the PoWPepe contract.
type PoWPepeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PoWPepe *PoWPepeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PoWPepeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PoWPepe.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PoWPepeApprovalIterator{contract: _PoWPepe.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PoWPepe *PoWPepeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PoWPepeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PoWPepe.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoWPepeApproval)
				if err := _PoWPepe.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_PoWPepe *PoWPepeFilterer) ParseApproval(log types.Log) (*PoWPepeApproval, error) {
	event := new(PoWPepeApproval)
	if err := _PoWPepe.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoWPepeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PoWPepe contract.
type PoWPepeTransferIterator struct {
	Event *PoWPepeTransfer // Event containing the contract specifics and raw log

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
func (it *PoWPepeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoWPepeTransfer)
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
		it.Event = new(PoWPepeTransfer)
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
func (it *PoWPepeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoWPepeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoWPepeTransfer represents a Transfer event raised by the PoWPepe contract.
type PoWPepeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PoWPepe *PoWPepeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PoWPepeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoWPepe.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PoWPepeTransferIterator{contract: _PoWPepe.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PoWPepe *PoWPepeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PoWPepeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PoWPepe.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoWPepeTransfer)
				if err := _PoWPepe.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_PoWPepe *PoWPepeFilterer) ParseTransfer(log types.Log) (*PoWPepeTransfer, error) {
	event := new(PoWPepeTransfer)
	if err := _PoWPepe.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
