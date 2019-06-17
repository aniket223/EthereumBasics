package main
import(
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func createKS() {
	ks:=keystore.NewKeyStore("./tmp",keystore.StandardScryptN,keystore.StandardScryptP)
	password:="newpassword"
	account,err:=ks.NewAccount(password)
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}

