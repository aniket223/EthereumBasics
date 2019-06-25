package main

import (
	"context"
	"crypto/ecdsa"
	store "ethereumBasics/smartContracts/contracts" // for demo
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//create a keyed transactor
	auth := bind.NewKeyedTransactor(privateKey)
	//set the standard transaction options
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)// in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice
	input := "1.0"
	//deploy ethereum contract binding an instance of Store to it
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())

	fmt.Println(tx.Hash().Hex())
	_ = instance

}
