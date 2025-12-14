package locker

import "time"

type DeploymentID string

const (
	DeploymentIDLiveFN = "62a9473a2dca46b29ccf17577fcf42d7"
)

type Loadout struct {
	LoadoutSlots []LoadoutSlot `json:"loadoutSlots"`
	ShuffleType string `json:"shuffleType"`
}

type ItemCustomization struct {
	ChannelTag     string `json:"channelTag"`
	VariantTag     string `json:"variantTag"`
	AdditionalData string `json:"additionalData,omitempty"`
}

type LoadoutSlot struct {
	SlotTemplate       string              `json:"slotTemplate"`
	EquippedItemId     string              `json:"equippedItemId"`
	ItemCustomizations []ItemCustomization `json:"itemCustomizations"`
}

type LoadoutPreset struct {
	DeploymentId string        `json:"deploymentId"`
	AccountId    string        `json:"accountId"`
	LoadoutType  string        `json:"loadoutType"`
	PresetId     string        `json:"presetId"`
	PresetIndex  int           `json:"presetIndex"`
	AthenaItemId string        `json:"athenaItemId"`
	CreationTime time.Time     `json:"creationTime"`
	UpdatedTime  time.Time     `json:"updatedTime"`
	LoadoutSlots []LoadoutSlot `json:"loadoutSlots"`
}

type ActiveLoadoutGroup struct {
	DeploymentId string               `json:"deploymentId"`
	AccountId    string               `json:"accountId"`
	AthenaItemId string               `json:"athenaItemId"`
	CreationTime time.Time            `json:"creationTime"`
	UpdatedTime  time.Time            `json:"updatedTime"`
	Loadouts     []map[string]Loadout `json:"loadouts"`
	ShuffleType  string               `json:"shuffleType"`
}

type LockerItems struct {
	ActiveLoadoutGroup  ActiveLoadoutGroup `json:"activeLoadoutGroup"`
	LoadoutGroupPresets []any              `json:"loadoutGroupPresets"`
	LoadoutPresets      []LoadoutPreset    `json:"loadoutPresets"`
}

type UpdateActiveLockerLoadoutPayload struct {
	Loadouts             []map[string]Loadout `json:"loadouts"`
	ShuffleType          string               `json:"shuffleType"`
	EquippedPresetItemId string               `json:"equippedPresetItemId"`
	AthenaItemId         string               `json:"athenaItemId"`
	CreationTime         time.Time            `json:"creationTime"`
}

type ChangeCompanionNamePayload struct {
	CosmeticItemID string `json:"cosmeticItemId"`
	CompanionName  string `json:"companionName"`
}