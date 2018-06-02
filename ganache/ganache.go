package ganache

import (
	"os/exec"

	"fmt"

	"log"
	"strings"
	"time"

	"bytes"

	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	attempt = 5
	message = "refused"
)

const DefaultURL = "http://localhost:8545"

func Launch(ganache string, opt *Option) (cmd *exec.Cmd, err error) {
	if opt.Accounts == nil {
		err = fmt.Errorf("Please insert account settings")
		return
	}
	if checkPortAlreadyUsed(Convert(opt.Port, "8545")) {
		err = fmt.Errorf("Port already used")
		return
	}
	cmd = exec.Command("bash", "-c", opt.generate(ganache))
	if err = cmd.Start(); err != nil {
		return
	}
	err = waitLaunched("http://localhost:" + Convert(opt.Port, "8545"))
	return
}

func checkPortAlreadyUsed(port string) (isUsed bool) {
	var stdout bytes.Buffer
	var cmd = exec.Command("lsof", "-t", "-i", ":"+port)
	cmd.Stdout = &stdout
	cmd.Run()
	return len(stdout.String()) != 0
}

func waitLaunched(rawurl string) (err error) {
	var count = 0
CONNECT:
	time.Sleep(300 * time.Millisecond)
	if _, err = ethclient.Dial(rawurl); err != nil {
		log.Println("1")
		if strings.Contains(err.Error(), message) {
			log.Println("Attempt", count+1)
			if count == attempt {
				log.Println("Timeout")
				return
			}
			count++
			goto CONNECT
		} else {
			return
		}
	}
	return
}
