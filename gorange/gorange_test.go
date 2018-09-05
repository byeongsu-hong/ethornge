package gorange

import (
	"context"
	"testing"

	"log"

	"github.com/frostornge/ethornge/account"
	"github.com/frostornge/ethornge/utils"
	"github.com/stretchr/testify/assert"
)

func TestLaunch(t *testing.T) {
	var ks account.Keys
	for i := 0; i < 100; i++ {
		k, err := account.NewKey()
		assert.NoError(t, err)
		ks = append(ks, k)
	}

	n, err := Launch(DefaultLocalConfig(ks.GetAddresses(), 100))
	assert.NoError(t, err)

	pv, err := n.WsProvider(context.Background())
	assert.NoError(t, err)
	defer pv.Close()

	for _, acc := range ks.GetAddresses() {
		balance, _ := pv.BalanceAt(pv.Context, acc, nil)
		log.Println(utils.WeiToEth(balance))
	}
	n.Stop()
}
