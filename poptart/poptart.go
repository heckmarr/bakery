package poptart

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

//Pop grabs a single frame via the linux command, "streamer"
//Slow and depreceated, only works in linux
func Pop(device string) {
	cmd := exec.Command("streamer", "-c", device, "-o", "poptart/101/test00.jpeg")
	cmd.Run()
}

//Person replaces all circle detected portions, detected as they use a particular color,
//with the words, "PERSON", a silly effect, but useful.
func Person(path string) {
	name := "PERSON"
	file, err := os.Open(path)
	fileReader := bufio.NewReader(file)
	if err != nil {
		fmt.Println("Error opening file.")
	}
	count := 0
	fileScanner := bufio.NewScanner(fileReader)
	newFile, err := os.Create(path[:18] + "00.txt")
	if err != nil {
		fmt.Println("Error creating file.")
	}
	newFileWriter := bufio.NewWriter(newFile)
	defer newFile.Close()
	for fileScanner.Scan() {
		words := fileScanner.Text()
		if count > 5 {
			count = 0
		}
		for c := 0; c < len(words); c++ {
			if count > 5 {
				count = 0
			}
			if string(words[c]) == "2" {
				newFileWriter.WriteString(string(name[count]))
				count++
			} else {
				newFileWriter.WriteString(" ")
			}
		}
		newFileWriter.WriteString("\n")
	}
	newFileWriter.Flush()

}

//Glaze grabs a ten second video via the linux command, "streamer"
//Depreceated as it only works in linux.
func Glaze(device string) {
	cmd := exec.Command("streamer", "-c", device, "-t", "10", "-r", "1", "-o", "poptart/101/test00.jpeg")
	cmd.Run()
}
