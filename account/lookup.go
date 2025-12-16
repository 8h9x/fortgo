package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetUserByID(accountID string) (GetUserResponseExtended, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	resp, err := request.Request(c.HTTPClient, "GET", fmt.Sprintf("%s/account/api/public/account/%s", consts.AccountService, accountID), headers, "")
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	res, err := request.ResponseParser[GetUserResponseExtended](resp)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	return res.Body, err
}

func (c *Client) GetUsersByIDBulk(accountIDs []string) ([]GetUserResponse, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	query := url.Values{}
	for _, accountID := range accountIDs {
		query.Add("accountId", accountID)
	}

	resp, err := request.Request(c.HTTPClient, "GET", fmt.Sprintf("%s/account/api/public/account?%s", consts.AccountService, query.Encode()), headers, "")
	if err != nil {
		return []GetUserResponse{}, err
	}

	res, err := request.ResponseParser[[]GetUserResponse](resp)
	if err != nil {
		return []GetUserResponse{}, err
	}

	return res.Body, err
}

func (c *Client) GetUserByDisplayName(displayName string) (GetUserResponseExtended, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	resp, err := request.Request(c.HTTPClient, "GET", fmt.Sprintf("%s/account/api/public/account/displayName/%s", consts.AccountService, displayName), headers, "")
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	res, err := request.ResponseParser[GetUserResponseExtended](resp)
	if err != nil {
		return GetUserResponseExtended{}, err
	}

	return res.Body, err
}

func (c *Client) GetUserByExternalDisplayName(externalAuthType ExternalAuthType, displayName string) ([]GetUserResponse, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	resp, err := request.Request(c.HTTPClient, "POST", fmt.Sprintf("%s/account/api/public/account/lookup/externalAuth/%s/displayName/%s?caseInsensitive=true", consts.AccountService, externalAuthType, displayName), headers, "")
	if err != nil {
		return []GetUserResponse{}, err
	}

	res, err := request.ResponseParser[[]GetUserResponse](resp)
	if err != nil {
		return []GetUserResponse{}, err
	}

	return res.Body, err
}

func (c *Client) GetUsersByExternalDisplayNameBulk(externalAuthType ExternalAuthType, displayNames []string) (data map[ExternalAuthType]ExternalAuth, err error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	bodyBytes, err := json.Marshal(getUsersByExternalDisplayNameBulkPayload{externalAuthType, displayNames})
	if err != nil {
		return data, err
	}

	resp, err := request.Request(c.HTTPClient, "POST", fmt.Sprintf("%s/account/api/public/account/lookup/externalDisplayName", consts.AccountService), headers, string(bodyBytes))
	if err != nil {
		return data, err
	}

	res, err := request.ResponseParser[map[ExternalAuthType]ExternalAuth](resp)
	if err != nil {
		return data, err
	}

	return res.Body, err
}

func (c *Client) GetUsersByExternalIDBulk(externalAuthType ExternalAuthType, ids []string) (data map[ExternalAuthType]ExternalAuth, err error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	bodyBytes, err := json.Marshal(getUsersByExternalIDBulkPayload{externalAuthType, ids})
	if err != nil {
		return data, err
	}

	resp, err := request.Request(c.HTTPClient, "POST", fmt.Sprintf("%s/account/api/public/account/lookup/externalId", consts.AccountService), headers, string(bodyBytes))
	if err != nil {
		return data, err
	}

	res, err := request.ResponseParser[map[ExternalAuthType]ExternalAuth](resp)
	if err != nil {
		return data, err
	}

	return res.Body, err
}