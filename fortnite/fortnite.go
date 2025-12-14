package fortnite

import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/8h9x/fortgo/auth"
	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func GetMCPVersion(httpClient *http.Client) (MCPVersionResponse, error) {
	req, err := http.NewRequest("GET", consts.FortniteMCPService+"/fortnite/api/version", nil)
	if err != nil {
		return MCPVersionResponse{}, err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return MCPVersionResponse{}, err
	}

	res, err := request.ResponseParser[MCPVersionResponse](resp)
	if err != nil {
		return MCPVersionResponse{}, err
	}

	return res.Body, err
}

func ComposeProfileOperation(httpClient *http.Client, credentials auth.TokenResponse, operation string, profileID string, payload string) (*http.Response, error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/fortnite/api/game/v2/profile/%s/client/%s?profileId=%s&rvn=-1", consts.FortniteMCPService, credentials.AccountID, operation, profileID)

	return request.Request(httpClient, "POST", reqUrl, headers, payload)
}

func ProfileOperation(httpClient *http.Client, credentials auth.TokenResponse, operation string, profileID string, payload any) (*http.Response, error) {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return &http.Response{}, err
	}

	return ComposeProfileOperation(httpClient, credentials, operation, profileID, string(bodyBytes))
}