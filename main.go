package main

import (
	"log"

	"github.com/frostornge/ethornge/gorange"
)

// main is simple command-line function that launches private network.
func main() {
	node, err := gorange.Launch(gorange.DefaultRemoteConfig(nil, 0))
	if err != nil {
		log.Fatalln(err)
	}
	node.Wait()
}
