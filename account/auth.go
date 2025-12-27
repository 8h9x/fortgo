package account

type TokenType string

const (
	TokenTypeDefault TokenType = "*"
	TokenTypeEG1     TokenType = "eg1"
	TokenTypeEP1     TokenType = "ep1"
	TokenTypeEpicID  TokenType = "epic_id"
	TokenTypeIDToken TokenType = "id_token"
)

type GrantType string

const (
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	GrantTypeClientCredentials GrantType = "client_credentials"
	GrantTypeContinuationToken GrantType = "continuation_token"
	GrantTypeDeviceAuth        GrantType = "device_auth"
	GrantTypeDeviceCode        GrantType = "device_code"
	GrantTypeExchangeCode      GrantType = "exchange_code"
	GrantTypeExternalAuth      GrantType = "exteral_auth"
	GrantTypeOTP               GrantType = "otp"
	GrantTypePassword          GrantType = "password"
	GrantTypeRefreshToken      GrantType = "refresh_token"
	GrantTypeTokenToToken      GrantType = "token_to_token"
)

type ClientCredentials struct {
}

type Credentials struct {
}
