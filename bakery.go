package main

import (
	"localtoast.net/localtoast/bakery/flour"
	//	"gitlab.com/localtoast/bakery/oven"
	//	"gitlab.com/localtoast/bakery/loaf"
	//	"gitlab.com/localtoast/bakery/dough"
	"bufio"
	"fmt"
	"math/rand"
	"strings"
	"time"
	//	"io"
	//	"io/ioutil"
	"os"
	"os/exec"
)

func readStdin(out chan string, in chan bool) {
	flour.Toast_Logger("readStdin")
	//no buffering
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	//no visible output
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	// restore the echoing state when exiting
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	var b = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		out <- string(b)

	}
}

func getNote(testToast []flour.Bread, fipath string) {
	xstart := 35
	xend := 74
	ystart := 2
	yend := 17
	//WIP GET FILENAME PROGRAMMATICALLY
	//    ALSO FILE THINGIES
	fe, err := os.Create("breadbox/toasting")
	defer fe.Close()
	var temptoast flour.Bread
	//	74 - 35 = length we need to go over
	//	17 - 2 = height we need to go over
	for y := ystart; y < yend; y++ {
		for x := xstart; x < xend; x++ {
			temptoast = flour.Bread_Getter(x, y, testToast)
			fe.WriteString(temptoast.Label)
		}
		fe.WriteString("\n")
		//		Sync() does the writing to a safe place
		fe.Sync()
	}
	fe.Close()
	if err != nil {
		//fmt.Println("SOMETHING WENT WRONG, AIEEE")
	}
}

func spatter(xvar int, yvar int, testToast []flour.Bread) {
	flour.Toast_Logger("spatter")
	rand.Seed(12)
	welcome := "MARCHELL"
	wel := strings.Split(welcome, "")
	fmt.Println(wel[0])
	for i := 0; i < len(testToast)/24; i++ {
		x := rand.Intn(xvar - 1)
		y := rand.Intn(yvar - 1)
		slice := flour.Bread_Getter(x, y, testToast)
		slice.Label = string(wel[rand.Intn(len(wel)-1)])
		testToast = flour.Bread_Setter(x, y, testToast, slice)
		flour.Toast(testToast)
		fmt.Printf("0\n<:o.o:>")
		//time.Sleep(75*time.Millisecond)
	}
}

func welcome(testToast []flour.Bread) {
	flour.Toast_Logger("welcome")
	welcome := "WELCOME_TO_DEEGEE"
	wel := strings.Split(welcome, "")
	for i := 0; i < len(welcome); i++ {
		//                                        DO STUFF HERE
		slice := flour.Bread_Getter(30+i, 11, testToast)
		slice.Label = string(wel[i])
		testToast = flour.Bread_Setter(30+i, 11, testToast, slice)
	}
	flour.Toast(testToast)
	//fmt.Printf("\n<:o.o:>")
}
func copyToast(welcome string, xvar int, yvar int, yend int, testToast []flour.Bread) {
	flour.Toast_Logger("copyToast")
	wel := strings.Split(welcome, "")
	if yend != 0 {
		for x := yend; x > 0; x-- {
			for i := 0; i < len(welcome); i++ {
				//                                        DO STUFF HERE
				slice := flour.Bread_Getter(xvar+i, yvar-x, testToast)
				slice.Label = string(wel[i])
				testToast = flour.Bread_Setter(xvar+i, yvar-x, testToast, slice)
			}
		}
	}
}

func flat(label string, testToast []flour.Bread) {
	flour.Toast_Logger("flat")
	for i := range testToast {
		testToast[i].Label = label
	}
	flour.Toast(testToast)
	//fmt.Printf("\n<:o.o:>")
}

func spawnButton(label string, xvar int, yvar int, testToast []flour.Bread) {
	flour.Toast_Logger("spawnButton")
	copyToast("=====", xvar, yvar, 1, testToast)
	copyToast(("| " + label + " |"), xvar, yvar+1, 1, testToast)
	copyToast("=====", xvar, yvar+2, 1, testToast)

}

func spawnContents(path string, xvar int, yvar int, testToast []flour.Bread) {
	flour.Toast_Logger("spawnContents")
	slice := flour.Dough(xvar+1, yvar+1)
	slice = flour.Oven(slice, "=", xvar, yvar)
	filo, err := os.Open(path)
	filscan := bufio.NewScanner(filo)
	for filscan.Scan() {
		yvar++
		copyToast(filscan.Text(), xvar, yvar, 1, testToast)

	}

	if err != nil {
		//fmt.Println("Something went wrong!")
	}
}

func spawnIndex(path string, xvar int, yvar int, testToast []flour.Bread, xlen int, yhei int) []flour.Bread {
	flour.Toast_Logger("spawnIndex")
	slice := flour.Dough(xvar+1, yvar+1)
	slice = flour.Oven(slice, "=", xvar, yvar)
	filo, err := os.Open(path)
	filscan := bufio.NewScanner(filo)
	for filscan.Scan() {
		yvar++
		copyToast(filscan.Text(), xvar, yvar, 1, testToast)

	}

	if err != nil {
		//fmt.Println("Something went wrong!")
	}
	return slice

}
func spawnContext(view string, testToast []flour.Bread) {
	//put different context triggers here
	switch view {
	case "owo":
		flat("_", testToast)
		spawnButton("$", 30, 2, testToast)
		spawnButton("@", 30, 19, testToast)
		spawnButton("#", 1, 2, testToast)
		spawnButton("4", 1, 19, testToast)
		spawnButton("5", 74, 2, testToast)
		spawnButton("6", 74, 19, testToast)
		spawnIndex("breadbox/000", 5, 4, testToast, 25, 14)
		spawnIndex("breadbox/001", 5, 5, testToast, 25, 14)
		//Update the screen
		flour.Toast(testToast)
	case "ono":
		//	ctx := context.Background()
		//	cmd := exec.CommandContext(ctx, "poptart/poptart.py")
		//	cmd.Run()

		flat("_", testToast)
		spawnButton("$", 59, 2, testToast)
		spawnButton("@", 59, 19, testToast)
		spawnButton("#", 1, 2, testToast)
		spawnButton("4", 1, 19, testToast)
		spawnButton("5", 74, 2, testToast)
		spawnButton("6", 74, 19, testToast)
		//update this with the autonoodly filename
		for {
			spawnIndex("poptart/101/localtoast.txt", 5, 4, testToast, 25, 14)
			spawnIndex("poptart/101/localtoast.txt", 5, 5, testToast, 25, 14)
			//fmt.Printf("0\n<:o.o:>")
			//Update the screen
			flour.Toast(testToast)
		}
	}
}

