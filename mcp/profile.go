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

type AthenaProfile Profile[AthenaProfileStats, map[string]any]
type CreativeProfile Profile[CreativeProfileStats, map[string]any]
type CampaignProfile Profile[CampaignProfileStats, CampaignNotification]
type CommonPublicProfile Profile[CommonPublicProfileStats, map[string]any]
type CollectionsProfile Profile[CollectionsProfileStats, map[string]any]
type CommonCoreProfile Profile[CommonCoreProfileStats, map[string]any]
type MetadataProfile Profile[MetadataProfileStats, map[string]any]
type CollectionBookPeople0Profile Profile[CollectionBookPeopleProfileStats, map[string]any]
type CollectionBookSchematics0Profile Profile[CollectionBookSchematicsProfileStats, map[string]any]
type Outpost0Profile Profile[OutpostProfileStats, map[string]any]
type Theater0Profile Profile[Theater0ProfileStats, map[string]any]
type Theater1Profile Profile[Theater1ProfileStats, map[string]any]
type Theater2Profile Profile[Theater2ProfileStats, map[string]any]
type RecycleBinProfile Profile[RecycleBinProfileStats, map[string]any]
type ProtoJunoProfile Profile[ProtoJunoProfileStats, map[string]any]
type Profile0Profile Profile[Profile0ProfileStats, map[string]any]
