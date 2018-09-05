# ethornge
A simple, vanila deployment toolkit (or framework), on the top of [go-ethereum](https://github.com/ethereum/go-ethereum).

 * Deployment tools and wrappers
     * Easy contract compile / deployment, without using Truffle
     * Ledger Nano (`usbledger`) support
     * Solidity compiler (solc) wrapper for contract deployment
 * Supports setting up lightweight, embedded private network without using Ganache
     * Only launched during the test period

## Usage

### Deployment (Using embedded private network)
```go
package main

func main() {
}
```

## Launch testnet
```
docker build -t airbloc/ethornge .
docker run -d -p 4471:4471 -p 4472:4472 --name ethornge-network airbloc/ethornge
```

`main.go` 참조

## TODO

- [] Revert message 지원
