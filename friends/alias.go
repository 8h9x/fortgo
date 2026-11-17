package friends

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) EditAlias(accountID string, friendID string, alias string) error {
	req, err := request.New(
		http.MethodPut,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends/%s/alias", accountID, friendID),
		request.WithTextBody(alias),
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

func (c *Client) RemoveAlias(accountID string, friendID string) error {
	req, err := request.New(
		http.MethodDelete,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends/%s/alias", accountID, friendID),
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
