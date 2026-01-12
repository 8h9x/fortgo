package mcp

import (
	"encoding/json"
	"time"

	"github.com/8h9x/fortgo/fortnite"
)

type ProfileChange[ST ProfileStatsType] struct {
	ChangeType           string                        `json:"changeType"`
	EnableConstructDelta bool                          `json:"enableConstructDelta"`
	Profile              ProfileChangeProfileEntry[ST] `json:"profile"`
}

type ProfileChangeProfileEntry[ST ProfileStatsType] struct {
	Created         string          `json:"created"`
	Updated         string          `json:"updated"`
	RVN             int             `json:"rvn"`
	WipeNumber      int             `json:"wipeNumber"`
	AccountID       string          `json:"accountId"`
	ProfileID       string          `json:"profileId"`
	Version         string          `json:"version"` // migration
	Stats           ST              `json:"stats"`
	CommandRevision int             `json:"commandRevision"`
	ID              string          `json:"_id"`
	Items           map[string]Item `json:"items"`
}

type ProfileStatsType interface {
	AthenaProfileStats | CampaignProfileStats | CollectionBookPeopleProfileStats | CollectionBookSchematicsProfileStats | CollectionsProfileStats | CommonPublicProfileStats |
		CommonCoreProfileStats | CreativeProfileStats | MetadataProfileStats | OutpostProfileStats | RecycleBinProfileStats | Theater0ProfileStats | Theater1ProfileStats |
		Theater2ProfileStats | Profile0ProfileStats | ProtoJunoProfileStats
}

type ProfileNotificationType interface {
	CampaignNotification | map[string]any
}

