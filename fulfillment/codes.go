package fulfillment

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) RedeemCodeForAccount(accountID string, code string) (RedeemCodeResponse, error) {
	req, err := request.New(
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

	resp, err := request.DecodeResponse[RedeemCodeResponse](res)
	if err != nil {
		return RedeemCodeResponse{}, err
	}

	return resp.Data, nil
}

func (c *Client) RedeemCodeForIdentity(identityID string, code string) (RedeemCodeResponse, error) {
	req, err := request.New(
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

	resp, err := request.DecodeResponse[RedeemCodeResponse](res)
	if err != nil {
		return RedeemCodeResponse{}, err
	}

	return resp.Data, nil
}
