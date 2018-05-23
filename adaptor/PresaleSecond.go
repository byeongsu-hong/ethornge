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

// PresaleSecondABI is the input ABI used to generate the binding from.
const PresaleSecondABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exceed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxcap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ignited\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"setDistributor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_whitelist\",\"type\":\"address\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"List\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"extinguish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"setWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"collect\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ignite\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_maxcap\",\"type\":\"uint256\"},{\"name\":\"_exceed\",\"type\":\"uint256\"},{\"name\":\"_minimum\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_distributor\",\"type\":\"address\"},{\"name\":\"_whitelist\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"Change\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resume\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Ignite\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Extinguish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_purchased\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_refund\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"Purchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Release\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PresaleSecond is an auto generated Go binding around an Ethereum contract.
type PresaleSecond struct {
	PresaleSecondCaller     // Read-only binding to the contract
	PresaleSecondTransactor // Write-only binding to the contract
	PresaleSecondFilterer   // Log filterer for contract events
}

// PresaleSecondCaller is an auto generated read-only Go binding around an Ethereum contract.
type PresaleSecondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresaleSecondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PresaleSecondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresaleSecondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PresaleSecondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresaleSecondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PresaleSecondSession struct {
	Contract     *PresaleSecond    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PresaleSecondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PresaleSecondCallerSession struct {
	Contract *PresaleSecondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PresaleSecondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PresaleSecondTransactorSession struct {
	Contract     *PresaleSecondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PresaleSecondRaw is an auto generated low-level Go binding around an Ethereum contract.
type PresaleSecondRaw struct {
	Contract *PresaleSecond // Generic contract binding to access the raw methods on
}

// PresaleSecondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PresaleSecondCallerRaw struct {
	Contract *PresaleSecondCaller // Generic read-only contract binding to access the raw methods on
}

// PresaleSecondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PresaleSecondTransactorRaw struct {
	Contract *PresaleSecondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPresaleSecond creates a new instance of PresaleSecond, bound to a specific deployed contract.
func NewPresaleSecond(address common.Address, backend bind.ContractBackend) (*PresaleSecond, error) {
	contract, err := bindPresaleSecond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PresaleSecond{PresaleSecondCaller: PresaleSecondCaller{contract: contract}, PresaleSecondTransactor: PresaleSecondTransactor{contract: contract}, PresaleSecondFilterer: PresaleSecondFilterer{contract: contract}}, nil
}

// NewPresaleSecondCaller creates a new read-only instance of PresaleSecond, bound to a specific deployed contract.
func NewPresaleSecondCaller(address common.Address, caller bind.ContractCaller) (*PresaleSecondCaller, error) {
	contract, err := bindPresaleSecond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PresaleSecondCaller{contract: contract}, nil
}

// NewPresaleSecondTransactor creates a new write-only instance of PresaleSecond, bound to a specific deployed contract.
func NewPresaleSecondTransactor(address common.Address, transactor bind.ContractTransactor) (*PresaleSecondTransactor, error) {
	contract, err := bindPresaleSecond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PresaleSecondTransactor{contract: contract}, nil
}

// NewPresaleSecondFilterer creates a new log filterer instance of PresaleSecond, bound to a specific deployed contract.
func NewPresaleSecondFilterer(address common.Address, filterer bind.ContractFilterer) (*PresaleSecondFilterer, error) {
	contract, err := bindPresaleSecond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PresaleSecondFilterer{contract: contract}, nil
}

// bindPresaleSecond binds a generic wrapper to an already deployed contract.
func bindPresaleSecond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PresaleSecondABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PresaleSecond *PresaleSecondRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PresaleSecond.Contract.PresaleSecondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PresaleSecond *PresaleSecondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleSecond.Contract.PresaleSecondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PresaleSecond *PresaleSecondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PresaleSecond.Contract.PresaleSecondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PresaleSecond *PresaleSecondCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PresaleSecond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PresaleSecond *PresaleSecondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleSecond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PresaleSecond *PresaleSecondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PresaleSecond.Contract.contract.Transact(opts, method, params...)
}

// List is a free data retrieval call binding the contract method 0xc258ff74.
//
// Solidity: function List() constant returns(address)
func (_PresaleSecond *PresaleSecondCaller) List(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "List")
	return *ret0, err
}

// List is a free data retrieval call binding the contract method 0xc258ff74.
//
// Solidity: function List() constant returns(address)
func (_PresaleSecond *PresaleSecondSession) List() (common.Address, error) {
	return _PresaleSecond.Contract.List(&_PresaleSecond.CallOpts)
}

