package provider

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"

	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/frostornge/ethornge/utils"
)

type PrivateKey struct {
	// PrivateKey provider Option
	Keys []string `json:"keys"`
}

type PrivateKeyOption struct {
	Context    context.Context
	Network    map[string]*Network `json:"network"`
	PrivateKey *PrivateKey         `json:"private_key"`
}

func getPrivateKeyAccount(key string) (opt *bind.TransactOpts, err error) {
	var buf []byte
	if buf, err = hex.DecodeString(key); err != nil {
		return
	}

	var tmp *ecdsa.PrivateKey
	if tmp, err = crypto.ToECDSA(buf); err != nil {
		return
	}

	opt = bind.NewKeyedTransactor(tmp)
	return
}

func (opt *PrivateKeyOption) getAccounts() (opts []*bind.TransactOpts, err error) {
	var privateKey = opt.PrivateKey
	for _, key := range privateKey.Keys {
		var txOpt *bind.TransactOpts
		if txOpt, err = getPrivateKeyAccount(key); err != nil {
			return
		}

		txOpt.Context = opt.Context
		opts = append(opts, txOpt)
	}
	return
}

func (opt *PrivateKeyOption) Provider(network string) (provider *Provider, err error) {
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

func (opt *PrivateKeyOption) Export(path string) (err error) {
	data, err := json.MarshalIndent(opt, "", "    ")
	if err != nil {
		return
	}
	return utils.WriteFile(path, data)
}

func (opt *PrivateKeyOption) Import(ctx context.Context, path string) (err error) {
	opt.Context = ctx
	data, err := utils.ReadFile(path)
	if err != nil {
		return
	}
	return json.Unmarshal(data, opt)
}

func (opt *PrivateKeyOption) ImportProvider(ctx context.Context, path string, network string) (provider *Provider, err error) {
	if err = opt.Import(ctx, path); err != nil {
		return
	}
	return opt.Provider(network)
}
