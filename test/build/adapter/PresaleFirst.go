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

// PresaleFirstABI is the input ABI used to generate the binding from.
const PresaleFirstABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_buyer\",\"type\":\"address\"}],\"name\":\"collect\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"endNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"checkRole\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exceed\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ROLE_WHITELISTED\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeAddressesFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAddressFromWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxcap\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAddressToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"release\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"buyers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"addAddressesToWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_startNumber\",\"type\":\"uint256\"},{\"name\":\"_endNumber\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Release\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"BuyTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"WhitelistedAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"RoleRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// PresaleFirstBin is the compiled bytecode used for deploying new contracts.
const PresaleFirstBin = `0x60806040526002805460ff19908116909155600a8054909116905534801561002657600080fd5b506040516080806117b2833981016040908152815160208301519183015160609093015160008054600160a060020a03191633179055909290600160a060020a038216151561007457600080fd5b600160a060020a038116151561008957600080fd5b60039390935560049190915560068054600160a060020a03928316600160a060020a0319918216179091556007805493909216921691909117905560006005556116da806100d86000396000f3006080604052600436106101695763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306ec16f88114610174578063089af913146101885780630988ca8c146101af5780630b1815671461021657806318b919e91461022b578063217fe6c6146102b557806324953eaa14610330578063286dd3f5146103855780632c4e722e146103a65780633ccfd60b146103bb5780633f4ba83a146103d05780634042b66f146103e5578063521eb273146103fa57806352d6804d1461042b57806355456f5814610440578063590e1ae3146104555780635c975abb1461046a578063715018a61461047f5780637b9417c8146104945780638456cb59146104b557806386d1a69f146104ca5780638da5cb5b146104df57806397a993aa146104f45780639b19251a146105155780639fc7a20014610536578063e2ec6ec31461054b578063f2fde38b146105a0578063fc0c546a146105c1575b610172336105d6565b005b610172600160a060020a03600435166105d6565b34801561019457600080fd5b5061019d610806565b60408051918252519081900360200190f35b3480156101bb57600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610172958335600160a060020a031695369560449491939091019190819084018382808284375094975061080c9650505050505050565b34801561022257600080fd5b5061019d61087a565b34801561023757600080fd5b50610240610887565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561027a578181015183820152602001610262565b50505050905090810190601f1680156102a75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102c157600080fd5b5060408051602060046024803582810135601f810185900485028601850190965285855261031c958335600160a060020a03169536956044949193909101919081908401838280828437509497506108ac9650505050505050565b604080519115158252519081900360200190f35b34801561033c57600080fd5b5060408051602060048035808201358381028086018501909652808552610172953695939460249493850192918291850190849080828437509497506109219650505050505050565b34801561039157600080fd5b50610172600160a060020a036004351661096e565b3480156103b257600080fd5b5061019d6109f1565b3480156103c757600080fd5b506101726109f7565b3480156103dc57600080fd5b50610172610b8c565b3480156103f157600080fd5b5061019d610be9565b34801561040657600080fd5b5061040f610bef565b60408051600160a060020a039092168252519081900360200190f35b34801561043757600080fd5b5061019d610bfe565b34801561044c57600080fd5b5061019d610c0a565b34801561046157600080fd5b50610172610c17565b34801561047657600080fd5b5061031c610d90565b34801561048b57600080fd5b50610172610d99565b3480156104a057600080fd5b50610172600160a060020a0360043516610e05565b3480156104c157600080fd5b50610172610e88565b3480156104d657600080fd5b50610172610ee7565b3480156104eb57600080fd5b5061040f61109d565b34801561050057600080fd5b5061019d600160a060020a03600435166110ac565b34801561052157600080fd5b5061031c600160a060020a03600435166110be565b34801561054257600080fd5b5061019d6110f5565b34801561055757600080fd5b5060408051602060048035808201358381028086018501909652808552610172953695939460249493850192918291850190849080828437509497506110fb9650505050505050565b3480156105ac57600080fd5b50610172600160a060020a0360043516611148565b3480156105cd57600080fd5b5061040f6111dc565b60008060006106083360408051908101604052806009815260200160008051602061168f83398151915281525061080c565b60025460ff161561061857600080fd5b600160a060020a038416151561062d57600080fd5b600554685150ae84a8cdf00000101561064557600080fd5b61064d6111eb565b151561065857600080fd5b600160a060020a038416600090815260086020526040902054681043561a88293000001161068557600080fd5b600160a060020a03841660009081526008602052604090205415156106fd57600980546001810182556000919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0386161790555b61070684611225565b9250610718348463ffffffff61123816565b604051909250600160a060020a0385169083156108fc029084906000818181858888f19350505050158015610751573d6000803e3d6000fd5b5061076483612cec63ffffffff61124a16565b60055490915061077a908463ffffffff61127316565b600555600160a060020a0384166000908152600860205260409020546107a6908463ffffffff61127316565b600160a060020a038516600081815260086020908152604091829020939093558051868152928301849052805191927f0a37b72bb67eee30e09084cf386f8a17817c57f620c3ab95fb25d6a20356ec77929081900390910190a250505050565b60045481565b610876826001836040518082805190602001908083835b602083106108425780518252601f199092019160209182019101610823565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050611280565b5050565b681043561a882930000081565b604080518082019091526009815260008051602061168f833981519152602082015281565b6000610918836001846040518082805190602001908083835b602083106108e45780518252601f1990920191602091820191016108c5565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050611295565b90505b92915050565b60008054600160a060020a0316331461093957600080fd5b5060005b815181101561087657610966828281518110151561095757fe5b9060200190602002015161096e565b60010161093d565b600054600160a060020a0316331461098557600080fd5b6109b28160408051908101604052806009815260200160008051602061168f8339815191528152506112b4565b60408051600160a060020a038316815290517ff1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a9181900360200190a150565b612cec81565b600054600160a060020a03163314610a0e57600080fd5b600654600754604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051610ac193600160a060020a039081169316916370a082319160248083019260209291908290030181600087803b158015610a7c57600080fd5b505af1158015610a90573d6000803e3d6000fd5b505050506040513d6020811015610aa657600080fd5b5051600754600160a060020a0316919063ffffffff6113d516565b600654600754604080517f70a082310000000000000000000000000000000000000000000000000000000081523060048201529051600160a060020a03938416937f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243649316916370a082319160248083019260209291908290030181600087803b158015610b4d57600080fd5b505af1158015610b61573d6000803e3d6000fd5b505050506040513d6020811015610b7757600080fd5b505160408051918252519081900360200190a2565b600054600160a060020a03163314610ba357600080fd5b60025460ff161515610bb457600080fd5b6002805460ff191690556040517f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3390600090a1565b60055481565b600654600160a060020a031681565b6706f05b59d3b2000081565b685150ae84a8cdf0000081565b60008054600160a060020a03163314610c2f57600080fd5b600a5460ff1615610c3f57600080fd5b610c47610e88565b610c4f6109f7565b5060005b600954811015610d80576009805482908110610c6b57fe5b600091825260208220015460098054600160a060020a03909216926108fc926008929086908110610c9857fe5b6000918252602080832090910154600160a060020a03168352820192909252604090810182205490518115909302929091818181858888f19350505050158015610ce6573d6000803e3d6000fd5b506009805482908110610cf557fe5b600091825260208220015460098054600160a060020a03909216927fbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d926008929086908110610d4057fe5b6000918252602080832090910154600160a060020a03168352828101939093526040918201902054815190815290519081900390910190a2600101610c53565b50600a805460ff19166001179055565b60025460ff1681565b600054600160a060020a03163314610db057600080fd5b60008054604051600160a060020a03909116917ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482091a26000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a03163314610e1c57600080fd5b610e498160408051908101604052806009815260200160008051602061168f83398151915281525061148d565b60408051600160a060020a038316815290517fd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f9181900360200190a150565b600054600160a060020a03163314610e9f57600080fd5b60025460ff1615610eaf57600080fd5b6002805460ff191660011790556040517f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62590600090a1565b60008054600160a060020a03163314610eff57600080fd5b600a5460ff1615610f0f57600080fd5b685150ae84a8cdf00000600554101580610f2b57506004544310155b1515610f3657600080fd5b600654604051600160a060020a0390911690303180156108fc02916000818181858888f19350505050158015610f70573d6000803e3d6000fd5b50600090505b6009548110156110955761100d600982815481101515610f9257fe5b600091825260208220015460098054600160a060020a0390921692610ff492612cec926008929188908110610fc357fe5b6000918252602080832090910154600160a060020a031683528201929092526040019020549063ffffffff61124a16565b600754600160a060020a0316919063ffffffff6113d516565b600980548290811061101b57fe5b9060005260206000200160009054906101000a9004600160a060020a0316600160a060020a03167ff6334794522b9db534a812aaae1af828a2e96aac68473b58e36d7d0bfd67477b61107c612cec60086000600987815481101515610fc357fe5b60408051918252519081900360200190a2600101610f76565b610d806109f7565b600054600160a060020a031681565b60086020526000908152604090205481565b60006110ed8260408051908101604052806009815260200160008051602061168f8339815191528152506108ac565b90505b919050565b60035481565b60008054600160a060020a0316331461111357600080fd5b5060005b815181101561087657611140828281518110151561113157fe5b90602001906020020151610e05565b600101611117565b600054600160a060020a0316331461115f57600080fd5b600160a060020a038116151561117457600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600754600160a060020a031681565b60008060006706f05b59d3b200003410159150600354431015801561121257506004544311155b905081801561121e5750805b9250505090565b60006110ed6112338361156e565b61160b565b60008282111561124457fe5b50900390565b600082151561125b5750600061091b565b5081810281838281151561126b57fe5b041461091b57fe5b8181018281101561091b57fe5b61128a8282611295565b151561087657600080fd5b600160a060020a03166000908152602091909152604090205460ff1690565b61131e826001836040518082805190602001908083835b602083106112ea5780518252601f1990920191602091820191016112cb565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050611647565b7fd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a82826040518083600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561139657818101518382015260200161137e565b50505050905090810190601f1680156113c35780820380516001836020036101000a031916815260200191505b50935050505060405180910390a15050565b82600160a060020a031663a9059cbb83836040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050602060405180830381600087803b15801561145157600080fd5b505af1158015611465573d6000803e3d6000fd5b505050506040513d602081101561147b57600080fd5b5051151561148857600080fd5b505050565b6114f7826001836040518082805190602001908083835b602083106114c35780518252601f1990920191602091820191016114a4565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922092915050611669565b7fbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b70048982826040518083600160a060020a0316600160a060020a0316815260200180602001828103825283818151815260200191508051906020019080838360008381101561139657818101518382015260200161137e565b6000681043561a882930000034106115905750681043561a88293000006110f0565b600160a060020a038216600090815260086020526040902054681043561a8829300000906115c590349063ffffffff61127316565b1061160457600160a060020a0382166000908152600860205260409020546115fd90681043561a88293000009063ffffffff61123816565b90506110f0565b50346110f0565b6000685150ae84a8cdf000006005548301101515611640576005546115fd90685150ae84a8cdf000009063ffffffff61123816565b50806110f0565b600160a060020a0316600090815260209190915260409020805460ff19169055565b600160a060020a0316600090815260209190915260409020805460ff19166001179055560077686974656c6973740000000000000000000000000000000000000000000000a165627a7a723058208909c96bae7940937554dbc624e8733ed68d668537a6747ff20feae4b387894f0029`

