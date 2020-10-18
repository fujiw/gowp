package client

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"gowp/request"
	"gowp/response"
)

type ApiClient struct {
	BaseUrl *url.URL
	HTTPClient *http.Client
	Logger *log.Logger
}

func (c *ApiClient) Run() {
	request := &request.Request{}

	fmt.Println(c.BaseUrl.String())

	response := &response.Response{}

	fmt.Println(request, response)
}