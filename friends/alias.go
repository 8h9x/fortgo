package friends

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) EditAlias(accountID string, friendID string, alias string) error {
	req, err := rest.NewRequest(
		http.MethodPut,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends/%s/alias", accountID, friendID),
		rest.WithTextBody(alias),
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

func (c *Client) RemoveAlias(accountID string, friendID string) error {
	req, err := rest.NewRequest(
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

	if err = rest.DecodeResponse(res, new(any)); err != nil {
		return err
	}

	return nil
}
