package eos

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) CreateConversation(title string, conversationType string, members []string, createIfExists bool) (Conversation, error) {
	query := url.Values{}
	if createIfExists {
		query.Add("createIfExists", "true")
	}

	req, err := request.New(
		http.MethodPost,
		consts.EOSService,
		"chat/v1/public/_/conversations",
		request.WithQueryParamaters(query),
		request.WithJSONBody(&CreateConversationPayload{
			Title:   title,
			Type:    conversationType,
			Members: members,
		}),
	)
	if err != nil {
		return Conversation{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return Conversation{}, err
	}

	resp, err := request.DecodeResponse[Conversation](res)
	if err != nil {
		return Conversation{}, err
	}

	return resp.Data, nil
}

func (c *Client) GetConversationSummary(accountID string) (ConversationSummary, error) {
	req, err := request.New(
		http.MethodGet,
		consts.EOSService,
		fmt.Sprintf("chat/v1/public/_/users/%s/summary", accountID),
	)
	if err != nil {
		return ConversationSummary{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return ConversationSummary{}, err
	}

	resp, err := request.DecodeResponse[ConversationSummary](res)
	if err != nil {
		return ConversationSummary{}, err
	}

	return resp.Data, nil
}
