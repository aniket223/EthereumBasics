package main
import (
"fmt"
"log"
"github.com/ethereum/go-ethereum/common"
"github.com/ethereum/go-ethereum/ethclient"
store "./contracts" // for demo
)
func main() {
client, err := ethclient.Dial("https://rinkeby.infura.io")
if err != nil {
log.Fatal(err)
}
address := common.HexToAddress("0x489cb1dcC13F6fD70D4a554A934C2F490EcF085e")
instance, err := store.NewStore(address, client)
if err != nil {
log.Fatal(err)
}
version, err := instance.Version(nil)
if err != nil {
log.Fatal(err)
}
fmt.Println(version) // "1.0"
}