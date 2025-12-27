package account

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetUserByID(accountID string) (GetUserResponseExtended, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.AccountService,
		fmt.Sprintf("account/api/public/account/%s", accountID),
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	resp, err := request.ParseResponse[GetUserResponseExtended](res)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetUsersByIDBulk(accountIDs []string) ([]GetUserResponse, error) {
	query := url.Values{}
	for _, accountID := range accountIDs {
		query.Add("accountId", accountID)
	}

	req, err := request.MakeRequest(
		http.MethodGet,
		consts.AccountService,
		fmt.Sprintf("account/api/public/account"),
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithQueryParamaters(query),
	)
	if err != nil {
		return []GetUserResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []GetUserResponse{}, err
	}

	resp, err := request.ParseResponse[[]GetUserResponse](res)
	if err != nil {
		return []GetUserResponse{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetUserByDisplayName(displayName string) (GetUserResponseExtended, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.AccountService,
		fmt.Sprintf("account/api/public/account/displayName/%s", displayName),
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	resp, err := request.ParseResponse[GetUserResponseExtended](res)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetUserByExternalDisplayName(externalAuthType ExternalAuthType, displayName string, caseSensitive bool) ([]GetUserResponse, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.AccountService,
		fmt.Sprintf("account/api/public/account/lookup/externalAuth/%s/displayName/%s?caseInsensitive=%t", externalAuthType, displayName, !caseSensitive),
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return []GetUserResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []GetUserResponse{}, err
	}

	resp, err := request.ParseResponse[[]GetUserResponse](res)
	if err != nil {
		return []GetUserResponse{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetUsersByExternalDisplayNameBulk(externalAuthType ExternalAuthType, displayNames []string) (data map[ExternalAuthType]ExternalAuth, err error) {
	req, err := request.MakeRequest(
		http.MethodPost,
		consts.AccountService,
		"account/api/public/account/lookup/externalDisplayName",
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithJSONBody(&getUsersByExternalDisplayNameBulkPayload{externalAuthType, displayNames}),
	)
	if err != nil {
		return data, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return data, err
	}

	resp, err := request.ParseResponse[map[ExternalAuthType]ExternalAuth](res)
	if err != nil {
		return data, err
	}

	return resp.Data, err
}

func (c *Client) GetUsersByExternalIDBulk(externalAuthType ExternalAuthType, ids []string) (data map[ExternalAuthType]ExternalAuth, err error) {
	req, err := request.MakeRequest(
		http.MethodPost,
		consts.AccountService,
		"account/api/public/account/lookup/externalId",
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithJSONBody(&getUsersByExternalIDBulkPayload{externalAuthType, ids}),
	)
	if err != nil {
		return data, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return data, err
	}

	resp, err := request.ParseResponse[map[ExternalAuthType]ExternalAuth](res)
	if err != nil {
		return data, err
	}

	return resp.Data, err
}
