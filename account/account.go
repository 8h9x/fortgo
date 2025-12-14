package account

import (
	"fmt"
	"net/http"
	"net/url"

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

func (c *Client) FetchUserByID(accountID string) (FetchUserResponseExtended, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	resp, err := request.Request(c.HTTPClient, "GET", fmt.Sprintf("%s/account/api/public/account/%s", consts.AccountService, accountID), headers, "")
	if err != nil {
		return FetchUserResponseExtended{}, err
	}

	res, err := request.ResponseParser[FetchUserResponseExtended](resp)
	if err != nil {
		return FetchUserResponseExtended{}, err
	}

	return res.Body, err
}

func (c *Client) FetchUsersByIDBulk(accountIDs []string) ([]FetchUserResponse, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	query := url.Values{}
	for _, accountID := range accountIDs {
		query.Add("accountId", accountID)
	}

	resp, err := request.Request(c.HTTPClient, "GET", fmt.Sprintf("%s/account/api/public/account?%s", consts.AccountService, query.Encode()), headers, "")
	if err != nil {
		return []FetchUserResponse{}, err
	}

	res, err := request.ResponseParser[[]FetchUserResponse](resp)
	if err != nil {
		return []FetchUserResponse{}, err
	}

	return res.Body, err
}

func (c *Client) FetchUserByDisplayName(displayName string) (FetchUserResponseExtended, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	resp, err := request.Request(c.HTTPClient, "GET", fmt.Sprintf("%s/account/api/public/account/displayName/%s", consts.AccountService, displayName), headers, "")
	if err != nil {
		return FetchUserResponseExtended{}, err
	}

	res, err := request.ResponseParser[FetchUserResponseExtended](resp)
	if err != nil {
		return FetchUserResponseExtended{}, err
	}

	return res.Body, err
}