package kitchen

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/buger/jsonparser"
	flour "gitlab.com/localtoast/flourPower"
	"golang.org/x/crypto/openpgp"
)

//spawnIndex writes out a file to toast.
func spawnIndex(path string, xvar int, yvar int, testToast []flour.Bread, xlen int, yhei int) {
	//flour.ToastLogger("spawnIndex")
	filo, err := os.Open(path)
	filscan := bufio.NewScanner(filo)
	for filscan.Scan() {
		yvar++
		flour.CopyToast(filscan.Text(), xvar, yvar, 1, testToast)

	}

	if err != nil {
		fmt.Println("Something went wrong!")
	}
	//return slice

}
func InTheKitchen(testToast []flour.Bread, testLoaf flour.Loaf) []flour.Bread {

	container, _ := flour.SpawnWin(testLoaf.Height, testLoaf.Width/2)

	spawnIndex("kitchen/chefs/chefs.txt", 5, 5, container, 25, 14)
	flour.RelWin(1, 0.5, 1, 1, container, testToast, testLoaf, true)

	fmt.Println("Currently no one is in the kitchen")
	return testToast
}

func decryptCarrot() {
	path := "/home/twotonne/.gnupg"
	secRingPrefix := path + "/secring.gpg"
	keyRingBuf, err := os.Open(secRingPrefix)
	defer keyRingBuf.Close()
	if err != nil {
		fmt.Println("Error opening ring file.")
	}
	entityList, err := openpgp.ReadKeyRing(keyRingBuf)
	if err != nil {
		fmt.Println("Error reading keyring.")
	}

	carrot, err := ioutil.ReadFile("kitchen/chefs/carrot")
	if err != nil {
		fmt.Println("Error reading.")
	}
	soup, err := os.Create("kitchen/chefs/soup")
	if err != nil {
		fmt.Println("Error creating soup.")
	}

	//Only needed if the private key has a passphrase
	//err = key[0].PrivateKey.Decrypt(carrot)
	mess, err := openpgp.ReadMessage(bytes.NewBuffer(carrot), entityList, nil, nil)
	if err != nil {
		fmt.Println("Error reading carrot.")
	}
	json, err := ioutil.ReadAll(mess.UnverifiedBody)
	if err != nil {
		fmt.Println("Error reading body of message.")
	}

	writer := bufio.NewWriter(soup)
	writer.Write(json)
	writer.Flush()
	fmt.Println(json)
	for {

	}
}
func EncryptUsers(name string, comment string, email string) openpgp.EntityList {
	//	Make sure to gpg --gen-keys

	//  IMPORTANT, CHANGE USER TO YOUR USER
	path := "/home/twotonne/.gnupg"
	//secRingPrefix := path + "/secring.gpg"
	pubRingPrefix := path + "/pubring.gpg"

	keyRingBuf, err := os.Open(pubRingPrefix)
	defer keyRingBuf.Close()

	entityList, err := openpgp.ReadKeyRing(keyRingBuf)
	if err != nil {
		fmt.Println("Error reading public ring!")
	}

	carrot, err := os.Create("kitchen/chefs/carrot")
	defer carrot.Close()
	carrotPencil := bufio.NewWriter(carrot)
	if err != nil {
		fmt.Println("Too few chefs in the kitchen!")
	}
	file, err := ioutil.ReadFile("kitchen/chefs/chefs.json")
	//reader := bufio.NewReader(file)
	//scanner := bufio.NewScanner(reader)
	if err != nil {
		fmt.Println("Error opening user list.")
	}
	text, err := openpgp.Encrypt(carrotPencil, entityList, nil, nil, nil)
	if err != nil {
		fmt.Println("Encryption error!")
	}

	bytesWritten, err := text.Write(file)
	if err != nil {
		fmt.Println("Error writing encoded text.")
	} else {
		fmt.Println(string(bytesWritten) + " bytes written.")
	}

	//	for scanner.Scan() {
	//		bytesWritten, err := text.Write([]byte(scanner.Text()))
	//		if err != nil {
	//			fmt.Println("Error writing encoded text")
	//		} else {
	//			fmt.Println(string(bytesWritten) + " bytes written.")

	//		}
	//	}
	text.Close()
	carrotPencil.Flush()

	decryptCarrot()
	return entityList
}
func Users() {
	file, err := ioutil.ReadFile("kitchen/chefs/chefs.json")
	if err != nil {
		fmt.Println("Error opening chefs list.")
	}
	KeyID, err := jsonparser.GetString(file, "Keys", "KeyID")
	if err != nil {
		fmt.Println("Error getting key!")
	}
	fmt.Println(KeyID)

}
