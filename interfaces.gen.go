package d2protocol

type FightResultListEntryIntrf interface {
	Type

	GetOutcome() uint16

	GetWave() uint8

	GetRewards() FightLoot
}

func (m *FightResultListEntry) GetOutcome() uint16 {
	return m.Outcome
}

func (m *FightResultListEntry) GetWave() uint8 {
	return m.Wave
}

func (m *FightResultListEntry) GetRewards() FightLoot {
	return m.Rewards
}

type GameFightMinimalStatsIntrf interface {
	Type

	GetLifePoints() uint32

	GetMaxLifePoints() uint32

	GetBaseMaxLifePoints() uint32

	GetPermanentDamagePercent() uint32

	GetShieldPoints() uint32

	GetActionPoints() int16

	GetMaxActionPoints() int16

	GetMovementPoints() int16

	GetMaxMovementPoints() int16

	GetSummoner() float64

	GetSummoned() bool

	GetNeutralElementResistPercent() int16

	GetEarthElementResistPercent() int16

	GetWaterElementResistPercent() int16

	GetAirElementResistPercent() int16

	GetFireElementResistPercent() int16

	GetNeutralElementReduction() int16

	GetEarthElementReduction() int16

	GetWaterElementReduction() int16

	GetAirElementReduction() int16

	GetFireElementReduction() int16

	GetCriticalDamageFixedResist() int16

	GetPushDamageFixedResist() int16

	GetPvpNeutralElementResistPercent() int16

	GetPvpEarthElementResistPercent() int16

	GetPvpWaterElementResistPercent() int16

	GetPvpAirElementResistPercent() int16

	GetPvpFireElementResistPercent() int16

	GetPvpNeutralElementReduction() int16

	GetPvpEarthElementReduction() int16

	GetPvpWaterElementReduction() int16

	GetPvpAirElementReduction() int16

	GetPvpFireElementReduction() int16

	GetDodgePALostProbability() uint16

	GetDodgePMLostProbability() uint16

	GetTackleBlock() int16

	GetTackleEvade() int16

	GetFixedDamageReflection() int16

	GetInvisibilityState() uint8
}

func (m *GameFightMinimalStats) GetLifePoints() uint32 {
	return m.LifePoints
}

func (m *GameFightMinimalStats) GetMaxLifePoints() uint32 {
	return m.MaxLifePoints
}

func (m *GameFightMinimalStats) GetBaseMaxLifePoints() uint32 {
	return m.BaseMaxLifePoints
}

func (m *GameFightMinimalStats) GetPermanentDamagePercent() uint32 {
	return m.PermanentDamagePercent
}

func (m *GameFightMinimalStats) GetShieldPoints() uint32 {
	return m.ShieldPoints
}

func (m *GameFightMinimalStats) GetActionPoints() int16 {
	return m.ActionPoints
}

func (m *GameFightMinimalStats) GetMaxActionPoints() int16 {
	return m.MaxActionPoints
}

func (m *GameFightMinimalStats) GetMovementPoints() int16 {
	return m.MovementPoints
}

func (m *GameFightMinimalStats) GetMaxMovementPoints() int16 {
	return m.MaxMovementPoints
}

func (m *GameFightMinimalStats) GetSummoner() float64 {
	return m.Summoner
}

func (m *GameFightMinimalStats) GetSummoned() bool {
	return m.Summoned
}

func (m *GameFightMinimalStats) GetNeutralElementResistPercent() int16 {
	return m.NeutralElementResistPercent
}

func (m *GameFightMinimalStats) GetEarthElementResistPercent() int16 {
	return m.EarthElementResistPercent
}

func (m *GameFightMinimalStats) GetWaterElementResistPercent() int16 {
	return m.WaterElementResistPercent
}

func (m *GameFightMinimalStats) GetAirElementResistPercent() int16 {
	return m.AirElementResistPercent
}

func (m *GameFightMinimalStats) GetFireElementResistPercent() int16 {
	return m.FireElementResistPercent
}

