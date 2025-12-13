package vinderman

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/8h9x/vinderman/consts"
)

type ProfileStatsType interface {
	AthenaProfileStats | CampaignProfileStats | CollectionBookPeopleProfileStats | CollectionBookSchematicsProfileStats | CollectionsProfileStats | CommonPublicProfileStats |
	CommonCoreProfileStats | CreativeProfileStats | MetadataProfileStats | OutpostProfileStats | RecycleBinProfileStats | Theater0ProfileStats | Theater1ProfileStats | Theater2ProfileStats
}

type Profile[ST ProfileStatsType, NT CampaignNotifications | []interface{}] struct {
	ProfileRevision            int    `json:"profileRevision"`
	ProfileID                  string `json:"profileId"`
	ProfileChangesBaseRevision int    `json:"profileChangesBaseRevision"`
	ProfileChanges             []struct {
		ChangeType string `json:"changeType"`
		Profile    struct {
			Created         string                     `json:"created"`
			Updated         string                     `json:"updated"`
			RVN             int                        `json:"rvn"`
			WipeNumber      int                        `json:"wipeNumber"`
			AccountID       string                     `json:"accountId"`
			ProfileID       string                     `json:"profileId"`
			Version         string                     `json:"version"`
			Items           map[string]json.RawMessage `json:"items"`
			Stats           ST                         `json:"stats"`
			CommandRevision int                        `json:"commandRevision"`
		} `json:"profile"`
	} `json:"profileChanges"`
	ProfileCommandRevision int       `json:"profileCommandRevision"`
	ServerTime             time.Time `json:"serverTime"`
	ResponseVersion        int       `json:"responseVersion"`
	Notifications          NT        `json:"notifications"`
}

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

func (c *Client) ComposeProfileOperation(operation string, profileID string, payload string) (resp *http.Response, err error) {
	credentials := c.CredentialsMap[c.ClientID]

	headers := http.Header{}
	headers.Set("Content-Type", "application/json")
	headers.Set("Authorization", "Bearer "+credentials.AccessToken)

	resp, err = c.Request("POST", fmt.Sprintf("%s/fortnite/api/game/v2/profile/%s/client/%s?profileId=%s&rvn=-1", consts.FortniteMCPService, credentials.AccountID, operation, profileID), headers, payload)
	return
}

func (c *Client) ProfileOperation(operation string, profileID string, payload any) (resp *http.Response, err error) {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return
	}

	return c.ComposeProfileOperation(operation, profileID, string(bodyBytes))
}

func (c *Client) AbandonExpedition(expeditionId string) (*http.Response, error) {
	return c.ProfileOperation("AbandonExpedition", "campaign", AbandonExpeditionPayload{ExpeditionID: expeditionId})
}

func (c *Client) ActivateConsumable(payload ActivateConsumablePayload) (*http.Response, error) {
	return c.ProfileOperation("ActivateConsumable", "campaign", payload)
}

func (c *Client) AddToCollection(payload AddToCollectionPayload) (*http.Response, error) {
	return c.ProfileOperation("AddToCollection", "collections", payload)
}

func (c *Client) ApplyVote(offerID string) (*http.Response, error) {
	return c.ProfileOperation("ApplyVote", "athena", ApplyVotePayload{OfferID: offerID})
}

func (c *Client) AssignGadgetToLoadout(payload AssignGadgetToLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation("AssignGadgetToLoadout", "campaign", payload)
}

func (c *Client) AssignHeroToLoadout(payload AssignHeroToLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation("AssignHeroToLoadout", "campaign", payload)
}

func (c *Client) AssignTeamPerkToLoadout(payload AssignTeamPerkToLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation("AssignTeamPerkToLoadout", "campaign", payload)
}

func (c *Client) AssignWorkerToSquad(payload AssignWorkerToSquadPayload) (*http.Response, error) {
	return c.ProfileOperation("AssignWorkerToSquad", "campaign", payload)
}

func (c *Client) AssignWorkerToSquadBatch(payload AssignWorkerToSquadBatchPayload) (*http.Response, error) {
	return c.ProfileOperation("AssignWorkerToSquadBatch", "campaign", payload)
}

func (c *Client) AthenaPinQuest(pinnedQuest string) (*http.Response, error) {
	return c.ProfileOperation("AthenaPinQuest", "athena", AthenaPinQuestPayload{PinnedQuest: pinnedQuest})
}

