package usersearch

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) Search(accountID string, platform Platform, term string) (SearchUsersResponse, error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	query := url.Values{}
	query.Set("platform", string(platform))
	query.Set("prefix", term)

	reqUrl := fmt.Sprintf("%s/api/v1/search/%s?%s", consts.UserSearchService, accountID, query.Encode())

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return SearchUsersResponse{}, err
	}

	res, err := request.ResponseParser[SearchUsersResponse](resp)
	if err != nil {
		return SearchUsersResponse{}, err
	}

	return res.Body, err
}