func (m *GameFightMinimalStats) GetNeutralElementReduction() int16 {
	return m.NeutralElementReduction
}

func (m *GameFightMinimalStats) GetEarthElementReduction() int16 {
	return m.EarthElementReduction
}

func (m *GameFightMinimalStats) GetWaterElementReduction() int16 {
	return m.WaterElementReduction
}

func (m *GameFightMinimalStats) GetAirElementReduction() int16 {
	return m.AirElementReduction
}

func (m *GameFightMinimalStats) GetFireElementReduction() int16 {
	return m.FireElementReduction
}

func (m *GameFightMinimalStats) GetCriticalDamageFixedResist() int16 {
	return m.CriticalDamageFixedResist
}

func (m *GameFightMinimalStats) GetPushDamageFixedResist() int16 {
	return m.PushDamageFixedResist
}

func (m *GameFightMinimalStats) GetPvpNeutralElementResistPercent() int16 {
	return m.PvpNeutralElementResistPercent
}

func (m *GameFightMinimalStats) GetPvpEarthElementResistPercent() int16 {
	return m.PvpEarthElementResistPercent
}

func (m *GameFightMinimalStats) GetPvpWaterElementResistPercent() int16 {
	return m.PvpWaterElementResistPercent
}

func (m *GameFightMinimalStats) GetPvpAirElementResistPercent() int16 {
	return m.PvpAirElementResistPercent
}

func (m *GameFightMinimalStats) GetPvpFireElementResistPercent() int16 {
	return m.PvpFireElementResistPercent
}

func (m *GameFightMinimalStats) GetPvpNeutralElementReduction() int16 {
	return m.PvpNeutralElementReduction
}

func (m *GameFightMinimalStats) GetPvpEarthElementReduction() int16 {
	return m.PvpEarthElementReduction
}

func (m *GameFightMinimalStats) GetPvpWaterElementReduction() int16 {
	return m.PvpWaterElementReduction
}

func (m *GameFightMinimalStats) GetPvpAirElementReduction() int16 {
	return m.PvpAirElementReduction
}

func (m *GameFightMinimalStats) GetPvpFireElementReduction() int16 {
	return m.PvpFireElementReduction
}

func (m *GameFightMinimalStats) GetDodgePALostProbability() uint16 {
	return m.DodgePALostProbability
}

func (m *GameFightMinimalStats) GetDodgePMLostProbability() uint16 {
	return m.DodgePMLostProbability
}

func (m *GameFightMinimalStats) GetTackleBlock() int16 {
	return m.TackleBlock
}

func (m *GameFightMinimalStats) GetTackleEvade() int16 {
	return m.TackleEvade
}

func (m *GameFightMinimalStats) GetFixedDamageReflection() int16 {
	return m.FixedDamageReflection
}

func (m *GameFightMinimalStats) GetInvisibilityState() uint8 {
	return m.InvisibilityState
}

type FightTeamInformationsIntrf interface {
	Type

	GetAbstractFightTeamInformations() AbstractFightTeamInformations

	GetTeamMembers() []FightTeamMemberInformationsIntrf
}

func (m *FightTeamInformations) GetAbstractFightTeamInformations() AbstractFightTeamInformations {
	return m.AbstractFightTeamInformations
}

func (m *FightTeamInformations) GetTeamMembers() []FightTeamMemberInformationsIntrf {
	return m.TeamMembers
}

type FightTeamMemberInformationsIntrf interface {
	Type

	GetId() float64
}

func (m *FightTeamMemberInformations) GetId() float64 {
	return m.Id
}

type CharacterBaseInformationsIntrf interface {
	Type

	GetCharacterMinimalPlusLookInformations() CharacterMinimalPlusLookInformations

	GetBreed() int8

	GetSex() bool
}

func (m *CharacterBaseInformations) GetCharacterMinimalPlusLookInformations() CharacterMinimalPlusLookInformations {
	return m.CharacterMinimalPlusLookInformations
}

func (m *CharacterBaseInformations) GetBreed() int8 {
	return m.Breed
}

