package fortgo

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/account"
	"github.com/8h9x/fortgo/auth"
	"github.com/8h9x/fortgo/fortnite"
	"github.com/8h9x/fortgo/links"
)

type Client struct {
	HTTPClient     *http.Client
	Header         http.Header
	ClientID       string
	CredentialsMap map[string]auth.TokenResponse

	Accounts *account.Client
	Fortnite *fortnite.Client
	Links *links.Client
}

func NewClient(httpClient *http.Client, initCredentials auth.TokenResponse) (*Client, error) {
	client := &Client{
		HTTPClient:     httpClient,
		Header:         make(http.Header),
		ClientID:       initCredentials.ClientID,
		CredentialsMap: make(map[string]auth.TokenResponse),
		Accounts: account.NewClient(httpClient, &initCredentials),
		Fortnite: fortnite.NewClient(httpClient, &initCredentials),
		Links: links.NewClient(httpClient, &initCredentials),
	}

	client.CredentialsMap[initCredentials.ClientID] = initCredentials

	_, err := auth.VerifyToken(httpClient, initCredentials.AccessToken, false)
	if err != nil {
		return client, err
	}

	mcpVersionData, err := client.Fortnite.GetMCPVersion()
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
	res, err := c.Links.FetchMnemonicInfo("fn", mnemonic, mnemonicType, version)
	if err != nil {
		return Mnemonic{}, err
	}

	return Mnemonic{c, 	res}, nil
}

func (c *Client) ComposeProfileOperation(operation string, profileID string, payload string) (*http.Response, error) {
	return c.Fortnite.ComposeProfileOperation(c.CurrentCredentials().AccountID, operation, profileID, payload)
}

func (c *Client) ProfileOperation(operation string, profileID string, payload any) (*http.Response, error) {
	return c.Fortnite.ProfileOperation(c.CurrentCredentials().AccountID, operation, profileID, payload)
}