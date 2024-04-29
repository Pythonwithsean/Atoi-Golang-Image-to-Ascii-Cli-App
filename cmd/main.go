package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Pythonwithsean/Atoi-Golang-Image-to-Ascii-Cli-App/ascii"
	"github.com/Pythonwithsean/Atoi-Golang-Image-to-Ascii-Cli-App/utils"
)

func main() {
	fmt.Printf(utils.Red + utils.IntialAscii)
	fmt.Println(utils.Yellow + "Welcome to Atoi")
	fmt.Println("A Image to Ascii Cli Application")
	fmt.Print("Search for an Image? ")
	buff := bufio.NewReader(os.Stdin)
	input, err := buff.ReadString('\n')
	if err != nil {
		panic(err)
	}

	ascii.ImageToAscii(input)
}
