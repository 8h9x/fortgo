package usersearch

import (
	"net/http"
)

type Client struct {
	HTTPClient *http.Client
}

func New(httpClient *http.Client) *Client {
	return &Client{httpClient}
}
