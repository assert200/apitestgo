package main

import (
	"log"
	"net/url"

	"github.com/assert200/gorest"
)

func main() {
	request := gorest.NewRestRequest()

	var url url.URL
	url.Scheme = "https"
	url.Host = "www.google.com"
	url.Path = "/"

	request.Method = "GET"
	request.URL = url

	test := gorest.RestTest{}
	test.Description = "google"
	test.RestRequest = request
	test.ExpectedStatusCode = 200

	doneTest := gorest.DoAndVerify(test)

	log.Println(doneTest.Description, doneTest.RestResponse.StatusCode)
}
