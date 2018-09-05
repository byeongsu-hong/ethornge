package account

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

// Key is a pair of public key address and private key.
type Key struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

// Keys is an array of keys.
type Keys []*Key

// NewKey generates new account.
func NewKey() (*Key, error) {
	priv, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}
	return &Key{
		crypto.PubkeyToAddress(priv.PublicKey),
		priv,
	}, nil
}

// SignTx signs a given transaction with the keypair.
func (k Key) SignTx(transaction *types.Transaction, chainId *big.Int) (*types.Transaction, error) {
	return types.SignTx(transaction, types.NewEIP155Signer(chainId), k.PrivateKey)
}

// GetKeys returns an array of addresses.
func (ks Keys) GetAddresses() []common.Address {
	var addrs []common.Address
	for _, k := range ks {
		addrs = append(addrs, k.Address)
	}
	return addrs
}

// GetKeys returns an array of private key.
func (ks Keys) GetKeys() (keys []*ecdsa.PrivateKey) {
	for _, acc := range ks {
		keys = append(keys, acc.PrivateKey)
	}
	return
}

// GetDividedMetadata returns a tuple of addresses and private keys.
func (ks Keys) GetDividedMetadata() (addrs []common.Address, keys []*ecdsa.PrivateKey) {
	for _, acc := range ks {
		addrs = append(addrs, acc.Address)
		keys = append(keys, acc.PrivateKey)
	}
	return
}
