package cannoli

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"gocv.io/x/gocv"
)

//DetectFace uses a Haar cascade face detection on the supplied gocv.Mat
func DetectFace(img gocv.Mat, classify gocv.CascadeClassifier) gocv.Mat {
	//classify := gocv.NewCascadeClassifier()
	//defer classify.Close()

	if !classify.Load("cannoli/pastry/haarcascade_frontalface_default.xml") {
		fmt.Println("Error loading classifier data.")
	}

	rect := classify.DetectMultiScale(img)
	fmt.Println("Found " + string(len(rect)) + " faces")
	green := color.RGBA{255, 255, 255, 255}
	for _, r := range rect {
		size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
		gocv.Rectangle(&img, r, green, 200)
		pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
		gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, green, 2)
	}
	return img
}

//Contour converts the image to grayscale, calculates contours in binary format
//and draws the results to the image
func Contour(img gocv.Mat) gocv.Mat {
	green := color.RGBA{0, 255, 0, 0}
	greyMat := gocv.NewMat()
	gocv.CvtColor(img, &greyMat, gocv.ColorRGBAToGray)
	contourMat := gocv.NewMat()
	gocv.AdaptiveThreshold(greyMat, &contourMat, 15, gocv.AdaptiveThresholdMean, gocv.ThresholdBinary, 5, 3)
	CalcContours := gocv.FindContours(contourMat, 1, 1)

	gocv.DrawContours(&contourMat, CalcContours, 1, green, 4)

	return contourMat
}

//ResizeAndGray resizes a Mat image, converts it to grayscale and returns it
func ResizeAndGray(img gocv.Mat) gocv.Mat {
	smallMat := gocv.NewMat()
	var point image.Point
	point.X = 64
	point.Y = 64
	gocv.Resize(img, &smallMat, point, 64, 64, 2)

	grayMat := gocv.NewMat()
	gocv.CvtColor(img, &grayMat, gocv.ColorRGBAToGray)

	return grayMat
}

//DetectHead detects average human-in-front-of-the-computer circles and draws them.
//Gives a nice effect for asciifying a person
func DetectHead(img gocv.Mat) gocv.Mat {
	circleMat := gocv.NewMat()
	gocv.HoughCirclesWithParams(img, &circleMat, 3, 2, 60, 61, 25, 57, 64)

	blue := color.RGBA{0, 0, 255, 0}

	for i := 0; i < circleMat.Cols(); i++ {
		v := circleMat.GetVecfAt(0, i)
		x := int(v[0])
		y := int(v[1])
		r := int(v[2])
		gocv.Circle(&img, image.Pt(x, y), r, blue, 64)
	}

	return img
}

//CaptureDetect captures an image from a webcamera as a blocking function.
//ie it will wait until the image is ready before it exits. It will also
//run any detection code we have defined in the Detect function on the image
//before writing it to disk.
func CaptureDetect(webcam *gocv.VideoCapture, path string, classify gocv.CascadeClassifier) bool {
	//webcam, _ := gocv.VideoCaptureDevice(1)
	//defer webcam.Close()
	img := gocv.NewMat()
	defer img.Close()

	//blocking read function
	for ok := webcam.Read(&img); !ok; ok = webcam.Read(&img) {
		if !ok {
			//fmt.Println("Device not ready.")
		} else {
			break
		}
	}
	//img = ResizeAndGray(img)
	img = Contour(img)
	img = DetectHead(img)
	//img = Contour(img)
	//img = DetectFace(img, classify)

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