func (m *CharacterBaseInformations) GetSex() bool {
	return m.Sex
}

type EntityDispositionInformationsIntrf interface {
	Type

	GetCellId() int16

	GetDirection() uint8
}

func (m *EntityDispositionInformations) GetCellId() int16 {
	return m.CellId
}

func (m *EntityDispositionInformations) GetDirection() uint8 {
	return m.Direction
}

type ObjectEffectIntrf interface {
	Type

	GetActionId() uint16
}

func (m *ObjectEffect) GetActionId() uint16 {
	return m.ActionId
}

type FriendSpouseInformationsIntrf interface {
	Type

	GetSpouseAccountId() uint32

	GetSpouseId() int64

	GetSpouseName() string

	GetSpouseLevel() uint8

	GetBreed() int8

	GetSex() int8

	GetSpouseEntityLook() EntityLook

	GetGuildInfo() GuildInformations

	GetAlignmentSide() int8
}

func (m *FriendSpouseInformations) GetSpouseAccountId() uint32 {
	return m.SpouseAccountId
}

func (m *FriendSpouseInformations) GetSpouseId() int64 {
	return m.SpouseId
}

func (m *FriendSpouseInformations) GetSpouseName() string {
	return m.SpouseName
}

func (m *FriendSpouseInformations) GetSpouseLevel() uint8 {
	return m.SpouseLevel
}

func (m *FriendSpouseInformations) GetBreed() int8 {
	return m.Breed
}

func (m *FriendSpouseInformations) GetSex() int8 {
	return m.Sex
}

func (m *FriendSpouseInformations) GetSpouseEntityLook() EntityLook {
	return m.SpouseEntityLook
}

func (m *FriendSpouseInformations) GetGuildInfo() GuildInformations {
	return m.GuildInfo
}

func (m *FriendSpouseInformations) GetAlignmentSide() int8 {
	return m.AlignmentSide
}

type FriendInformationsIntrf interface {
	Type

	GetAbstractContactInformations() AbstractContactInformations

	GetPlayerState() uint8

	GetLastConnection() uint16

	GetAchievementPoints() int32
}

func (m *FriendInformations) GetAbstractContactInformations() AbstractContactInformations {
	return m.AbstractContactInformations
}

func (m *FriendInformations) GetPlayerState() uint8 {
	return m.PlayerState
}

func (m *FriendInformations) GetLastConnection() uint16 {
	return m.LastConnection
}

func (m *FriendInformations) GetAchievementPoints() int32 {
	return m.AchievementPoints
}

type InteractiveElementIntrf interface {
	Type

	GetElementId() uint32

	GetElementTypeId() int32

	GetEnabledSkills() []InteractiveElementSkillIntrf

	GetDisabledSkills() []InteractiveElementSkillIntrf

	GetOnCurrentMap() bool
}

func (m *InteractiveElement) GetElementId() uint32 {
	return m.ElementId
}

func (m *InteractiveElement) GetElementTypeId() int32 {
	return m.ElementTypeId
}

func (m *InteractiveElement) GetEnabledSkills() []InteractiveElementSkillIntrf {
	return m.EnabledSkills
}

func (m *InteractiveElement) GetDisabledSkills() []InteractiveElementSkillIntrf {
	return m.DisabledSkills
}

func (m *InteractiveElement) GetOnCurrentMap() bool {
	return m.OnCurrentMap
}

type PartyMemberInformationsIntrf interface {
	Type

	GetCharacterBaseInformations() CharacterBaseInformations

	GetLifePoints() uint32

	GetMaxLifePoints() uint32

	GetProspecting() uint16

	GetRegenRate() uint8

	GetInitiative() uint16

	GetAlignmentSide() int8

	GetWorldX() int16

	GetWorldY() int16

	GetMapId() int32

	GetSubAreaId() uint16

	GetStatus() PlayerStatusIntrf

	GetCompanions() []PartyCompanionMemberInformations
}

func (m *PartyMemberInformations) GetCharacterBaseInformations() CharacterBaseInformations {
	return m.CharacterBaseInformations
}

