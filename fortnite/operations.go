package fortnite

type AbandonExpeditionPayload struct {
	ExpeditionID string `json:"expeditionId"`
}

type ActivateConsumablePayload struct {
	TargetItemID    string `json:"targetItemId"`
	TargetAccountID string `json:"targetAccountId"`
}

type AddToCollectionPayload struct {
	Category    string              `json:"category"`
	Variant     string              `json:"variant"`
	ContextTags []string            `json:"contextTags"`
	Properties  interface{}         `json:"properties"`
	SeenState   EFortCollectedState `json:"seenState"`
	Count       int                 `json:"count"`
}

type ApplyVotePayload struct {
	OfferID string `json:"offerId"`
}

type AssignGadgetToLoadoutPayload struct {
	GadgetID  string `json:"gadgetId"`
	LoadoutID string `json:"loadoutId"`
	SlotIndex uint8  `json:"slotIndex"` // either 0 or 1
}

type AssignHeroToLoadoutPayload struct {
	HeroID    string `json:"heroId"`
	LoadoutID string `json:"loadoutId"`
	SlotIndex uint8  `json:"slotIndex"` // either 0 or 1
}

type AssignTeamPerkToLoadoutPayload struct {
	TeamPerkID string `json:"teamPerkId"`
	LoadoutID  string `json:"loadoutId"`
}

type AssignWorkerToSquadPayload struct {
	CharacterID string         `json:"characterId"`
	SquadID     SquadAttribute `json:"squadId"`
	SlotIndex   uint8          `json:"slotIndex"`
}

type AssignWorkerToSquadBatchPayload struct {
	CharacterIDs []string         `json:"characterIds"`
	SquadIDs     []SquadAttribute `json:"squadIds"`
	SlotIndices  []uint8          `json:"slotIndices"`
}

type AthenaPinQuestPayload struct {
	PinnedQuest string `json:"pinnedQuest"`
}

type AthenaRemoveQuestsPayload struct {
	RemovedQuests []string `json:"removedQuests"`
}

type BulkUpdateCollectionsPayload struct {
	Items []AddToCollectionPayload `json:"items"`
}

type CancelOrResumeSubscriptionPayload struct {
	AppStore             string `json:"appStore"`
	UniqueSubscriptionId string `json:"uniqueSubscriptionId"`
	WillAutoRenew        bool   `json:"willAutoRenew"`
}

type ChallengeBundleLevelUpPayload struct {
	BundleIdToLevel string `json:"bundleIdToLevel"`
}

type ClaimCollectedResourcesPayload struct {
	CollectorsToClaim []string `json:"collectorsToClaim"`
}

type ClaimCollectionBookPageRewardsPayload struct {
	PageTemplateID      string `json:"pageTemplateId"`
	SectionID           string `json:"sectionId"`
	SelectedRewardIndex int    `json:"selectedRewardIndex"`
}

type ClaimCollectionBookRewardsPayload struct {
	RequiredXP          int `json:"requiredXp"`
	SelectedRewardIndex int `json:"selectedRewardIndex"`
}

type ClaimImportFriendsRewardPayload struct {
	Network ESocialImportPanelPlatform `json:"network"`
}

type ClaimMFAEnabledPayload struct {
	ClaimForSTW bool `json:"bClaimForStw"`
}

type ClaimQuestRewardPayload struct {
	QuestID             string `json:"questId"`
	SelectedRewardIndex int    `json:"selectedRewardIndex"`
}

type ClaimSubscriptionRewardsPayload struct {
	AppStore             string `json:"appStore"`
	UniqueSubscriptionId string `json:"uniqueSubscriptionId"`
	ReceiptInfo          string `json:"receiptInfo"`
}

type ClearHeroLoadoutPayload struct {
	LoadoutID string `json:"loadoutId"`
}

type ClientQuestLoginPayload struct {
	StreamingAppKey string `json:"streamingAppKey"`
}

type CollectExpeditionPayload struct {
	ExpeditionTemplate string `json:"expeditionTemplate"`
	ExpeditionId       string `json:"expeditionId"`
}

type CompletePlayerSurveyPayload struct {
	SurveyID                 string `json:"surveyId"`
	UpdateAllSurveysMetadata bool   `json:"bUpdateAllSurveysMetadata"`
}

