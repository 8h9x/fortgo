package fortgo

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/account"
	"github.com/8h9x/fortgo/auth"
	"github.com/8h9x/fortgo/avatars"
	"github.com/8h9x/fortgo/caldera"
	"github.com/8h9x/fortgo/fortnite"
	"github.com/8h9x/fortgo/friends"
	"github.com/8h9x/fortgo/fulfillment"
	"github.com/8h9x/fortgo/links"
	"github.com/8h9x/fortgo/locker"
	"github.com/8h9x/fortgo/party"
	"github.com/8h9x/fortgo/usersearch"
)

type Client struct {
	HTTPClient     *http.Client
	Header         http.Header
	ClientID       string
	CredentialsMap map[string]auth.TokenResponse
	OnTokenRefresh OnTokenRefresh

	AccountService *account.Client
	AvatarService *avatars.Client
	CalderaService *caldera.Client
	FortniteService *fortnite.Client
	FriendService *friends.Client
	FulfillmentService *fulfillment.Client
	LinkService *links.Client
	LockerService *locker.Client
	PartyService *party.Client
	UserSearchService *usersearch.Client
}

func NewClient(httpClient *http.Client, credentials auth.TokenResponse) *Client {
	client := &Client{
		HTTPClient:     httpClient,
		Header:         make(http.Header),
		ClientID:       credentials.ClientID,
		CredentialsMap: make(map[string]auth.TokenResponse),
	}

	client.CredentialsMap[credentials.ClientID] = credentials

	client.AccountService = account.NewClient(httpClient, &credentials)
	client.AvatarService = avatars.NewClient(httpClient, &credentials)
	client.CalderaService = caldera.NewClient(httpClient, &credentials)
	client.FortniteService = fortnite.NewClient(httpClient, &credentials)
	client.FriendService = friends.NewClient(httpClient, &credentials)
	client.FulfillmentService = fulfillment.NewClient(httpClient, &credentials)
	client.LinkService = links.NewClient(httpClient, &credentials)
	client.LockerService = locker.NewClient(httpClient, &credentials)
	client.PartyService = party.NewClient(httpClient, &credentials)
	client.UserSearchService = usersearch.NewClient(httpClient, &credentials)

	return client
}

func Startup() {

}

func (c *Client) Connect() error {
	credentials := c.CurrentCredentials()

	if _, err := auth.VerifyToken(c.HTTPClient, credentials.AccessToken, false); err != nil {
		return fmt.Errorf("verify token: %w", err)
	}

	mcpVersionData, err := c.FortniteService.GetMCPVersion()
	if err != nil {
		return fmt.Errorf("fetch mcp version: %w", err)
	}

	c.Header.Set("User-Agent", fmt.Sprintf(
		"Fortnite/++Fortnite+%s-CL-%s Windows/10.0.26100.1.256.64bit",
		mcpVersionData.Branch, mcpVersionData.CLN,
	))

	return nil
}


func (c *Client) CurrentCredentials() auth.TokenResponse {
	credentials := c.CredentialsMap[c.ClientID]

	_, err := auth.VerifyToken(c.HTTPClient, credentials.AccessToken, false)
	if err != nil {
		res, err := auth.Authenticate(c.HTTPClient, auth.FortnitePS4USClient, auth.PayloadRefreshToken{credentials.RefreshToken}, true)
		if err != nil {
			println(err)
			// TODO: handle
		}

		c.CredentialsMap[c.ClientID] = res
	}

	return c.CredentialsMap[c.ClientID]
}

func (c *Client) GetMnemonic(mnemonic string, mnemonicType links.MnemonicType, version int) (Mnemonic, error) {
	res, err := c.LinkService.GetMnemonicInfo("fn", mnemonic, mnemonicType, version)
	if err != nil {
		return Mnemonic{}, err
	}

	return Mnemonic{c, 	res}, nil
}

func (c *Client) ComposeProfileOperation(operation string, profileID string, payload string) (*http.Response, error) {
	return c.FortniteService.ComposeProfileOperation(c.CurrentCredentials().AccountID, operation, profileID, payload)
}

func (c *Client) ProfileOperation(operation string, profileID string, payload any) (*http.Response, error) {
	return c.FortniteService.ProfileOperation(c.CurrentCredentials().AccountID, operation, profileID, payload)
}