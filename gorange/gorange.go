package gorange

import (
	"fmt"

	"runtime"

	"github.com/ethereum/go-ethereum/eth"
	"github.com/frostornge/ethornge/utils"
)

func Launch(c Config) (*Node, error) {
	node, err := c.getNode()
	if err != nil {
		return nil, err
	}
	if err = node.start(); err != nil {
		return nil, err
	}
	if err = node.preAlloc(c); err != nil {
		return nil, err
	}

	return node, nil
}

func (n *Node) start() error {
	if err := n.Start(); err != nil {
		return fmt.Errorf("Error starting protocol stack: %v", err)
	}

	var e *eth.Ethereum
	if err := n.Service(&e); err != nil {
		return fmt.Errorf("Ethereum service not running: %v", err)
	}

	e.ChainDb()

	e.TxPool().SetGasPrice(utils.Gwei(1))
	if err := e.StartMining(runtime.NumCPU()); err != nil {
		return fmt.Errorf("Failed to start mining: %v", err)
	}
	return nil
}
