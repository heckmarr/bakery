package poptart

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
	flour "gitlab.com/localtoast/flourPower"
	"gocv.io/x/gocv"
)

//Small converts an pic passed in via filename to an ascii text file.
func Small(filename string) {

	pic, err := imaging.Open(filename)
	if err != nil {
		fmt.Println("Error opening pic.")
	}
	asciicode := strings.Split(" -.-,-:-;-i-r-s-X-A-2-5-3-h-M-H-G-S-#-9-B-&-@", "-")
	pic = imaging.Grayscale(pic)
	size := 16
	filter := imaging.NearestNeighbor
	pic = imaging.Resize(pic, size, size, filter)
	//imaging.Save(pic, "poptart/101/greypic.jpeg")
	fileString := fmt.Sprint(filename[0:18] + ".txt")
	asciiFile, err := os.Create(fileString)
	if err != nil {
		fmt.Println("Error opening file.")
	}
	asciipic := bufio.NewWriter(asciiFile)
	defer asciiFile.Close()
	if err != nil {
		fmt.Println("Error opening text file.")
	}

	for column := 1; column <= size; column++ {
		for row := 1; row <= size; row++ {
			_, g, _, _ := pic.At(row, column).RGBA()
			asciinum := int(g) / (25 * 32)
			num := asciinum / 8
			_, err := asciipic.WriteString(asciicode[num])
			//fmt.Println(bytesWritten)
			if err != nil {
				fmt.Println("Error writing string.")
			}
		}
		asciipic.WriteString("\n")
	}
	asciipic.Flush()
}

