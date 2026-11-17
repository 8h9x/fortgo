package account

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) VerifyToken(includePerms bool) (VerifyTokenResponse, error) {
	v := url.Values{}

	if includePerms {
		v.Set("includePerms", "true")
	}

	req, err := request.New(
		http.MethodGet,
		consts.AccountService,
		"account/api/oauth/verify",
		request.WithFormBody(v),
	)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return VerifyTokenResponse{}, err
	}

	resp, err := request.DecodeResponse[VerifyTokenResponse](res)
	if err != nil {
		return VerifyTokenResponse{}, err
	}

	return resp.Data, err
}

func (c *Client) CreateDeviceAuth(accountId string) (DeviceAuth, error) {
	req, err := request.New(
		http.MethodPost,
		consts.AccountService,
		fmt.Sprintf("account/api/public/account/%s/deviceAuth", accountId),
	)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return DeviceAuth{}, err
	}

	resp, err := request.DecodeResponse[DeviceAuth](res)
	if err != nil {
		return DeviceAuth{}, err
	}

	return resp.Data, err
}

func (c *Client) GetExchangeCode() (ExchangeResponse, error) {
	req, err := request.New(
		http.MethodGet,
		consts.AccountService,
		"account/api/oauth/exchange",
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
		consts.AccountService,
		"account/api/oauth/deviceAuthorization",
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

func (c *Client) GetToken(authClient *AuthClient, tokenGrant TokenGrant, tokenType TokenType) (Token, error) {
	tokenGrant.Set("token_type", string(tokenType))

	req, err := request.New(
		http.MethodPost,
		consts.AccountService,
		"account/api/oauth/token",
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
