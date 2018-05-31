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

// PresaleSecondABI is the input ABI used to generate the binding from.
const PresaleSecondABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exceed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"release\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxcap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ignited\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawEther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_distributor\",\"type\":\"address\"}],\"name\":\"setDistributor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_whitelist\",\"type\":\"address\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"finalized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"distributor\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"Token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"List\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"extinguish\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_wallet\",\"type\":\"address\"}],\"name\":\"setWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"collect\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ignite\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"refund\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_maxcap\",\"type\":\"uint256\"},{\"name\":\"_exceed\",\"type\":\"uint256\"},{\"name\":\"_minimum\",\"type\":\"uint256\"},{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_distributor\",\"type\":\"address\"},{\"name\":\"_whitelist\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"Change\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resume\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Ignite\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Extinguish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_purchased\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_refund\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tokens\",\"type\":\"uint256\"}],\"name\":\"Purchase\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Release\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawEther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PresaleSecondBin is the compiled bytecode used for deploying new contracts.
const PresaleSecondBin = `0x60806040526005805461ffff191690556000600655600c805460ff1916905534801561002a57600080fd5b50604051610100806113cd83398101604090815281516020830151918301516060840151608085015160a086015160c087015160e09097015160008054600160a060020a031916331790559496939492939192909190600160a060020a038416151561009557600080fd5b600160a060020a03821615156100aa57600080fd5b600160a060020a03831615156100bf57600080fd5b600160a060020a03811615156100d457600080fd5b60019790975560029590955560039390935560049190915560078054600160a060020a03928316600160a060020a0319918216179091556008805493831693821693909317909255600a80549482169483169490941790935560098054929093169116179055611284806101496000396000f3006080604052600436106101745763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663046f7da2811461017e5780630b1815671461019357806319165587146101ba5780632c4e722e146101ef5780634042b66f146102045780634bb278f314610219578063521eb2731461022e57806352d6804d1461025f57806355456f58146102745780635c975abb1461028957806366092ea81461029e578063715018a6146102b35780637362377b146102c857806375619ab5146102dd5780638456cb59146102fe578063854cff2f146103135780638da5cb5b1461033457806397a993aa14610349578063b3f05b971461036a578063bfe109281461037f578063c241267614610394578063c258ff74146103a9578063c54837a4146103be578063ca628c78146103d3578063deaa59df146103e8578063e522538114610174578063f2fde38b14610409578063f768923a1461042a578063fa89401a1461043f575b61017c610460565b005b34801561018a57600080fd5b5061017c6106c3565b34801561019f57600080fd5b506101a861070f565b60408051918252519081900360200190f35b3480156101c657600080fd5b506101db600160a060020a0360043516610715565b604080519115158252519081900360200190f35b3480156101fb57600080fd5b506101a861083b565b34801561021057600080fd5b506101a8610841565b34801561022557600080fd5b5061017c610847565b34801561023a57600080fd5b506102436108a3565b60408051600160a060020a039092168252519081900360200190f35b34801561026b57600080fd5b506101a86108b2565b34801561028057600080fd5b506101a86108b8565b34801561029557600080fd5b506101db6108be565b3480156102aa57600080fd5b506101db6108c7565b3480156102bf57600080fd5b5061017c6108d5565b3480156102d457600080fd5b5061017c610941565b3480156102e957600080fd5b5061017c600160a060020a03600435166109ed565b34801561030a57600080fd5b5061017c610aad565b34801561031f57600080fd5b5061017c600160a060020a0360043516610afc565b34801561034057600080fd5b50610243610bbd565b34801561035557600080fd5b506101a8600160a060020a0360043516610bcc565b34801561037657600080fd5b506101db610bde565b34801561038b57600080fd5b50610243610be7565b3480156103a057600080fd5b50610243610bf6565b3480156103b557600080fd5b50610243610c05565b3480156103ca57600080fd5b5061017c610c14565b3480156103df57600080fd5b5061017c610c61565b3480156103f457600080fd5b5061017c600160a060020a0360043516610e0b565b34801561041557600080fd5b5061017c600160a060020a0360043516610ecb565b34801561043657600080fd5b5061017c610f5f565b34801561044b57600080fd5b506101db600160a060020a0360043516610fb0565b600554339034906000908190610100900460ff168015610483575060055460ff16155b151561048e57600080fd5b600954604080517f9b19251a000000000000000000000000000000000000000000000000000000008152600160a060020a03878116600483015291519190921691639b19251a9160248083019260209291908290030181600087803b1580156104f657600080fd5b505af115801561050a573d6000803e3d6000fd5b505050506040513d602081101561052057600080fd5b5051151561052d57600080fd5b600160a060020a038416151561054257600080fd5b600354600160a060020a0385166000908152600b602052604090205461056e908563ffffffff6110c316565b101561057957600080fd5b600254600160a060020a0385166000908152600b60205260409020541061059f57600080fd5b600154600654106105af57600080fd5b6105b984846110d6565b60065491935091506105d1908363ffffffff6110c316565b6006819055600154116105ea576005805461ff00191690555b600160a060020a0384166000908152600b6020526040902054610613908363ffffffff6110c316565b600160a060020a0385166000818152600b60205260409020919091556004547f5bc97d73357ac0d035d4b9268a69240988a5776b8a4fcced3dbc223960123f40908490849061066990839063ffffffff61116516565b60408051938452602084019290925282820152519081900360600190a2604051600160a060020a0385169082156108fc029083906000818181858888f193505050501580156106bc573d6000803e3d6000fd5b5050505050565b600054600160a060020a031633146106da57600080fd5b6005805460ff191690556040517f490d6d11e278f168be9be39e46297f72ea877136d5bccad9cf4993e63a29568f90600090a1565b60025481565b6005546000908190610100900460ff161580156107355750600c5460ff16155b151561074057600080fd5b600854600160a060020a0316331461075757600080fd5b600160a060020a038316151561076c57600080fd5b600160a060020a0383166000908152600b602052604090205415156107945760009150610835565b600454600160a060020a0384166000908152600b60205260409020546107bf9163ffffffff61116516565b600160a060020a038085166000908152600b6020526040812055600a549192506107f19116848363ffffffff61118e16565b604080518281529051600160a060020a038516917ff6334794522b9db534a812aaae1af828a2e96aac68473b58e36d7d0bfd67477b919081900360200190a2600191505b50919050565b60045481565b60065481565b600054600160a060020a0316331461085e57600080fd5b600554610100900460ff161580156108795750600c5460ff16155b151561088457600080fd5b61088c610941565b610894610c61565b600c805460ff19166001179055565b600754600160a060020a031681565b60035481565b60015481565b60055460ff1681565b600554610100900460ff1681565b600054600160a060020a031633146108ec57600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a0316331461095857600080fd5b600554610100900460ff161561096d57600080fd5b600754604051600160a060020a0390911690303180156108fc02916000818181858888f193505050501580156109a7573d6000803e3d6000fd5b5060075460408051303181529051600160a060020a03909216917fdb35132c111efe920cede025e819975671cfd1b8fcc1174762c8670c4e94c2119181900360200190a2565b600054600160a060020a03163314610a0457600080fd5b600160a060020a0381161515610a1957600080fd5b60088054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911681179091556040805191825260208201819052600b828201527f6469737472696275746f720000000000000000000000000000000000000000006060830152517fa785fc346da73c9ad6c933dde68fe85362a97b7b07706c3e23ff3400cc5d93b59181900360800190a150565b600054600160a060020a03163314610ac457600080fd5b6005805460ff191660011790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b600054600160a060020a03163314610b1357600080fd5b600160a060020a0381161515610b2857600080fd5b60098054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811782556040805191825260208201819052818101929092527f77686974656c6973740000000000000000000000000000000000000000000000606082015290517fa785fc346da73c9ad6c933dde68fe85362a97b7b07706c3e23ff3400cc5d93b59181900360800190a150565b600054600160a060020a031681565b600b6020526000908152604090205481565b600c5460ff1681565b600854600160a060020a031681565b600a54600160a060020a031681565b600954600160a060020a031681565b600054600160a060020a03163314610c2b57600080fd5b6005805461ff00191690556040517fd53036fa90376b59ca8afb9d95483e6b47ffa785d597814fcfd7405a91bba67a90600090a1565b600054600160a060020a03163314610c7857600080fd5b600554610100900460ff1615610c8d57600080fd5b600754600a54604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051610d4093600160a060020a039081169316916370a082319160248083019260209291908290030181600087803b158015610cfb57600080fd5b505af1158015610d0f573d6000803e3d6000fd5b505050506040513d6020811015610d2557600080fd5b5051600a54600160a060020a0316919063ffffffff61118e16565b600754600a54604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a03938416937f992ee874049a42cae0757a765cd7f641b6028cc35c3478bde8330bf417c3a7a99316916370a082319160248083019260209291908290030181600087803b158015610dcc57600080fd5b505af1158015610de0573d6000803e3d6000fd5b505050506040513d6020811015610df657600080fd5b505160408051918252519081900360200190a2565b600054600160a060020a03163314610e2257600080fd5b600160a060020a0381161515610e3757600080fd5b60078054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff19909116811790915560408051918252602082018190526006828201527f77616c6c657400000000000000000000000000000000000000000000000000006060830152517fa785fc346da73c9ad6c933dde68fe85362a97b7b07706c3e23ff3400cc5d93b59181900360800190a150565b600054600160a060020a03163314610ee257600080fd5b600160a060020a0381161515610ef757600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600054600160a060020a03163314610f7657600080fd5b6005805461ff0019166101001790556040517fb66ce7cc84acb9e2cdfa3c16073eec63d39b0540887b91247afebaee4645611a90600090a1565b6005546000908190610100900460ff16158015610fd05750600c5460ff16155b1515610fdb57600080fd5b600854600160a060020a03163314610ff257600080fd5b600160a060020a038316151561100757600080fd5b600160a060020a0383166000908152600b6020526040902054151561102f5760009150610835565b50600160a060020a0382166000818152600b6020526040808220805490839055905190929183156108fc02918491818181858888f1935050505015801561107a573d6000803e3d6000fd5b50604080518281529051600160a060020a038516917fbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d919081900360200190a250600192915050565b818101828110156110d057fe5b92915050565b60008060008060006110f560065460015461124690919063ffffffff16565b600160a060020a0388166000908152600b6020526040902054600254919450611124919063ffffffff61124616565b91508183116111335782611135565b815b905080861161114657856000611157565b80611157878263ffffffff61124616565b945094505050509250929050565b6000821515611176575060006110d0565b5081810281838281151561118657fe5b04146110d057fe5b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561120a57600080fd5b505af115801561121e573d6000803e3d6000fd5b505050506040513d602081101561123457600080fd5b5051151561124157600080fd5b505050565b60008282111561125257fe5b509003905600a165627a7a72305820d641178c70b4cdb127ac0fde22fae228fdad0bcbe4fd027cedfdfcc3fa21df860029`

