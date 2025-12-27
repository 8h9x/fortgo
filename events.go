package fortgo

import "github.com/8h9x/fortgo/auth"

type OnTokenRefresh func(credentials auth.TokenResponse)
