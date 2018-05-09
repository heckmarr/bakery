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
	//flour.ToastLogger("readStdin")
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
			temptoast = flour.BreadGetter(x, y, testToast)
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
	//flour.ToastLogger("spatter")
	rand.Seed(12)
	welcome := "MARCHELL"
	wel := strings.Split(welcome, "")
	fmt.Println(wel[0])
	for i := 0; i < len(testToast); i++ {
		x := rand.Intn(xvar - 1)
		y := rand.Intn(yvar - 1)
		slice := flour.BreadGetter(x, y, testToast)
		slice.Label = string(wel[rand.Intn(len(wel)-1)])
		testToast = flour.BreadSetter(x, y, testToast, slice)
		flour.PrintFleck(i, testToast)
		//fmt.Printf("0\n<:o.o:>")
		//time.Sleep(75*time.Millisecond)
	}
}

func welcome(testToast []flour.Bread) {
	//flour.ToastLogger("welcome")
	welcome := "WELCOME_TO_DEEGEE"
	wel := strings.Split(welcome, "")
	for i := 0; i < len(welcome); i++ {
		//                                        DO STUFF HERE
		slice := flour.BreadGetter(30+i, 11, testToast)
		slice.Label = string(wel[i])
		testToast = flour.BreadSetter(30+i, 11, testToast, slice)
	}
	flour.Toast(testToast)
	testToast = flour.CleanFlecks(testToast)
	//fmt.Printf("\n<:o.o:>")
}

func flat(label string, testToast []flour.Bread) {
	//flour.ToastLogger("flat")
	for i := range testToast {
		testToast[i].Label = label
		testToast[i].Dirty = true
		//flour.CopyToast(label, testToast[i].X, testToast[i].Y, testToast[i].Y+1, testToast)

	}
	flour.Toast(testToast)
	//fmt.Printf("\n<:o.o:>")
}

func spawnButton(label string, xvar int, yvar int, testToast []flour.Bread) {
	//flour.ToastLogger("spawnButton")
	flour.CopyToast("=====", xvar, yvar, 1, testToast)
	flour.CopyToast(("| " + label + " |"), xvar, yvar+1, 1, testToast)
	flour.CopyToast("=====", xvar, yvar+2, 1, testToast)

}

func spawnContents(path string, xvar int, yvar int, testToast []flour.Bread) {
	//flour.ToastLogger("spawnContents")
	filo, err := os.Open(path)
	filscan := bufio.NewScanner(filo)
	for filscan.Scan() {
		yvar++
		flour.CopyToast(filscan.Text(), xvar, yvar, 1, testToast)

	}

	if err != nil {
		//fmt.Println("Something went wrong!")
	}
}

func spawnIndex(path string, xvar int, yvar int, testToast []flour.Bread, xlen int, yhei int) {
	//flour.ToastLogger("spawnIndex")
	filo, err := os.Open(path)
	filscan := bufio.NewScanner(filo)
	for filscan.Scan() {
		yvar++
		flour.CopyToast(filscan.Text(), xvar, yvar, 1, testToast)

	}

	if err != nil {
		//fmt.Println("Something went wrong!")
	}
	//return slice

}
func spawnContext(view string, testToast []flour.Bread) {
	//put different context triggers here
	switch view {
	case "ouo":
		flat("+", testToast)
		//testToast = flour.CleanFlecks(testToast)
		win := flour.SpawnWin(5, 5)
		testToast = flour.RelWin(0.5, 0.5, 5, 5, win, testToast)
		flour.Toast(testToast)
	case "owo":
		//flat("_", testToast)
		testToast = flour.CleanFlecks(testToast)
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
		//testToast = flour.CleanFlecks(testToast)
	case "ono":
		//	ctx := context.Background()
		//	cmd := exec.CommandContext(ctx, "poptart/poptart.py")
		//	cmd.Run()
		testToast = flour.CleanFlecks(testToast)

		//flat("_", testToast)
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
			//testToast = flour.CleanFlecks(testToast)
		}
	}
}

