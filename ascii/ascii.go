package ascii

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"
)



func ImageToAscii(addr string) (string, error) {

	// asciiChars := map[string]string{
		
	// }
	
	// body,err := ioutil.ReadFile(addr)
	file, err := os.Open(string(strings.TrimSpace(addr)))
	if err != nil {
		log.Fatal("Error Opening File", err)
	}

	// Close Reader When Function gets popped off Call stack
	defer file.Close()

	// switch ext := filepath.Ext(string(reader.Name()));  ext {
	// case ".jpeg":
	// 	fmt.Println("Your File is a JPEG file")
	// case ".png":
	// 	fmt.Println("Your File is a PNG file")
	// }

		/* To Calculate the Brightmess from Rgb values it can be done bt converting the rgb values of each pixel into a Greyscale value 
		I have Chosen to use the Luminance formula grayscale = 0.2126 * Red + 0.7152 * Green + 0.00722 * Blue this formula will take into account
		the humans eyes sensitivity to different colors and provies a better representation of brighness 
		*/
		//Luminance (perceived option 1): (0.299*R + 0.587*G + 0.114*B)

		//Ascii symbols
	brightAscii := " "
	darkAscii := "#"
		
	/* Decoded image  after opening reader  */
	imgData,_,err := image.Decode(file)
	if err != nil { 
		log.Fatal(err)
	}

	// Iterate Over bounds 
	bounds := imgData.Bounds()
	for x := bounds.Min.X; x < bounds.Max.X; x++{ 
		for y := bounds.Min.Y;  y< bounds.Max.Y; y++{ 
			r,g,b,_ := imgData.At(x,y).RGBA()
			brigthness := (0.299* float64(r) + 0.587* float64(g) + 0.114* float64(b))
			bright :=  brigthness >=  150 
			dark  := brigthness <= 100  
			if bright {
				fmt.Print(brightAscii + " ")

			}else if dark {
				fmt.Print(darkAscii + " ")
			}
		}
	} 
	
	
	return "", nil

}