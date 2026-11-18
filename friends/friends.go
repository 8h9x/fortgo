package friends

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) AcceptBulk(accountID string, targetIDs []string) ([]string, error) {
	query := url.Values{}
	query.Set("targetIds", strings.Join(targetIDs, ","))

	req, err := rest.NewRequest(
		http.MethodPost,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/incoming/accept", accountID),
		rest.WithQueryParamaters(query),
	)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var data []string
	if err = rest.DecodeResponse(res, &data); err != nil {
		return nil, err
	}

	return data, err
}

func (c *Client) AddFriend(accountID string, targetID string, parentalControlsPin string) error {
	req, err := rest.NewRequest(
		http.MethodPost,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends/%s", accountID, targetID),
		rest.WithJSONBody(&ParentalControlsPinPayload{parentalControlsPin}),
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

func (c *Client) GetFriend(accountID string, friendID string) (Friend, error) {
	req, err := rest.NewRequest(
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

	var data Friend
	if err = rest.DecodeResponse(res, &data); err != nil {
		return Friend{}, err
	}

	return data, err
}

func (c *Client) GetFriendsList(accountID string) ([]Friend, error) {
	req, err := rest.NewRequest(
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

	var data []Friend
	if err = rest.DecodeResponse(res, &data); err != nil {
		return []Friend{}, err
	}

	return data, err
}

func (c *Client) GetMutual(accountID string, friendID string) ([]string, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/friends/%s/mutual", accountID, friendID),
	)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var data []string
	if err = rest.DecodeResponse(res, &data); err != nil {
		return nil, err
	}

	return data, err
}

func (c *Client) GetSummary(accountID string) (FriendSummaryResponse, error) {
	req, err := rest.NewRequest(
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

	var data FriendSummaryResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return FriendSummaryResponse{}, err
	}

	return data, err
}

func (c *Client) GetSuggested(accountID string) ([]SuggestedFriend, error) {
	req, err := rest.NewRequest(
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

	var data []SuggestedFriend
	if err = rest.DecodeResponse(res, &data); err != nil {
		return []SuggestedFriend{}, err
	}

	return data, err
}

func (c *Client) GetRelations(accountID string) (any, error) {
	req, err := rest.NewRequest(
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

	var data []SuggestedFriend
	if err = rest.DecodeResponse(res, &data); err != nil {
		return []SuggestedFriend{}, err
	}

	return data, err
}

func (c *Client) GetRequestsIncoming(accountID string) ([]FriendRequest, error) {
	req, err := rest.NewRequest(
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

	var data []FriendRequest
	if err = rest.DecodeResponse(res, &data); err != nil {
		return []FriendRequest{}, err
	}

	return data, err
}

func (c *Client) GetRequestsOutgoing(accountID string) ([]FriendRequest, error) {
	req, err := rest.NewRequest(
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

	var data []FriendRequest
	if err = rest.DecodeResponse(res, &data); err != nil {
		return []FriendRequest{}, err
	}

	return data, err
}

func (c *Client) RemoveAllFriends(accountID string) error {
	req, err := rest.NewRequest(
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

	if err = rest.DecodeResponse(res, new(any)); err != nil {
		return err
	}

	return nil
}

func (c *Client) RemoveFriend(accountID string, targetID string) error {
	req, err := rest.NewRequest(
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

	if err = rest.DecodeResponse(res, new(any)); err != nil {
		return err
	}

	return nil
}
