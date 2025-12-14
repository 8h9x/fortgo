package locker

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) ChangeCompanionName(accountID string, cosmeticItemID string, companionName string) error {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	bodyBytes, err := json.Marshal(&ChangeCompanionNamePayload{cosmeticItemID, companionName})
	if err != nil {
		return err
	}

	reqUrl := fmt.Sprintf("/%s/api/locker/v4/%s/account/%s/companion-name", consts.FortniteLockerService, DeploymentIDLiveFN, accountID)

	resp, err := request.Request(c.HTTPClient, "PATCH", reqUrl, headers, string(bodyBytes))
	if err != nil {
		return err
	}

	_, err = request.ResponseParser[interface{}](resp)
	return err
}

func (c *Client) QueryItems(accountID string) (LockerItems, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("/%s/api/locker/v4/%s/account/%s/items", consts.FortniteLockerService, DeploymentIDLiveFN, accountID)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return LockerItems{}, err
	}

	res, err := request.ResponseParser[LockerItems](resp)
	if err != nil {
		return LockerItems{}, err
	}

	return res.Body, nil
}

func (c *Client) UpdateActiveLockerLoadout(accountID string, payload UpdateActiveLockerLoadoutPayload) (ActiveLoadoutGroup, error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("/%s/api/locker/v4/%s/account/%s/active-loadout-group", consts.FortniteLockerService, DeploymentIDLiveFN, accountID)

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	resp, err := request.Request(c.HTTPClient, "PUT", reqUrl, headers, string(bodyBytes))
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	res, err := request.ResponseParser[ActiveLoadoutGroup](resp)
	if err != nil {
		return ActiveLoadoutGroup{}, err
	}

	return res.Body, nil
}