package gorange

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/node"
	"github.com/frostornge/ethornge/provider"
)

const network = "gorange"

type Node struct {
	node.Node
}

func (n Node) Keystore() *keystore.KeyStore {
	return n.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
}

func (n Node) getProvider(ctx context.Context, url string, phrases []string) (pv *provider.Provider, err error) {
	net := make(map[string]*provider.Network)
	net[network] = &provider.Network{
		URL:     url,
		Network: network,
	}

	pv = new(provider.Provider)
	pv.Context = ctx
	pv.Client, err = net[network].GetClient(ctx)
	if err != nil {
		return
	}

	var e *eth.Ethereum
	if err = n.Service(&e); err != nil {
		err = fmt.Errorf("Ethereum service not running: %v", err)
		return
	}
	chainId := e.BlockChain().Config().ChainID

	ks := n.Keystore()
	accounts := ks.Accounts()
	for i := range accounts {
		if len(phrases) > i {
			if err = ks.Unlock(accounts[i], phrases[i]); err != nil {
				err = fmt.Errorf(
					"error occured while unlock account\n addr : %s, error : %v",
					accounts[i].Address.Hex(), err,
				)
				return
			}
			pv.Accounts = append(pv.Accounts, &bind.TransactOpts{
				Context: ctx,
				From:    accounts[i].Address,
				Signer: func(signer types.Signer, addresses common.Address, transaction *types.Transaction) (*types.Transaction, error) {
					return ks.SignTx(accounts[i], transaction, chainId)
				},
			})
		} else {
			break
		}
	}

	return
}

func (n Node) HttpProvider(ctx context.Context, phrases []string) (*provider.Provider, error) {
	return n.getProvider(ctx, "http://"+n.HTTPEndpoint(), phrases)
}

func (n Node) WsProvider(ctx context.Context, phrases []string) (*provider.Provider, error) {
	return n.getProvider(ctx, "ws://"+n.WSEndpoint(), phrases)
}
