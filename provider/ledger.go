package provider

import "context"

type LedgerOption struct {
	URL       string `json:"url"`
	Context   context.Context
	Network   string `json:"network"`
	NetworkID int    `json:"network_id"`

	// Ledger provider Option
	Path  string `json:"path"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

func (opt *LedgerOption) LedgerProvider() (provider *Provider, err error) {
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = getClient(opt)
	if err != nil {
		return
	}
	provider.Accounts, err = getLedgerOpts(opt)
	return
}

func (opt *LedgerOption) Export() (err error) {
	return
}

func (opt *LedgerOption) Import(ctx context.Context, path string) (err error) {
	opt.Context = ctx
	return
}

func (opt *LedgerOption) ImportProvider(ctx context.Context, path string) (provider *Provider, err error) {
	opt.Import(ctx, path)
	return opt.LedgerProvider()
}
