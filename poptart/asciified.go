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

	var returnString []string
	pic := gocv.IMRead(filename, gocv.IMReadUnchanged)
	var asciiCode []string
	asciiCode = strings.Split("1-_-+-,-.-i-r-s-X-A-a-e-B-h-M-K-G-S-9-B-A-Z", "-")
	var p gocv.Veci
	var r float64
	var g float64
	var b float64
	var G int
	var rS int
	var gS int
	var bS int
	eightBit := gocv.NewMat()
	pic.ConvertTo(&eightBit, gocv.MatTypeCV8S)
	bigFile, err := os.Open(fileText)
	if err != nil {
		fmt.Println("Error opening file")
	}
	bigScanner := bufio.NewScanner(bigFile)
	var bigChar []string
	for bigScanner.Scan() {
		bigChar = append(bigChar, bigScanner.Text())
	}
	var stringToReturn string
	for row := 1; row < len(bigChar); row++ {
		for column := 1; column < len(bigChar[row]); column++ {

			p = eightBit.GetVeciAt((row*10)+100, column*5)

			r = math.Floor((float64(p[2]) / 21474836))
			g = math.Floor((float64(p[1]) / 21474836))
			b = math.Floor((float64(p[0]) / 21474836))
			G = int(int(g))
			num := int(G) / 8
			rS = int(r) - 25
			gS = int(g) - 25
			bS = int(b) - 25
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
			columnToAdd := column - 3
			if columnToAdd <= 0 {
				columnToAdd = column + 3
			}
			stringToReturn = fmt.Sprintln(rS, "-", gS, "-", bS, "-", bigChar[row][column], "-", columnToAdd, "-", row)
			returnString = append(returnString, stringToReturn)
		}
	}
	count := 0
	for _, v := range returnString {
		stringToSplit := v
		if count >= 32 {
			count = 0
		} else {
			count++
		}
		words := strings.Split(stringToSplit, "-")
		xvar, err := strconv.Atoi(strings.TrimSpace(words[4]))
		yvar, err := strconv.Atoi(strings.TrimSpace(words[5]))

		if err != nil {
			fmt.Println("Not a valid int!")
		}
		slice := flour.BreadGetter(yvar, xvar, testToast)
		rrS := strings.TrimSpace(words[0])
		ggS := strings.TrimSpace(words[1])
		bbS := strings.TrimSpace(words[2])
		code := strings.TrimSpace(words[3])
		word := fmt.Sprint("\033[48;2;", rrS, ";", ggS, ";", bbS, "m", code, "\033[0m")
		word = fmt.Sprint("\033[38;2;", rrS, ";", ggS, ";", bbS, "m", word, "\033[0m")
		slice.Label = word
		slice.Dirty = true
		testToast = flour.BreadSetter(yvar, xvar, testToast, slice)

	}

	return testToast
}

//Big converts an pic passed in via filename to an ascii text file.
func Big(filename string) {

	pic, err := imaging.Open(filename)
	if err != nil {
		fmt.Println("Error opening pic.")
	}
	asciicode := strings.Split(" -.-,-:-.-i-r-s-X-A-2-5-3-h-M-H-G-S-#-9-B-&-@", "-")
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
			r, g, b, _ := pic.At(row, column).RGBA()
			asciinum := int((r + g + b/3)) / (25 * 32)
			num := asciinum / 20
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
