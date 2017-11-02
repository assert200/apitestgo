package main

import (
	"log"
	"net/url"

	"github.com/assert200/gorest"
)

func main() {
	restRequest := gorest.Request{}

	var url url.URL
	url.Scheme = "https"
	url.Host = "www.google.com"
	url.Path = "/"

	restRequest.Method = "GET"
	restRequest.URL = url

	session := gorest.CreateNewSession()

	restResponse, err := gorest.Do(session, restRequest)

	if err != nil {
		log.Fatal("There was an error excuting the rest request: " + err.Error())
	}

	if restResponse.StatusCode != 200 {
		log.Fatal("Expecting Status Code 200, Recieved: ", restResponse.StatusCode)
	}

	log.Println("PASS: Response Status Code: ", restResponse.StatusCode)
}
