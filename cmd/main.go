package main

import (
	"github.com/Pythonwithsean/Atoi-Golang-Image-to-Ascii-Cli-App/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	Dir = "uploads"
)

func main() {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}                   // Adjust this to your frontend URL
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}            // Allowed HTTP methods
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"} // Allowed headers

	r := gin.Default()

	// Apply CORS middleware to the router
	r.Use(cors.New(corsConfig))

	r.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	r.POST("/convert", func(c *gin.Context) {
		utils.UploadFile(c)
	})

	r.Run()
}
