package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/fatih/color"
)

type ApiClient struct {
	BaseUrl *url.URL
	HTTPClient *http.Client
	Logger *log.Logger
}

func (c *ApiClient) Run() {
	notice := color.New(color.Bold, color.FgGreen).PrintlnFunc()
	notice(c.BaseUrl.String())

	req, err := http.NewRequest(http.MethodGet, c.BaseUrl.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	notice(resp.StatusCode)

	body, _ := ioutil.ReadAll(resp.Body)
	notice(string(body))
}