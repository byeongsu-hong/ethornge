package provider

import (
	"fmt"
	"log"
	"math/big"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

func (pv *Provider) NewSignedTransaction(
	opts *bind.TransactOpts,
	to common.Address,
) (*types.Transaction, error) {
	return pv.NewSignedTransactionWithData(opts, to, nil)
}

func (pv *Provider) NewSignedTransactionWithData(
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

	if opts.GasPrice == nil {
		gasPrice, err := pv.SuggestGasPrice(pv.Context)
		if err != nil {
			return nil, err
		}
		opts.GasPrice = gasPrice
	}

	if opts.GasLimit == 0 {
		gasLimit, err := pv.EstimateGas(pv.Context, ethereum.CallMsg{
			From:     opts.From,
			To:       &to,
			Gas:      0,
			GasPrice: opts.GasPrice,
			Value:    opts.Value,
			Data:     data,
		})
		if err != nil {
			return nil, err
		}
		opts.GasLimit = gasLimit
	}

	if opts.Context == nil {
		return nil, fmt.Errorf("nil context option")
	}

	nonce, err := pv.NonceAt(opts.Context, opts.From, nil)
	if err != nil {
		return nil, err
	}
	tx := types.NewTransaction(nonce, to, opts.Value, opts.GasLimit, opts.GasPrice, data)
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
		big.NewInt(params.GWei),
	)

	log.Println("TxHash           : ", receipt.TxHash.Hex())
	log.Println("TxReceipt Status : ", status)
	log.Println("Gas Limit        : ", tx.Gas())
	log.Println("Gas Used         : ", receipt.GasUsed)
	log.Println("Gas Price        : ", gasPrice, "Gwei")
}

func PrintTxResultWithTime(tx *types.Transaction, receipt *types.Receipt, start time.Time) {
	log.Printf("Confirmed        : %.3f s\n", time.Since(start).Seconds())
	PrintTxResult(tx, receipt)
}
