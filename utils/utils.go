package utils

import (
	"math/big"

	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
)

func Gwei(x int64) *big.Int {
	return new(big.Int).Mul(
		big.NewInt(x),
		big.NewInt(params.GWei),
	)
}

func Ether(x int64) *big.Int {
	return new(big.Int).Mul(
		big.NewInt(x),
		big.NewInt(params.Ether),
	)
}

func WeiToGwei(wei *big.Int) *big.Float {
	return new(big.Float).Quo(
		new(big.Float).SetInt(wei),
		new(big.Float).SetInt64(params.GWei),
	)
}

func WeiToEth(wei *big.Int) *big.Float {
	return new(big.Float).Quo(
		new(big.Float).SetInt(wei),
		new(big.Float).SetInt64(params.Ether),
	)
}

func KeyToAddr(key string) (addr common.Address, err error) {
	var buf []byte
	if buf, err = hex.DecodeString(key); err != nil {
		return
	}

	var tmp *ecdsa.PrivateKey
	if tmp, err = crypto.ToECDSA(buf); err != nil {
		return
	}
	addr = crypto.PubkeyToAddress(tmp.PublicKey)
	return
}
