package mcp

import (
	"encoding/json"
	"fmt"
)

type Profile struct {
	ProfileRevision            int                   `json:"profileRevision"`
	ProfileID                  string                `json:"profileId"`
	ProfileChangesBaseRevision int                   `json:"profileChangesBaseRevision"`
	ProfileChanges             []ProfileChange       `json:"profileChanges"`
	ProfileCommandRevision     int                   `json:"profileCommandRevision"`
	CreationTime               string                `json:"creationTime"`
	ServerTime                 string                `json:"serverTime"`
	ResponseVersion            int                   `json:"responseVersion"`
	Notifications              []ProfileNotification `json:"notifications,omitempty"`
}

type ProfileChange struct {
	ChangeType           string                    `json:"changeType"`
	EnableConstructDelta bool                      `json:"enableConstructDelta"`
	Profile              ProfileChangeProfileEntry `json:"profile"`
}

type ProfileChangeProfileEntry struct {
	Created         string                 `json:"created"`
	Updated         string                 `json:"updated"`
	RVN             int                    `json:"rvn"`
	WipeNumber      int                    `json:"wipeNumber"`
	AccountID       string                 `json:"accountId"`
	ProfileID       string                 `json:"profileId"`
	Version         string                 `json:"version"` // migrations
	Stats           ProfileStats           `json:"stats"`
	CommandRevision int                    `json:"commandRevision"`
	ID              string                 `json:"_id"`
	Items           map[string]ProfileItem `json:"items"`
}

type ProfileStatsAttributes = json.RawMessage

type ProfileStats struct {
	Attributes ProfileStatsAttributes `json:"attributes"`
}

func (i *ProfileStats) ReadAttributes(v any) error {
	return json.Unmarshal(i.Attributes, v)
}

func (i *ProfileStats) WriteAttributes(v any) error {
	attrs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	i.Attributes = attrs
	return nil
}

func (i *ProfileItem) GetAttribute(key string, v any) error {
	var m map[string]json.RawMessage
	if err := i.ReadAttributes(&m); err != nil {
		return err
	}
	raw, ok := m[key]
	if !ok {
		return fmt.Errorf("attribute %q not found", key)
	}
	return json.Unmarshal(raw, v)
}

func (i *ProfileItem) SetAttribute(key string, value any) error {
	var m map[string]json.RawMessage
	if err := i.ReadAttributes(&m); err != nil {
		return err
	}
	if m == nil {
		m = make(map[string]json.RawMessage)
	}

	encoded, err := json.Marshal(value)
	if err != nil {
		return err
	}
	m[key] = encoded

	return i.WriteAttributes(m)
}

type ProfileItemAttributes = json.RawMessage

type ProfileItem struct {
	TemplateID string                `json:"templateId"`
	Attributes ProfileItemAttributes `json:"attributes"`
	Quantity   int                   `json:"quantity"`
}

func (i *ProfileItem) ReadAttributes(v any) error {
	return json.Unmarshal(i.Attributes, v)
}

func (i *ProfileItem) WriteAttributes(v any) error {
	attrs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	i.Attributes = attrs
	return nil
}

type ProfileNotification struct {
}
