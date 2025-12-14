package fortgo

import (
	"net/http"

	"github.com/8h9x/fortgo/fortnite"
)

func (c *Client) AbandonExpedition(expeditionId string) (*http.Response, error) {
	return c.ProfileOperation("AbandonExpedition", "campaign", fortnite.AbandonExpeditionPayload{ExpeditionID: expeditionId})
}

func (c *Client) ActivateConsumable(payload fortnite.ActivateConsumablePayload) (*http.Response, error) {
	return c.ProfileOperation("ActivateConsumable", "campaign", payload)
}

func (c *Client) AddToCollection(payload fortnite.AddToCollectionPayload) (*http.Response, error) {
	return c.ProfileOperation("AddToCollection", "collections", payload)
}

func (c *Client) ApplyVote(offerID string) (*http.Response, error) {
	return c.ProfileOperation("ApplyVote", "athena", fortnite.ApplyVotePayload{OfferID: offerID})
}

func (c *Client) AssignGadgetToLoadout(payload fortnite.AssignGadgetToLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation("AssignGadgetToLoadout", "campaign", payload)
}

func (c *Client) AssignHeroToLoadout(payload fortnite.AssignHeroToLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation("AssignHeroToLoadout", "campaign", payload)
}

func (c *Client) AssignTeamPerkToLoadout(payload fortnite.AssignTeamPerkToLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation("AssignTeamPerkToLoadout", "campaign", payload)
}

func (c *Client) AssignWorkerToSquad(payload fortnite.AssignWorkerToSquadPayload) (*http.Response, error) {
	return c.ProfileOperation("AssignWorkerToSquad", "campaign", payload)
}

func (c *Client) AssignWorkerToSquadBatch(payload fortnite.AssignWorkerToSquadBatchPayload) (*http.Response, error) {
	return c.ProfileOperation("AssignWorkerToSquadBatch", "campaign", payload)
}

func (c *Client) AthenaPinQuest(pinnedQuest string) (*http.Response, error) {
	return c.ProfileOperation("AthenaPinQuest", "athena", fortnite.AthenaPinQuestPayload{PinnedQuest: pinnedQuest})
}

func (c *Client) AthenaRemoveQuests(removedQuests []string) (*http.Response, error) {
	return c.ProfileOperation("AthenaRemoveQuests", "athena", fortnite.AthenaRemoveQuestsPayload{RemovedQuests: removedQuests})
}

func (c *Client) BulkUpdateCollections(payload fortnite.BulkUpdateCollectionsPayload) (*http.Response, error) {
	return c.ProfileOperation("BulkUpdateCollections", "collections", payload)
}

func (c *Client) CancelOrResumeSubscription(payload fortnite.CancelOrResumeSubscriptionPayload) (*http.Response, error) {
	return c.ProfileOperation("CancelOrResumeSubscription", "common_core", payload)
}

func (c *Client) ChallengeBundleLevelUp(bundleIdToLevel string) (*http.Response, error) {
	return c.ProfileOperation("ChallengeBundleLevelUp", "athena", fortnite.ChallengeBundleLevelUpPayload{BundleIdToLevel: bundleIdToLevel})
}

func (c *Client) ClaimCollectedResources(collectorsToClaim []string) (*http.Response, error) {
	return c.ProfileOperation("ClaimCollectedResources", "campaign", fortnite.ClaimCollectedResourcesPayload{CollectorsToClaim: collectorsToClaim})
}

// ClaimCollectionBookPageRewards collection_book_people0, collection_book_schematics0
func (c *Client) ClaimCollectionBookPageRewards(profileID string, payload fortnite.ClaimCollectionBookPageRewardsPayload) (*http.Response, error) {
	return c.ProfileOperation("ClaimCollectionBookPageRewards", profileID, payload)
}

func (c *Client) ClaimCollectionBookRewards(payload fortnite.ClaimCollectionBookRewardsPayload) (*http.Response, error) {
	return c.ProfileOperation("ClaimCollectionBookPageRewards", "campaign", payload)
}