func (c *Client) AthenaRemoveQuests(removedQuests []string) (*http.Response, error) {
	return c.ProfileOperation("AthenaRemoveQuests", "athena", AthenaRemoveQuestsPayload{RemovedQuests: removedQuests})
}

func (c *Client) BulkUpdateCollections(payload BulkUpdateCollectionsPayload) (*http.Response, error) {
	return c.ProfileOperation("BulkUpdateCollections", "collections", payload)
}

func (c *Client) CancelOrResumeSubscription(payload CancelOrResumeSubscriptionPayload) (*http.Response, error) {
	return c.ProfileOperation("CancelOrResumeSubscription", "common_core", payload)
}

func (c *Client) ChallengeBundleLevelUp(bundleIdToLevel string) (*http.Response, error) {
	return c.ProfileOperation("ChallengeBundleLevelUp", "athena", ChallengeBundleLevelUpPayload{BundleIdToLevel: bundleIdToLevel})
}

func (c *Client) ClaimCollectedResources(collectorsToClaim []string) (*http.Response, error) {
	return c.ProfileOperation("ClaimCollectedResources", "campaign", ClaimCollectedResourcesPayload{CollectorsToClaim: collectorsToClaim})
}

// ClaimCollectionBookPageRewards collection_book_people0, collection_book_schematics0
func (c *Client) ClaimCollectionBookPageRewards(profileID string, payload ClaimCollectionBookPageRewardsPayload) (*http.Response, error) {
	return c.ProfileOperation("ClaimCollectionBookPageRewards", profileID, payload)
}

func (c *Client) ClaimCollectionBookRewards(payload ClaimCollectionBookRewardsPayload) (*http.Response, error) {
	return c.ProfileOperation("ClaimCollectionBookPageRewards", "campaign", payload)
}

func (c *Client) ClaimImportFriendsReward(network ESocialImportPanelPlatform) (*http.Response, error) {
	return c.ProfileOperation("ClaimImportFriendsReward", "common_core", ClaimImportFriendsRewardPayload{Network: network})
}

func (c *Client) ClaimLoginReward() (*http.Response, error) {
	return c.ComposeProfileOperation("ClaimLoginReward", "campaign", "{}")
}

func (c *Client) ClaimMFAEnabled(claimForSTW bool) (*http.Response, error) {
	return c.ProfileOperation("ClaimMfaEnabled", "common_core", ClaimMFAEnabledPayload{ClaimForSTW: claimForSTW})
}

func (c *Client) ClaimMissionAlertRewards() (*http.Response, error) {
	return c.ComposeProfileOperation("ClaimMissionAlertRewards", "campaign", "{}")
}

// ClaimQuestReward athena, campaign
func (c *Client) ClaimQuestReward(profileID string, payload ClaimQuestRewardPayload) (*http.Response, error) {
	return c.ProfileOperation("ClaimQuestReward", profileID, payload)
}

func (c *Client) ClaimSubscriptionRewards(payload ClaimSubscriptionRewardsPayload) (*http.Response, error) {
	return c.ProfileOperation("ClaimSubscriptionRewards", "common_core", payload)
}

func (c *Client) ClearHeroLoadout(loadoutID string) (*http.Response, error) {
	return c.ProfileOperation("ClearHeroLoadout", "campaign", ClearHeroLoadoutPayload{LoadoutID: loadoutID})
}

// ClientQuestLogin athena, campaign
func (c *Client) ClientQuestLogin(profileID string, streamingAppKey string) (*http.Response, error) {
	return c.ProfileOperation("ClientQuestLogin", profileID, ClientQuestLoginPayload{StreamingAppKey: streamingAppKey})
}

func (c *Client) CollectExpedition(payload CollectExpeditionPayload) (*http.Response, error) {
	return c.ProfileOperation("CollectExpedition", "campaign", payload)
}

func (c *Client) CompletePlayerSurvey(payload CompletePlayerSurveyPayload) (*http.Response, error) {
	return c.ProfileOperation("CompletePlayerSurvey", "common_core", payload)
}

func (c *Client) ConvertItem(payload ConvertItemPayload) (*http.Response, error) {
	return c.ProfileOperation("ConvertItem", "campaign", payload)
}

// ConvertSlottedItem collection_book_people0, collection_book_schematics0
func (c *Client) ConvertSlottedItem(profileID string, payload ConvertSlottedItemPayload) (*http.Response, error) {
	return c.ProfileOperation("ConvertSlottedItem", profileID, payload)
}

