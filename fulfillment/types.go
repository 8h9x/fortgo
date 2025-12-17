package fulfillment

type RedemetionCodeDetails struct {
	EntitlementID   string `json:"entitlementId"`
	EntitlementName string `json:"entitlementName"`
	ItemID          string `json:"itemId"`
	Namespace       string `json:"namespace"`
	Country         string `json:"country"`
}


type RedeemCodeResponse struct {
	OfferID    string `json:"offerId"`
	AccountID  string `json:"accountId"`
	IdentityID string `json:"identityId"`
	Details    []RedemetionCodeDetails `json:"details"`
}