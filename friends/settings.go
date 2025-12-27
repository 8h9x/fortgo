package friends

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetExternalSettings(accountID string, source ExternalSource) (ExternalSourcesSettings, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/settings/externalSources/%s", accountID, source),
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	resp, err := request.ParseResponse[ExternalSourcesSettings](res)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	return resp.Data, nil
}

func (c *Client) SetExternalSettings(accountID string, source ExternalSource, doNotShowLinkingProposal bool) (ExternalSourcesSettings, error) {
	req, err := request.MakeRequest(
		http.MethodPut,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/settings/externalSources/%s", accountID, source),
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithJSONBody(&ExternalSourcesSettings{doNotShowLinkingProposal}),
	)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	resp, err := request.ParseResponse[ExternalSourcesSettings](res)
	if err != nil {
		return ExternalSourcesSettings{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetPrivacySettings(accountID string) (FriendsPrivacySettings, error) {
	req, err := request.MakeRequest(
		http.MethodGet,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/settings", accountID),
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	resp, err := request.ParseResponse[FriendsPrivacySettings](res)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	return resp.Data, nil
}

func (c *Client) SetPrivacySettings(accountID string, acceptInvites PrivacySettingAcceptInvites, mutualPrivacy PrivacySettingMutualPrivacy) (FriendsPrivacySettings, error) {
	req, err := request.MakeRequest(
		http.MethodPatch,
		consts.FriendsService,
		fmt.Sprintf("friends/api/v1/%s/settings", accountID),
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithJSONBody(&FriendsPrivacySettings{acceptInvites, mutualPrivacy}),
	)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	resp, err := request.ParseResponse[FriendsPrivacySettings](res)
	if err != nil {
		return FriendsPrivacySettings{}, err
	}

	return resp.Data, nil
}
