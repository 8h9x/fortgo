package fortnite

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetSystemCloudstorageConfig() (CloudstorageConfig, error) {
	req, err := request.New(
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

	resp, err := request.DecodeResponse[CloudstorageConfig](res)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetUserCloudstorageConfig() (CloudstorageConfig, error) {
	req, err := request.New(
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

	resp, err := request.DecodeResponse[CloudstorageConfig](res)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	return resp.Data, nil
}

func (c *Client) ListSystemCloudstorage() ([]CloudstorageFilePointerSystem, error) {
	req, err := request.New(
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

	resp, err := request.DecodeResponse[[]CloudstorageFilePointerSystem](res)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (c *Client) ListUserCloudstorage(accountID string) ([]CloudstorageFilePointerUser, error) {
	req, err := request.New(
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

	resp, err := request.DecodeResponse[[]CloudstorageFilePointerUser](res)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (c *Client) UpdateUserFilePointer(accountID string, filename string, shouldMigrate bool, payload UpdateUserFilePointerPayload) (CloudstorageFilePointerUser, error) {
	query := url.Values{}
	if shouldMigrate {
		query.Set("shouldMigrate", "true")
	}

	req, err := request.New(
		http.MethodPost,
		consts.FortniteMCPService,
		fmt.Sprintf("fortnite/api/cloudstorage/user/%s/%s/updatePointer", accountID, filename),
		request.WithQueryParamaters(query),
		request.WithJSONBody(&payload),
	)
	if err != nil {
		return CloudstorageFilePointerUser{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return CloudstorageFilePointerUser{}, err
	}

	resp, err := request.DecodeResponse[CloudstorageFilePointerUser](res)
	if err != nil {
		return CloudstorageFilePointerUser{}, err
	}

	return resp.Data, nil
}
