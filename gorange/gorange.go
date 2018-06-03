package gorange

import "github.com/ethereum/go-ethereum/core"

func Launch(opt *Option) (err error) {
	var genesis *core.Genesis
	if genesis, err = GetCliqueGenesis(opt); err != nil {
		return
	}
}
