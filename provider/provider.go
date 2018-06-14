package provider

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Provider struct {
	*ethclient.Client
	Context  context.Context
	Accounts []*bind.TransactOpts
}

func (pv *Provider) WaitMinedWithTimeout(tx *types.Transaction, t time.Duration) (*types.Receipt, error) {
	ctx, close := context.WithTimeout(pv.Context, t)
	defer close()
	return bind.WaitMined(ctx, pv.Client, tx)
}

func (pv *Provider) WaitDeployedWithTimeout(tx *types.Transaction, t time.Duration) (common.Address, error) {
	ctx, close := context.WithTimeout(pv.Context, t)
	defer close()
	return bind.WaitDeployed(ctx, pv.Client, tx)
}