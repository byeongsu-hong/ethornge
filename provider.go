package ethornge

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"strconv"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/usbwallet"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func LedgerProvider(p string, index int) (account accounts.Account, wallet accounts.Wallet, err error) {
	var hub *usbwallet.Hub
	if hub, err = usbwallet.NewLedgerHub(); err != nil {
		return
	}

	if len(hub.Wallets()) == 0 {
		err = fmt.Errorf("No wallet found")
		return
	}

	wallet = hub.Wallets()[0]
	if err = wallet.Open(""); err != nil {
		return
	}

	var path accounts.DerivationPath
	if path, err = accounts.ParseDerivationPath(p + strconv.Itoa(index)); err != nil {
		return
	}

	if account, err = wallet.Derive(path, true); err != nil {
		return
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
