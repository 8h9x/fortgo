package fortgo

import (
	"fmt"
	"github.com/8h9x/fortgo/auth"
	"github.com/8h9x/fortgo/fortnite"
	"github.com/8h9x/fortgo/links"
	"net/http"
)

type Client struct {
	HttpClient     *http.Client
	Header         http.Header
	ClientID       string
	CredentialsMap map[string]auth.TokenResponse
}

func NewClient(httpClient *http.Client, initCredentials auth.TokenResponse) (*Client, error) {
	client := &Client{
		HttpClient:     httpClient,
		Header:         make(http.Header),
		ClientID:       initCredentials.ClientID,
		CredentialsMap: make(map[string]auth.TokenResponse),
	}

	client.CredentialsMap[initCredentials.ClientID] = initCredentials

	_, err := auth.VerifyToken(httpClient, initCredentials.AccessToken, false)
	if err != nil {
		return client, err
	}

	mcpVersionData, err := fortnite.GetMCPVersion(httpClient)
	if err != nil {
		return client, err
	}

	client.Header.Set("User-Agent", fmt.Sprintf("Fortnite/++Fortnite+%s-CL-%s Windows/10.0.26100.1.256.64bit", mcpVersionData.Branch, mcpVersionData.CLN))

	return client, nil
}


func (c *Client) CurrentCredentials() auth.TokenResponse {
	return c.CredentialsMap[c.ClientID]
}

func (c *Client) GetMnemonic(mnemonic string, mnemonicType links.MnemonicType, version int) (Mnemonic, error) {
	res, err := links.FetchMnemonicInfo(c.HttpClient, c.CurrentCredentials(), "fn", mnemonic, mnemonicType, version)
	if err != nil {
		return Mnemonic{}, err
	}

	return Mnemonic{c, 	res}, nil
}

func (c *Client) ComposeProfileOperation(operation string, profileID string, payload string) (*http.Response, error) {
	return fortnite.ComposeProfileOperation(c.HttpClient, c.CurrentCredentials(), operation, profileID, payload)
}

func (c *Client) ProfileOperation(operation string, profileID string, payload any) (*http.Response, error) {
	return fortnite.ProfileOperation(c.HttpClient, c.CurrentCredentials(), operation, profileID, payload)
}