//BigColour converts an pic passed in via filename to an ascii text file.
func BigColour(filename string, testToast []flour.Bread, fileText string) []flour.Bread {

	//	var stringToReturn string
	var returnString []string
	//	returnString := make([]string, len(testToast))
	//	pic := imaging.New(64, 64, nil)
	//	picd, err := imaging.Open(filename)
	pic := gocv.IMRead(filename, gocv.IMReadUnchanged)
	//pic.Open(filename)
	//if err != nil {
	//	fmt.Println("Error opening pic.")
	//}
	//sizer, err := pic.ToImage()
	//if err != nil {
	//	fmt.Println("Error converting image.")
	//}
	//sizeX := sizer.Bounds().Dx()
	//sizeY := sizer.Bounds().Dy()
	//fmt.Println(sizer.Bounds().Dx())
	//fmt.Println(sizer.Bounds().Dy())
	var asciiCode []string
	asciiCode = strings.Split("1-_-+-,-.-i-r-s-X-A-a-e-B-h-M-K-G-S-9-B-A-Z", "-")
	//var code string
	//pic = imaging.Grayscale(pic)
	//size := 32
	//filter := imaging.NearestNeighbor
	//colours := pic.ColorModel()
	//fmt.Println(colours)
	//picResized := gocv.NewMat()
	//var point image.Point
	//point.X = 24
	//point.Y = 32

	//gocv.Resize(pic, &picResized, point, 0, 0, 0)
	//needed for pinhole view
	//rect := image.Rect(64, 64, 164, 164)
	//rect := image.Rect(144, 104, 176, 136)

	//needed for pinhole view
	//picResized = pic.Region(rect)

	//gocv.Resize(pic, &picResized, point, 0, 0, gocv.InterpolationNearestNeighbor)
	//pic = imaging.Resize(pic, size, size, filter)
	//imaging.Save(pic, "poptart/101/greypic.jpeg")
	//	fileString := fmt.Sprint(filename[0:18] + ".txt")
	//	asciiFile, err := os.Create(fileString)
	//	if err != nil {
	//		fmt.Println("Error opening file.")
	//	}
	//	asciipic := bufio.NewWriter(asciiFile)
	//	defer asciiFile.Close()
	//	if err != nil {
	//		fmt.Println("Error opening text file.")
	//	}

	//count := 1
	//	fmt.Println(len(testToast))
	//var returnString string
	//returnString := make([]string, 4096, 4096)
	var p gocv.Veci
	var r float64
	var g float64
	var b float64
	var G int
	var rS int
	var gS int
	var bS int
	//var dont bool
	//	var returnString []string
	//eightBit := gocv.NewMat()
	//picResized.ConvertTo(&eightBit, gocv.MatTypeCV8S)
	//	grayPic := gocv.NewMat()
	eightBit := gocv.NewMat()
	pic.ConvertTo(&eightBit, gocv.MatTypeCV8S)
	//grayImage, err := picResized.ToImage()
	//grayPic := imaging.Grayscale(grayImage)
	//if err != nil {
	//	fmt.Println("Error converting to grayscale.")
	//}
	//picResized = pic.Reshape(32, 24)
	//pinhole view
	//picResized.ConvertTo(&eightBit, gocv.MatTypeCV8S)
	//Big("/poptart/101/server.jpeg")
	bigFile, err := os.Open(fileText)
	if err != nil {
		fmt.Println("Error opening file")
	}
	bigScanner := bufio.NewScanner(bigFile)
	var bigChar []string
	for bigScanner.Scan() {
		//	for _, v := range bigScanner.Text() {
		bigChar = append(bigChar, bigScanner.Text())
		//fmt.Println(bigChar)
		//	}
	}
	var stringToReturn string
	//	fmt.Println(len(bigChar))
	for row := 1; row < len(bigChar); row++ {
		for column := 1; column < len(bigChar[row]); column++ {
			//fmt.Println(count)
			//fmt.Println("cols", column)
			//fmt.Println("rows", row)
			//var imageColor color.Color
			//valueColor := grayPic.NRGBAAt(row, column)
			//rG := valueColor.R
			//gG := valueColor.G
			//bG := valueColor.B
			//imageResized, err := picResized.ToImage()
			//		fmt.Println(picResized.Type())

			p = eightBit.GetVeciAt((row*10)+100, column*5)
			//fmt.Println(column, "cols")
			//fmt.Println(row, "rows")
			//p = eightBit.GetVeciAt(column, row)

			//p = picResized.GetVeciAt(row, column)
			//fmt.Println(p)

			//picResized.Channels()
			//fmt.Println(picResized.Channels())
			//for _, v := range p {
			//	fmt.Println(v)
			//}
			//			fmt.Println(p[0], p[1], p[2])

			//p := picResized.Ptr()

			//imageColor := imageResized.At(row, column)
			//fmt.Println(imageColor)
			//imageColor

			r = math.Floor((float64(p[2]) / 21474836))
			g = math.Floor((float64(p[1]) / 21474836))
			b = math.Floor((float64(p[0]) / 21474836))
			G = int(int(g))
			//fmt.Println(r, g, b, G)
			//		a := "255"
			//			a := imageColor.A
			//imageColor
			num := int(G) / 8
			//num := asciinum / 8
			rS = int(r) - 25
			gS = int(g) - 25
			bS = int(b) - 25
			//rGI := int(rG)
			//gGI := int(gG)
			//bGI := int(bG)
			//num = (rGI + gGI + bGI) / 22
			//aS := string(int(a))
			//aS := "255"
			//fmt.Println(G)
			//num = int(G) / 4

			//fmt.Println(gG)
			if rS <= 0 {
				rS = 1
			}
			if gS <= 0 {
				gS = 1
			}
			if bS <= 0 {
				bS = 1
			}
			if num >= len(asciiCode) || num < 0 {
				num = len(asciiCode) - 1
			}
			//fmt.Println(column)
			//fmt.Println(row)
			columnToAdd := column - 3
			if columnToAdd <= 0 {
				columnToAdd = column + 3
			}
			stringToReturn = fmt.Sprintln(rS, "-", gS, "-", bS, "-", bigChar[row][column], "-", columnToAdd, "-", row)
			returnString = append(returnString, stringToReturn)
			//fmt.Println(returnString)
			//fmt.Println(len(returnString))
		}
	}
	//	fmt.Println(returnString)
	//fmt.Println(rS, gS, bS)

	//if count == column && column == size-1 {
	//testToast = flour.Dye256(asciiCode[0], rS, gS, bS, aS, false, true, testToast, int(row*column), true)
	count := 0
	for _, v := range returnString {
		//if i >= len(returnString) {
		//	i = len(returnString) - 1
		//}
		stringToSplit := v
		//fmt.Println(stringToSplit)
		if count >= 32 {
			count = 0
		} else {
			count++
		}
		words := strings.Split(stringToSplit, "-")
		//fmt.Println(words)

		//fmt.Println(testToast[i].Y)
		//xvar := testToast[i].X
		//yvar := testToast[i].Y
		//		fmt.Println(xvar, yvar)
		xvar, err := strconv.Atoi(strings.TrimSpace(words[4]))
		yvar, err := strconv.Atoi(strings.TrimSpace(words[5]))

		//		xvar = xvar / 2
		//		yvar = yvar * 2

		if err != nil {
			fmt.Println("Not a valid int!")
		}
		slice := flour.BreadGetter(yvar, xvar, testToast)
		rrS := strings.TrimSpace(words[0])
		ggS := strings.TrimSpace(words[1])
		bbS := strings.TrimSpace(words[2])
		code := strings.TrimSpace(words[3])
		word := fmt.Sprint("\033[48;2;", rrS, ";", ggS, ";", bbS, "m", code, "\033[0m")
		slice.Label = word
		slice.Dirty = true
		testToast = flour.BreadSetter(yvar, xvar, testToast, slice)

		//fmt.Printf(word)
	}
	//	return testToast
	//	count = 0
	//} else {
	//	testToast = flour.Dye256(asciiCode[0], rS, gS, bS, aS, false, true, &testToast, int(row*column), false)
	//}
	//fmt.Println(string(asciiCode[num]))
	//fmt.Printf(code)

	//returnString[column] += code
	//_, err := asciipic.WriteString(code)
	//fmt.Println(bytesWritten)
	//if err != nil {
	//	fmt.Println("Error writing string.")
	//}

	//fmt.Println("")
	//returnString += "\n"
	//asciipic.WriteString("\n")

	//asciipic.Flush()
	//flour.Toast256(testToast)
	//flour.Toast(testToast, "none", "none")
	return testToast
}