// CopyCosmeticLoadout athena, campaign
func (c *Client) CopyCosmeticLoadout(profileID string, payload CopyCosmeticLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation("CopyCosmeticLoadout", profileID, payload)
}

func (c *Client) CraftWorldItem(payload CraftWorldItemPayload) (*http.Response, error) {
	return c.ProfileOperation("CraftWorldItem", "theater0", payload)
}

func (c *Client) DeleteBattleLabIsland() (*http.Response, error) {
	return c.ComposeProfileOperation("DeleteBattleLabIsland", "creative", "{}")
}

// DeleteCosmeticLoadout athena, campaign
func (c *Client) DeleteCosmeticLoadout(profileID string, payload DeleteCosmeticLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation("DeleteCosmeticLoadout", profileID, payload)
}

// DestroyWorldItems outpost0, theater0, theater1, theater2
func (c *Client) DestroyWorldItems(profileID string, itemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("DestroyWorldItems", profileID, DestroyWorldItemsPayload{ItemIDs: itemIDs})
}

// DisassembleWorldItems theater0, theater1, theater2
func (c *Client) DisassembleWorldItems(profileID string, payload DisassembleWorldItemsPayload) (*http.Response, error) {
	return c.ProfileOperation("DisassembleWorldItems", profileID, payload)
}

func (c *Client) ExchangeGameCurrencyForBattlePassOffer(offerItemIDList []string) (*http.Response, error) {
	return c.ProfileOperation("ExchangeGameCurrencyForBattlePassOffer", "athena", ExchangeGameCurrencyForBattlePassOfferPayload{OfferItemIDList: offerItemIDList})
}

func (c *Client) ExchangeGiftToken() (*http.Response, error) {
	return c.ComposeProfileOperation("ExchangeGiftToken", "athena", "{}")
}

// FortRerollDailyQuest athena, campaign
func (c *Client) FortRerollDailyQuest(profileID string, questID string) (*http.Response, error) {
	return c.ProfileOperation("FortRerollDailyQuest", profileID, FortRerollDailyQuestPayload{QuestID: questID})
}

func (c *Client) GiftCatalogEntry(payload GiftCatalogEntryPayload) (*http.Response, error) {
	return c.ProfileOperation("GiftCatalogEntry", "common_core", payload)
}

func (c *Client) IssueFriendCode(codeTokenType string) (*http.Response, error) {
	return c.ProfileOperation("IssueFriendCode", "common_core", IssueFriendCodePayload{CodeTokenType: codeTokenType})
}

func (c *Client) MarkCollectedItemsSeen(payload MarkCollectedItemsSeenPayload) (*http.Response, error) {
	return c.ProfileOperation("MarkCollectedItemsSeen", "collections", payload)
}

// MarkItemSeen athena, campaign, common_core
func (c *Client) MarkItemSeen(profileID string, itemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("MarkItemSeen", profileID, MarkItemSeenPayload{ItemIDs: itemIDs})
}

// MarkNewQuestNotificationSent athena, campaign
func (c *Client) MarkNewQuestNotificationSent(profileID string, itemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("MarkNewQuestNotificationSent", profileID, MarkNewQuestNotificationSentPayload{ItemIDs: itemIDs})
}

// ModifyQuickbar theater0, theater1, theater2
func (c *Client) ModifyQuickbar(profileID string, payload ModifyQuickbarPayload) (*http.Response, error) {
	return c.ProfileOperation("ModifyQuickbar", profileID, payload)
}

func (c *Client) OpenCardPack(payload OpenCardPackPayload) (*http.Response, error) {
	return c.ProfileOperation("OpenCardPack", "campaign", payload)
}

func (c *Client) OpenCardPackBatch(cardPackitemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("OpenCardPackBatch", "campaign", OpenCardPackBatchPayload{CardPackItemIDs: cardPackitemIDs})
}

func (c *Client) PopulatePrerolledOffers() (*http.Response, error) {
	return c.ComposeProfileOperation("PopulatePrerolledOffers", "campaign", "{}")
}

// PromoteItem campaign, collection_book_people0, collection_book_schematics0
func (c *Client) PromoteItem(profileID string, targetItemID string) (*http.Response, error) {
	return c.ProfileOperation("PromoteItem", profileID, PromoteItemPayload{TargetItemID: targetItemID})
}

