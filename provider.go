package ethornge

import (
	"context"

	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	DefaultLedgerOption = &ProviderOption{
		Path:  "m/44'/60'/0'/",
		Start: 0,
		End:   1,
	}
)

type ProviderOption struct {
	// Common option
	URL       string
	Context   context.Context
	Network   string
	NetworkID int

	// Private Key Option
	Keys []string

	// Ledger Provider Option
	Path  string
	Start int
	End   int
}

type Provider struct {
	Client   *ethclient.Client
	Context  context.Context
	Accounts TxOpts
}

func LedgerProvider(opt *ProviderOption) (provider *Provider, err error) {
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = ethclient.Dial(opt.URL)
	if err != nil {
		return
	}
	provider.Accounts, err = GetLedgerOpts(opt)
	return
}

func PrivateKeyProvider(opt *ProviderOption) (provider *Provider, err error) {
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = ethclient.Dial(opt.URL)
	if err != nil {
		return
	}
	provider.Accounts, err = GetPrivateKeyOpts(opt)
	return
}
