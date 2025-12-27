package links

import (
	"time"
)

type Namespace string

const (
	NamespaceFortnite Namespace = "fn"
)

type MnemonicType string

const (
	MnemonicTypeBRPlaylist     MnemonicType = "BR:Playlist"
	MnemonicTypeCreativeIsland MnemonicType = "Creative:Island"
)

type ImageMap struct {
	ExtraSmall string `json:"url_xs"`
	Small      string `json:"url_s"`
	Medium     string `json:"url_m"`
	Large      string `json:"url_l"`
	Base       string `json:"url"`
}

type LocalizationMap map[string]string

type GetMnemonicInfoBulkPayload struct {
	Mnemonic string       `json:"mnemonic"`
	Type     MnemonicType `json:"type"`
	Filter   bool         `json:"filter"`
	Version  int          `json:"v"`
}

type ActivationHistoryEntry struct {
	Mnemonic  string    `json:"mnemonic"`
	Version   int       `json:"version"`
	Activated time.Time `json:"activated"`
}

type MnemonicMetadataDynamicXPData struct {
	UniqueGameVersion int    `json:"uniqueGameVersion"`
	CalibrationPhase  string `json:"calibrationPhase"`
}

type MnemonicMetadataMatchmakingData struct {
	UniqueGameVersion int    `json:"uniqueGameVersion"`
	CalibrationPhase  string `json:"calibrationPhase"`
}

type MnemonicMetadata struct {
	ParentSet           string                          `json:"parent_set"`
	FavoriteOverride    string                          `json:"favorite_override"`
	PlayHistoryOverride string                          `json:"play_history_override"`
	AltTitle            LocalizationMap                 `json:"alt_title"`
	ImageURL            string                          `json:"image_url"`
	ProductTag          string                          `json:"product_tag"`
	ImageURLs           ImageMap                        `json:"image_urls"`
	DynamicXP           MnemonicMetadataDynamicXPData   `json:"dynamicXp"`
	Matchmaking         MnemonicMetadataMatchmakingData `json:"matchmaking"`
	VideoVUID           string                          `json:"video_vuid"`
	Title               string                          `json:"title"`
}

type MnemonicMetadataExtended struct {
	MnemonicMetadata
	ExtraVideoVUIDs          []string `json:"extra_video_vuids"`
	LobbyBackgroundImageUrls struct {
		Url string `json:"url"`
	} `json:"lobby_background_image_urls"`
	BlogCategory     string `json:"blog_category"`
	FrontendPlugin   string `json:"frontend_plugin"`
	Locale           string `json:"locale"`
	UnlockConditions struct {
		AllOf struct {
			Conditions []interface{} `json:"conditions"`
		} `json:"allOf"`
		PartyEligibility string `json:"partyEligibility"`
	} `json:"unlockConditions"`
	SubLinkCodes  []string `json:"sub_link_codes"`
	MatchmakingV2 struct {
		RatingType string `json:"ratingType"`
		IsRanked   bool   `json:"isRanked"`
	} `json:"matchmakingV2"`
	AltTagline         LocalizationMap `json:"alt_tagline"`
	MatchmakingPlugins []string        `json:"matchmaking_plugins"`
	Ratings            struct {
		RatingReceivedTime time.Time `json:"rating_received_time"`
		Boards             struct {
			USK struct {
				Descriptors         []interface{} `json:"descriptors"`
				RatingOverridden    bool          `json:"rating_overridden"`
				Rating              string        `json:"rating"`
				InitialRating       string        `json:"initial_rating"`
				InteractiveElements []string      `json:"interactive_elements"`
			} `json:"USK"`
		} `json:"boards"`
	} `json:"ratings"`
	CcuSourceLinks []string `json:"ccu_source_links"`
	FallbackLinks  struct {
		Graceful string `json:"graceful"`
	} `json:"fallback_links"`
	Tagline           string     `json:"tagline"`
	ExtraImageUrls    []ImageMap `json:"extra_image_urls"`
	SquareImageUrls   ImageMap   `json:"square_image_urls"`
	CorrespondingSets struct {
		Unranked string `json:"unranked"`
	} `json:"corresponding_sets"`
	DefaultSubLinkCode string `json:"default_sub_link_code"`
}

type MnemonicData struct {
	Namespace        string           `json:"namespace"`
	AccountID        string           `json:"accountId"`
	CreatorName      string           `json:"creatorName"`
	Mnemonic         string           `json:"mnemonic"`
	LinkType         string           `json:"linkType"`
	Metadata         MnemonicMetadata `json:"metadata"`
	Version          int              `json:"version"`
	Active           bool             `json:"active"`
	Disabled         bool             `json:"disabled"`
	Created          time.Time        `json:"created"`
	Published        time.Time        `json:"published"`
	DescriptionTags  []interface{}    `json:"descriptionTags"`
	ModerationStatus string           `json:"moderationStatus"`
}

type MnemonicDataWithActivationHistory struct {
	MnemonicData
	Metadata          MnemonicMetadataExtended `json:"metadata"`
	DiscoveryIntent   string                   `json:"discoveryIntent"`
	ActivationHistory []ActivationHistoryEntry `json:"activationHistory"`
	LinkCategory      string                   `json:"linkCategory"`
}

type GetRelatedMnemonicsResponse struct {
	ParentLinks []MnemonicData `json:"parentLinks"`
	Links       []MnemonicData `json:"links"`
}
