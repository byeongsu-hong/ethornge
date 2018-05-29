package ethornge

import (
	"log"
	"testing"

	"context"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	URL = "http://localhost:8545" // Ganache only Plz
)

var (
	privKeys = []string{
		"e8dd70be4f88ad68e2d833bbefa8b85cd39610857fdbc7b08630c32be4540d50",
		"58bd6d7b4d0d90a4bf8798bb645b7e0c66f2f39808a1c6259c7be62b9d9c4eb0",
		"549dbcddf98498822043bfc3ae4b4b8b1edd5da7e9b5c1d464bcd86551b2d8a8",
		"64f58fc583408f22c10c74eac0d87de9b8ace8a3b1f59ad30464d19f554b12d2",
		"d253029f169a71575b276a95982205034580ff172c01c8f3613c1ff77d1bf2f2",
		"68ea0da277ddfecc886a5a73d3934eb704e1e803841f979a26b3366451bcafc6",
		"89960bf9569f28c1c26dbbd6e24ed809777fb58c0caf0463196e0a9aac7e2581",
		"ca0728e8a7945774d44f0a765cfd399af2be08a89144c76a465e47c1083c1bd9",
		"c7338ebea5734ff6f825247313736dc37b20319b9acf6aa9eed172bbb0774f90",
		"5f5885cf467b64879a954c15553f3ce38e9922b4e554d5d051aa429d9adcda23",
	}
)

var ctx = context.Background()

func TestLedgerProvider(t *testing.T) {
	var err error
	var client *ethclient.Client
	if client, err = ethclient.Dial(URL); err != nil {
		t.Error("Error : ", err)
		return
	}

	var ganacheProviders = make(map[common.Address]*bind.TransactOpts)
	if ganacheProviders, err = PrivKeyProvider(privKeys); err != nil {
		t.Error("Error : ", err)
		return
	}

	var to = common.HexToAddress("0x93486fb06d2c18f3802d2c6c7628034067030e9a")
	for _, opts := range ganacheProviders {
		opts = &bind.TransactOpts{
			From:     opts.From,
			Signer:   opts.Signer,
			GasLimit: 21000,
			GasPrice: Gwei(15),
			Value:    Ether(10),
			Context:  ctx,
		}

		var tx *types.Transaction
		if tx, err = NewSignedTransaction(client, opts, to); err != nil {
			t.Error("Error : ", err)
			return
		}

		if err = client.SendTransaction(ctx, tx); err != nil {
			t.Error("Error : ", err)
			return
		}

		var balance *big.Int
		if balance, err = client.BalanceAt(ctx, to, nil); err != nil {
			t.Error("Error : ", err)
			return
		}

		log.Println(WeiToEth(balance))
	}

	//var account accounts.Account
	//var wallet accounts.Wallet
	//if account, wallet, err = LedgerProvider("m/44'/60'/0'/", 1); err != nil {
	//	t.Error("Error : ", err)
	//	return
	//}

	//var a = GetLedgerTransactor(account, wallet)
	//
	//var addr common.Address
	//var tx *types.Transaction
	//var token *adapter.MintableToken
	//if addr, tx, token, err = adapter.DeployMintableToken(a, client); err != nil {
	//	t.Error("Error : ", err)
	//	return
	//}
	//
	//var owner common.Address
	//if owner, err = token.Owner(nil); err != nil {
	//	t.Error("Error : ", err)
	//	return
	//}
	//log.Println(owner.Hex())
}
