package main

import (
	"gitlab.com/localtoast/flourPower"
	"gocv.io/x/gocv"
	"localtoast.net/localtoast/bakery/cannoli"
	"localtoast.net/localtoast/bakery/kitchen"
	//"localtoast.net/localtoast/bakery/flour"
	"github.com/gordonklaus/portaudio"
	"github.com/xlab/pocketsphinx-go/sphinx"
	"localtoast.net/localtoast/bakery/olive"
	"localtoast.net/localtoast/bakery/poptart"
	"localtoast.net/localtoast/bakery/taste"
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

//readStdin uses the "stty" linux command to grab all keyboard inputs and let
//the program handle it
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

//getNote gets the note currently at a certain positon on the screen
//and copies it to the file, "breadbox/toasting"
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

//spatter randomly spatters the word "MARCHELL" to a toast
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

//welcome copies the words "WELCOME_TO_DEEGEE" to a toast
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
	flour.Toast(testToast, "green", "blue")
	testToast = flour.CleanFlecks(testToast)
	//fmt.Printf("\n<:o.o:>")
}

//spawnButton copies a set shape to toast
func spawnButton(label string, xvar int, yvar int, testToast []flour.Bread) {
	//flour.ToastLogger("spawnButton")
	flour.CopyToast("=====", xvar, yvar, 1, testToast)
	flour.CopyToast(("| " + label + " |"), xvar, yvar+1, 1, testToast)
	flour.CopyToast("=====", xvar, yvar+2, 1, testToast)

}

