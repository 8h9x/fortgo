package fortgo

import (
	"github.com/8h9x/fortgo/account"
	"github.com/8h9x/fortgo/auth"
)

func (c *Client) GetExchangeCode() (auth.ExchangeResponse, error) {
	return auth.GetExchangeCode(c.HTTPClient, c.CredentialsMap[c.ClientID])
}

func (c *Client) CreateDeviceAuth() (auth.DeviceAuthResponse, error) {
	return auth.CreateDeviceAuth(c.HTTPClient, c.CredentialsMap[c.ClientID])
}

func (c *Client) FetchMe() (account.FetchUserResponseExtended, error) {
	return c.AccountService.FetchUserByID(c.CurrentCredentials().AccountID)
}
