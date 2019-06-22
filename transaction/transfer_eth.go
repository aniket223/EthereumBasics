// this code is to show how to transfer ether
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	// dial in the connection
	//also keep here the testnet/mainnet in which you have ethercl
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	// generate a private key or load a private key as in the next comment
	/*privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}*/

	// load private key, enter a private key of a account which has some ether
	privateKey, err := crypto.HexToECDSA("b51f95c6528edd8fb1a8bef7a67546e39a51d4400a9575ac4a870c6109f3cbdb")
	if err != nil {
		log.Fatal(err)
	}

	//get an interface containing public key
	publicKey := privateKey.Public()
	//convert public key into any type you want, here we convert it into *ecdsa.PublicKey
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	//sender's address
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// It's hard to keep manual track of all the nonces so the ethereum client provides a
	//helper method PendingNonceAt that will return the next nonce you should use.
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	// just to check balance in your account
	balance, err := client.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)
	//set the value you want to transfer, gasLimit and get the gas Price
	value := big.NewInt(1000000000000000000) //in wei
	gasLimit := uint64(21000)                //in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//address to which we need to send ether
	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	//creating a transaction, it takes the parameters and we put data in nil if you want to
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, nil)
	//chainID for EIP155signer
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//get the signed transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	//send the signed transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
