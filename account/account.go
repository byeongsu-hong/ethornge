package account

import (
	"fmt"
	"math/big"
	"path/filepath"

	"github.com/ethereum/go-ethereum/common"
	"github.com/frostornge/ethornge/utils"
)

var (
	DefaultBalance     = utils.Ether(1000)
	DefaultPrivateKeys = []string{
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

type Account struct {
	PrivateKey string
	PassPhrase string
	Address    common.Address
	Balance    *big.Int
}

func (account *Account) Export(key string, pass string) (err error) {
	if err = utils.WriteFile(key, []byte(account.PrivateKey)); err != nil {
		return
	}
	return utils.WriteFile(pass, []byte(account.PassPhrase))
}

func NewAccount(key string, pass string, balance *big.Int) (*Account, error) {
	var addr, err = utils.KeyToAddr(key)
	return &Account{
		PrivateKey: key,
		PassPhrase: pass,
		Address:    addr,
		Balance:    balance,
	}, err
}

type Accounts []*Account

func (accounts Accounts) Export(path string) (err error) {
	if err = utils.CreateDir(path); err != nil {
		return
	}
	for index, account := range accounts {
		if err = account.Export(
			filepath.Join(path, fmt.Sprint(index+1)+"pv.txt"),
			filepath.Join(path, fmt.Sprint(index+1)+"pw.txt"),
		); err != nil {
			return
		}
	}
	return
}

func (accounts Accounts) GetKeys() (keys []string) {
	for _, account := range accounts {
		keys = append(keys, account.PrivateKey)
	}
	return
}

func (accounts Accounts) GetAddrs() (addrs []common.Address) {
	for _, account := range accounts {
		addrs = append(addrs, account.Address)
	}
	return
}

func GetDefaultAccounts() (accounts Accounts) {
	for _, key := range DefaultPrivateKeys {
		var account, _ = NewAccount(key, "foobar", DefaultBalance)
		accounts = append(accounts, account)
	}
	return
}