func (m *PartyMemberInformations) GetLifePoints() uint32 {
	return m.LifePoints
}

func (m *PartyMemberInformations) GetMaxLifePoints() uint32 {
	return m.MaxLifePoints
}

func (m *PartyMemberInformations) GetProspecting() uint16 {
	return m.Prospecting
}

func (m *PartyMemberInformations) GetRegenRate() uint8 {
	return m.RegenRate
}

func (m *PartyMemberInformations) GetInitiative() uint16 {
	return m.Initiative
}

func (m *PartyMemberInformations) GetAlignmentSide() int8 {
	return m.AlignmentSide
}

func (m *PartyMemberInformations) GetWorldX() int16 {
	return m.WorldX
}

func (m *PartyMemberInformations) GetWorldY() int16 {
	return m.WorldY
}

func (m *PartyMemberInformations) GetMapId() int32 {
	return m.MapId
}

func (m *PartyMemberInformations) GetSubAreaId() uint16 {
	return m.SubAreaId
}

func (m *PartyMemberInformations) GetStatus() PlayerStatusIntrf {
	return m.Status
}

func (m *PartyMemberInformations) GetCompanions() []PartyCompanionMemberInformations {
	return m.Companions
}

type SkillActionDescriptionIntrf interface {
	Type

	GetSkillId() uint16
}

func (m *SkillActionDescription) GetSkillId() uint16 {
	return m.SkillId
}

type IgnoredInformationsIntrf interface {
	Type

	GetAbstractContactInformations() AbstractContactInformations
}

func (m *IgnoredInformations) GetAbstractContactInformations() AbstractContactInformations {
	return m.AbstractContactInformations
}

type HouseInformationsIntrf interface {
	Type

	GetHouseId() uint32

	GetModelId() uint16
}

func (m *HouseInformations) GetHouseId() uint32 {
	return m.HouseId
}

func (m *HouseInformations) GetModelId() uint16 {
	return m.ModelId
}

type PaddockBuyableInformationsIntrf interface {
	Type

	GetPrice() int64

	GetLocked() bool
}

func (m *PaddockBuyableInformations) GetPrice() int64 {
	return m.Price
}

func (m *PaddockBuyableInformations) GetLocked() bool {
	return m.Locked
}

type GroupMonsterStaticInformationsIntrf interface {
	Type

	GetMainCreatureLightInfos() MonsterInGroupLightInformations

	GetUnderlings() []MonsterInGroupInformations
}

func (m *GroupMonsterStaticInformations) GetMainCreatureLightInfos() MonsterInGroupLightInformations {
	return m.MainCreatureLightInfos
}

func (m *GroupMonsterStaticInformations) GetUnderlings() []MonsterInGroupInformations {
	return m.Underlings
}

type GameRolePlayActorInformationsIntrf interface {
	Type

	GetGameContextActorInformations() GameContextActorInformations
}

func (m *GameRolePlayActorInformations) GetGameContextActorInformations() GameContextActorInformations {
	return m.GameContextActorInformations
}

type GameFightFighterInformationsIntrf interface {
	Type

	GetGameContextActorInformations() GameContextActorInformations

	GetTeamId() uint8

	GetWave() uint8

	GetAlive() bool

	GetStats() GameFightMinimalStatsIntrf

	GetPreviousPositions() []uint16
}

func (m *GameFightFighterInformations) GetGameContextActorInformations() GameContextActorInformations {
	return m.GameContextActorInformations
}

func (m *GameFightFighterInformations) GetTeamId() uint8 {
	return m.TeamId
}

func (m *GameFightFighterInformations) GetWave() uint8 {
	return m.Wave
}

func (m *GameFightFighterInformations) GetAlive() bool {
	return m.Alive
}

func (m *GameFightFighterInformations) GetStats() GameFightMinimalStatsIntrf {
	return m.Stats
}

func (m *GameFightFighterInformations) GetPreviousPositions() []uint16 {
	return m.PreviousPositions
}