func main() {
	input := ""
	//init can be changed
	xvar := 81
	yvar := 23
	testToast := flour.Dough(xvar, yvar)
	testToast = flour.Oven(testToast, "_", xvar, yvar)
	flat("_", testToast)
	//just toasting something
	for {
		flour.Toast(testToast)
		//fmt.Printf("\n<:o.o:>")
		fmt.Scan(&input)
		switch input {
		case "@":
			//WIP FILE THINGIES
			getNote(testToast, "blah")
		//	spawnContents(path,
		case "$":
			stdin := make(chan string, 1)
			kill := make(chan bool, 1)
			xpos := 0
			go readStdin(stdin, kill)
			for {
				flour.Toast(testToast)
				fmt.Printf("0\n<:o.o:>")
				fmt.Printf("_")
				str := <-stdin

				if str == "0" {
					//give echo back to the terminal
					exec.Command("stty", "-F", "/dev/tty", "echo").Run()

					//						kill <- true
					//close(stdin)
					break
				} else {
					copyToast(str, 35+xpos, 5, 1, testToast)
					xpos++
					// this is a good place to grab the rune printed
				}

			}

		case "#":
			stdin := make(chan string, 1)
			kill := make(chan bool, 1)
			xpos := 0
			thread := "0"

			go readStdin(stdin, kill)
			for {
				flour.Toast(testToast)
				fmt.Printf("0\n<:o.o:>")
				fmt.Printf("_")
				str := <-stdin
				if str == "0" {
					exec.Command("stty", "-F", "/dev/tty", "echo").Run()
					//kill <- true
					//close(stdin)
					break
				}
				if str == "k" {
					if xpos < 10 {
						thread = fmt.Sprint("00", xpos)
					}
					if xpos >= 100 {
						thread = fmt.Sprint(xpos)
					}
					if xpos >= 10 {
						thread = fmt.Sprint("0", xpos)
					}
					spawnContext("owo", testToast)
					spawnContents(fmt.Sprint("breadbox/"+thread+".1"), 35, 4, testToast)
					fmt.Printf("0\n<:o.o:>")
					copyToast("#", 4, 5+xpos, 1, testToast)
					copyToast("#", 30, 5+xpos, 1, testToast)

					xpos--
					//END WIP, don't forget to do it upwards too
				}
				if str == "j" {
					//if xpos == 0 {
					//	copyToast("#", 4, 5, 1, testToast)
					//	copyToast("#", 30, 5, 1, testToast)
					//} else{
					//			copyToast("_", 4, 5+xpos-1, 1, testToast)
					//			copyToast("_", 30, 5+xpos-1, 1, testToast)
					//pre and post title hash
					//Spawn Context clears, so we don't need pre
					if xpos < 10 {
						thread = fmt.Sprint("00", xpos)
					}
					if xpos >= 100 {
						thread = fmt.Sprint(xpos)
					}
					if xpos >= 10 {
						thread = fmt.Sprint("0", xpos)
					}
					//}
					spawnContext("owo", testToast)
					//spawnIndex("breadbox/000.1", 35, 4, testToast, 39, 14)
					spawnContents(fmt.Sprint("breadbox/"+thread+".1"), 35, 4, testToast)
					fmt.Printf("0\n<:o.o:>")
					copyToast("#", 4, 5+xpos, 1, testToast)
					copyToast("#", 30, 5+xpos, 1, testToast)

					xpos++
					// this is a good place to grab the rune printed
				}
			}

		case "spawn":
			flat("_", testToast)
			time.Sleep(1 * time.Second)
			copyToast("DEEGEE", 35, 11, 1, testToast)
			flour.Toast(testToast)
			fmt.Printf("_<:o.o:>")
			time.Sleep(1 * time.Second)
			spatter(xvar, yvar, testToast)
			//Now spawn where we want to go
			//				spawnContext("ono", testToast)
			spawnContext("owo", testToast)
			//from here
			//turn this into spawn_content
			spawnIndex("breadbox/000.1", 35, 4, testToast, 39, 14)
			spawnIndex("breadbox/000", 5, 4, testToast, 25, 14)
			spawnIndex("breadbox/001", 5, 5, testToast, 25, 14)
			//do things with them
		case "owo":
			flat("_", testToast)
			fmt.Printf("_<:o.o:>")
			spawnContext("owo", testToast)
			spawnIndex("breadbox/000.1", 35, 4, testToast, 39, 14)

		case "ono":
			flat("_", testToast)
			fmt.Printf("_<:o.o:>")
			spawnContext("ono", testToast)
		case "spatter":
			spatter(xvar, yvar, testToast)
		case "welcome":
			welcome(testToast)
		case "flat":
			flat("_", testToast)
		case "exit":
			os.Exit(1)
			break
		default:
			flour.Toast(testToast)
		}
	}
}