func (c *Client) PurchaseCatalogEntry(payload PurchaseCatalogEntryPayload) (*http.Response, error) {
	return c.ProfileOperation("PurchaseCatalogEntry", "common_core", payload)
}

func (c *Client) PurchaseMultipleCatalogEntries(payload PurchaseMultipleCatalogEntriesPayload) (*http.Response, error) {
	return c.ProfileOperation("PurchaseMultipleCatalogEntries", "common_core", payload)
}

func (c *Client) PurchaseOrUpgradeHomebaseNode(nodeID string) (*http.Response, error) {
	return c.ProfileOperation("PurchaseOrUpgradeHomebaseNode", "campaign", PurchaseOrUpgradeHomebaseNodePayload{NodeID: nodeID})
}

func (c *Client) PurchaseResearchStatUpgrade(statID string) (*http.Response, error) {
	return c.ProfileOperation("PurchaseResearchStatUpgrade", "campaign", PurchaseResearchStatUpgradePayload{StatID: statID})
}

func (c *Client) QueryProfile(profileID string) (*http.Response, error) {
	return c.ComposeProfileOperation("QueryProfile", profileID, "{}")
}

// TODO: QueryPublicProfile

func (c *Client) RecycleItem(targetItemID string) (*http.Response, error) {
	return c.ProfileOperation("RecycleItem", "campaign", RecycleItemPayload{TargetItemID: targetItemID})
}

func (c *Client) RecycleItemBatch(targetItemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("RecycleItemBatch", "campaign", RecycleItemBatchPayload{TargetItemIDs: targetItemIDs})
}

func (c *Client) RedeemRealMoneyPurchases(payload RedeemRealMoneyPurchasesPayload) (*http.Response, error) {
	return c.ProfileOperation("RedeemRealMoneyPurchases", "common_core", payload)
}

func (c *Client) RedeemSTWAccoladeTokens() (*http.Response, error) {
	return c.ComposeProfileOperation("RedeemSTWAccoladeTokens", "athena", "{}")
}

func (c *Client) RefreshExpeditions() (*http.Response, error) {
	return c.ComposeProfileOperation("RefreshExpeditions", "campaign", "{}")
}

func (c *Client) RefundItem(targetItemID string) (*http.Response, error) {
	return c.ProfileOperation("RefundItem", "campaign", RefundItemPayload{TargetItemID: targetItemID})
}

func (c *Client) RefundMtxPurchase(payload RefundMtxPurchasePayload) (*http.Response, error) {
	return c.ProfileOperation("RefundMtxPurchase", "common_core", payload)
}

func (c *Client) RemoveGiftBox(profileID string, giftBoxItemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("RemoveGiftBox", profileID, RemoveGiftBoxPayload{GiftBoxItemIDs: giftBoxItemIDs})
}

func (c *Client) RequestRestedStateIncrease(payload RequestRestedStateIncreasePayload) (*http.Response, error) {
	return c.ProfileOperation("RequestRestedStateIncrease", "athena", payload)
}

// ResearchItemFromCollectionBook campaign, theater0, theater1, theater2
func (c *Client) ResearchItemFromCollectionBook(profileID string, templateID string) (*http.Response, error) {
	return c.ProfileOperation("ResearchItemFromCollectionBook", profileID, ResearchItemFromCollectionBookPayload{TemplateID: templateID})
}

func (c *Client) RespecAlteration(payload RespecAlterationPayload) (*http.Response, error) {
	return c.ProfileOperation("RespecAlteration", "campaign", payload)
}

func (c *Client) RespecResearch() (*http.Response, error) {
	return c.ComposeProfileOperation("RespecResearch", "campaign", "{}")
}

func (c *Client) RespecUpgrades() (*http.Response, error) {
	return c.ComposeProfileOperation("RespecUpgrades", "campaign", "{}")
}

func (c *Client) SetActiveHeroLoadout(selectedLoadout string) (*http.Response, error) {
	return c.ProfileOperation("SetActiveHeroLoadout", "campaign", SetActiveHeroLoadoutPayload{SelectedLoadout: selectedLoadout})
}

func (c *Client) SetAffiliateNameLoadout(affiliateName string) (*http.Response, error) {
	return c.ProfileOperation("SetActiveHeroLoadout", "common_core", SetAffiliateNamePayload{AffiliateName: affiliateName})
}

