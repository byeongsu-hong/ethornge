package main

import (
	"os"
	"github.com/urfave/cli"
	"log"
)

func main() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
