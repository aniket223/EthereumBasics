package main
import (
"context"
"encoding/hex"
"fmt"
"log"
"github.com/ethereum/go-ethereum/common"
"github.com/ethereum/go-ethereum/ethclient"
)
func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
if err != nil {
log.Fatal(err)
}
contractAddress := common.HexToAddress("0x489cb1dcC13F6fD70D4a554A934C2F490EcF085e")
bytecode, err := client.CodeAt(context.Background(), contractAddress, nil) 
if err != nil {
log.Fatal(err)
}
fmt.Println(hex.EncodeToString(bytecode))
}