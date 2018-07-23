package gorange

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/common/fdlimit"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/params"
	whisper "github.com/ethereum/go-ethereum/whisper/whisperv6"
)

type Config struct {
	NetworkId uint64
	HTTPPort  int
	WSPort    int
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
	nodeConfig.HTTPHost = "127.0.0.1"
	nodeConfig.HTTPPort = c.HTTPPort
	nodeConfig.WSHost = "127.0.0.1"
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
	developer, err := ks.NewAccount("")
	if err != nil {
		return nil, fmt.Errorf("Failed to create developer account: %v", err)
	}
	if err := ks.Unlock(developer, ""); err != nil {
		return nil, fmt.Errorf("Failed to unlock developer account: %v", err)
	}
	log.Println("Using developer account", "address", developer.Address.Hex())
	ethConfig.Genesis = core.DeveloperGenesisBlock(uint64(0), developer.Address)

	utils.RegisterEthService(stack, ethConfig)
	utils.RegisterShhService(stack, &whisper.DefaultConfig)

	return &Node{*stack}, nil
}
