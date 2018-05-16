package cannoli

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

//Detect uses a Haar cascade face detection on the supplied gocv.Mat
func Detect(img gocv.Mat) gocv.Mat {
	classify := gocv.NewCascadeClassifier()
	defer classify.Close()

	if !classify.Load("cannoli/pastry/haarcascade_frontalface_default.xml") {
		fmt.Println("Error loading classifier data.")
	}

	rect := classify.DetectMultiScale(img)
	fmt.Println("Found " + string(len(rect)) + " faces")
	green := color.RGBA{0, 0, 255, 0}
	for _, r := range rect {
		size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
		gocv.Rectangle(&img, r, green, 5)
		pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
		gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, green, 2)
	}
	return img
}

//CaptureDetect captures an image from a webcamera as a blocking function.
//ie it will wait until the image is ready before it exits. It will also
//run any detection code we have defined in the Detect function on the image
//before writing it to disk.
func CaptureDetect(path string) bool {
	webcam, _ := gocv.VideoCaptureDevice(1)
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
	img = Detect(img)
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
