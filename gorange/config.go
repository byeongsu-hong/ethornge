package gorange

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	ethutils "github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/fdlimit"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	whisper "github.com/ethereum/go-ethereum/whisper/whisperv6"
	"github.com/frostornge/ethornge/utils"
)

type Config struct {
	NetworkId uint64
	HTTPHost  string
	HTTPPort  int
	WSHost    string
	WSPort    int
	Account   int
	Accounts  []common.Address // count
	Balances  int64            // ETH
}

func defaultNodeConfig() node.Config {
	cfg := node.DefaultConfig
	cfg.Name = "geth"
	cfg.Version = params.VersionWithCommit("")
	cfg.HTTPModules = append(cfg.HTTPModules, "eth", "shh", "miner", "personal", "admin")
	cfg.WSModules = append(cfg.WSModules, "eth", "shh", "miner", "personal", "admin")
	cfg.IPCPath = "geth.ipc"
	return cfg
}

func makeDatabaseHandles() int {
	limit, err := fdlimit.Current()
	if err != nil {
		log.Fatalf("Failed to retrieve file descriptor allowance: %v", err)
	}
	if limit < 2048 {

		if err := fdlimit.Raise(2048); err != nil {
			log.Fatalf("Failed to raise file descriptor allowance: %v", err)
		}
	}
	if limit > 2048 { // cap database file descriptors even if more is available
		limit = 2048
	}
	return limit / 2 // Leave half for networking and other stuff
}

func developerGenesisBlock(period uint64, accounts []common.Address, balance int64) *core.Genesis {
	config := *params.AllCliqueProtocolChanges
	config.Clique.Period = period

	genesis := &core.Genesis{
		Config:     &config,
		ExtraData:  append(append(make([]byte, 32), accounts[0][:]...), make([]byte, 65)...),
		GasLimit:   6283185,
		Difficulty: big.NewInt(1),
		Alloc: map[common.Address]core.GenesisAccount{
			common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
			common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
			common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
			common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
			common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
			common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
			common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
			common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
			accounts[0]:                      {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
		},
	}

	for _, account := range accounts[1:] {
		genesis.Alloc[account] = core.GenesisAccount{Balance: utils.Ether(balance)}
	}

	return genesis
}

func getNode(c Config) (*Node, error) {
	n := defaultNodeConfig()
	nodeConfig := &n

	// set up p2p config
	p2pConfig := &nodeConfig.P2P
	p2pConfig.MaxPeers = 0
	p2pConfig.ListenAddr = ":0"
	p2pConfig.NoDiscovery = true
	p2pConfig.DiscoveryV5 = false

	// set up node config
	nodeConfig.UserIdent = "gorange"
	nodeConfig.DataDir = ""
	nodeConfig.HTTPHost = c.HTTPHost
	nodeConfig.HTTPPort = c.HTTPPort
	nodeConfig.WSHost = c.WSHost
	nodeConfig.WSPort = c.WSPort

	stack, err := node.New(nodeConfig)
	if err != nil {
		return nil, fmt.Errorf("Failed to create the protocol stack: %v", err)
	}

	// set up eth config
	ethConfig := &eth.DefaultConfig
	ethConfig.NetworkId = c.NetworkId
	ethConfig.NoPruning = false
	ethConfig.EnablePreimageRecording = true
	ethConfig.GasPrice = big.NewInt(1)
	ethConfig.DatabaseHandles = makeDatabaseHandles()

	ks := stack.AccountManager().Backends(keystore.KeyStoreType)[0].(*keystore.KeyStore)
	var accounts []common.Address
	for i := 0; i <= c.Account; i++ { // at least one account
		account, err := ks.NewAccount("")
		if err != nil {
			return nil, fmt.Errorf("Failed to create account: %v", err)
		}
		if err := ks.Unlock(account, ""); err != nil {
			return nil, fmt.Errorf("Failed to unlock account: %v", err)
		}

		if i == 0 {
			log.Println("Faucet :", account.Address.Hex())
		} else {
			log.Println("Account [", i, "] :", account.Address.Hex())
		}
		accounts = append(accounts, account.Address)
	}
	ethConfig.Genesis = developerGenesisBlock(uint64(0), append(accounts, c.Accounts...), c.Balances)
	ethConfig.Genesis.Config.ChainID = big.NewInt(int64(c.NetworkId))

	ethutils.RegisterEthService(stack, ethConfig)
	ethutils.RegisterShhService(stack, &whisper.DefaultConfig)

	return &Node{*stack}, nil
}
