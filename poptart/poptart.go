package main

import (
	"fmt"
	"os"
	"os/exec"
)

func readStdin(out chan string, in chan bool) {
	//no buffering
	exec.Command("stty","-F", "/dev/tty", "cbreak", "min", "1").Start()
	//no visible output
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	// restore the echoing state when exiting
	defer exec.Command("stty", "-F", "/dev/tty", "echo").Run()

	var b []byte = make([]byte, 1)
	for {
		os.Stdin.Read(b)
		if b != nil {
		fmt.Println("==="+string(b)+"===")
		}
//		select {
//		case <-in:
//			return
//		default:
//			os.Stdin.Read(b)
//			fmt.Printf(">>> %v: ", b)
//			out <- string(b)
//		}
	}
}

func main() {
	defer func() {
		exec.Command("stty", "-f", "/dev/tty", "echo").Run()
	}()
	stdin := make(chan string, 1)
	kill := make(chan bool, 1)

	readStdin(stdin, kill)
	for {
		str := <-stdin

		if str == "0" {
			kill <- true
			close(stdin)
			break
		} else {
			fmt.Println(str)
		}

	}

}