func (c *Client) ClaimImportFriendsReward(network fortnite.ESocialImportPanelPlatform) (*http.Response, error) {
	return c.ProfileOperation("ClaimImportFriendsReward", "common_core", fortnite.ClaimImportFriendsRewardPayload{Network: network})
}

func (c *Client) ClaimLoginReward() (*http.Response, error) {
	return c.ComposeProfileOperation("ClaimLoginReward", "campaign", "{}")
}

func (c *Client) ClaimMFAEnabled(claimForSTW bool) (*http.Response, error) {
	return c.ProfileOperation("ClaimMfaEnabled", "common_core", fortnite.ClaimMFAEnabledPayload{ClaimForSTW: claimForSTW})
}

func (c *Client) ClaimMissionAlertRewards() (*http.Response, error) {
	return c.ComposeProfileOperation("ClaimMissionAlertRewards", "campaign", "{}")
}

// ClaimQuestReward athena, campaign
func (c *Client) ClaimQuestReward(profileID string, payload fortnite.ClaimQuestRewardPayload) (*http.Response, error) {
	return c.ProfileOperation("ClaimQuestReward", profileID, payload)
}

func (c *Client) ClaimSubscriptionRewards(payload fortnite.ClaimSubscriptionRewardsPayload) (*http.Response, error) {
	return c.ProfileOperation("ClaimSubscriptionRewards", "common_core", payload)
}

func (c *Client) ClearHeroLoadout(loadoutID string) (*http.Response, error) {
	return c.ProfileOperation("ClearHeroLoadout", "campaign", fortnite.ClearHeroLoadoutPayload{LoadoutID: loadoutID})
}

// ClientQuestLogin athena, campaign
func (c *Client) ClientQuestLogin(profileID string, streamingAppKey string) (*http.Response, error) {
	return c.ProfileOperation("ClientQuestLogin", profileID, fortnite. ClientQuestLoginPayload{StreamingAppKey: streamingAppKey})
}

func (c *Client) CollectExpedition(payload fortnite.CollectExpeditionPayload) (*http.Response, error) {
	return c.ProfileOperation("CollectExpedition", "campaign", payload)
}

func (c *Client) CompletePlayerSurvey(payload fortnite.CompletePlayerSurveyPayload) (*http.Response, error) {
	return c.ProfileOperation("CompletePlayerSurvey", "common_core", payload)
}

func (c *Client) ConvertItem(payload fortnite.ConvertItemPayload) (*http.Response, error) {
	return c.ProfileOperation("ConvertItem", "campaign", payload)
}

// ConvertSlottedItem collection_book_people0, collection_book_schematics0
func (c *Client) ConvertSlottedItem(profileID string, payload fortnite.ConvertSlottedItemPayload) (*http.Response, error) {
	return c.ProfileOperation("ConvertSlottedItem", profileID, payload)
}

// CopyCosmeticLoadout athena, campaign
func (c *Client) CopyCosmeticLoadout(profileID string, payload fortnite.CopyCosmeticLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation("CopyCosmeticLoadout", profileID, payload)
}

func (c *Client) CraftWorldItem(payload fortnite.CraftWorldItemPayload) (*http.Response, error) {
	return c.ProfileOperation("CraftWorldItem", "theater0", payload)
}

func (c *Client) DeleteBattleLabIsland() (*http.Response, error) {
	return c.ComposeProfileOperation("DeleteBattleLabIsland", "creative", "{}")
}

// DeleteCosmeticLoadout athena, campaign
func (c *Client) DeleteCosmeticLoadout(profileID string, payload fortnite.DeleteCosmeticLoadoutPayload) (*http.Response, error) {
	return c.ProfileOperation("DeleteCosmeticLoadout", profileID, payload)
}

// DestroyWorldItems outpost0, theater0, theater1, theater2
func (c *Client) DestroyWorldItems(profileID string, itemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("DestroyWorldItems", profileID, fortnite. DestroyWorldItemsPayload{ItemIDs: itemIDs})
}

