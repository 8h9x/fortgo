package friends

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) Block(accountID string, targetID string) error {
	req, err := rest.NewRequest(
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

	if err = rest.DecodeResponse(res, new(any)); err != nil {
		return err
	}

	return nil
}

func (c *Client) GetBlockedList(accountID string) ([]BlockedPlayer, error) {
	req, err := rest.NewRequest(
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

	var data []BlockedPlayer
	if err = rest.DecodeResponse(res, &data); err != nil {
		return []BlockedPlayer{}, err
	}

	return data, err
}

func (c *Client) Unblock(accountID string, targetID string) error {
	req, err := rest.NewRequest(
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

	if err = rest.DecodeResponse(res, new(any)); err != nil {
		return err
	}

	return nil
}
