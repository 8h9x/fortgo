package fortgo

import (
	"fmt"
	"github.com/8h9x/fortgo/auth"
	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
	"net/http"
	"time"
)

func (c *Client) GetExchangeCode() (auth.ExchangeResponse, error) {
	return auth.GetExchangeCode(c.HttpClient, c.CredentialsMap[c.ClientID])
}

func (c *Client) CreateDeviceAuth() (auth.DeviceAuthResponse, error) {
	return auth.CreateDeviceAuth(c.HttpClient, c.CredentialsMap[c.ClientID])
}

type GetPublicKeyResponse struct {
	KTY string `json:"kty"`
	E   string `json:"e"`
	KID string `json:"kid"`
	N   string `json:"n"`
}

func (c *Client) GetPublicKey(thumbprint string) (GetPublicKeyResponse, error) {
	resp, err := c.Request("GET", consts.AccountService+"/account/api/publickeys/"+thumbprint, nil, "")
	if err != nil {
		return GetPublicKeyResponse{}, err
	}

	res, err := request.ResponseParser[GetPublicKeyResponse](resp)
	if err != nil {
		return GetPublicKeyResponse{}, err
	}

	return res.Body, err
}

type FetchUserResponse struct {
	ID                           string    `json:"id"`
	DisplayName                  string    `json:"displayName"`
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

func (c *Client) FetchUserByID(accountID string) (FetchUserResponse, error) {
	credentials := c.CredentialsMap[c.ClientID]

	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/account/api/public/account/%s", consts.AccountService, accountID), headers, "")
	if err != nil {
		return FetchUserResponse{}, err
	}

	res, err := request.ResponseParser[FetchUserResponse](resp)
	if err != nil {
		return FetchUserResponse{}, err
	}

	return res.Body, err
}

func (c *Client) FetchUserByDisplayName(displayName string) (FetchUserResponse, error) {
	credentials := c.CredentialsMap[c.ClientID]

	headers := http.Header{}
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err := c.Request("GET", fmt.Sprintf("%s/account/api/public/account/displayName/%s", consts.AccountService, displayName), headers, "")
	if err != nil {
		return FetchUserResponse{}, err
	}

	res, err := request.ResponseParser[FetchUserResponse](resp)
	if err != nil {
		return FetchUserResponse{}, err
	}

	return res.Body, err
}

func (c *Client) FetchMe() (FetchUserResponse, error) {
	credentials := c.CredentialsMap[c.ClientID]
	return c.FetchUserByID(credentials.AccountID)
}
