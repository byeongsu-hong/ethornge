package provider

import (
	"context"
	"os/exec"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/frostornge/ethornge/account"
	"github.com/frostornge/ethornge/ganache"
	"github.com/frostornge/ethornge/test/build/adapter"
)

var ctx = context.Background()

func TestPrivateKeyProvider(t *testing.T) {
	var (
		err error
		cmd *exec.Cmd
		opt = &ganache.Option{Accounts: account.GetDefaultAccounts()}
	)
	if cmd, err = ganache.Launch("ganache-cli", opt); err != nil {
		t.Error(err)
		return
	}
	defer cmd.Process.Kill()

	var pv *Provider
	if pv, err = PrivateKeyProvider(&Option{
		URL:     opt.Url(),
		Context: ctx,
		Keys:    opt.Accounts.GetKeys(),
	}); err != nil {
		t.Error("Error : ", err)
		return
	}
	defer pv.Close()

	var mint *adapter.MintableToken
	if _, _, mint, err = adapter.DeployMintableToken(pv.Accounts[0], pv.Client); err != nil {
		t.Error("Error : ", err)
		return
	}

	var owner common.Address
	if owner, err = mint.Owner(nil); err != nil {
		t.Error("Error : ", err)
		return
	}
	if owner != pv.Accounts[0].From {
		t.Error("Deployer and Contract owner mismatching")
	}
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
