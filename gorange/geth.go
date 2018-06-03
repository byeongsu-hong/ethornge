package gorange

import (
	"os/exec"

	"log"

	"fmt"
	"path/filepath"
	"sync"

	"../account"
)

type WSConfig struct {
	Addr    string
	Port    string
	API     string
	Origins string
}

type RPCConfig struct {
	Addr string
	Port string
	API  string
}

type GethOption struct {
	Geth       string
	DataDir    string
	AccountDir string
	Keystore   string
	WS         *WSConfig
	RPC        *RPCConfig
}

func (gethOpt *GethOption) geth() {
	gethOpt.command("--dev")
}

func (gethOpt *GethOption) importAccounts(accounts account.Accounts) (err error) {
	if err = accounts.Export(gethOpt.AccountDir); err != nil {
		return
	}

	var ws = new(sync.WaitGroup)
	for n := range accounts {
		ws.Add(1)
		go func(n int) {
			defer ws.Done()
			gethOpt.command(
				"account", "import", "--password",
				filepath.Join(gethOpt.AccountDir, fmt.Sprint(n)+"pw.txt"),
				filepath.Join(gethOpt.AccountDir, fmt.Sprint(n)+"pv.txt"),
			)
		}(n)
	}
	ws.Wait()
	return
}

func (gethOpt *GethOption) initGenesis(filePath string) {
	gethOpt.command("init", filePath)
}

func (gethOpt *GethOption) command(args ...string) {
	args = append([]string{"--datadir", gethOpt.DataDir}, args...)
	var out, _ = exec.Command(gethOpt.Geth, args...).Output()
	log.Println(string(out))
	return
}