type AthenaProfileStats struct {
	Attributes struct {
		UseRandomLoadout bool `json:"use_random_loadout"`
		PastSeasons      []struct {
			SeasonXP         int  `json:"seasonXp,omitempty"`
			SeasonLevel      int  `json:"seasonLevel"`
			BookLevel        int  `json:"bookLevel"`
			SeasonNumber     int  `json:"seasonNumber"`
			NumWins          int  `json:"numWins,omitempty"`
			PurchasedVIP     bool `json:"purchasedVIP,omitempty"`
			BookXP           int  `json:"bookXp,omitempty"`
			NumRoyalRoyales  int  `json:"numRoyalRoyales,omitempty"`
			SurvivorPrestige int  `json:"survivorPrestige,omitempty"`
			SurvivorTier     int  `json:"survivorTier,omitempty"`
		} `json:"past_seasons"`
		SeasonMatchBoost          int       `json:"season_match_boost"`
		RestedXPOverflow          int       `json:"rested_xp_overflow"`
		MfaRewardClaimed          bool      `json:"mfa_reward_claimed"`
		LastXPInteraction         time.Time `json:"last_xp_interaction"`
		RestedXpGoldenPathGranted int       `json:"rested_xp_golden_path_granted"`
		QuestManager              struct {
			DailyLoginInterval time.Time `json:"dailyLoginInterval"`
			DailyQuestRerolls  int       `json:"dailyQuestRerolls"`
		} `json:"quest_manager"`
		BookLevel            int    `json:"book_level"`
		SeasonNum            int    `json:"season_num"`
		LockerTwoPhaseCommit string `json:"locker_two_phase_commit"`
		BookXP               int    `json:"book_xp"`
		CreativeDynamicXP    struct {
			WeeklyExcessXPMult float64 `json:"weeklyExcessXpMult"`
			BankXP             int     `json:"bankXp"`
			BankXPMult         float64 `json:"bankXpMult"`
			CurrentWeekXP      int     `json:"currentWeekXp"`
			BoosterBucketXP    int     `json:"boosterBucketXp"`
			BoosterXPMult      float64 `json:"boosterXpMult"`
			CurrentWeek        int     `json:"currentWeek"`
			Timespan           float64 `json:"timespan"`
			BucketXP           int     `json:"bucketXp"`
		} `json:"creative_dynamic_xp"`
		Season struct {
			NumWins int `json:"numWins"`
		} `json:"season"`
		LockerServiceCosmeticItemsMigrationStatus string `json:"locker_service_cosmetic_items_migration_status"`
		VoteData                                  struct {
			ElectionID  string `json:"electionId"`
			VoteHistory struct {
				Vote7Mr570D4119Meh78Jo4I562God02 struct {
					VoteCount   int       `json:"voteCount"`
					FirstVoteAt time.Time `json:"firstVoteAt"`
					LastVoteAt  time.Time `json:"lastVoteAt"`
				} `json:"vote://7mr570d4119meh78jo4i562god[0]:2"`
			} `json:"voteHistory"`
			VotesRemaining  int       `json:"votesRemaining"`
			LastVoteGranted time.Time `json:"lastVoteGranted"`
		} `json:"vote_data"`
		LifetimeWins                  int    `json:"lifetime_wins"`
		PartyAssistQuest              string `json:"party_assist_quest"`
		PurchasedBattlePassTierOffers []struct {
			ID    string `json:"id"`
			Count int    `json:"count"`
		} `json:"purchased_battle_pass_tier_offers"`
		LockerLoadoutMigration string   `json:"locker_loadout_migration"`
		TrackedQuests          []string `json:"tracked_quests"`
		HabaneroUnlocked       bool     `json:"habanero_unlocked"`
		RestedXpExchange       float64  `json:"rested_xp_exchange"`
		PlaytimeXP             struct {
			CurrentWeekXP int `json:"currentWeekXp"`
			CurrentWeek   int `json:"currentWeek"`
		} `json:"playtime_xp"`
		Level                           int       `json:"level"`
		RestedXPMult                    float64   `json:"rested_xp_mult"`
		AccountLevel                    int       `json:"accountLevel"`
		RestedXPConsumedCumulative      int       `json:"rested_xp_consumed_cumulative"`
		RestedXPCumulative              int       `json:"rested_xp_cumulative"`
		XP                              int       `json:"xp"`
		SeasonFriendMatchBoost          int       `json:"season_friend_match_boost"`
		LastMatchEndDatetime            time.Time `json:"last_match_end_datetime"`
		LastStwAccoladeTransferDatetime time.Time `json:"last_stw_accolade_transfer_datetime"`
		PastSeasonPasses                []struct {
			Level            int    `json:"level"`
			SeasonTemplateID string `json:"seasonTemplateId"`
		} `json:"past_season_passes"`
		PastSeasonPurchaseContextHistories struct {
			HistoriesPerPassType struct {
				BR []struct {
					PurchaseContextHistory []struct {
						Level           int       `json:"level"`
						PurchaseContext string    `json:"purchaseContext"`
						Timestamp       time.Time `json:"timestamp"`
					} `json:"purchaseContextHistory"`
					SeasonTemplateID string `json:"seasonTemplateId"`
				} `json:"br"`
				Figment []struct {
					PurchaseContextHistory []struct {
						Level           int       `json:"level"`
						PurchaseContext string    `json:"purchaseContext"`
						Timestamp       time.Time `json:"timestamp"`
					} `json:"purchaseContextHistory"`
					SeasonTemplateID string `json:"seasonTemplateId"`
				} `json:"figment"`
				Musicpass []struct {
					PurchaseContextHistory []struct {
						Level           int       `json:"level"`
						PurchaseContext string    `json:"purchaseContext"`
						Timestamp       time.Time `json:"timestamp"`
					} `json:"purchaseContextHistory"`
					SeasonTemplateID string `json:"seasonTemplateId"`
				} `json:"musicpass"`
				Juno []struct {
					PurchaseContextHistory []struct {
						Level           int       `json:"level"`
						PurchaseContext string    `json:"purchaseContext"`
						Timestamp       time.Time `json:"timestamp"`
					} `json:"purchaseContextHistory"`
					SeasonTemplateID string `json:"seasonTemplateId"`
				} `json:"juno"`
			} `json:"historiesPerPassType"`
		} `json:"past_season_purchase_context_histories"`
	} `json:"attributes"`
}

type CampaignProfileStats fortnite.CampaignProfileStats

type CollectionBookPeopleProfileStats struct {
	Attributes map[string]any `json:"attributes"`
}

