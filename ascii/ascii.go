package ascii

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

func ImageToAscii(addr string) string {

	file, err := os.Open(string(strings.TrimSpace(addr)))
	if err != nil {
		log.Fatal("Error Opening File: ", err)
	}
	defer file.Close()

	darkAscii := []string{"@", "#", "*", "&", "£", "$", "€", "¥", "o", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "l", "m", "n", "p", "q", "r", "s", "t", "u", "v", "x", "z", "1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	brightAscii := []string{".", ".", "|", ":", "-", " "}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal("Error Decoding File:  ", err)
	}

	resizedImg := resize.Resize(150, 0, img, resize.Lanczos3)
	fmt.Println(resizedImg.Bounds().Max.X, resizedImg.Bounds().Max.Y)
	matrix := make([][]any, resizedImg.Bounds().Max.Y)

	result := ""

	for i := resizedImg.Bounds().Min.Y; i < resizedImg.Bounds().Max.Y; i++ {
		matrix[i] = make([]any, resizedImg.Bounds().Max.X)
		for j := resizedImg.Bounds().Min.X; j < resizedImg.Bounds().Max.X; j++ {
			r, g, b, a := resizedImg.At(j, i).RGBA()
			// Shifting 16 bits to the right to get the 8 bit value
			// 65535 -> 255
			// RGBA uses 16 bits to represent the color value
			// RIGHT SHIFTING removes the 8 bits of padding
			r = r >> 8
			g = g >> 8
			b = b >> 8
			a = a >> 8
			// Brightness Formula -> 0.299×R+0.587×G+0.114×B
			brightness := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
			if a < 255 {
				brightness *= float64(a) / 255.0
			}
			if brightness < 128 {
				result += string(darkAscii[rand.Intn(len(darkAscii))])
			} else {
				result += string(brightAscii[rand.Intn(len(brightAscii))])
			}
		}
		result += "\n"
	}

	fmt.Println(result)
	return result
}
