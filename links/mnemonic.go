package links

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func getMnemonicInfoRaw[T MnemonicData | MnemonicDataWithActivationHistory](httpClient *http.Client, namespace Namespace, mnemonic string, mnemonicType MnemonicType, version int) (T, error) {
	includeActivationHistory := false

	var data T
	switch any(data).(type) {
	case MnemonicDataWithActivationHistory:
		includeActivationHistory = true
	}

	query := url.Values{}
	query.Set("type", string(mnemonicType))
	if version != -1 {
		query.Set("v", strconv.Itoa(version))
	}
	if includeActivationHistory {
		query.Set("includeActivationHistory", "true")
	}

	req, err := rest.NewRequest(
		http.MethodGet,
		consts.LinksService,
		fmt.Sprintf("links/api/%s/mnemonic/%s", namespace, mnemonic),
		rest.WithQueryParameters(query),
	)
	if err != nil {
		return data, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return data, err
	}

	if err = rest.DecodeResponse(res, &data); err != nil {
		return data, err
	}

	return data, err
}

// GetMnemonicInfo fetches information about provided mnemonic on the namespace given a matching mnemonicType,
// set version to '-1' in order to fetch latest
func (c *Client) GetMnemonicInfo(namespace Namespace, mnemonic string, mnemonicType MnemonicType, version int) (MnemonicData, error) {
	return getMnemonicInfoRaw[MnemonicData](c.HTTPClient, namespace, mnemonic, mnemonicType, version)
}

// GetMnemonicInfoWithActivationHistory fetches information with activation history and extended metadata about provided mnemonic on the namespace given a matching mnemonicType,
// set version to '-1' in order to fetch latest
func (c *Client) GetMnemonicInfoWithActivationHistory(namespace Namespace, mnemonic string, mnemonicType MnemonicType, version int) (MnemonicDataWithActivationHistory, error) {
	return getMnemonicInfoRaw[MnemonicDataWithActivationHistory](c.HTTPClient, namespace, mnemonic, mnemonicType, version)
}

func (c *Client) GetRelatedMnemonics(namespace Namespace, mnemonic string, version int) (GetRelatedMnemonicsResponse, error) {
	query := url.Values{}
	if version != -1 {
		query.Set("v", string(rune(version)))
	}

	req, err := rest.NewRequest(
		http.MethodGet,
		consts.LinksService,
		fmt.Sprintf("links/api/%s/mnemonic/%s/related", namespace, mnemonic),
		rest.WithQueryParameters(query),
	)
	if err != nil {
		return GetRelatedMnemonicsResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return GetRelatedMnemonicsResponse{}, err
	}

	var data GetRelatedMnemonicsResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return GetRelatedMnemonicsResponse{}, err
	}

	return data, err
}

func (c *Client) GetMnemonicInfoBulk(namespace Namespace, mnemonics []GetMnemonicInfoBulkPayload, ignoreFailures bool) ([]MnemonicData, error) {
	query := url.Values{}
	if ignoreFailures {
		query.Set("ignoreFailures", "true")
	}

	req, err := rest.NewRequest(
		http.MethodPost,
		consts.LinksService,
		fmt.Sprintf("links/api/%s/mnemonic", namespace),
		rest.WithQueryParameters(query),
		rest.WithJSONBody(&mnemonics),
	)
	if err != nil {
		return []MnemonicData{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return []MnemonicData{}, err
	}

	var data []MnemonicData
	if err = rest.DecodeResponse(res, &data); err != nil {
		return []MnemonicData{}, err
	}

	return data, err
}