type CollectionBookSchematicsProfileStats struct {
	Attributes map[string]any `json:"attributes"`
}

type CollectionsProfileStats struct {
	Attributes struct {
		CurrentSeason int `json:"current_season"`
	} `json:"attributes"`
}

type CommonPublicProfileStats struct {
	Attributes struct {
		BannerColor  string `json:"banner_color"`
		HomebaseName string `json:"homebase_name"`
		BannerIcon   string `json:"banner_icon"`
	} `json:"attributes"`
}

type CommonCoreProfileStats fortnite.CommonCoreProfileStats

type CreativeProfileStats struct {
	Attributes struct {
		LastUsedProject string   `json:"last_used_project"`
		MaxIslandPlots  int      `json:"max_island_plots"`
		LastUsedPlot    string   `json:"last_used_plot"`
		Permissions     []string `json:"permissions"`
	} `json:"attributes"`
}

type MetadataProfileStats struct {
	Attributes map[string]any `json:"attributes"`
}

type OutpostProfileStats struct {
	Attributes map[string]any `json:"attributes"`
}

type RecycleBinProfileStats struct {
	Attributes struct {
		NextReceiptSequence int `json:"next_receipt_sequence"`
	} `json:"attributes"`
}

type ProtoJunoProfileStats struct {
	Attributes map[string]any `json:"attributes"`
}

type Profile0ProfileStats struct {
	Attributes map[string]any `json:"attributes"`
}

type Theater0ProfileStats struct {
	Attributes struct {
		PlayerLoadout struct {
			BPlayerIsNew          bool `json:"bPlayerIsNew"`
			PrimaryQuickBarRecord struct {
				Slots []struct {
					Items []string `json:"items"`
				} `json:"slots"`
			} `json:"primaryQuickBarRecord"`
			SecondaryQuickBarRecord struct {
				Slots []struct {
					Items []string `json:"items"`
				} `json:"slots"`
			} `json:"secondaryQuickBarRecord"`
			ZonesCompleted int `json:"zonesCompleted"`
		} `json:"player_loadout"`
	} `json:"attributes"`
}

type Theater1ProfileStats struct {
	Attributes struct {
		PlayerLoadout struct {
			BPlayerIsNew             bool          `json:"bPlayerIsNew"`
			PinnedSchematicInstances []interface{} `json:"pinnedSchematicInstances"`
			PrimaryQuickBarRecord    struct {
				CurrentFocusedSlot   int `json:"currentFocusedSlot"`
				PreviousFocusedSlot  int `json:"previousFocusedSlot"`
				SecondaryFocusedSlot int `json:"secondaryFocusedSlot"`
				Slots                []struct {
					Items []string `json:"items"`
				} `json:"slots"`
				DataDefinition struct {
					QuickbarSlots []struct {
						AcceptedItemTypes []string `json:"acceptedItemTypes"`
						BStaticSlot       bool     `json:"bStaticSlot"`
						DefaultItem       string   `json:"defaultItem"`
					} `json:"quickbarSlots"`
				} `json:"dataDefinition"`
			} `json:"primaryQuickBarRecord"`
			SecondaryQuickBarRecord struct {
				CurrentFocusedSlot   int `json:"currentFocusedSlot"`
				PreviousFocusedSlot  int `json:"previousFocusedSlot"`
				SecondaryFocusedSlot int `json:"secondaryFocusedSlot"`
				Slots                []struct {
					Items []string `json:"items"`
				} `json:"slots"`
				DataDefinition struct {
					QuickbarSlots []struct {
						AcceptedItemTypes []string `json:"acceptedItemTypes"`
						BStaticSlot       bool     `json:"bStaticSlot"`
						DefaultItem       string   `json:"defaultItem"`
					} `json:"quickbarSlots"`
				} `json:"dataDefinition"`
			} `json:"secondaryQuickBarRecord"`
			ZonesCompleted int `json:"zonesCompleted"`
		} `json:"player_loadout"`
	} `json:"attributes"`
}

