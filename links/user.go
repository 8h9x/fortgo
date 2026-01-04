package links

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) CreateUserMnemonic(namespace string, accountID string, payload CreateUserMnemonicPayload) (MnemonicDataWithActivationHistory, error) {
	req, err := request.MakeRequest(
		http.MethodPost,
		consts.LinksService,
		fmt.Sprintf("links/api/%s/author/%s", namespace, accountID),
		request.WithBearerToken(c.Credentials.AccessToken),
		request.WithJSONBody(payload),
	)
	if err != nil {
		return MnemonicDataWithActivationHistory{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return MnemonicDataWithActivationHistory{}, err
	}

	resp, err := request.ParseResponse[MnemonicDataWithActivationHistory](res)
	if err != nil {
		return MnemonicDataWithActivationHistory{}, err
	}

	return resp.Data, nil
}
