package account

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func VerifyToken(httpClient *http.Client, accessToken string, includePerms bool) (VerifyTokenResponse, error) {
	v := url.Values{}

	if includePerms {
		v.Set("includePerms", "true")
	}

	req, err := request.MakeRequest(
		http.MethodGet,
		consts.AccountService,
		"account/api/oauth/verify",
		request.WithBearerToken(accessToken),
		request.WithFormBody(v),
	)

	res, err := httpClient.Do(req)
	if err != nil {
		return VerifyTokenResponse{}, err
	}

	resp, err := request.ParseResponse[VerifyTokenResponse](res)
	if err != nil {
		return VerifyTokenResponse{}, err
	}

	return resp.Data, err
}

func CreateDeviceAuth(httpClient *http.Client, credentials *TokenResponse) (DeviceAuthResponse, error) {
	req, err := request.MakeRequest(
		http.MethodPost,
		consts.AccountService,
		fmt.Sprintf("account/api/public/account/%s/deviceAuth", credentials.AccountID),
		request.WithBearerToken(credentials.AccessToken),
	)

	res, err := httpClient.Do(req)
	if err != nil {
		return DeviceAuthResponse{}, err
	}

	resp, err := request.ParseResponse[DeviceAuthResponse](res)
	if err != nil {
		return DeviceAuthResponse{}, err
	}

	return resp.Data, err
}

func GetExchangeCode(httpClient *http.Client, credentials *TokenResponse) (ExchangeResponse, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.AccountService,
		"account/api/oauth/exchange",
		request.WithBearerToken(credentials.AccessToken),
	)

	res, err := httpClient.Do(req)
	if err != nil {
		return ExchangeResponse{}, err
	}

	resp, err := request.ParseResponse[ExchangeResponse](res)
	if err != nil {
		return ExchangeResponse{}, err
	}

	return resp.Data, err
}

func CreateDeviceCode(httpClient *http.Client, credentials *TokenResponse) (GetDeviceCodeResponse, error) {
	req, err := request.MakeRequest(
		http.MethodPost,
		consts.AccountService,
		"account/api/oauth/deviceAuthorization",
		request.WithBearerToken(credentials.AccessToken),
	)

	res, err := httpClient.Do(req)
	if err != nil {
		return GetDeviceCodeResponse{}, err
	}

	resp, err := request.ParseResponse[GetDeviceCodeResponse](res)
	if err != nil {
		return GetDeviceCodeResponse{}, err
	}

	return resp.Data, err
}

func Authenticate[T TokenPayload](httpClient *http.Client, client *AuthClient, payload T, eg1 bool) (TokenResponse, error) {
	v := url.Values{}
	if eg1 == true {
		v.Set("token_type", "eg1")
	}

	switch p := any(payload).(type) {
	case TokenPayloadAuthorizationCode:
		v.Set("grant_type", "authorization_code")
		v.Set("code", p.Code)
	case TokenPayloadClientCredentials:
		v.Set("grant_type", "client_credentials")
	case TokenPayloadContinuationToken:
		v.Set("grant_type", "continuation_token")
		v.Set("continuation_token", p.ContinuationToken)
	case TokenPayloadDeviceAuth:
		v.Set("grant_type", "device_auth")
		v.Set("account_id", p.AccountID)
		v.Set("device_id", p.DeviceID)
		v.Set("secret", p.Secret)
	case TokenPayloadDeviceCode:
		v.Set("grant_type", "device_code")
		v.Set("device_code", p.DeviceCode)
	case TokenPayloadExchangeCode:
		v.Set("grant_type", "exchange_code")
		v.Set("exchange_code", p.ExchangeCode)
	case TokenPayloadExternalAuth:
		v.Set("grant_type", "external_auth")
		v.Set("external_auth_token", p.ExternalAuthToken)
	case TokenPayloadOTP:
		v.Set("grant_type", "otp")
		v.Set("otp", p.OTP)
	case TokenPayloadPassword:
		v.Set("grant_type", "password")
		v.Set("username", p.Username)
		v.Set("password", p.Password)
	case TokenPayloadRefreshToken:
		v.Set("grant_type", "refresh_token")
		v.Set("refresh_token", p.RefreshToken)
	case TokenPayloadTokenToToken:
		v.Set("grant_type", "token_to_token")
		v.Set("access_token", p.AccessToken)
	}

	req, err := request.MakeRequest(
		http.MethodPost,
		consts.AccountService,
		"account/api/oauth/token",
		request.WithBasicToken(client.Base64()),
		request.WithFormBody(v),
	)

	res, err := httpClient.Do(req)
	if err != nil {
		return TokenResponse{}, err
	}

	resp, err := request.ParseResponse[TokenResponse](res)
	if err != nil {
		return TokenResponse{}, err
	}

	return resp.Data, err
}
