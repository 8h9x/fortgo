package fortnite

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetMCPVersion() (MCPVersionResponse, error) {
	req, err := request.MakeRequest(
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

	resp, err := request.ParseResponse[MCPVersionResponse](res)
	if err != nil {
		return MCPVersionResponse{}, err
	}

	return resp.Data, err
}

func (c *Client) ComposeProfileOperation(accountID string, operation string, profileID string, payload string) (*http.Response, error) {
	req, err := request.MakeRequest(
		http.MethodPost,
		consts.FortniteMCPService,
		fmt.Sprintf("fortnite/api/game/v2/profile/%s/client/%s?profileId=%s&rvn=-1", accountID, operation, profileID),
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithTextBody(payload),
		request.WithHeader("Content-Type", "application/json"), // TODO: fix, this is yucky
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