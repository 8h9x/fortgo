package account

import (
	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetPublicKey(thumbprint string) (GetPublicKeyResponse, error) {
	resp, err := request.Request(c.HTTPClient, "GET", consts.AccountService+"/account/api/publickeys/"+thumbprint, nil, "")
	if err != nil {
		return GetPublicKeyResponse{}, err
	}

	res, err := request.ResponseParser[GetPublicKeyResponse](resp)
	if err != nil {
		return GetPublicKeyResponse{}, err
	}

	return res.Body, err
}