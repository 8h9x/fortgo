package eos

import (
	"net/http"
	"strings"

	"github.com/8h9x/fortgo/account"
	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetExchangeCode() (ExchangeResponse, error) {
	req, err := request.New(
		http.MethodGet,
		consts.EOSServiceAuth,
		"exchange",
	)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ExchangeResponse{}, err
	}

	resp, err := request.DecodeResponse[ExchangeResponse](res)
	if err != nil {
		return ExchangeResponse{}, err
	}

	return resp.Data, err
}

func (c *Client) CreateDeviceCode(clientToken Token) (GetDeviceCodeResponse, error) {
	req, err := request.New(
		http.MethodPost,
		consts.EOSServiceAuth,
		"deviceAuthorization",
		request.WithBearerToken(clientToken.AccessToken),
	)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return GetDeviceCodeResponse{}, err
	}

	resp, err := request.DecodeResponse[GetDeviceCodeResponse](res)
	if err != nil {
		return GetDeviceCodeResponse{}, err
	}

	return resp.Data, err
}

func (c *Client) GetToken(authClient *account.AuthClient, tokenGrant account.TokenGrant, tokenType account.TokenType, scope ...string) (Token, error) {
	tokenGrant.Set("token_type", string(tokenType))
	if len(scope) > 0 {
		tokenGrant.Set("scope", strings.Join(scope, " "))
	}

	req, err := request.New(
		http.MethodPost,
		consts.EOSServiceAuth,
		"token",
		request.WithBasicToken(authClient.Base64()),
		request.WithFormBody(tokenGrant),
	)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return Token{}, err
	}

	resp, err := request.DecodeResponse[Token](res)
	if err != nil {
		return Token{}, err
	}

	return resp.Data, err
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
