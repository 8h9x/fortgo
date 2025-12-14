package account

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

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

func (c *Client) FetchUserByExternalDisplayName(externalAuthType ExternalAuthType, displayName string) ([]FetchUserResponse, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	resp, err := request.Request(c.HTTPClient, "POST", fmt.Sprintf("%s/account/api/public/account/lookup/externalAuth/%s/displayName/%s?caseInsensitive=true", consts.AccountService, externalAuthType, displayName), headers, "")
	if err != nil {
		return []FetchUserResponse{}, err
	}

	res, err := request.ResponseParser[[]FetchUserResponse](resp)
	if err != nil {
		return []FetchUserResponse{}, err
	}

	return res.Body, err
}

func (c *Client) FetchUsersByExternalDisplayNameBulk(externalAuthType ExternalAuthType, displayNames []string) (data map[ExternalAuthType]ExternalAuth, err error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	bodyBytes, err := json.Marshal(fetchUsersByExternalDisplayNameBulkPayload{externalAuthType, displayNames})
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

func (c *Client) FetchUsersByExternalIDBulk(externalAuthType ExternalAuthType, ids []string) (data map[ExternalAuthType]ExternalAuth, err error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	bodyBytes, err := json.Marshal(fetchUsersByExternalIDBulkPayload{externalAuthType, ids})
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