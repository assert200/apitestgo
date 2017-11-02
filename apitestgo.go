package main

import (
	"log"
	"net/url"

	"github.com/assert200/gorest"
)

func main() {
	request := gorest.NewRequest()

	var url url.URL
	url.Scheme = "https"
	url.Host = "www.google.com"
	url.Path = "/"

	request.Method = "GET"
	request.URL = url

	session := gorest.NewSession()

	response, err := gorest.Do(session, request)

	if err != nil {
		log.Fatal("There was an error excuting the rest request: " + err.Error())
	}

	if response.StatusCode != 200 {
		log.Fatal("Expecting Status Code 200, Recieved: ", response.StatusCode)
	}

	log.Println("PASS: Response Status Code: ", response.StatusCode)
}