// Deprecated: SetCosmeticLockerBanner Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c *Client) SetCosmeticLockerBanner(profileID string, payload SetCosmeticLockerBannerPayload) (*http.Response, error) {
	return c.ProfileOperation("SetCosmeticLockerBanner", profileID, payload)
}

// Deprecated: SetCosmeticLockerName Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c *Client) SetCosmeticLockerName(profileID string, payload SetCosmeticLockerNamePayload) (*http.Response, error) {
	return c.ProfileOperation("SetCosmeticLockerName", profileID, payload)
}

// Deprecated: SetCosmeticLockerSlot Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c *Client) SetCosmeticLockerSlot(profileID string, payload SetCosmeticLockerSlotPayload) (*http.Response, error) {
	return c.ProfileOperation("SetCosmeticLockerSlot", profileID, payload)
}

// Deprecated: SetCosmeticLockerSlots Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c *Client) SetCosmeticLockerSlots(profileID string, payload SetCosmeticLockerSlotsPayload) (*http.Response, error) {
	return c.ProfileOperation("SetCosmeticLockerSlots", profileID, payload)
}

func (c *Client) SetForcedIntroPlayed(forcedIntroName string) (*http.Response, error) {
	return c.ProfileOperation("SetForcedIntroPlayed", "common_core", SetForcedIntroPlayedPayload{ForcedIntroName: forcedIntroName})
}

func (c *Client) SetHardcoreModifier(payload SetHardcoreModifierPayload) (*http.Response, error) {
	return c.ProfileOperation("SetHardcoreModifier", "athena", payload)
}

func (c *Client) SetHeroCosmeticVariants(payload SetHeroCosmeticVariantsPayload) (*http.Response, error) {
	return c.ProfileOperation("SetHeroCosmeticVariants", "campaign", payload)
}

func (c *Client) SetHomebaseBanner(payload SetHomebaseBannerPayload) (*http.Response, error) {
	return c.ProfileOperation("SetHomebaseBanner", "common_public", payload)
}

// Deprecated: SetHomebaseName The home base naming system has been removed from the game.
func (c *Client) SetHomebaseName(homebaseName string) (*http.Response, error) {
	return c.ProfileOperation("SetHomebaseName", "common_core", SetHomebaseNamePayload{HomebaseName: homebaseName})
}

func (c *Client) SetIntroGamePlayed() (*http.Response, error) {
	return c.ComposeProfileOperation("SetIntroGamePlayed", "common_core", "{}")
}

func (c *Client) SetItemArchivedStatusBatch(payload SetItemArchivedStatusBatchPayload) (*http.Response, error) {
	return c.ProfileOperation("SetItemArchivedStatusBatch", "athena", payload)
}

// SetItemFavoriteStatus athena, campaign
func (c *Client) SetItemFavoriteStatus(profileID string, payload SetItemFavoriteStatusPayload) (*http.Response, error) {
	return c.ProfileOperation("SetItemFavoriteStatus", profileID, payload)
}

// SetItemFavoriteStatusBatch athena, campaign
func (c *Client) SetItemFavoriteStatusBatch(profileID string, payload SetItemFavoriteStatusBatchPayload) (*http.Response, error) {
	return c.ProfileOperation("SetItemFavoriteStatusBatch", profileID, payload)
}

func (c *Client) SetMatchmakingBansViewed() (*http.Response, error) {
	return c.ComposeProfileOperation("SetMatchmakingBansViewed", "common_core", "{}")
}

func (c *Client) SetMtxPlatform(newPlatform string) (*http.Response, error) {
	return c.ProfileOperation("SetMtxPlatform", "common_core", SetMtxPlatformPayload{NewPlatform: newPlatform})
}

func (c *Client) SetPinnedQuests(pinnedQuestIDs []string) (*http.Response, error) {
	return c.ProfileOperation("SetPinnedQuests", "common_core", SetPinnedQuestsPayload{PinnedQuestIDs: pinnedQuestIDs})
}

// SetRandomCosmeticLoadoutFlag athena, campaign
func (c *Client) SetRandomCosmeticLoadoutFlag(profileID string, random bool) (*http.Response, error) {
	return c.ProfileOperation("SetRandomCosmeticLoadoutFlag", profileID, SetRandomCosmeticLoadoutFlagPayload{Random: random})
}