// DeployPresaleSecond deploys a new Ethereum contract, binding an instance of PresaleSecond to it.
func DeployPresaleSecond(auth *bind.TransactOpts, backend bind.ContractBackend, _maxcap *big.Int, _exceed *big.Int, _minimum *big.Int, _rate *big.Int, _wallet common.Address, _distributor common.Address, _whitelist common.Address, _token common.Address) (common.Address, *types.Transaction, *PresaleSecond, error) {
	parsed, err := abi.JSON(strings.NewReader(PresaleSecondABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PresaleSecondBin), backend, _maxcap, _exceed, _minimum, _rate, _wallet, _distributor, _whitelist, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PresaleSecond{PresaleSecondCaller: PresaleSecondCaller{contract: contract}, PresaleSecondTransactor: PresaleSecondTransactor{contract: contract}, PresaleSecondFilterer: PresaleSecondFilterer{contract: contract}}, nil
}

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
// Solidity: e Change(_addr address, _name string)
func (_PresaleSecond *PresaleSecondFilterer) FilterChange(opts *bind.FilterOpts) (*PresaleSecondChangeIterator, error) {

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Change")
	if err != nil {
		return nil, err
	}
	return &PresaleSecondChangeIterator{contract: _PresaleSecond.contract, event: "Change", logs: logs, sub: sub}, nil
}

// WatchChange is a free log subscription operation binding the contract event 0xa785fc346da73c9ad6c933dde68fe85362a97b7b07706c3e23ff3400cc5d93b5.
//
// Solidity: e Change(_addr address, _name string)
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
// Solidity: e Extinguish()
func (_PresaleSecond *PresaleSecondFilterer) FilterExtinguish(opts *bind.FilterOpts) (*PresaleSecondExtinguishIterator, error) {

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Extinguish")
	if err != nil {
		return nil, err
	}
	return &PresaleSecondExtinguishIterator{contract: _PresaleSecond.contract, event: "Extinguish", logs: logs, sub: sub}, nil
}

// WatchExtinguish is a free log subscription operation binding the contract event 0xd53036fa90376b59ca8afb9d95483e6b47ffa785d597814fcfd7405a91bba67a.
//
// Solidity: e Extinguish()
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
// Solidity: e Ignite()
func (_PresaleSecond *PresaleSecondFilterer) FilterIgnite(opts *bind.FilterOpts) (*PresaleSecondIgniteIterator, error) {

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Ignite")
	if err != nil {
		return nil, err
	}
	return &PresaleSecondIgniteIterator{contract: _PresaleSecond.contract, event: "Ignite", logs: logs, sub: sub}, nil
}

// WatchIgnite is a free log subscription operation binding the contract event 0xb66ce7cc84acb9e2cdfa3c16073eec63d39b0540887b91247afebaee4645611a.
//
// Solidity: e Ignite()
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
// Solidity: e OwnershipRenounced(previousOwner indexed address)
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
// Solidity: e OwnershipRenounced(previousOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
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
// Solidity: e Pause()
func (_PresaleSecond *PresaleSecondFilterer) FilterPause(opts *bind.FilterOpts) (*PresaleSecondPauseIterator, error) {

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PresaleSecondPauseIterator{contract: _PresaleSecond.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
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
// Solidity: e Purchase(_buyer indexed address, _purchased uint256, _refund uint256, _tokens uint256)
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
// Solidity: e Purchase(_buyer indexed address, _purchased uint256, _refund uint256, _tokens uint256)
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
// Solidity: e Refund(_to indexed address, _amount uint256)
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
// Solidity: e Refund(_to indexed address, _amount uint256)
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
// Solidity: e Release(_to indexed address, _amount uint256)
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
// Solidity: e Release(_to indexed address, _amount uint256)
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
// Solidity: e Resume()
func (_PresaleSecond *PresaleSecondFilterer) FilterResume(opts *bind.FilterOpts) (*PresaleSecondResumeIterator, error) {

	logs, sub, err := _PresaleSecond.contract.FilterLogs(opts, "Resume")
	if err != nil {
		return nil, err
	}
	return &PresaleSecondResumeIterator{contract: _PresaleSecond.contract, event: "Resume", logs: logs, sub: sub}, nil
}

// WatchResume is a free log subscription operation binding the contract event 0x490d6d11e278f168be9be39e46297f72ea877136d5bccad9cf4993e63a29568f.
//
// Solidity: e Resume()
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
// Solidity: e WithdrawEther(_from indexed address, _amount uint256)
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
// Solidity: e WithdrawEther(_from indexed address, _amount uint256)
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
// Solidity: e WithdrawToken(_from indexed address, _amount uint256)
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
// Solidity: e WithdrawToken(_from indexed address, _amount uint256)
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
