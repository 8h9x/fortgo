package account

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) GetUserByID(accountID string) (GetUserResponseExtended, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.AccountService,
		fmt.Sprintf("account/api/public/account/%s", accountID),
	)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	var data GetUserResponseExtended
	if err = rest.DecodeResponse(res, &data); err != nil {
		return GetUserResponseExtended{}, err
	}

	return data, err
}

func (c *Client) GetUsersByIDBulk(accountIDs []string) ([]GetUserResponse, error) {
	query := url.Values{}
	for _, accountID := range accountIDs {
		query.Add("accountId", accountID)
	}

	req, err := rest.NewRequest(
		http.MethodGet,
		consts.AccountService,
		"account/api/public/account",
		rest.WithQueryParamaters(query),
	)
	if err != nil {
		return []GetUserResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []GetUserResponse{}, err
	}

	var data []GetUserResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return []GetUserResponse{}, err
	}

	return data, err
}

func (c *Client) GetUserByDisplayName(displayName string) (GetUserResponseExtended, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.AccountService,
		fmt.Sprintf("account/api/public/account/displayName/%s", displayName),
	)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	var data GetUserResponseExtended
	if err = rest.DecodeResponse(res, &data); err != nil {
		return GetUserResponseExtended{}, err
	}

	return data, err
}

func (c *Client) GetUserByExternalDisplayName(externalAuthType ExternalAuthType, displayName string, caseSensitive bool) ([]GetUserResponse, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.AccountService,
		fmt.Sprintf("account/api/public/account/lookup/externalAuth/%s/displayName/%s?caseInsensitive=%t", externalAuthType, displayName, !caseSensitive),
	)
	if err != nil {
		return []GetUserResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []GetUserResponse{}, err
	}

	var data []GetUserResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return []GetUserResponse{}, err
	}

	return data, err
}

func (c *Client) GetUsersByExternalDisplayNameBulk(externalAuthType ExternalAuthType, displayNames []string) (ExternalAuthGroup, error) {
	req, err := rest.NewRequest(
		http.MethodPost,
		consts.AccountService,
		"account/api/public/account/lookup/externalDisplayName",
		rest.WithJSONBody(&getUsersByExternalDisplayNameBulkPayload{externalAuthType, displayNames}),
	)
	if err != nil {
		return ExternalAuthGroup{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ExternalAuthGroup{}, err
	}

	var data ExternalAuthGroup
	if err = rest.DecodeResponse(res, &data); err != nil {
		return ExternalAuthGroup{}, err
	}

	return data, err
}

func (c *Client) GetUsersByExternalIDBulk(externalAuthType ExternalAuthType, ids []string) (ExternalAuthGroup, error) {
	req, err := rest.NewRequest(
		http.MethodPost,
		consts.AccountService,
		"account/api/public/account/lookup/externalId",
		rest.WithJSONBody(&getUsersByExternalIDBulkPayload{externalAuthType, ids}),
	)
	if err != nil {
		return ExternalAuthGroup{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ExternalAuthGroup{}, err
	}

	var data ExternalAuthGroup
	if err = rest.DecodeResponse(res, &data); err != nil {
		return ExternalAuthGroup{}, err
	}

	return data, err
}
