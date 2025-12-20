package auth

import (
	"fmt"
	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
	"net/http"
	"net/url"
	"time"
)

type PayloadAuthorizationCode struct {
	Code string `json:"code"`
}

type PayloadClientCredentials struct{}

type PayloadContinuationToken struct {
	ContinuationToken string `json:"continuation_token"`
}

type PayloadDeviceAuth struct {
	AccountID string `json:"account_id"`
	DeviceID  string `json:"device_id"`
	Secret    string `json:"secret"`
}

type PayloadDeviceCode struct {
	DeviceCode string `json:"device_code"`
}

type PayloadExchangeCode struct {
	ExchangeCode string `json:"exchange_code"`
}

type PayloadExternalAuth struct {
	ExternalAuthToken string `json:"external_auth_token"`
}

type PayloadOTP struct {
	OTP string `json:"otp"`
}

type PayloadPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PayloadRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type PayloadTokenToToken struct {
	AccessToken string `json:"access_token"`
}

type Payload interface {
	PayloadAuthorizationCode | PayloadClientCredentials | PayloadContinuationToken |
		PayloadDeviceAuth | PayloadDeviceCode | PayloadExchangeCode | PayloadExternalAuth |
		PayloadOTP | PayloadPassword | PayloadRefreshToken | PayloadTokenToToken
}

func Authenticate[T Payload](httpClient *http.Client, client *AuthClient, payload T, eg1 bool) (TokenResponse, error) {
	v := url.Values{}
	if eg1 == true {
		v.Set("token_type", "eg1")
	}

	switch p := any(payload).(type) {
	case PayloadAuthorizationCode:
		v.Set("grant_type", "authorization_code")
		v.Set("code", p.Code)
		break
	case PayloadClientCredentials:
		v.Set("grant_type", "client_credentials")
		break
	case PayloadContinuationToken:
		v.Set("grant_type", "continuation_token")
		v.Set("continuation_token", p.ContinuationToken)
		break
	case PayloadDeviceAuth:
		v.Set("grant_type", "device_auth")
		v.Set("account_id", p.AccountID)
		v.Set("device_id", p.DeviceID)
		v.Set("secret", p.Secret)
		break
	case PayloadDeviceCode:
		v.Set("grant_type", "device_code")
		v.Set("device_code", p.DeviceCode)
		break
	case PayloadExchangeCode:
		v.Set("grant_type", "exchange_code")
		v.Set("exchange_code", p.ExchangeCode)
		break
	case PayloadExternalAuth:
		v.Set("grant_type", "external_auth")
		v.Set("external_auth_token", p.ExternalAuthToken)
		break
	case PayloadOTP:
		v.Set("grant_type", "otp")
		v.Set("otp", p.OTP)
		break
	case PayloadPassword:
		v.Set("grant_type", "password")
		v.Set("username", p.Username)
		v.Set("password", p.Password)
		break
	case PayloadRefreshToken:
		v.Set("grant_type", "refresh_token")
		v.Set("refresh_token", p.RefreshToken)
		break
	case PayloadTokenToToken:
		v.Set("grant_type", "token_to_token")
		v.Set("access_token", p.AccessToken)
		break
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

type TokenResponse struct {
	AccessToken      string    `json:"access_token"`
	ExpiresIn        int       `json:"expires_in"`
	ExpiresAt        time.Time `json:"expires_at"`
	TokenType        string    `json:"token_type"`
	RefreshToken     string    `json:"refresh_token"`
	RefreshExpires   int       `json:"refresh_expires"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
	AccountID        string    `json:"account_id"`
	ClientID         string    `json:"client_id"`
	InternalClient   bool      `json:"internal_client"`
	ClientService    string    `json:"client_service"`
	DisplayName      string    `json:"displayName"`
	App              string    `json:"app"`
	InAppID          string    `json:"in_app_id"`
	DeviceID         string    `json:"device_id"`
	ProductID        string    `json:"product_id"`
	ApplicationID    string    `json:"application_id"`
	ACR              string    `json:"acr"`
	AuthTime         time.Time `json:"auth_time"`
}

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

type VerifyTokenResponsePermission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

type VerifyTokenResponse struct {
	Token          string                          `json:"token"`
	SessionID      string                          `json:"session_id"`
	TokenType      string                          `json:"token_type"`
	ClientID       string                          `json:"client_id"`
	InternalClient bool                            `json:"internal_client"`
	ClientService  string                          `json:"client_service"`
	AccountID      string                          `json:"account_id"`
	ExpiresIn      int                             `json:"expires_in"`
	ExpiresAt      time.Time                       `json:"expires_at"`
	AuthMethod     string                          `json:"auth_method"`
	DisplayName    string                          `json:"display_name"`
	App            string                          `json:"app"`
	InAppID        string                          `json:"in_app_id"`
	DeviceID       string                          `json:"device_id"`
	Scope          []string                        `json:"scope"`
	ProductID      string                          `json:"product_id"`
	SandboxID      string                          `json:"sandbox_id"`
	DeploymentID   string                          `json:"deployment_id"`
	ApplicationID  string                          `json:"application_id"`
	ACR            string                          `json:"acr"`
	AuthTime       time.Time                       `json:"auth_time"`
	Perms          []VerifyTokenResponsePermission `json:"perms"`
}

func CreateDeviceAuth(httpClient *http.Client, credentials TokenResponse) (DeviceAuthResponse, error) {
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

type DeviceAuthResponse struct {
	AccountID string `json:"accountId"`
	Created   struct {
		DateTime  time.Time `json:"dateTime"`
		IpAddress string    `json:"ipAddress"`
		Location  string    `json:"location"`
	} `json:"created"`
	DeviceID  string `json:"deviceId"`
	Secret    string `json:"secret"`
	UserAgent string `json:"userAgent"`
}

func GetExchangeCode(httpClient *http.Client, credentials TokenResponse) (ExchangeResponse, error) {
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

type ExchangeResponse struct {
	Code             string `json:"code"`
	CreatingClientID string `json:"creatingClientId"`
	ExpiresInSeconds int    `json:"expiresInSeconds"`
}

func CreateDeviceCode(httpClient *http.Client, credentials TokenResponse) (GetDeviceCodeResponse, error) {
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

type GetDeviceCodeResponse struct {
	UserCode                string `json:"user_code"`
	DeviceCode              string `json:"device_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	Prompt                  string `json:"prompt"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
	ClientID                string `json:"client_id"`
}
