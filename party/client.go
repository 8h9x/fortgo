package party

import (
	"net/http"

	"github.com/8h9x/fortgo/auth"
)

type Client struct {
    HTTPClient  *http.Client
    Credentials *auth.TokenResponse
}

func NewClient(httpClient *http.Client, credentials *auth.TokenResponse) *Client {
	return &Client{httpClient, credentials}
}