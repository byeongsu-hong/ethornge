package account

import (
	"math/big"

	"time"

	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/frostornge/ethornge/provider"
)

type Account struct {
	*bind.TransactOpts
}

func (acc Account) GetOpt(
	ctx context.Context,
	value *big.Int,
	gasPrice *big.Int,
	gasLimit uint64,
) *bind.TransactOpts {
	return &bind.TransactOpts{
		From:     acc.From,
		Signer:   acc.Signer,
		Value:    value,
		GasPrice: gasPrice, // Gas price to use for the transaction execution (nil = gas price oracle)
		GasLimit: gasLimit, // Gas limit to set for the transaction execution (0 = estimate)
		Context:  ctx,      // Network context to support cancellation and timeouts (nil = no timeout)
	}
}

func (acc Account) SendToWithTimeout(
	pv *provider.Provider,
	to common.Address,
	val *big.Int,
	t time.Duration,
) (*types.Transaction, *types.Receipt, error) {
	tx, err := pv.NewSignedTransaction(
		acc.GetOpt(pv.Context, val, nil, 0), to,
	)
	if err != nil {
		return nil, nil, err
	}
	if err = pv.SendTransaction(pv.Context, tx); err != nil {
		return nil, nil, err
	}
	receipt, err := pv.WaitMinedWithTimeout(tx, t)
	return tx, receipt, err
}

func (acc Account) SendTo(
	pv *provider.Provider,
	to common.Address,
	val *big.Int,
) (*types.Transaction, *types.Receipt, error) {
	return acc.SendToWithTimeout(pv, to, val, 30*time.Minute)
}