type ConvertItemPayload struct {
	TargetItemID    string `json:"targetItemId"`
	ConversionIndex uint8  `json:"conversionIndex"`
}

type ConvertSlottedItemPayload struct {
	TargetItemID    string `json:"targetItemId"`
	ConversionIndex uint8  `json:"conversionIndex"`
}

type CopyCosmeticLoadoutPayload struct {
	SourceIndex         int    `json:"sourceIndex"`
	TargetIndex         int    `json:"targetIndex"`
	OptNewNameForTarget string `json:"optNewNameForTarget"`
}

type CraftWorldItemPayload struct {
	TargetSchematicItemID string `json:"targetSchematicItemId"`
	NumTimesToCraft       int    `json:"numTimesToCraft"`
	TargetSchematicTier   string `json:"targetSchematicTier"`
}

type DeleteCosmeticLoadoutPayload struct {
	Index                int  `json:"index"`
	FallbackLoadoutIndex int  `json:"fallbackLoadoutIndex"`
	LeaveNullSlot        bool `json:"leaveNullSlot"`
}

type DestroyWorldItemsPayload struct {
	ItemIDs []string `json:"itemIds"`
}

type DisassembleWorldItemsPayload struct {
	TargetItemIdAndQuantityPairs []struct {
		ItemID   string `json:"itemId"`
		Quantity int    `json:"quantity"`
	} `json:"targetItemIdAndQuantityPairs"`
}

type ExchangeGameCurrencyForBattlePassOfferPayload struct {
	OfferItemIDList []string `json:"offerItemIdList"`
}

type FortRerollDailyQuestPayload struct {
	QuestID string `json:"questId"`
}

type GiftCatalogEntryPayload struct {
	OfferID            string   `json:"offerId"`
	Currency           string   `json:"currency"`
	CurrencySubType    string   `json:"currencySubType"`
	ExpectedTotalPrice int      `json:"expectedTotalPrice"`
	GameContext        string   `json:"gameContext"`
	ReceiverAccountIds []string `json:"receiverAccountIds"`
	GiftWrapTemplateID string   `json:"giftWrapTemplateId"`
	PersonalMessage    string   `json:"personalMessage"`
}

type InitializeTheaterPayload struct {
	TheaterGUID string `json:"theaterGuid"`
}

type IssueFriendCodePayload struct {
	CodeTokenType string `json:"codeTokenType"`
}

type MarkCollectedItemsSeenPayload struct {
	Variants []MarkCollectedItemsSeenPayloadVariants `json:"variants"`
}

type MarkCollectedItemsSeenPayloadVariants struct {
	Category string `json:"category"`
	Variant  string `json:"variant"`
}

type MarkItemSeenPayload struct {
	ItemIDs []string `json:"itemIds"`
}

type MarkNewQuestNotificationSentPayload struct {
	ItemIDs []string `json:"itemIds"`
}

type ModifyQuickbarPayload struct {
	PrimaryQuickbarChoices  []string `json:"primaryQuickbarChoices"`
	SecondaryQuickbarChoice string   `json:"secondaryQuickbarChoice"`
}

type OpenCardPackPayload struct {
	CardPackItemID string `json:"cardPackItemId"`
	SelectionIndex int    `json:"selectionIdx"`
}

type OpenCardPackBatchPayload struct {
	CardPackItemIDs []string `json:"cardPackItemIds"`
}

type PromoteItemPayload struct {
	TargetItemID string `json:"targetItemId"`
}

type PurchaseCatalogEntryPayload struct {
	OfferID            string `json:"offerId"`
	PurchaseQuantity   int    `json:"purchaseQuantity"`
	Currency           string `json:"currency"`
	CurrencySubType    string `json:"currencySubType"`
	ExpectedTotalPrice int    `json:"expectedTotalPrice"`
	GameContext        string `json:"gameContext"`
}

type PurchaseMultipleCatalogEntriesPayload struct {
	PurchaseInfoList []PurchaseCatalogEntryPayload `json:"purchaseInfoList"`
}

type PurchaseOrUpgradeHomebaseNodePayload struct {
	NodeID string `json:"nodeId"`
}

type PurchaseResearchStatUpgradePayload struct {
	StatID string `json:"statId"`
}

