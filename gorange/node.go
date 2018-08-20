package gorange

import (
	"context"
	"fmt"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/node"
	"github.com/frostornge/ethornge/account"
	"github.com/frostornge/ethornge/provider"
)

const network = "gorange"

type Node struct {
	node.Node
}

func (n Node) Keystore() *keystore.KeyStore {
	return n.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
}

func (n Node) appendPrivateKeyAccount(pv *provider.Provider, chainId *big.Int, accs account.Keys) (err error) {
	for _, acc := range accs {
		pv.Accounts = append(pv.Accounts, &bind.TransactOpts{
			Context: pv.Context,
			From:    acc.Address,
			Signer: func(signer types.Signer, addresses common.Address, transaction *types.Transaction) (*types.Transaction, error) {
				return acc.SignTx(transaction, chainId)
			},
		})
	}
	return
}

func (n Node) appendKeystoreAccount(pv *provider.Provider, chainId *big.Int) (err error) {
	ks := n.Keystore()
	accounts := ks.Accounts()
	for i := range accounts {
		pv.Accounts = append(pv.Accounts, &bind.TransactOpts{
			Context: pv.Context,
			From:    accounts[i].Address,
			Signer: func(signer types.Signer, addresses common.Address, transaction *types.Transaction) (*types.Transaction, error) {
				return ks.SignTxWithPassphrase(accounts[i], "", transaction, chainId)
			},
		})
	}
	return
}

func (n Node) getProvider(ctx context.Context, url string, accs account.Keys) (pv *provider.Provider, err error) {
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

	err = n.appendKeystoreAccount(pv, chainId)
	if err != nil {
		return
	}
	err = n.appendPrivateKeyAccount(pv, chainId, accs)
	if err != nil {
		return
	}

	return
}

func (n Node) HttpProvider(ctx context.Context) (*provider.Provider, error) {
	return n.getProvider(ctx, "http://"+n.HTTPEndpoint(), nil)
}

func (n Node) HttpProviderWithAccounts(ctx context.Context, accs account.Keys) (*provider.Provider, error) {
	return n.getProvider(ctx, "http://"+n.HTTPEndpoint(), accs)
}

func (n Node) WsProvider(ctx context.Context) (*provider.Provider, error) {
	return n.getProvider(ctx, "ws://"+n.WSEndpoint(), nil)
}

func (n Node) WsProviderWithAccounts(ctx context.Context, accs account.Keys) (*provider.Provider, error) {
	return n.getProvider(ctx, "ws://"+n.WSEndpoint(), accs)
}
