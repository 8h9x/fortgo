package friends

import "time"

type Friend struct {
	AccountID string    `json:"accountId"`
	Groups    []any     `json:"groups"`
	Mutual    int       `json:"mutual"`
	Alias     string    `json:"alias"`
	Note      string    `json:"note"`
	Favorite  bool      `json:"favorite"`
	Created   time.Time `json:"created"`
}

type FriendRequest struct {
	AccountID string    `json:"accountId"`
	Mutual    int       `json:"mutual"`
	Favorite  bool      `json:"favorite"`
	Created   time.Time `json:"created"`
}

type SuggestedFriend struct{}

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
	PrivacySettingMutualPrivacyAll      PrivacySettingAcceptInvites = "ALL"
	PrivacySettingMutualPrivacyFriends  PrivacySettingAcceptInvites = "FRIENDS"
	PrivacySettingMutualPrivacyNone     PrivacySettingAcceptInvites = "NONE"
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
	Friends []Friend `json:"friends"`
	Incoming []FriendRequest `json:"incoming"`
	Outgoing []FriendRequest `json:"outgoing"`
	Suggested []any `json:"suggested"` // possibly FriendRequest
	Blocklist []BlockedPlayer `json:"blocklist"`
	Settings FriendsPrivacySettings `json:"settings"`
	LimitsReached FriendsLimits `json:"limitsReached"`
}

type ParentalControlsPinPayload struct {
	Pin string `json:"pin"`
}

type ExternalSource string

const (
	ExternalSourceDefault ExternalSource = "default"
	ExternalSourceSteam   ExternalSource = "steam"
)

type ExternalSourcesSettings struct {
	DoNotShowLinkingProposal bool `json:"doNotShowLinkingProposal"`
}