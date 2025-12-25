package eos

import "time"

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