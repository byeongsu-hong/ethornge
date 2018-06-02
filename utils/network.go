package utils

const (
	MAINNET  = 1
	ROPSTEN  = 3
	RINKEBY  = 4
	KOVAN    = 42
	ETHORNGE = 2470
)

func ConvertId(id int) string {
	switch id {
	case MAINNET:
		return "mainnet"
	case ROPSTEN:
		return "ropsten"
	case RINKEBY:
		return "rinkeby"
	case KOVAN:
		return "koven"
	case ETHORNGE:
		return "ethornge"
	default:
		return ""
	}
}
