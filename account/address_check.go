// this code is used to check whether an address is a smart contract address or a random user account address
// a smart contract address always have a bytecode so check for that
package main

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	fmt.Printf("is valid:%v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d"))
	fmt.Printf("is valid:%v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d"))

	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract) // is contract: true
	// a random user account address
	address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	isContract = len(bytecode) > 0
	fmt.Printf("is contract: %v\n", isContract) // is contract: false
}
