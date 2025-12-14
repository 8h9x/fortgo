package account

import "time"

type GetPublicKeyResponse struct {
	KTY string `json:"kty"`
	E   string `json:"e"`
	KID string `json:"kid"`
	N   string `json:"n"`
}

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
	ExternalAuths map[string]ExternalAuth `json:"externalAuths"`
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