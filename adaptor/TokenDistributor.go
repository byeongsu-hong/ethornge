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

// TokenDistributorABI is the input ABI used to generate the binding from.
const TokenDistributorABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"x\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Lock\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Genesis\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"y\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addrs\",\"type\":\"address[]\"}],\"name\":\"releaseMany\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_ratioX\",\"type\":\"uint256\"},{\"name\":\"_ratioY\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_genesis\",\"type\":\"address\"},{\"name\":\"_tokenLock\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_safeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_lockAmount\",\"type\":\"uint256\"}],\"name\":\"Release\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// TokenDistributor is an auto generated Go binding around an Ethereum contract.
type TokenDistributor struct {
	TokenDistributorCaller     // Read-only binding to the contract
	TokenDistributorTransactor // Write-only binding to the contract
	TokenDistributorFilterer   // Log filterer for contract events
}

// TokenDistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenDistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenDistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenDistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenDistributorSession struct {
	Contract     *TokenDistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenDistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenDistributorCallerSession struct {
	Contract *TokenDistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TokenDistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenDistributorTransactorSession struct {
	Contract     *TokenDistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TokenDistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenDistributorRaw struct {
	Contract *TokenDistributor // Generic contract binding to access the raw methods on
}

// TokenDistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenDistributorCallerRaw struct {
	Contract *TokenDistributorCaller // Generic read-only contract binding to access the raw methods on
}

// TokenDistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenDistributorTransactorRaw struct {
	Contract *TokenDistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenDistributor creates a new instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributor(address common.Address, backend bind.ContractBackend) (*TokenDistributor, error) {
	contract, err := bindTokenDistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenDistributor{TokenDistributorCaller: TokenDistributorCaller{contract: contract}, TokenDistributorTransactor: TokenDistributorTransactor{contract: contract}, TokenDistributorFilterer: TokenDistributorFilterer{contract: contract}}, nil
}

// NewTokenDistributorCaller creates a new read-only instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributorCaller(address common.Address, caller bind.ContractCaller) (*TokenDistributorCaller, error) {
	contract, err := bindTokenDistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorCaller{contract: contract}, nil
}

// NewTokenDistributorTransactor creates a new write-only instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenDistributorTransactor, error) {
	contract, err := bindTokenDistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorTransactor{contract: contract}, nil
}

// NewTokenDistributorFilterer creates a new log filterer instance of TokenDistributor, bound to a specific deployed contract.
func NewTokenDistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenDistributorFilterer, error) {
	contract, err := bindTokenDistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorFilterer{contract: contract}, nil
}

// bindTokenDistributor binds a generic wrapper to an already deployed contract.
func bindTokenDistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenDistributorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenDistributor *TokenDistributorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenDistributor.Contract.TokenDistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenDistributor *TokenDistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenDistributor.Contract.TokenDistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenDistributor *TokenDistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenDistributor.Contract.TokenDistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenDistributor *TokenDistributorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenDistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenDistributor *TokenDistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenDistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenDistributor *TokenDistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenDistributor.Contract.contract.Transact(opts, method, params...)
}

// Genesis is a free data retrieval call binding the contract method 0x6bf6eaff.
//
// Solidity: function Genesis() constant returns(address)
func (_TokenDistributor *TokenDistributorCaller) Genesis(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenDistributor.contract.Call(opts, out, "Genesis")
	return *ret0, err
}

// Genesis is a free data retrieval call binding the contract method 0x6bf6eaff.
//
// Solidity: function Genesis() constant returns(address)
func (_TokenDistributor *TokenDistributorSession) Genesis() (common.Address, error) {
	return _TokenDistributor.Contract.Genesis(&_TokenDistributor.CallOpts)
}

