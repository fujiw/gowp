package main

import (
	"gowp/client"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	logger := log.New(os.Stderr, "[LOG]", log.LstdFlags)
	baseURL, err := url.Parse("http://checkip.amazonaws.com/")

	if err != nil {
		panic("error")
	}

	c := client.ApiClient{
		BaseUrl: baseURL,
		HTTPClient: http.DefaultClient,
		Logger: logger,
	}

	c.Run()
}