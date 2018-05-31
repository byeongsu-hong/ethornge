package ethornge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

func Gwei(x int64) *big.Int {
	return new(big.Int).Mul(
		big.NewInt(x),
		big.NewInt(params.Shannon),
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
		new(big.Float).SetInt64(params.Shannon),
	)
}

func WeiToEth(wei *big.Int) *big.Float {
	return new(big.Float).Quo(
		new(big.Float).SetInt(wei),
		new(big.Float).SetInt64(params.Ether),
	)
}
