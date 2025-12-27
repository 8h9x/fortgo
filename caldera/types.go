package caldera

type AntiCheatProvider string

const (
	ProviderEasyAntiCheat AntiCheatProvider = "EasyAntiCheat"
)

type CloudGamingProvider string

const (
	CloudGamingProviderGeforceNow CloudGamingProvider = "nvidia"
	CloudGamingProviderAmazonLuna CloudGamingProvider = "luna"
	CloudGamingProviderSalmon     CloudGamingProvider = "salmon"
)

type AnticheatProviderResponse struct {
	Provider string `json:"provider"`
	JWT      string `json:"jwt"`
}

type GetAnticheatProviderPayload struct {
	AccountID    string `json:"account_id"`
	ExchangeCode string `json:"exchange_code"`
	EpicApp      string `json:"epic_app"`
	TestMode     bool   `json:"test_mode"`
	GDKMode      bool   `json:"gdk_mode"`
	Nvidia       bool   `json:"nvidia"`
	Luna         bool   `json:"luna"`
	Salmon       bool   `json:"salmon"`
}
