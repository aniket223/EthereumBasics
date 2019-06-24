Smart Contracts

In this section, we will learn how to compile, deploy, query, read and write
smart contracts using golang.

First, you must have a [solidity
compiler](https://solidity.readthedocs.io/en/latest/installing-solidity.html)
installed. To install it use:

> sudo snap install solc --edge

After that, you need to get ethereum.

> go get -u github.com/ethereum/go-ethereum

Then go into the package.

> cd $GOPATH/src/github.com/ethereum/go-ethereum

We also need to install a tool called abigen for generating the ABI from a
solidity smart contract.

> go  install ./cmd/abigen

Now we can generate the ABI from solidity source file.(assuming we keep the file
name as Store.sol)

> solc --abi Store.sol -o build

To convert the ABI to a go file use the following command.

> abigen --abi=./build/Store.abi --pkg=store --out=Store.go

We also need to compile the solidity smart contract to EVM bytecode to deploy
using golang.We will send EVM bytecode in the data field of the transaction. The
bin file is required for generating the deploy methods on the Go contract file.

> solc --bin Store.sol -o build

Now we compile the Go contract file which will include the deploy methods
because we have included the bin file.

> abigen --bin=./build/Store.bin --abi=./build/Store.abi --pkg=store --out=Store.go