// DeployPresaleFirst deploys a new Ethereum contract, binding an instance of PresaleFirst to it.
func DeployPresaleFirst(auth *bind.TransactOpts, backend bind.ContractBackend, _startNumber *big.Int, _endNumber *big.Int, _wallet common.Address, _token common.Address) (common.Address, *types.Transaction, *PresaleFirst, error) {
	parsed, err := abi.JSON(strings.NewReader(PresaleFirstABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PresaleFirstBin), backend, _startNumber, _endNumber, _wallet, _token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PresaleFirst{PresaleFirstCaller: PresaleFirstCaller{contract: contract}, PresaleFirstTransactor: PresaleFirstTransactor{contract: contract}, PresaleFirstFilterer: PresaleFirstFilterer{contract: contract}}, nil
}

// PresaleFirst is an auto generated Go binding around an Ethereum contract.
type PresaleFirst struct {
	PresaleFirstCaller     // Read-only binding to the contract
	PresaleFirstTransactor // Write-only binding to the contract
	PresaleFirstFilterer   // Log filterer for contract events
}

// PresaleFirstCaller is an auto generated read-only Go binding around an Ethereum contract.
type PresaleFirstCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresaleFirstTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PresaleFirstTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresaleFirstFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PresaleFirstFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PresaleFirstSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PresaleFirstSession struct {
	Contract     *PresaleFirst     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PresaleFirstCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PresaleFirstCallerSession struct {
	Contract *PresaleFirstCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PresaleFirstTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PresaleFirstTransactorSession struct {
	Contract     *PresaleFirstTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PresaleFirstRaw is an auto generated low-level Go binding around an Ethereum contract.
type PresaleFirstRaw struct {
	Contract *PresaleFirst // Generic contract binding to access the raw methods on
}

// PresaleFirstCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PresaleFirstCallerRaw struct {
	Contract *PresaleFirstCaller // Generic read-only contract binding to access the raw methods on
}

// PresaleFirstTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PresaleFirstTransactorRaw struct {
	Contract *PresaleFirstTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPresaleFirst creates a new instance of PresaleFirst, bound to a specific deployed contract.
func NewPresaleFirst(address common.Address, backend bind.ContractBackend) (*PresaleFirst, error) {
	contract, err := bindPresaleFirst(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PresaleFirst{PresaleFirstCaller: PresaleFirstCaller{contract: contract}, PresaleFirstTransactor: PresaleFirstTransactor{contract: contract}, PresaleFirstFilterer: PresaleFirstFilterer{contract: contract}}, nil
}

// NewPresaleFirstCaller creates a new read-only instance of PresaleFirst, bound to a specific deployed contract.
func NewPresaleFirstCaller(address common.Address, caller bind.ContractCaller) (*PresaleFirstCaller, error) {
	contract, err := bindPresaleFirst(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PresaleFirstCaller{contract: contract}, nil
}

// NewPresaleFirstTransactor creates a new write-only instance of PresaleFirst, bound to a specific deployed contract.
func NewPresaleFirstTransactor(address common.Address, transactor bind.ContractTransactor) (*PresaleFirstTransactor, error) {
	contract, err := bindPresaleFirst(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PresaleFirstTransactor{contract: contract}, nil
}

// NewPresaleFirstFilterer creates a new log filterer instance of PresaleFirst, bound to a specific deployed contract.
func NewPresaleFirstFilterer(address common.Address, filterer bind.ContractFilterer) (*PresaleFirstFilterer, error) {
	contract, err := bindPresaleFirst(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PresaleFirstFilterer{contract: contract}, nil
}

// bindPresaleFirst binds a generic wrapper to an already deployed contract.
func bindPresaleFirst(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PresaleFirstABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PresaleFirst *PresaleFirstRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PresaleFirst.Contract.PresaleFirstCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PresaleFirst *PresaleFirstRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleFirst.Contract.PresaleFirstTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PresaleFirst *PresaleFirstRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PresaleFirst.Contract.PresaleFirstTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PresaleFirst *PresaleFirstCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PresaleFirst.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PresaleFirst *PresaleFirstTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleFirst.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PresaleFirst *PresaleFirstTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PresaleFirst.Contract.contract.Transact(opts, method, params...)
}

// ROLEWHITELISTED is a free data retrieval call binding the contract method 0x18b919e9.
//
// Solidity: function ROLE_WHITELISTED() constant returns(string)
func (_PresaleFirst *PresaleFirstCaller) ROLEWHITELISTED(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "ROLE_WHITELISTED")
	return *ret0, err
}

// ROLEWHITELISTED is a free data retrieval call binding the contract method 0x18b919e9.
//
// Solidity: function ROLE_WHITELISTED() constant returns(string)
func (_PresaleFirst *PresaleFirstSession) ROLEWHITELISTED() (string, error) {
	return _PresaleFirst.Contract.ROLEWHITELISTED(&_PresaleFirst.CallOpts)
}

// ROLEWHITELISTED is a free data retrieval call binding the contract method 0x18b919e9.
//
// Solidity: function ROLE_WHITELISTED() constant returns(string)
func (_PresaleFirst *PresaleFirstCallerSession) ROLEWHITELISTED() (string, error) {
	return _PresaleFirst.Contract.ROLEWHITELISTED(&_PresaleFirst.CallOpts)
}

// Buyers is a free data retrieval call binding the contract method 0x97a993aa.
//
// Solidity: function buyers( address) constant returns(uint256)
func (_PresaleFirst *PresaleFirstCaller) Buyers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "buyers", arg0)
	return *ret0, err
}

// Buyers is a free data retrieval call binding the contract method 0x97a993aa.
//
// Solidity: function buyers( address) constant returns(uint256)
func (_PresaleFirst *PresaleFirstSession) Buyers(arg0 common.Address) (*big.Int, error) {
	return _PresaleFirst.Contract.Buyers(&_PresaleFirst.CallOpts, arg0)
}

// Buyers is a free data retrieval call binding the contract method 0x97a993aa.
//
// Solidity: function buyers( address) constant returns(uint256)
func (_PresaleFirst *PresaleFirstCallerSession) Buyers(arg0 common.Address) (*big.Int, error) {
	return _PresaleFirst.Contract.Buyers(&_PresaleFirst.CallOpts, arg0)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_PresaleFirst *PresaleFirstCaller) CheckRole(opts *bind.CallOpts, addr common.Address, roleName string) error {
	var ()
	out := &[]interface{}{}
	err := _PresaleFirst.contract.Call(opts, out, "checkRole", addr, roleName)
	return err
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_PresaleFirst *PresaleFirstSession) CheckRole(addr common.Address, roleName string) error {
	return _PresaleFirst.Contract.CheckRole(&_PresaleFirst.CallOpts, addr, roleName)
}

// CheckRole is a free data retrieval call binding the contract method 0x0988ca8c.
//
// Solidity: function checkRole(addr address, roleName string) constant returns()
func (_PresaleFirst *PresaleFirstCallerSession) CheckRole(addr common.Address, roleName string) error {
	return _PresaleFirst.Contract.CheckRole(&_PresaleFirst.CallOpts, addr, roleName)
}

// EndNumber is a free data retrieval call binding the contract method 0x089af913.
//
// Solidity: function endNumber() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCaller) EndNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "endNumber")
	return *ret0, err
}

// EndNumber is a free data retrieval call binding the contract method 0x089af913.
//
// Solidity: function endNumber() constant returns(uint256)
func (_PresaleFirst *PresaleFirstSession) EndNumber() (*big.Int, error) {
	return _PresaleFirst.Contract.EndNumber(&_PresaleFirst.CallOpts)
}

// EndNumber is a free data retrieval call binding the contract method 0x089af913.
//
// Solidity: function endNumber() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCallerSession) EndNumber() (*big.Int, error) {
	return _PresaleFirst.Contract.EndNumber(&_PresaleFirst.CallOpts)
}

// Exceed is a free data retrieval call binding the contract method 0x0b181567.
//
// Solidity: function exceed() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCaller) Exceed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "exceed")
	return *ret0, err
}

// Exceed is a free data retrieval call binding the contract method 0x0b181567.
//
// Solidity: function exceed() constant returns(uint256)
func (_PresaleFirst *PresaleFirstSession) Exceed() (*big.Int, error) {
	return _PresaleFirst.Contract.Exceed(&_PresaleFirst.CallOpts)
}

// Exceed is a free data retrieval call binding the contract method 0x0b181567.
//
// Solidity: function exceed() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCallerSession) Exceed() (*big.Int, error) {
	return _PresaleFirst.Contract.Exceed(&_PresaleFirst.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_PresaleFirst *PresaleFirstCaller) HasRole(opts *bind.CallOpts, addr common.Address, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "hasRole", addr, roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_PresaleFirst *PresaleFirstSession) HasRole(addr common.Address, roleName string) (bool, error) {
	return _PresaleFirst.Contract.HasRole(&_PresaleFirst.CallOpts, addr, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0x217fe6c6.
//
// Solidity: function hasRole(addr address, roleName string) constant returns(bool)
func (_PresaleFirst *PresaleFirstCallerSession) HasRole(addr common.Address, roleName string) (bool, error) {
	return _PresaleFirst.Contract.HasRole(&_PresaleFirst.CallOpts, addr, roleName)
}

// Maxcap is a free data retrieval call binding the contract method 0x55456f58.
//
// Solidity: function maxcap() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCaller) Maxcap(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "maxcap")
	return *ret0, err
}

// Maxcap is a free data retrieval call binding the contract method 0x55456f58.
//
// Solidity: function maxcap() constant returns(uint256)
func (_PresaleFirst *PresaleFirstSession) Maxcap() (*big.Int, error) {
	return _PresaleFirst.Contract.Maxcap(&_PresaleFirst.CallOpts)
}

// Maxcap is a free data retrieval call binding the contract method 0x55456f58.
//
// Solidity: function maxcap() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCallerSession) Maxcap() (*big.Int, error) {
	return _PresaleFirst.Contract.Maxcap(&_PresaleFirst.CallOpts)
}

// Minimum is a free data retrieval call binding the contract method 0x52d6804d.
//
// Solidity: function minimum() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCaller) Minimum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "minimum")
	return *ret0, err
}

