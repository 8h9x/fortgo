package links

import (
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) CreateUserMnemonic(namespace string, accountID string, payload CreateUserMnemonicPayload) (MnemonicDataWithActivationHistory, error) {
	req, err := rest.NewRequest(
		http.MethodPost,
		consts.LinksService,
		fmt.Sprintf("links/api/%s/author/%s", namespace, accountID),
		rest.WithJSONBody(payload),
	)
	if err != nil {
		return MnemonicDataWithActivationHistory{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return MnemonicDataWithActivationHistory{}, err
	}

	var data MnemonicDataWithActivationHistory
	if err = rest.DecodeResponse(res, &data); err != nil {
		return MnemonicDataWithActivationHistory{}, err
	}

	return data, err
}
