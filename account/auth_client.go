package account

import (
	"encoding/base64"

	"github.com/8h9x/fortgo/consts"
)

type AuthClient struct {
	ID     string
	Secret string
}

func NewAuthClient(clientId, secret string) *AuthClient {
	return &AuthClient{
		ID:     clientId,
		Secret: secret,
	}
}

func (ac *AuthClient) String() string {
	return ac.ID + ":" + ac.Secret
}

func (ac *AuthClient) Base64() string {
	return base64.StdEncoding.EncodeToString([]byte(ac.String()))
}

var FortnitePS4USClient = NewAuthClient(consts.FortnitePS4USClientID, consts.FortnitePS4USClientSecret)
var FortniteNewIOSClient = NewAuthClient(consts.FortniteNewIOSClientID, consts.FortniteNewIOSClientSecret)
