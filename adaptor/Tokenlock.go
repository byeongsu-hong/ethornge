// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adaptor

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// TokenlockABI is the input ABI used to generate the binding from.
const TokenlockABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setLocker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"started\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_buyer\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"setDistributor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"releaseTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"locker\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_distributor\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// Tokenlock is an auto generated Go binding around an Ethereum contract.
type Tokenlock struct {
	TokenlockCaller     // Read-only binding to the contract
	TokenlockTransactor // Write-only binding to the contract
	TokenlockFilterer   // Log filterer for contract events
}

// TokenlockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenlockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenlockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenlockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenlockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenlockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenlockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenlockSession struct {
	Contract     *Tokenlock        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenlockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenlockCallerSession struct {
	Contract *TokenlockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenlockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenlockTransactorSession struct {
	Contract     *TokenlockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenlockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenlockRaw struct {
	Contract *Tokenlock // Generic contract binding to access the raw methods on
}

// TokenlockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenlockCallerRaw struct {
	Contract *TokenlockCaller // Generic read-only contract binding to access the raw methods on
}

// TokenlockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenlockTransactorRaw struct {
	Contract *TokenlockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenlock creates a new instance of Tokenlock, bound to a specific deployed contract.
func NewTokenlock(address common.Address, backend bind.ContractBackend) (*Tokenlock, error) {
	contract, err := bindTokenlock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenlock{TokenlockCaller: TokenlockCaller{contract: contract}, TokenlockTransactor: TokenlockTransactor{contract: contract}, TokenlockFilterer: TokenlockFilterer{contract: contract}}, nil
}

// NewTokenlockCaller creates a new read-only instance of Tokenlock, bound to a specific deployed contract.
func NewTokenlockCaller(address common.Address, caller bind.ContractCaller) (*TokenlockCaller, error) {
	contract, err := bindTokenlock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenlockCaller{contract: contract}, nil
}

// NewTokenlockTransactor creates a new write-only instance of Tokenlock, bound to a specific deployed contract.
func NewTokenlockTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenlockTransactor, error) {
	contract, err := bindTokenlock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenlockTransactor{contract: contract}, nil
}

// NewTokenlockFilterer creates a new log filterer instance of Tokenlock, bound to a specific deployed contract.
func NewTokenlockFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenlockFilterer, error) {
	contract, err := bindTokenlock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenlockFilterer{contract: contract}, nil
}

// bindTokenlock binds a generic wrapper to an already deployed contract.
func bindTokenlock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenlockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenlock *TokenlockRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tokenlock.Contract.TokenlockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenlock *TokenlockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenlock.Contract.TokenlockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenlock *TokenlockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenlock.Contract.TokenlockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenlock *TokenlockCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Tokenlock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenlock *TokenlockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenlock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenlock *TokenlockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenlock.Contract.contract.Transact(opts, method, params...)
}

