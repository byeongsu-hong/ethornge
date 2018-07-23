package provider

import (
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Provider struct {
	*ethclient.Client
	Context  context.Context
	Accounts []*bind.TransactOpts
}

func (pv *Provider) ProcessTransaction(
	tx *types.Transaction,
	terr error,
) (receipt *types.Receipt, loop bool, err error) {
	if terr != nil {
		if strings.Contains(terr.Error(), "nonce too low") ||
			strings.Contains(terr.Error(), "known transaction") ||
			strings.Contains(terr.Error(), "replacement transaction underpriced") {
			return nil, true, terr
		} else {
			return nil, false, terr
		}
	}
	log.Println(tx.Hash().Hex())
	receipt, err = bind.WaitMined(pv.Context, pv.Client, tx)
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, true, terr
		}
		return nil, false, terr
	}
	return receipt, false, nil
}

func (pv *Provider) ProcessTransactionWithTimeout(
	tx *types.Transaction,
	timeout time.Duration,
	terr error,
) (receipt *types.Receipt, loop bool, err error) {
	if terr != nil {
		if strings.Contains(terr.Error(), "nonce too low") ||
			strings.Contains(terr.Error(), "known transaction") ||
			strings.Contains(terr.Error(), "replacement transaction underpriced") {
			return nil, true, terr
		} else {
			return nil, false, terr
		}
	}
	log.Println(tx.Hash().Hex())
	receipt, err = pv.WaitMinedWithTimeout(tx, timeout)
	if err != nil {
		if err == context.DeadlineExceeded {
			return nil, true, terr
		}
		return nil, false, terr
	}
	return receipt, false, nil
}

func (pv *Provider) GetSuggestedGasPrice(
	opt *bind.TransactOpts,
	crit *big.Int,
	set *big.Int,
) (err error) {
	for {
		var price *big.Int
		if price, err = pv.SuggestGasPrice(pv.Context); err != nil {
			return
		}

		if price.Cmp(crit) <= 0 {
			opt.GasPrice = set
			break
		}

		time.Sleep(1 * time.Second)
	}
	return
}

func (pv *Provider) WaitMinedWithTimeout(tx *types.Transaction, t time.Duration) (*types.Receipt, error) {
	ctx, c := context.WithTimeout(pv.Context, t)
	defer c()
	return bind.WaitMined(ctx, pv.Client, tx)
}

func (pv *Provider) WaitDeployedWithTimeout(tx *types.Transaction, t time.Duration) (common.Address, error) {
	ctx, c := context.WithTimeout(pv.Context, t)
	defer c()
	return bind.WaitDeployed(ctx, pv.Client, tx)
}
