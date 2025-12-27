package eos

import (
	"net/http"
	"time"

	"github.com/8h9x/fortgo/auth"
)

type ClientCredentials struct {
	AccessToken   string    `json:"access_token"`
	ApplicationID string    `json:"application_id"`
	ClientID      string    `json:"client_id"`
	ExpiresAt     time.Time `json:"expires_at"`
	ExpiresIn     int       `json:"expires_in"`
	TokenType     string    `json:"token_type"`
}

type DeviceAuthorization struct {
	ClientID                string `json:"client_id"`
	DeviceCode              string `json:"device_code"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
	Prompt                  string `json:"prompt"`
	UserCode                string `json:"user_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
}

type UserCredentials struct {
	AccessToken      string    `json:"access_token"`
	AccountID        string    `json:"account_id"`
	ApplicationID    string    `json:"application_id"`
	ClientID         string    `json:"client_id"`
	ExpiresAt        time.Time `json:"expires_at"`
	ExpiresIn        int       `json:"expires_in"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at"`
	RefreshExpiresIn int       `json:"refresh_expires_in"`
	RefreshToken     string    `json:"refresh_token"`
	Scope            string    `json:"scope"`
	TokenType        string    `json:"token_type"`
}

type Exchange struct {
	Code             string `json:"code"`
	CreatingClientID string `json:"creatingClientId"`
	ExpiresInSeconds int    `json:"expiresInSeconds"`
}

type AuthPayload auth.Payload

func Authenticate[T AuthPayload](httpClient *http.Client, clientId, clientSecret string, payload T, eg1 bool) (T, error) {
	return payload, nil
}

//func (c Client) DeviceCodeLogin(clientId string, clientSecret string, deviceCode string) (credentials UserCredentials, err error) {
//	encodedClientToken := Base64Encode(clientId + ":" + clientSecret)
//
//	headers := http.Header{}
//	headers.Set("Content-Type", "application/x-www-form-urlencoded")
//	headers.Set("Authorization", fmt.Sprint("Basic ", encodedClientToken))
//
//	v := url.Values{}
//	v.Set("grant_type", "device_code")
//	v.Set("device_code", deviceCode)
//	body := v.Encode()
//
//	resp, err := c.Request("POST", consts.EOS_AUTH+"/token", headers, body)
//	if err != nil {
//		return
//	}
//
//	res, err := request.ResponseParser[UserCredentials](resp)
//
//	return res.Body, err
//}

//func (c Client) ExchangeCodeLogin(clientId string, clientSecret string, code string) (credentials UserCredentials, err error) {
//	encodedClientToken := Base64Encode(clientId + ":" + clientSecret)
//
//	headers := http.Header{}
//	headers.Set("Content-Type", "application/x-www-form-urlencoded")
//	headers.Set("Authorization", fmt.Sprint("Basic ", encodedClientToken))
//
//	v := url.Values{}
//	v.Set("grant_type", "exchange_code")
//	v.Set("exchange_code", code)
//	v.Set("scope", "offline_access")
//	body := v.Encode()
//
//	resp, err := c.Request("POST", consts.EOS_AUTH+"/token", headers, body)
//	if err != nil {
//		return
//	}
//
//	res, err := request.ResponseParser[UserCredentials](resp)
//
//	return res.Body, err
//}
//
//func (c Client) GetClientCredentials(clientId string, clientSecret string) (credentials ClientCredentials, err error) {
//	encodedClientToken := Base64Encode(clientId + ":" + clientSecret)
//
//	headers := http.Header{}
//	headers.Set("Content-Type", "application/x-www-form-urlencoded")
//	headers.Set("Authorization", fmt.Sprint("Basic ", encodedClientToken))
//
//	v := url.Values{}
//	v.Set("grant_type", "client_credentials")
//	body := v.Encode()
//
//	resp, err := c.Request("POST", consts.EOS_AUTH+"/token", headers, body)
//	if err != nil {
//		return
//	}
//
//	res, err := request.ResponseParser[ClientCredentials](resp)
//
//	return res.Body, err
//}
//
//func (c Client) GetExchangeCode(credentials UserCredentials) (exchange Exchange, err error) {
//	headers := http.Header{}
//	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))
//
//	resp, err := c.Request("GET", consts.EOS_AUTH+"/exchange", headers, "")
//	if err != nil {
//		return
//	}
//
//	res, err := request.ResponseParser[Exchange](resp)
//
//	return res.Body, err
//}
//
//func (c Client) GetDeviceCode(credentials ClientCredentials) (deviceAuth DeviceAuthorization, err error) {
//	headers := http.Header{}
//	headers.Set("Content-Type", "application/x-www-form-urlencoded")
//	headers.Set("Authorization", fmt.Sprint("Bearer ", credentials.AccessToken))
//
//	v := url.Values{}
//	v.Set("prompt", "login")
//	body := v.Encode()
//
//	resp, err := c.Request("POST", consts.EOS_AUTH+"/deviceAuthorization", headers, body)
//	if err != nil {
//		return
//	}
//
//	res, err := request.ResponseParser[DeviceAuthorization](resp)
//
//	return res.Body, err
//}
//
//func (c Client) WaitForDeviceCodeAccept(clientId string, clientSecret string, deviceCode string) (credentials UserCredentials, err error) {
//	credentials, err = c.DeviceCodeLogin(clientId, clientSecret, deviceCode)
//
//	if err != nil {
//		if err.(*request.Error[EpicErrorResponse]).Raw.ErrorCode == consts.ErrorAuthorizationPending {
//			time.Sleep(10 * time.Second)
//			return c.WaitForDeviceCodeAccept(clientId, clientSecret, deviceCode)
//		}
//
//		return
//	}
//
//	return
//}
//
//func Base64Encode(s string) string {
//	return base64.StdEncoding.EncodeToString([]byte(s))
//}
