package utils

import (
	"fmt"
	_ "image/png"
	"log"
	"net/http"
)

// type image struct { }

func Fetch(url string) (string, error) {
	body, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(body)

	return "", nil
}
