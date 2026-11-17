package locker

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) ChangeCompanionName(accountID string, cosmeticItemID string, companionName string) error {
	payload := &ChangeCompanionNamePayload{cosmeticItemID, companionName}

	req, err := request.New(
		http.MethodPatch,
		consts.FortniteLockerService,
		fmt.Sprintf("api/locker/v4/%s/account/%s/companion-name", consts.DeploymentIDLiveFN, accountID),
		request.WithJSONBody(payload),
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

func (c *Client) FetchCosmeticData(accountID string, limit uint8, nextToken string) (CosmeticData, error) {
	query := url.Values{}
	if limit > 0 {
		query.Set("limit", strconv.Itoa(int(limit)))
	}
	if nextToken != "" {
		query.Set("nextToken", nextToken)
	}

	req, err := request.New(
		http.MethodGet,
		consts.FortniteLockerService,
		"api/locker/v4/%s/account/%s/cosmetic-data",
		request.WithQueryParamaters(query),
	)
	if err != nil {
		return CosmeticData{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return CosmeticData{}, err
	}

	resp, err := request.DecodeResponse[CosmeticData](res)
	if err != nil {
		return CosmeticData{}, err
	}

	return resp.Data, nil
}

func (c *Client) LockInImmutableItem(accountID string, templateID string, itemGUID string, variants map[string]ItemCustomization) error {
	payload := &LockInImmutableItemPayload{variants}

	req, err := request.New(
		http.MethodPost,
		consts.FortniteLockerService,
		fmt.Sprintf("api/locker/v4/%s/account/%s/lock-in-immutable-item/%s:%s", consts.DeploymentIDLiveFN, accountID, templateID, itemGUID),
		request.WithJSONBody(payload),
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

func (c *Client) QueryItems(accountID string) (LockerItems, error) {
	req, err := request.New(
		http.MethodGet,
		consts.FortniteLockerService,
		"api/locker/v4/%s/account/%s/items",
	)
	if err != nil {
		return LockerItems{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return LockerItems{}, err
	}

	resp, err := request.DecodeResponse[LockerItems](res)
	if err != nil {
		return LockerItems{}, err
	}

	return resp.Data, nil
}

func (c *Client) UpdateActiveLoadoutGroup(accountID string, payload UpdateActiveLockerLoadoutPayload) (ActiveLoadoutGroup, error) {
	req, err := request.New(
		http.MethodPut,
		consts.FortniteLockerService,
		fmt.Sprintf("api/locker/v4/%s/account/%s/active-loadout-group", consts.DeploymentIDLiveFN, accountID),
		request.WithJSONBody(payload),
	)
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	resp, err := request.DecodeResponse[ActiveLoadoutGroup](res)
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	return resp.Data, nil
}
