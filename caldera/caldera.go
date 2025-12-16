package caldera

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/8h9x/fortgo/consts"
	"github.com/8h9x/fortgo/request"
)

func (c *Client) GetAnticheatProvider(accountID string, exchangeCode string, epicApp string, testMode bool, cloudGamingProvider CloudGamingProvider) (AnticheatProviderResponse, error) {
	headers := http.Header{}
	headers.Set("Content-Type", "application/json")

	payload := &GetAnticheatProviderPayload{
		AccountID: accountID,
		ExchangeCode: exchangeCode,
		TestMode: testMode,
		EpicApp: epicApp,
		Nvidia: false,
		Luna: false,
		Salmon: false,
	}

	if cloudGamingProvider == CloudGamingProviderGeforceNow {
		payload.Nvidia = true
	} else if cloudGamingProvider == CloudGamingProviderAmazonLuna {
		payload.Luna = true
	} else if cloudGamingProvider == CloudGamingProviderSalmon {
		payload.Salmon = true
	}

	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return AnticheatProviderResponse{}, err
	}

	reqUrl := fmt.Sprintf("%s/caldera/api/v1/launcher/racp", consts.CalderaService)

	resp, err := request.Request(c.HTTPClient, "POST", reqUrl, headers, string(bodyBytes))
	if err != nil {
		return AnticheatProviderResponse{}, err
	}

	res, err := request.ResponseParser[AnticheatProviderResponse](resp)
	if err != nil {
		return AnticheatProviderResponse{}, err
	}

	return res.Body, nil
}