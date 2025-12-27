package party

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) Invite(buildID string, partyID string, friendID string, platform consts.Platform, senderDisplayName string, sendPing bool) error {
	query := url.Values{}
	if sendPing {
		query.Set("sendPing", "true")
	}

	payload := &InvitePayload{
		BuildID:                    buildID,
		Platform:                   string(platform),
		UrnEpicConnTypeS:           "game",
		UrnEpicInvitePlatformdataS: "",
		DisplayName:                senderDisplayName,
	}

	req, err := request.MakeRequest(
		http.MethodPost,
		consts.PartyService,
		fmt.Sprintf("party/api/v1/Fortnite/parties/%s/invites/%s", partyID, friendID),
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithQueryParamaters(query),
		request.WithJSONBody(payload),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	_, err = request.ParseResponse[any](res)
	return err
}

func (c *Client) Remove(partyID string, memberID string) error {
	req, err := request.MakeRequest(
		http.MethodDelete,
		consts.PartyService,
		fmt.Sprintf("party/api/v1/Fortnite/parties/%s/members/%s", partyID, memberID),
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	_, err = request.ParseResponse[any](res)
	return err
}

func (c *Client) Promote(partyID string, memberID string) error {
	req, err := request.MakeRequest(
		http.MethodPost,
		consts.PartyService,
		fmt.Sprintf("party/api/v1/Fortnite/parties/%s/members/%s/promote", partyID, memberID),
		request.WithBearerToken(c.Credentials.AccessToken),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	_, err = request.ParseResponse[any](res)
	return err
}
