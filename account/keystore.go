// this code is used to generate and import keystore
package main
import(
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

//create a keystore
func createKS() {
	//create keystore for given directory
	ks:=keystore.NewKeyStore("./tmp",keystore.StandardScryptN,keystore.StandardScryptP)
	password:="newpassword"
	//generate a new key and store it into the key directory
	account,err:=ks.NewAccount(password)
	if err!=nil{
		log.Fatal(err)
	}
	// print address
	fmt.Println(account.Address.Hex())
}

//import a keystore
func importKS() {

	//after ./tmp/ enter the name of the generated keystore in the tmp file
	file := "./tmp/UTC--2019-06-17T07-25-37.439572263Z--2ac1eb92a9366e30d0547ba9c817ef3f218cb198"
ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
//reading the file
jsonBytes,err:= ioutil.ReadFile(file)
if err!=nil{
	log.Fatal(err)
}
//setting the passphrase
password:="newpassword"
// we keep the passphrase same, for this example we can change it as well
account, err := ks.Import(jsonBytes, password, password)
if err != nil {
	log.Fatal(err)
	}
	//print address
	fmt.Println(account.Address.Hex())
	// no use of keeping the same thing twice so delete old one
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
		}
}

func main(){
	createKS()
	//first use only createKS then when you run code second time import the with importKS using the name of the file
	//importKS()
}

