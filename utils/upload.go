package utils

import (
	"fmt"
	"io"
	"os"

	"github.com/Pythonwithsean/Atoi-Golang-Image-to-Ascii-Cli-App/ascii"
	"github.com/gin-gonic/gin"
)

const (
	dir = "uploads"
)

func UploadFile(c *gin.Context) {
	fmt.Println("Convert Endpoint Hit")

	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "No file found",
		})
		return
	}
	defer file.Close()

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to create directory",
		})
		return
	}

	filePath := dir + "/" + fileHeader.Filename
	fmt.Println(filePath)

	out, err := os.Create(filePath)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "iFailed to create file",
		})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Failed to save file",
		})
		return
	}
	imageAscii := ascii.ImageToAscii(filePath)
	c.JSON(200, gin.H{
		"art": imageAscii,
	})

}