// DisassembleWorldItems theater0, theater1, theater2
func (c *Client) DisassembleWorldItems(profileID string, payload fortnite.DisassembleWorldItemsPayload) (*http.Response, error) {
	return c.ProfileOperation("DisassembleWorldItems", profileID, payload)
}

func (c *Client) ExchangeGameCurrencyForBattlePassOffer(offerItemIDList []string) (*http.Response, error) {
	return c.ProfileOperation("ExchangeGameCurrencyForBattlePassOffer", "athena", fortnite.ExchangeGameCurrencyForBattlePassOfferPayload{OfferItemIDList: offerItemIDList})
}

func (c *Client) ExchangeGiftToken() (*http.Response, error) {
	return c.ComposeProfileOperation("ExchangeGiftToken", "athena", "{}")
}

// FortRerollDailyQuest athena, campaign
func (c *Client) FortRerollDailyQuest(profileID string, questID string) (*http.Response, error) {
	return c.ProfileOperation("FortRerollDailyQuest", profileID, fortnite. FortRerollDailyQuestPayload{QuestID: questID})
}

func (c *Client) GiftCatalogEntry(payload fortnite.GiftCatalogEntryPayload) (*http.Response, error) {
	return c.ProfileOperation("GiftCatalogEntry", "common_core", payload)
}

func (c *Client) IssueFriendCode(codeTokenType string) (*http.Response, error) {
	return c.ProfileOperation("IssueFriendCode", "common_core", fortnite.IssueFriendCodePayload{CodeTokenType: codeTokenType})
}

func (c *Client) MarkCollectedItemsSeen(payload fortnite.MarkCollectedItemsSeenPayload) (*http.Response, error) {
	return c.ProfileOperation("MarkCollectedItemsSeen", "collections", payload)
}

// MarkItemSeen athena, campaign, common_core
func (c *Client) MarkItemSeen(profileID string, itemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("MarkItemSeen", profileID, fortnite. MarkItemSeenPayload{ItemIDs: itemIDs})
}

// MarkNewQuestNotificationSent athena, campaign
func (c *Client) MarkNewQuestNotificationSent(profileID string, itemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("MarkNewQuestNotificationSent", profileID, fortnite. MarkNewQuestNotificationSentPayload{ItemIDs: itemIDs})
}

// ModifyQuickbar theater0, theater1, theater2
func (c *Client) ModifyQuickbar(profileID string, payload fortnite.ModifyQuickbarPayload) (*http.Response, error) {
	return c.ProfileOperation("ModifyQuickbar", profileID, payload)
}

func (c *Client) OpenCardPack(payload fortnite.OpenCardPackPayload) (*http.Response, error) {
	return c.ProfileOperation("OpenCardPack", "campaign", payload)
}

func (c *Client) OpenCardPackBatch(cardPackitemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("OpenCardPackBatch", "campaign", fortnite.OpenCardPackBatchPayload{CardPackItemIDs: cardPackitemIDs})
}

func (c *Client) PopulatePrerolledOffers() (*http.Response, error) {
	return c.ComposeProfileOperation("PopulatePrerolledOffers", "campaign", "{}")
}

// PromoteItem campaign, collection_book_people0, collection_book_schematics0
func (c *Client) PromoteItem(profileID string, targetItemID string) (*http.Response, error) {
	return c.ProfileOperation("PromoteItem", profileID, fortnite. PromoteItemPayload{TargetItemID: targetItemID})
}

func (c *Client) PurchaseCatalogEntry(payload fortnite.PurchaseCatalogEntryPayload) (*http.Response, error) {
	return c.ProfileOperation("PurchaseCatalogEntry", "common_core", payload)
}

func (c *Client) PurchaseMultipleCatalogEntries(payload fortnite.PurchaseMultipleCatalogEntriesPayload) (*http.Response, error) {
	return c.ProfileOperation("PurchaseMultipleCatalogEntries", "common_core", payload)
}

func (c *Client) PurchaseOrUpgradeHomebaseNode(nodeID string) (*http.Response, error) {
	return c.ProfileOperation("PurchaseOrUpgradeHomebaseNode", "campaign", fortnite.PurchaseOrUpgradeHomebaseNodePayload{NodeID: nodeID})
}