type RecycleItemPayload struct {
	TargetItemID string `json:"targetItemId"`
}

type RecycleItemBatchPayload struct {
	TargetItemIDs []string `json:"targetItemIds"`
}

type RedeemRealMoneyPurchasesPayload struct {
	AppStore              string                                `json:"appStore"`
	AuthTokens            []string                              `json:"authTokens"`
	ReceiptIDs            []string                              `json:"receiptIds"`
	RefreshType           RealMoneyPurchaseRefreshType          `json:"refreshType"`
	VerifierModeOverride  RealMoneyPurchaseVerifierModeOverride `json:"verifierModeOverride"`
	PurchaseCorrelationID string                                `json:"purchaseCorrelationId"`
}

type RefundItemPayload struct {
	TargetItemID string `json:"targetItemId"`
}

type RefundMtxPurchasePayload struct {
	PurchaseID  string `json:"purchaseId"`
	QuickReturn bool   `json:"quickReturn"`
	GameContext string `json:"gameContext"`
}

type RemoveGiftBoxPayload struct {
	GiftBoxItemIDs []string `json:"giftBoxItemIds"`
}

type RequestRestedStateIncreasePayload struct {
	TimeToCompensateFor    int `json:"timeToCompensateFor"`
	RestedXPGenAccumulated int `json:"restedXpGenAccumulated"`
}

type ResearchItemFromCollectionBookPayload struct {
	TemplateID string `json:"templateId"`
}

type RespecAlterationPayload struct {
	TargetItemID   string `json:"targetItemId"`
	AlterationSlot int    `json:"alterationSlot"`
	AlterationID   string `json:"alterationId"`
}

type SetActiveHeroLoadoutPayload struct {
	SelectedLoadout string `json:"selectedLoadout"`
}

type SetAffiliateNamePayload struct {
	AffiliateName string `json:"affiliateName"`
}

type SetCosmeticLockerBannerPayload struct {
	LockerItem              string `json:"lockerItem"`
	BannerIconTemplateName  string `json:"bannerIconTemplateName"`
	BannerColorTemplateName string `json:"bannerColorTemplateName"`
}

type SetCosmeticLockerNamePayload struct {
	LockerItem string `json:"lockerItem"`
	Name       string `json:"name"`
}

type SetCosmeticLockerSlotPayload struct {
	LockerItem                string                                      `json:"lockerItem"`
	Category                  string                                      `json:"category"`
	ItemToSlot                string                                      `json:"itemToSlot"`
	SlotIndex                 uint8                                       `json:"slotIndex"`
	VariantUpdates            []SetCosmeticLockerSlotPayloadVariantUpdate `json:"variantUpdates"`
	OptLockerUseCountOverride int                                         `json:"optLockerUseCountOverride"`
}

type SetCosmeticLockerSlotPayloadVariantUpdate struct {
	Channel string   `json:"channel"`
	Active  string   `json:"active"`
	Owned   []string `json:"owned"`
}

type SetCosmeticLockerSlotsPayload struct {
	LockerItem  string                                     `json:"lockerItem"`
	LoadoutData []SetCosmeticLockerSlotsPayloadLoadoutData `json:"loadoutData"`
}

type SetCosmeticLockerSlotsPayloadLoadoutData struct {
	Category        string `json:"category"`
	ItemToSlot      string `json:"itemToSlot"`
	IndexWithinSlot uint8  `json:"indexWithinSlot"`
}

type SetForcedIntroPlayedPayload struct {
	ForcedIntroName string `json:"forcedIntroName"`
}

type SetHardcoreModifierPayload struct {
	Updates []SetHardcoreModifierPayloadUpdates `json:"updates"`
}

type SetHardcoreModifierPayloadUpdates struct {
	ModifierID string `json:"modifierId"`
	Enabled    bool   `json:"bEnabled"`
}

type SetHeroCosmeticVariantsPayload struct {
	HeroItem          string                                   `json:"heroItem"`
	OutfitVariants    []SetHeroCosmeticVariantsPayloadVariants `json:"outfitVariants"`
	BackblingVariants []SetHeroCosmeticVariantsPayloadVariants `json:"backblingVariants"`
}

type SetHeroCosmeticVariantsPayloadVariants struct {
	Channel string        `json:"channel"`
	Active  string        `json:"active"`
	Owned   []interface{} `json:"owned"` // TODO: proper type
}

