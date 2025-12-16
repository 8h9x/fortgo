package friends

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) EditNode(accountID string, friendID string, note string) error {
	headers := http.Header{}
	headers.Set("Content-Type", "text/plain")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/friends/%s/note", consts.FriendsService, accountID, friendID)

	resp, err := request.Request(c.HTTPClient, "PUT", reqUrl, headers, note)
	if err != nil {
		return err
	}

	_, err = request.ResponseParser[any](resp)
	return err
}

func (c *Client) RemoveNode(accountID string, friendID string) error {
	headers := http.Header{}
	headers.Set("Content-Type", "text/plain")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/friends/%s/note", consts.FriendsService, accountID, friendID)

	resp, err := request.Request(c.HTTPClient, "DELETE", reqUrl, headers, "")
	if err != nil {
		return err
	}

	_, err = request.ResponseParser[any](resp)
	return err
}