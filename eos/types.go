package eos

import "time"

type Token struct {
	AccessToken      string    `json:"access_token"`
	AccountID        string    `json:"account_id,omitempty"`
	ApplicationID    string    `json:"application_id"`
	ClientID         string    `json:"client_id"`
	ExpiresAt        time.Time `json:"expires_at"`
	ExpiresIn        int       `json:"expires_in"`
	RefreshExpiresAt time.Time `json:"refresh_expires_at,omitempty"`
	RefreshExpiresIn int       `json:"refresh_expires_in,omitempty"`
	RefreshToken     string    `json:"refresh_token,omitempty"`
	Scope            string    `json:"scope,omitempty"`
	TokenType        string    `json:"token_type"`
}

type ExchangeResponse struct {
	Code             string `json:"code"`
	CreatingClientID string `json:"creatingClientId"`
	ExpiresInSeconds int    `json:"expiresInSeconds"`
}

type GetDeviceCodeResponse struct {
	UserCode                string `json:"user_code"`
	DeviceCode              string `json:"device_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	Prompt                  string `json:"prompt"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
	ClientID                string `json:"client_id"`
}

type CreateConversationPayload struct {
	Title   string   `json:"title"`
	Type    string   `json:"type"`
	Members []string `json:"members"`
}

type Conversation struct {
	ConversationID string    `json:"conversationId"`
	DateCreated    time.Time `json:"dateCreated"`
	IsActive       bool      `json:"isActive"`
	IsReportable   bool      `json:"isReportable"`
	Members        []string  `json:"members"`
}

type ConversationSummary struct {
	ConversationCount           int  `json:"conversationCount"`
	UnreadMessages              int  `json:"unreadMessages"`
	UnreadPendingInvites        int  `json:"unreadPendingInvites"`
	HasMoreUnreadMessages       bool `json:"hasMoreUnreadMessages"`
	HasMoreUnreadPendingInvites bool `json:"hasMoreUnreadPendingInvites"`
}
