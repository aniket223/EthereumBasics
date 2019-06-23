//in this code we setup a subscription to get events when a new block is mined
//First thing is we need an Ethereum provider that supports RPC over websockets
//we use infura websocket endpoint
package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	//dial in the connection
	client, err := ethclient.Dial("wss://ropsten.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}
	//make a channel
	headers := make(chan *types.Header)
	//takes in header channel and returns a subscription object
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	//select the case whether it is a failure or subscription
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println(header.Hash().Hex())
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Number().Uint64())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
		}
	}
}
