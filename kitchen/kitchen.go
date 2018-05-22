package kitchen

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/buger/jsonparser"
	flour "gitlab.com/localtoast/flourPower"
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
