package account

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetPublicKey(thumbprint string) (GetPublicKeyResponse, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.AccountService,
		fmt.Sprintf("account/api/publickeys/%s", thumbprint),
	)
	if err != nil {
		return GetPublicKeyResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return GetPublicKeyResponse{}, err
	}

	resp, err := request.ParseResponse[GetPublicKeyResponse](res)
	if err != nil {
		return GetPublicKeyResponse{}, err
	}

	return resp.Data, nil
}