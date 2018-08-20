package account

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Key struct {
	Address    common.Address
	PrivateKey *ecdsa.PrivateKey
}

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

func (k Key) SignTx(transaction *types.Transaction, chainId *big.Int) (*types.Transaction, error) {
	return types.SignTx(transaction, types.NewEIP155Signer(chainId), k.PrivateKey)
}

type Keys []*Key

func (ks Keys) GetAddresses() []common.Address {
	var addrs []common.Address
	for _, k := range ks {
		addrs = append(addrs, k.Address)
	}
	return addrs
}
