package avatars

type AccountAvatarData struct {
	AccountID string `json:"accountId"`
	Namespace string `json:"namespace"`
	AvatarID  string `json:"avatarId"`
}

type GetAvatarsResponse []AccountAvatarData
