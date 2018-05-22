package kitchen

import (
	"bufio"
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
func EncryptUsers(name string, comment string, email string) openpgp.EntityList {
	//	var hints openpgp.FileHints
	//	hints.FileName = "chefs.json"
	//	var entityList openpgp.EntityList
	//	entity, err := openpgp.NewEntity(name, comment, email, nil)
	//	if err != nil {
	//		fmt.Println("Error creating Entity.")
	//	}
	//	entityList = append(entityList, entity)
	//	fmt.Println(entityList)

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
	carrotPencil := bufio.NewWriter(carrot)
	//carrotPencil.WriteString(entity.Serialize)
	//carrotPencil.Flush()
	food, err := os.Open("kitchen/chefs/cabbage")
	defer food.Close()
	if err != nil {
		fmt.Println("Too few chefs in the kitchen!")
	}
	//	reader, err := io.Reader("kitchen/chefs/chefs.json")
	file, err := os.Open("kitchen/chefs/chefs.json")
	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	//writer := bufio.NewWriter(food)
	//	readWriter := bufio.NewReadWriter(reader, writer)
	if err != nil {
		fmt.Println("Error opening user list.")
	}
	//	write, err := os.Writer(file)
	//	var write io.Writer
	//writer := bufio.NewWriter(file)
	//buf := new(bytes.Buffer)
	text, err := openpgp.Encrypt(carrotPencil, entityList, nil, nil, nil)
	if err != nil {
		fmt.Println("Encryption error!")
	}

	for scanner.Scan() {
		bytesWritten, err := text.Write([]byte(scanner.Text()))
		if err != nil {
			fmt.Println("Error writing encoded text")
		} else {
			fmt.Println(string(bytesWritten) + " bytes written.")

			//		carrotPencil.Write([]byte(bytesWritten))
		}
	}
	text.Close()
	carrotPencil.Flush()
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
