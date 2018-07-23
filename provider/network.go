package provider

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrNilNetwork = fmt.Errorf("")
	ErrOnSync     = fmt.Errorf("Sync on progress")
)

type Network struct {
	URL       string   `json:"url"`
	Network   string   `json:"network"`
	NetworkID *big.Int `json:"network_id"`
}

func (opt *Network) GetClient(ctx context.Context) (client *ethclient.Client, err error) {
	client, err = ethclient.Dial(opt.URL)
	if err != nil {
		return
	}

	var sync *ethereum.SyncProgress
	if sync, err = client.SyncProgress(ctx); err != nil {
		return
	}

	if sync != nil {
		err = ErrOnSync
	}
	return
}
