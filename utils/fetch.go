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

	switch ext := filepath.Ext(string(reader.Name()));  ext {
	case ".jpeg":
		fmt.Println("Your File is a JPEG file")
	case ".png":
		fmt.Println("Your File is a PNG file")
	}

	// reader := base64.NewDecoder() 

	
	return "", nil

}
