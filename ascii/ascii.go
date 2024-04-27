package ascii

import (
	"fmt"
	"image"
	"log"
	"os"
	"strings"

	"github.com/nfnt/resize"
)

// func convertToGreyScale( img image.Image) image.Image{
// 	//Create New Grey Image Frame with the bounds of the Old Image
// 	greyImg := image.NewGray(img.Bounds())

// 	for y := greyImg.Bounds().Min.Y; y < greyImg.Bounds().Max.Y; y++{
// 		for x:= greyImg.Bounds().Min.X; x < greyImg.Bounds().Max.X; x++{
// 				pixelAtLocation := img.At(x,y)
// 				grayColor := color.GrayModel.Convert(pixelAtLocation)
// 				greyImg.Set(x,y,grayColor)

// 		}
// 	}
// 	return greyImg
// }

func ImageToAscii(addr string)  {

	// asciiChars := map[string]string{
		
	// }
	
	// body,err := ioutil.ReadFile(addr)
	file, err := os.Open(string(strings.TrimSpace(addr)))
	if err != nil {
		log.Fatal("Error Opening File", err)
	}


	// Close Reader When Function gets popped off Call stack
	defer file.Close()

		/* To Calculate the Brightmess from Rgb values it can be done bt converting the rgb values of each pixel into a Greyscale value 
		I have Chosen to use the Luminance formula grayscale = 0.2126 * Red + 0.7152 * Green + 0.00722 * Blue this formula will take into account
		the humans eyes sensitivity to different colors and provies a better representation of brighness 
		*/
		//Luminance (perceived option 1): (0.299*R + 0.587*G + 0.114*B)

		//Ascii symbols
	brightAscii := []string{"@", "+", "="}

	darkAscii := []string{" ", ".", "-" }
		
	/* Decoded image so that i can work with the image data like pixels for each pixel   */
	imgData,_,err := image.Decode(file)
	if err != nil { 
		log.Fatal(err)
	}

	// greyScaledImage := convertToGreyScale(imgData)


	// Resizen Image to smaller res for Terminal
	resizedImg := resize.Resize(0 , 100 ,imgData,resize.Lanczos2)
	bounds := resizedImg.Bounds()
	for y := bounds.Min.Y;  y < bounds.Max.Y; y++{ 
		for x := bounds.Min.X ; x < bounds.Max.X; x++{ 
		//This Returns the Color Codes for 16bit Color Channels from 0 - 2^16
			r,g,b,_ := resizedImg.At(x,y).RGBA()
			// color := image.RGBA()
			brigthness := (0.299* float64(r) + 0.587* float64(g) + 0.114* float64(b))
			bright :=  brigthness >=  20000
			brightest := brigthness >= 30000 
			dark  := brigthness <= 1500
			darked := brigthness <= 300
			// fmt.Println("Color for RGB ad Brightness",r,g,b,brigthness)
			if bright {
				fmt.Print(brightAscii[0]  + " ")	
			}else if brightest{
				fmt.Print(brightAscii[2] + " ")
			}else if darked{
				fmt.Print(darkAscii[1] + " ")
			}else if dark {
				fmt.Print(darkAscii[0] + " ")
			}
		}
		fmt.Println()
	} 
}