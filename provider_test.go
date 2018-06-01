package ethornge

import (
	"log"
	"testing"

	"context"

	"math/big"

	"./test/build/adapter"
	"./testornge"

	"os/exec"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var ctx = context.Background()

func TestLedgerProvider(t *testing.T) {
	var ganache *exec.Cmd
	var keys []string
	var err error
	if ganache, keys, err = testornge.Launch(8545, 10000); err != nil {
		t.Error("Error : ", err)
		return
	}
	defer testornge.Kill(ganache)

	var ganacheProvider *Provider
	if ganacheProvider, err = PrivateKeyProvider(&ProviderOption{
		URL:     testornge.URL,
		Keys:    keys,
		Context: context.Background(),
	}); err != nil {
		t.Error("Error : ", err)
		return
	}

	var ledgerProvider *Provider
	if ledgerProvider, err = LedgerProvider(&ProviderOption{
		URL:     testornge.URL,
		Context: context.Background(),
		Path:    DefaultLedgerOption.Path,
		Start:   0,
		End:     10,
	}); err != nil {
		t.Error("Error : ", err)
		return
	}

	var to = ledgerProvider.Accounts[0].From
	for _, opts := range ganacheProvider.Accounts {
		opts = &bind.TransactOpts{
			From:     opts.From,
			Signer:   opts.Signer,
			GasLimit: 21000,
			GasPrice: Gwei(15),
			Value:    Ether(10),
			Context:  ctx,
		}

		var tx *types.Transaction
		if tx, err = NewSignedTransaction(ganacheProvider.Client, opts, to); err != nil {
			t.Error("Error : ", err)
			return
		}

		if err = ganacheProvider.Client.SendTransaction(ctx, tx); err != nil {
			t.Error("Error : ", err)
			return
		}

		var balance *big.Int
		if balance, err = ganacheProvider.Client.BalanceAt(ctx, to, nil); err != nil {
			t.Error("Error : ", err)
			return
		}

		log.Println(WeiToEth(balance))
	}

	var token *adapter.MintableToken
	var addr common.Address
	var tx *types.Transaction
	if addr, tx, token, err = adapter.DeployMintableToken(ledgerProvider.Accounts[0], ledgerProvider.Client); err != nil {
		t.Error("Error : ", err)
		return
	}
	log.Println("Mintable Token :", addr.Hex())
	log.Println("TxHash :", tx.Hash().Hex())

	var owner common.Address
	if owner, err = token.Owner(nil); err != nil {
		t.Error("Error : ", err)
		return
	}
	log.Println(owner.Hex())
}
