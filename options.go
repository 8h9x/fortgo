package fortgo

import (
	"net/http"

	"github.com/8h9x/fortgo/account"
)

type Option func(*config)

func WithHttpClient(httpClient *http.Client) Option {
	return func(cfg *config) {
		cfg.httpClient = httpClient
	}
}

func WithAuthClient(authClient *account.AuthClient) Option {
	return func(cfg *config) {
		cfg.authClient = authClient
	}
}

func WithTokenGrant(tokenGrant account.TokenGrant) Option {
	return func(cfg *config) {
		cfg.tokenGrant = tokenGrant
	}
}
