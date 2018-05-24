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
func BigColour(filename string, testToast []flour.Bread) string {
	//	pic := imaging.New(64, 64, nil)
	//	picd, err := imaging.Open(filename)
	pic := gocv.IMRead(filename, gocv.IMReadReducedColor2)
	//pic.Open(filename)
	//if err != nil {
	//	fmt.Println("Error opening pic.")
	//}
	var asciiCode []string
	asciiCode = strings.Split("1-_-+-,-.-i-r-s-X-A-a-e-B-h-M-K-G-S-9-B-A-Z", "-")

	//pic = imaging.Grayscale(pic)
	size := 31
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
	var rS int
	var gS int
	var bS int
	var returnString string
	//returnString := make([]string, 4096, 4096)
	for column := 1; column < size; column++ {
		for row := 1; row < size; row++ {
			//var imageColor color.Color
			//r, g, b, a := pic.At(row, column)
			//imageResized, err := picResized.ToImage()
			p := picResized.GetVeciAt(row, column)
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
			r := math.Floor((float64(p[2]) / 21474836))
			g := math.Floor((float64(p[1]) / 21474836))
			b := math.Floor((float64(p[0]) / 21474836))
			G := int(int(g))

			//fmt.Println(r, g, b, G)
			//		a := "255"
			//			a := imageColor.A
			//imageColor
			num := int(G) / 8
			//num := asciinum / 8
			rS = int(r)
			gS = int(g)
			bS = int(b)
			//aS := string(int(a))
			aS := "255"
			if rS <= 0 {
				rS = 1
			}
			if gS <= 0 {
				gS = 1
			}
			if bS <= 0 {
				bS = 1
			}
			//fmt.Println(rS, gS, bS, aS)
			if num >= len(asciiCode) {
				num = len(asciiCode) - 1
			}
			//fmt.Println(string(asciiCode[num]))
			flour.Dye256(asciiCode[7], rS, gS, bS, aS, false, true, &testToast, int(row*column))
			//fmt.Printf(code)

			//returnString += code
			//_, err := asciipic.WriteString(code)
			//fmt.Println(bytesWritten)
			//if err != nil {
			//	fmt.Println("Error writing string.")
			//}
		}
		fmt.Println("")
		//returnString += "\n"
		//asciipic.WriteString("\n")
	}
	//asciipic.Flush()
	//flour.Toast256(testToast)
	//flour.Toast(testToast, "none", "none")
	return returnString
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
