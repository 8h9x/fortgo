package account

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) VerifyToken(includePerms bool) (VerifyTokenResponse, error) {
	v := url.Values{}

	if includePerms {
		v.Set("includePerms", "true")
	}

	req, err := rest.NewRequest(
		http.MethodGet,
		consts.AccountService,
		"account/api/oauth/verify",
		rest.WithFormBody(v),
	)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return VerifyTokenResponse{}, err
	}

	var data VerifyTokenResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return VerifyTokenResponse{}, err
	}

	return data, err
}

func (c *Client) CreateDeviceAuth(accountId string) (DeviceAuth, error) {
	req, err := rest.NewRequest(
		http.MethodPost,
		consts.AccountService,
		fmt.Sprintf("account/api/public/account/%s/deviceAuth", accountId),
	)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return DeviceAuth{}, err
	}

	var data DeviceAuth
	if err = rest.DecodeResponse(res, &data); err != nil {
		return DeviceAuth{}, err
	}

	return data, err
}

func (c *Client) GetExchangeCode() (ExchangeResponse, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.AccountService,
		"account/api/oauth/exchange",
	)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ExchangeResponse{}, err
	}

	var data ExchangeResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return ExchangeResponse{}, err
	}

	return data, err
}

func (c *Client) CreateDeviceCode(clientToken Token) (GetDeviceCodeResponse, error) {
	req, err := rest.NewRequest(
		http.MethodPost,
		consts.AccountService,
		"account/api/oauth/deviceAuthorization",
		rest.WithBearerToken(clientToken.AccessToken),
	)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return GetDeviceCodeResponse{}, err
	}

	var data GetDeviceCodeResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return GetDeviceCodeResponse{}, err
	}

	return data, err
}

func (c *Client) GetToken(authClient *AuthClient, tokenGrant TokenGrant, tokenType TokenType) (Token, error) {
	tokenGrant.Set("token_type", string(tokenType))

	req, err := rest.NewRequest(
		http.MethodPost,
		consts.AccountService,
		"account/api/oauth/token",
		rest.WithBasicToken(authClient.Base64()),
		rest.WithFormBody(tokenGrant),
	)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return Token{}, err
	}

	var data Token
	if err = rest.DecodeResponse(res, &data); err != nil {
		return Token{}, err
	}

	return data, err
}
