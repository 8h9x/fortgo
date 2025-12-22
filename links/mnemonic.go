package links

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/auth"
	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func getMnemonicInfoRaw[T MnemonicData | MnemonicDataWithActivationHistory](httpClient *http.Client, credentials *auth.TokenResponse, namespace Namespace, mnemonic string, mnemonicType MnemonicType, version int) (T, error) {
	includeActivationHistory := false

	var data T
	switch any(data).(type) {
	case MnemonicDataWithActivationHistory:
		includeActivationHistory = true
	}

	query := url.Values{}
	query.Set("type", string(mnemonicType))
	if version != -1 {
		query.Set("v", string(rune(version)))
	}
	if includeActivationHistory {
		query.Set("includeActivationHistory", "true")
	}

	req, err := request.MakeRequest(
		http.MethodGet,
		consts.LinksService,
		fmt.Sprintf("links/api/%s/mnemonic/%s?%s", namespace, mnemonic, query.Encode()),
		request.WithBearerToken(credentials.AccessToken),
	)
	if err != nil {
		return data, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return data, err
	}

	resp, err := request.ParseResponse[T](res)
	if err != nil {
		return data, err
	}

	return resp.Data, nil
}

// GetMnemonicInfo fetches information about provided mnemonic on the namespace given a matching mnemonicType,
// set version to '-1' in order to fetch latest
func (c *Client) GetMnemonicInfo(namespace Namespace, mnemonic string, mnemonicType MnemonicType, version int) (MnemonicData, error) {
	return getMnemonicInfoRaw[MnemonicData](c.HTTPClient, c.Credentials, namespace, mnemonic, mnemonicType, version)
}

// GetMnemonicInfoWithActivationHistory fetches information with activation history and extended metadata about provided mnemonic on the namespace given a matching mnemonicType,
// set version to '-1' in order to fetch latest
func (c *Client) GetMnemonicInfoWithActivationHistory(httpClient *http.Client, credentials auth.TokenResponse, namespace Namespace, mnemonic string, mnemonicType MnemonicType, version int) (MnemonicDataWithActivationHistory, error) {
	return getMnemonicInfoRaw[MnemonicDataWithActivationHistory](c.HTTPClient, c.Credentials, namespace, mnemonic, mnemonicType, version)
}

func (c *Client) GetRelatedMnemonics(namespace Namespace, mnemonic string, version int) (GetRelatedMnemonicsResponse, error) {
	query := url.Values{}
	if version != -1 {
		query.Set("v", string(rune(version)))
	}

	req, err := request.MakeRequest(
		http.MethodGet,
		consts.LinksService,
		fmt.Sprintf("links/api/%s/mnemonic/%s/related?%s", namespace, mnemonic, query.Encode()),
	)
	if err != nil {
		return GetRelatedMnemonicsResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return GetRelatedMnemonicsResponse{}, err
	}

	resp, err := request.ParseResponse[GetRelatedMnemonicsResponse](res)
	if err != nil {
		return GetRelatedMnemonicsResponse{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetMnemonicInfoBulk(namespace Namespace, mnemonics []GetMnemonicInfoBulkPayload, ignoreFailures bool) ([]MnemonicData, error) {
	query := url.Values{}
	if ignoreFailures {
		query.Set("ignoreFailures", "true")
	}

	req, err := request.MakeRequest(
		http.MethodPost,
		consts.LinksService,
		fmt.Sprintf("links/api/%s/mnemonic?%s", namespace, query.Encode()),
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithJSONBody(&mnemonics),
	)
	if err != nil {
		return []MnemonicData{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []MnemonicData{}, err
	}

	resp, err := request.ParseResponse[[]MnemonicData](res)
	if err != nil {
		return []MnemonicData{}, err
	}

	return resp.Data, nil
}