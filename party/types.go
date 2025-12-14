package party

type InvitePayload struct {
	BuildID         string `json:"urn:epic:cfg:build-id_s"`
	Platform       string `json:"urn:epic:conn:platform_s"`
	UrnEpicConnTypeS           string `json:"urn:epic:conn:type_s"`
	UrnEpicInvitePlatformdataS string `json:"urn:epic:invite:platformdata_s"`
	DisplayName           string `json:"urn:epic:member:dn_s"`
}