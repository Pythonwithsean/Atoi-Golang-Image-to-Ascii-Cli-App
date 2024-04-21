package utils

import (
	"image"
	_ "image/png"
	"log"
	"os"
	"strings"
)

// type image struct { }

func Fetch(url string) (string, error) {
	return "", nil
}

func FetchImage(addr string) (string, error) {

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

		//Ascii symbols
		
	/* Decoded image  after opening reader  */
	imgData,_,err := image.Decode(file)
	if err != nil { 
		log.Fatal(err)
	}
	bounds := imgData.Bounds()
	for y := bounds.Min.Y;  y< bounds.Max.Y; y++{ 
		for x := bounds.Min.X; x < bounds.Max.X; x++{ 
			r,g,b,a := imgData.At(x,y).RGBA()


		}
	} 
	
	
	return "", nil

}
