package avatars

import (
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) Get(accountIDs ...string) (GetAvatarsResponse, error) {
	query := url.Values{}
	query.Set("accountIds", strings.Join(accountIDs, ","))

	req, err := request.New(
		http.MethodGet,
		consts.AvatarService,
		"v1/avatar/fortnite/ids",
		request.WithQueryParamaters(query),
	)
	if err != nil {
		return GetAvatarsResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return GetAvatarsResponse{}, err
	}

	resp, err := request.DecodeResponse[GetAvatarsResponse](res)
	if err != nil {
		return GetAvatarsResponse{}, err
	}

	return resp.Data, err
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
