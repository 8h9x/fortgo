package account

import "time"

type TokenType string

const (
	TokenTypeDefault TokenType = "*"
	TokenTypeEG1     TokenType = "eg1"
	TokenTypeEP1     TokenType = "ep1"
	TokenTypeEpicID  TokenType = "epic_id"
	TokenTypeIDToken TokenType = "id_token"
)

type GrantType string

const (
	GrantTypeAuthorizationCode GrantType = "authorization_code"
	GrantTypeClientCredentials GrantType = "client_credentials"
	GrantTypeContinuationToken GrantType = "continuation_token"
	GrantTypeDeviceAuth        GrantType = "device_auth"
	GrantTypeDeviceCode        GrantType = "device_code"
	GrantTypeExchangeCode      GrantType = "exchange_code"
	GrantTypeExternalAuth      GrantType = "exteral_auth"
	GrantTypeOTP               GrantType = "otp"
	GrantTypePassword          GrantType = "password"
	GrantTypeRefreshToken      GrantType = "refresh_token"
	GrantTypeTokenToToken      GrantType = "token_to_token"
)

type GetPublicKeyResponse struct {
	KTY string `json:"kty"`
	E   string `json:"e"`
	KID string `json:"kid"`
	N   string `json:"n"`
}

type ExternalAuthType string

const (
	ExternalAuthTypeApple    ExternalAuthType = "apple"
	ExternalAuthTypeFacebook ExternalAuthType = "facebook"
	ExternalAuthTypeGithub   ExternalAuthType = "github"
	ExternalAuthTypeGoogle   ExternalAuthType = "google"
	ExternalAuthTypeInternal ExternalAuthType = "internal"
	ExternalAuthTypeNintendo ExternalAuthType = "nintendo"
	ExternalAuthTypePSN      ExternalAuthType = "psn"
	ExternalAuthTypeSteam    ExternalAuthType = "steam"
	ExternalAuthTypeTwitch   ExternalAuthType = "twitch"
	ExternalAuthTypeUbisoft  ExternalAuthType = "ubisoft"
	ExternalAuthTypeVK       ExternalAuthType = "vk"
	ExternalAuthTypeXBL      ExternalAuthType = "xbl"
)

type ExternalAuthID struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type ExternalAuth struct {
	AccountID               string           `json:"accountId"`
	Type                    string           `json:"type"`
	ExternalAuthID          string           `json:"externalAuthId"`
	ExternalAuthIDType      string           `json:"externalAuthIdType"`
	ExternalAuthSecondaryID string           `json:"externalAuthSecondaryId"`
	ExternalDisplayName     string           `json:"externalDisplayName"`
	Avatar                  string           `json:"avatar"`
	AuthIDs                 []ExternalAuthID `json:"authIds"`
	// DateAdded               time.Time `json:"dateAdded"`
	// RegionInfo              string `json:"regionInfo"`
}

type GetUserResponse struct {
	ID            string                            `json:"id"`
	DisplayName   string                            `json:"displayName"`
	ExternalAuths map[ExternalAuthType]ExternalAuth `json:"externalAuths"`
}

type GetUserResponseExtended struct {
	GetUserResponse
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

type TokenResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	ExpiresAt        string `json:"expires_at"`
	TokenType        string `json:"token_type"`
	RefreshToken     string `json:"refresh_token,omitempty"`
	RefreshExpires   int64  `json:"refresh_expires,omitempty"`
	RefreshExpiresAt string `json:"refresh_expires_at,omitempty"`
	AccountID        string `json:"account_id,omitempty"`
	ClientID         string `json:"client_id"`
	InternalClient   bool   `json:"internal_client"`
	ClientService    string `json:"client_service"`
	DisplayName      string `json:"displayName,omitempty"`
	App              string `json:"app,omitempty"`
	InAppID          string `json:"in_app_id,omitempty"`
	ProductID        string `json:"product_id"`
	ApplicationID    string `json:"application_id"`
}

