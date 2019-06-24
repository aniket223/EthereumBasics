package main
import (
"fmt"
"log"
"github.com/ethereum/go-ethereum/common"
"github.com/ethereum/go-ethereum/ethclient"
store "ethereumBasics/smartContracts/contracts" // for demo
)

func main() {
	client,err:=ethclient.Dial("https://rinkeby.infura.io")
	if err!=nil{
		log.Fatal(err)
	}
	//adress of smart contract
	address:=common.HexToAddress("0x489cb1dcC13F6fD70D4a554A934C2F490EcF085e")
	//NewStore is a funtion in Store.go used to create an instance of Store bound to a deployed contract 
	instance,err:= store.NewStore(address,client)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("contract loaded")
	_ = instance
}