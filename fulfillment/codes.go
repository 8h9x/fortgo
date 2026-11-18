package fulfillment

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) RedeemCodeForAccount(accountID string, code string) (RedeemCodeResponse, error) {
	req, err := rest.NewRequest(
		http.MethodPost,
		consts.FulfillmentService,
		fmt.Sprintf("fulfillment/api/public/accounts/%s/codes/%s", accountID, code),
	)
	if err != nil {
		return RedeemCodeResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return RedeemCodeResponse{}, err
	}

	var data RedeemCodeResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return RedeemCodeResponse{}, err
	}

	return data, err
}

func (c *Client) RedeemCodeForIdentity(identityID string, code string) (RedeemCodeResponse, error) {
	req, err := rest.NewRequest(
		http.MethodPost,
		consts.FulfillmentService,
		fmt.Sprintf("fulfillment/api/public/identities/%s/codes/%s", identityID, code),
	)
	if err != nil {
		return RedeemCodeResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return RedeemCodeResponse{}, err
	}

	var data RedeemCodeResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return RedeemCodeResponse{}, err
	}

	return data, err
}
