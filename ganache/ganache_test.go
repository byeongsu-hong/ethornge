package ganache

import (
	"os/exec"
	"testing"
)

func TestLaunch(t *testing.T) {
	var (
		err      error
		cmd      *exec.Cmd
		accounts = GetDefaultAccounts()
	)
	if cmd, err = Launch("ganache-cli", &Option{Accounts: accounts}); err != nil {
		t.Error(err)
		return
	}
	defer cmd.Process.Kill()
}
