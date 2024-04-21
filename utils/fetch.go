package utils

import (
	"fmt"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// type image struct { }
func Fetch(url string) (string, error) {
	return "", nil
}

func FetchImage(addr string) (string, error) {
	
	// body,err := ioutil.ReadFile(addr)
	reader, err := os.Open(string(strings.TrimSpace(addr)))
	if err != nil {
		log.Fatal("Error Opening File", err)
	}

	switch {
	case filepath.Ext(string(strings.TrimSpace(reader.Name()))) == "jpg":
		fmt.Println("Your File is a JPEG file")
	case filepath.Ext(string(strings.TrimSpace(reader.Name()))) == "png":
		fmt.Println("Your File is a PNG file")
	}

	fmt.Println(reader.Name())

	fmt.Println(reader)
	return "", nil

}
