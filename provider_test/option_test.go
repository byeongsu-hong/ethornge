package provider_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/frostornge/ethornge/provider"
	"github.com/frostornge/ethornge/test/build/adt"
	"github.com/frostornge/ethornge/utils"
	"github.com/stretchr/testify/assert"
)

func TestLedgerOption_Import(t *testing.T) {
	lopt := provider.LedgerOption{}
	popt := provider.PrivateKeyOption{}
	assert.NoError(t, lopt.Import(context.Background(), filepath.Join(TestImportEnv, TestLedgerEnv)))
	assert.NoError(t, popt.Import(context.Background(), filepath.Join(TestImportEnv, TestPrivateKeyEnv)))

	lpv, err := lopt.Provider("test")
	assert.NoError(t, err)
	ppv, err := popt.Provider("test")
	assert.NoError(t, err)

	var a = ppv.Accounts[0]
	a.Value = utils.Ether(5)
	a.GasPrice = utils.Gwei(5)
	a.GasLimit = 50000

	tx, err := ppv.NewSignedTransaction(a, lpv.Accounts[0].From)
	assert.NoError(t, err)
	assert.NoError(t, ppv.SendTransaction(ppv.Context, tx))

	_, _, _, err = adapter.DeployMintableToken(lpv.Accounts[0], lpv.Client)
	assert.NoError(t, err)
}

func TestLedgerOption_Export(t *testing.T) {
	opt := provider.LedgerOption{}
	assert.NoError(t, opt.Import(context.Background(), filepath.Join(TestImportEnv, TestLedgerEnv)))
	assert.NoError(t, opt.Export(filepath.Join(TestExportEnv, TestLedgerEnv)))
}

func TestKeystoreOption_Import(t *testing.T) {
	kopt := provider.KeystoreOption{}
	popt := provider.PrivateKeyOption{}
	assert.NoError(t, kopt.Import(context.Background(), filepath.Join(TestImportEnv, TestKeystoreEnv)))
	assert.NoError(t, popt.Import(context.Background(), filepath.Join(TestImportEnv, TestPrivateKeyEnv)))

	kpv, err := kopt.Provider("test")
	assert.NoError(t, err)
	ppv, err := popt.Provider("test")
	assert.NoError(t, err)

	var a = ppv.Accounts[0]
	a.Value = utils.Ether(5)
	a.GasPrice = utils.Gwei(5)
	a.GasLimit = 50000

	tx, err := ppv.NewSignedTransaction(a, kpv.Accounts[0].From)
	assert.NoError(t, err)
	assert.NoError(t, ppv.SendTransaction(ppv.Context, tx))

	_, _, _, err = adapter.DeployMintableToken(kpv.Accounts[0], kpv.Client)
	assert.NoError(t, err)
}

func TestKeystoreOption_Export(t *testing.T) {
	opt := provider.KeystoreOption{}
	assert.NoError(t, opt.Import(context.Background(), filepath.Join(TestImportEnv, TestKeystoreEnv)))
	assert.NoError(t, opt.Export(filepath.Join(TestExportEnv, TestKeystoreEnv)))
}

func TestPrivateKeyOption_Import(t *testing.T) {
	opt := provider.PrivateKeyOption{}
	assert.NoError(t, opt.Import(context.Background(), filepath.Join(TestImportEnv, TestPrivateKeyEnv)))

	pv, err := opt.Provider("test")
	assert.NoError(t, err)

	_, _, _, err = adapter.DeployMintableToken(pv.Accounts[0], pv.Client)
	assert.NoError(t, err)
}

func TestPrivateKeyOption_Export(t *testing.T) {
	opt := provider.PrivateKeyOption{}
	assert.NoError(t, opt.Import(context.Background(), filepath.Join(TestImportEnv, TestPrivateKeyEnv)))
	assert.NoError(t, opt.Export(filepath.Join(TestExportEnv, TestPrivateKeyEnv)))
}
