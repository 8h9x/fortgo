package friends

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) GetExternalSettings(accountID string, source ExternalSource) (ExternalSourcesSettings, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/settings/externalSources/%s", accountID, source),
	)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	var data ExternalSourcesSettings
	if err = rest.DecodeResponse(res, &data); err != nil {
		return ExternalSourcesSettings{}, err
	}

	return data, err
}

func (c *Client) SetExternalSettings(accountID string, source ExternalSource, doNotShowLinkingProposal bool) (ExternalSourcesSettings, error) {
	req, err := rest.NewRequest(
		http.MethodPut,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/settings/externalSources/%s", accountID, source),
		rest.WithJSONBody(&ExternalSourcesSettings{doNotShowLinkingProposal}),
	)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	var data ExternalSourcesSettings
	if err = rest.DecodeResponse(res, &data); err != nil {
		return ExternalSourcesSettings{}, err
	}

	return data, err
}

func (c *Client) GetPrivacySettings(accountID string) (FriendsPrivacySettings, error) {
	req, err := rest.NewRequest(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/settings", accountID),
	)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	var data FriendsPrivacySettings
	if err = rest.DecodeResponse(res, &data); err != nil {
		return FriendsPrivacySettings{}, err
	}

	return data, err
}

func (c *Client) SetPrivacySettings(accountID string, acceptInvites PrivacySettingAcceptInvites, mutualPrivacy PrivacySettingMutualPrivacy) (FriendsPrivacySettings, error) {
	req, err := rest.NewRequest(
		http.MethodPatch,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/settings", accountID),
		rest.WithJSONBody(&FriendsPrivacySettings{acceptInvites, mutualPrivacy}),
	)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	var data FriendsPrivacySettings
	if err = rest.DecodeResponse(res, &data); err != nil {
		return FriendsPrivacySettings{}, err
	}

	return data, err
}