//spawnContext switches on the view string and spawns the correct view on to
//the toast passed in
func spawnContext(view []string, testToast []flour.Bread, testLoaf flour.Loaf) {
	//put different context triggers here
	var args bool
	if len(view) < 2 {
		args = false
	} else {
		args = true
	}

	switch view[0] {
	case "ovo":
		kitchen.Users()
		kitchen.EncryptUsers("snowcrash", "the network", "info@snowcrash.network")
	case "ouo":
		//flour.Flat("_", testToast)
		//testToast = flour.CleanFlecks(testToast)
		win, winLoaf := flour.SpawnWin(11, 11)
		win = flour.CopyToast("@@@@@", 3, 6, 3, win)
		testToast, _ := flour.RelWin(0.5, 0.5, 1, 1, win, testToast, testLoaf, true)
		testToast, _ = flour.RelWin(0.25, 0.25, 1, 1, win, testToast, winLoaf, false)
		flour.Toast(testToast, "green", "black")
	case "owo":
		//flour.Flat("_", testToast)
		testToast = flour.CleanFlecks(testToast)
		spawnButton("$", 30, 2, testToast)
		spawnButton("@", 30, 19, testToast)
		spawnButton("#", 1, 2, testToast)
		spawnButton("4", 1, 19, testToast)
		spawnButton("5", 74, 2, testToast)
		spawnButton("6", 74, 19, testToast)
		flour.SpawnIndex("breadbox/000", 5, 4, testToast, 25, 14)
		flour.SpawnIndex("breadbox/001", 5, 5, testToast, 25, 14)
		//Update the screen
		flour.Toast(testToast, "green", "blue")
		//testToast = flour.CleanFlecks(testToast)
	case "ewe":
		poptart.Glaze("/dev/video0")
		poptart.Small("poptart/101/test00.jpeg")
		poptart.Small("poptart/101/test01.jpeg")
		poptart.Small("poptart/101/test02.jpeg")
		poptart.Small("poptart/101/test03.jpeg")
		poptart.Small("poptart/101/test04.jpeg")
		poptart.Small("poptart/101/test05.jpeg")
		poptart.Small("poptart/101/test06.jpeg")
		poptart.Small("poptart/101/test07.jpeg")
		poptart.Small("poptart/101/test08.jpeg")
		poptart.Small("poptart/101/test09.jpeg")
		indexes := strings.Split("0,1,2,3,4,5,6,7,8,9,0", ",")
		for {
			for i := 0; i < 10; i++ {
				filename := fmt.Sprint("poptart/101/test0" + indexes[i] + ".txt")
				//fmt.Println(filename)
				flour.SpawnIndex(filename, 5, 5, testToast, 25, 14)
				//flour.SpawnIndex("poptart/101/test00.txt", 5, 5, testToast, 25, 14)
				flour.Toast(testToast, "red", "blue")
			}
		}
	case "testColour":
		hello := "WELCOME"
		for _, _ = range hello {
			fmt.Println("\x1b[38;2;255;0;0;255mhallo\x1b[0m")
			fmt.Println("\x1b[38;2;255;0;0;255mhenlo\x1b[0m")

		}
		for i, v := range hello {
			switch i {
			case 0:
				fore := "red"
				back := "black"
				flour.CopyColourToast(string(v), 5+i, 5, 6, fore, back, testToast)
			case 1:
				fore := "black"
				back := "red"
				flour.CopyColourToast(string(v), 5+i, 5, 6, fore, back, testToast)
			case 2:
				fore := "red"
				back := "green"
				flour.CopyColourToast(string(v), 5+i, 5, 6, fore, back, testToast)
			case 3:
				fore := "red"
				back := "yellow"
				flour.CopyColourToast(string(v), 5+i, 5, 6, fore, back, testToast)
			case 4:
				fore := "black"
				back := "blue"
				flour.CopyColourToast(string(v), 5+i, 5, 6, fore, back, testToast)
			case 5:
				fore := "black"
				back := "magenta"
				flour.CopyColourToast(string(v), 5+i, 5, 6, fore, back, testToast)
			case 6:
				fore := "red"
				back := "cyan"
				flour.CopyColourToast(string(v), 5+i, 5, 6, fore, back, testToast)
			}

		}
		flour.Toast(testToast, "none", "none")
		for {
			//fmt.Println("CBREAK to exit")
		}
	case "oco":
		//	ctx := context.Background()
		//	cmd := exec.CommandContext(ctx, "poptart/poptart.py")
		//	cmd.Run()
		testToast = flour.CleanFlecks(testToast)
		button, _ := flour.SpawnWin(5, 5)
		container, containerLoaf := flour.SpawnWin(100, 38)
		//flour.Flat("_", testToast)
		button = flour.CopyToast("$", 2, 3, 1, button)
		container, _ = flour.RelWin(0.33, 0.05, 1, 1, button, container, containerLoaf, true)
		//spawnButton("$", 59, 2, testToast)
		button = flour.CopyToast("@", 2, 3, 1, button)
		container, _ = flour.RelWin(0.33, 0.85, 1, 1, button, container, containerLoaf, true)
		//spawnButton("@", 59, 19, testToast)
		button = flour.CopyToast("#", 2, 3, 1, button)
		container, _ = flour.RelWin(0.03, 0.85, 1, 1, button, container, containerLoaf, true)
		//spawnButton("#", 1, 2, testToast)
		button = flour.CopyToast("4", 2, 3, 1, button)
		container, _ = flour.RelWin(0.03, 0.05, 1, 1, button, container, containerLoaf, true)
		//spawnButton("4", 1, 19, testToast)
		button = flour.CopyToast("5", 2, 3, 1, button)
		container, _ = flour.RelWin(0.85, 0.05, 1, 1, button, container, containerLoaf, true)
		//spawnButton("5", 74, 2, testToast)
		button = flour.CopyToast("6", 2, 3, 1, button)
		container, _ = flour.RelWin(0.85, 0.85, 1, 1, button, container, containerLoaf, true)
		//bag, bagLoaf := flour.SpawnWin(testLoaf.Height, testLoaf.Width)
		//bag, bagLoaf = flour.RelWin(0.03, 0.03, 1, 1, container, bag, bagLoaf, true)
		testToast, _ = flour.RelWin(0.03, 0.99, 1, 1, container, testToast, testLoaf, true)

		testToast, _ = flour.RelWin(0.55, 0.99, 1, 1, container, testToast, testLoaf, true)
		//spawnButton("6", 74, 19, testToast)
		//update this with the autonoodly filename
		webcam, err := gocv.VideoCaptureDevice(0)
		if err != nil {
			fmt.Println("Error opening webcam")
		}
		defer webcam.Close()

		for {
			olive.CreateClient(testToast)
			flour.SpawnIndex("poptart/101/server.txt", 5, 5, testToast, 25, 14)
			flour.SpawnIndex("poptart/101/serve0.txt", 80, 5, testToast, 25, 14)

			ok := cannoli.Capture(webcam, "poptart/101/server.jpeg")
			if !ok {
				fmt.Println("Error capturing picture")
			}
			poptart.Big("poptart/101/server.jpeg")
			flour.Toast(testToast, "red", "blue")

		}

	case "ozo":
		var decoder *sphinx.Decoder
		var in []int16
		var err error
		var stream *portaudio.Stream

		if args && len(view) >= 2 {
			if view[1] == "-sphinx" {
				/////////////////////////////BEGIN SPHINX CODE/////////////////////////////
				const (
					samplesPerChannel = 1024
					sampleRate        = 16000
					channels          = 1
				)
				//	app     = cli.App("tasteTest", "This is a test system to see how well sphinx works with golang.")
				hmm := "/usr/local/share/pocketsphinx/model/en-us/en-us"
				dict := "/usr/local/share/pocketsphinx/model/en-us/cmudict-en-us.dict"
				lm := "/usr/local/share/pocketsphinx/model/en-us/en-us.lm.bin"
				logfile := "/dev/null"
				//	logfile = app.StringOpt("log", "taste.log", "Log file to write log to.")
				stdout := false

				//	outraw  = app.StringOpt("outraw", "", "Specify output dir for RAW recorded sound files (s16le). Directory must exist.")

				portaudio.Initialize()
				defer portaudio.Terminate()
				//	defer portaudio.Terminate()
				//defer listener.Close()
				cfg := sphinx.NewConfig(
					sphinx.HMMDirOption(hmm),
					sphinx.DictFileOption(dict),
					sphinx.LMFileOption(lm),
					sphinx.SampleRateOption(sampleRate),
				)
				if stdout == false {
					sphinx.LogFileOption(logfile)(cfg)
				}
				fmt.Println("Loading CMU sphinx...")
				decoder, err = sphinx.NewDecoder(cfg)
				if err != nil {
					fmt.Println("Error creating decoder!")
				}

				//file, err := os.Create("taste/sound.wav")
				//fileWriter := bufio.NewWriter(file)
				in = make([]int16, 10240)
				stream, err = portaudio.OpenDefaultStream(1, 0, 16000, len(in), in)
				defer stream.Close()
				if err != nil {
					fmt.Println("Error opening default stream.")
				}

				stream.Start()
				defer stream.Stop()
				fmt.Println("Processing")
			}
		}
		/////////////////////////////END SPHINX CODE////////////////////////////

		testToast = flour.CleanFlecks(testToast)
		button, _ := flour.SpawnWin(5, 5)
		container, containerLoaf := flour.SpawnWin(100, 38)
		//flour.Flat("_", testToast)
		button = flour.CopyToast("$", 2, 3, 1, button)
		container, _ = flour.RelWin(0.33, 0.05, 1, 1, button, container, containerLoaf, true)
		//spawnButton("$", 59, 2, testToast)
		button = flour.CopyToast("@", 2, 3, 1, button)
		container, _ = flour.RelWin(0.33, 0.85, 1, 1, button, container, containerLoaf, true)
		//spawnButton("@", 59, 19, testToast)
		button = flour.CopyToast("#", 2, 3, 1, button)
		container, _ = flour.RelWin(0.03, 0.85, 1, 1, button, container, containerLoaf, true)
		//spawnButton("#", 1, 2, testToast)
		button = flour.CopyToast("4", 2, 3, 1, button)
		container, _ = flour.RelWin(0.03, 0.05, 1, 1, button, container, containerLoaf, true)
		//spawnButton("4", 1, 19, testToast)
		button = flour.CopyToast("5", 2, 3, 1, button)
		container, _ = flour.RelWin(0.85, 0.05, 1, 1, button, container, containerLoaf, true)
		//spawnButton("5", 74, 2, testToast)
		button = flour.CopyToast("6", 2, 3, 1, button)
		container, _ = flour.RelWin(0.85, 0.85, 1, 1, button, container, containerLoaf, true)
		//bag, bagLoaf := flour.SpawnWin(testLoaf.Height, testLoaf.Width)
		//bag, bagLoaf = flour.RelWin(0.03, 0.03, 1, 1, container, bag, bagLoaf, true)
		testToast, _ = flour.RelWin(0.03, 0.99, 1, 1, container, testToast, testLoaf, true)

		testToast, _ = flour.RelWin(0.55, 0.99, 1, 1, container, testToast, testLoaf, true)
		webcam, err := gocv.VideoCaptureDevice(0)
		if err != nil {
			fmt.Println("Error opening webcam")
		}
		defer webcam.Close()
		for {
			if args {
				if view[1] == "-sphinx" {
					decoder.StartUtt()
					//fmt.Println("Listening...")
					stream.Read()
					var words string
					decoder.ProcessRaw(in, false, true)
					words, _ = decoder.Hypothesis()
					//fmt.Println("Heard " + words)
					if decoder.IsInSpeech() {
						//fmt.Println("Listening...")
						decoder.ProcessRaw(in, false, true)
						decoder.EndUtt()

						words, _ = decoder.Hypothesis()
						//		fmt.Println("Done listening!")
					}
					decoder.EndUtt()
					if words != "" {
						fmt.Println(words)
					}
					flour.CopyColourToast(words, 50, 7, 1, "red", "yellow", testToast)
				}
			}
			//hack to keep decoder alive
			//if args && len(view) > 2 {
			//	decoder.StartUtt()
			//	decoder.EndUtt()
			//}
			//end hack
			olive.CreateServer(testToast)

			flour.SpawnIndex("poptart/101/serve0.txt", 80, 5, testToast, 25, 14)
			//words := taste.Interpret(in, stream, decoder)

			ok := cannoli.Capture(webcam, "poptart/101/server.jpeg")
			if !ok {
				fmt.Println("Error capturing picture")
			}
			if args && len(view) >= 3 {
				if view[2] == "-personrec" {
					classify := gocv.NewCascadeClassifier()
					defer classify.Close()
					cannoli.CaptureDetect(webcam, "poptart/101/server.jpeg", classify)
					poptart.Big("poptart/101/server.jpeg")
					poptart.Person("poptart/101/server.txt")
					flour.SpawnIndex("poptart/101/server00.txt", 5, 5, testToast, 25, 14)
				}
			} else {
				poptart.Big("poptart/101/server.jpeg")
				flour.SpawnIndex("poptart/101/server.txt", 5, 5, testToast, 25, 14)

			}
			flour.Toast(testToast, "red", "blue")

		}
	case "taste":
		in, stream, decoder := taste.Listen()
		defer taste.Plug(stream, decoder)
		for {
			words := taste.Interpret(in, stream, decoder)
			if words != "" {
				fmt.Println(words)
			}
		}
	case "zmq":
		olive.CreateServer(testToast)
	case "help":
		flour.SpawnIndex("breadbox/help", 5, 5, testToast, 55, 2)
	case "ono":
		//	ctx := context.Background()
		//	cmd := exec.CommandContext(ctx, "poptart/poptart.py")
		//	cmd.Run()
		testToast = flour.CleanFlecks(testToast)
		button, _ := flour.SpawnWin(5, 5)
		container, containerLoaf := flour.SpawnWin(100, 38)
		//flour.Flat("_", testToast)
		button = flour.CopyToast("$", 2, 3, 1, button)
		container, _ = flour.RelWin(0.33, 0.05, 1, 1, button, container, containerLoaf, true)
		//spawnButton("$", 59, 2, testToast)
		button = flour.CopyToast("@", 2, 3, 1, button)
		container, _ = flour.RelWin(0.33, 0.85, 1, 1, button, container, containerLoaf, true)
		//spawnButton("@", 59, 19, testToast)
		button = flour.CopyToast("#", 2, 3, 1, button)
		container, _ = flour.RelWin(0.03, 0.85, 1, 1, button, container, containerLoaf, true)
		//spawnButton("#", 1, 2, testToast)
		button = flour.CopyToast("4", 2, 3, 1, button)
		container, _ = flour.RelWin(0.03, 0.05, 1, 1, button, container, containerLoaf, true)
		//spawnButton("4", 1, 19, testToast)
		button = flour.CopyToast("5", 2, 3, 1, button)
		container, _ = flour.RelWin(0.85, 0.05, 1, 1, button, container, containerLoaf, true)
		//spawnButton("5", 74, 2, testToast)
		button = flour.CopyToast("6", 2, 3, 1, button)
		container, _ = flour.RelWin(0.85, 0.85, 1, 1, button, container, containerLoaf, true)
		//bag, bagLoaf := flour.SpawnWin(testLoaf.Height, testLoaf.Width)
		//bag, bagLoaf = flour.RelWin(0.03, 0.03, 1, 1, container, bag, bagLoaf, true)
		testToast, _ = flour.RelWin(0.03, 0.05, 1, 1, container, testToast, testLoaf, true)

		testToast, _ = flour.RelWin(0.03, 0.55, 1, 1, container, testToast, testLoaf, true)

		testToast, _ = flour.RelWin(0.55, 0.05, 1, 1, container, testToast, testLoaf, true)

		testToast, _ = flour.RelWin(0.55, 0.55, 1, 1, container, testToast, testLoaf, true)
		//spawnButton("6", 74, 19, testToast)
		//update this with the autonoodly filename
		webcam, err := gocv.VideoCaptureDevice(0)
		if err != nil {
			fmt.Println("Error opening webcam")
		}
		defer webcam.Close()

		classify := gocv.NewCascadeClassifier()
		defer classify.Close()

		for {

			flour.SpawnIndex("poptart/101/matt00.txt", 5, 5, testToast, 25, 14)
			flour.SpawnIndex("poptart/101/test00.txt", 5, 26, testToast, 25, 14)
			flour.SpawnIndex("poptart/101/test00.txt", 80, 5, testToast, 25, 14)
			flour.SpawnIndex("poptart/101/test00.txt", 80, 26, testToast, 25, 14)
			//fmt.Printf("0\n<:o.o:>")
			//Update the screen
			//			poptart.Pop("/dev/video1", in)
			//go poptart.Pop("/dev/video1")

			ok := cannoli.CaptureDetect(webcam, "poptart/101/matt00.jpeg", classify)
			if !ok {
				fmt.Println("Error capturing picture")
			}
			//cannoli.Write("poptart/101/mat00.png", img)
			poptart.Small("poptart/101/matt00.jpeg")
			poptart.Person("poptart/101/matt00.txt")
			flour.Toast(testToast, "red", "blue")
			//testToast = flour.CleanFlecks(testToast)
		}
	}
}

