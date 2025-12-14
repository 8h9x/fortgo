package party

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) Invite(buildID, partyID, friendID string, platform consts.Platform, senderDisplayName string, sendPing bool) error {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	query := url.Values{}
	if sendPing {
		query.Set("sendPing", "true")
	}

	bodyBytes, err := json.Marshal(&InvitePayload{
		BuildID: buildID,
		Platform: string(platform),
		UrnEpicConnTypeS: "game",
		UrnEpicInvitePlatformdataS: "",
		DisplayName: senderDisplayName,
	})
	if err != nil {
		return err
	}

	reqUrl := fmt.Sprintf("%s/party/api/v1/Fortnite/parties/%s/invites/%s?%s", consts.PartyService, partyID, friendID, query.Encode())

	resp, err := request.Request(c.HTTPClient, "POST", reqUrl, headers, string(bodyBytes))
	if err != nil {
		return err
	}

	_, err = request.ResponseParser[interface{}](resp)
	return err
}

func (c *Client) Remove(partyID, memberID string) error {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/party/api/v1/Fortnite/parties/%s/members/%s", consts.PartyService, partyID, memberID)

	resp, err := request.Request(c.HTTPClient, "DELETE", reqUrl, headers, "")
	if err != nil {
		return err
	}

	_, err = request.ResponseParser[interface{}](resp)
	return err
}

func (c *Client) Promote(partyID, memberID string) error {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+c.Credentials.AccessToken)

	reqUrl := fmt.Sprintf("%s/party/api/v1/Fortnite/parties/%s/members/%s/promote", consts.PartyService, partyID, memberID)

	resp, err := request.Request(c.HTTPClient, "POST", reqUrl, headers, "")
	if err != nil {
		return err
	}

	_, err = request.ResponseParser[interface{}](resp)
	return err
}