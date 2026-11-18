package fortnite

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) GetMCPVersion() (MCPVersionResponse, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.FortniteMCPService,
		"fortnite/api/version",
	)
	if err != nil {
		return MCPVersionResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return MCPVersionResponse{}, err
	}

	var data MCPVersionResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return MCPVersionResponse{}, err
	}

	return data, err
}

func (c *Client) ComposeProfileOperation(accountID string, operation string, profileID string, payload string) (*http.Response, error) {
	req, err := rest.NewRequest(
		http.MethodPost,
		consts.FortniteMCPService,
		fmt.Sprintf("fortnite/api/game/v2/profile/%s/client/%s", accountID, operation),
		rest.WithQueryParameters(url.Values{
			"profileId": {profileID},
			"rvn":       {"-1"},
		}),
		rest.WithTextBody(payload),
		rest.WithHeader("Content-Type", "application/json"), // TODO: fix, this is yucky
	)
	if err != nil {
		return &http.Response{}, err
	}

	return c.HTTPClient.Do(req)
}

func (c *Client) ProfileOperation(accountID string, operation string, profileID string, payload any) (*http.Response, error) {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return &http.Response{}, err
	}

	return c.ComposeProfileOperation(accountID, operation, profileID, string(bodyBytes))
}
