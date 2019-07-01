package main

import (
	"bytes"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func main() {
	//load private key
	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	if err != nil {
		log.Fatal(err)
	}
	//get public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal(err)
	}
	//get public key in bytes format
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//data to be signed
	data := []byte("hello")
	//hash of the data
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())
	//sign the hash
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hexutil.Encode(signature))
	//get public key of signed data
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	//check if both public keys are sam, if same true as both public keys are same, verfied
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches)
	//different ways to check signature
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches)

	signatureNoRecoverID := signature[:len(signature)-1]
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified)
}
