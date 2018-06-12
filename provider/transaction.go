package provider

import (
	"fmt"
	"log"
	"time"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
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

func PrintTxResult(tx *types.Transaction, receipt *types.Receipt) {
	var status string
	if receipt.Status == 0 {
		status = "Failed"
	} else {
		status = "Success"
	}

	var gasPrice = new(big.Int).Div(
		tx.GasPrice(),
		big.NewInt(params.Shannon),
	)

	log.Println("TxHash           : ", receipt.TxHash.Hex())
	log.Println("TxReceipt Status : ", status)
	log.Println("Gas Limit        : ", tx.Gas())
	log.Println("Gas Used         : ", receipt.GasUsed)
	log.Println("Gas Price        : ", gasPrice, "Gwei")
}

func PrintTxResultWithTime(tx *types.Transaction, receipt *types.Receipt, start time.Time) {
	log.Printf("Confirmed : %.3f s\n", time.Since(start).Seconds())
	PrintTxResult(tx, receipt)
}