func (c *Client) PurchaseResearchStatUpgrade(statID string) (*http.Response, error) {
	return c.ProfileOperation("PurchaseResearchStatUpgrade", "campaign", fortnite.PurchaseResearchStatUpgradePayload{StatID: statID})
}

func (c *Client) QueryProfile(profileID string) (*http.Response, error) {
	return c.ComposeProfileOperation("QueryProfile", profileID, "{}")
}

// TODO: QueryPublicProfile

func (c *Client) RecycleItem(targetItemID string) (*http.Response, error) {
	return c.ProfileOperation("RecycleItem", "campaign", fortnite.RecycleItemPayload{TargetItemID: targetItemID})
}

func (c *Client) RecycleItemBatch(targetItemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("RecycleItemBatch", "campaign", fortnite.RecycleItemBatchPayload{TargetItemIDs: targetItemIDs})
}

func (c *Client) RedeemRealMoneyPurchases(payload fortnite.RedeemRealMoneyPurchasesPayload) (*http.Response, error) {
	return c.ProfileOperation("RedeemRealMoneyPurchases", "common_core", payload)
}

func (c *Client) RedeemSTWAccoladeTokens() (*http.Response, error) {
	return c.ComposeProfileOperation("RedeemSTWAccoladeTokens", "athena", "{}")
}

func (c *Client) RefreshExpeditions() (*http.Response, error) {
	return c.ComposeProfileOperation("RefreshExpeditions", "campaign", "{}")
}

func (c *Client) RefundItem(targetItemID string) (*http.Response, error) {
	return c.ProfileOperation("RefundItem", "campaign", fortnite.RefundItemPayload{TargetItemID: targetItemID})
}

func (c *Client) RefundMtxPurchase(payload fortnite.RefundMtxPurchasePayload) (*http.Response, error) {
	return c.ProfileOperation("RefundMtxPurchase", "common_core", payload)
}

func (c *Client) RemoveGiftBox(profileID string, giftBoxItemIDs []string) (*http.Response, error) {
	return c.ProfileOperation("RemoveGiftBox", profileID, fortnite. RemoveGiftBoxPayload{GiftBoxItemIDs: giftBoxItemIDs})
}

func (c *Client) RequestRestedStateIncrease(payload fortnite.RequestRestedStateIncreasePayload) (*http.Response, error) {
	return c.ProfileOperation("RequestRestedStateIncrease", "athena", payload)
}

// ResearchItemFromCollectionBook campaign, theater0, theater1, theater2
func (c *Client) ResearchItemFromCollectionBook(profileID string, templateID string) (*http.Response, error) {
	return c.ProfileOperation("ResearchItemFromCollectionBook", profileID, fortnite. ResearchItemFromCollectionBookPayload{TemplateID: templateID})
}

func (c *Client) RespecAlteration(payload fortnite.RespecAlterationPayload) (*http.Response, error) {
	return c.ProfileOperation("RespecAlteration", "campaign", payload)
}

func (c *Client) RespecResearch() (*http.Response, error) {
	return c.ComposeProfileOperation("RespecResearch", "campaign", "{}")
}

func (c *Client) RespecUpgrades() (*http.Response, error) {
	return c.ComposeProfileOperation("RespecUpgrades", "campaign", "{}")
}

func (c *Client) SetActiveHeroLoadout(selectedLoadout string) (*http.Response, error) {
	return c.ProfileOperation("SetActiveHeroLoadout", "campaign", fortnite.SetActiveHeroLoadoutPayload{SelectedLoadout: selectedLoadout})
}

func (c *Client) SetAffiliateNameLoadout(affiliateName string) (*http.Response, error) {
	return c.ProfileOperation("SetActiveHeroLoadout", "common_core", fortnite.SetAffiliateNamePayload{AffiliateName: affiliateName})
}

