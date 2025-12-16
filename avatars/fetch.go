package avatars

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) Get(accountIDs ...string) (GetAvatarsResponse, error) {
	if len(accountIDs) == 0 {
		accountIDs = []string{c.Credentials.AccountID}
	}

	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	query := url.Values{}
	query.Set("accountIds", strings.Join(accountIDs, ","))

	reqUrl := fmt.Sprintf("%s/v1/avatar/fortnite/ids?%s", consts.AvatarService, query.Encode())

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return GetAvatarsResponse{}, err
	}

	res, err := request.ResponseParser[GetAvatarsResponse](resp)
	if err != nil {
		return GetAvatarsResponse{}, err
	}

	return res.Body, err
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