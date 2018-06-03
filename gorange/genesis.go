package gorange

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"../account"
	"../utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/params"
)

type Genesis core.Genesis

func (g Genesis) export(path string) (err error) {
	var marshaled []byte
	if marshaled, err = core.Genesis(g).MarshalJSON(); err != nil {
		return
	}

	var out bytes.Buffer
	if err = json.Indent(&out, marshaled, "", "    "); err != nil {
		return
	}
	return utils.WriteFile(path, out.Bytes())
}

type GenesisOption struct {
	FilePath  string
	Consensus bool // true - clique, false - ethash
	Period    uint64
	ChainId   *big.Int
	Signers   []common.Address
	Accounts  account.Accounts
}

func (gnOpt *GenesisOption) init(geth *GethOption) (err error) {
	// get genesis config
	var genesis *Genesis
	if gnOpt.Consensus {
		if genesis, err = gnOpt.getClique(); err != nil {
			return
		}
	} else {
		if genesis, err = gnOpt.getEthash(); err != nil {
			return
		}
	}

	// export genesis file
	if err = genesis.export(gnOpt.FilePath); err != nil {
		return
	}

	return
}

func (gnOpt *GenesisOption) newGenesis() (*Genesis, error) {
	if gnOpt == nil {
		return nil, fmt.Errorf("nil option")
	}

	// Construct a default genesis block
	return &Genesis{
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
				Period: gnOpt.Period,
				Epoch:  30000,
			},
			ChainId: gnOpt.ChainId,
		},
	}, nil
}

func (gnOpt *GenesisOption) setInitialFund(genesis *Genesis) {
	// set initial funds
	for _, ac := range gnOpt.Accounts {
		genesis.Alloc[ac.Address] = core.GenesisAccount{Balance: ac.Balance}
	}
	// Add a batch of precompile balances to avoid them getting deleted
	for i := int64(0); i < 256; i++ {
		genesis.Alloc[common.BigToAddress(big.NewInt(i))] = core.GenesisAccount{Balance: big.NewInt(1)}
	}
}

func (gnOpt *GenesisOption) getEthash() (genesis *Genesis, err error) {
	if genesis, err = gnOpt.newGenesis(); err != nil {
		return
	}

	genesis.Config.Ethash = new(params.EthashConfig)
	genesis.ExtraData = make([]byte, 32)
	gnOpt.setInitialFund(genesis)

	return
}

func (gnOpt *GenesisOption) getClique() (genesis *Genesis, err error) {
	if genesis, err = gnOpt.newGenesis(); err != nil {
		return
	}

	// Sort the signers and embed into the extra-data section
	var signers = gnOpt.Signers
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
	gnOpt.setInitialFund(genesis)

	return
}
