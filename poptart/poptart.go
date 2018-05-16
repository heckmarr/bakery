package poptart

import (
	"os/exec"
)

//Pop grabs a single frame via the linux command, "streamer"
//Slow and depreceated, only works in linux
func Pop(device string) {
	cmd := exec.Command("streamer", "-c", device, "-o", "poptart/101/test00.jpeg")
	cmd.Run()
}

//Glaze grabs a ten second video via the linux command, "streamer"
//Depreceated as it only works in linux.
func Glaze(device string) {
	cmd := exec.Command("streamer", "-c", device, "-t", "10", "-r", "1", "-o", "poptart/101/test00.jpeg")
	cmd.Run()
}
