# ethornge
A simple, vanila deployment toolkit (or framework), on the top of [go-ethereum](https://github.com/ethereum/go-ethereum).

 * Deployment tools and wrappers
     * Easy contract compile / deployment, without using Truffle
     * Ledger Nano (`usbledger`) support
     * Solidity compiler (solc) wrapper for contract deployment
 * Supports setting up lightweight, embedded private network without using Ganache
     * Only launched during the test period

## Usage

### Contract Deployment (Using embedded private network)
```go
func main() {
    // launch an embedded private network that'll be run
    // during the deployment (or test) session, with an account with 100 ETH.
    key, _ := account.NewKey()
    node, _ := gorange.Launch(
        gorange.DefaultLocalConfig([]common.Address{key.Address}, 100),
    )
    defer node.Stop()

    // connect to the private node
    provider, _ := node.WsProviderWithAccounts(context.Background(), account.Keys{key})
    defer provider.Close()

    // suppose that test.go, a biniding of the contract "Test" exists.
    _, tx, _, err := test.DeployTest(provider.Accounts[0], provider)
    if err != nil {
        log.Println("Failed to deploy contract Test :", err)
        return
    }
    address, err := provider.WaitDeployedWithTimeout(tx, 5 * time.Minute)
    if err != nil {
        log.Println("Timeout! :", err)
        return
    }
    log.Println("Contract address :", address.Hex())
}
```

See `DeploySmartContract()` on `examples.go`.

## Test (Using embedded private network)

```go
// TODO : example
```

## Launch testnet
```
docker build -t airbloc/ethornge .
docker run -d -p 4471:4471 -p 4472:4472 --name ethornge-network airbloc/ethornge
```

`main.go` 참조

## TODO

- [] Revert message 지원
