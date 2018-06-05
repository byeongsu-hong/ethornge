package provider

import (
	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getClient(opt *Option) (client *ethclient.Client, err error) {
	client, err = ethclient.Dial(opt.URL)
	if err != nil {
		return
	}

	var sync *ethereum.SyncProgress
	if sync, err = client.SyncProgress(opt.Context); err != nil {
		return
	}

	if sync != nil {
		err = fmt.Errorf("Current provider is syncing to blockchain")
	}
	return
}
