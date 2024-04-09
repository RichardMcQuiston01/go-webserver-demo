package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func goDotEnvVariable(envFilename string, key string) string {

	// load .env file
	err := godotenv.Load(envFilename)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	r := gin.Default()
	var environmentFile = ".env.develop"

	godotenv.Load(environmentFile)
	envBindUrl := goDotEnvVariable(environmentFile, "BIND")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(envBindUrl) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
