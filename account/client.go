package account

import (
	"net/http"
)

type Client struct {
	HTTPClient  *http.Client
	Credentials *TokenResponse
}

func NewClient(httpClient *http.Client, credentials *TokenResponse) *Client {
	return &Client{httpClient, credentials}
}