type TaxCollectorStaticInformationsIntrf interface {
	Type

	GetFirstNameId() uint16

	GetLastNameId() uint16

	GetGuildIdentity() GuildInformations
}

func (m *TaxCollectorStaticInformations) GetFirstNameId() uint16 {
	return m.FirstNameId
}

func (m *TaxCollectorStaticInformations) GetLastNameId() uint16 {
	return m.LastNameId
}

func (m *TaxCollectorStaticInformations) GetGuildIdentity() GuildInformations {
	return m.GuildIdentity
}

type GameContextActorInformationsIntrf interface {
	Type

	GetContextualId() float64

	GetLook() EntityLook

	GetDisposition() EntityDispositionInformationsIntrf
}

func (m *GameContextActorInformations) GetContextualId() float64 {
	return m.ContextualId
}

func (m *GameContextActorInformations) GetLook() EntityLook {
	return m.Look
}

func (m *GameContextActorInformations) GetDisposition() EntityDispositionInformationsIntrf {
	return m.Disposition
}

type HumanInformationsIntrf interface {
	Type

	GetRestrictions() ActorRestrictionsInformations

	GetSex() bool

	GetOptions() []HumanOptionIntrf
}

func (m *HumanInformations) GetRestrictions() ActorRestrictionsInformations {
	return m.Restrictions
}

func (m *HumanInformations) GetSex() bool {
	return m.Sex
}

func (m *HumanInformations) GetOptions() []HumanOptionIntrf {
	return m.Options
}

type CharacterMinimalPlusLookInformationsIntrf interface {
	Type

	GetCharacterMinimalInformations() CharacterMinimalInformations

	GetEntityLook() EntityLook
}

func (m *CharacterMinimalPlusLookInformations) GetCharacterMinimalInformations() CharacterMinimalInformations {
	return m.CharacterMinimalInformations
}

func (m *CharacterMinimalPlusLookInformations) GetEntityLook() EntityLook {
	return m.EntityLook
}

type TaxCollectorInformationsIntrf interface {
	Type

	GetUniqueId() int32

	GetFirtNameId() uint16

	GetLastNameId() uint16

	GetAdditionalInfos() AdditionalTaxCollectorInformations

	GetWorldX() int16

	GetWorldY() int16

	GetSubAreaId() uint16

	GetState() uint8

	GetLook() EntityLook

	GetComplements() []TaxCollectorComplementaryInformationsIntrf
}

func (m *TaxCollectorInformations) GetUniqueId() int32 {
	return m.UniqueId
}

func (m *TaxCollectorInformations) GetFirtNameId() uint16 {
	return m.FirtNameId
}

func (m *TaxCollectorInformations) GetLastNameId() uint16 {
	return m.LastNameId
}

func (m *TaxCollectorInformations) GetAdditionalInfos() AdditionalTaxCollectorInformations {
	return m.AdditionalInfos
}

func (m *TaxCollectorInformations) GetWorldX() int16 {
	return m.WorldX
}

func (m *TaxCollectorInformations) GetWorldY() int16 {
	return m.WorldY
}

func (m *TaxCollectorInformations) GetSubAreaId() uint16 {
	return m.SubAreaId
}

func (m *TaxCollectorInformations) GetState() uint8 {
	return m.State
}

func (m *TaxCollectorInformations) GetLook() EntityLook {
	return m.Look
}

func (m *TaxCollectorInformations) GetComplements() []TaxCollectorComplementaryInformationsIntrf {
	return m.Complements
}

type MapCoordinatesIntrf interface {
	Type

	GetWorldX() int16

	GetWorldY() int16
}

func (m *MapCoordinates) GetWorldX() int16 {
	return m.WorldX
}

func (m *MapCoordinates) GetWorldY() int16 {
	return m.WorldY
}

type FightResultAdditionalDataIntrf interface {
	Type
}

type AbstractFightDispellableEffectIntrf interface {
	Type

	GetUid() uint32

	GetTargetId() float64

	GetTurnDuration() int16

	GetDispelable() uint8

	GetSpellId() uint16

	GetEffectId() uint32

	GetParentBoostUid() uint32
}

