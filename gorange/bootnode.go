package gorange

import (
	"net"

	"github.com/ethereum/go-ethereum/p2p/discv5"
	"github.com/frostornge/ethornge/account"
)

func NewBootNode(key *account.Key) *discv5.Node {
	return discv5.NewNode(
		discv5.PubkeyID(&key.PrivateKey.PublicKey),
		net.ParseIP("127.0.0.1"),
		2478, 2479,
	)
}