func main() {
	input := ""
	//init can be changed
	//	xvar := 81
	//	yvar := 23
	xvar, yvar, testToast := flour.DoughMax()
	fmt.Println(xvar)
	fmt.Println(yvar)
	testToast = flour.Oven(testToast, "_", xvar, yvar)
	//flour.CleanFlecks(testToast)
	flat("_", testToast)

	//just toasting something
	for {
		flour.Toast(testToast)
		//testToast = flour.CleanFlecks(testToast)
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
				//testToast = flour.CleanFlecks(testToast)
				//fmt.Printf("0\n<:o.o:>")
				//fmt.Printf("_")
				str := <-stdin

				if str == "0" {
					//give echo back to the terminal
					exec.Command("stty", "-F", "/dev/tty", "echo").Run()

					//						kill <- true
					//close(stdin)
					break
				} else {
					flour.CopyToast(str, 35+xpos, 5, 1, testToast)
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
				//testToast = flour.CleanFlecks(testToast)
				//fmt.Printf("0\n<:o.o:>")
				//fmt.Printf("_")
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
					//flour.CleanFlecks(testToast)
					//flat("_", testToast)
					spawnContext("owo", testToast)
					spawnContents(fmt.Sprint("breadbox/"+thread+".1"), 35, 4, testToast)
					//fmt.Printf("0\n<:o.o:>")
					flour.CopyToast("#", 4, 5+xpos, 1, testToast)
					flour.CopyToast("#", 30, 5+xpos, 1, testToast)

					xpos--
					//END WIP, don't forget to do it upwards too
				}
				if str == "j" {
					//if xpos == 0 {
					//	flour.CopyToast("#", 4, 5, 1, testToast)
					//	flour.CopyToast("#", 30, 5, 1, testToast)
					//} else{
					//			flour.CopyToast("_", 4, 5+xpos-1, 1, testToast)
					//			flour.CopyToast("_", 30, 5+xpos-1, 1, testToast)
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
					//flour.CleanFlecks(testToast)
					//flat("_", testToast)
					spawnContext("owo", testToast)
					//spawnIndex("breadbox/000.1", 35, 4, testToast, 39, 14)
					spawnContents(fmt.Sprint("breadbox/"+thread+".1"), 35, 4, testToast)
					//fmt.Printf("0\n<:o.o:>")
					flour.CopyToast("#", 4, 5+xpos, 1, testToast)
					flour.CopyToast("#", 30, 5+xpos, 1, testToast)

					xpos++
					// this is a good place to grab the rune printed
				}
			}

		case "spawn":
			flat("_", testToast)
			time.Sleep(1 * time.Second)
			flour.CopyToast("DEEGEE", 35, 11, 1, testToast)
			flour.Toast(testToast)
			//fmt.Printf("_<:o.o:>")
			time.Sleep(1 * time.Second)
			spatter(xvar, yvar, testToast)
			//Now spawn where we want to go
			//				spawnContext("ono", testToast)
			//flour.CleanFlecks(testToast)
			flat("_", testToast)

			spawnContext("owo", testToast)
			//from here
			//turn this into spawn_content
			spawnIndex("breadbox/000.1", 35, 4, testToast, 39, 14)
			spawnIndex("breadbox/000", 5, 4, testToast, 25, 14)
			spawnIndex("breadbox/001", 5, 5, testToast, 25, 14)
			//do things with them
		case "owo":
			//flat("_", testToast)
			//fmt.Printf("_<:o.o:>")
			spawnContext("owo", testToast)
			spawnIndex("breadbox/000.1", 35, 4, testToast, 39, 14)
		case "ouo":
			spawnContext("ouo", testToast)
		case "ono":
			//flat("_", testToast)
			//fmt.Printf("_<:o.o:>")
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
			//testToast = flour.CleanFlecks(testToast)
		}
	}
}
