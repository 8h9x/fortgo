package friends

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) Block(accountID string, targetID string) error {
	req, err := request.New(
		http.MethodPost,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/blocklist/%s", accountID, targetID),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	_, err = request.DecodeResponse[any](res)
	return err
}

func (c *Client) GetBlockedList(accountID string) ([]BlockedPlayer, error) {
	req, err := request.New(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/blocklist", accountID),
	)
	if err != nil {
		return []BlockedPlayer{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []BlockedPlayer{}, err
	}

	resp, err := request.DecodeResponse[[]BlockedPlayer](res)
	if err != nil {
		return []BlockedPlayer{}, err
	}

	return resp.Data, nil
}

func (c *Client) Unblock(accountID string, targetID string) error {
	req, err := request.New(
		http.MethodDelete,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/blocklist/%s", accountID, targetID),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	_, err = request.DecodeResponse[any](res)
	return err
}
