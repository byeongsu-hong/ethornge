package ethornge

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"

	"io"
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

func GetTransactorByPrivKey(key string) (txOpts *bind.TransactOpts, err error) {
	var buf []byte
	if buf, err = hex.DecodeString(key); err != nil {
		return nil, err
	}

	var tmp *ecdsa.PrivateKey
	if tmp, err = crypto.ToECDSA(buf); err != nil {
		return nil, err
	}

	txOpts = bind.NewKeyedTransactor(tmp)
	return
}

func GetTransactorByKeystore(path string, pass string) (txOpts *bind.TransactOpts, err error) {
	var r io.Reader
	if r, err = os.Open(path); err != nil {
		return nil, err
	}

	return bind.NewTransactor(r, pass)
}

func GetLedgerTransactor(account accounts.Account, wallet accounts.Wallet) (txOpts *bind.TransactOpts) {
	return &bind.TransactOpts{
		From: account.Address,
		Signer: func(signer types.Signer, addr common.Address, tx *types.Transaction) (*types.Transaction, error) {
			return wallet.SignTx(account, tx, big.NewInt(5777))
		},
	}
}
