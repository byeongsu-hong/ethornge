package provider

import (
	"context"
	"log"
	"os/exec"
	"testing"

	"../ganache"

	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/frostornge/ethornge/test/build/adapter"
)

var ctx = context.Background()

func TestPrivateKeyProvider(t *testing.T) {
	var (
		err      error
		cmd      *exec.Cmd
		accounts = ganache.GetDefaultAccounts()
	)
	if cmd, err = ganache.Launch("ganache-cli", &ganache.Option{
		Port:     big.NewInt(5455),
		Accounts: accounts,
	}); err != nil {
		t.Error(err)
		return
	}
	defer cmd.Process.Kill()

	var pv *Provider
	if pv, err = PrivateKeyProvider(&Option{
		URL:     ganache.DefaultURL,
		Context: ctx,
		Keys:    accounts.GetKeys(),
	}); err != nil {
		t.Error("Error : ", err)
		return
	}
	defer pv.Close()

	var token *adapter.MintableToken
	var addr common.Address
	var tx *types.Transaction
	if addr, tx, token, err = adapter.DeployMintableToken(pv.Accounts[0], pv); err != nil {
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

//func TestLedgerProvider(t *testing.T) {
//	var (
//		cmd  *exec.Cmd
//		keys []string
//		err  error
//	)
//	if cmd, keys, err = testornge.Launch(8545, 10000, 2470); err != nil {
//		t.Error("Error : ", err)
//		return
//	}
//	defer testornge.Kill(cmd)
//
//	var ganache *Provider
//	if ganache, err = PrivateKeyProvider(&Option{
//		URL:     testornge.URL,
//		Keys:    keys,
//		Context: context.Background(),
//	}); err != nil {
//		t.Error("Error : ", err)
//		return
//	}
//
//	var provider *Provider
//	if provider, err = LedgerProvider(&Option{
//		URL:     testornge.URL,
//		Context: context.Background(),
//		Path:    DefaultLedgerOption.Path,
//		Start:   0,
//		End:     10,
//	}); err != nil {
//		t.Error("Error : ", err)
//		return
//	}
//
//	var tx *types.Transaction
//	var owner = provider.Accounts[0]
//	for _, account := range ganache.Accounts {
//		if tx, err = NewSignedTransaction(ganache.Client, &bind.TransactOpts{
//			From:    account.From,
//			Signer:  account.Signer,
//			Value:   utils.Ether(3),
//			Context: account.Context,
//		}, owner.From); err != nil {
//			t.Error("Error : ", err)
//			return
//		}
//
//		if err = ganache.Client.SendTransaction(ctx, tx); err != nil {
//			t.Error("Error : ", err)
//			return
//		}
//	}
//
//	var token *adapter.MintableToken
//	var addr common.Address
//	if addr, tx, token, err = adapter.DeployMintableToken(provider.Accounts[0], provider.Client); err != nil {
//		t.Error("Error : ", err)
//		return
//	}
//	log.Println("Mintable Token :", addr.Hex())
//	log.Println("TxHash :", tx.Hash().Hex())
//
//	if addr, err = token.Owner(nil); err != nil {
//		t.Error("Error : ", err)
//		return
//	}
//	log.Println(addr.Hex())
//}
