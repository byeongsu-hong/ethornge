// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adapter

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

// TokenDistributorBin is the compiled bytecode used for deploying new contracts.
const TokenDistributorBin = `0x608060405234801561001057600080fd5b5060405160a08061088c8339810160409081528151602083015191830151606084015160809094015160008054600160a060020a031916331790559193909184841161005b57600080fd5b6000841161006857600080fd5b600160a060020a038316151561007d57600080fd5b600160a060020a038216151561009257600080fd5b600160a060020a03811615156100a757600080fd5b60019490945560029290925560038054600160a060020a03928316600160a060020a03199182161790915560048054938316938216939093179092556005805491909316911617905561078d806100ff6000396000f3006080604052600436106100a35763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630c55699c81146100a857806319165587146100cf57806346620e39146100f25780636bf6eaff14610123578063715018a6146101385780638da5cb5b1461014d578063a56dfe4a14610162578063c241267614610177578063e52f64ce1461018c578063f2fde38b146101ac575b600080fd5b3480156100b457600080fd5b506100bd6101cd565b60408051918252519081900360200190f35b3480156100db57600080fd5b506100f0600160a060020a03600435166101d3565b005b3480156100fe57600080fd5b506101076104b8565b60408051600160a060020a039092168252519081900360200190f35b34801561012f57600080fd5b506101076104c7565b34801561014457600080fd5b506100f06104d6565b34801561015957600080fd5b50610107610542565b34801561016e57600080fd5b506100bd610551565b34801561018357600080fd5b50610107610557565b34801561019857600080fd5b506100f06004803560248101910135610566565b3480156101b857600080fd5b506100f0600160a060020a03600435166105c8565b60015481565b60008054819081908190600160a060020a031633146101f157600080fd5b600554604080517ff8b2cb4f000000000000000000000000000000000000000000000000000000008152600160a060020a0388811660048301529151919092169163f8b2cb4f9160248083019260209291908290030181600087803b15801561025957600080fd5b505af115801561026d573d6000803e3d6000fd5b505050506040513d602081101561028357600080fd5b505160048054604080517f70a08231000000000000000000000000000000000000000000000000000000008152600160a060020a038a81169482019490945290519397509116916370a08231916024808201926020929091908290030181600087803b1580156102f257600080fd5b505af1158015610306573d6000803e3d6000fd5b505050506040513d602081101561031c57600080fd5b50519250600160a060020a038516151561033557600080fd5b831561034057600080fd5b6000831161034d57600080fd5b6103746001546103686002548661065c90919063ffffffff16565b9063ffffffff61067316565b91506103a461039060015460025461069c90919063ffffffff16565b60025461036890869063ffffffff61065c16565b6005546003549192506103ca91600160a060020a0390811691168363ffffffff6106ae16565b600554604080517f282d3fdf000000000000000000000000000000000000000000000000000000008152600160a060020a038881166004830152602482018590529151919092169163282d3fdf91604480830192600092919082900301818387803b15801561043857600080fd5b505af115801561044c573d6000803e3d6000fd5b505060035461046e9250600160a060020a03169050868463ffffffff6106ae16565b60408051838152602081018390528151600160a060020a038816927f2c57dec1db0095a6b800c2698d5bbceef2c180c6ac43429769a719658983f677928290030190a25050505050565b600554600160a060020a031681565b600454600160a060020a031681565b600054600160a060020a031633146104ed57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031681565b60025481565b600354600160a060020a031681565b60008054600160a060020a0316331461057e57600080fd5b601e821061058b57600080fd5b5060005b818110156105c3576105bb8383838181106105a657fe5b90506020020135600160a060020a03166101d3565b60010161058f565b505050565b600054600160a060020a031633146105df57600080fd5b600160a060020a03811615156105f457600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b6000818381151561066957fe5b0490505b92915050565b60008215156106845750600061066d565b5081810281838281151561069457fe5b041461066d57fe5b6000828211156106a857fe5b50900390565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561072a57600080fd5b505af115801561073e573d6000803e3d6000fd5b505050506040513d602081101561075457600080fd5b505115156105c357600080fd00a165627a7a7230582089ac940cf6e0eb55ee81994b3436ab920664e535d56feb8968f500be3fe1a6680029`

// DeployTokenDistributor deploys a new Ethereum contract, binding an instance of TokenDistributor to it.
func DeployTokenDistributor(auth *bind.TransactOpts, backend bind.ContractBackend, _ratioX *big.Int, _ratioY *big.Int, _token common.Address, _genesis common.Address, _tokenLock common.Address) (common.Address, *types.Transaction, *TokenDistributor, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenDistributorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenDistributorBin), backend, _ratioX, _ratioY, _token, _genesis, _tokenLock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenDistributor{TokenDistributorCaller: TokenDistributorCaller{contract: contract}, TokenDistributorTransactor: TokenDistributorTransactor{contract: contract}, TokenDistributorFilterer: TokenDistributorFilterer{contract: contract}}, nil
}

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
// Solidity: e OwnershipRenounced(previousOwner indexed address)
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
// Solidity: e OwnershipRenounced(previousOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e Release(_to indexed address, _safeAmount uint256, _lockAmount uint256)
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
// Solidity: e Release(_to indexed address, _safeAmount uint256, _lockAmount uint256)
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
