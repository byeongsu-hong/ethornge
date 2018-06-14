package provider_test

import (
	"log"
	"math/big"
	"os"
	"testing"

	"github.com/frostornge/ethornge/account"
	"github.com/frostornge/ethornge/ganache"
	"github.com/frostornge/ethornge/utils"
)

var (
	TestPort      = big.NewInt(2432)
	TestAccounts  = account.GetDefaultAccounts()
	TestNetworkId = big.NewInt(2432)

	TestExportEnv     = "../test/env/export"
	TestImportEnv     = "../test/env/import"
	TestLedgerEnv     = "ledger.json"
	TestKeystoreEnv   = "keystore.json"
	TestPrivateKeyEnv = "privateKey.json"
)

func TestMain(m *testing.M) {
	log.SetFlags(log.Lshortfile)
	utils.RemoveDir(TestExportEnv)
	utils.CreateDir(TestExportEnv)

	opt := &ganache.Option{
		Accounts:  TestAccounts,
		Port:      TestPort,
		NetworkId: TestNetworkId,
	}
	if cmd, err := ganache.Launch("ganache-cli", opt); err != nil {
		log.Fatal(err)
	} else {
		defer os.Exit(m.Run())
		defer cmd.Process.Kill()
	}
}