func (c *Client) SetReceiveGiftsEnabled(receiveGifts bool) (*http.Response, error) {
	return c.ProfileOperation("SetReceiveGiftsEnabled", "common_core", SetReceiveGiftsEnabledPayload{ReceiveGifts: receiveGifts})
}

func (c *Client) SetRewardGraphConfig(payload SetRewardGraphConfigPayload) (*http.Response, error) {
	return c.ProfileOperation("SetRewardGraphConfig", "athena", payload)
}

func (c *Client) StartExpedition(payload StartExpeditionPayload) (*http.Response, error) {
	return c.ProfileOperation("StartExpedition", "campaign", payload)
}

func (c *Client) StorageTransfer(payload StorageTransferPayload) (*http.Response, error) {
	return c.ProfileOperation("StorageTransfer", "theater0", payload)
}

func (c *Client) ToggleQuestActiveState(questIDs []string) (*http.Response, error) {
	return c.ProfileOperation("ToggleQuestActiveState", "athena", ToggleQuestActiveStatePayload{QuestIDs: questIDs})
}

func (c *Client) UnassignAllSquads(squadIDs []SquadAttribute) (*http.Response, error) {
	return c.ProfileOperation("UnassignAllSquads", "campaign", UnassignAllSquadsPayload{SquadIDs: squadIDs})
}

func (c *Client) UnlockRewardNode(payload UnlockRewardNodePayload) (*http.Response, error) {
	return c.ProfileOperation("UnlockRewardNode", "athena", payload)
}

// UnslotItemFromCollectionBook campaign, theater0, theater1, theater2
func (c *Client) UnslotItemFromCollectionBook(profileID string, payload UnslotItemFromCollectionBookPayload) (*http.Response, error) {
	return c.ProfileOperation("UnslotItemFromCollectionBook", profileID, payload)
}

// UpdateQuestClientObjectives athena, campaign
func (c *Client) UpdateQuestClientObjectives(profileID string, payload UpdateQuestClientObjectivesPayload) (*http.Response, error) {
	return c.ProfileOperation("UpdateQuestClientObjectives", profileID, payload)
}

func (c *Client) UpgradeAlteration(payload UpgradeAlterationPayload) (*http.Response, error) {
	return c.ProfileOperation("UpgradeAlteration", "campaign", payload)
}

func (c *Client) UpgradeItem(targetItemID string) (*http.Response, error) {
	return c.ProfileOperation("UpgradeItem", "campaign", UpgradeItemPayload{TargetItemID: targetItemID})
}

func (c *Client) UpgradeItemBulk(payload UpgradeItemBulkPayload) (*http.Response, error) {
	return c.ProfileOperation("UpgradeItemBulk", "campaign", payload)
}

func (c *Client) UpgradeItemRarity(targetItemID string) (*http.Response, error) {
	return c.ProfileOperation("UpgradeItemRarity", "campaign", UpgradeItemRarityPayload{TargetItemID: targetItemID})
}

// UpgradeSlottedItem collection_book_people0, collection_book_schematics0
func (c *Client) UpgradeSlottedItem(profileID string, payload UpgradeSlottedItemPayload) (*http.Response, error) {
	return c.ProfileOperation("UpgradeSlottedItem", profileID, payload)
}

func (c *Client) VerifyRealMoneyPurchase(payload VerifyRealMoneyPurchasePayload) (*http.Response, error) {
	return c.ProfileOperation("VerifyRealMoneyPurchase", "common_core", payload)
}

//func ProfileOperationExample() () {
//	client := New()
//
//	res, err := client.QueryProfile(UserCredentials{}, "athena")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	data, err := request.ResponseParser[Profile[AthenaProfileStats]](res)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var skins []AthenaCosmeticItem
//
//	for _, item := range data.Body.ProfileChanges[0].Profile.Items {
//		var cosmetic AthenaCosmeticItem
//		if err = json.Unmarshal(item, &cosmetic); err != nil {
//			// not a skin; (you should probably add an additional check to ensure that it isn't some other type of error occurring); TODO: abstract this to a helper function that properly error checks and returns an empty state of the type passed if the type of data doesnt match
//			continue
//		}
//
//		if strings.HasPrefix(cosmetic.TemplateID, "AthenaCharacter") {
//			skins = append(skins, cosmetic)
//		}
//	}
//
//	log.Println("Account Level:", data.Body.ProfileChanges[0].Profile.Stats.Attributes.AccountLevel)
//	log.Println("Skin Count:", len(skins))
//}
