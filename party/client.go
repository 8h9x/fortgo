package party

import (
	"net/http"

	"github.com/8h9x/fortgo/account"
	"github.com/8h9x/fortgo/base"
)

type Client struct {
	base.Client
}

func NewClient(httpClient *http.Client, credentials *account.TokenResponse) *Client {
	return &Client{base.New(httpClient, credentials)}
}
