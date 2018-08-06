package gorange

import (
	"fmt"

	"github.com/ethereum/go-ethereum/eth"
	"github.com/frostornge/ethornge/utils"
)

func Launch(c Config) (*Node, error) {
	node, err := getNode(c)
	if err != nil {
		return nil, err
	}
	err = start(node)
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
