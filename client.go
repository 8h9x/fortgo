package fortgo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/account"
	"github.com/8h9x/fortgo/avatars"
	"github.com/8h9x/fortgo/caldera"
	"github.com/8h9x/fortgo/eos"
	"github.com/8h9x/fortgo/fortnite"
	"github.com/8h9x/fortgo/friends"
	"github.com/8h9x/fortgo/fulfillment"
	"github.com/8h9x/fortgo/links"
	"github.com/8h9x/fortgo/locker"
	"github.com/8h9x/fortgo/party"
	"github.com/8h9x/fortgo/publickey"
	"github.com/8h9x/fortgo/usersearch"
)

type config struct {
	httpClient *http.Client
	authClient *account.AuthClient
	tokenGrant account.TokenGrant
}

type Client struct {
	Me *ClientUser

	config         *config
	Header         http.Header
	CredentialsMap map[string]*account.Token
	OnTokenRefresh OnTokenRefresh

	tokenSource    *ClientTokenSource
	tokenSourceEOS *ClientTokenSourceEOS

	AccountService     *account.Client
	AvatarService      *avatars.Client
	CalderaService     *caldera.Client
	EOS                *eos.Client
	FortniteService    *fortnite.Client
	FriendService      *friends.Client
	FulfillmentService *fulfillment.Client
	LinkService        *links.Client
	LockerService      *locker.Client
	PartyService       *party.Client
	PublicKeyService   *publickey.Client
	UserSearchService  *usersearch.Client
}

type ClientTokenSource struct {
	token *account.Token
}

func (s *ClientTokenSource) Token(ctx context.Context) (string, error) {
	return s.token.AccessToken, nil
}

type ClientTokenSourceEOS struct {
	token *eos.Token
}

func (s *ClientTokenSourceEOS) Token(ctx context.Context) (string, error) {
	return s.token.AccessToken, nil
}

func NewClient(opts ...Option) *Client {
	tokenSource := &ClientTokenSource{}
	tokenSourceEOS := &ClientTokenSourceEOS{}

	cfg := &config{
		httpClient: NewHTTPClient(tokenSource),
		authClient: account.FortnitePS4USClient,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	client := &Client{
		config:         cfg,
		Header:         make(http.Header),
		CredentialsMap: make(map[string]*account.Token),
		tokenSource:    tokenSource,
		tokenSourceEOS: tokenSourceEOS,
	}

	//	client.config.httpClient.Timeout = 30 * time.Second
	//	client.config.httpClient.Transport = &authenticatedTransport{src: src, base: http.DefaultTransport}

	client.Me = &ClientUser{
		client: client,
	}

	client.AccountService = account.New(client.config.httpClient)
	client.AvatarService = avatars.New(client.config.httpClient)
	client.CalderaService = caldera.New(client.config.httpClient)
	client.EOS = eos.New(NewHTTPClient(tokenSourceEOS))
	client.FortniteService = fortnite.New(client.config.httpClient)
	client.FriendService = friends.New(client.config.httpClient)
	client.FulfillmentService = fulfillment.New(client.config.httpClient)
	client.LinkService = links.New(client.config.httpClient)
	client.LockerService = locker.New(client.config.httpClient)
	client.PartyService = party.New(client.config.httpClient)
	client.PublicKeyService = publickey.New(client.config.httpClient)
	client.UserSearchService = usersearch.New(client.config.httpClient)

	return client
}

func (c *Client) AuthClient() *account.AuthClient {
	return c.config.authClient
}

func (c *Client) ConnectWithToken(token account.Token) error {
	c.tokenSource.token = &token

	c.CredentialsMap[c.config.authClient.ID] = &token

	if _, err := c.AccountService.VerifyToken(false); err != nil {
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

func (c *Client) Connect() error {
	token, err := c.AccountService.GetToken(
		c.config.authClient,
		c.config.tokenGrant,
		account.TokenTypeEG1,
	)
	if err != nil {
		return fmt.Errorf("get token: %w", err)
	}

	return c.ConnectWithToken(token)
}

func (c *Client) CurrentCredentials() *account.Token {
	credentials := c.CredentialsMap[c.config.authClient.ID]

	_, err := c.AccountService.VerifyToken(false)
	if err != nil {
		token, err := c.AccountService.GetToken(
			c.AuthClient(),
			account.TokenGrantRefreshToken(credentials.RefreshToken),
			account.TokenTypeEG1,
		)
		if err != nil {
			println(err)
			// TODO: handle
		}

		*c.CredentialsMap[c.config.authClient.ID] = token
	}

	return c.CredentialsMap[c.config.authClient.ID]
}

func (c *Client) GetMnemonic(mnemonic string, mnemonicType links.MnemonicType, version int) (Mnemonic, error) {
	res, err := c.LinkService.GetMnemonicInfo("fn", mnemonic, mnemonicType, version)
	if err != nil {
		return Mnemonic{}, err
	}

	return Mnemonic{c, res}, nil
}

func (c *Client) ComposeProfileOperation(operation string, profileID string, payload string) (*http.Response, error) {
	return c.FortniteService.ComposeProfileOperation(c.CurrentCredentials().AccountID, operation, profileID, payload)
}

func (c *Client) ProfileOperation(operation string, profileID string, payload any) (*http.Response, error) {
	return c.FortniteService.ProfileOperation(c.CurrentCredentials().AccountID, operation, profileID, payload)
}
