package provider

import "context"

type PrivateKeyOption struct {
	URL       string `json:"url"`
	Context   context.Context
	Network   string `json:"network"`
	NetworkID int    `json:"network_id"`

	// PrivateKey provider Option
	Keys []string
}

func (opt *PrivateKeyOption) PrivateKeyProvider() (provider *Provider, err error) {
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = getClient(opt)
	if err != nil {
		return
	}
	provider.Accounts, err = getPrivateKeyOpts(opt)
	return
}

func (opt *PrivateKeyOption) Export() (err error) {
	return
}

func (opt *PrivateKeyOption) Import(ctx context.Context, path string) (err error) {
	return
}
