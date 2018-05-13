package poptart

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)

func Poptart(filename string) {

	image, err := imaging.Open(filename)
	if err != nil {
		fmt.Println("Error opening image.")
	}
	asciicode := strings.Split(" -.-,-:-;-i-r-s-X-A-2-5-3-h-M-H-G-S-#-9-B-&-@", "-")
	image = imaging.Grayscale(image)
	size := 16
	filter := imaging.NearestNeighbor
	image = imaging.Resize(image, size, size, filter)
	//imaging.Save(image, "poptart/101/greyImage.jpeg")
	fileString := fmt.Sprint(filename[0:18] + ".txt")
	asciiFile, err := os.Create(fileString)
	if err != nil {
		fmt.Println("Error opening file.")
	}
	asciiImage := bufio.NewWriter(asciiFile)
	defer asciiFile.Close()
	if err != nil {
		fmt.Println("Error opening text file.")
	}

	for column := 1; column <= size; column++ {
		for row := 1; row <= size; row++ {
			_, g, _, _ := image.At(row, column).RGBA()
			asciinum := int(g) / (25 * 32)
			num := asciinum / 8
			_, err := asciiImage.WriteString(asciicode[num])
			//fmt.Println(bytesWritten)
			if err != nil {
				fmt.Println("Error writing string.")
			}
		}
		asciiImage.WriteString("\n")
	}
	asciiImage.Flush()
}
