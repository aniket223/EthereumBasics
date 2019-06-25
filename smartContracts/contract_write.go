//code to write to smart contracts
package main

import (
	"context"
	"crypto/ecdsa"
	store "ethereumBasics/smartContracts/contracts"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	//dial connection
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	//load privat key
	privateKey, err := crypto.HexToECDSA("303738E4E54FABF06D362104EE3EAB94C134A2CFB8AE8A78A309E346D1AEB76F")
	if err != nil {
		log.Fatal(err)
	}
	//convert to public key and get address
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//get the nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	//get gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//create a keyed transactor
	auth := bind.NewKeyedTransactor(privateKey)
	//set the standard transaction options
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice
	//address of smart contract
	address := common.HexToAddress("0x489cb1dcC13F6fD70D4a554A934C2F490EcF085e")
	//create a new instance of store, which is bound to a specific contract
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	//key and value
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))
	//This method will encode this function call with
	//it's arguments, set it as the data property of the transaction, and sign it with the private key. The
	//result will be a signed transaction object.
	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx sent:", tx.Hash().Hex())
	//to verify key/value were set read the mapping from smart contract
	result, err := instance.Items(nil, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(result[:]))
}