func (m *AbstractFightDispellableEffect) GetUid() uint32 {
	return m.Uid
}

func (m *AbstractFightDispellableEffect) GetTargetId() float64 {
	return m.TargetId
}

func (m *AbstractFightDispellableEffect) GetTurnDuration() int16 {
	return m.TurnDuration
}

func (m *AbstractFightDispellableEffect) GetDispelable() uint8 {
	return m.Dispelable
}

func (m *AbstractFightDispellableEffect) GetSpellId() uint16 {
	return m.SpellId
}

func (m *AbstractFightDispellableEffect) GetEffectId() uint32 {
	return m.EffectId
}

func (m *AbstractFightDispellableEffect) GetParentBoostUid() uint32 {
	return m.ParentBoostUid
}

type InteractiveElementSkillIntrf interface {
	Type

	GetSkillId() uint32

	GetSkillInstanceUid() uint32
}

func (m *InteractiveElementSkill) GetSkillId() uint32 {
	return m.SkillId
}

func (m *InteractiveElementSkill) GetSkillInstanceUid() uint32 {
	return m.SkillInstanceUid
}

type UpdateMountBoostIntrf interface {
	Type

	GetType() uint8
}

func (m *UpdateMountBoost) GetType() uint8 {
	return m.Type
}

type ShortcutIntrf interface {
	Type

	GetSlot() uint8
}

func (m *Shortcut) GetSlot() uint8 {
	return m.Slot
}

type QuestActiveInformationsIntrf interface {
	Type

	GetQuestId() uint16
}

func (m *QuestActiveInformations) GetQuestId() uint16 {
	return m.QuestId
}

type QuestObjectiveInformationsIntrf interface {
	Type

	GetObjectiveId() uint16

	GetObjectiveStatus() bool

	GetDialogParams() []string
}

func (m *QuestObjectiveInformations) GetObjectiveId() uint16 {
	return m.ObjectiveId
}

func (m *QuestObjectiveInformations) GetObjectiveStatus() bool {
	return m.ObjectiveStatus
}

func (m *QuestObjectiveInformations) GetDialogParams() []string {
	return m.DialogParams
}

type HumanOptionIntrf interface {
	Type
}

type GameFightFighterLightInformationsIntrf interface {
	Type

	GetId() float64

	GetWave() uint8

	GetLevel() uint16

	GetBreed() int8

	GetSex() bool

	GetAlive() bool
}

func (m *GameFightFighterLightInformations) GetId() float64 {
	return m.Id
}

func (m *GameFightFighterLightInformations) GetWave() uint8 {
	return m.Wave
}

func (m *GameFightFighterLightInformations) GetLevel() uint16 {
	return m.Level
}

func (m *GameFightFighterLightInformations) GetBreed() int8 {
	return m.Breed
}

func (m *GameFightFighterLightInformations) GetSex() bool {
	return m.Sex
}

func (m *GameFightFighterLightInformations) GetAlive() bool {
	return m.Alive
}

type PlayerStatusIntrf interface {
	Type

	GetStatusId() uint8
}

func (m *PlayerStatus) GetStatusId() uint8 {
	return m.StatusId
}

type AbstractSocialGroupInfosIntrf interface {
	Type
}

type AllianceFactSheetInformationsIntrf interface {
	Type

	GetAllianceInformations() AllianceInformations

	GetCreationDate() uint32
}

func (m *AllianceFactSheetInformations) GetAllianceInformations() AllianceInformations {
	return m.AllianceInformations
}

func (m *AllianceFactSheetInformations) GetCreationDate() uint32 {
	return m.CreationDate
}

type GuildFactSheetInformationsIntrf interface {
	Type

	GetGuildInformations() GuildInformations

	GetLeaderId() int64

	GetNbMembers() uint16
}

func (m *GuildFactSheetInformations) GetGuildInformations() GuildInformations {
	return m.GuildInformations
}

