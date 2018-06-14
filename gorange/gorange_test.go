package gorange

import (
	"testing"
)

const (
	testGenesisPath = "../test/gorange/genesis.json"
	testAccountDir  = "../test/gorange/account/"
	testGethDatadir = "../test/gorange/datadir/"
)

func TestLaunch(t *testing.T) {
	//utils.RemoveDir("../test/gorange")
	//
	//var (
	//	err      error
	//	accounts = account.GetDefaultAccounts()
	//	gnOpt    = &GenesisOption{
	//		FilePath:  testGenesisPath,
	//		Consensus: true,
	//		Period:    5,
	//		ChainId:   big.NewInt(4349),
	//		Signers: []common.Address{
	//			accounts[0].Address,
	//			accounts[1].Address,
	//		},
	//		Accounts: accounts,
	//	}
	//	gethOpt = &GethOption{
	//		Geth:       "geth",
	//		DataDir:    testGethDatadir,
	//		AccountDir: testAccountDir,
	//	}
	//)
	//
	//if err = Launch(gnOpt, gethOpt); err != nil {
	//	t.Error(err)
	//	return
	//}
}
