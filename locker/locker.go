package locker

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) ChangeCompanionName(accountID string, cosmeticItemID string, companionName string) error {
	payload := &ChangeCompanionNamePayload{cosmeticItemID, companionName}

	req, err := rest.NewRequest(
		http.MethodPatch,
		consts.FortniteLockerService,
		fmt.Sprintf("api/locker/v4/%s/account/%s/companion-name", consts.DeploymentIDLiveFN, accountID),
		rest.WithJSONBody(payload),
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

func (c *Client) FetchCosmeticData(accountID string, limit uint8, nextToken string) (CosmeticData, error) {
	query := url.Values{}
	if limit > 0 {
		query.Set("limit", strconv.Itoa(int(limit)))
	}
	if nextToken != "" {
		query.Set("nextToken", nextToken)
	}

	req, err := rest.NewRequest(
		http.MethodGet,
		consts.FortniteLockerService,
		"api/locker/v4/%s/account/%s/cosmetic-data",
		rest.WithQueryParamaters(query),
	)
	if err != nil {
		return CosmeticData{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return CosmeticData{}, err
	}

	var data CosmeticData
	if err = rest.DecodeResponse(res, &data); err != nil {
		return CosmeticData{}, err
	}

	return data, err
}

func (c *Client) LockInImmutableItem(accountID string, templateID string, itemGUID string, variants map[string]ItemCustomization) error {
	payload := &LockInImmutableItemPayload{variants}

	req, err := rest.NewRequest(
		http.MethodPost,
		consts.FortniteLockerService,
		fmt.Sprintf("api/locker/v4/%s/account/%s/lock-in-immutable-item/%s:%s", consts.DeploymentIDLiveFN, accountID, templateID, itemGUID),
		rest.WithJSONBody(payload),
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

func (c *Client) QueryItems(accountID string) (LockerItems, error) {
	req, err := rest.NewRequest(
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

	var data LockerItems
	if err = rest.DecodeResponse(res, &data); err != nil {
		return LockerItems{}, err
	}

	return data, err
}

func (c *Client) UpdateActiveLoadoutGroup(accountID string, payload UpdateActiveLockerLoadoutPayload) (ActiveLoadoutGroup, error) {
	req, err := rest.NewRequest(
		http.MethodPut,
		consts.FortniteLockerService,
		fmt.Sprintf("api/locker/v4/%s/account/%s/active-loadout-group", consts.DeploymentIDLiveFN, accountID),
		rest.WithJSONBody(payload),
	)
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	var data ActiveLoadoutGroup
	if err = rest.DecodeResponse(res, &data); err != nil {
		return ActiveLoadoutGroup{}, err
	}

	return data, err
}