func (m *GuildFactSheetInformations) GetLeaderId() int64 {
	return m.LeaderId
}

func (m *GuildFactSheetInformations) GetNbMembers() uint16 {
	return m.NbMembers
}

type PrismInformationIntrf interface {
	Type

	GetTypeId() uint8

	GetState() uint8

	GetNextVulnerabilityDate() uint32

	GetPlacementDate() uint32

	GetRewardTokenCount() uint32
}

func (m *PrismInformation) GetTypeId() uint8 {
	return m.TypeId
}

func (m *PrismInformation) GetState() uint8 {
	return m.State
}

func (m *PrismInformation) GetNextVulnerabilityDate() uint32 {
	return m.NextVulnerabilityDate
}

func (m *PrismInformation) GetPlacementDate() uint32 {
	return m.PlacementDate
}

func (m *PrismInformation) GetRewardTokenCount() uint32 {
	return m.RewardTokenCount
}

type ServerSessionConstantIntrf interface {
	Type

	GetId() uint16
}

func (m *ServerSessionConstant) GetId() uint16 {
	return m.Id
}

type GuildVersatileInformationsIntrf interface {
	Type

	GetGuildId() uint32

	GetLeaderId() int64

	GetGuildLevel() uint8

	GetNbMembers() uint8
}

func (m *GuildVersatileInformations) GetGuildId() uint32 {
	return m.GuildId
}

func (m *GuildVersatileInformations) GetLeaderId() int64 {
	return m.LeaderId
}

func (m *GuildVersatileInformations) GetGuildLevel() uint8 {
	return m.GuildLevel
}

func (m *GuildVersatileInformations) GetNbMembers() uint8 {
	return m.NbMembers
}

type PrismSubareaEmptyInfoIntrf interface {
	Type

	GetSubAreaId() uint16

	GetAllianceId() uint32
}

func (m *PrismSubareaEmptyInfo) GetSubAreaId() uint16 {
	return m.SubAreaId
}

func (m *PrismSubareaEmptyInfo) GetAllianceId() uint32 {
	return m.AllianceId
}

type TaxCollectorComplementaryInformationsIntrf interface {
	Type
}

type TreasureHuntStepIntrf interface {
	Type
}

type PortalInformationIntrf interface {
	Type

	GetPortalId() int32

	GetAreaId() int16
}

func (m *PortalInformation) GetPortalId() int32 {
	return m.PortalId
}

func (m *PortalInformation) GetAreaId() int16 {
	return m.AreaId
}

type StatisticDataIntrf interface {
	Type
}

type IdolIntrf interface {
	Type

	GetId() uint16

	GetXpBonusPercent() uint16

	GetDropBonusPercent() uint16
}

func (m *Idol) GetId() uint16 {
	return m.Id
}

func (m *Idol) GetXpBonusPercent() uint16 {
	return m.XpBonusPercent
}

func (m *Idol) GetDropBonusPercent() uint16 {
	return m.DropBonusPercent
}

type PartyIdolIntrf interface {
	Type

	GetIdol() Idol

	GetOwnersIds() []int64
}

func (m *PartyIdol) GetIdol() Idol {
	return m.Idol
}

func (m *PartyIdol) GetOwnersIds() []int64 {
	return m.OwnersIds
}

type HouseInstanceInformationsIntrf interface {
	Type

	GetInstanceId() uint32

	GetSecondHand() bool

	GetOwnerName() string

	GetIsOnSale() bool

	GetIsSaleLocked() bool
}

func (m *HouseInstanceInformations) GetInstanceId() uint32 {
	return m.InstanceId
}

func (m *HouseInstanceInformations) GetSecondHand() bool {
	return m.SecondHand
}

func (m *HouseInstanceInformations) GetOwnerName() string {
	return m.OwnerName
}

func (m *HouseInstanceInformations) GetIsOnSale() bool {
	return m.IsOnSale
}

func (m *HouseInstanceInformations) GetIsSaleLocked() bool {
	return m.IsSaleLocked
}
