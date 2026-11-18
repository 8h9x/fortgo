package account

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) GetPublicKey(thumbprint string) (GetPublicKeyResponse, error) {
	req, err := rest.NewRequest(
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

	var data GetPublicKeyResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return GetPublicKeyResponse{}, err
	}

	return data, err
}
