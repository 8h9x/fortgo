package locker

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) ChangeCompanionName(accountID string, cosmeticItemID string, companionName string) error {
	payload := &ChangeCompanionNamePayload{cosmeticItemID, companionName}

	req, err := request.MakeRequest(
		http.MethodPatch,
		consts.FortniteLockerService,
		fmt.Sprintf("api/locker/v4/%s/account/%s/companion-name", DeploymentIDLiveFN, accountID),
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithJSONBody(payload),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	_, err = request.ParseResponse[any](res)
	return err
}

func (c *Client) QueryItems(accountID string) (LockerItems, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.FortniteLockerService,
		"api/locker/v4/%s/account/%s/items",
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return LockerItems{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return LockerItems{}, err
	}

	resp, err := request.ParseResponse[LockerItems](res)
	if err != nil {
		return LockerItems{}, err
	}

	return resp.Data, nil
}

func (c *Client) UpdateActiveLockerLoadout(accountID string, payload UpdateActiveLockerLoadoutPayload) (ActiveLoadoutGroup, error) {
	req, err := request.MakeRequest(
		http.MethodPut,
		consts.FortniteLockerService,
		fmt.Sprintf("api/locker/v4/%s/account/%s/active-loadout-group", DeploymentIDLiveFN, accountID),
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithJSONBody(payload),
	)
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	resp, err := request.ParseResponse[ActiveLoadoutGroup](res)
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	return resp.Data, nil
}

func (c *Client) LockInImmutableItem(accountID string, templateID string, itemGUID string, variants map[string]ItemCustomization) error {
	payload := &LockInImmutableItemPayload{variants}

	req, err := request.MakeRequest(
		http.MethodPost,
		consts.FortniteLockerService,
		fmt.Sprintf("api/locker/v4/%s/account/%s/lock-in-immutable-item/%s:%s", DeploymentIDLiveFN, accountID, templateID, itemGUID),
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithJSONBody(payload),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	_, err = request.ParseResponse[any](res)
	return err
}