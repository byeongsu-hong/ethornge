package ganache

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"os/exec"
	"strings"
	"time"

	"github.com/frostornge/ethornge/account"
	"github.com/frostornge/ethornge/provider"
)

const (
	attempt = 50
	message = "Listening on 127.0.0.1:"
)

func Launch(ganache string, opt *Option) (cmd *exec.Cmd, err error) {
	if opt.Accounts == nil {
		err = fmt.Errorf("Please insert account settings")
		return
	}
	var port = convert(opt.Port, "8545")
	if checkPortAlreadyUsed(port) {
		err = fmt.Errorf("Port already used")
		return
	}
	var stdout bytes.Buffer
	cmd = exec.Command("bash", "-c", opt.generate(ganache))
	cmd.Stdout = &stdout
	if err = cmd.Start(); err != nil {
		return
	}
	var cnt = 0
	for !strings.Contains(stdout.String(), message+fmt.Sprint(port)) {
		if cnt == attempt {
			err = fmt.Errorf("Launch timeout")
			return
		}
		time.Sleep(1 * time.Second)
		cnt++
	}
	return
}

func LaunchDefault(ganache string, port *big.Int, network string) (cmd *exec.Cmd, pv *provider.Provider, err error) {
	var opt = &Option{
		Port:     port,
		Accounts: account.GetDefaultAccounts(),
	}
	if cmd, err = Launch(ganache, opt); err != nil {
		return
	}
	if pv, err = opt.Provider(context.Background(), network); err != nil {
		return
	}
	return
}

func checkPortAlreadyUsed(port string) (isUsed bool) {
	var stdout bytes.Buffer
	var cmd = exec.Command("lsof", "-t", "-i", ":"+port)
	cmd.Stdout = &stdout
	cmd.Run()
	return len(stdout.String()) != 0
}
