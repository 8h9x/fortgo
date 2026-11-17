package friends

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) EditNode(accountID string, friendID string, note string) error {
	req, err := request.New(
		http.MethodPut,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends/%s/note", accountID, friendID),
		request.WithTextBody(note),
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

func (c *Client) RemoveNode(accountID string, friendID string) error {
	req, err := request.New(
		http.MethodDelete,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends/%s/note", accountID, friendID),
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
