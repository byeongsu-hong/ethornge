package main

import (
	"log"

	"github.com/frostornge/ethornge/gorange"
)

func main() {
	node, err := gorange.Launch(gorange.Config{
		4470,
		"0.0.0.0", 4471,
		"0.0.0.0", 4472,
		20,
		100,
	})
	if err != nil {
		log.Fatalln(err)
	}
	node.Wait()
}
