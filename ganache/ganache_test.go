package ganache

import (
	"context"
	"math/big"
	"os/exec"
	"testing"

	"../account"
	"../provider"
)

func TestLaunch(t *testing.T) {
	var (
		err error
		cmd *exec.Cmd
		opt = &Option{
			Accounts: account.GetDefaultAccounts(),
			Port:     big.NewInt(8546),
		}
	)
	if cmd, err = Launch("ganache-cli", opt); err != nil {
		t.Error(err)
		return
	}
	defer cmd.Process.Kill()

	if _, err = provider.PrivateKeyProvider(&provider.Option{
		URL:     opt.Url(),
		Context: context.Background(),
		Keys:    opt.Accounts.GetKeys(),
	}); err != nil {
		t.Error(err)
		return
	}
}
