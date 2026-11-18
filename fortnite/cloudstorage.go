package fortnite

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) GetSystemCloudstorageConfig() (CloudstorageConfig, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.FortniteMCPService,
		"fortnite/api/cloudstorage/system/config",
	)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	var data CloudstorageConfig
	if err = rest.DecodeResponse(res, &data); err != nil {
		return CloudstorageConfig{}, err
	}

	return data, err
}

func (c *Client) GetUserCloudstorageConfig() (CloudstorageConfig, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.FortniteMCPService,
		"fortnite/api/cloudstorage/user/config",
	)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	var data CloudstorageConfig
	if err = rest.DecodeResponse(res, &data); err != nil {
		return CloudstorageConfig{}, err
	}

	return data, err
}

func (c *Client) ListSystemCloudstorage() ([]CloudstorageFilePointerSystem, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.FortniteMCPService,
		"fortnite/api/cloudstorage/system",
	)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var data []CloudstorageFilePointerSystem
	if err = rest.DecodeResponse(res, &data); err != nil {
		return nil, err
	}

	return data, err
}

func (c *Client) ListUserCloudstorage(accountID string) ([]CloudstorageFilePointerUser, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.FortniteMCPService,
		fmt.Sprintf("fortnite/api/cloudstorage/user/%s", accountID),
	)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	var data []CloudstorageFilePointerUser
	if err = rest.DecodeResponse(res, &data); err != nil {
		return nil, err
	}

	return data, err
}

func (c *Client) UpdateUserFilePointer(accountID string, filename string, shouldMigrate bool, payload UpdateUserFilePointerPayload) (CloudstorageFilePointerUser, error) {
	query := url.Values{}
	if shouldMigrate {
		query.Set("shouldMigrate", "true")
	}

	req, err := rest.NewRequest(
		http.MethodPost,
		consts.FortniteMCPService,
		fmt.Sprintf("fortnite/api/cloudstorage/user/%s/%s/updatePointer", accountID, filename),
		rest.WithQueryParamaters(query),
		rest.WithJSONBody(&payload),
	)
	if err != nil {
		return CloudstorageFilePointerUser{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return CloudstorageFilePointerUser{}, err
	}

	var data CloudstorageFilePointerUser
	if err = rest.DecodeResponse(res, &data); err != nil {
		return CloudstorageFilePointerUser{}, err
	}

	return data, err
}
