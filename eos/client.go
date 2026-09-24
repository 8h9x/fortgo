package eos

import (
	"net/http"

	"github.com/8h9x/fortgo/account"
)

type Client struct {
	HTTPClient  *http.Client
	Credentials *account.TokenResponse
}

func NewClient(httpClient *http.Client, credentials *account.TokenResponse) *Client {
	return &Client{httpClient, credentials}
}
