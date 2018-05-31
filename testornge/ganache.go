package testornge

import (
	"bytes"
	"log"
	"os/exec"
	"time"
)

const URL = "http://localhost:8545"

func Launch() (cmd *exec.Cmd, keys []string, err error) {
	var buf bytes.Buffer
	cmd = exec.Command("ganache-cli")
	cmd.Stdout = &buf
	if err = cmd.Start(); err != nil {
		return
	}
	time.Sleep(5 * time.Second)
	keys = parse(buf.String())
	log.Println("Ganache Launched")
	return
}

func Kill(cmd *exec.Cmd) (err error) {
	if err = cmd.Process.Kill(); err != nil {
		return
	}
	log.Println("Ganache Killed")
	return
}