func main() {
	//inputs := ""
	//init can be changed
	//	xvar := 81
	//	yvar := 23
	xvar, yvar, testToast, testLoaf := flour.DoughMax()
	fmt.Println(xvar)
	fmt.Println(yvar)
	testToast = flour.Oven(testToast, "_", xvar, yvar)
	//flour.CleanFlecks(testToast)
	flour.Flat("_", testToast)

	//just toasting something
	for {
		flour.Toast(testToast, "green", "black")
		//testToast = flour.CleanFlecks(testToast)
		//fmt.Printf("\n<:o.o:>")
		//fmt.Scan(&inputs)
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		intermediate := strings.Split(input, "\n")
		inputs := strings.Split(intermediate[0], " ")
		fmt.Println(inputs)
		switch inputs[0] {
		case "@":
			//WIP FILE THINGIES
			getNote(testToast, "blah")
		//	flour.SpawnContents(path,
		case "$":
			stdin := make(chan string, 1)
			kill := make(chan bool, 1)
			xpos := 0
			go readStdin(stdin, kill)
			for {
				flour.Toast(testToast, "green", "black")
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

				flour.Toast(testToast, "red", "black")
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
					//flour.Flat("_", testToast)
					spawnContext(inputs, testToast, testLoaf)
					flour.SpawnContents(fmt.Sprint("breadbox/"+thread+".1"), 35, 4, testToast)
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
					//flour.Flat("_", testToast)
					spawnContext(inputs, testToast, testLoaf)
					//flour.SpawnIndex("breadbox/000.1", 35, 4, testToast, 39, 14)
					flour.SpawnContents(fmt.Sprint("breadbox/"+thread+".1"), 35, 4, testToast)
					//fmt.Printf("0\n<:o.o:>")
					flour.CopyToast("#", 4, 5+xpos, 1, testToast)
					flour.CopyToast("#", 30, 5+xpos, 1, testToast)

					xpos++
					// this is a good place to grab the rune printed
				}
			}

		case "spawn":
			flour.Flat("_", testToast)
			time.Sleep(1 * time.Second)
			words := "DEEGEE"
			flour.CopyToast(words, 35, 11, 1, testToast)
			flour.Toast(testToast, "red", "blue")
			//fmt.Printf("_<:o.o:>")
			time.Sleep(1 * time.Second)
			spatter(xvar, yvar, testToast)
			//Now spawn where we want to go
			//				spawnContext("ono", testToast)
			//flour.CleanFlecks(testToast)
			flour.Flat("_", testToast)

			spawnContext(inputs, testToast, testLoaf)
			//from here
			//turn this into spawn_content
			flour.SpawnIndex("breadbox/000.1", 35, 4, testToast, 39, 14)
			flour.SpawnIndex("breadbox/000", 5, 4, testToast, 25, 14)
			flour.SpawnIndex("breadbox/001", 5, 5, testToast, 25, 14)
			//do things with them
		case "ewe":
			spawnContext(inputs, testToast, testLoaf)
		case "owo":
			//flour.Flat("_", testToast)
			//fmt.Printf("_<:o.o:>")
			spawnContext(inputs, testToast, testLoaf)
			flour.SpawnIndex("breadbox/000.1", 35, 4, testToast, 39, 14)
		case "oco":
			spawnContext(inputs, testToast, testLoaf)
		case "ozo":
			spawnContext(inputs, testToast, testLoaf)
		case "taste":
			spawnContext(inputs, testToast, testLoaf)
		case "ovo":
			spawnContext(inputs, testToast, testLoaf)
		case "ouo":
			spawnContext(inputs, testToast, testLoaf)
		case "ono":
			//flour.Flat("_", testToast)
			//fmt.Printf("_<:o.o:>")
			spawnContext(inputs, testToast, testLoaf)
		case "spatter":
			spatter(xvar, yvar, testToast)
		case "welcome":
			welcome(testToast)
		case "flour.Flat":
			flour.Flat("_", testToast)
		case "exit":
			os.Exit(1)
			break
		case "testColour":
			spawnContext(inputs, testToast, testLoaf)
		case "zmq":
			spawnContext(inputs, testToast, testLoaf)
		case "help":
			flour.SpawnIndex("breadbox/help", 5, 5, testToast, 55, 2)
		default:
			flour.SpawnIndex("breadbox/help", 5, 5, testToast, 55, 2)
			//testToast = flour.CleanFlecks(testToast)
		}
	}
}