type getUsersByExternalDisplayNameBulkPayload struct {
	AuthType     ExternalAuthType `json:"authType"`
	DisplayNames []string         `json:"displayNames"`
}

type getUsersByExternalIDBulkPayload struct {
	AuthType ExternalAuthType `json:"authType"`
	IDs      []string         `json:"ids"`
}

type TokenPayloadAuthorizationCode struct {
	Code string `json:"code"`
}

type TokenPayloadClientCredentials struct{}

type TokenPayloadContinuationToken struct {
	ContinuationToken string `json:"continuation_token"`
}

type TokenPayloadDeviceAuth struct {
	AccountID string `json:"account_id"`
	DeviceID  string `json:"device_id"`
	Secret    string `json:"secret"`
}

type TokenPayloadDeviceCode struct {
	DeviceCode string `json:"device_code"`
}

type TokenPayloadExchangeCode struct {
	ExchangeCode string `json:"exchange_code"`
}

type TokenPayloadExternalAuth struct {
	ExternalAuthToken string `json:"external_auth_token"`
}

type TokenPayloadOTP struct {
	OTP string `json:"otp"`
}

type TokenPayloadPassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenPayloadRefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenPayloadTokenToToken struct {
	AccessToken string `json:"access_token"`
}

type TokenPayload interface {
	TokenPayloadAuthorizationCode | TokenPayloadClientCredentials | TokenPayloadContinuationToken |
		TokenPayloadDeviceAuth | TokenPayloadDeviceCode | TokenPayloadExchangeCode | TokenPayloadExternalAuth |
		TokenPayloadOTP | TokenPayloadPassword | TokenPayloadRefreshToken | TokenPayloadTokenToToken
}

type VerifyTokenResponsePermission struct {
	Resource string `json:"resource"`
	Action   string `json:"action"`
}

type VerifyTokenResponse struct {
	Token          string                          `json:"token"`
	SessionID      string                          `json:"session_id"`
	TokenType      string                          `json:"token_type"`
	ClientID       string                          `json:"client_id"`
	InternalClient bool                            `json:"internal_client"`
	ClientService  string                          `json:"client_service"`
	AccountID      string                          `json:"account_id"`
	ExpiresIn      int                             `json:"expires_in"`
	ExpiresAt      time.Time                       `json:"expires_at"`
	AuthMethod     string                          `json:"auth_method"`
	DisplayName    string                          `json:"display_name"`
	App            string                          `json:"app"`
	InAppID        string                          `json:"in_app_id"`
	DeviceID       string                          `json:"device_id"`
	Scope          []string                        `json:"scope"`
	ProductID      string                          `json:"product_id"`
	SandboxID      string                          `json:"sandbox_id"`
	DeploymentID   string                          `json:"deployment_id"`
	ApplicationID  string                          `json:"application_id"`
	ACR            string                          `json:"acr"`
	AuthTime       time.Time                       `json:"auth_time"`
	Perms          []VerifyTokenResponsePermission `json:"perms"`
}

type DeviceAuthResponse struct {
	AccountID string `json:"accountId"`
	Created   struct {
		DateTime  time.Time `json:"dateTime"`
		IpAddress string    `json:"ipAddress"`
		Location  string    `json:"location"`
	} `json:"created"`
	DeviceID  string `json:"deviceId"`
	Secret    string `json:"secret"`
	UserAgent string `json:"userAgent"`
}

type ExchangeResponse struct {
	Code             string `json:"code"`
	CreatingClientID string `json:"creatingClientId"`
	ExpiresInSeconds int    `json:"expiresInSeconds"`
}

type GetDeviceCodeResponse struct {
	UserCode                string `json:"user_code"`
	DeviceCode              string `json:"device_code"`
	VerificationURI         string `json:"verification_uri"`
	VerificationURIComplete string `json:"verification_uri_complete"`
	Prompt                  string `json:"prompt"`
	ExpiresIn               int    `json:"expires_in"`
	Interval                int    `json:"interval"`
	ClientID                string `json:"client_id"`
}

type AuthClient struct {
	ID     string
	Secret string
}
