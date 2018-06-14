package ganache

import (
	"fmt"
	"math/big"
	"strings"

	"context"

	"github.com/frostornge/ethornge/account"
	"github.com/frostornge/ethornge/provider"
	"github.com/frostornge/ethornge/utils"
)

type Option struct {
	Accounts  account.Accounts
	Port      *big.Int
	NetworkId *big.Int
}

func convert(x *big.Int, def string) string {
	if x == nil {
		return def
	} else {
		return x.String()
	}
}

func (opt *Option) Provider(ctx context.Context, network string) (pv *provider.Provider, err error) {
	n := make(map[string]*provider.Network)
	n[network] = &provider.Network{
		URL:       opt.Url(),
		Network:   network,
		NetworkID: opt.NetworkId,
	}
	return (&provider.PrivateKeyOption{
		Context: ctx,
		Network: n,
		PrivateKey: &provider.PrivateKey{
			Keys: opt.Accounts.GetKeys(),
		},
	}).Provider(network)
}

func (opt *Option) Url() string {
	return "http://localhost:" + convert(opt.Port, "8545")
}

func (opt *Option) generate(ganache string) string {
	var args []string
	args = append(args, ganache)
	args = append(args, "-p", convert(opt.Port, "8545"))
	args = append(args, "-i", convert(opt.NetworkId, fmt.Sprint(utils.ETHORNGE)))
	for _, account := range opt.Accounts {
		args = append(args, "--account", fmt.Sprintf(`"0x%s,%s"`, account.PrivateKey, account.Balance))
	}
	var builder strings.Builder
	for _, arg := range args {
		builder.WriteString(arg + " ")
	}
	return builder.String()
}
