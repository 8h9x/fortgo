package fortnite

import "time"

type MCPVersionResponse struct {
	App                       string    `json:"app"`
	ServerDate                time.Time `json:"serverDate"`
	OverridePropertiesVersion string    `json:"overridePropertiesVersion"`
	CLN                       string    `json:"cln"`
	Build                     string    `json:"build"`
	ModuleName                string    `json:"moduleName"`
	BuildDate                 time.Time `json:"buildDate"`
	Version                   string    `json:"version"`
	Branch                    string    `json:"branch"`
	Modules                   struct {
		EpicLightSwitchAccessControlCore struct {
			CLN       string    `json:"cln"`
			Build     string    `json:"build"`
			BuildDate time.Time `json:"buildDate"`
			Version   string    `json:"version"`
			Branch    string    `json:"branch"`
		} `json:"Epic-LightSwitch-AccessControlCore"`
		EpicCommonCore struct {
			CLN       string    `json:"cln"`
			Build     string    `json:"build"`
			BuildDate time.Time `json:"buildDate"`
			Version   string    `json:"version"`
			Branch    string    `json:"branch"`
		} `json:"epic-common-core"`
	} `json:"modules"`
}
