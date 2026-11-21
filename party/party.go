package party

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) Invite(buildID, partyID, friendID string, platform consts.Platform, senderDisplayName string, sendPing bool) error {
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

	req, err := rest.NewRequest(
		http.MethodPost,
		consts.PartyService,
		fmt.Sprintf("party/api/v1/Fortnite/parties/%s/invites/%s", partyID, friendID),
		rest.WithQueryParameters(query),
		rest.WithJSONBody(payload),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if err = rest.DecodeResponse(res, new(any)); err != nil {
		return err
	}

	return nil
}

func (c *Client) Remove(partyID string, memberID string) error {
	req, err := rest.NewRequest(
		http.MethodDelete,
		consts.PartyService,
		fmt.Sprintf("party/api/v1/Fortnite/parties/%s/members/%s", partyID, memberID),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if err = rest.DecodeResponse(res, new(any)); err != nil {
		return err
	}

	return nil
}

func (c *Client) Promote(partyID string, memberID string) error {
	req, err := rest.NewRequest(
		http.MethodPost,
		consts.PartyService,
		fmt.Sprintf("party/api/v1/Fortnite/parties/%s/members/%s/promote", partyID, memberID),
	)
	if err != nil {
		return err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if err = rest.DecodeResponse(res, new(any)); err != nil {
		return err
	}

	return nil
}
