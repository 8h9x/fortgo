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

type FriendsSettings struct {
	AcceptInvites string `json:"acceptInvites"`
	MutualPrivacy string `json:"mutualPrivacy"`
}

type FriendsLimits struct {
	Incoming bool `json:"incoming"`
	Outgoing bool `json:"outgoing"`
	Accepted bool `json:"accepted"`
}

type FriendSuggestionResponse []SuggestedFriend

type FriendSummaryResponse struct {
	Friends []Friend `json:"friends"`
	Incoming []FriendRequest `json:"incoming"`
	Outgoing []FriendRequest `json:"outgoing"`
	Suggested []any `json:"suggested"` // possibly FriendRequest
	Blocklist []BlockedPlayer `json:"blocklist"`
	Settings FriendsSettings `json:"settings"`
	LimitsReached FriendsLimits `json:"limitsReached"`
}