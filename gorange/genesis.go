package gorange

import (
	"math/big"
	"time"

	"bytes"

	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/params"
)

type Account struct {
	PrivateKey string
	Address    common.Address
	Balance    *big.Int
}

type Option struct {
	Period   uint64
	ChainId  *big.Int
	Signers  []common.Address
	Accounts []Account
}

func GetCliqueGenesis(opt *Option) (genesis *core.Genesis, err error) {
	if opt == nil {
		err = fmt.Errorf("nil option")
		return
	}

	// Construct a default genesis block
	genesis = &core.Genesis{
		Timestamp:  uint64(time.Now().Unix()),
		GasLimit:   4700000,
		Difficulty: big.NewInt(1),
		Alloc:      make(core.GenesisAlloc),
		Config: &params.ChainConfig{
			HomesteadBlock: big.NewInt(1),
			EIP150Block:    big.NewInt(2),
			EIP155Block:    big.NewInt(3),
			EIP158Block:    big.NewInt(3),
			ByzantiumBlock: big.NewInt(4),
			Clique: &params.CliqueConfig{
				Period: opt.Period,
				Epoch:  30000,
			},
			ChainId: opt.ChainId,
		},
	}

	// Sort the signers and embed into the extra-data section
	var signers = opt.Signers
	for i := 0; i < len(signers); i++ {
		for j := i + 1; j < len(signers); j++ {
			if bytes.Compare(signers[i][:], signers[j][:]) > 0 {
				signers[i], signers[j] = signers[j], signers[i]
			}
		}
	}
	genesis.ExtraData = make([]byte, 32+len(signers)*common.AddressLength+65)
	for i, signer := range signers {
		copy(genesis.ExtraData[32+i*common.AddressLength:], signer[:])
	}
	// set initial funds
	for _, account := range opt.Accounts {
		genesis.Alloc[account.Address] = core.GenesisAccount{Balance: account.Balance}
	}
	// Add a batch of precompile balances to avoid them getting deleted
	for i := int64(0); i < 256; i++ {
		genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
	}
	return
}
