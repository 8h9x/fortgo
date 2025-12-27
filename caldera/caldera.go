package caldera

import (
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetAnticheatProvider(accountID string, exchangeCode string, epicApp string, testMode bool, cloudGamingProvider CloudGamingProvider) (AnticheatProviderResponse, error) {
	payload := &GetAnticheatProviderPayload{
		AccountID:    accountID,
		ExchangeCode: exchangeCode,
		TestMode:     testMode,
		EpicApp:      epicApp,
		Nvidia:       false,
		Luna:         false,
		Salmon:       false,
	}

	if cloudGamingProvider == CloudGamingProviderGeforceNow {
		payload.Nvidia = true
	} else if cloudGamingProvider == CloudGamingProviderAmazonLuna {
		payload.Luna = true
	} else if cloudGamingProvider == CloudGamingProviderSalmon {
		payload.Salmon = true
	}

	req, err := request.MakeRequest(
		http.MethodPost,
		consts.CalderaService,
		"caldera/api/v1/launcher/racp",
		request.WithJSONBody(payload),
	)
	if err != nil {
		return AnticheatProviderResponse{}, err
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return AnticheatProviderResponse{}, err
	}

	res, err := request.ParseResponse[AnticheatProviderResponse](resp)
	if err != nil {
		return AnticheatProviderResponse{}, err
	}

	return res.Data, nil
}