// List is a free data retrieval call binding the contract method 0xc258ff74.
//
// Solidity: function List() constant returns(address)
func (_PresaleSecond *PresaleSecondCallerSession) List() (common.Address, error) {
	return _PresaleSecond.Contract.List(&_PresaleSecond.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xc2412676.
//
// Solidity: function Token() constant returns(address)
func (_PresaleSecond *PresaleSecondCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "Token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xc2412676.
//
// Solidity: function Token() constant returns(address)
func (_PresaleSecond *PresaleSecondSession) Token() (common.Address, error) {
	return _PresaleSecond.Contract.Token(&_PresaleSecond.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xc2412676.
//
// Solidity: function Token() constant returns(address)
func (_PresaleSecond *PresaleSecondCallerSession) Token() (common.Address, error) {
	return _PresaleSecond.Contract.Token(&_PresaleSecond.CallOpts)
}

// Buyers is a free data retrieval call binding the contract method 0x97a993aa.
//
// Solidity: function buyers( address) constant returns(uint256)
func (_PresaleSecond *PresaleSecondCaller) Buyers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "buyers", arg0)
	return *ret0, err
}

// Buyers is a free data retrieval call binding the contract method 0x97a993aa.
//
// Solidity: function buyers( address) constant returns(uint256)
func (_PresaleSecond *PresaleSecondSession) Buyers(arg0 common.Address) (*big.Int, error) {
	return _PresaleSecond.Contract.Buyers(&_PresaleSecond.CallOpts, arg0)
}

// Buyers is a free data retrieval call binding the contract method 0x97a993aa.
//
// Solidity: function buyers( address) constant returns(uint256)
func (_PresaleSecond *PresaleSecondCallerSession) Buyers(arg0 common.Address) (*big.Int, error) {
	return _PresaleSecond.Contract.Buyers(&_PresaleSecond.CallOpts, arg0)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() constant returns(address)
func (_PresaleSecond *PresaleSecondCaller) Distributor(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "distributor")
	return *ret0, err
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() constant returns(address)
func (_PresaleSecond *PresaleSecondSession) Distributor() (common.Address, error) {
	return _PresaleSecond.Contract.Distributor(&_PresaleSecond.CallOpts)
}

// Distributor is a free data retrieval call binding the contract method 0xbfe10928.
//
// Solidity: function distributor() constant returns(address)
func (_PresaleSecond *PresaleSecondCallerSession) Distributor() (common.Address, error) {
	return _PresaleSecond.Contract.Distributor(&_PresaleSecond.CallOpts)
}

// Exceed is a free data retrieval call binding the contract method 0x0b181567.
//
// Solidity: function exceed() constant returns(uint256)
func (_PresaleSecond *PresaleSecondCaller) Exceed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "exceed")
	return *ret0, err
}

// Exceed is a free data retrieval call binding the contract method 0x0b181567.
//
// Solidity: function exceed() constant returns(uint256)
func (_PresaleSecond *PresaleSecondSession) Exceed() (*big.Int, error) {
	return _PresaleSecond.Contract.Exceed(&_PresaleSecond.CallOpts)
}

// Exceed is a free data retrieval call binding the contract method 0x0b181567.
//
// Solidity: function exceed() constant returns(uint256)
func (_PresaleSecond *PresaleSecondCallerSession) Exceed() (*big.Int, error) {
	return _PresaleSecond.Contract.Exceed(&_PresaleSecond.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() constant returns(bool)
func (_PresaleSecond *PresaleSecondCaller) Finalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "finalized")
	return *ret0, err
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() constant returns(bool)
func (_PresaleSecond *PresaleSecondSession) Finalized() (bool, error) {
	return _PresaleSecond.Contract.Finalized(&_PresaleSecond.CallOpts)
}

// Finalized is a free data retrieval call binding the contract method 0xb3f05b97.
//
// Solidity: function finalized() constant returns(bool)
func (_PresaleSecond *PresaleSecondCallerSession) Finalized() (bool, error) {
	return _PresaleSecond.Contract.Finalized(&_PresaleSecond.CallOpts)
}

// Ignited is a free data retrieval call binding the contract method 0x66092ea8.
//
// Solidity: function ignited() constant returns(bool)
func (_PresaleSecond *PresaleSecondCaller) Ignited(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "ignited")
	return *ret0, err
}

// Ignited is a free data retrieval call binding the contract method 0x66092ea8.
//
// Solidity: function ignited() constant returns(bool)
func (_PresaleSecond *PresaleSecondSession) Ignited() (bool, error) {
	return _PresaleSecond.Contract.Ignited(&_PresaleSecond.CallOpts)
}

// Ignited is a free data retrieval call binding the contract method 0x66092ea8.
//
// Solidity: function ignited() constant returns(bool)
func (_PresaleSecond *PresaleSecondCallerSession) Ignited() (bool, error) {
	return _PresaleSecond.Contract.Ignited(&_PresaleSecond.CallOpts)
}

// Maxcap is a free data retrieval call binding the contract method 0x55456f58.
//
// Solidity: function maxcap() constant returns(uint256)
func (_PresaleSecond *PresaleSecondCaller) Maxcap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "maxcap")
	return *ret0, err
}

