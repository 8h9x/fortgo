package fortnite

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetSystemCloudstorageConfig() (CloudstorageConfig, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.FortniteMCPService,
		"fortnite/api/cloudstorage/system/config",
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	resp, err := request.ParseResponse[CloudstorageConfig](res)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetUserCloudstorageConfig() (CloudstorageConfig, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.FortniteMCPService,
		"fortnite/api/cloudstorage/user/config",
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	resp, err := request.ParseResponse[CloudstorageConfig](res)
	if err != nil {
		return CloudstorageConfig{}, err
	}

	return resp.Data, nil
}

func (c *Client) ListSystemCloudstorage() ([]CloudstorageFilePointerSystem, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.FortniteMCPService,
		"fortnite/api/cloudstorage/system",
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	resp, err := request.ParseResponse[[]CloudstorageFilePointerSystem](res)
	if err != nil {
		return nil, err
	}

	return resp.Data, nil
}

func (c *Client) ListUserCloudstorage(accountID string) ([]CloudstorageFilePointerUser, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.FortniteMCPService,
		fmt.Sprintf("fortnite/api/cloudstorage/user/%s", accountID),
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	resp, err := request.ParseResponse[[]CloudstorageFilePointerUser](res)
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

	req, err := request.MakeRequest(
		http.MethodPost,
		consts.FortniteMCPService,
		fmt.Sprintf("fortnite/api/cloudstorage/user/%s/%s/updatePointer", accountID, filename),
		request.WithBearerToken(c.Credentials.AccessToken),
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

	resp, err := request.ParseResponse[CloudstorageFilePointerUser](res)
	if err != nil {
		return CloudstorageFilePointerUser{}, err
	}

	return resp.Data, nil
}
