package friends

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) Block(accountID string, targetID string) error {
	req, err := request.MakeRequest(
		http.MethodPost,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/blocklist/%s", accountID, targetID),
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	_, err = request.ParseResponse[any](res)
	return err
}

//func (c *Client) GetBlockedList(accountID string) ([]BlockedPlayer, error) {
//	headers := http.Header{}
//	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)
//
//	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/blocklist", consts.FriendsService, accountID)
//
//	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
//	if err != nil {
//		return []BlockedPlayer{}, err
//	}
//
//	res, err := request.ResponseParser[[]BlockedPlayer](resp)
//	if err != nil {
//		return []BlockedPlayer{}, err
//	}
//
//	return res.Body, nil
//}

func (c *Client) Unblock(accountID string, targetID string) error {
	req, err := request.MakeRequest(
		http.MethodDelete,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/blocklist/%s", accountID, targetID),
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	_, err = request.ParseResponse[any](res)
	return err
}