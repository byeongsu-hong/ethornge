package testornge

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"time"
)

const URL = "http://localhost:8545"

func checkPortAlreadyUsed(port int) bool {
	var stdout bytes.Buffer
	var cmd = exec.Command("lsof", "-t", "-i", ":"+fmt.Sprint(port))
	cmd.Stdout = &stdout
	cmd.Run()
	return len(stdout.String()) != 0
}

func Launch(port int, ether int) (cmd *exec.Cmd, keys []string, err error) {
	if checkPortAlreadyUsed(port) {
		err = fmt.Errorf("Port already in use : %d", port)
		return
	}

	var buf bytes.Buffer
	cmd = exec.Command("ganache-cli",
		"-p", fmt.Sprint(port),
		"-e", fmt.Sprint(ether),
	)
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
