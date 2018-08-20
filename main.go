package main

import (
	"log"

	"github.com/frostornge/ethornge/gorange"
)

func main() {
	node, err := gorange.Launch(gorange.DefaultRemoteConfig(nil, 0))
	if err != nil {
		log.Fatalln(err)
	}
	node.Wait()
}