// Deprecated: SetCosmeticLockerBanner Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c *Client) SetCosmeticLockerBanner(profileID string, payload fortnite.SetCosmeticLockerBannerPayload) (*http.Response, error) {
	return c.ProfileOperation("SetCosmeticLockerBanner", profileID, payload)
}

// Deprecated: SetCosmeticLockerName Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c *Client) SetCosmeticLockerName(profileID string, payload fortnite.SetCosmeticLockerNamePayload) (*http.Response, error) {
	return c.ProfileOperation("SetCosmeticLockerName", profileID, payload)
}

// Deprecated: SetCosmeticLockerSlot Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c *Client) SetCosmeticLockerSlot(profileID string, payload fortnite.SetCosmeticLockerSlotPayload) (*http.Response, error) {
	return c.ProfileOperation("SetCosmeticLockerSlot", profileID, payload)
}

// Deprecated: SetCosmeticLockerSlots Modern versions of the fortnite mcp system use a separate locker api--use locker.UpdateActiveLockerLoadout()
func (c *Client) SetCosmeticLockerSlots(profileID string, payload fortnite.SetCosmeticLockerSlotsPayload) (*http.Response, error) {
	return c.ProfileOperation("SetCosmeticLockerSlots", profileID, payload)
}

func (c *Client) SetForcedIntroPlayed(forcedIntroName string) (*http.Response, error) {
	return c.ProfileOperation("SetForcedIntroPlayed", "common_core", fortnite.SetForcedIntroPlayedPayload{ForcedIntroName: forcedIntroName})
}

func (c *Client) SetHardcoreModifier(payload fortnite.SetHardcoreModifierPayload) (*http.Response, error) {
	return c.ProfileOperation("SetHardcoreModifier", "athena", payload)
}

func (c *Client) SetHeroCosmeticVariants(payload fortnite.SetHeroCosmeticVariantsPayload) (*http.Response, error) {
	return c.ProfileOperation("SetHeroCosmeticVariants", "campaign", payload)
}

func (c *Client) SetHomebaseBanner(payload fortnite.SetHomebaseBannerPayload) (*http.Response, error) {
	return c.ProfileOperation("SetHomebaseBanner", "common_public", payload)
}

// Deprecated: SetHomebaseName The home base naming system has been removed from the game.
func (c *Client) SetHomebaseName(homebaseName string) (*http.Response, error) {
	return c.ProfileOperation("SetHomebaseName", "common_core", fortnite.SetHomebaseNamePayload{HomebaseName: homebaseName})
}

func (c *Client) SetIntroGamePlayed() (*http.Response, error) {
	return c.ComposeProfileOperation("SetIntroGamePlayed", "common_core", "{}")
}

func (c *Client) SetItemArchivedStatusBatch(payload fortnite.SetItemArchivedStatusBatchPayload) (*http.Response, error) {
	return c.ProfileOperation("SetItemArchivedStatusBatch", "athena", payload)
}

// SetItemFavoriteStatus athena, campaign
func (c *Client) SetItemFavoriteStatus(profileID string, payload fortnite.SetItemFavoriteStatusPayload) (*http.Response, error) {
	return c.ProfileOperation("SetItemFavoriteStatus", profileID, payload)
}

// SetItemFavoriteStatusBatch athena, campaign
func (c *Client) SetItemFavoriteStatusBatch(profileID string, payload fortnite.SetItemFavoriteStatusBatchPayload) (*http.Response, error) {
	return c.ProfileOperation("SetItemFavoriteStatusBatch", profileID, payload)
}

func (c *Client) SetMatchmakingBansViewed() (*http.Response, error) {
	return c.ComposeProfileOperation("SetMatchmakingBansViewed", "common_core", "{}")
}

func (c *Client) SetMtxPlatform(newPlatform string) (*http.Response, error) {
	return c.ProfileOperation("SetMtxPlatform", "common_core", fortnite.SetMtxPlatformPayload{NewPlatform: newPlatform})
}

func (c *Client) SetPinnedQuests(pinnedQuestIDs []string) (*http.Response, error) {
	return c.ProfileOperation("SetPinnedQuests", "common_core", fortnite.SetPinnedQuestsPayload{PinnedQuestIDs: pinnedQuestIDs})
}

