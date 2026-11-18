package friends

import (
	"encoding/json"
	"time"
)

type Friend struct {
	AccountID string    `json:"accountId"`
	Groups    []any     `json:"groups"`
	Mutual    int       `json:"mutual"`
	Alias     string    `json:"alias"`
	Note      string    `json:"note"`
	Favorite  bool      `json:"favorite"`
	Created   time.Time `json:"created"`
}

type SuggestedFriend json.RawMessage

type Relation json.RawMessage

type BlockedPlayer struct {
	AccountID string    `json:"accountId"`
	Created   time.Time `json:"created"`
}

type PrivacySettingAcceptInvites string

const (
	PrivacySettingAcceptInvitesPublic  PrivacySettingAcceptInvites = "PUBLIC"
	PrivacySettingAcceptInvitesMutual  PrivacySettingAcceptInvites = "FRIENDS_OF_FRIENDS"
	PrivacySettingAcceptInvitesPrivate PrivacySettingAcceptInvites = "PRIVATE"
)

type PrivacySettingMutualPrivacy string

const (
	PrivacySettingMutualPrivacyAll     PrivacySettingMutualPrivacy = "ALL"
	PrivacySettingMutualPrivacyFriends PrivacySettingMutualPrivacy = "FRIENDS"
	PrivacySettingMutualPrivacyNone    PrivacySettingMutualPrivacy = "NONE"
)

type FriendsPrivacySettings struct {
	AcceptInvites PrivacySettingAcceptInvites `json:"acceptInvites"`
	MutualPrivacy PrivacySettingMutualPrivacy `json:"mutualPrivacy"`
}

type FriendsLimits struct {
	Incoming bool `json:"incoming"`
	Outgoing bool `json:"outgoing"`
	Accepted bool `json:"accepted"`
}

type FriendSummaryResponse struct {
	Friends       []Friend               `json:"friends"`
	Incoming      []Friend               `json:"incoming"`
	Outgoing      []Friend               `json:"outgoing"`
	Suggested     []any                  `json:"suggested"` // possibly Friend
	Blocklist     []BlockedPlayer        `json:"blocklist"`
	Settings      FriendsPrivacySettings `json:"settings"`
	LimitsReached FriendsLimits          `json:"limitsReached"`
}

type ParentalControlsPinPayload struct {
	Pin string `json:"pin,omitempty"`
}

type ExternalSource string

const (
	ExternalSourceDefault ExternalSource = "default"
	ExternalSourceSteam   ExternalSource = "steam"
)

type ExternalSourcesSettings struct {
	DoNotShowLinkingProposal bool `json:"doNotShowLinkingProposal"`
}
