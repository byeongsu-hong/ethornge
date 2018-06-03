package gorange

import (
	"math/big"
	"testing"

	"../account"

	"github.com/ethereum/go-ethereum/common"
)

func TestLaunch(t *testing.T) {
	var err error
	var accounts = account.GetDefaultAccounts()
	if err = Launch(&GenesisOption{
		Period:  5,
		ChainId: big.NewInt(4349),
		Signers: []common.Address{
			accounts[0].Address,
			accounts[1].Address,
		},
		Accounts: accounts,
	}); err != nil {
		t.Error(err)
		return
	}
}
