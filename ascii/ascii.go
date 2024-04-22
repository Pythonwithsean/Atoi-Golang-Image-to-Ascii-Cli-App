package ascii

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"

	"github.com/nfnt/resize"
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
	brightAscii := "@"
	darkAscii := " "
		
	/* Decoded image  after opening reader  */
	imgData,_,err := image.Decode(file)
	if err != nil { 
		log.Fatal(err)
	}
	// Aspect ratio of Old Image 
	// a := float64(imgData.Bounds().Dx()) / float64(imgData.Bounds().Dy())

	//new Widht and Height Using Aspect Ratio 
	newwidth := 0
	// newHeight := int(float64(newwidth) / a)

	// Iterate Over each pixel in image 
	resizedImg := resize.Resize(uint(newwidth), 100 ,imgData,resize.Lanczos2)
	bounds := resizedImg.Bounds()
	

	for y := bounds.Min.Y;  y < bounds.Max.Y; y++{ 
	for x := bounds.Min.X ; x < bounds.Max.X; x++{ 
			r,g,b,_ := resizedImg.At(x,y).RGBA()
			brigthness := (0.299* float64(r) + 0.587* float64(g) + 0.114* float64(b))
			bright :=  brigthness >=  150 
			dark  := brigthness <= 50  
			if bright {
				fmt.Print(brightAscii + " ")	
			}else if dark {
				fmt.Print(darkAscii + " ")
			}
		}
		fmt.Println()
	} 


	
	
	return "", nil

}