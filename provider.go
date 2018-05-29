package ethornge

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"strconv"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/usbwallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func getLedgerProvider(account accounts.Account, wallet accounts.Wallet) (txOpts *bind.TransactOpts) {
	return &bind.TransactOpts{
		From: account.Address,
		Signer: func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return wallet.SignTx(account, tx, big.NewInt(5777))
		},
	}
}

func LedgerProvider(p string, startIndex int, endIndex int) (providers map[common.Address]*bind.TransactOpts, err error) {
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

	providers = make(map[common.Address]*bind.TransactOpts)
	for i := startIndex; i < endIndex; i++ {
		var path accounts.DerivationPath
		if path, err = accounts.ParseDerivationPath(p + strconv.Itoa(i)); err != nil {
			return
		}

		var account accounts.Account
		if account, err = wallet.Derive(path, true); err != nil {
			return
		}

		providers[account.Address] = getLedgerProvider(account, wallet)
	}
	return
}

func getKeyProvider(key string) (txOpts *bind.TransactOpts, err error) {
	var buf []byte
	if buf, err = hex.DecodeString(key); err != nil {
		return
	}

	var tmp *ecdsa.PrivateKey
	if tmp, err = crypto.ToECDSA(buf); err != nil {
		return
	}

	txOpts = bind.NewKeyedTransactor(tmp)
	return
}

func PrivKeyProvider(keys []string) (providers map[common.Address]*bind.TransactOpts, err error) {
	providers = make(map[common.Address]*bind.TransactOpts)
	for _, key := range keys {
		var txOpts *bind.TransactOpts
		if txOpts, err = getKeyProvider(key); err != nil {
			return
		}
		providers[txOpts.From] = txOpts
	}
	return
}
