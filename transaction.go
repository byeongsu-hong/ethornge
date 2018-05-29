package ethornge

import (
	"fmt"

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
	if len(opts.From) == 0 {
		return nil, fmt.Errorf("nil from option")
	}

	if opts.Signer == nil {
		return nil, fmt.Errorf("nil signer option")
	}

	if opts.GasPrice == nil || opts.GasLimit == 0 {
		return nil, fmt.Errorf("no gas option")
	}

	if opts.Context == nil {
		return nil, fmt.Errorf("nil context option")
	}

	var nonce, err = client.NonceAt(opts.Context, opts.From, nil)
	if err != nil {
		return nil, err
	}
	var tx = types.NewTransaction(nonce, to, opts.Value, opts.GasLimit, opts.GasPrice, data)
	return opts.Signer(types.HomesteadSigner{}, opts.From, tx)
}
