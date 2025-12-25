package usersearch

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) Search(accountID string, platform Platform, term string) (SearchUsersResponse, error) {
	query := url.Values{}
	query.Set("platform", string(platform))
	query.Set("prefix", term)

	req, err := request.MakeRequest(
	    http.MethodGet,
		consts.UserSearchService,
		fmt.Sprintf("api/v1/search/%s", accountID),
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithQueryParamaters(query),
	)
	if err != nil {
		return SearchUsersResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return SearchUsersResponse{}, err
	}

	resp, err := request.ParseResponse[SearchUsersResponse](res)
	if err != nil {
		return SearchUsersResponse{}, err
	}

	return resp.Data, err
}