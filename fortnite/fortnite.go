package fortnite

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetMCPVersion() (MCPVersionResponse, error) {
	req, err := http.NewRequest("GET", consts.FortniteMCPService+"/fortnite/api/version", nil)
	if err != nil {
		return MCPVersionResponse{}, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return MCPVersionResponse{}, err
	}

	res, err := request.ResponseParser[MCPVersionResponse](resp)
	if err != nil {
		return MCPVersionResponse{}, err
	}

	return res.Body, err
}

func (c *Client) ComposeProfileOperation(accountID string, operation string, profileID string, payload string) (*http.Response, error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/fortnite/api/game/v2/profile/%s/client/%s?profileId=%s&rvn=-1", consts.FortniteMCPService, accountID, operation, profileID)

	return request.Request(c.HTTPClient, "POST", reqUrl, headers, payload)
}

func (c *Client) ProfileOperation(accountID string, operation string, profileID string, payload any) (*http.Response, error) {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return &http.Response{}, err
	}

	return c.ComposeProfileOperation(accountID, operation, profileID, string(bodyBytes))
}