// Token is a free data retrieval call binding the contract method 0xc2412676.
//
// Solidity: function Token() constant returns(address)
func (_Tokenlock *TokenlockCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tokenlock.contract.Call(opts, out, "Token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xc2412676.
//
// Solidity: function Token() constant returns(address)
func (_Tokenlock *TokenlockSession) Token() (common.Address, error) {
	return _Tokenlock.Contract.Token(&_Tokenlock.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xc2412676.
//
// Solidity: function Token() constant returns(address)
func (_Tokenlock *TokenlockCallerSession) Token() (common.Address, error) {
	return _Tokenlock.Contract.Token(&_Tokenlock.CallOpts)
}

// Buyers is a free data retrieval call binding the contract method 0x97a993aa.
//
// Solidity: function buyers( address) constant returns(uint256)
func (_Tokenlock *TokenlockCaller) Buyers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tokenlock.contract.Call(opts, out, "buyers", arg0)
	return *ret0, err
}

// Buyers is a free data retrieval call binding the contract method 0x97a993aa.
//
// Solidity: function buyers( address) constant returns(uint256)
func (_Tokenlock *TokenlockSession) Buyers(arg0 common.Address) (*big.Int, error) {
	return _Tokenlock.Contract.Buyers(&_Tokenlock.CallOpts, arg0)
}

// Buyers is a free data retrieval call binding the contract method 0x97a993aa.
//
// Solidity: function buyers( address) constant returns(uint256)
func (_Tokenlock *TokenlockCallerSession) Buyers(arg0 common.Address) (*big.Int, error) {
	return _Tokenlock.Contract.Buyers(&_Tokenlock.CallOpts, arg0)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() constant returns(address)
func (_Tokenlock *TokenlockCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tokenlock.contract.Call(opts, out, "distributor")
	return *ret0, err
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() constant returns(address)
func (_Tokenlock *TokenlockSession) Distributor() (common.Address, error) {
	return _Tokenlock.Contract.Distributor(&_Tokenlock.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() constant returns(address)
func (_Tokenlock *TokenlockCallerSession) Distributor() (common.Address, error) {
	return _Tokenlock.Contract.Distributor(&_Tokenlock.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(_buyer address) constant returns(uint256)
func (_Tokenlock *TokenlockCaller) GetBalance(opts *bind.CallOpts, _buyer common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tokenlock.contract.Call(opts, out, "getBalance", _buyer)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(_buyer address) constant returns(uint256)
func (_Tokenlock *TokenlockSession) GetBalance(_buyer common.Address) (*big.Int, error) {
	return _Tokenlock.Contract.GetBalance(&_Tokenlock.CallOpts, _buyer)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(_buyer address) constant returns(uint256)
func (_Tokenlock *TokenlockCallerSession) GetBalance(_buyer common.Address) (*big.Int, error) {
	return _Tokenlock.Contract.GetBalance(&_Tokenlock.CallOpts, _buyer)
}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() constant returns(address)
func (_Tokenlock *TokenlockCaller) Locker(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tokenlock.contract.Call(opts, out, "locker")
	return *ret0, err
}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() constant returns(address)
func (_Tokenlock *TokenlockSession) Locker() (common.Address, error) {
	return _Tokenlock.Contract.Locker(&_Tokenlock.CallOpts)
}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() constant returns(address)
func (_Tokenlock *TokenlockCallerSession) Locker() (common.Address, error) {
	return _Tokenlock.Contract.Locker(&_Tokenlock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tokenlock *TokenlockCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Tokenlock.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tokenlock *TokenlockSession) Owner() (common.Address, error) {
	return _Tokenlock.Contract.Owner(&_Tokenlock.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Tokenlock *TokenlockCallerSession) Owner() (common.Address, error) {
	return _Tokenlock.Contract.Owner(&_Tokenlock.CallOpts)
}

// ReleaseTime is a free data retrieval call binding the contract method 0xb91d4001.
//
// Solidity: function releaseTime() constant returns(uint256)
func (_Tokenlock *TokenlockCaller) ReleaseTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Tokenlock.contract.Call(opts, out, "releaseTime")
	return *ret0, err
}

// ReleaseTime is a free data retrieval call binding the contract method 0xb91d4001.
//
// Solidity: function releaseTime() constant returns(uint256)
func (_Tokenlock *TokenlockSession) ReleaseTime() (*big.Int, error) {
	return _Tokenlock.Contract.ReleaseTime(&_Tokenlock.CallOpts)
}

// ReleaseTime is a free data retrieval call binding the contract method 0xb91d4001.
//
// Solidity: function releaseTime() constant returns(uint256)
func (_Tokenlock *TokenlockCallerSession) ReleaseTime() (*big.Int, error) {
	return _Tokenlock.Contract.ReleaseTime(&_Tokenlock.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() constant returns(bool)
func (_Tokenlock *TokenlockCaller) Started(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Tokenlock.contract.Call(opts, out, "started")
	return *ret0, err
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() constant returns(bool)
func (_Tokenlock *TokenlockSession) Started() (bool, error) {
	return _Tokenlock.Contract.Started(&_Tokenlock.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() constant returns(bool)
func (_Tokenlock *TokenlockCallerSession) Started() (bool, error) {
	return _Tokenlock.Contract.Started(&_Tokenlock.CallOpts)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(_buyer address, _amount uint256) returns()
func (_Tokenlock *TokenlockTransactor) Lock(opts *bind.TransactOpts, _buyer common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tokenlock.contract.Transact(opts, "lock", _buyer, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(_buyer address, _amount uint256) returns()
func (_Tokenlock *TokenlockSession) Lock(_buyer common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tokenlock.Contract.Lock(&_Tokenlock.TransactOpts, _buyer, _amount)
}

// Lock is a paid mutator transaction binding the contract method 0x282d3fdf.
//
// Solidity: function lock(_buyer address, _amount uint256) returns()
func (_Tokenlock *TokenlockTransactorSession) Lock(_buyer common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Tokenlock.Contract.Lock(&_Tokenlock.TransactOpts, _buyer, _amount)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_buyer address) returns()
func (_Tokenlock *TokenlockTransactor) Release(opts *bind.TransactOpts, _buyer common.Address) (*types.Transaction, error) {
	return _Tokenlock.contract.Transact(opts, "release", _buyer)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_buyer address) returns()
func (_Tokenlock *TokenlockSession) Release(_buyer common.Address) (*types.Transaction, error) {
	return _Tokenlock.Contract.Release(&_Tokenlock.TransactOpts, _buyer)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_buyer address) returns()
func (_Tokenlock *TokenlockTransactorSession) Release(_buyer common.Address) (*types.Transaction, error) {
	return _Tokenlock.Contract.Release(&_Tokenlock.TransactOpts, _buyer)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenlock *TokenlockTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenlock.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenlock *TokenlockSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenlock.Contract.RenounceOwnership(&_Tokenlock.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tokenlock *TokenlockTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tokenlock.Contract.RenounceOwnership(&_Tokenlock.TransactOpts)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(_addr address) returns()
func (_Tokenlock *TokenlockTransactor) SetDistributor(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Tokenlock.contract.Transact(opts, "setDistributor", _addr)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(_addr address) returns()
func (_Tokenlock *TokenlockSession) SetDistributor(_addr common.Address) (*types.Transaction, error) {
	return _Tokenlock.Contract.SetDistributor(&_Tokenlock.TransactOpts, _addr)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(_addr address) returns()
func (_Tokenlock *TokenlockTransactorSession) SetDistributor(_addr common.Address) (*types.Transaction, error) {
	return _Tokenlock.Contract.SetDistributor(&_Tokenlock.TransactOpts, _addr)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(_addr address) returns()
func (_Tokenlock *TokenlockTransactor) SetLocker(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _Tokenlock.contract.Transact(opts, "setLocker", _addr)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(_addr address) returns()
func (_Tokenlock *TokenlockSession) SetLocker(_addr common.Address) (*types.Transaction, error) {
	return _Tokenlock.Contract.SetLocker(&_Tokenlock.TransactOpts, _addr)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(_addr address) returns()
func (_Tokenlock *TokenlockTransactorSession) SetLocker(_addr common.Address) (*types.Transaction, error) {
	return _Tokenlock.Contract.SetLocker(&_Tokenlock.TransactOpts, _addr)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Tokenlock *TokenlockTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenlock.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Tokenlock *TokenlockSession) Start() (*types.Transaction, error) {
	return _Tokenlock.Contract.Start(&_Tokenlock.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Tokenlock *TokenlockTransactorSession) Start() (*types.Transaction, error) {
	return _Tokenlock.Contract.Start(&_Tokenlock.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Tokenlock *TokenlockTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tokenlock.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Tokenlock *TokenlockSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenlock.Contract.TransferOwnership(&_Tokenlock.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Tokenlock *TokenlockTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tokenlock.Contract.TransferOwnership(&_Tokenlock.TransactOpts, newOwner)
}

// TokenlockOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Tokenlock contract.
type TokenlockOwnershipRenouncedIterator struct {
	Event *TokenlockOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *TokenlockOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenlockOwnershipRenounced)
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
		it.Event = new(TokenlockOwnershipRenounced)
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
func (it *TokenlockOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenlockOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenlockOwnershipRenounced represents a OwnershipRenounced event raised by the Tokenlock contract.
type TokenlockOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_Tokenlock *TokenlockFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*TokenlockOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Tokenlock.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenlockOwnershipRenouncedIterator{contract: _Tokenlock.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_Tokenlock *TokenlockFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *TokenlockOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Tokenlock.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenlockOwnershipRenounced)
				if err := _Tokenlock.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// TokenlockOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tokenlock contract.
type TokenlockOwnershipTransferredIterator struct {
	Event *TokenlockOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenlockOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenlockOwnershipTransferred)
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
		it.Event = new(TokenlockOwnershipTransferred)
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
func (it *TokenlockOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenlockOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenlockOwnershipTransferred represents a OwnershipTransferred event raised by the Tokenlock contract.
type TokenlockOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Tokenlock *TokenlockFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenlockOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenlock.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenlockOwnershipTransferredIterator{contract: _Tokenlock.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Tokenlock *TokenlockFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenlockOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tokenlock.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenlockOwnershipTransferred)
				if err := _Tokenlock.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
