package gorange

import (
	"context"
	"testing"

	"math/big"

	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/frostornge/ethornge/test/build/adt"
	"github.com/frostornge/ethornge/utils"
	"github.com/stretchr/testify/assert"
)

func TestLaunch(t *testing.T) {
	n, err := Launch(Config{2470, 2471, 2472})
	assert.NoError(t, err)

	pv, err := n.WsProvider(context.Background(), []string{""})
	assert.NoError(t, err)

	_, tx, revert, err := adapter.DeployRevertMessage(pv.Accounts[0], pv.Client)
	assert.NoError(t, err)
	bind.WaitDeployed(pv.Context, pv.Client, tx)

	pv.Accounts[0].GasPrice = utils.Gwei(1)
	pv.Accounts[0].GasLimit = 50000

	tx, err = revert.Set(pv.Accounts[0], big.NewInt(60))
	assert.NoError(t, err)

	receipt, err := bind.WaitMined(pv.Context, pv.Client, tx)
	assert.NoError(t, err)

	log.Println(receipt.Bloom)

	n.Stop()
}
