package mcp

import (
	"fmt"
	"time"
)

type Profile[ST ProfileStatsType, NT ProfileNotificationType] struct {
	ProfileRevision            int                 `json:"profileRevision"`
	ProfileID                  string              `json:"profileId"`
	ProfileChangesBaseRevision int                 `json:"profileChangesBaseRevision"`
	ProfileChanges             []ProfileChange[ST] `json:"profileChanges"`
	ProfileCommandRevision     int                 `json:"profileCommandRevision"`
	CreationTime               time.Time           `json:"creationTime"`
	ServerTime                 time.Time           `json:"serverTime"`
	ResponseVersion            int                 `json:"responseVersion"`
	Notifications              []NT                `json:"notifications,omitempty"`
}

func (p *Profile[ST, NT]) ReadItem(templateID string, v any) error {
	for _, change := range p.ProfileChanges {
		if item, exists := change.Profile.Items[templateID]; exists {
			return item.ReadInto(v)
		}
	}
	return fmt.Errorf("item with templateID %s not found", templateID)
}

type AthenaProfile Profile[AthenaProfileStats, []any]
type CampaignProfile Profile[CampaignProfileStats, []any]
