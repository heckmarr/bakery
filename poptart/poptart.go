package poptart

import (
	"os/exec"
)

func Pop(device string) {
	cmd := exec.Command("streamer", "-c", device, "-o", "poptart/101/test.jpeg")
	cmd.Run()
}

func Glaze(device string) {
	cmd := exec.Command("streamer", "-c", device, "-t", "10", "-r", "1", "-o", "poptart/101/test00.jpeg")
	cmd.Run()
}
