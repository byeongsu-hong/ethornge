package gorange

import (
	"context"
	"fmt"

	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/frostornge/ethornge/utils"
)

func Launch(c Config) (*Node, error) {
	node, err := getNode(c)
	if err != nil {
		return nil, err
	}
	err = start(node)

	client, err := ethclient.Dial("ws://" + node.WSEndpoint())
	if err != nil {
		return nil, err
	}

	keystore := node.Keystore()
	owner := keystore.Accounts()[0]
	if err = keystore.Unlock(owner, ""); err != nil {
		return nil, err
	}

	for i := 0; i < c.Accounts; i++ {
		account, err := keystore.NewAccount("")
		if err != nil {
			return nil, err
		}

		nonce, err := client.NonceAt(context.Background(), owner.Address, nil)
		if err != nil {
			return nil, err
		}

		tx := types.NewTransaction(
			nonce,
			account.Address,
			utils.Ether(c.Balances),
			21000,
			utils.Gwei(1),
			nil,
		)

		signTx, err := keystore.SignTx(owner, tx, nil)
		if err != nil {
			return nil, err
		}

		err = client.SendTransaction(context.Background(), signTx)
		if err != nil {
			return nil, err
		}

		_, err = bind.WaitMined(context.Background(), client, signTx)
		if err != nil {
			return nil, err
		}

		log.Println(account.Address.Hex(), ":", c.Balances, "ETH")
	}
	return node, nil
}

func start(stack *Node) error {
	if err := stack.Start(); err != nil {
		return fmt.Errorf("Error starting protocol stack: %v", err)
	}

	var e *eth.Ethereum
	if err := stack.Service(&e); err != nil {
		return fmt.Errorf("Ethereum service not running: %v", err)
	}

	e.TxPool().SetGasPrice(utils.Gwei(1))
	if err := e.StartMining(true); err != nil {
		return fmt.Errorf("Failed to start mining: %v", err)
	}
	return nil
}