//Big converts an pic passed in via filename to an ascii text file.
func Big(filename string) {

	pic, err := imaging.Open(filename)
	if err != nil {
		fmt.Println("Error opening pic.")
	}
	asciicode := strings.Split(" -.-,-:-;-i-r-s-X-A-2-5-3-h-M-H-G-S-#-9-B-&-@", "-")
	pic = imaging.Grayscale(pic)
	size := 32
	filter := imaging.NearestNeighbor
	pic = imaging.Resize(pic, size, size, filter)
	//imaging.Save(pic, "poptart/101/greypic.jpeg")
	fileString := fmt.Sprint(filename[0:18] + ".txt")
	asciiFile, err := os.Create(fileString)
	if err != nil {
		fmt.Println("Error opening file.")
	}
	asciipic := bufio.NewWriter(asciiFile)
	defer asciiFile.Close()
	if err != nil {
		fmt.Println("Error opening text file.")
	}

	for column := 1; column <= size; column++ {
		for row := 1; row <= size; row++ {
			_, g, _, _ := pic.At(row, column).RGBA()
			asciinum := int(g) / (25 * 32)
			num := asciinum / 8
			_, err := asciipic.WriteString(asciicode[num])
			//fmt.Println(bytesWritten)
			if err != nil {
				fmt.Println("Error writing string.")
			}
		}
		asciipic.WriteString("\n")
	}
	asciipic.Flush()
}
