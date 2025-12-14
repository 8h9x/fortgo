package account

import "time"

type GetPublicKeyResponse struct {
	KTY string `json:"kty"`
	E   string `json:"e"`
	KID string `json:"kid"`
	N   string `json:"n"`
}

type ExternalAuthType string

const (
	ExternalAuthTypeApple   ExternalAuthType = "apple"
	ExternalAuthTypeFacebook   ExternalAuthType = "facebook"
	ExternalAuthTypeGithub   ExternalAuthType = "github"
	ExternalAuthTypeGoogle   ExternalAuthType = "google"
	ExternalAuthTypeInternal   ExternalAuthType = "internal"
	ExternalAuthTypeNintendo ExternalAuthType = "nintendo"
	ExternalAuthTypePSN      ExternalAuthType = "psn"
	ExternalAuthTypeSteam    ExternalAuthType = "steam"
	ExternalAuthTypeTwitch   ExternalAuthType = "twitch"
	ExternalAuthTypeUbisoft      ExternalAuthType = "ubisoft"
	ExternalAuthTypeVK      ExternalAuthType = "vk"
	ExternalAuthTypeXBL      ExternalAuthType = "xbl"
)

type ExternalAuthID struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type ExternalAuth struct {
	AccountID               string `json:"accountId"`
	Type                    string `json:"type"`
	ExternalAuthID          string `json:"externalAuthId"`
	ExternalAuthIDType      string `json:"externalAuthIdType"`
	ExternalAuthSecondaryID string `json:"externalAuthSecondaryId"`
	ExternalDisplayName     string `json:"externalDisplayName"`
	Avatar                  string `json:"avatar"`
	AuthIDs                 []ExternalAuthID `json:"authIds"`
}

type FetchUserResponse struct {
	ID            string `json:"id"`
	DisplayName   string `json:"displayName"`
	ExternalAuths map[ExternalAuthType]ExternalAuth `json:"externalAuths"`
}

type FetchUserResponseExtended struct {
	FetchUserResponse
	Name                         string    `json:"name"`
	Email                        string    `json:"email"`
	FailedLoginAttempts          int       `json:"failedLoginAttempts"`
	LastLogin                    time.Time `json:"lastLogin"`
	NumberOfDisplayNameChanges   int       `json:"numberOfDisplayNameChanges"`
	AgeGroup                     string    `json:"ageGroup"`
	Headless                     bool      `json:"headless"`
	Country                      string    `json:"country"`
	LastName                     string    `json:"lastName"`
	PhoneNumber                  string    `json:"phoneNumber"`
	Company                      string    `json:"company"`
	PreferredLanguage            string    `json:"preferredLanguage"`
	LastDisplayNameChange        time.Time `json:"lastDisplayNameChange"`
	CanUpdateDisplayName         bool      `json:"canUpdateDisplayName"`
	TFAEnabled                   bool      `json:"tfaEnabled"`
	EmailVerified                bool      `json:"emailVerified"`
	MinorVerified                bool      `json:"minorVerified"`
	MinorExpected                bool      `json:"minorExpected"`
	MinorStatus                  string    `json:"minorStatus"`
	GuardianChallengeTimestamp   time.Time `json:"guardianChallengeTimestamp"`
	SIWENotificationEnabled      bool      `json:"siweNotificationEnabled"`
	CabinedMode                  bool      `json:"cabinedMode"`
	HasHashedEmail               bool      `json:"hasHashedEmail"`
	LastReviewedSecuritySettings time.Time `json:"lastReviewedSecuritySettings"`
}

type fetchUsersByExternalDisplayNameBulkPayload struct {
	AuthType     ExternalAuthType   `json:"authType"`
	DisplayNames []string `json:"displayNames"`
}

type fetchUsersByExternalIDBulkPayload struct {
	AuthType     ExternalAuthType   `json:"authType"`
	IDs []string `json:"ids"`
}

