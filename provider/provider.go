package provider

import (
	"context"

	"math/big"

	"time"

	"fmt"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ErrTimeout = fmt.Errorf("Timeout")

	DefaultLedgerOption = &Option{
		Path:  "m/44'/60'/0'/",
		Start: 0,
		End:   1,
	}
)

type Option struct {
	// Common option
	URL       string
	Context   context.Context
	Network   string
	NetworkID int

	// PrivateKey provider Option
	Keys []string

	// Keystore provider Option
	Keypath  string
	Password string

	// Ledger provider Option
	Path  string
	Start int
	End   int
}

type Provider struct {
	*ethclient.Client
	Context  context.Context
	Accounts []*bind.TransactOpts
}

func LedgerProvider(opt *Option) (provider *Provider, err error) {
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = getClient(opt)
	if err != nil {
		return
	}
	provider.Accounts, err = getLedgerOpts(opt)
	return
}

func KeystoreProvider(opt *Option) (provider *Provider, err error) {
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = getClient(opt)
	if err != nil {
		return
	}
	var account *bind.TransactOpts
	if account, err = getKeystoreOpt(opt); err != nil {
		return
	}
	provider.Accounts = []*bind.TransactOpts{account}
	return
}

func PrivateKeyProvider(opt *Option) (provider *Provider, err error) {
	provider = new(Provider)
	provider.Context = opt.Context
	provider.Client, err = getClient(opt)
	if err != nil {
		return
	}
	provider.Accounts, err = getPrivateKeyOpts(opt)
	return
}

/* Provider only */
func (pv *Provider) WaitMinedWithTimeout(tx *types.Transaction, d time.Duration) (*types.Receipt, error) {
	ctx, cancel := context.WithTimeout(pv.Context, d)
	defer cancel()

	receipt, err := bind.WaitMined(ctx, pv.Client, tx)
	if err != nil {
		return nil, err
	}
	if receipt == nil && err == nil {
		return nil, ErrTimeout
	}
	return receipt, nil
}

func (pv *Provider) WaitDeployedWithTimeout(tx *types.Transaction, d time.Duration) (common.Address, error) {
	ctx, cancel := context.WithTimeout(pv.Context, d)
	defer cancel()

	addr, err := bind.WaitDeployed(ctx, pv.Client, tx)
	if err != nil {
		return common.Address{}, err
	}
	if addr == (common.Address{}) && err == nil {
		return common.Address{}, ErrTimeout
	}
	return addr, nil
}

/* Provider bind */
func (pv *Provider) BlockByHash(hash common.Hash) (*types.Block, error) {
	return pv.Client.BlockByHash(pv.Context, hash)
}

func (pv *Provider) BlockByNumber(number *big.Int) (*types.Block, error) {
	return pv.Client.BlockByNumber(pv.Context, number)
}

func (pv *Provider) HeaderByHash(hash common.Hash) (*types.Header, error) {
	return pv.Client.HeaderByHash(pv.Context, hash)
}

func (pv *Provider) HeaderByNumber(number *big.Int) (*types.Header, error) {
	return pv.Client.HeaderByNumber(pv.Context, number)
}

func (pv *Provider) TransactionByHash(hash common.Hash) (*types.Transaction, bool, error) {
	return pv.Client.TransactionByHash(pv.Context, hash)
}

func (pv *Provider) TransactionSender(tx *types.Transaction, block common.Hash, index uint) (common.Address, error) {
	return pv.Client.TransactionSender(pv.Context, tx, block, index)
}

func (pv *Provider) TransactionCount(blockHash common.Hash) (uint, error) {
	return pv.Client.TransactionCount(pv.Context, blockHash)
}

func (pv *Provider) TransactionInBlock(blockHash common.Hash, index uint) (*types.Transaction, error) {
	return pv.Client.TransactionInBlock(pv.Context, blockHash, index)
}

func (pv *Provider) TransactionReceipt(txHash common.Hash) (*types.Receipt, error) {
	return pv.Client.TransactionReceipt(pv.Context, txHash)
}

func (pv *Provider) SyncProgress() (*ethereum.SyncProgress, error) {
	return pv.Client.SyncProgress(pv.Context)
}

func (pv *Provider) SubscribeNewHead(ch chan<- *types.Header) (ethereum.Subscription, error) {
	return pv.Client.SubscribeNewHead(pv.Context, ch)
}

func (pv *Provider) NetworkID() (*big.Int, error) {
	return pv.Client.NetworkID(pv.Context)
}

func (pv *Provider) BalanceAt(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return pv.Client.BalanceAt(pv.Context, account, blockNumber)
}

func (pv *Provider) StorageAt(account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	return pv.Client.StorageAt(pv.Context, account, key, blockNumber)
}

func (pv *Provider) CodeAt(account common.Address, blockNumber *big.Int) ([]byte, error) {
	return pv.Client.CodeAt(pv.Context, account, blockNumber)
}

func (pv *Provider) NonceAt(account common.Address, blockNumber *big.Int) (uint64, error) {
	return pv.Client.NonceAt(pv.Context, account, blockNumber)
}

func (pv *Provider) FilterLogs(q ethereum.FilterQuery) ([]types.Log, error) {
	return pv.Client.FilterLogs(pv.Context, q)
}

func (pv *Provider) SubscribeFilterLogs(q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	return pv.Client.SubscribeFilterLogs(pv.Context, q, ch)
}

func (pv *Provider) PendingBalanceAt(account common.Address) (*big.Int, error) {
	return pv.Client.PendingBalanceAt(pv.Context, account)
}

func (pv *Provider) PendingStorageAt(account common.Address, key common.Hash) ([]byte, error) {
	return pv.Client.PendingStorageAt(pv.Context, account, key)
}

func (pv *Provider) PendingCodeAt(account common.Address) ([]byte, error) {
	return pv.Client.PendingCodeAt(pv.Context, account)
}

func (pv *Provider) PendingNonceAt(account common.Address) (uint64, error) {
	return pv.Client.PendingNonceAt(pv.Context, account)
}

func (pv *Provider) PendingTransactionCount() (uint, error) {
	return pv.Client.PendingTransactionCount(pv.Context)
}

func (pv *Provider) CallContract(msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	return pv.Client.CallContract(pv.Context, msg, blockNumber)
}

func (pv *Provider) PendingCallContract(msg ethereum.CallMsg) ([]byte, error) {
	return pv.Client.PendingCallContract(pv.Context, msg)
}

func (pv *Provider) SuggestGasPrice() (*big.Int, error) {
	return pv.Client.SuggestGasPrice(pv.Context)
}

func (pv *Provider) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	return pv.Client.EstimateGas(pv.Context, msg)
}

func (pv *Provider) SendTransaction(tx *types.Transaction) error {
	return pv.Client.SendTransaction(pv.Context, tx)
}

func (pv *Provider) WaitMined(tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(pv.Context, pv.Client, tx)
}

func (pv *Provider) WaitDeployed(tx *types.Transaction) (common.Address, error) {
	return bind.WaitDeployed(pv.Context, pv.Client, tx)
}
