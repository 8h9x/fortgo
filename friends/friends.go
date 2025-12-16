package friends

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetSuggested(accountID string) (FriendSuggestionResponse, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("/friends/api/v1/%s/suggested", consts.FriendsService, accountID)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return FriendSuggestionResponse{}, err
	}

	res, err := request.ResponseParser[FriendSuggestionResponse](resp)
	if err != nil {
		return FriendSuggestionResponse{}, err
	}

	return res.Body, nil
}
