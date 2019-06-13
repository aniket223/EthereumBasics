// this code is used to get ethereum address of a particular hex

package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	//for getting address
	address := common.HexToAddress("0x71c7656ec7ab88b098defb751b7401b5f6d8976f")
	//for getting the hex
	fmt.Println(address.Hex())
	//for getting the hash
	fmt.Println(address.Hash().Hex())
	//for getting the adress, address and address.Bytes() is the same thing
	fmt.Println(address.Bytes())

}
