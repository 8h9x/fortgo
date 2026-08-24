package locker

import "time"

type ArchivedStatus string

const (
	ArchivedStatusArchived ArchivedStatus = "ARCHIVED"
)

type CosmeticLoadoutSchema string

const (
	CosmeticLoadoutSchemaCharacter CosmeticLoadoutSchema = "CosmeticLoadout:LoadoutSchema_Character"
	CosmeticLoadoutSchemaEmotes    CosmeticLoadoutSchema = "CosmeticLoadout:LoadoutSchema_Emotes"
	CosmeticLoadoutSchemaJam       CosmeticLoadoutSchema = "CosmeticLoadout:LoadoutSchema_Jam"
	CosmeticLoadoutSchemaMimosa    CosmeticLoadoutSchema = "CosmeticLoadout:LoadoutSchema_Mimosa"
	CosmeticLoadoutSchemaPlatform  CosmeticLoadoutSchema = "CosmeticLoadout:LoadoutSchema_Platform"
	CosmeticLoadoutSchemaSparks    CosmeticLoadoutSchema = "CosmeticLoadout:LoadoutSchema_Sparks"
	CosmeticLoadoutSchemaVehicle   CosmeticLoadoutSchema = "CosmeticLoadout:LoadoutSchema_Vehicle"
	CosmeticLoadoutSchemaWraps     CosmeticLoadoutSchema = "CosmeticLoadout:LoadoutSchema_Wraps"
)

type PresetFavoriteStatus string

const (
	PresetFavoriteStatusEmpty PresetFavoriteStatus = "EMPTY"
)

type ShuffleType string

const (
	ShuffleTypeEnabled  ShuffleType = "ENABLED"
	ShuffleTypeDisabled ShuffleType = "DISABLED"
)

type Loadout struct {
	LoadoutSlots []LoadoutSlot `json:"loadoutSlots"`
	ShuffleType  ShuffleType   `json:"shuffleType"`
}

type ItemCustomization struct {
	ChannelTag     string `json:"channelTag"`
	VariantTag     string `json:"variantTag"`
	AdditionalData string `json:"additionalData,omitempty"`
}

type LoadoutSlot struct {
	SlotTemplate       string              `json:"slotTemplate"`
	EquippedItemID     string              `json:"equippedItemId,omitempty"`
	ItemCustomizations []ItemCustomization `json:"itemCustomizations"`
}

type LoadoutPreset struct {
	DeploymentID string        `json:"deploymentId"`
	AccountID    string        `json:"accountId"`
	LoadoutType  string        `json:"loadoutType"`
	PresetId     string        `json:"presetId"`
	PresetIndex  int           `json:"presetIndex"`
	AthenaItemID string        `json:"athenaItemId"`
	CreationTime time.Time     `json:"creationTime"`
	UpdatedTime  time.Time     `json:"updatedTime"`
	LoadoutSlots []LoadoutSlot `json:"loadoutSlots"`
}

type ActiveLoadoutGroup struct {
	DeploymentID string                              `json:"deploymentId"`
	AccountID    string                              `json:"accountId"`
	AthenaItemId string                              `json:"athenaItemId"`
	CreationTime time.Time                           `json:"creationTime"`
	UpdatedTime  time.Time                           `json:"updatedTime"`
	Loadouts     map[CosmeticLoadoutSchema]Loadout   `json:"loadouts"`
	ShuffleType  ShuffleType                         `json:"shuffleType"`
	Namespace    string                              `json:"namespace"`
}

type LoadoutGroupPreset struct {
	Loadouts             []map[string]Loadout `json:"loadouts"`
	DisplayName          string               `json:"displayName"`
	PresetFavoriteStatus PresetFavoriteStatus `json:"presetFavoriteStatus"`
	DeploymentID         string               `json:"deploymentId"`
	AccountID            string               `json:"accountId"`
	PresetID             string               `json:"presetId"`
	AthenaItemID         string               `json:"athenaItemId"`
	CreationTime         time.Time            `json:"creationTime"`
	UpdatedTime          time.Time            `json:"updatedTime"`
}

type LockerItems struct {
	ActiveLoadoutGroup  ActiveLoadoutGroup   `json:"activeLoadoutGroup"`
	LoadoutGroupPresets []LoadoutGroupPreset `json:"loadoutGroupPresets"`
	LoadoutPresets      []LoadoutPreset      `json:"loadoutPresets"`
}

type UpdateActiveLockerLoadoutPayload struct {
	Loadouts             map[CosmeticLoadoutSchema]Loadout   `json:"loadouts"`
	ShuffleType          ShuffleType                         `json:"shuffleType"`
	EquippedPresetItemID string                              `json:"equippedPresetItemId"`
	AthenaItemID         string                              `json:"athenaItemId"`
	CreationTime         time.Time                           `json:"creationTime"`
}

type ChangeCompanionNamePayload struct {
	CosmeticItemID string `json:"cosmeticItemId"`
	CompanionName  string `json:"companionName"`
}

type LockInImmutableItemPayload struct {
	Variants map[string]ItemCustomization `json:"variants"`
}

type CosmeticItemVariant struct {
	VariantTag string `json:"variantTag"`
	AdditionalData string `json:"additionalData,omitempty"`
}

type CosmeticItem struct {
	TemplateID     string    						`json:"templateId"`
	CosmeticItemID string    						`json:"cosmeticItemId"`
	AthenaItemID   string   						`json:"athenaItemId"`
	UpdatedTime    time.Time 						`json:"updatedTime"`
	ActiveVariants map[string]CosmeticItemVariant	`json:"activeVariants"`
	OwnedVariants  map[string][]CosmeticItemVariant	`json:"ownedVariants,omitempty"`
	CreationTime   time.Time						`json:"creationTime,omitempty"`
	LockedIn       bool      						`json:"lockedIn,omitempty"`
}

type CosmeticData struct {
	CosmeticItems            []CosmeticItem	`json:"cosmeticItems"`
	CosmeticItemAccessTokens []string		`json:"cosmeticItemAccessTokens"`
	RequestTime              time.Time		`json:"requestTime"`
	NextToken                string			`json:"nextToken"`
}
