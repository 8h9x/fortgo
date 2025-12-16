package friends

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetExternalSettings(accountID string, source ExternalSource) (ExternalSourcesSettings, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/settings/externalSources/%s", consts.FriendsService, accountID, source)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	res, err := request.ResponseParser[ExternalSourcesSettings](resp)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	return res.Body, nil
}

func (c *Client) SetExternalSettings(accountID string, source ExternalSource, doNotShowLinkingProposal bool) (ExternalSourcesSettings, error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	bodyBytes, err := json.Marshal(&ExternalSourcesSettings{doNotShowLinkingProposal})
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/settings/externalSources/%s", consts.FriendsService, accountID, source)

	resp, err := request.Request(c.HTTPClient, "PUT", reqUrl, headers, string(bodyBytes))
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	res, err := request.ResponseParser[ExternalSourcesSettings](resp)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	return res.Body, nil
}

func (c *Client) GetPrivacySettings(accountID string) (FriendsPrivacySettings, error) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/settings", consts.FriendsService, accountID)

	resp, err := request.Request(c.HTTPClient, "GET", reqUrl, headers, "")
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	res, err := request.ResponseParser[FriendsPrivacySettings](resp)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	return res.Body, nil
}

func (c *Client) SetPrivacySettings(accountID string, acceptInvites PrivacySettingAcceptInvites, mutualPrivacy PrivacySettingMutualPrivacy) (FriendsPrivacySettings, error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	bodyBytes, err := json.Marshal(&FriendsPrivacySettings{acceptInvites, mutualPrivacy})
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	reqUrl := fmt.Sprintf("%s/friends/api/v1/%s/settings", consts.FriendsService, accountID)

	resp, err := request.Request(c.HTTPClient, "PATCH", reqUrl, headers, string(bodyBytes))
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	res, err := request.ResponseParser[FriendsPrivacySettings](resp)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	return res.Body, nil
}