package links

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/auth"
	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func fetchMnemonicInfoRaw[T MnemonicData | MnemonicDataWithActivationHistory](httpClient *http.Client, credentials auth.TokenResponse, namespace Namespace, mnemonic string, mnemonicType MnemonicType, version int) (T, error) {
	includeActivationHistory := false

	var data T
	switch any(data).(type) {
	case MnemonicDataWithActivationHistory:
		includeActivationHistory = true
	}

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	query := url.Values{}
	query.Set("type", string(mnemonicType))
	if version != -1 {
		query.Set("v", string(rune(version)))
	}
	if includeActivationHistory {
		query.Set("includeActivationHistory", "true")
	}

	reqUrl := fmt.Sprintf("%s/links/api/%s/mnemonic/%s?%s", consts.LinksService, namespace, mnemonic, query.Encode())

	resp, err := request.Request(httpClient, "GET", reqUrl, headers, "")
	if err != nil {
		return data, err
	}

	res, err := request.ResponseParser[T](resp)
	if err != nil {
		return data, err
	}

	return res.Body, err
}

// FetchMnemonicInfo fetches information about provided mnemonic on the namespace given a matching mnemonicType,
// set version to '-1' in order to fetch latest
func FetchMnemonicInfo(httpClient *http.Client, credentials auth.TokenResponse, namespace Namespace, mnemonic string, mnemonicType MnemonicType, version int) (MnemonicData, error) {
	return fetchMnemonicInfoRaw[MnemonicData](httpClient, credentials, namespace, mnemonic, mnemonicType, version)
}

// FetchMnemonicInfoWithActivationHistory fetches information with activation history and extended metadata about provided mnemonic on the namespace given a matching mnemonicType,
// set version to '-1' in order to fetch latest
func FetchMnemonicInfoWithActivationHistory(httpClient *http.Client, credentials auth.TokenResponse, namespace Namespace, mnemonic string, mnemonicType MnemonicType, version int) (MnemonicDataWithActivationHistory, error) {
	return fetchMnemonicInfoRaw[MnemonicDataWithActivationHistory](httpClient, credentials, namespace, mnemonic, mnemonicType, version)
}

func FetchRelatedMnemonics(httpClient *http.Client, credentials auth.TokenResponse, namespace Namespace, mnemonic string, version int) (FetchRelatedMnemonicsResponse, error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	query := url.Values{}
	if version != -1 {
		query.Set("v", string(rune(version)))
	}

	reqUrl := fmt.Sprintf("%s/links/api/%s/mnemonic/%s/related?%s", consts.LinksService, namespace, mnemonic, query.Encode())

	resp, err := request.Request(httpClient, "GET", reqUrl, headers, "")
	if err != nil {
		return FetchRelatedMnemonicsResponse{}, err
	}

	res, err := request.ResponseParser[FetchRelatedMnemonicsResponse](resp)
	if err != nil {
		return FetchRelatedMnemonicsResponse{}, err
	}

	return res.Body, err
}

func FetchMnemonicInfoBulk(httpClient *http.Client, credentials auth.TokenResponse, namespace Namespace, mnemonics []FetchMnemonicInfoBulkPayload, ignoreFailures bool) ([]MnemonicData, error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	query := url.Values{}
	if ignoreFailures {
		query.Set("ignoreFailures", "true")
	}

	bodyBytes, err := json.Marshal(mnemonics)
	if err != nil {
		return []MnemonicData{}, err
	}

	reqUrl := fmt.Sprintf("%s/links/api/%s/mnemonic?%s", consts.LinksService, namespace, query.Encode())

	resp, err := request.Request(httpClient, "POST", reqUrl, headers, string(bodyBytes))
	if err != nil {
		return []MnemonicData{}, err
	}

	res, err := request.ResponseParser[[]MnemonicData](resp)
	if err != nil {
		return []MnemonicData{}, err
	}

	return res.Body, err
}