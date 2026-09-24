package fortgo

import "github.com/8h9x/fortgo/account"

type OnTokenRefresh func(credentials account.TokenResponse)
