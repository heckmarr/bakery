package asciified

import (
	"fmt"
	"os"
	"strings"

	"github.com/disintegration/imaging"
)

func main() {
	image, err := imaging.Open("101/image.png")
	if err != nil {
		fmt.Println("Error opening image.")
	}
	asciicode := strings.Split(" -.-,-:-;-i-r-s-X-A-2-5-3-h-M-H-G-S-#-9-B-&-@", "-")
	imageG := imaging.Grayscale(image)
	size := 32
	var filter imaging.ResampleFilter
	imageG = imaging.Resize(image, size, size, filter)
	asciiImage, err := os.Open("101/local.txt")
	defer asciiImage.Close()
	if err != nil {
		fmt.Println("Error opening text file.")
	}

	for column := 0; column < size; column++ {
		for row := 0; row < size; row++ {
			r, g, b, _ := imageG.At(row, column).RGBA()
			asciinum := int(r) + int(g) + int(b)/24

			asciiImage.WriteString(asciicode[asciinum])
		}
		asciiImage.WriteString("\n")
	}
}
