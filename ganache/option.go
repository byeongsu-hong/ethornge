package ganache

import (
	"math/big"

	"fmt"

	"strings"

	"../utils"
)

type Option struct {
	Accounts  Accounts
	Port      *big.Int
	NetworkId *big.Int
}

func Convert(x *big.Int, def string) string {
	if x == nil {
		return def
	} else {
		return x.String()
	}
}

func (opt *Option) Url() string {
	return "http://localhost:" + Convert(opt.Port, "8545")
}

func (opt *Option) generate(ganache string) string {
	var args []string
	args = append(args, ganache)
	args = append(args, "-p", Convert(opt.Port, "8545"))
	args = append(args, "-i", Convert(opt.NetworkId, fmt.Sprint(utils.ETHORNGE)))
	for _, account := range opt.Accounts {
		args = append(args, "--account", fmt.Sprintf(`"0x%s,%s"`, account.PrivateKey, account.Balance))
	}
	var builder strings.Builder
	for _, arg := range args {
		builder.WriteString(arg + " ")
	}
	return builder.String()
}
