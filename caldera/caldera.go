package caldera

import (
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/rest"
)

func (c *Client) GetAnticheatProvider(accountID string, exchangeCode string, epicApp string, testMode bool, cloudGamingProvider CloudGamingProvider) (AnticheatProviderResponse, error) {
	payload := &GetAnticheatProviderPayload{
		AccountID:    accountID,
		ExchangeCode: exchangeCode,
		EpicApp:      epicApp,
		TestMode:     testMode,
		GDKMode:      false,
		Nvidia:       false,
		Luna:         false,
		Salmon:       false,
	}

	switch cloudGamingProvider {
	case CloudGamingProviderGeforceNow:
		payload.Nvidia = true
	case CloudGamingProviderAmazonLuna:
		payload.Luna = true
	case CloudGamingProviderSalmon:
		payload.Salmon = true
	}

	req, err := rest.NewRequest(
		http.MethodPost,
		consts.CalderaService,
		"caldera/api/v1/launcher/racp",
		rest.WithJSONBody(payload),
	)
	if err != nil {
		return AnticheatProviderResponse{}, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return AnticheatProviderResponse{}, err
	}

	var data AnticheatProviderResponse
	if err = rest.DecodeResponse(res, &data); err != nil {
		return AnticheatProviderResponse{}, err
	}

	return data, err
}
