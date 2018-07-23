package provider

import (
	"context"
	"io"

	"os"

	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/frostornge/ethornge/utils"
)

type Keystore struct {
	// Keystore provider Option
	Keypath  string `json:"keypath"`
	Password string `json:"password"`
}

type KeystoreOption struct {
	Context  context.Context
	Network  map[string]*Network `json:"network"`
	Keystore *Keystore           `json:"keystore"`
}

func (opt *KeystoreOption) getAccounts() (opts []*bind.TransactOpts, err error) {
	opts = make([]*bind.TransactOpts, 1)
	var keystore = opt.Keystore
	var r io.Reader
	if r, err = os.Open(keystore.Keypath); err != nil {
		return
	}
	if opts[0], err = bind.NewTransactor(r, keystore.Password); err != nil {
		return
	}
	opts[0].Context = opt.Context
	return
}

func (opt *KeystoreOption) Provider(network string) (provider *Provider, err error) {
	if opt.Network[network] == nil {
		return nil, ErrNilNetwork
	}
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = opt.Network[network].GetClient(opt.Context)
	if err != nil {
		return
	}
	provider.Accounts, err = opt.getAccounts()
	return
}

func (opt *KeystoreOption) Export(path string) (err error) {
	data, err := json.MarshalIndent(opt, "", "    ")
	if err != nil {
		return
	}
	return utils.WriteFile(path, data)
}

func (opt *KeystoreOption) Import(ctx context.Context, path string) (err error) {
	opt.Context = ctx
	data, err := utils.ReadFile(path)
	if err != nil {
		return
	}
	return json.Unmarshal(data, opt)
}

func (opt *KeystoreOption) ImportProvider(ctx context.Context, path string, network string) (provider *Provider, err error) {
	if err = opt.Import(ctx, path); err != nil {
		return
	}
	return opt.Provider(network)
}
