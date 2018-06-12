package provider

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type KeyStoreOption struct {
	URL       string `json:"url"`
	Context   context.Context
	Network   string `json:"network"`
	NetworkID int    `json:"network_id"`

	// Keystore provider Option
	Keypath  string
	Password string
}

func (opt *KeyStoreOption) KeystoreProvider() (provider *Provider, err error) {
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = getClient(opt)
	if err != nil {
		return
	}
	var account *bind.TransactOpts
	if account, err = getKeystoreOpt(opt); err != nil {
		return
	}
	provider.Accounts = []*bind.TransactOpts{account}
	return
}

func (opt *KeyStoreOption) Export() (err error) {
	return
}

func (opt *KeyStoreOption) Import(ctx context.Context, path string) (err error) {
	return
}
