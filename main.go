package main

import (
	"log"

	"github.com/frostornge/ethornge/gorange"
)

func main() {
	node, err := gorange.Launch(gorange.Config{
		NetworkId: 4470,
		HTTPPort:  4471,
		WSPort:    4472,
	})
	if err != nil {
		log.Fatalln(err)
	}
	node.Wait()
}