// Genesis is a free data retrieval call binding the contract method 0x6bf6eaff.
//
// Solidity: function Genesis() constant returns(address)
func (_TokenDistributor *TokenDistributorCallerSession) Genesis() (common.Address, error) {
	return _TokenDistributor.Contract.Genesis(&_TokenDistributor.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0x46620e39.
//
// Solidity: function Lock() constant returns(address)
func (_TokenDistributor *TokenDistributorCaller) Lock(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenDistributor.contract.Call(opts, out, "Lock")
	return *ret0, err
}

// Lock is a free data retrieval call binding the contract method 0x46620e39.
//
// Solidity: function Lock() constant returns(address)
func (_TokenDistributor *TokenDistributorSession) Lock() (common.Address, error) {
	return _TokenDistributor.Contract.Lock(&_TokenDistributor.CallOpts)
}

// Lock is a free data retrieval call binding the contract method 0x46620e39.
//
// Solidity: function Lock() constant returns(address)
func (_TokenDistributor *TokenDistributorCallerSession) Lock() (common.Address, error) {
	return _TokenDistributor.Contract.Lock(&_TokenDistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xc2412676.
//
// Solidity: function Token() constant returns(address)
func (_TokenDistributor *TokenDistributorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenDistributor.contract.Call(opts, out, "Token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xc2412676.
//
// Solidity: function Token() constant returns(address)
func (_TokenDistributor *TokenDistributorSession) Token() (common.Address, error) {
	return _TokenDistributor.Contract.Token(&_TokenDistributor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xc2412676.
//
// Solidity: function Token() constant returns(address)
func (_TokenDistributor *TokenDistributorCallerSession) Token() (common.Address, error) {
	return _TokenDistributor.Contract.Token(&_TokenDistributor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenDistributor *TokenDistributorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenDistributor.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenDistributor *TokenDistributorSession) Owner() (common.Address, error) {
	return _TokenDistributor.Contract.Owner(&_TokenDistributor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenDistributor *TokenDistributorCallerSession) Owner() (common.Address, error) {
	return _TokenDistributor.Contract.Owner(&_TokenDistributor.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() constant returns(uint256)
func (_TokenDistributor *TokenDistributorCaller) X(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenDistributor.contract.Call(opts, out, "x")
	return *ret0, err
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() constant returns(uint256)
func (_TokenDistributor *TokenDistributorSession) X() (*big.Int, error) {
	return _TokenDistributor.Contract.X(&_TokenDistributor.CallOpts)
}

// X is a free data retrieval call binding the contract method 0x0c55699c.
//
// Solidity: function x() constant returns(uint256)
func (_TokenDistributor *TokenDistributorCallerSession) X() (*big.Int, error) {
	return _TokenDistributor.Contract.X(&_TokenDistributor.CallOpts)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() constant returns(uint256)
func (_TokenDistributor *TokenDistributorCaller) Y(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenDistributor.contract.Call(opts, out, "y")
	return *ret0, err
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() constant returns(uint256)
func (_TokenDistributor *TokenDistributorSession) Y() (*big.Int, error) {
	return _TokenDistributor.Contract.Y(&_TokenDistributor.CallOpts)
}

// Y is a free data retrieval call binding the contract method 0xa56dfe4a.
//
// Solidity: function y() constant returns(uint256)
func (_TokenDistributor *TokenDistributorCallerSession) Y() (*big.Int, error) {
	return _TokenDistributor.Contract.Y(&_TokenDistributor.CallOpts)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_addr address) returns()
func (_TokenDistributor *TokenDistributorTransactor) Release(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _TokenDistributor.contract.Transact(opts, "release", _addr)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_addr address) returns()
func (_TokenDistributor *TokenDistributorSession) Release(_addr common.Address) (*types.Transaction, error) {
	return _TokenDistributor.Contract.Release(&_TokenDistributor.TransactOpts, _addr)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_addr address) returns()
func (_TokenDistributor *TokenDistributorTransactorSession) Release(_addr common.Address) (*types.Transaction, error) {
	return _TokenDistributor.Contract.Release(&_TokenDistributor.TransactOpts, _addr)
}

// ReleaseMany is a paid mutator transaction binding the contract method 0xe52f64ce.
//
// Solidity: function releaseMany(_addrs address[]) returns()
func (_TokenDistributor *TokenDistributorTransactor) ReleaseMany(opts *bind.TransactOpts, _addrs []common.Address) (*types.Transaction, error) {
	return _TokenDistributor.contract.Transact(opts, "releaseMany", _addrs)
}

// ReleaseMany is a paid mutator transaction binding the contract method 0xe52f64ce.
//
// Solidity: function releaseMany(_addrs address[]) returns()
func (_TokenDistributor *TokenDistributorSession) ReleaseMany(_addrs []common.Address) (*types.Transaction, error) {
	return _TokenDistributor.Contract.ReleaseMany(&_TokenDistributor.TransactOpts, _addrs)
}

// ReleaseMany is a paid mutator transaction binding the contract method 0xe52f64ce.
//
// Solidity: function releaseMany(_addrs address[]) returns()
func (_TokenDistributor *TokenDistributorTransactorSession) ReleaseMany(_addrs []common.Address) (*types.Transaction, error) {
	return _TokenDistributor.Contract.ReleaseMany(&_TokenDistributor.TransactOpts, _addrs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenDistributor *TokenDistributorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenDistributor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenDistributor *TokenDistributorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenDistributor.Contract.RenounceOwnership(&_TokenDistributor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TokenDistributor *TokenDistributorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TokenDistributor.Contract.RenounceOwnership(&_TokenDistributor.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenDistributor *TokenDistributorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TokenDistributor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenDistributor *TokenDistributorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenDistributor.Contract.TransferOwnership(&_TokenDistributor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_TokenDistributor *TokenDistributorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TokenDistributor.Contract.TransferOwnership(&_TokenDistributor.TransactOpts, newOwner)
}

// TokenDistributorOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the TokenDistributor contract.
type TokenDistributorOwnershipRenouncedIterator struct {
	Event *TokenDistributorOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *TokenDistributorOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenDistributorOwnershipRenounced)
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
		it.Event = new(TokenDistributorOwnershipRenounced)
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
func (it *TokenDistributorOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenDistributorOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenDistributorOwnershipRenounced represents a OwnershipRenounced event raised by the TokenDistributor contract.
type TokenDistributorOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_TokenDistributor *TokenDistributorFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*TokenDistributorOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _TokenDistributor.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorOwnershipRenouncedIterator{contract: _TokenDistributor.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_TokenDistributor *TokenDistributorFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *TokenDistributorOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _TokenDistributor.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenDistributorOwnershipRenounced)
				if err := _TokenDistributor.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// TokenDistributorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TokenDistributor contract.
type TokenDistributorOwnershipTransferredIterator struct {
	Event *TokenDistributorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TokenDistributorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenDistributorOwnershipTransferred)
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
		it.Event = new(TokenDistributorOwnershipTransferred)
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
func (it *TokenDistributorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenDistributorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenDistributorOwnershipTransferred represents a OwnershipTransferred event raised by the TokenDistributor contract.
type TokenDistributorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_TokenDistributor *TokenDistributorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TokenDistributorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenDistributor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorOwnershipTransferredIterator{contract: _TokenDistributor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_TokenDistributor *TokenDistributorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TokenDistributorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TokenDistributor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenDistributorOwnershipTransferred)
				if err := _TokenDistributor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// TokenDistributorReleaseIterator is returned from FilterRelease and is used to iterate over the raw logs and unpacked data for Release events raised by the TokenDistributor contract.
type TokenDistributorReleaseIterator struct {
	Event *TokenDistributorRelease // Event containing the contract specifics and raw log

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
func (it *TokenDistributorReleaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenDistributorRelease)
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
		it.Event = new(TokenDistributorRelease)
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
func (it *TokenDistributorReleaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenDistributorReleaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenDistributorRelease represents a Release event raised by the TokenDistributor contract.
type TokenDistributorRelease struct {
	To         common.Address
	SafeAmount *big.Int
	LockAmount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRelease is a free log retrieval operation binding the contract event 0x2c57dec1db0095a6b800c2698d5bbceef2c180c6ac43429769a719658983f677.
//
// Solidity: event Release(_to indexed address, _safeAmount uint256, _lockAmount uint256)
func (_TokenDistributor *TokenDistributorFilterer) FilterRelease(opts *bind.FilterOpts, _to []common.Address) (*TokenDistributorReleaseIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TokenDistributor.contract.FilterLogs(opts, "Release", _toRule)
	if err != nil {
		return nil, err
	}
	return &TokenDistributorReleaseIterator{contract: _TokenDistributor.contract, event: "Release", logs: logs, sub: sub}, nil
}

// WatchRelease is a free log subscription operation binding the contract event 0x2c57dec1db0095a6b800c2698d5bbceef2c180c6ac43429769a719658983f677.
//
// Solidity: event Release(_to indexed address, _safeAmount uint256, _lockAmount uint256)
func (_TokenDistributor *TokenDistributorFilterer) WatchRelease(opts *bind.WatchOpts, sink chan<- *TokenDistributorRelease, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _TokenDistributor.contract.WatchLogs(opts, "Release", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenDistributorRelease)
				if err := _TokenDistributor.contract.UnpackLog(event, "Release", log); err != nil {
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
