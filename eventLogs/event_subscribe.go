//code to subscribe to events 
//the output of the events are stored in transaction receipts under a logs section

package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	//dial into a websocket enabled ethereum client
	client, err := ethclient.Dial("wss://rinkeby.infura.io/ws")
	if err != nil {
		log.Fatal(err)
	}
	//contract address of contract we created in smartContracts section(Store.sol)
	contractAddress := common.HexToAddress("0x489cb1dcC13F6fD70D4a554A934C2F490EcF085e")
	//filterQuery to read all the events
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	//channel to receive all events
	logs := make(chan types.Log)
	//subscribe to logs
	//SubscribeFilterLogs subscribes to the results of a streaming filter query.
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	//select statement to read in either new log events or subscription error
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vlog := <-logs:
			fmt.Println(vlog)
		}
	}

}
