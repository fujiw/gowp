package client

import (
	"log"
	"net/http"
	"net/url"

	"github.com/fatih/color"

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

	notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
	notice(c.BaseUrl.String())

	response := &response.Response{}

	notice(request, response)
}