package publickey

import "time"

type registerKeyPayload struct {
	Algorithm string `json:"algorithm"`
	PublicKey string `json:"key"`
}

type RegisterPublicKeyResponse struct {
	Key        string    `json:"key"`
	AccountID  string    `json:"account_id"`
	KeyGUID    string    `json:"key_guid"`
	KID        string    `json:"kid"`
	Expiration time.Time `json:"expiration"`
	JWT        string    `json:"jwt"`
	Type       string    `json:"type"`
}