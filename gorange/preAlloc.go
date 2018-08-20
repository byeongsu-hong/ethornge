// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gorange

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// PreAllocABI is the input ABI used to generate the binding from.
const PreAllocABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addrs\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"distribute\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// PreAllocBin is the compiled bytecode used for deploying new contracts.
const PreAllocBin = `0x608060405234801561001057600080fd5b506101bf806100206000396000f3006080604052600436106100405763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416631826c1198114610045575b600080fd5b6040805160206004803580820135838102808601850190965280855261008f9536959394602494938501929182918501908490808284375094975050933594506100919350505050565b005b600034828451021115151561010757604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f556e73756666696369656e742042616c616e6365000000000000000000000000604482015290519081900360640190fd5b5060005b82518167ffffffffffffffff16101561018e57828167ffffffffffffffff1681518110151561013657fe5b9060200190602002015173ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015610185573d6000803e3d6000fd5b5060010161010b565b5050505600a165627a7a723058208eb1416356232f602ab81d26c24f711db178bcb5dc156af5046ee2e9086d5c880029`

// DeployPreAlloc deploys a new Ethereum contract, binding an instance of PreAlloc to it.
func DeployPreAlloc(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PreAlloc, error) {
	parsed, err := abi.JSON(strings.NewReader(PreAllocABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PreAllocBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PreAlloc{PreAllocCaller: PreAllocCaller{contract: contract}, PreAllocTransactor: PreAllocTransactor{contract: contract}, PreAllocFilterer: PreAllocFilterer{contract: contract}}, nil
}

// PreAlloc is an auto generated Go binding around an Ethereum contract.
type PreAlloc struct {
	PreAllocCaller     // Read-only binding to the contract
	PreAllocTransactor // Write-only binding to the contract
	PreAllocFilterer   // Log filterer for contract events
}

// PreAllocCaller is an auto generated read-only Go binding around an Ethereum contract.
type PreAllocCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreAllocTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PreAllocTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreAllocFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PreAllocFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PreAllocSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PreAllocSession struct {
	Contract     *PreAlloc         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PreAllocCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PreAllocCallerSession struct {
	Contract *PreAllocCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PreAllocTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PreAllocTransactorSession struct {
	Contract     *PreAllocTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PreAllocRaw is an auto generated low-level Go binding around an Ethereum contract.
type PreAllocRaw struct {
	Contract *PreAlloc // Generic contract binding to access the raw methods on
}

// PreAllocCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PreAllocCallerRaw struct {
	Contract *PreAllocCaller // Generic read-only contract binding to access the raw methods on
}

// PreAllocTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PreAllocTransactorRaw struct {
	Contract *PreAllocTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPreAlloc creates a new instance of PreAlloc, bound to a specific deployed contract.
func NewPreAlloc(address common.Address, backend bind.ContractBackend) (*PreAlloc, error) {
	contract, err := bindPreAlloc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PreAlloc{PreAllocCaller: PreAllocCaller{contract: contract}, PreAllocTransactor: PreAllocTransactor{contract: contract}, PreAllocFilterer: PreAllocFilterer{contract: contract}}, nil
}

// NewPreAllocCaller creates a new read-only instance of PreAlloc, bound to a specific deployed contract.
func NewPreAllocCaller(address common.Address, caller bind.ContractCaller) (*PreAllocCaller, error) {
	contract, err := bindPreAlloc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PreAllocCaller{contract: contract}, nil
}

// NewPreAllocTransactor creates a new write-only instance of PreAlloc, bound to a specific deployed contract.
func NewPreAllocTransactor(address common.Address, transactor bind.ContractTransactor) (*PreAllocTransactor, error) {
	contract, err := bindPreAlloc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PreAllocTransactor{contract: contract}, nil
}

// NewPreAllocFilterer creates a new log filterer instance of PreAlloc, bound to a specific deployed contract.
func NewPreAllocFilterer(address common.Address, filterer bind.ContractFilterer) (*PreAllocFilterer, error) {
	contract, err := bindPreAlloc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PreAllocFilterer{contract: contract}, nil
}

// bindPreAlloc binds a generic wrapper to an already deployed contract.
func bindPreAlloc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PreAllocABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreAlloc *PreAllocRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PreAlloc.Contract.PreAllocCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreAlloc *PreAllocRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreAlloc.Contract.PreAllocTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreAlloc *PreAllocRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreAlloc.Contract.PreAllocTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PreAlloc *PreAllocCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PreAlloc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PreAlloc *PreAllocTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PreAlloc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PreAlloc *PreAllocTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PreAlloc.Contract.contract.Transact(opts, method, params...)
}

// Distribute is a paid mutator transaction binding the contract method 0x1826c119.
//
// Solidity: function distribute(_addrs address[], _amount uint256) returns()
func (_PreAlloc *PreAllocTransactor) Distribute(opts *bind.TransactOpts, _addrs []common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PreAlloc.contract.Transact(opts, "distribute", _addrs, _amount)
}

// Distribute is a paid mutator transaction binding the contract method 0x1826c119.
//
// Solidity: function distribute(_addrs address[], _amount uint256) returns()
func (_PreAlloc *PreAllocSession) Distribute(_addrs []common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PreAlloc.Contract.Distribute(&_PreAlloc.TransactOpts, _addrs, _amount)
}

// Distribute is a paid mutator transaction binding the contract method 0x1826c119.
//
// Solidity: function distribute(_addrs address[], _amount uint256) returns()
func (_PreAlloc *PreAllocTransactorSession) Distribute(_addrs []common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PreAlloc.Contract.Distribute(&_PreAlloc.TransactOpts, _addrs, _amount)
}
