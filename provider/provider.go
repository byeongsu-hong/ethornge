package provider

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	DefaultLedgerOption = &Option{
		Path:  "m/44'/60'/0'/",
		Start: 0,
		End:   1,
	}
)

type Option struct {
	// Common option
	URL       string
	Context   context.Context
	Network   string
	NetworkID int

	// PrivateKey provider Option
	Keys []string

	// Ledger provider Option
	Path  string
	Start int
	End   int
}

type Provider struct {
	*ethclient.Client
	Context  context.Context
	Accounts []*bind.TransactOpts
}

func LedgerProvider(opt *Option) (provider *Provider, err error) {
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = ethclient.Dial(opt.URL)
	if err != nil {
		return
	}
	provider.Accounts, err = getLedgerOpts(opt)
	return
}

func PrivateKeyProvider(opt *Option) (provider *Provider, err error) {
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = ethclient.Dial(opt.URL)
	if err != nil {
		return
	}
	provider.Accounts, err = getPrivateKeyOpts(opt)
	return
}
