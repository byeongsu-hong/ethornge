package provider

import (
	"context"
	"fmt"
	"math/big"
	"strconv"

	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/usbwallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/frostornge/ethornge/utils"
)

var (
	DefaultLedger = &Ledger{
		Path:  "m/44'/60'/0'/",
		Start: 0,
		End:   1,
	}
)

type Ledger struct {
	// Ledger provider Option
	Path  string `json:"path"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

type LedgerOption struct {
	Context context.Context
	Network map[string]*Network `json:"network"`
	Ledger  *Ledger             `json:"ledger"`
}

func getLedgerAccount(account accounts.Account, wallet accounts.Wallet, id *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: account.Address,
		Signer: func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return wallet.SignTx(account, tx, id)
		},
	}
}

func (opt *LedgerOption) getAccounts(network string) (opts []*bind.TransactOpts, err error) {
	var ledger = opt.Ledger
	var hub *usbwallet.Hub
	if hub, err = usbwallet.NewLedgerHub(); err != nil {
		return
	}

	if len(hub.Wallets()) == 0 {
		err = fmt.Errorf("No wallet found")
		return
	}

	var wallet = hub.Wallets()[0]
	if err = wallet.Open(""); err != nil {
		return
	}

	for i := ledger.Start; i < ledger.End; i++ {
		var path accounts.DerivationPath
		if path, err = accounts.ParseDerivationPath(ledger.Path + strconv.Itoa(i)); err != nil {
			return
		}

		var account accounts.Account
		if account, err = wallet.Derive(path, true); err != nil {
			return
		}

		var txOpt = getLedgerAccount(account, wallet, opt.Network[network].NetworkID)
		txOpt.Context = opt.Context
		opts = append(opts, txOpt)
	}
	return
}

func (opt *LedgerOption) Provider(network string) (provider *Provider, err error) {
	if opt.Network[network] == nil {
		return nil, ErrNilNetwork
	}
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = opt.Network[network].GetClient(opt.Context)
	if err != nil {
		return
	}
	provider.Accounts, err = opt.getAccounts(network)
	return
}

func (opt *LedgerOption) Export(path string) (err error) {
	data, err := json.MarshalIndent(opt, "", "    ")
	if err != nil {
		return
	}
	return utils.WriteFile(path, data)
}

func (opt *LedgerOption) Import(ctx context.Context, path string) (err error) {
	opt.Context = ctx
	data, err := utils.ReadFile(path)
	if err != nil {
		return
	}
	return json.Unmarshal(data, opt)
}

func (opt *LedgerOption) ImportProvider(ctx context.Context, path string, network string) (provider *Provider, err error) {
	if err = opt.Import(ctx, path); err != nil {
		return
	}
	return opt.Provider(network)
}
