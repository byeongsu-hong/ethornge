package ethornge

import (
	"fmt"
	"strconv"

	"math/big"

	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/usbwallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type TxOpts []*bind.TransactOpts

func getLedgerOpt(account accounts.Account, wallet accounts.Wallet) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: account.Address,
		Signer: func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return wallet.SignTx(account, tx, big.NewInt(5777))
		},
	}
}

func GetLedgerOpts(opt *ProviderOption) (opts TxOpts, err error) {
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

	for i := opt.Start; i < opt.End; i++ {
		var path accounts.DerivationPath
		if path, err = accounts.ParseDerivationPath(opt.Path + strconv.Itoa(i)); err != nil {
			return
		}

		var account accounts.Account
		if account, err = wallet.Derive(path, true); err != nil {
			return
		}

		opts = append(opts, getLedgerOpt(account, wallet))
	}
	return
}

func getPrivateKeyOpt(key string) (opt *bind.TransactOpts, err error) {
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

func GetPrivateKeyOpts(opt *ProviderOption) (opts TxOpts, err error) {
	for _, key := range opt.Keys {
		var txOpts *bind.TransactOpts
		if txOpts, err = getPrivateKeyOpt(key); err != nil {
			return
		}
		opts = append(opts, txOpts)
	}
	return
}
