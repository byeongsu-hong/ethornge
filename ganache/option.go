package ganache

import (
	"math/big"

	"fmt"

	"strings"

	"../utils"
)

type Option struct {
	// account settings
	Accounts Accounts

	// chain settings
	GasPrice  *big.Int
	GasLimit  *big.Int
	BlockTime *big.Int

	// network settings
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

func (opt *Option) generate(ganache string) string {
	var args []string
	/*
		GasPrice  *big.Int
		GasLimit  *big.Int
		BlockTime *big.Int
	*/
	args = append(args, ganache)
	args = append(args, "-b", Convert(opt.BlockTime, "0"))
	args = append(args, "-g", Convert(opt.GasPrice, "20000000000"))
	args = append(args, "-l", Convert(opt.GasLimit, "90000"))

	/*
		Port      *big.Int
		Hostname  string
		NetworkId *big.Int
	*/
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
