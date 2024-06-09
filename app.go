package main

import (
<<<<<<< HEAD
=======
	b64 "encoding/base64"
	"fmt"
>>>>>>> develop
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

<<<<<<< HEAD
=======
type apiResponse struct {
	success bool
	error   error
	data    any
}

type apiRequestParams struct {
	start string
	end   string
	total string
}

>>>>>>> develop
func goDotEnvVariable(envFilename string, key string) string {

	// load .env file
	err := godotenv.Load(envFilename)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

<<<<<<< HEAD
=======
func get_request(token string, url string, params apiRequestParams) apiResponse {
	client := &http.Client{}

	response := apiResponse{}
	response.success = false

	req, _ := http.NewRequest("GET", url, nil)

	req.Header = http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {"Basic " + token},
	}

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		response.error = err
	} else {
		response.success = true
	}

	response.data = res

	return response
}

>>>>>>> develop
func main() {
	r := gin.Default()
	var environmentFile = ".env.develop"

	godotenv.Load(environmentFile)
	envBindUrl := goDotEnvVariable(environmentFile, "BIND")

<<<<<<< HEAD
=======
	onetUsername := goDotEnvVariable(environmentFile, "ONET_USERNAME")
	onetPassword := goDotEnvVariable(environmentFile, "ONET_PASSWORD")
	onetBaseUrl := goDotEnvVariable(environmentFile, "ONET_BASE_URL")
	onetToken := b64.StdEncoding.EncodeToString([]byte(onetUsername + ":" + onetPassword))

>>>>>>> develop
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
<<<<<<< HEAD
	r.Run(envBindUrl) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
=======

	r.GET("/job_zones", func(c *gin.Context) {
		url := onetBaseUrl + "mnm/interestprofiler/job_zones"
		params := apiRequestParams{}

		result := get_request(onetToken, url, params)

		if result.success == true {
			c.JSON(http.StatusOK, gin.H{
				"data": result.data,
			})

		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"url":  url,
				"data": result.error,
			})
		}
	})

	r.Run(envBindUrl)
>>>>>>> develop
}