// SetRandomCosmeticLoadoutFlag athena, campaign
func (c *Client) SetRandomCosmeticLoadoutFlag(profileID string, random bool) (*http.Response, error) {
	return c.ProfileOperation("SetRandomCosmeticLoadoutFlag", profileID, fortnite. SetRandomCosmeticLoadoutFlagPayload{Random: random})
}

func (c *Client) SetReceiveGiftsEnabled(receiveGifts bool) (*http.Response, error) {
	return c.ProfileOperation("SetReceiveGiftsEnabled", "common_core", fortnite.SetReceiveGiftsEnabledPayload{ReceiveGifts: receiveGifts})
}

func (c *Client) SetRewardGraphConfig(payload fortnite.SetRewardGraphConfigPayload) (*http.Response, error) {
	return c.ProfileOperation("SetRewardGraphConfig", "athena", payload)
}

func (c *Client) StartExpedition(payload fortnite.StartExpeditionPayload) (*http.Response, error) {
	return c.ProfileOperation("StartExpedition", "campaign", payload)
}

func (c *Client) StorageTransfer(payload fortnite.StorageTransferPayload) (*http.Response, error) {
	return c.ProfileOperation("StorageTransfer", "theater0", payload)
}

func (c *Client) ToggleQuestActiveState(questIDs []string) (*http.Response, error) {
	return c.ProfileOperation("ToggleQuestActiveState", "athena", fortnite.ToggleQuestActiveStatePayload{QuestIDs: questIDs})
}

func (c *Client) UnassignAllSquads(squadIDs []fortnite.SquadAttribute) (*http.Response, error) {
	return c.ProfileOperation("UnassignAllSquads", "campaign", fortnite.UnassignAllSquadsPayload{SquadIDs: squadIDs})
}

func (c *Client) UnlockRewardNode(payload fortnite.UnlockRewardNodePayload) (*http.Response, error) {
	return c.ProfileOperation("UnlockRewardNode", "athena", payload)
}

// UnslotItemFromCollectionBook campaign, theater0, theater1, theater2
func (c *Client) UnslotItemFromCollectionBook(profileID string, payload fortnite.UnslotItemFromCollectionBookPayload) (*http.Response, error) {
	return c.ProfileOperation("UnslotItemFromCollectionBook", profileID, payload)
}

// UpdateQuestClientObjectives athena, campaign
func (c *Client) UpdateQuestClientObjectives(profileID string, payload fortnite.UpdateQuestClientObjectivesPayload) (*http.Response, error) {
	return c.ProfileOperation("UpdateQuestClientObjectives", profileID, payload)
}

func (c *Client) UpgradeAlteration(payload fortnite.UpgradeAlterationPayload) (*http.Response, error) {
	return c.ProfileOperation("UpgradeAlteration", "campaign", payload)
}

func (c *Client) UpgradeItem(targetItemID string) (*http.Response, error) {
	return c.ProfileOperation("UpgradeItem", "campaign", fortnite.UpgradeItemPayload{TargetItemID: targetItemID})
}

func (c *Client) UpgradeItemBulk(payload fortnite.UpgradeItemBulkPayload) (*http.Response, error) {
	return c.ProfileOperation("UpgradeItemBulk", "campaign", payload)
}

func (c *Client) UpgradeItemRarity(targetItemID string) (*http.Response, error) {
	return c.ProfileOperation("UpgradeItemRarity", "campaign", fortnite.UpgradeItemRarityPayload{TargetItemID: targetItemID})
}

// UpgradeSlottedItem collection_book_people0, collection_book_schematics0
func (c *Client) UpgradeSlottedItem(profileID string, payload fortnite.UpgradeSlottedItemPayload) (*http.Response, error) {
	return c.ProfileOperation("UpgradeSlottedItem", profileID, payload)
}

func (c *Client) VerifyRealMoneyPurchase(payload fortnite.VerifyRealMoneyPurchasePayload) (*http.Response, error) {
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
