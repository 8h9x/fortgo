package fortnite

import "time"

type CloudstorageFilePointer struct {
	UniqueFilename string    `json:"uniqueFilename"`
	Filename       string    `json:"filename"`
	Hash           string    `json:"hash"`
	Hash256        string    `json:"hash256"`
	Length         int       `json:"length"`
	ContentType    string    `json:"contentType"`
	Uploaded       time.Time `json:"uploaded"`
	StorageType    string    `json:"storageType"`
	StorageIDs     struct {
		DSS string `json:"DSS"`
	} `json:"storageIds"`
}

type CloudstorageFilePointerSystem struct {
	CloudstorageFilePointer
	DoNotCache bool `json:"doNotCache"`
}

type CloudstorageFilePointerUser struct {
	CloudstorageFilePointer
	Metadata struct {
	} `json:"metadata"`
	AccountID string `json:"accountId"`
}

type CloudstorageConfigTransport struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	AppName        string `json:"appName"`
	IsEnabled      bool   `json:"isEnabled"`
	IsRequired     bool   `json:"isRequired"`
	IsPrimary      bool   `json:"isPrimary"`
	TimeoutSeconds int    `json:"timeoutSeconds"`
	Priority       int    `json:"priority"`
}

type CloudstorageConfig struct {
	LastUpdated        time.Time                              `json:"lastUpdated"`
	DisableV2          bool                                   `json:"disableV2"`
	IsAuthenticated    bool                                   `json:"isAuthenticated"`
	EnumerateFilesPath string                                 `json:"enumerateFilesPath"`
	EnableMigration    bool                                   `json:"enableMigration"`
	EnableWrites       bool                                   `json:"enableWrites"`
	EpicAppName        string                                 `json:"epicAppName"`
	Transports         map[string]CloudstorageConfigTransport `json:"transports"`
	// Transports         struct {
	// 	McpProxyTransport     CloudstorageConfigTransport `json:"McpProxyTransport"`
	// 	McpSignatoryTransport CloudstorageConfigTransport `json:"McpSignatoryTransport"`
	// 	DssDirectTransport    CloudstorageConfigTransport `json:"DssDirectTransport"`
	// } `json:"transports"`
}

type UpdateUserFilePointerPayload struct {
	ModuleName string `json:"moduleName"`
	Pointer    string `json:"pointer"`
	FileHash   string `json:"fileHash"`
	Length     int    `json:"length"`
}