// Minimum is a free data retrieval call binding the contract method 0x52d6804d.
//
// Solidity: function minimum() constant returns(uint256)
func (_PresaleFirst *PresaleFirstSession) Minimum() (*big.Int, error) {
	return _PresaleFirst.Contract.Minimum(&_PresaleFirst.CallOpts)
}

// Minimum is a free data retrieval call binding the contract method 0x52d6804d.
//
// Solidity: function minimum() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCallerSession) Minimum() (*big.Int, error) {
	return _PresaleFirst.Contract.Minimum(&_PresaleFirst.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PresaleFirst *PresaleFirstCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PresaleFirst *PresaleFirstSession) Owner() (common.Address, error) {
	return _PresaleFirst.Contract.Owner(&_PresaleFirst.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PresaleFirst *PresaleFirstCallerSession) Owner() (common.Address, error) {
	return _PresaleFirst.Contract.Owner(&_PresaleFirst.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PresaleFirst *PresaleFirstCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PresaleFirst *PresaleFirstSession) Paused() (bool, error) {
	return _PresaleFirst.Contract.Paused(&_PresaleFirst.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PresaleFirst *PresaleFirstCallerSession) Paused() (bool, error) {
	return _PresaleFirst.Contract.Paused(&_PresaleFirst.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PresaleFirst *PresaleFirstSession) Rate() (*big.Int, error) {
	return _PresaleFirst.Contract.Rate(&_PresaleFirst.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCallerSession) Rate() (*big.Int, error) {
	return _PresaleFirst.Contract.Rate(&_PresaleFirst.CallOpts)
}

// StartNumber is a free data retrieval call binding the contract method 0x9fc7a200.
//
// Solidity: function startNumber() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCaller) StartNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "startNumber")
	return *ret0, err
}

// StartNumber is a free data retrieval call binding the contract method 0x9fc7a200.
//
// Solidity: function startNumber() constant returns(uint256)
func (_PresaleFirst *PresaleFirstSession) StartNumber() (*big.Int, error) {
	return _PresaleFirst.Contract.StartNumber(&_PresaleFirst.CallOpts)
}

// StartNumber is a free data retrieval call binding the contract method 0x9fc7a200.
//
// Solidity: function startNumber() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCallerSession) StartNumber() (*big.Int, error) {
	return _PresaleFirst.Contract.StartNumber(&_PresaleFirst.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PresaleFirst *PresaleFirstCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PresaleFirst *PresaleFirstSession) Token() (common.Address, error) {
	return _PresaleFirst.Contract.Token(&_PresaleFirst.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_PresaleFirst *PresaleFirstCallerSession) Token() (common.Address, error) {
	return _PresaleFirst.Contract.Token(&_PresaleFirst.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PresaleFirst *PresaleFirstCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PresaleFirst *PresaleFirstSession) Wallet() (common.Address, error) {
	return _PresaleFirst.Contract.Wallet(&_PresaleFirst.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_PresaleFirst *PresaleFirstCallerSession) Wallet() (common.Address, error) {
	return _PresaleFirst.Contract.Wallet(&_PresaleFirst.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PresaleFirst *PresaleFirstSession) WeiRaised() (*big.Int, error) {
	return _PresaleFirst.Contract.WeiRaised(&_PresaleFirst.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_PresaleFirst *PresaleFirstCallerSession) WeiRaised() (*big.Int, error) {
	return _PresaleFirst.Contract.WeiRaised(&_PresaleFirst.CallOpts)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(addr address) constant returns(bool)
func (_PresaleFirst *PresaleFirstCaller) Whitelist(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PresaleFirst.contract.Call(opts, out, "whitelist", addr)
	return *ret0, err
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(addr address) constant returns(bool)
func (_PresaleFirst *PresaleFirstSession) Whitelist(addr common.Address) (bool, error) {
	return _PresaleFirst.Contract.Whitelist(&_PresaleFirst.CallOpts, addr)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(addr address) constant returns(bool)
func (_PresaleFirst *PresaleFirstCallerSession) Whitelist(addr common.Address) (bool, error) {
	return _PresaleFirst.Contract.Whitelist(&_PresaleFirst.CallOpts, addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns()
func (_PresaleFirst *PresaleFirstTransactor) AddAddressToWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "addAddressToWhitelist", addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns()
func (_PresaleFirst *PresaleFirstSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.AddAddressToWhitelist(&_PresaleFirst.TransactOpts, addr)
}

// AddAddressToWhitelist is a paid mutator transaction binding the contract method 0x7b9417c8.
//
// Solidity: function addAddressToWhitelist(addr address) returns()
func (_PresaleFirst *PresaleFirstTransactorSession) AddAddressToWhitelist(addr common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.AddAddressToWhitelist(&_PresaleFirst.TransactOpts, addr)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns()
func (_PresaleFirst *PresaleFirstTransactor) AddAddressesToWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "addAddressesToWhitelist", addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns()
func (_PresaleFirst *PresaleFirstSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.AddAddressesToWhitelist(&_PresaleFirst.TransactOpts, addrs)
}

// AddAddressesToWhitelist is a paid mutator transaction binding the contract method 0xe2ec6ec3.
//
// Solidity: function addAddressesToWhitelist(addrs address[]) returns()
func (_PresaleFirst *PresaleFirstTransactorSession) AddAddressesToWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.AddAddressesToWhitelist(&_PresaleFirst.TransactOpts, addrs)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(_buyer address) returns()
func (_PresaleFirst *PresaleFirstTransactor) Collect(opts *bind.TransactOpts, _buyer common.Address) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "collect", _buyer)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(_buyer address) returns()
func (_PresaleFirst *PresaleFirstSession) Collect(_buyer common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.Collect(&_PresaleFirst.TransactOpts, _buyer)
}

// Collect is a paid mutator transaction binding the contract method 0x06ec16f8.
//
// Solidity: function collect(_buyer address) returns()
func (_PresaleFirst *PresaleFirstTransactorSession) Collect(_buyer common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.Collect(&_PresaleFirst.TransactOpts, _buyer)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PresaleFirst *PresaleFirstTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PresaleFirst *PresaleFirstSession) Pause() (*types.Transaction, error) {
	return _PresaleFirst.Contract.Pause(&_PresaleFirst.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PresaleFirst *PresaleFirstTransactorSession) Pause() (*types.Transaction, error) {
	return _PresaleFirst.Contract.Pause(&_PresaleFirst.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_PresaleFirst *PresaleFirstTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_PresaleFirst *PresaleFirstSession) Refund() (*types.Transaction, error) {
	return _PresaleFirst.Contract.Refund(&_PresaleFirst.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns()
func (_PresaleFirst *PresaleFirstTransactorSession) Refund() (*types.Transaction, error) {
	return _PresaleFirst.Contract.Refund(&_PresaleFirst.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_PresaleFirst *PresaleFirstTransactor) Release(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "release")
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_PresaleFirst *PresaleFirstSession) Release() (*types.Transaction, error) {
	return _PresaleFirst.Contract.Release(&_PresaleFirst.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_PresaleFirst *PresaleFirstTransactorSession) Release() (*types.Transaction, error) {
	return _PresaleFirst.Contract.Release(&_PresaleFirst.TransactOpts)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns()
func (_PresaleFirst *PresaleFirstTransactor) RemoveAddressFromWhitelist(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "removeAddressFromWhitelist", addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns()
func (_PresaleFirst *PresaleFirstSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.RemoveAddressFromWhitelist(&_PresaleFirst.TransactOpts, addr)
}

// RemoveAddressFromWhitelist is a paid mutator transaction binding the contract method 0x286dd3f5.
//
// Solidity: function removeAddressFromWhitelist(addr address) returns()
func (_PresaleFirst *PresaleFirstTransactorSession) RemoveAddressFromWhitelist(addr common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.RemoveAddressFromWhitelist(&_PresaleFirst.TransactOpts, addr)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns()
func (_PresaleFirst *PresaleFirstTransactor) RemoveAddressesFromWhitelist(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "removeAddressesFromWhitelist", addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns()
func (_PresaleFirst *PresaleFirstSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.RemoveAddressesFromWhitelist(&_PresaleFirst.TransactOpts, addrs)
}

// RemoveAddressesFromWhitelist is a paid mutator transaction binding the contract method 0x24953eaa.
//
// Solidity: function removeAddressesFromWhitelist(addrs address[]) returns()
func (_PresaleFirst *PresaleFirstTransactorSession) RemoveAddressesFromWhitelist(addrs []common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.RemoveAddressesFromWhitelist(&_PresaleFirst.TransactOpts, addrs)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PresaleFirst *PresaleFirstTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PresaleFirst *PresaleFirstSession) RenounceOwnership() (*types.Transaction, error) {
	return _PresaleFirst.Contract.RenounceOwnership(&_PresaleFirst.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PresaleFirst *PresaleFirstTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PresaleFirst.Contract.RenounceOwnership(&_PresaleFirst.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PresaleFirst *PresaleFirstTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PresaleFirst *PresaleFirstSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.TransferOwnership(&_PresaleFirst.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PresaleFirst *PresaleFirstTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PresaleFirst.Contract.TransferOwnership(&_PresaleFirst.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PresaleFirst *PresaleFirstTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PresaleFirst *PresaleFirstSession) Unpause() (*types.Transaction, error) {
	return _PresaleFirst.Contract.Unpause(&_PresaleFirst.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PresaleFirst *PresaleFirstTransactorSession) Unpause() (*types.Transaction, error) {
	return _PresaleFirst.Contract.Unpause(&_PresaleFirst.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PresaleFirst *PresaleFirstTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PresaleFirst.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PresaleFirst *PresaleFirstSession) Withdraw() (*types.Transaction, error) {
	return _PresaleFirst.Contract.Withdraw(&_PresaleFirst.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_PresaleFirst *PresaleFirstTransactorSession) Withdraw() (*types.Transaction, error) {
	return _PresaleFirst.Contract.Withdraw(&_PresaleFirst.TransactOpts)
}

// PresaleFirstBuyTokensIterator is returned from FilterBuyTokens and is used to iterate over the raw logs and unpacked data for BuyTokens events raised by the PresaleFirst contract.
type PresaleFirstBuyTokensIterator struct {
	Event *PresaleFirstBuyTokens // Event containing the contract specifics and raw log

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
func (it *PresaleFirstBuyTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstBuyTokens)
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
		it.Event = new(PresaleFirstBuyTokens)
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
func (it *PresaleFirstBuyTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstBuyTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstBuyTokens represents a BuyTokens event raised by the PresaleFirst contract.
type PresaleFirstBuyTokens struct {
	Buyer  common.Address
	Price  *big.Int
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuyTokens is a free log retrieval operation binding the contract event 0x0a37b72bb67eee30e09084cf386f8a17817c57f620c3ab95fb25d6a20356ec77.
//
// Solidity: e BuyTokens(buyer indexed address, price uint256, tokens uint256)
func (_PresaleFirst *PresaleFirstFilterer) FilterBuyTokens(opts *bind.FilterOpts, buyer []common.Address) (*PresaleFirstBuyTokensIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "BuyTokens", buyerRule)
	if err != nil {
		return nil, err
	}
	return &PresaleFirstBuyTokensIterator{contract: _PresaleFirst.contract, event: "BuyTokens", logs: logs, sub: sub}, nil
}

// WatchBuyTokens is a free log subscription operation binding the contract event 0x0a37b72bb67eee30e09084cf386f8a17817c57f620c3ab95fb25d6a20356ec77.
//
// Solidity: e BuyTokens(buyer indexed address, price uint256, tokens uint256)
func (_PresaleFirst *PresaleFirstFilterer) WatchBuyTokens(opts *bind.WatchOpts, sink chan<- *PresaleFirstBuyTokens, buyer []common.Address) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "BuyTokens", buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstBuyTokens)
				if err := _PresaleFirst.contract.UnpackLog(event, "BuyTokens", log); err != nil {
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

// PresaleFirstOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the PresaleFirst contract.
type PresaleFirstOwnershipRenouncedIterator struct {
	Event *PresaleFirstOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *PresaleFirstOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstOwnershipRenounced)
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
		it.Event = new(PresaleFirstOwnershipRenounced)
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
func (it *PresaleFirstOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstOwnershipRenounced represents a OwnershipRenounced event raised by the PresaleFirst contract.
type PresaleFirstOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PresaleFirst *PresaleFirstFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*PresaleFirstOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PresaleFirstOwnershipRenouncedIterator{contract: _PresaleFirst.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_PresaleFirst *PresaleFirstFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *PresaleFirstOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstOwnershipRenounced)
				if err := _PresaleFirst.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// PresaleFirstOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PresaleFirst contract.
type PresaleFirstOwnershipTransferredIterator struct {
	Event *PresaleFirstOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PresaleFirstOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstOwnershipTransferred)
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
		it.Event = new(PresaleFirstOwnershipTransferred)
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
func (it *PresaleFirstOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstOwnershipTransferred represents a OwnershipTransferred event raised by the PresaleFirst contract.
type PresaleFirstOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PresaleFirst *PresaleFirstFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PresaleFirstOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PresaleFirstOwnershipTransferredIterator{contract: _PresaleFirst.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PresaleFirst *PresaleFirstFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PresaleFirstOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstOwnershipTransferred)
				if err := _PresaleFirst.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PresaleFirstPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PresaleFirst contract.
type PresaleFirstPauseIterator struct {
	Event *PresaleFirstPause // Event containing the contract specifics and raw log

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
func (it *PresaleFirstPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstPause)
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
		it.Event = new(PresaleFirstPause)
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
func (it *PresaleFirstPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstPause represents a Pause event raised by the PresaleFirst contract.
type PresaleFirstPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PresaleFirst *PresaleFirstFilterer) FilterPause(opts *bind.FilterOpts) (*PresaleFirstPauseIterator, error) {

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PresaleFirstPauseIterator{contract: _PresaleFirst.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PresaleFirst *PresaleFirstFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PresaleFirstPause) (event.Subscription, error) {

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstPause)
				if err := _PresaleFirst.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PresaleFirstRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the PresaleFirst contract.
type PresaleFirstRefundIterator struct {
	Event *PresaleFirstRefund // Event containing the contract specifics and raw log

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
func (it *PresaleFirstRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstRefund)
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
		it.Event = new(PresaleFirstRefund)
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
func (it *PresaleFirstRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstRefund represents a Refund event raised by the PresaleFirst contract.
type PresaleFirstRefund struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0xbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d.
//
// Solidity: e Refund(_to indexed address, _amount uint256)
func (_PresaleFirst *PresaleFirstFilterer) FilterRefund(opts *bind.FilterOpts, _to []common.Address) (*PresaleFirstRefundIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "Refund", _toRule)
	if err != nil {
		return nil, err
	}
	return &PresaleFirstRefundIterator{contract: _PresaleFirst.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0xbb28353e4598c3b9199101a66e0989549b659a59a54d2c27fbb183f1932c8e6d.
//
// Solidity: e Refund(_to indexed address, _amount uint256)
func (_PresaleFirst *PresaleFirstFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *PresaleFirstRefund, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "Refund", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstRefund)
				if err := _PresaleFirst.contract.UnpackLog(event, "Refund", log); err != nil {
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

// PresaleFirstReleaseIterator is returned from FilterRelease and is used to iterate over the raw logs and unpacked data for Release events raised by the PresaleFirst contract.
type PresaleFirstReleaseIterator struct {
	Event *PresaleFirstRelease // Event containing the contract specifics and raw log

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
func (it *PresaleFirstReleaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstRelease)
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
		it.Event = new(PresaleFirstRelease)
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
func (it *PresaleFirstReleaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstReleaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstRelease represents a Release event raised by the PresaleFirst contract.
type PresaleFirstRelease struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRelease is a free log retrieval operation binding the contract event 0xf6334794522b9db534a812aaae1af828a2e96aac68473b58e36d7d0bfd67477b.
//
// Solidity: e Release(_to indexed address, _amount uint256)
func (_PresaleFirst *PresaleFirstFilterer) FilterRelease(opts *bind.FilterOpts, _to []common.Address) (*PresaleFirstReleaseIterator, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "Release", _toRule)
	if err != nil {
		return nil, err
	}
	return &PresaleFirstReleaseIterator{contract: _PresaleFirst.contract, event: "Release", logs: logs, sub: sub}, nil
}

// WatchRelease is a free log subscription operation binding the contract event 0xf6334794522b9db534a812aaae1af828a2e96aac68473b58e36d7d0bfd67477b.
//
// Solidity: e Release(_to indexed address, _amount uint256)
func (_PresaleFirst *PresaleFirstFilterer) WatchRelease(opts *bind.WatchOpts, sink chan<- *PresaleFirstRelease, _to []common.Address) (event.Subscription, error) {

	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "Release", _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstRelease)
				if err := _PresaleFirst.contract.UnpackLog(event, "Release", log); err != nil {
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

// PresaleFirstRoleAddedIterator is returned from FilterRoleAdded and is used to iterate over the raw logs and unpacked data for RoleAdded events raised by the PresaleFirst contract.
type PresaleFirstRoleAddedIterator struct {
	Event *PresaleFirstRoleAdded // Event containing the contract specifics and raw log

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
func (it *PresaleFirstRoleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstRoleAdded)
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
		it.Event = new(PresaleFirstRoleAdded)
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
func (it *PresaleFirstRoleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstRoleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstRoleAdded represents a RoleAdded event raised by the PresaleFirst contract.
type PresaleFirstRoleAdded struct {
	Addr     common.Address
	RoleName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleAdded is a free log retrieval operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(addr address, roleName string)
func (_PresaleFirst *PresaleFirstFilterer) FilterRoleAdded(opts *bind.FilterOpts) (*PresaleFirstRoleAddedIterator, error) {

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "RoleAdded")
	if err != nil {
		return nil, err
	}
	return &PresaleFirstRoleAddedIterator{contract: _PresaleFirst.contract, event: "RoleAdded", logs: logs, sub: sub}, nil
}

// WatchRoleAdded is a free log subscription operation binding the contract event 0xbfec83d64eaa953f2708271a023ab9ee82057f8f3578d548c1a4ba0b5b700489.
//
// Solidity: e RoleAdded(addr address, roleName string)
func (_PresaleFirst *PresaleFirstFilterer) WatchRoleAdded(opts *bind.WatchOpts, sink chan<- *PresaleFirstRoleAdded) (event.Subscription, error) {

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "RoleAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstRoleAdded)
				if err := _PresaleFirst.contract.UnpackLog(event, "RoleAdded", log); err != nil {
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

// PresaleFirstRoleRemovedIterator is returned from FilterRoleRemoved and is used to iterate over the raw logs and unpacked data for RoleRemoved events raised by the PresaleFirst contract.
type PresaleFirstRoleRemovedIterator struct {
	Event *PresaleFirstRoleRemoved // Event containing the contract specifics and raw log

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
func (it *PresaleFirstRoleRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstRoleRemoved)
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
		it.Event = new(PresaleFirstRoleRemoved)
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
func (it *PresaleFirstRoleRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstRoleRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstRoleRemoved represents a RoleRemoved event raised by the PresaleFirst contract.
type PresaleFirstRoleRemoved struct {
	Addr     common.Address
	RoleName string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRoleRemoved is a free log retrieval operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(addr address, roleName string)
func (_PresaleFirst *PresaleFirstFilterer) FilterRoleRemoved(opts *bind.FilterOpts) (*PresaleFirstRoleRemovedIterator, error) {

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "RoleRemoved")
	if err != nil {
		return nil, err
	}
	return &PresaleFirstRoleRemovedIterator{contract: _PresaleFirst.contract, event: "RoleRemoved", logs: logs, sub: sub}, nil
}

// WatchRoleRemoved is a free log subscription operation binding the contract event 0xd211483f91fc6eff862467f8de606587a30c8fc9981056f051b897a418df803a.
//
// Solidity: e RoleRemoved(addr address, roleName string)
func (_PresaleFirst *PresaleFirstFilterer) WatchRoleRemoved(opts *bind.WatchOpts, sink chan<- *PresaleFirstRoleRemoved) (event.Subscription, error) {

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "RoleRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstRoleRemoved)
				if err := _PresaleFirst.contract.UnpackLog(event, "RoleRemoved", log); err != nil {
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

// PresaleFirstUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PresaleFirst contract.
type PresaleFirstUnpauseIterator struct {
	Event *PresaleFirstUnpause // Event containing the contract specifics and raw log

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
func (it *PresaleFirstUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstUnpause)
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
		it.Event = new(PresaleFirstUnpause)
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
func (it *PresaleFirstUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstUnpause represents a Unpause event raised by the PresaleFirst contract.
type PresaleFirstUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PresaleFirst *PresaleFirstFilterer) FilterUnpause(opts *bind.FilterOpts) (*PresaleFirstUnpauseIterator, error) {

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PresaleFirstUnpauseIterator{contract: _PresaleFirst.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PresaleFirst *PresaleFirstFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PresaleFirstUnpause) (event.Subscription, error) {

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstUnpause)
				if err := _PresaleFirst.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// PresaleFirstWhitelistedAddressAddedIterator is returned from FilterWhitelistedAddressAdded and is used to iterate over the raw logs and unpacked data for WhitelistedAddressAdded events raised by the PresaleFirst contract.
type PresaleFirstWhitelistedAddressAddedIterator struct {
	Event *PresaleFirstWhitelistedAddressAdded // Event containing the contract specifics and raw log

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
func (it *PresaleFirstWhitelistedAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstWhitelistedAddressAdded)
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
		it.Event = new(PresaleFirstWhitelistedAddressAdded)
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
func (it *PresaleFirstWhitelistedAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstWhitelistedAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstWhitelistedAddressAdded represents a WhitelistedAddressAdded event raised by the PresaleFirst contract.
type PresaleFirstWhitelistedAddressAdded struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressAdded is a free log retrieval operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: e WhitelistedAddressAdded(addr address)
func (_PresaleFirst *PresaleFirstFilterer) FilterWhitelistedAddressAdded(opts *bind.FilterOpts) (*PresaleFirstWhitelistedAddressAddedIterator, error) {

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return &PresaleFirstWhitelistedAddressAddedIterator{contract: _PresaleFirst.contract, event: "WhitelistedAddressAdded", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressAdded is a free log subscription operation binding the contract event 0xd1bba68c128cc3f427e5831b3c6f99f480b6efa6b9e80c757768f6124158cc3f.
//
// Solidity: e WhitelistedAddressAdded(addr address)
func (_PresaleFirst *PresaleFirstFilterer) WatchWhitelistedAddressAdded(opts *bind.WatchOpts, sink chan<- *PresaleFirstWhitelistedAddressAdded) (event.Subscription, error) {

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "WhitelistedAddressAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstWhitelistedAddressAdded)
				if err := _PresaleFirst.contract.UnpackLog(event, "WhitelistedAddressAdded", log); err != nil {
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

// PresaleFirstWhitelistedAddressRemovedIterator is returned from FilterWhitelistedAddressRemoved and is used to iterate over the raw logs and unpacked data for WhitelistedAddressRemoved events raised by the PresaleFirst contract.
type PresaleFirstWhitelistedAddressRemovedIterator struct {
	Event *PresaleFirstWhitelistedAddressRemoved // Event containing the contract specifics and raw log

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
func (it *PresaleFirstWhitelistedAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstWhitelistedAddressRemoved)
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
		it.Event = new(PresaleFirstWhitelistedAddressRemoved)
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
func (it *PresaleFirstWhitelistedAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstWhitelistedAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstWhitelistedAddressRemoved represents a WhitelistedAddressRemoved event raised by the PresaleFirst contract.
type PresaleFirstWhitelistedAddressRemoved struct {
	Addr common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterWhitelistedAddressRemoved is a free log retrieval operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: e WhitelistedAddressRemoved(addr address)
func (_PresaleFirst *PresaleFirstFilterer) FilterWhitelistedAddressRemoved(opts *bind.FilterOpts) (*PresaleFirstWhitelistedAddressRemovedIterator, error) {

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return &PresaleFirstWhitelistedAddressRemovedIterator{contract: _PresaleFirst.contract, event: "WhitelistedAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchWhitelistedAddressRemoved is a free log subscription operation binding the contract event 0xf1abf01a1043b7c244d128e8595cf0c1d10743b022b03a02dffd8ca3bf729f5a.
//
// Solidity: e WhitelistedAddressRemoved(addr address)
func (_PresaleFirst *PresaleFirstFilterer) WatchWhitelistedAddressRemoved(opts *bind.WatchOpts, sink chan<- *PresaleFirstWhitelistedAddressRemoved) (event.Subscription, error) {

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "WhitelistedAddressRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstWhitelistedAddressRemoved)
				if err := _PresaleFirst.contract.UnpackLog(event, "WhitelistedAddressRemoved", log); err != nil {
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

// PresaleFirstWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PresaleFirst contract.
type PresaleFirstWithdrawIterator struct {
	Event *PresaleFirstWithdraw // Event containing the contract specifics and raw log

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
func (it *PresaleFirstWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PresaleFirstWithdraw)
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
		it.Event = new(PresaleFirstWithdraw)
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
func (it *PresaleFirstWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PresaleFirstWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PresaleFirstWithdraw represents a Withdraw event raised by the PresaleFirst contract.
type PresaleFirstWithdraw struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: e Withdraw(_from indexed address, _amount uint256)
func (_PresaleFirst *PresaleFirstFilterer) FilterWithdraw(opts *bind.FilterOpts, _from []common.Address) (*PresaleFirstWithdrawIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _PresaleFirst.contract.FilterLogs(opts, "Withdraw", _fromRule)
	if err != nil {
		return nil, err
	}
	return &PresaleFirstWithdrawIterator{contract: _PresaleFirst.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: e Withdraw(_from indexed address, _amount uint256)
func (_PresaleFirst *PresaleFirstFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PresaleFirstWithdraw, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _PresaleFirst.contract.WatchLogs(opts, "Withdraw", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PresaleFirstWithdraw)
				if err := _PresaleFirst.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
