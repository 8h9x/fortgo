package friends

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) AcceptBulk(accountID string, targetIDs []string) ([]string, error) {
	query := url.Values{}
	query.Set("targetIds", strings.Join(targetIDs, ","))

	req, err := request.New(
		http.MethodPost,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/incoming/accept", accountID),
		request.WithQueryParamaters(query),
	)
	if err != nil {
		return []string{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []string{}, err
	}

	resp, err := request.DecodeResponse[[]string](res)
	if err != nil {
		return []string{}, err
	}

	return resp.Data, nil
}

func (c *Client) AddFriend(accountID string, targetID string, parentalControlsPin string) error {
	req, err := request.New(
		http.MethodPost,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends/%s", accountID, targetID),
		request.WithJSONBody(&ParentalControlsPinPayload{parentalControlsPin}),
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

func (c *Client) GetFriend(accountID string, friendID string) (Friend, error) {
	req, err := request.New(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends/%s", accountID, friendID),
	)
	if err != nil {
		return Friend{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return Friend{}, err
	}

	resp, err := request.DecodeResponse[Friend](res)
	if err != nil {
		return Friend{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetFriendsList(accountID string) ([]Friend, error) {
	req, err := request.New(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends", accountID),
	)
	if err != nil {
		return []Friend{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []Friend{}, err
	}

	resp, err := request.DecodeResponse[[]Friend](res)
	if err != nil {
		return []Friend{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetMutual(accountID string, friendID string) ([]string, error) {
	req, err := request.New(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends/%s/mutual", accountID, friendID),
	)
	if err != nil {
		return []string{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []string{}, err
	}

	resp, err := request.DecodeResponse[[]string](res)
	if err != nil {
		return []string{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetSummary(accountID string) (FriendSummaryResponse, error) {
	req, err := request.New(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/summary", accountID),
	)
	if err != nil {
		return FriendSummaryResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return FriendSummaryResponse{}, err
	}

	resp, err := request.DecodeResponse[FriendSummaryResponse](res)
	if err != nil {
		return FriendSummaryResponse{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetSuggested(accountID string) ([]SuggestedFriend, error) {
	req, err := request.New(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/suggested", accountID),
	)
	if err != nil {
		return []SuggestedFriend{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []SuggestedFriend{}, err
	}

	resp, err := request.DecodeResponse[[]SuggestedFriend](res)
	if err != nil {
		return []SuggestedFriend{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetRelations(accountID string) (any, error) {
	req, err := request.New(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/relations", accountID),
	)
	if err != nil {
		return []SuggestedFriend{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []SuggestedFriend{}, err
	}

	resp, err := request.DecodeResponse[[]SuggestedFriend](res)
	if err != nil {
		return []SuggestedFriend{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetRequestsIncoming(accountID string) ([]FriendRequest, error) {
	req, err := request.New(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/incoming", accountID),
	)
	if err != nil {
		return []FriendRequest{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []FriendRequest{}, err
	}

	resp, err := request.DecodeResponse[[]FriendRequest](res)
	if err != nil {
		return []FriendRequest{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetRequestsOutgoing(accountID string) ([]FriendRequest, error) {
	req, err := request.New(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/outgoing", accountID),
	)
	if err != nil {
		return []FriendRequest{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []FriendRequest{}, err
	}

	resp, err := request.DecodeResponse[[]FriendRequest](res)
	if err != nil {
		return []FriendRequest{}, err
	}

	return resp.Data, nil
}

func (c *Client) RemoveAllFriends(accountID string) error {
	req, err := request.New(
		http.MethodDelete,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends", accountID),
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

func (c *Client) RemoveFriend(accountID string, targetID string) error {
	req, err := request.New(
		http.MethodDelete,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends/%s", accountID, targetID),
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