type SetHomebaseBannerPayload struct {
	HomebaseBannerIconID  string `json:"homebaseBannerIconId"`
	HomebaseBannerColorID string `json:"homebaseBannerColorId"`
}

type SetHomebaseNamePayload struct {
	HomebaseName string `json:"homebaseName"`
}

type SetItemArchivedStatusBatchPayload struct {
	ItemIDs  []string `json:"itemIds"`
	Archived bool     `json:"archived"`
}

type SetItemFavoriteStatusPayload struct {
	TargetItemID string `json:"targetItemId"`
	Favorite     bool   `json:"bFavorite"`
}

type SetItemFavoriteStatusBatchPayload struct {
	ItemIDs       []string `json:"itemIds"`
	ItemFavStatus []bool   `json:"itemFavStatus"`
}

type SetMtxPlatformPayload struct {
	NewPlatform string `json:"newPlatform"`
}

type SetPinnedQuestsPayload struct {
	PinnedQuestIDs []string `json:"pinnedQuestIds"`
}

type SetRandomCosmeticLoadoutFlagPayload struct {
	Random bool `json:"random"`
}

type SetReceiveGiftsEnabledPayload struct {
	ReceiveGifts bool `json:"bReceiveGifts"`
}

type SetRewardGraphConfigPayload struct {
	State         []string `json:"state"`
	RewardGraphID string   `json:"rewardGraphId"`
}

type StartExpeditionPayload struct {
	ExpeditionID string   `json:"expeditionId"`
	SquadID      string   `json:"squadId"`
	ItemIDs      []string `json:"itemIds"`
	SlotIndices  []int    `json:"slotIndices"`
}

type StorageTransferPayload struct {
	TransferOperations []StorageTransferPayloadTransferOperations `json:"transferOperations"`
}

type StorageTransferPayloadTransferOperations struct {
	ItemID        string `json:"itemId"`
	Quantity      int    `json:"quantity"`
	ToStorage     bool   `json:"toStorage"`
	NewItemIdHint string `json:"newItemIdHint"`
}

type ToggleQuestActiveStatePayload struct {
	QuestIDs []string `json:"questIds"`
}

type UnassignAllSquadsPayload struct {
	SquadIDs []SquadAttribute `json:"squadIds"`
}

type UnlockRewardNodePayload struct {
	NodeID        string `json:"nodeId"`
	RewardGraphID string `json:"rewardGraphId"`
	RewardCFG     string `json:"rewardCfg"`
}

type UnslotItemFromCollectionBookPayload struct {
	TemplateID string `json:"templateId"`
	ItemID     string `json:"itemId"`
	Specific   string `json:"specific"`
}

type UpdateQuestClientObjectivesPayload struct {
	Advance []UpdateQuestClientObjectivesPayloadAdvance `json:"advance"`
}

type UpdateQuestClientObjectivesPayloadAdvance struct {
	StatName        string `json:"statName"`
	Count           int    `json:"count"`
	TimestampOffset int    `json:"timestampOffset"`
}

type UpgradeAlterationPayload struct {
	TargetItemID   string `json:"targetItemId"`
	AlterationSlot uint8  `json:"alterationSlot"`
}

type UpgradeItemPayload struct {
	TargetItemID string `json:"targetItemId"`
}

type UpgradeItemBulkPayload struct {
	TargetItemID                string `json:"targetItemId"`
	DesiredLevel                int    `json:"desiredLevel"`
	DesiredTier                 string `json:"desiredTier"`
	ConversionRecipeIndexChoice int    `json:"conversionRecipeIndexChoice"`
}

type UpgradeItemRarityPayload struct {
	TargetItemID string `json:"targetItemId"`
}

type UpgradeSlottedItemPayload struct {
	TargetItemID string `json:"targetItemId"`
	DesiredLevel int    `json:"desiredLevel"`
}

type VerifyRealMoneyPurchasePayload struct {
	AppStore              string `json:"appStore"`
	AppStoreID            string `json:"appStoreId"`
	ReceiptID             string `json:"receiptId"`
	ReceiptInfo           string `json:"receiptInfo"`
	PurchaseCorrelationID string `json:"purchaseCorrelationId"`
}