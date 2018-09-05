package main

import (
	"context"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/frostornge/ethornge/account"
	"github.com/frostornge/ethornge/gorange"
	"github.com/frostornge/ethornge/provider"
	"github.com/frostornge/ethornge/test/build/bind"
	"github.com/frostornge/ethornge/utils"
)

// SendTx shows an example of deploying smart contract using embedded private network.
func DeploySmartContract() {
	k, _ := account.NewKey()
	n, _ := gorange.Launch(
		gorange.DefaultLocalConfig(
			[]common.Address{k.Address},
			100,
		),
	)
	defer n.Stop()

	p, _ := n.WsProviderWithAccounts(context.Background(), account.Keys{k})
	defer p.Close()

	_, tx, _, err := adapter.DeployTest(p.Accounts[0], p)
	if err != nil {
		log.Println("Failed to deploy test contract :", err)
		return
	}
	addr, err := p.WaitDeployedWithTimeout(tx, 5*time.Minute)
	if err != nil {
		log.Println("Timeout! :", err)
		return
	}

	receipt, _ := p.TransactionReceipt(p.Context, tx.Hash())
	log.Println("Contract address :", addr.Hex())
	provider.PrintTxResult(tx, receipt)
}

// SendTx shows an example of sending 10 ETH using embedded private network.
func SendTx() {
	k1, _ := account.NewKey()
	k2, _ := account.NewKey()
	n, _ := gorange.Launch(
		gorange.DefaultLocalConfig(
			[]common.Address{
				k1.Address,
				k2.Address,
			},
			100,
		),
	)
	defer n.Stop()

	p, _ := n.WsProviderWithAccounts(context.Background(), account.Keys{k1, k2})
	defer p.Close()

	acc := account.Account{TransactOpts: p.Accounts[0]}
	tx, receipt, err := acc.SendTo(p, k2.Address, utils.Ether(10))
	if err != nil {
		log.Println("Failed to send ether :", err)
		return
	}
	provider.PrintTxResult(tx, receipt)
}
