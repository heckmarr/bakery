package poptart

import (
	"os/exec"
)

func Pop(device string) {
	cmd := exec.Command("streamer", "-c", device, "-o", "poptart/101/test.jpeg")
	cmd.Run()
}