// Maxcap is a free data retrieval call binding the contract method 0x55456f58.
//
// Solidity: function maxcap() constant returns(uint256)
func (_PresaleSecond *PresaleSecondSession) Maxcap() (*big.Int, error) {
	return _PresaleSecond.Contract.Maxcap(&_PresaleSecond.CallOpts)
}

// Maxcap is a free data retrieval call binding the contract method 0x55456f58.
//
// Solidity: function maxcap() constant returns(uint256)
func (_PresaleSecond *PresaleSecondCallerSession) Maxcap() (*big.Int, error) {
	return _PresaleSecond.Contract.Maxcap(&_PresaleSecond.CallOpts)
}

// Minimum is a free data retrieval call binding the contract method 0x52d6804d.
//
// Solidity: function minimum() constant returns(uint256)
func (_PresaleSecond *PresaleSecondCaller) Minimum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "minimum")
	return *ret0, err
}

// Minimum is a free data retrieval call binding the contract method 0x52d6804d.
//
// Solidity: function minimum() constant returns(uint256)
func (_PresaleSecond *PresaleSecondSession) Minimum() (*big.Int, error) {
	return _PresaleSecond.Contract.Minimum(&_PresaleSecond.CallOpts)
}

// Minimum is a free data retrieval call binding the contract method 0x52d6804d.
//
// Solidity: function minimum() constant returns(uint256)
func (_PresaleSecond *PresaleSecondCallerSession) Minimum() (*big.Int, error) {
	return _PresaleSecond.Contract.Minimum(&_PresaleSecond.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PresaleSecond *PresaleSecondCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PresaleSecond *PresaleSecondSession) Owner() (common.Address, error) {
	return _PresaleSecond.Contract.Owner(&_PresaleSecond.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PresaleSecond *PresaleSecondCallerSession) Owner() (common.Address, error) {
	return _PresaleSecond.Contract.Owner(&_PresaleSecond.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PresaleSecond *PresaleSecondCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PresaleSecond *PresaleSecondSession) Paused() (bool, error) {
	return _PresaleSecond.Contract.Paused(&_PresaleSecond.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PresaleSecond *PresaleSecondCallerSession) Paused() (bool, error) {
	return _PresaleSecond.Contract.Paused(&_PresaleSecond.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PresaleSecond *PresaleSecondCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PresaleSecond *PresaleSecondSession) Rate() (*big.Int, error) {
	return _PresaleSecond.Contract.Rate(&_PresaleSecond.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PresaleSecond *PresaleSecondCallerSession) Rate() (*big.Int, error) {
	return _PresaleSecond.Contract.Rate(&_PresaleSecond.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PresaleSecond *PresaleSecondCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PresaleSecond *PresaleSecondSession) Wallet() (common.Address, error) {
	return _PresaleSecond.Contract.Wallet(&_PresaleSecond.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PresaleSecond *PresaleSecondCallerSession) Wallet() (common.Address, error) {
	return _PresaleSecond.Contract.Wallet(&_PresaleSecond.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PresaleSecond *PresaleSecondCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleSecond.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PresaleSecond *PresaleSecondSession) WeiRaised() (*big.Int, error) {
	return _PresaleSecond.Contract.WeiRaised(&_PresaleSecond.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PresaleSecond *PresaleSecondCallerSession) WeiRaised() (*big.Int, error) {
	return _PresaleSecond.Contract.WeiRaised(&_PresaleSecond.CallOpts)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_PresaleSecond *PresaleSecondTransactor) Collect(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "collect")
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_PresaleSecond *PresaleSecondSession) Collect() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Collect(&_PresaleSecond.TransactOpts)
}

// Collect is a paid mutator transaction binding the contract method 0xe5225381.
//
// Solidity: function collect() returns()
func (_PresaleSecond *PresaleSecondTransactorSession) Collect() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Collect(&_PresaleSecond.TransactOpts)
}

// Extinguish is a paid mutator transaction binding the contract method 0xc54837a4.
//
// Solidity: function extinguish() returns()
func (_PresaleSecond *PresaleSecondTransactor) Extinguish(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "extinguish")
}

// Extinguish is a paid mutator transaction binding the contract method 0xc54837a4.
//
// Solidity: function extinguish() returns()
func (_PresaleSecond *PresaleSecondSession) Extinguish() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Extinguish(&_PresaleSecond.TransactOpts)
}

// Extinguish is a paid mutator transaction binding the contract method 0xc54837a4.
//
// Solidity: function extinguish() returns()
func (_PresaleSecond *PresaleSecondTransactorSession) Extinguish() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Extinguish(&_PresaleSecond.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_PresaleSecond *PresaleSecondTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_PresaleSecond *PresaleSecondSession) Finalize() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Finalize(&_PresaleSecond.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_PresaleSecond *PresaleSecondTransactorSession) Finalize() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Finalize(&_PresaleSecond.TransactOpts)
}

// Ignite is a paid mutator transaction binding the contract method 0xf768923a.
//
// Solidity: function ignite() returns()
func (_PresaleSecond *PresaleSecondTransactor) Ignite(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "ignite")
}

// Ignite is a paid mutator transaction binding the contract method 0xf768923a.
//
// Solidity: function ignite() returns()
func (_PresaleSecond *PresaleSecondSession) Ignite() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Ignite(&_PresaleSecond.TransactOpts)
}

// Ignite is a paid mutator transaction binding the contract method 0xf768923a.
//
// Solidity: function ignite() returns()
func (_PresaleSecond *PresaleSecondTransactorSession) Ignite() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Ignite(&_PresaleSecond.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PresaleSecond *PresaleSecondTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PresaleSecond *PresaleSecondSession) Pause() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Pause(&_PresaleSecond.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PresaleSecond *PresaleSecondTransactorSession) Pause() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Pause(&_PresaleSecond.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(_addr address) returns(bool)
func (_PresaleSecond *PresaleSecondTransactor) Refund(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "refund", _addr)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(_addr address) returns(bool)
func (_PresaleSecond *PresaleSecondSession) Refund(_addr common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.Refund(&_PresaleSecond.TransactOpts, _addr)
}

// Refund is a paid mutator transaction binding the contract method 0xfa89401a.
//
// Solidity: function refund(_addr address) returns(bool)
func (_PresaleSecond *PresaleSecondTransactorSession) Refund(_addr common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.Refund(&_PresaleSecond.TransactOpts, _addr)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_addr address) returns(bool)
func (_PresaleSecond *PresaleSecondTransactor) Release(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "release", _addr)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_addr address) returns(bool)
func (_PresaleSecond *PresaleSecondSession) Release(_addr common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.Release(&_PresaleSecond.TransactOpts, _addr)
}

// Release is a paid mutator transaction binding the contract method 0x19165587.
//
// Solidity: function release(_addr address) returns(bool)
func (_PresaleSecond *PresaleSecondTransactorSession) Release(_addr common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.Release(&_PresaleSecond.TransactOpts, _addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PresaleSecond *PresaleSecondTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PresaleSecond *PresaleSecondSession) RenounceOwnership() (*types.Transaction, error) {
	return _PresaleSecond.Contract.RenounceOwnership(&_PresaleSecond.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PresaleSecond *PresaleSecondTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PresaleSecond.Contract.RenounceOwnership(&_PresaleSecond.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_PresaleSecond *PresaleSecondTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_PresaleSecond *PresaleSecondSession) Resume() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Resume(&_PresaleSecond.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_PresaleSecond *PresaleSecondTransactorSession) Resume() (*types.Transaction, error) {
	return _PresaleSecond.Contract.Resume(&_PresaleSecond.TransactOpts)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(_distributor address) returns()
func (_PresaleSecond *PresaleSecondTransactor) SetDistributor(opts *bind.TransactOpts, _distributor common.Address) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "setDistributor", _distributor)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(_distributor address) returns()
func (_PresaleSecond *PresaleSecondSession) SetDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.SetDistributor(&_PresaleSecond.TransactOpts, _distributor)
}

// SetDistributor is a paid mutator transaction binding the contract method 0x75619ab5.
//
// Solidity: function setDistributor(_distributor address) returns()
func (_PresaleSecond *PresaleSecondTransactorSession) SetDistributor(_distributor common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.SetDistributor(&_PresaleSecond.TransactOpts, _distributor)
}

// SetWallet is a paid mutator transaction binding the contract method 0xdeaa59df.
//
// Solidity: function setWallet(_wallet address) returns()
func (_PresaleSecond *PresaleSecondTransactor) SetWallet(opts *bind.TransactOpts, _wallet common.Address) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "setWallet", _wallet)
}

// SetWallet is a paid mutator transaction binding the contract method 0xdeaa59df.
//
// Solidity: function setWallet(_wallet address) returns()
func (_PresaleSecond *PresaleSecondSession) SetWallet(_wallet common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.SetWallet(&_PresaleSecond.TransactOpts, _wallet)
}

// SetWallet is a paid mutator transaction binding the contract method 0xdeaa59df.
//
// Solidity: function setWallet(_wallet address) returns()
func (_PresaleSecond *PresaleSecondTransactorSession) SetWallet(_wallet common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.SetWallet(&_PresaleSecond.TransactOpts, _wallet)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(_whitelist address) returns()
func (_PresaleSecond *PresaleSecondTransactor) SetWhitelist(opts *bind.TransactOpts, _whitelist common.Address) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "setWhitelist", _whitelist)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(_whitelist address) returns()
func (_PresaleSecond *PresaleSecondSession) SetWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.SetWhitelist(&_PresaleSecond.TransactOpts, _whitelist)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x854cff2f.
//
// Solidity: function setWhitelist(_whitelist address) returns()
func (_PresaleSecond *PresaleSecondTransactorSession) SetWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.SetWhitelist(&_PresaleSecond.TransactOpts, _whitelist)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PresaleSecond *PresaleSecondTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PresaleSecond *PresaleSecondSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.TransferOwnership(&_PresaleSecond.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PresaleSecond *PresaleSecondTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PresaleSecond.Contract.TransferOwnership(&_PresaleSecond.TransactOpts, newOwner)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_PresaleSecond *PresaleSecondTransactor) WithdrawEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "withdrawEther")
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_PresaleSecond *PresaleSecondSession) WithdrawEther() (*types.Transaction, error) {
	return _PresaleSecond.Contract.WithdrawEther(&_PresaleSecond.TransactOpts)
}

