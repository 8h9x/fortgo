package fortgo

import (
	"github.com/8h9x/fortgo/request"
	"net/http"
)

func (c *Client) Request(method string, url string, header http.Header, body string) (*http.Response, error) {
	for key, value := range c.Header {
		header.Set(key, value[0])
	}

	return request.Request(c.HttpClient, method, url, header, body)
}