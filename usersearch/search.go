package usersearch

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) Search(accountID string, platform Platform, term string) (SearchUsersResponse, error) {
	query := url.Values{}
	query.Set("platform", string(platform))
	query.Set("prefix", term)

	req, err := rest.NewRequest(
		http.MethodGet,
		consts.UserSearchService,
		fmt.Sprintf("api/v1/search/%s", accountID),
		rest.WithQueryParameters(query),
	)
	if err != nil {
		return SearchUsersResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return SearchUsersResponse{}, err
	}

	var data SearchUsersResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return SearchUsersResponse{}, err
	}

	return data, err
}