type Theater2ProfileStats struct {
	Attributes struct {
		PlayerLoadout struct {
			BPlayerIsNew          bool `json:"bPlayerIsNew"`
			PrimaryQuickBarRecord struct {
				Slots []struct {
					Items []string `json:"items,omitempty"`
				} `json:"slots"`
			} `json:"primaryQuickBarRecord"`
			ZonesCompleted int `json:"zonesCompleted"`
		} `json:"player_loadout"`
		LastEventInstanceKey string `json:"last_event_instance_key"`
	} `json:"attributes"`
}

type CampaignNotification struct {
	Type         string `json:"type"`
	Primary      bool   `json:"primary"`
	DaysLoggedIn int    `json:"daysLoggedIn"`
	Items        []struct {
		ItemType    string `json:"itemType"`
		ItemGuid    string `json:"itemGuid"`
		ItemProfile string `json:"itemProfile"`
		Quantity    int    `json:"quantity"`
	} `json:"items"`
}

type AthenaCosmeticItem baseItem[AthenaCosmeticItemAttributes]

type AthenaCosmeticItemAttributes struct {
	CreationTime time.Time                             `json:"creation_time"`
	Level        int                                   `json:"level"`
	ItemSeen     bool                                  `json:"item_seen"`
	Variants     []AthenaCosmeticItemAttributesVariant `json:"variants,omitempty"`
}

// type AthenaCosmeticItem struct {
// 	TemplateID string `json:"templateId"`
// 	Attributes struct {
// 		Favorite      bool `json:"favorite,omitempty"`
// 		Archived      bool `json:"archived,omitempty"`
// 		ItemSeen      bool `json:"item_seen"`
// 		Level         int  `json:"level"`
// 		MaxLevelBonus int  `json:"max_level_bonus"`
// 		RndSelCnt     int  `json:"rnd_sel_cnt"`
// 		Variants      []struct {
// 			Channel string   `json:"channel"`
// 			Active  string   `json:"active"`
// 			Owned   []string `json:"owned"`
// 		} `json:"variants,omitempty"`
// 		XP int `json:"xp"`
// 	} `json:"attributes"`
// 	Quantity int `json:"quantity"`
// }

type AthenaCosmeticItemAttributesVariant struct {
	Channel string   `json:"channel"`
	Active  string   `json:"active"`
	Owned   []string `json:"owned"`
}

type AthenaQuestItem baseItem[AthenaQuestItemAttributes]

type AthenaQuestItemAttributes struct {
	CreationTime        time.Time      `json:"creation_time"`
	QuestState          string         `json:"quest_state"`
	LastStateChangeTime time.Time      `json:"last_state_change_time"`
	Level               int            `json:"level"`
	ChallengeBundleID   string         `json:"challenge_bundle_id"`
	QuestRarity         string         `json:"quest_rarity"`
	Extra               map[string]any `json:"-"`
}

func (qa *AthenaQuestItemAttributes) UnmarshalJSON(data []byte) error {
	var raw map[string]any
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	type Alias AthenaQuestItemAttributes
	aux := (*Alias)(qa)
	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	qa.Extra = make(map[string]any)
	knownFields := map[string]bool{
		"creation_time": true, "quest_state": true, "last_state_change_time": true,
		"level": true, "challenge_bundle_id": true, "quest_rarity": true,
	}

	for k, v := range raw {
		if !knownFields[k] {
			qa.Extra[k] = v
		}
	}

	return nil
}

type RewardGraphItem baseItem[RewardGraphItemAttributes]

type RewardGraphItemAttributes struct {
	UnlockEpoch                   time.Time   `json:"unlock_epoch"`
	PlayerRandomSeed              int         `json:"player_random_seed"`
	RewardGraphPurchasedTimestamp int64       `json:"reward_graph_purchased_timestamp"`
	RewardGraphPurchased          bool        `json:"reward_graph_purchased"`
	RewardKeys                    []RewardKey `json:"reward_keys"`
}

type RewardKey struct {
	StaticKeyTemplateID string `json:"static_key_template_id"`
	UnlockKeysUsed      int    `json:"unlock_keys_used"`
	KeysGrantedToday    int    `json:"keys_granted_today"`
}
