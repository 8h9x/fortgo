package avatars

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) Get(accountIDs ...string) (GetAvatarsResponse, error) {
	query := url.Values{}
	query.Set("accountIds", strings.Join(accountIDs, ","))

	req, err := rest.NewRequest(
		http.MethodGet,
		consts.AvatarService,
		"v1/avatar/fortnite/ids",
		rest.WithQueryParameters(query),
	)
	if err != nil {
		return GetAvatarsResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return GetAvatarsResponse{}, err
	}

	var data GetAvatarsResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return GetAvatarsResponse{}, err
	}

	return data, err
}

func (c *Client) GetOne(accountID string) (AccountAvatarData, error) {
	res, err := c.Get(accountID)
	if err != nil {
		return AccountAvatarData{}, err
	}

	if len(res) == 0 {
		return AccountAvatarData{}, errors.New("returned avatar data list has no entries")
	}

	return res[0], nil
}
