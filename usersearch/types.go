package usersearch

type Platform string

const (
	PlatformEpic           Platform = "epic"
	PlatformNintendoSwitch Platform = "nsw"
	PlatformPlaystation    Platform = "psn"
	PlatformSteam          Platform = "steam"
	PlatformXbox           Platform = "xbl"
)

type SearchMatchPlatformed struct {
	Value    string   `json:"value"`
	Platform Platform `json:"platform"`
}

type SearchMatch struct {
	AccountID    string                  `json:"accountId"`
	Matches      []SearchMatchPlatformed `json:"matches"`
	MatchType    string                  `json:"matchType"`
	EpicMutuals  int                     `json:"epicMutuals"`
	SortPosition int                     `json:"sortPosition"`
}

type SearchUsersResponse []SearchMatch
