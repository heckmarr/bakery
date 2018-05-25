package poptart

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
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
func BigColour(filename string, testToast []flour.Bread) []flour.Bread {

	//	var stringToReturn string
	var returnString []string
	//	returnString := make([]string, len(testToast))
	//	pic := imaging.New(64, 64, nil)
	//	picd, err := imaging.Open(filename)
	pic := gocv.IMRead(filename, gocv.IMReadReducedColor2)
	//pic.Open(filename)
	//if err != nil {
	//	fmt.Println("Error opening pic.")
	//}
	var asciiCode []string
	asciiCode = strings.Split("1-_-+-,-.-i-r-s-X-A-a-e-B-h-M-K-G-S-9-B-A-Z", "-")
	//var code string
	//pic = imaging.Grayscale(pic)
	size := 32
	//filter := imaging.NearestNeighbor
	//colours := pic.ColorModel()
	//fmt.Println(colours)
	picResized := gocv.NewMat()
	var point image.Point
	point.X = 32
	point.Y = 32
	gocv.Resize(pic, &picResized, point, 32, 32, 0)
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
	eightBit := gocv.NewMat()
	picResized.ConvertTo(&eightBit, gocv.MatTypeCV8S)
	for column := 1; column < size; column++ {
		for row := 1; row < size; row++ {
			//fmt.Println(count)
			//var imageColor color.Color
			//r, g, b, a := pic.At(row, column)
			//imageResized, err := picResized.ToImage()
			//		fmt.Println(picResized.Type())
			p = eightBit.GetVeciAt(column, row)
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
			r = math.Floor((float64(p[2]) / 21474820))
			g = math.Floor((float64(p[1]) / 21474820))
			b = math.Floor((float64(p[0]) / 21474820))
			G = int(int(g))
			//fmt.Println(r, g, b, G)
			//		a := "255"
			//			a := imageColor.A
			//imageColor
			num := int(G) / 8
			//num := asciinum / 8
			rS = int(r) + 99
			gS = int(g) + 99
			bS = int(b) + 99
			//aS := string(int(a))
			//aS := "255"
			num = (int(g) + 99) / 9
			if rS <= 0 {
				rS = 1
			}
			if gS <= 0 {
				gS = 1
			}
			if bS <= 0 {
				bS = 1
			}
			if num >= len(asciiCode) {
				num = len(asciiCode) - 1
			}
			stringToReturn := fmt.Sprintln(rS, "-", gS, "-", bS, "-", asciiCode[num])
			returnString = append(returnString, stringToReturn)
			//fmt.Println(returnString)
			//fmt.Println(len(returnString))
		}
	}
	//fmt.Println(rS, gS, bS)

	//if count == column && column == size-1 {
	//testToast = flour.Dye256(asciiCode[0], rS, gS, bS, aS, false, true, testToast, int(row*column), true)
	for i, _ := range testToast {
		if i >= 960 {
			i = 960
		}

		stringToSplit := returnString[i]
		words := strings.Split(stringToSplit, "-")
		//fmt.Println(words)
		slice := flour.BreadGetter(testToast[i].X, testToast[i].Y, testToast)
		rrS := strings.TrimSpace(words[0])
		ggS := strings.TrimSpace(words[1])
		bbS := strings.TrimSpace(words[2])
		code := strings.TrimSpace(words[3])
		word := fmt.Sprint("\033[48;2;", rrS, ";", ggS, ";", bbS, "m", code, "\033[0m")
		slice.Label = word
		slice.Dirty = true
		testToast = flour.BreadSetter(testToast[i].X, testToast[i].Y, testToast, slice)

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
