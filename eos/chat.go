package eos

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) CreateConversation(title string, conversationType string, members []string, createIfExists bool) (Conversation, error) {
	query := url.Values{}
	if createIfExists {
		query.Add("createIfExists", "true")
	}

	req, err := rest.NewRequest(
		http.MethodPost,
		consts.EOSService,
		"chat/v1/public/_/conversations",
		rest.WithQueryParamaters(query),
		rest.WithJSONBody(&CreateConversationPayload{
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

	var data Conversation
	if err = rest.DecodeResponse(res, &data); err != nil {
		return Conversation{}, err
	}

	return data, err
}

func (c *Client) GetConversationSummary(accountID string) (ConversationSummary, error) {
	req, err := rest.NewRequest(
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

	var data ConversationSummary
	if err = rest.DecodeResponse(res, &data); err != nil {
		return ConversationSummary{}, err
	}

	return data, err
}
