package gorange

import (
	"log"

	"github.com/ethereum/go-ethereum/core"
)

func Launch(opt *GenesisOption) (err error) {
	var genesis *core.Genesis
	if genesis, err = GetCliqueGenesis(opt); err != nil {
		return
	}

	log.Println(genesis.MarshalJSON())
	return
}
