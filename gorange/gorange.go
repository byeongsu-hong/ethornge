package gorange

import (
	"fmt"

	"runtime"

	"github.com/ethereum/go-ethereum/eth"
	"github.com/frostornge/ethornge/utils"
)

// Launch starts an embedded private network that'll be used 
// during gorange deployment session, with given config.
func Launch(config Config) (*Node, error) {
	node, err := config.getNode()
	if err != nil {
		return nil, err
	}
	if err = node.start(); err != nil {
		return nil, err
	}
	if err = node.preAlloc(config); err != nil {
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

	// use minimum gas price for convinience.
	e.TxPool().SetGasPrice(utils.Gwei(1))

	// start executing transactions
	// TODO: runtime.NumCPU 조절
	if err := e.StartMining(runtime.NumCPU()); err != nil {
		return fmt.Errorf("Failed to start mining: %v", err)
	}
	return nil
}
