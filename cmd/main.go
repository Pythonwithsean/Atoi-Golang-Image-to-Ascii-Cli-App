package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Pythonwithsean/Atoi-Golang-Image-to-Ascii-Cli-App/utils"
)

func main() { 
	fmt.Printf(utils.Yellow + utils.IntialAscii)
	fmt.Println(utils.Cyan + "Welcome to Atoi")
	fmt.Println("A Image to Ascii Cli Application")
	fmt.Print("Search for an Image? ")
	buff := bufio.NewReader(os.Stdin)
	input,err :=buff.ReadString('\n')
	if err != nil { 
		panic(err)
	}
	fmt.Println(input)
}