package base

import (
	"net/http"

	"github.com/8h9x/fortgo/account"
)

type Client struct {
	HTTPClient  *http.Client
	Credentials *account.Token
}

func New(httpClient *http.Client, credentials *account.Token) Client {
	return Client{httpClient, credentials}
}
