package ethornge

const (
	MAINNET = 1
	ROPSTEN = 3
	RINKEBY = 4
	KOVAN   = 42
)

func convertId(id int) string {
	switch id {
	case MAINNET:
		return "mainnet"
	case ROPSTEN:
		return "ropsten"
	case RINKEBY:
		return "rinkeby"
	case KOVAN:
		return "koven"
	default:
		return ""
	}
}
