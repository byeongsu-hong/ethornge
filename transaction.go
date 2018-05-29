package ethornge

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewSignedTransaction(
	client *ethclient.Client,
	opts *bind.TransactOpts,
	to common.Address,
) (*types.Transaction, error) {
	return NewSignedTransactionWithData(client, opts, to, nil)
}

func NewSignedTransactionWithData(
	client *ethclient.Client,
	opts *bind.TransactOpts,
	to common.Address,
	data []byte,
) (*types.Transaction, error) {
	// nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte
	var err error
	var nonce uint64
	if nonce, err = client.NonceAt(opts.Context, opts.From, nil); err != nil {
		return nil, err
	}

	var tx = types.NewTransaction(nonce, to, opts.Value, opts.GasLimit, opts.GasPrice, data)
	return opts.Signer(types.HomesteadSigner{}, opts.From, tx)
}
