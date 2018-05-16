package cannoli

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

func Capture(path string) bool {
	webcam, _ := gocv.VideoCaptureDevice(0)
	defer webcam.Close()
	img := gocv.NewMat()
	defer img.Close()

	for ok := webcam.Read(&img); !ok; ok = webcam.Read(&img) {
		if !ok {
			//fmt.Println("Device not ready.")
		} else {
			break
		}
	}
	file, err := os.Create(path)
	defer file.Close()
	if err != nil {
		fmt.Println("Error creating file.")
	}
	ok := gocv.IMWrite(path, img)
	if !ok {
		fmt.Println("Error writing file.")
	}
	return ok
}
