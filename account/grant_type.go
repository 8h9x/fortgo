package account

import "net/url"

type GrantType string

const (
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	GrantTypeClientCredentials GrantType = "client_credentials"
	GrantTypeContinuationToken GrantType = "continuation_token"
	GrantTypeDeviceAuth        GrantType = "device_auth"
	GrantTypeDeviceCode        GrantType = "device_code"
	GrantTypeExchangeCode      GrantType = "exchange_code"
	GrantTypeExternalAuth      GrantType = "external_auth"
	GrantTypeOTP               GrantType = "otp"
	GrantTypePassword          GrantType = "password"
	GrantTypeRefreshToken      GrantType = "refresh_token"
	GrantTypeTokenToToken      GrantType = "token_to_token"
)

type TokenGrant = url.Values

func TokenGrantAuthorizationCode(code string) TokenGrant {
	return TokenGrant{
		"grant_type": {"authorization_code"},
		"code":       {code},
	}
}

func TokenGrantClientCredentials() TokenGrant {
	return TokenGrant{
		"grant_type": {"client_credentials"},
	}
}

func TokenGrantContinuationToken(continuationToken string) TokenGrant {
	return TokenGrant{
		"grant_type":         {"continuation_token"},
		"continuation_token": {continuationToken},
	}
}

func TokenGrantDeviceAuth(accountId, deviceId, secret string) TokenGrant {
	return TokenGrant{
		"grant_type": {"device_auth"},
		"account_id": {accountId},
		"device_id":  {deviceId},
		"secret":     {secret},
	}
}

func TokenGrantDeviceCode(deviceCode string) TokenGrant {
	return TokenGrant{
		"grant_type":  {"device_code"},
		"device_code": {deviceCode},
	}
}

func TokenGrantExchangeCode(code string) TokenGrant {
	return TokenGrant{
		"grant_type":    {"exchange_code"},
		"exchange_code": {code},
	}
}

func TokenGrantExternalAuth(externalAuthToken string) TokenGrant {
	return TokenGrant{
		"grant_type":          {"external_auth"},
		"external_auth_token": {externalAuthToken},
	}
}

func TokenGrantOTP(otp string) TokenGrant {
	return TokenGrant{
		"grant_type": {"otp"},
		"otp":        {otp},
	}
}

func TokenGrantPassword(username, password string) TokenGrant {
	return TokenGrant{
		"grant_type": {"password"},
		"username":   {username},
		"password":   {password},
	}
}

func TokenGrantRefreshToken(refreshToken string) TokenGrant {
	return TokenGrant{
		"grant_type":    {"refresh_token"},
		"refresh_token": {refreshToken},
	}
}

func TokenGrantAccessToken(accessToken string) TokenGrant {
	return TokenGrant{
		"grant_type":   {"token_to_token"},
		"access_token": {accessToken},
	}
}