// WithdrawEther is a paid mutator transaction binding the contract method 0x7362377b.
//
// Solidity: function withdrawEther() returns()
func (_PresaleSecond *PresaleSecondTransactorSession) WithdrawEther() (*types.Transaction, error) {
	return _PresaleSecond.Contract.WithdrawEther(&_PresaleSecond.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_PresaleSecond *PresaleSecondTransactor) WithdrawToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleSecond.contract.Transact(opts, "withdrawToken")
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_PresaleSecond *PresaleSecondSession) WithdrawToken() (*types.Transaction, error) {
	return _PresaleSecond.Contract.WithdrawToken(&_PresaleSecond.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_PresaleSecond *PresaleSecondTransactorSession) WithdrawToken() (*types.Transaction, error) {
	return _PresaleSecond.Contract.WithdrawToken(&_PresaleSecond.TransactOpts)
}

// PresaleSecondChangeIterator is returned from FilterChange and is used to iterate over the raw logs and unpacked data for Change events raised by the PresaleSecond contract.
type PresaleSecondChangeIterator struct {
	Event *PresaleSecondChange // Event containing the contract specifics and raw log

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
func (it *PresaleSecondChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondChange)
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
		it.Event = new(PresaleSecondChange)
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
func (it *PresaleSecondChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondChange represents a Change event raised by the PresaleSecond contract.
type PresaleSecondChange struct {
	Addr common.Address
	Name string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterChange is a free log retrieval operation binding the contract event 0xa785fc346da73c9ad6c933dde68fe85362a97b7b07706c3e23ff3400cc5d93b5.
//
// Solidity: event Change(_addr address, _name string)
func (_PresaleSecond *PresaleSecondFilterer) FilterChange(opts *bind.FilterOpts) (*PresaleSecondChangeIterator, error) {

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Change")
	if err != nil {
		return nil, err
	}
	return &PresaleSecondChangeIterator{contract: _PresaleSecond.contract, event: "Change", logs: logs, sub: sub}, nil
}

// WatchChange is a free log subscription operation binding the contract event 0xa785fc346da73c9ad6c933dde68fe85362a97b7b07706c3e23ff3400cc5d93b5.
//
// Solidity: event Change(_addr address, _name string)
func (_PresaleSecond *PresaleSecondFilterer) WatchChange(opts *bind.WatchOpts, sink chan<- *PresaleSecondChange) (event.Subscription, error) {

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "Change")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondChange)
				if err := _PresaleSecond.contract.UnpackLog(event, "Change", log); err != nil {
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

// PresaleSecondExtinguishIterator is returned from FilterExtinguish and is used to iterate over the raw logs and unpacked data for Extinguish events raised by the PresaleSecond contract.
type PresaleSecondExtinguishIterator struct {
	Event *PresaleSecondExtinguish // Event containing the contract specifics and raw log

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
func (it *PresaleSecondExtinguishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondExtinguish)
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
		it.Event = new(PresaleSecondExtinguish)
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
func (it *PresaleSecondExtinguishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondExtinguishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondExtinguish represents a Extinguish event raised by the PresaleSecond contract.
type PresaleSecondExtinguish struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterExtinguish is a free log retrieval operation binding the contract event 0xd53036fa90376b59ca8afb9d95483e6b47ffa785d597814fcfd7405a91bba67a.
//
// Solidity: event Extinguish()
func (_PresaleSecond *PresaleSecondFilterer) FilterExtinguish(opts *bind.FilterOpts) (*PresaleSecondExtinguishIterator, error) {

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Extinguish")
	if err != nil {
		return nil, err
	}
	return &PresaleSecondExtinguishIterator{contract: _PresaleSecond.contract, event: "Extinguish", logs: logs, sub: sub}, nil
}

// WatchExtinguish is a free log subscription operation binding the contract event 0xd53036fa90376b59ca8afb9d95483e6b47ffa785d597814fcfd7405a91bba67a.
//
// Solidity: event Extinguish()
func (_PresaleSecond *PresaleSecondFilterer) WatchExtinguish(opts *bind.WatchOpts, sink chan<- *PresaleSecondExtinguish) (event.Subscription, error) {

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "Extinguish")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondExtinguish)
				if err := _PresaleSecond.contract.UnpackLog(event, "Extinguish", log); err != nil {
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

// PresaleSecondIgniteIterator is returned from FilterIgnite and is used to iterate over the raw logs and unpacked data for Ignite events raised by the PresaleSecond contract.
type PresaleSecondIgniteIterator struct {
	Event *PresaleSecondIgnite // Event containing the contract specifics and raw log

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
func (it *PresaleSecondIgniteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondIgnite)
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
		it.Event = new(PresaleSecondIgnite)
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
func (it *PresaleSecondIgniteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondIgniteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondIgnite represents a Ignite event raised by the PresaleSecond contract.
type PresaleSecondIgnite struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIgnite is a free log retrieval operation binding the contract event 0xb66ce7cc84acb9e2cdfa3c16073eec63d39b0540887b91247afebaee4645611a.
//
// Solidity: event Ignite()
func (_PresaleSecond *PresaleSecondFilterer) FilterIgnite(opts *bind.FilterOpts) (*PresaleSecondIgniteIterator, error) {

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Ignite")
	if err != nil {
		return nil, err
	}
	return &PresaleSecondIgniteIterator{contract: _PresaleSecond.contract, event: "Ignite", logs: logs, sub: sub}, nil
}

// WatchIgnite is a free log subscription operation binding the contract event 0xb66ce7cc84acb9e2cdfa3c16073eec63d39b0540887b91247afebaee4645611a.
//
// Solidity: event Ignite()
func (_PresaleSecond *PresaleSecondFilterer) WatchIgnite(opts *bind.WatchOpts, sink chan<- *PresaleSecondIgnite) (event.Subscription, error) {

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "Ignite")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondIgnite)
				if err := _PresaleSecond.contract.UnpackLog(event, "Ignite", log); err != nil {
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

// PresaleSecondOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PresaleSecond contract.
type PresaleSecondOwnershipRenouncedIterator struct {
	Event *PresaleSecondOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PresaleSecondOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondOwnershipRenounced)
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
		it.Event = new(PresaleSecondOwnershipRenounced)
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
func (it *PresaleSecondOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondOwnershipRenounced represents a OwnershipRenounced event raised by the PresaleSecond contract.
type PresaleSecondOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_PresaleSecond *PresaleSecondFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PresaleSecondOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PresaleSecondOwnershipRenouncedIterator{contract: _PresaleSecond.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(previousOwner indexed address)
func (_PresaleSecond *PresaleSecondFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PresaleSecondOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondOwnershipRenounced)
				if err := _PresaleSecond.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PresaleSecondOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PresaleSecond contract.
type PresaleSecondOwnershipTransferredIterator struct {
	Event *PresaleSecondOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PresaleSecondOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondOwnershipTransferred)
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
		it.Event = new(PresaleSecondOwnershipTransferred)
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
func (it *PresaleSecondOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondOwnershipTransferred represents a OwnershipTransferred event raised by the PresaleSecond contract.
type PresaleSecondOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PresaleSecond *PresaleSecondFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PresaleSecondOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PresaleSecondOwnershipTransferredIterator{contract: _PresaleSecond.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PresaleSecond *PresaleSecondFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PresaleSecondOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondOwnershipTransferred)
				if err := _PresaleSecond.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PresaleSecondPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PresaleSecond contract.
type PresaleSecondPauseIterator struct {
	Event *PresaleSecondPause // Event containing the contract specifics and raw log

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
func (it *PresaleSecondPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondPause)
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
		it.Event = new(PresaleSecondPause)
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
func (it *PresaleSecondPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondPause represents a Pause event raised by the PresaleSecond contract.
type PresaleSecondPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PresaleSecond *PresaleSecondFilterer) FilterPause(opts *bind.FilterOpts) (*PresaleSecondPauseIterator, error) {

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PresaleSecondPauseIterator{contract: _PresaleSecond.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_PresaleSecond *PresaleSecondFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PresaleSecondPause) (event.Subscription, error) {

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondPause)
				if err := _PresaleSecond.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PresaleSecondPurchaseIterator is returned from FilterPurchase and is used to iterate over the raw logs and unpacked data for Purchase events raised by the PresaleSecond contract.
type PresaleSecondPurchaseIterator struct {
	Event *PresaleSecondPurchase // Event containing the contract specifics and raw log

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
func (it *PresaleSecondPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondPurchase)
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
		it.Event = new(PresaleSecondPurchase)
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
func (it *PresaleSecondPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondPurchase represents a Purchase event raised by the PresaleSecond contract.
type PresaleSecondPurchase struct {
	Buyer     common.Address
	Purchased *big.Int
	Refund    *big.Int
	Tokens    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPurchase is a free log retrieval operation binding the contract event 0x5bc97d73357ac0d035d4b9268a69240988a5776b8a4fcced3dbc223960123f40.
//
// Solidity: event Purchase(_buyer indexed address, _purchased uint256, _refund uint256, _tokens uint256)
func (_PresaleSecond *PresaleSecondFilterer) FilterPurchase(opts *bind.FilterOpts, _buyer []common.Address) (*PresaleSecondPurchaseIterator, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Purchase", _buyerRule)
	if err != nil {
		return nil, err
	}
	return &PresaleSecondPurchaseIterator{contract: _PresaleSecond.contract, event: "Purchase", logs: logs, sub: sub}, nil
}

// WatchPurchase is a free log subscription operation binding the contract event 0x5bc97d73357ac0d035d4b9268a69240988a5776b8a4fcced3dbc223960123f40.
//
// Solidity: event Purchase(_buyer indexed address, _purchased uint256, _refund uint256, _tokens uint256)
func (_PresaleSecond *PresaleSecondFilterer) WatchPurchase(opts *bind.WatchOpts, sink chan<- *PresaleSecondPurchase, _buyer []common.Address) (event.Subscription, error) {

	var _buyerRule []interface{}
	for _, _buyerItem := range _buyer {
		_buyerRule = append(_buyerRule, _buyerItem)
	}

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "Purchase", _buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondPurchase)
				if err := _PresaleSecond.contract.UnpackLog(event, "Purchase", log); err != nil {
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

// PresaleSecondRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the PresaleSecond contract.
type PresaleSecondRefundIterator struct {
	Event *PresaleSecondRefund // Event containing the contract specifics and raw log

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
func (it *PresaleSecondRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondRefund)
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
		it.Event = new(PresaleSecondRefund)
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
func (it *PresaleSecondRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondRefund represents a Refund event raised by the PresaleSecond contract.
type PresaleSecondRefund struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0xbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d.
//
// Solidity: event Refund(_to indexed address, _amount uint256)
func (_PresaleSecond *PresaleSecondFilterer) FilterRefund(opts *bind.FilterOpts, _to []common.Address) (*PresaleSecondRefundIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Refund", _toRule)
	if err != nil {
		return nil, err
	}
	return &PresaleSecondRefundIterator{contract: _PresaleSecond.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0xbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d.
//
// Solidity: event Refund(_to indexed address, _amount uint256)
func (_PresaleSecond *PresaleSecondFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *PresaleSecondRefund, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "Refund", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondRefund)
				if err := _PresaleSecond.contract.UnpackLog(event, "Refund", log); err != nil {
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

// PresaleSecondReleaseIterator is returned from FilterRelease and is used to iterate over the raw logs and unpacked data for Release events raised by the PresaleSecond contract.
type PresaleSecondReleaseIterator struct {
	Event *PresaleSecondRelease // Event containing the contract specifics and raw log

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
func (it *PresaleSecondReleaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondRelease)
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
		it.Event = new(PresaleSecondRelease)
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
func (it *PresaleSecondReleaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondReleaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondRelease represents a Release event raised by the PresaleSecond contract.
type PresaleSecondRelease struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRelease is a free log retrieval operation binding the contract event 0xf6334794522b9db534a812aaae1af828a2e96aac68473b58e36d7d0bfd67477b.
//
// Solidity: event Release(_to indexed address, _amount uint256)
func (_PresaleSecond *PresaleSecondFilterer) FilterRelease(opts *bind.FilterOpts, _to []common.Address) (*PresaleSecondReleaseIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Release", _toRule)
	if err != nil {
		return nil, err
	}
	return &PresaleSecondReleaseIterator{contract: _PresaleSecond.contract, event: "Release", logs: logs, sub: sub}, nil
}

// WatchRelease is a free log subscription operation binding the contract event 0xf6334794522b9db534a812aaae1af828a2e96aac68473b58e36d7d0bfd67477b.
//
// Solidity: event Release(_to indexed address, _amount uint256)
func (_PresaleSecond *PresaleSecondFilterer) WatchRelease(opts *bind.WatchOpts, sink chan<- *PresaleSecondRelease, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "Release", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondRelease)
				if err := _PresaleSecond.contract.UnpackLog(event, "Release", log); err != nil {
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

// PresaleSecondResumeIterator is returned from FilterResume and is used to iterate over the raw logs and unpacked data for Resume events raised by the PresaleSecond contract.
type PresaleSecondResumeIterator struct {
	Event *PresaleSecondResume // Event containing the contract specifics and raw log

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
func (it *PresaleSecondResumeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondResume)
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
		it.Event = new(PresaleSecondResume)
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
func (it *PresaleSecondResumeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondResumeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondResume represents a Resume event raised by the PresaleSecond contract.
type PresaleSecondResume struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResume is a free log retrieval operation binding the contract event 0x490d6d11e278f168be9be39e46297f72ea877136d5bccad9cf4993e63a29568f.
//
// Solidity: event Resume()
func (_PresaleSecond *PresaleSecondFilterer) FilterResume(opts *bind.FilterOpts) (*PresaleSecondResumeIterator, error) {

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Resume")
	if err != nil {
		return nil, err
	}
	return &PresaleSecondResumeIterator{contract: _PresaleSecond.contract, event: "Resume", logs: logs, sub: sub}, nil
}

// WatchResume is a free log subscription operation binding the contract event 0x490d6d11e278f168be9be39e46297f72ea877136d5bccad9cf4993e63a29568f.
//
// Solidity: event Resume()
func (_PresaleSecond *PresaleSecondFilterer) WatchResume(opts *bind.WatchOpts, sink chan<- *PresaleSecondResume) (event.Subscription, error) {

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "Resume")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondResume)
				if err := _PresaleSecond.contract.UnpackLog(event, "Resume", log); err != nil {
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

// PresaleSecondWithdrawEtherIterator is returned from FilterWithdrawEther and is used to iterate over the raw logs and unpacked data for WithdrawEther events raised by the PresaleSecond contract.
type PresaleSecondWithdrawEtherIterator struct {
	Event *PresaleSecondWithdrawEther // Event containing the contract specifics and raw log

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
func (it *PresaleSecondWithdrawEtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondWithdrawEther)
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
		it.Event = new(PresaleSecondWithdrawEther)
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
func (it *PresaleSecondWithdrawEtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondWithdrawEtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondWithdrawEther represents a WithdrawEther event raised by the PresaleSecond contract.
type PresaleSecondWithdrawEther struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawEther is a free log retrieval operation binding the contract event 0xdb35132c111efe920cede025e819975671cfd1b8fcc1174762c8670c4e94c211.
//
// Solidity: event WithdrawEther(_from indexed address, _amount uint256)
func (_PresaleSecond *PresaleSecondFilterer) FilterWithdrawEther(opts *bind.FilterOpts, _from []common.Address) (*PresaleSecondWithdrawEtherIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "WithdrawEther", _fromRule)
	if err != nil {
		return nil, err
	}
	return &PresaleSecondWithdrawEtherIterator{contract: _PresaleSecond.contract, event: "WithdrawEther", logs: logs, sub: sub}, nil
}

// WatchWithdrawEther is a free log subscription operation binding the contract event 0xdb35132c111efe920cede025e819975671cfd1b8fcc1174762c8670c4e94c211.
//
// Solidity: event WithdrawEther(_from indexed address, _amount uint256)
func (_PresaleSecond *PresaleSecondFilterer) WatchWithdrawEther(opts *bind.WatchOpts, sink chan<- *PresaleSecondWithdrawEther, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "WithdrawEther", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondWithdrawEther)
				if err := _PresaleSecond.contract.UnpackLog(event, "WithdrawEther", log); err != nil {
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

// PresaleSecondWithdrawTokenIterator is returned from FilterWithdrawToken and is used to iterate over the raw logs and unpacked data for WithdrawToken events raised by the PresaleSecond contract.
type PresaleSecondWithdrawTokenIterator struct {
	Event *PresaleSecondWithdrawToken // Event containing the contract specifics and raw log

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
func (it *PresaleSecondWithdrawTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleSecondWithdrawToken)
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
		it.Event = new(PresaleSecondWithdrawToken)
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
func (it *PresaleSecondWithdrawTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleSecondWithdrawTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleSecondWithdrawToken represents a WithdrawToken event raised by the PresaleSecond contract.
type PresaleSecondWithdrawToken struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawToken is a free log retrieval operation binding the contract event 0x992ee874049a42cae0757a765cd7f641b6028cc35c3478bde8330bf417c3a7a9.
//
// Solidity: event WithdrawToken(_from indexed address, _amount uint256)
func (_PresaleSecond *PresaleSecondFilterer) FilterWithdrawToken(opts *bind.FilterOpts, _from []common.Address) (*PresaleSecondWithdrawTokenIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "WithdrawToken", _fromRule)
	if err != nil {
		return nil, err
	}
	return &PresaleSecondWithdrawTokenIterator{contract: _PresaleSecond.contract, event: "WithdrawToken", logs: logs, sub: sub}, nil
}

// WatchWithdrawToken is a free log subscription operation binding the contract event 0x992ee874049a42cae0757a765cd7f641b6028cc35c3478bde8330bf417c3a7a9.
//
// Solidity: event WithdrawToken(_from indexed address, _amount uint256)
func (_PresaleSecond *PresaleSecondFilterer) WatchWithdrawToken(opts *bind.WatchOpts, sink chan<- *PresaleSecondWithdrawToken, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _PresaleSecond.contract.WatchLogs(opts, "WithdrawToken", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleSecondWithdrawToken)
				if err := _PresaleSecond.contract.UnpackLog(event, "WithdrawToken", log); err != nil {
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
