package friends

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) AcceptBulk(accountID string, targetIDs []string) ([]string, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	query := url.Values{}
	query.Set("targetIds", strings.Join(targetIDs, ","))

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/incoming/accept%s", consts.FriendsService, accountID, query.Encode())

	resp, err := request.Request(c.HTTPClient, "POST", reqUrl, headers, "")
	if err != nil {
		return []string{}, err
	}

	res, err := request.ResponseParser[[]string](resp)
	if err != nil {
		return []string{}, err
	}

	return res.Body, nil
}

func (c *Client) AddFriend(accountID string, targetID string, parentalControlsPin string) error {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	bodyBytes, err := json.Marshal(&ParentalControlsPinPayload{parentalControlsPin})
	if err != nil {
		return err
	}

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/friends/%s", consts.FriendsService, accountID, targetID)

	resp, err := request.Request(c.HTTPClient, "POST", reqUrl, headers, string(bodyBytes))
	if err != nil {
		return err
	}

	_, err = request.ResponseParser[any](resp)
	return err
}

func (c *Client) GetFriend(accountID string, friendID string) (Friend, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/friends/%s", consts.FriendsService, accountID, friendID)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return Friend{}, err
	}

	res, err := request.ResponseParser[Friend](resp)
	if err != nil {
		return Friend{}, err
	}

	return res.Body, nil
}

func (c *Client) GetFriendsList(accountID string) ([]Friend, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/friends", consts.FriendsService, accountID)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return []Friend{}, err
	}

	res, err := request.ResponseParser[[]Friend](resp)
	if err != nil {
		return []Friend{}, err
	}

	return res.Body, nil
}

func (c *Client) GetMutual(accountID string, friendID string) ([]string, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/friends/%s/mutual", consts.FriendsService, accountID, friendID)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return []string{}, err
	}

	res, err := request.ResponseParser[[]string](resp)
	if err != nil {
		return []string{}, err
	}

	return res.Body, nil
}

func (c *Client) GetSummary(accountID string) (FriendSummaryResponse, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/summary", consts.FriendsService, accountID)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return FriendSummaryResponse{}, err
	}

	res, err := request.ResponseParser[FriendSummaryResponse](resp)
	if err != nil {
		return FriendSummaryResponse{}, err
	}

	return res.Body, nil
}

func (c *Client) GetSuggested(accountID string) ([]SuggestedFriend, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/suggested", consts.FriendsService, accountID)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return []SuggestedFriend{}, err
	}

	res, err := request.ResponseParser[[]SuggestedFriend](resp)
	if err != nil {
		return []SuggestedFriend{}, err
	}

	return res.Body, nil
}

func (c *Client) GetRelations(accountID string) (any, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/relations", consts.FriendsService, accountID)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return struct{}{}, err
	}

	res, err := request.ResponseParser[any](resp)
	if err != nil {
		return struct{}{}, err
	}

	return res.Body, nil
}

func (c *Client) GetRequestsIncoming(accountID string) ([]FriendRequest, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/incoming", consts.FriendsService, accountID)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return []FriendRequest{}, err
	}

	res, err := request.ResponseParser[[]FriendRequest](resp)
	if err != nil {
		return []FriendRequest{}, err
	}

	return res.Body, nil
}

func (c *Client) GetRequestsOutgoing(accountID string) ([]FriendRequest, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/outgoing", consts.FriendsService, accountID)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return []FriendRequest{}, err
	}

	res, err := request.ResponseParser[[]FriendRequest](resp)
	if err != nil {
		return []FriendRequest{}, err
	}

	return res.Body, nil
}

func (c *Client) RemoveAllFriends(accountID string) error {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/friends", consts.FriendsService, accountID)

	resp, err := request.Request(c.HTTPClient, "DELETE", reqUrl, headers, "")
	if err != nil {
		return err
	}

	_, err = request.ResponseParser[[]FriendRequest](resp)
	return err
}

func (c *Client) RemoveFriend(accountID string, targetID string) error {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/friends/%s", consts.FriendsService, accountID, targetID)

	resp, err := request.Request(c.HTTPClient, "DELETE", reqUrl, headers, "")
	if err != nil {
		return err
	}

	_, err = request.ResponseParser[any](resp)
	return err
}