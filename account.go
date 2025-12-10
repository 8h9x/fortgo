package vinderman

import (
	"github.com/8h9x/vinderman/auth"
	"github.com/8h9x/vinderman/consts"
	"github.com/8h9x/vinderman/request"
)

func (c *Client) GetExchangeCode() (auth.ExchangeResponse, error) {
	return auth.GetExchangeCode(c.HttpClient, c.CredentialsMap[c.ClientID])
}

func (c *Client) CreateDeviceAuth() (auth.DeviceAuthResponse, error) {
	return auth.CreateDeviceAuth(c.HttpClient, c.CredentialsMap[c.ClientID])
}

type GetPublicKeyResponse struct {
	KTY string `json:"kty"`
	E   string `json:"e"`
	KID string `json:"kid"`
	N   string `json:"n"`
}

func (c *Client) GetPublicKey(thumbprint string) (GetPublicKeyResponse, error) {
	resp, err := c.Request("GET", consts.AccountService+"/account/api/publickeys/"+thumbprint, nil, "")
	if err != nil {
		return GetPublicKeyResponse{}, err
	}

	res, err := request.ResponseParser[GetPublicKeyResponse](resp)
	if err != nil {
		return GetPublicKeyResponse{}, err
	}

	return res.Body, err
}
