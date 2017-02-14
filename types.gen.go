package d2protocol

import (
	"reflect"
)

var types = map[uint16]reflect.Type{

	350: reflect.TypeOf((*ContentPart)(nil)),

	150: reflect.TypeOf((*GameContextActorInformations)(nil)),

	143: reflect.TypeOf((*GameFightFighterInformations)(nil)),

	151: reflect.TypeOf((*GameFightAIInformations)(nil)),

	29: reflect.TypeOf((*GameFightMonsterInformations)(nil)),

	141: reflect.TypeOf((*GameRolePlayActorInformations)(nil)),

	154: reflect.TypeOf((*GameRolePlayNamedActorInformations)(nil)),

	180: reflect.TypeOf((*GameRolePlayMountInformations)(nil)),

	80: reflect.TypeOf((*InteractiveElement)(nil)),

	467: reflect.TypeOf((*GameRolePlayPortalInformations)(nil)),

	160: reflect.TypeOf((*GameRolePlayGroupMonsterInformations)(nil)),

	464: reflect.TypeOf((*GameRolePlayGroupMonsterWaveInformations)(nil)),

	161: reflect.TypeOf((*GameRolePlayPrismInformations)(nil)),

	48: reflect.TypeOf((*GameFightTaxCollectorInformations)(nil)),

	156: reflect.TypeOf((*GameRolePlayNpcInformations)(nil)),

	148: reflect.TypeOf((*GameRolePlayTaxCollectorInformations)(nil)),

	158: reflect.TypeOf((*GameFightFighterNamedInformations)(nil)),

	46: reflect.TypeOf((*GameFightCharacterInformations)(nil)),

	450: reflect.TypeOf((*GameFightCompanionInformations)(nil)),

	159: reflect.TypeOf((*GameRolePlayHumanoidInformations)(nil)),

	36: reflect.TypeOf((*GameRolePlayCharacterInformations)(nil)),

	383: reflect.TypeOf((*GameRolePlayNpcWithQuestInformations)(nil)),

	129: reflect.TypeOf((*GameRolePlayMerchantInformations)(nil)),

	3: reflect.TypeOf((*GameRolePlayMutantInformations)(nil)),

	54: reflect.TypeOf((*SubEntity)(nil)),

	55: reflect.TypeOf((*EntityLook)(nil)),

	8: reflect.TypeOf((*CharacterCharacteristicsInformations)(nil)),

	428: reflect.TypeOf((*PrismInformation)(nil)),

	427: reflect.TypeOf((*AlliancePrismInformation)(nil)),

	505: reflect.TypeOf((*DareReward)(nil)),

	431: reflect.TypeOf((*AllianceInsiderPrismInformation)(nil)),

	400: reflect.TypeOf((*AbstractCharacterInformation)(nil)),

	503: reflect.TypeOf((*CharacterBasicMinimalInformations)(nil)),

	110: reflect.TypeOf((*CharacterMinimalInformations)(nil)),

	88: reflect.TypeOf((*GuildMember)(nil)),

	60: reflect.TypeOf((*EntityDispositionInformations)(nil)),

	31: reflect.TypeOf((*GameFightMinimalStats)(nil)),

	219: reflect.TypeOf((*InteractiveElementSkill)(nil)),

	76: reflect.TypeOf((*ObjectEffect)(nil)),

	70: reflect.TypeOf((*ObjectEffectInteger)(nil)),

	466: reflect.TypeOf((*PortalInformation)(nil)),

	140: reflect.TypeOf((*GroupMonsterStaticInformations)(nil)),

	412: reflect.TypeOf((*AchievementRewardable)(nil)),

	381: reflect.TypeOf((*QuestActiveInformations)(nil)),

	101: reflect.TypeOf((*JobDescription)(nil)),

	102: reflect.TypeOf((*SkillActionDescription)(nil)),

	100: reflect.TypeOf((*SkillActionDescriptionCraft)(nil)),

	103: reflect.TypeOf((*SkillActionDescriptionTimed)(nil)),

	99: reflect.TypeOf((*SkillActionDescriptionCollect)(nil)),

	204: reflect.TypeOf((*ActorRestrictionsInformations)(nil)),

	406: reflect.TypeOf((*HumanOption)(nil)),

	408: reflect.TypeOf((*HumanOptionTitle)(nil)),

	163: reflect.TypeOf((*CharacterMinimalPlusLookInformations)(nil)),

	45: reflect.TypeOf((*CharacterBaseInformations)(nil)),

	215: reflect.TypeOf((*CharacterSpellModification)(nil)),

	111: reflect.TypeOf((*HouseInformations)(nil)),

	390: reflect.TypeOf((*AccountHouseInformations)(nil)),

	147: reflect.TypeOf((*TaxCollectorStaticInformations)(nil)),

	4: reflect.TypeOf((*CharacterBaseCharacteristic)(nil)),

	25: reflect.TypeOf((*GameServerInformations)(nil)),

	415: reflect.TypeOf((*PlayerStatus)(nil)),

	201: reflect.TypeOf((*ActorAlignmentInformations)(nil)),

	411: reflect.TypeOf((*HumanOptionOrnament)(nil)),

	168: reflect.TypeOf((*ItemDurability)(nil)),

	157: reflect.TypeOf((*HumanInformations)(nil)),

	395: reflect.TypeOf((*MonsterInGroupLightInformations)(nil)),

	396: reflect.TypeOf((*GroupMonsterStaticInformationsWithAlternatives)(nil)),

	144: reflect.TypeOf((*MonsterInGroupInformations)(nil)),

	394: reflect.TypeOf((*AlternativeMonstersInGroupLightInformations)(nil)),

	132: reflect.TypeOf((*PaddockInformations)(nil)),

	509: reflect.TypeOf((*PaddockInstancesInformations)(nil)),

	130: reflect.TypeOf((*PaddockBuyableInformations)(nil)),

	218: reflect.TypeOf((*HouseInformationsInside)(nil)),

	510: reflect.TypeOf((*HouseOnMapInformations)(nil)),

	511: reflect.TypeOf((*HouseInstanceInformations)(nil)),

	384: reflect.TypeOf((*GameRolePlayNpcQuestFlag)(nil)),

	373: reflect.TypeOf((*DungeonPartyFinderPlayer)(nil)),

	116: reflect.TypeOf((*AbstractFightTeamInformations)(nil)),

	33: reflect.TypeOf((*FightTeamInformations)(nil)),

	376: reflect.TypeOf((*PartyInvitationMemberInformations)(nil)),

	90: reflect.TypeOf((*PartyMemberInformations)(nil)),

	453: reflect.TypeOf((*PartyCompanionBaseInformations)(nil)),

	44: reflect.TypeOf((*FightTeamMemberInformations)(nil)),

	374: reflect.TypeOf((*PartyGuestInformations)(nil)),

	452: reflect.TypeOf((*PartyCompanionMemberInformations)(nil)),

	43: reflect.TypeOf((*FightCommonInformations)(nil)),

	391: reflect.TypeOf((*PartyMemberArenaInformations)(nil)),

	170: reflect.TypeOf((*HouseInformationsForGuild)(nil)),

	7: reflect.TypeOf((*Item)(nil)),

	483: reflect.TypeOf((*ObjectItemGenericQuantity)(nil)),

	380: reflect.TypeOf((*AbstractContactInformations)(nil)),

	106: reflect.TypeOf((*IgnoredInformations)(nil)),

	105: reflect.TypeOf((*IgnoredOnlineInformations)(nil)),

	78: reflect.TypeOf((*FriendInformations)(nil)),

	92: reflect.TypeOf((*FriendOnlineInformations)(nil)),

	183: reflect.TypeOf((*PaddockContentInformations)(nil)),

	493: reflect.TypeOf((*TaxCollectorMovement)(nil)),

	167: reflect.TypeOf((*TaxCollectorInformations)(nil)),

	414: reflect.TypeOf((*PlayerStatusExtended)(nil)),

	416: reflect.TypeOf((*AbstractSocialGroupInfos)(nil)),

	365: reflect.TypeOf((*BasicGuildInformations)(nil)),

	127: reflect.TypeOf((*GuildInformations)(nil)),

	424: reflect.TypeOf((*GuildFactSheetInformations)(nil)),

	423: reflect.TypeOf((*GuildInsiderFactSheetInformations)(nil)),

	438: reflect.TypeOf((*PrismSubareaEmptyInfo)(nil)),

	434: reflect.TypeOf((*PrismGeolocalizedInformation)(nil)),

	420: reflect.TypeOf((*GuildInAllianceInformations)(nil)),

	494: reflect.TypeOf((*ObjectItemGenericQuantityPrice)(nil)),

	37: reflect.TypeOf((*ObjectItem)(nil)),

	419: reflect.TypeOf((*BasicAllianceInformations)(nil)),

	217: reflect.TypeOf((*FightEntityDispositionInformations)(nil)),

	108: reflect.TypeOf((*StatedElement)(nil)),

	360: reflect.TypeOf((*GameFightMinimalStatsPreparation)(nil)),

	200: reflect.TypeOf((*MapObstacle)(nil)),

	107: reflect.TypeOf((*IdentifiedEntityDispositionInformations)(nil)),

	198: reflect.TypeOf((*ObjectItemInRolePlay)(nil)),

	185: reflect.TypeOf((*PaddockItem)(nil)),

	425: reflect.TypeOf((*HumanOptionAlliance)(nil)),

	405: reflect.TypeOf((*IndexedEntityLook)(nil)),

	495: reflect.TypeOf((*HumanOptionSkillUse)(nil)),

	407: reflect.TypeOf((*HumanOptionEmote)(nil)),

	410: reflect.TypeOf((*HumanOptionFollowers)(nil)),

	353: reflect.TypeOf((*ActorOrientation)(nil)),

	449: reflect.TypeOf((*HumanOptionObjectUse)(nil)),

	377: reflect.TypeOf((*TrustCertificate)(nil)),

	206: reflect.TypeOf((*AbstractFightDispellableEffect)(nil)),

	209: reflect.TypeOf((*FightTemporaryBoostEffect)(nil)),

	205: reflect.TypeOf((*GameFightSpellCooldown)(nil)),

	85: reflect.TypeOf((*GameActionMarkedCell)(nil)),

	16: reflect.TypeOf((*FightResultListEntry)(nil)),

	189: reflect.TypeOf((*FightResultFighterListEntry)(nil)),

	208: reflect.TypeOf((*FightDispellableEffectExtendedInformations)(nil)),

	24: reflect.TypeOf((*FightResultPlayerListEntry)(nil)),

	469: reflect.TypeOf((*NamedPartyTeam)(nil)),

	489: reflect.TypeOf((*Idol)(nil)),

	351: reflect.TypeOf((*GameActionMark)(nil)),

	470: reflect.TypeOf((*NamedPartyTeamWithOutcome)(nil)),

	50: reflect.TypeOf((*GameFightMutantInformations)(nil)),

	364: reflect.TypeOf((*GameFightResumeSlaveInfo)(nil)),

	84: reflect.TypeOf((*FightResultTaxCollectorListEntry)(nil)),

	222: reflect.TypeOf((*PaddockInformationsForSell)(nil)),

	221: reflect.TypeOf((*HouseInformationsForSell)(nil)),

	211: reflect.TypeOf((*FightTemporaryBoostWeaponDamagesEffect)(nil)),

	207: reflect.TypeOf((*FightTemporarySpellBoostEffect)(nil)),

	214: reflect.TypeOf((*FightTemporaryBoostStateEffect)(nil)),

	210: reflect.TypeOf((*FightTriggeredEffect)(nil)),

	366: reflect.TypeOf((*FightTemporarySpellImmunityEffect)(nil)),

	202: reflect.TypeOf((*ActorExtendedAlignmentInformations)(nil)),

	443: reflect.TypeOf((*PrismFightersInformation)(nil)),

	169: reflect.TypeOf((*TaxCollectorFightersInformation)(nil)),

	124: reflect.TypeOf((*ObjectItemMinimalInformation)(nil)),

	387: reflect.TypeOf((*ObjectItemInformationWithQuantity)(nil)),

	52: reflect.TypeOf((*StartupActionAddObject)(nil)),

	502: reflect.TypeOf((*DareInformations)(nil)),

	504: reflect.TypeOf((*DareVersatileInformations)(nil)),

	501: reflect.TypeOf((*DareCriteria)(nil)),

	87: reflect.TypeOf((*GuildEmblem)(nil)),

	422: reflect.TypeOf((*AlliancedGuildFactSheetInformations)(nil)),

	435: reflect.TypeOf((*GuildVersatileInformations)(nil)),

	432: reflect.TypeOf((*AllianceVersatileInformations)(nil)),

	418: reflect.TypeOf((*BasicNamedAllianceInformations)(nil)),

	417: reflect.TypeOf((*AllianceInformations)(nil)),

	421: reflect.TypeOf((*AllianceFactSheetInformations)(nil)),

	448: reflect.TypeOf((*TaxCollectorComplementaryInformations)(nil)),

	446: reflect.TypeOf((*TaxCollectorGuildInformations)(nil)),

	165: reflect.TypeOf((*AdditionalTaxCollectorInformations)(nil)),

	372: reflect.TypeOf((*TaxCollectorLootInformations)(nil)),

	447: reflect.TypeOf((*TaxCollectorWaitingForHelpInformations)(nil)),

	77: reflect.TypeOf((*FriendSpouseInformations)(nil)),

	93: reflect.TypeOf((*FriendSpouseOnlineInformations)(nil)),

	119: reflect.TypeOf((*ObjectItemQuantity)(nil)),

	369: reflect.TypeOf((*Shortcut)(nil)),

	367: reflect.TypeOf((*ShortcutObject)(nil)),

	492: reflect.TypeOf((*ShortcutObjectIdolsPreset)(nil)),

	355: reflect.TypeOf((*Preset)(nil)),

	388: reflect.TypeOf((*ShortcutSmiley)(nil)),

	491: reflect.TypeOf((*IdolsPreset)(nil)),

	389: reflect.TypeOf((*ShortcutEmote)(nil)),

	370: reflect.TypeOf((*ShortcutObjectPreset)(nil)),

	371: reflect.TypeOf((*ShortcutObjectItem)(nil)),

	368: reflect.TypeOf((*ShortcutSpell)(nil)),

	178: reflect.TypeOf((*MountClientData)(nil)),

	356: reflect.TypeOf((*UpdateMountBoost)(nil)),

	357: reflect.TypeOf((*UpdateMountIntBoost)(nil)),

	484: reflect.TypeOf((*StatisticData)(nil)),

	485: reflect.TypeOf((*StatisticDataInt)(nil)),

	445: reflect.TypeOf((*CharacterMinimalGuildInformations)(nil)),

	486: reflect.TypeOf((*StatisticDataByte)(nil)),

	385: reflect.TypeOf((*QuestObjectiveInformations)(nil)),

	177: reflect.TypeOf((*FightTeamMemberTaxCollectorInformations)(nil)),

	430: reflect.TypeOf((*ServerSessionConstant)(nil)),

	436: reflect.TypeOf((*ServerSessionConstantString)(nil)),

	174: reflect.TypeOf((*MapCoordinates)(nil)),

	392: reflect.TypeOf((*MapCoordinatesAndId)(nil)),

	176: reflect.TypeOf((*MapCoordinatesExtended)(nil)),

	463: reflect.TypeOf((*TreasureHuntStep)(nil)),

	462: reflect.TypeOf((*TreasureHuntStepFight)(nil)),

	382: reflect.TypeOf((*QuestActiveDetailedInformations)(nil)),

	71: reflect.TypeOf((*ObjectEffectCreature)(nil)),

	74: reflect.TypeOf((*ObjectEffectString)(nil)),

	482: reflect.TypeOf((*StatisticDataBoolean)(nil)),

	440: reflect.TypeOf((*TaxCollectorStaticExtendedInformations)(nil)),

	191: reflect.TypeOf((*FightResultAdditionalData)(nil)),

	190: reflect.TypeOf((*FightResultPvpData)(nil)),

	444: reflect.TypeOf((*CharacterMinimalAllianceInformations)(nil)),

	6: reflect.TypeOf((*FightTeamMemberMonsterInformations)(nil)),

	82: reflect.TypeOf((*ObjectEffectMinMax)(nil)),

	439: reflect.TypeOf((*FightAllianceTeamInformations)(nil)),

	490: reflect.TypeOf((*PartyIdol)(nil)),

	203: reflect.TypeOf((*GameFightMonsterWithAlignmentInformations)(nil)),

	398: reflect.TypeOf((*InteractiveElementWithAgeBonus)(nil)),

	413: reflect.TypeOf((*GameFightFighterLightInformations)(nil)),

	13: reflect.TypeOf((*FightTeamMemberCharacterInformations)(nil)),

	179: reflect.TypeOf((*ObjectEffectMount)(nil)),

	487: reflect.TypeOf((*StatisticDataString)(nil)),

	429: reflect.TypeOf((*ServerSessionConstantLong)(nil)),

	457: reflect.TypeOf((*GameFightFighterTaxCollectorLightInformations)(nil)),

	426: reflect.TypeOf((*FightTeamMemberWithAllianceCharacterInformations)(nil)),

	474: reflect.TypeOf((*CharacterHardcoreOrEpicInformations)(nil)),

	471: reflect.TypeOf((*GameRolePlayTreasureHintInformations)(nil)),

	75: reflect.TypeOf((*ObjectEffectDuration)(nil)),

	455: reflect.TypeOf((*GameFightFighterMonsterLightInformations)(nil)),

	409: reflect.TypeOf((*HumanOptionGuild)(nil)),

	512: reflect.TypeOf((*HouseGuildedInformations)(nil)),

	72: reflect.TypeOf((*ObjectEffectDate)(nil)),

	81: reflect.TypeOf((*ObjectEffectLadder)(nil)),

	454: reflect.TypeOf((*GameFightFighterCompanionLightInformations)(nil)),

	73: reflect.TypeOf((*ObjectEffectDice)(nil)),

	465: reflect.TypeOf((*TreasureHuntStepDig)(nil)),

	468: reflect.TypeOf((*TreasureHuntStepFollowDirection)(nil)),

	508: reflect.TypeOf((*PaddockGuildedInformations)(nil)),

	437: reflect.TypeOf((*GuildInAllianceVersatileInformations)(nil)),

	456: reflect.TypeOf((*GameFightFighterNamedLightInformations)(nil)),

	192: reflect.TypeOf((*FightResultExperienceData)(nil)),

	451: reflect.TypeOf((*FightTeamMemberCompanionInformations)(nil)),

	220: reflect.TypeOf((*InteractiveElementNamedSkill)(nil)),

	488: reflect.TypeOf((*StatisticDataShort)(nil)),

	216: reflect.TypeOf((*FightResultMutantListEntry)(nil)),

	433: reflect.TypeOf((*ServerSessionConstantInteger)(nil)),

	461: reflect.TypeOf((*TreasureHuntStepFollowDirectionToPOI)(nil)),

	193: reflect.TypeOf((*CharacterMinimalPlusLookAndGradeInformations)(nil)),

	472: reflect.TypeOf((*TreasureHuntStepFollowDirectionToHint)(nil)),

	386: reflect.TypeOf((*QuestObjectiveInformationsWithCompletion)(nil)),

	473: reflect.TypeOf((*TreasureHuntFlag)(nil)),

	98: reflect.TypeOf((*JobExperience)(nil)),

	500: reflect.TypeOf((*JobBookSubscription)(nil)),

	97: reflect.TypeOf((*JobCrafterDirectorySettings)(nil)),

	354: reflect.TypeOf((*PresetItem)(nil)),

	479: reflect.TypeOf((*CharacterRemodelingInformation)(nil)),

	477: reflect.TypeOf((*CharacterToRemodelInformations)(nil)),

	480: reflect.TypeOf((*RemodelingInformation)(nil)),

	11: reflect.TypeOf((*Version)(nil)),

	393: reflect.TypeOf((*VersionExtended)(nil)),

	497: reflect.TypeOf((*MonsterBoosts)(nil)),

	506: reflect.TypeOf((*FinishMoveInformations)(nil)),

	49: reflect.TypeOf((*SpellItem)(nil)),

	499: reflect.TypeOf((*ArenaRankInfos)(nil)),

	378: reflect.TypeOf((*PartyMemberGeoPosition)(nil)),

	20: reflect.TypeOf((*FightOptionsInformations)(nil)),

	513: reflect.TypeOf((*FightStartingPositions)(nil)),

	96: reflect.TypeOf((*TaxCollectorBasicInformations)(nil)),

	184: reflect.TypeOf((*MountInformationsForPaddock)(nil)),

	41: reflect.TypeOf((*FightLoot)(nil)),

	120: reflect.TypeOf((*ObjectItemToSell)(nil)),

	359: reflect.TypeOf((*ObjectItemToSellInHumanVendorShop)(nil)),

	164: reflect.TypeOf((*ObjectItemToSellInBid)(nil)),

	122: reflect.TypeOf((*BidExchangerObjectInfo)(nil)),

	121: reflect.TypeOf((*SellerBuyerDescriptor)(nil)),

	117: reflect.TypeOf((*FightExternalInformations)(nil)),

	352: reflect.TypeOf((*ObjectItemToSellInNpcShop)(nil)),

	186: reflect.TypeOf((*ProtectedEntityWaitingForHelpInfo)(nil)),

	481: reflect.TypeOf((*DecraftedItemStackInfo)(nil)),

	363: reflect.TypeOf((*Achievement)(nil)),

	196: reflect.TypeOf((*JobCrafterDirectoryListEntry)(nil)),

	498: reflect.TypeOf((*HavenBagFurnitureInformation)(nil)),

	397: reflect.TypeOf((*KrosmasterFigure)(nil)),

	63: reflect.TypeOf((*EntityMovementInformations)(nil)),

	134: reflect.TypeOf((*ObjectItemNotInContainer)(nil)),

	123: reflect.TypeOf((*GoldItem)(nil)),

	475: reflect.TypeOf((*AbstractCharacterToRefurbishInformation)(nil)),

	212: reflect.TypeOf((*CharacterToRecolorInformation)(nil)),

	399: reflect.TypeOf((*CharacterToRelookInformation)(nil)),

	195: reflect.TypeOf((*JobCrafterDirectoryEntryJobInfo)(nil)),

	194: reflect.TypeOf((*JobCrafterDirectoryEntryPlayerInfo)(nil)),

	175: reflect.TypeOf((*AtlasPointsInformations)(nil)),

	115: reflect.TypeOf((*FightTeamLightInformations)(nil)),

	404: reflect.TypeOf((*AchievementObjective)(nil)),

	402: reflect.TypeOf((*AchievementStartedObjective)(nil)),
}

type ContentPart struct {
	Id string

	State uint8
}

func (m *ContentPart) ID() uint16 {
	return 350
}

func (m *ContentPart) Serialize(w Writer) error {

	if err := w.WriteString(m.Id); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.State); err != nil {
		return err
	}

	return nil
}

func (m *ContentPart) Deserialize(r Reader) error {

	lid, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Id = lid

	lstate, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.State = lstate

	return nil
}

type GameContextActorInformations struct {
	ContextualId float64

	Look *EntityLook

	Disposition *EntityDispositionInformations
}

func (m *GameContextActorInformations) ID() uint16 {
	return 150
}

func (m *GameContextActorInformations) Serialize(w Writer) error {

	if err := w.WriteDouble(m.ContextualId); err != nil {
		return err
	}

	if err := m.Look.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Disposition.ID()); err != nil {
		return err
	}

	if err := m.Disposition.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameContextActorInformations) Deserialize(r Reader) error {

	lcontextualId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.ContextualId = lcontextualId

	var llook EntityLook

	llook.Deserialize(r)

	m.Look = &llook

	typedispositionID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	ldisposition, err := GetType(typedispositionID)
	if err != nil {
		return err
	}

	ldisposition.Deserialize(r)

	m.Disposition = ldisposition.(*EntityDispositionInformations)

	return nil
}

type GameFightFighterInformations struct {
	GameContextActorInformations

	TeamId uint8

	Wave uint8

	Alive bool

	Stats *GameFightMinimalStats

	PreviousPositions []uint16
}

func (m *GameFightFighterInformations) ID() uint16 {
	return 143
}

func (m *GameFightFighterInformations) Serialize(w Writer) error {

	if err := m.GameContextActorInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.TeamId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Wave); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Alive); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Stats.ID()); err != nil {
		return err
	}

	if err := m.Stats.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.PreviousPositions))); err != nil {
		return err
	}

	for i := range m.PreviousPositions {

		if err := w.WriteVarUInt16(m.PreviousPositions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameFightFighterInformations) Deserialize(r Reader) error {

	if err := m.GameContextActorInformations.Deserialize(r); err != nil {
		return err
	}

	lteamId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TeamId = lteamId

	lwave, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Wave = lwave

	lalive, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Alive = lalive

	typestatsID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lstats, err := GetType(typestatsID)
	if err != nil {
		return err
	}

	lstats.Deserialize(r)

	m.Stats = lstats.(*GameFightMinimalStats)

	lpreviousPositionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PreviousPositions = make([]uint16, lpreviousPositionsLen)

	for i := range m.PreviousPositions {

		lpreviousPositions, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.PreviousPositions[i] = lpreviousPositions

	}

	return nil
}

type GameFightAIInformations struct {
	GameFightFighterInformations
}

func (m *GameFightAIInformations) ID() uint16 {
	return 151
}

func (m *GameFightAIInformations) Serialize(w Writer) error {

	if err := m.GameFightFighterInformations.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameFightAIInformations) Deserialize(r Reader) error {

	if err := m.GameFightFighterInformations.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type GameFightMonsterInformations struct {
	GameFightAIInformations

	CreatureGenericId uint16

	CreatureGrade uint8
}

func (m *GameFightMonsterInformations) ID() uint16 {
	return 29
}

func (m *GameFightMonsterInformations) Serialize(w Writer) error {

	if err := m.GameFightAIInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.CreatureGenericId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.CreatureGrade); err != nil {
		return err
	}

	return nil
}

func (m *GameFightMonsterInformations) Deserialize(r Reader) error {

	if err := m.GameFightAIInformations.Deserialize(r); err != nil {
		return err
	}

	lcreatureGenericId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CreatureGenericId = lcreatureGenericId

	lcreatureGrade, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CreatureGrade = lcreatureGrade

	return nil
}

type GameRolePlayActorInformations struct {
	GameContextActorInformations
}

func (m *GameRolePlayActorInformations) ID() uint16 {
	return 141
}

func (m *GameRolePlayActorInformations) Serialize(w Writer) error {

	if err := m.GameContextActorInformations.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayActorInformations) Deserialize(r Reader) error {

	if err := m.GameContextActorInformations.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type GameRolePlayNamedActorInformations struct {
	GameRolePlayActorInformations

	Name string
}

func (m *GameRolePlayNamedActorInformations) ID() uint16 {
	return 154
}

func (m *GameRolePlayNamedActorInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayActorInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayNamedActorInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayActorInformations.Deserialize(r); err != nil {
		return err
	}

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	return nil
}

type GameRolePlayMountInformations struct {
	GameRolePlayNamedActorInformations

	OwnerName string

	Level uint8
}

func (m *GameRolePlayMountInformations) ID() uint16 {
	return 180
}

func (m *GameRolePlayMountInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayNamedActorInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.OwnerName); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayMountInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayNamedActorInformations.Deserialize(r); err != nil {
		return err
	}

	lownerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.OwnerName = lownerName

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	return nil
}

type InteractiveElement struct {
	ElementId uint32

	ElementTypeId int32

	EnabledSkills []*InteractiveElementSkill

	DisabledSkills []*InteractiveElementSkill

	OnCurrentMap bool
}

func (m *InteractiveElement) ID() uint16 {
	return 80
}

func (m *InteractiveElement) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.ElementId); err != nil {
		return err
	}

	if err := w.WriteInt32(m.ElementTypeId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.EnabledSkills))); err != nil {
		return err
	}

	for i := range m.EnabledSkills {

		if err := w.WriteUInt16(m.EnabledSkills[i].ID()); err != nil {
			return err
		}

		if err := m.EnabledSkills[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.DisabledSkills))); err != nil {
		return err
	}

	for i := range m.DisabledSkills {

		if err := w.WriteUInt16(m.DisabledSkills[i].ID()); err != nil {
			return err
		}

		if err := m.DisabledSkills[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteBoolean(m.OnCurrentMap); err != nil {
		return err
	}

	return nil
}

func (m *InteractiveElement) Deserialize(r Reader) error {

	lelementId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.ElementId = lelementId

	lelementTypeId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.ElementTypeId = lelementTypeId

	lenabledSkillsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.EnabledSkills = make([]*InteractiveElementSkill, lenabledSkillsLen)

	for i := range m.EnabledSkills {

		typeenabledSkillsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lenabledSkills, err := GetType(typeenabledSkillsID)
		if err != nil {
			return err
		}

		lenabledSkills.Deserialize(r)

		m.EnabledSkills[i] = lenabledSkills.(*InteractiveElementSkill)

	}

	ldisabledSkillsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DisabledSkills = make([]*InteractiveElementSkill, ldisabledSkillsLen)

	for i := range m.DisabledSkills {

		typedisabledSkillsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		ldisabledSkills, err := GetType(typedisabledSkillsID)
		if err != nil {
			return err
		}

		ldisabledSkills.Deserialize(r)

		m.DisabledSkills[i] = ldisabledSkills.(*InteractiveElementSkill)

	}

	lonCurrentMap, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.OnCurrentMap = lonCurrentMap

	return nil
}

type GameRolePlayPortalInformations struct {
	GameRolePlayActorInformations

	Portal *PortalInformation
}

func (m *GameRolePlayPortalInformations) ID() uint16 {
	return 467
}

func (m *GameRolePlayPortalInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayActorInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Portal.ID()); err != nil {
		return err
	}

	if err := m.Portal.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayPortalInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayActorInformations.Deserialize(r); err != nil {
		return err
	}

	typeportalID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lportal, err := GetType(typeportalID)
	if err != nil {
		return err
	}

	lportal.Deserialize(r)

	m.Portal = lportal.(*PortalInformation)

	return nil
}

type GameRolePlayGroupMonsterInformations struct {
	GameRolePlayActorInformations

	StaticInfos *GroupMonsterStaticInformations

	CreationTime float64

	AgeBonusRate uint32

	LootShare int8

	AlignmentSide int8

	KeyRingBonus bool

	HasHardcoreDrop bool

	HasAVARewardToken bool
}

func (m *GameRolePlayGroupMonsterInformations) ID() uint16 {
	return 160
}

func (m *GameRolePlayGroupMonsterInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayActorInformations.Serialize(w); err != nil {
		return err
	}

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.KeyRingBonus)

	setWrappedFlag(bbw0, 1, m.HasHardcoreDrop)

	setWrappedFlag(bbw0, 2, m.HasAVARewardToken)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.StaticInfos.ID()); err != nil {
		return err
	}

	if err := m.StaticInfos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.CreationTime); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.AgeBonusRate); err != nil {
		return err
	}

	if err := w.WriteInt8(m.LootShare); err != nil {
		return err
	}

	if err := w.WriteInt8(m.AlignmentSide); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayGroupMonsterInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayActorInformations.Deserialize(r); err != nil {
		return err
	}

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.KeyRingBonus = getWrappedFlag(bbw0, 0)

	m.HasHardcoreDrop = getWrappedFlag(bbw0, 1)

	m.HasAVARewardToken = getWrappedFlag(bbw0, 2)

	typestaticInfosID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lstaticInfos, err := GetType(typestaticInfosID)
	if err != nil {
		return err
	}

	lstaticInfos.Deserialize(r)

	m.StaticInfos = lstaticInfos.(*GroupMonsterStaticInformations)

	lcreationTime, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.CreationTime = lcreationTime

	lageBonusRate, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AgeBonusRate = lageBonusRate

	llootShare, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.LootShare = llootShare

	lalignmentSide, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.AlignmentSide = lalignmentSide

	return nil
}

type GameRolePlayGroupMonsterWaveInformations struct {
	GameRolePlayGroupMonsterInformations

	NbWaves uint8

	Alternatives []*GroupMonsterStaticInformations
}

func (m *GameRolePlayGroupMonsterWaveInformations) ID() uint16 {
	return 464
}

func (m *GameRolePlayGroupMonsterWaveInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayGroupMonsterInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.NbWaves); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Alternatives))); err != nil {
		return err
	}

	for i := range m.Alternatives {

		if err := w.WriteUInt16(m.Alternatives[i].ID()); err != nil {
			return err
		}

		if err := m.Alternatives[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameRolePlayGroupMonsterWaveInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayGroupMonsterInformations.Deserialize(r); err != nil {
		return err
	}

	lnbWaves, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NbWaves = lnbWaves

	lalternativesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Alternatives = make([]*GroupMonsterStaticInformations, lalternativesLen)

	for i := range m.Alternatives {

		typealternativesID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lalternatives, err := GetType(typealternativesID)
		if err != nil {
			return err
		}

		lalternatives.Deserialize(r)

		m.Alternatives[i] = lalternatives.(*GroupMonsterStaticInformations)

	}

	return nil
}

type GameRolePlayPrismInformations struct {
	GameRolePlayActorInformations

	Prism *PrismInformation
}

func (m *GameRolePlayPrismInformations) ID() uint16 {
	return 161
}

func (m *GameRolePlayPrismInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayActorInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Prism.ID()); err != nil {
		return err
	}

	if err := m.Prism.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayPrismInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayActorInformations.Deserialize(r); err != nil {
		return err
	}

	typeprismID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lprism, err := GetType(typeprismID)
	if err != nil {
		return err
	}

	lprism.Deserialize(r)

	m.Prism = lprism.(*PrismInformation)

	return nil
}

type GameFightTaxCollectorInformations struct {
	GameFightAIInformations

	FirstNameId uint16

	LastNameId uint16

	Level uint8
}

func (m *GameFightTaxCollectorInformations) ID() uint16 {
	return 48
}

func (m *GameFightTaxCollectorInformations) Serialize(w Writer) error {

	if err := m.GameFightAIInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.FirstNameId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.LastNameId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	return nil
}

func (m *GameFightTaxCollectorInformations) Deserialize(r Reader) error {

	if err := m.GameFightAIInformations.Deserialize(r); err != nil {
		return err
	}

	lfirstNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.FirstNameId = lfirstNameId

	llastNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.LastNameId = llastNameId

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	return nil
}

type GameRolePlayNpcInformations struct {
	GameRolePlayActorInformations

	NpcId uint16

	Sex bool

	SpecialArtworkId uint16
}

func (m *GameRolePlayNpcInformations) ID() uint16 {
	return 156
}

func (m *GameRolePlayNpcInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayActorInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NpcId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Sex); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SpecialArtworkId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayNpcInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayActorInformations.Deserialize(r); err != nil {
		return err
	}

	lnpcId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NpcId = lnpcId

	lsex, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Sex = lsex

	lspecialArtworkId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpecialArtworkId = lspecialArtworkId

	return nil
}

type GameRolePlayTaxCollectorInformations struct {
	GameRolePlayActorInformations

	Identification *TaxCollectorStaticInformations

	GuildLevel uint8

	TaxCollectorAttack int32
}

func (m *GameRolePlayTaxCollectorInformations) ID() uint16 {
	return 148
}

func (m *GameRolePlayTaxCollectorInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayActorInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Identification.ID()); err != nil {
		return err
	}

	if err := m.Identification.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.GuildLevel); err != nil {
		return err
	}

	if err := w.WriteInt32(m.TaxCollectorAttack); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayTaxCollectorInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayActorInformations.Deserialize(r); err != nil {
		return err
	}

	typeidentificationID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lidentification, err := GetType(typeidentificationID)
	if err != nil {
		return err
	}

	lidentification.Deserialize(r)

	m.Identification = lidentification.(*TaxCollectorStaticInformations)

	lguildLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.GuildLevel = lguildLevel

	ltaxCollectorAttack, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.TaxCollectorAttack = ltaxCollectorAttack

	return nil
}

type GameFightFighterNamedInformations struct {
	GameFightFighterInformations

	Name string

	Status *PlayerStatus
}

func (m *GameFightFighterNamedInformations) ID() uint16 {
	return 158
}

func (m *GameFightFighterNamedInformations) Serialize(w Writer) error {

	if err := m.GameFightFighterInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := m.Status.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameFightFighterNamedInformations) Deserialize(r Reader) error {

	if err := m.GameFightFighterInformations.Deserialize(r); err != nil {
		return err
	}

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	var lstatus PlayerStatus

	lstatus.Deserialize(r)

	m.Status = &lstatus

	return nil
}

type GameFightCharacterInformations struct {
	GameFightFighterNamedInformations

	Level uint8

	AlignmentInfos *ActorAlignmentInformations

	Breed int8

	Sex bool
}

func (m *GameFightCharacterInformations) ID() uint16 {
	return 46
}

func (m *GameFightCharacterInformations) Serialize(w Writer) error {

	if err := m.GameFightFighterNamedInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	if err := m.AlignmentInfos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Sex); err != nil {
		return err
	}

	return nil
}

func (m *GameFightCharacterInformations) Deserialize(r Reader) error {

	if err := m.GameFightFighterNamedInformations.Deserialize(r); err != nil {
		return err
	}

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	var lalignmentInfos ActorAlignmentInformations

	lalignmentInfos.Deserialize(r)

	m.AlignmentInfos = &lalignmentInfos

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	lsex, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Sex = lsex

	return nil
}

type GameFightCompanionInformations struct {
	GameFightFighterInformations

	CompanionGenericId uint8

	Level uint8

	MasterId float64
}

func (m *GameFightCompanionInformations) ID() uint16 {
	return 450
}

func (m *GameFightCompanionInformations) Serialize(w Writer) error {

	if err := m.GameFightFighterInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.CompanionGenericId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	if err := w.WriteDouble(m.MasterId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightCompanionInformations) Deserialize(r Reader) error {

	if err := m.GameFightFighterInformations.Deserialize(r); err != nil {
		return err
	}

	lcompanionGenericId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CompanionGenericId = lcompanionGenericId

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	lmasterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.MasterId = lmasterId

	return nil
}

type GameRolePlayHumanoidInformations struct {
	GameRolePlayNamedActorInformations

	HumanoidInfo *HumanInformations

	AccountId uint32
}

func (m *GameRolePlayHumanoidInformations) ID() uint16 {
	return 159
}

func (m *GameRolePlayHumanoidInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayNamedActorInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.HumanoidInfo.ID()); err != nil {
		return err
	}

	if err := m.HumanoidInfo.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayHumanoidInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayNamedActorInformations.Deserialize(r); err != nil {
		return err
	}

	typehumanoidInfoID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lhumanoidInfo, err := GetType(typehumanoidInfoID)
	if err != nil {
		return err
	}

	lhumanoidInfo.Deserialize(r)

	m.HumanoidInfo = lhumanoidInfo.(*HumanInformations)

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	return nil
}

type GameRolePlayCharacterInformations struct {
	GameRolePlayHumanoidInformations

	AlignmentInfos *ActorAlignmentInformations
}

func (m *GameRolePlayCharacterInformations) ID() uint16 {
	return 36
}

func (m *GameRolePlayCharacterInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayHumanoidInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.AlignmentInfos.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayCharacterInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayHumanoidInformations.Deserialize(r); err != nil {
		return err
	}

	var lalignmentInfos ActorAlignmentInformations

	lalignmentInfos.Deserialize(r)

	m.AlignmentInfos = &lalignmentInfos

	return nil
}

type GameRolePlayNpcWithQuestInformations struct {
	GameRolePlayNpcInformations

	QuestFlag *GameRolePlayNpcQuestFlag
}

func (m *GameRolePlayNpcWithQuestInformations) ID() uint16 {
	return 383
}

func (m *GameRolePlayNpcWithQuestInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayNpcInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.QuestFlag.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayNpcWithQuestInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayNpcInformations.Deserialize(r); err != nil {
		return err
	}

	var lquestFlag GameRolePlayNpcQuestFlag

	lquestFlag.Deserialize(r)

	m.QuestFlag = &lquestFlag

	return nil
}

type GameRolePlayMerchantInformations struct {
	GameRolePlayNamedActorInformations

	SellType uint8

	Options []*HumanOption
}

func (m *GameRolePlayMerchantInformations) ID() uint16 {
	return 129
}

func (m *GameRolePlayMerchantInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayNamedActorInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.SellType); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Options))); err != nil {
		return err
	}

	for i := range m.Options {

		if err := w.WriteUInt16(m.Options[i].ID()); err != nil {
			return err
		}

		if err := m.Options[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameRolePlayMerchantInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayNamedActorInformations.Deserialize(r); err != nil {
		return err
	}

	lsellType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.SellType = lsellType

	loptionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Options = make([]*HumanOption, loptionsLen)

	for i := range m.Options {

		typeoptionsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		loptions, err := GetType(typeoptionsID)
		if err != nil {
			return err
		}

		loptions.Deserialize(r)

		m.Options[i] = loptions.(*HumanOption)

	}

	return nil
}

type GameRolePlayMutantInformations struct {
	GameRolePlayHumanoidInformations

	MonsterId uint16

	PowerLevel int8
}

func (m *GameRolePlayMutantInformations) ID() uint16 {
	return 3
}

func (m *GameRolePlayMutantInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayHumanoidInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MonsterId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.PowerLevel); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayMutantInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayHumanoidInformations.Deserialize(r); err != nil {
		return err
	}

	lmonsterId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MonsterId = lmonsterId

	lpowerLevel, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.PowerLevel = lpowerLevel

	return nil
}

type SubEntity struct {
	BindingPointCategory uint8

	BindingPointIndex uint8

	SubEntityLook *EntityLook
}

func (m *SubEntity) ID() uint16 {
	return 54
}

func (m *SubEntity) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.BindingPointCategory); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.BindingPointIndex); err != nil {
		return err
	}

	if err := m.SubEntityLook.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *SubEntity) Deserialize(r Reader) error {

	lbindingPointCategory, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BindingPointCategory = lbindingPointCategory

	lbindingPointIndex, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BindingPointIndex = lbindingPointIndex

	var lsubEntityLook EntityLook

	lsubEntityLook.Deserialize(r)

	m.SubEntityLook = &lsubEntityLook

	return nil
}

type EntityLook struct {
	BonesId uint16

	Skins []uint16

	IndexedColors []int32

	Scales []int16

	Subentities []*SubEntity
}

func (m *EntityLook) ID() uint16 {
	return 55
}

func (m *EntityLook) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.BonesId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Skins))); err != nil {
		return err
	}

	for i := range m.Skins {

		if err := w.WriteVarUInt16(m.Skins[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.IndexedColors))); err != nil {
		return err
	}

	for i := range m.IndexedColors {

		if err := w.WriteInt32(m.IndexedColors[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Scales))); err != nil {
		return err
	}

	for i := range m.Scales {

		if err := w.WriteVarInt16(m.Scales[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Subentities))); err != nil {
		return err
	}

	for i := range m.Subentities {

		if err := m.Subentities[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *EntityLook) Deserialize(r Reader) error {

	lbonesId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.BonesId = lbonesId

	lskinsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Skins = make([]uint16, lskinsLen)

	for i := range m.Skins {

		lskins, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Skins[i] = lskins

	}

	lindexedColorsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.IndexedColors = make([]int32, lindexedColorsLen)

	for i := range m.IndexedColors {

		lindexedColors, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.IndexedColors[i] = lindexedColors

	}

	lscalesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Scales = make([]int16, lscalesLen)

	for i := range m.Scales {

		lscales, err := r.ReadVarInt16()
		if err != nil {
			return err
		}

		m.Scales[i] = lscales

	}

	lsubentitiesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Subentities = make([]*SubEntity, lsubentitiesLen)

	for i := range m.Subentities {

		var lsubentities SubEntity

		lsubentities.Deserialize(r)

		m.Subentities[i] = &lsubentities

	}

	return nil
}

type CharacterCharacteristicsInformations struct {
	Experience int64

	ExperienceLevelFloor int64

	ExperienceNextLevelFloor int64

	ExperienceBonusLimit int64

	Kamas int64

	StatsPoints uint16

	AdditionnalPoints uint16

	SpellsPoints uint16

	AlignmentInfos *ActorExtendedAlignmentInformations

	LifePoints uint32

	MaxLifePoints uint32

	EnergyPoints uint16

	MaxEnergyPoints uint16

	ActionPointsCurrent int16

	MovementPointsCurrent int16

	Initiative *CharacterBaseCharacteristic

	Prospecting *CharacterBaseCharacteristic

	ActionPoints *CharacterBaseCharacteristic

	MovementPoints *CharacterBaseCharacteristic

	Strength *CharacterBaseCharacteristic

	Vitality *CharacterBaseCharacteristic

	Wisdom *CharacterBaseCharacteristic

	Chance *CharacterBaseCharacteristic

	Agility *CharacterBaseCharacteristic

	Intelligence *CharacterBaseCharacteristic

	Range *CharacterBaseCharacteristic

	SummonableCreaturesBoost *CharacterBaseCharacteristic

	Reflect *CharacterBaseCharacteristic

	CriticalHit *CharacterBaseCharacteristic

	CriticalHitWeapon uint16

	CriticalMiss *CharacterBaseCharacteristic

	HealBonus *CharacterBaseCharacteristic

	AllDamagesBonus *CharacterBaseCharacteristic

	WeaponDamagesBonusPercent *CharacterBaseCharacteristic

	DamagesBonusPercent *CharacterBaseCharacteristic

	TrapBonus *CharacterBaseCharacteristic

	TrapBonusPercent *CharacterBaseCharacteristic

	GlyphBonusPercent *CharacterBaseCharacteristic

	RuneBonusPercent *CharacterBaseCharacteristic

	PermanentDamagePercent *CharacterBaseCharacteristic

	TackleBlock *CharacterBaseCharacteristic

	TackleEvade *CharacterBaseCharacteristic

	PAAttack *CharacterBaseCharacteristic

	PMAttack *CharacterBaseCharacteristic

	PushDamageBonus *CharacterBaseCharacteristic

	CriticalDamageBonus *CharacterBaseCharacteristic

	NeutralDamageBonus *CharacterBaseCharacteristic

	EarthDamageBonus *CharacterBaseCharacteristic

	WaterDamageBonus *CharacterBaseCharacteristic

	AirDamageBonus *CharacterBaseCharacteristic

	FireDamageBonus *CharacterBaseCharacteristic

	DodgePALostProbability *CharacterBaseCharacteristic

	DodgePMLostProbability *CharacterBaseCharacteristic

	NeutralElementResistPercent *CharacterBaseCharacteristic

	EarthElementResistPercent *CharacterBaseCharacteristic

	WaterElementResistPercent *CharacterBaseCharacteristic

	AirElementResistPercent *CharacterBaseCharacteristic

	FireElementResistPercent *CharacterBaseCharacteristic

	NeutralElementReduction *CharacterBaseCharacteristic

	EarthElementReduction *CharacterBaseCharacteristic

	WaterElementReduction *CharacterBaseCharacteristic

	AirElementReduction *CharacterBaseCharacteristic

	FireElementReduction *CharacterBaseCharacteristic

	PushDamageReduction *CharacterBaseCharacteristic

	CriticalDamageReduction *CharacterBaseCharacteristic

	PvpNeutralElementResistPercent *CharacterBaseCharacteristic

	PvpEarthElementResistPercent *CharacterBaseCharacteristic

	PvpWaterElementResistPercent *CharacterBaseCharacteristic

	PvpAirElementResistPercent *CharacterBaseCharacteristic

	PvpFireElementResistPercent *CharacterBaseCharacteristic

	PvpNeutralElementReduction *CharacterBaseCharacteristic

	PvpEarthElementReduction *CharacterBaseCharacteristic

	PvpWaterElementReduction *CharacterBaseCharacteristic

	PvpAirElementReduction *CharacterBaseCharacteristic

	PvpFireElementReduction *CharacterBaseCharacteristic

	SpellModifications []*CharacterSpellModification

	ProbationTime uint32
}

func (m *CharacterCharacteristicsInformations) ID() uint16 {
	return 8
}

func (m *CharacterCharacteristicsInformations) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.Experience); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceLevelFloor); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceNextLevelFloor); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceBonusLimit); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Kamas); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.StatsPoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.AdditionnalPoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SpellsPoints); err != nil {
		return err
	}

	if err := m.AlignmentInfos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.LifePoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MaxLifePoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.EnergyPoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MaxEnergyPoints); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.ActionPointsCurrent); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.MovementPointsCurrent); err != nil {
		return err
	}

	if err := m.Initiative.Serialize(w); err != nil {
		return err
	}

	if err := m.Prospecting.Serialize(w); err != nil {
		return err
	}

	if err := m.ActionPoints.Serialize(w); err != nil {
		return err
	}

	if err := m.MovementPoints.Serialize(w); err != nil {
		return err
	}

	if err := m.Strength.Serialize(w); err != nil {
		return err
	}

	if err := m.Vitality.Serialize(w); err != nil {
		return err
	}

	if err := m.Wisdom.Serialize(w); err != nil {
		return err
	}

	if err := m.Chance.Serialize(w); err != nil {
		return err
	}

	if err := m.Agility.Serialize(w); err != nil {
		return err
	}

	if err := m.Intelligence.Serialize(w); err != nil {
		return err
	}

	if err := m.Range.Serialize(w); err != nil {
		return err
	}

	if err := m.SummonableCreaturesBoost.Serialize(w); err != nil {
		return err
	}

	if err := m.Reflect.Serialize(w); err != nil {
		return err
	}

	if err := m.CriticalHit.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.CriticalHitWeapon); err != nil {
		return err
	}

	if err := m.CriticalMiss.Serialize(w); err != nil {
		return err
	}

	if err := m.HealBonus.Serialize(w); err != nil {
		return err
	}

	if err := m.AllDamagesBonus.Serialize(w); err != nil {
		return err
	}

	if err := m.WeaponDamagesBonusPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.DamagesBonusPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.TrapBonus.Serialize(w); err != nil {
		return err
	}

	if err := m.TrapBonusPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.GlyphBonusPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.RuneBonusPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.PermanentDamagePercent.Serialize(w); err != nil {
		return err
	}

	if err := m.TackleBlock.Serialize(w); err != nil {
		return err
	}

	if err := m.TackleEvade.Serialize(w); err != nil {
		return err
	}

	if err := m.PAAttack.Serialize(w); err != nil {
		return err
	}

	if err := m.PMAttack.Serialize(w); err != nil {
		return err
	}

	if err := m.PushDamageBonus.Serialize(w); err != nil {
		return err
	}

	if err := m.CriticalDamageBonus.Serialize(w); err != nil {
		return err
	}

	if err := m.NeutralDamageBonus.Serialize(w); err != nil {
		return err
	}

	if err := m.EarthDamageBonus.Serialize(w); err != nil {
		return err
	}

	if err := m.WaterDamageBonus.Serialize(w); err != nil {
		return err
	}

	if err := m.AirDamageBonus.Serialize(w); err != nil {
		return err
	}

	if err := m.FireDamageBonus.Serialize(w); err != nil {
		return err
	}

	if err := m.DodgePALostProbability.Serialize(w); err != nil {
		return err
	}

	if err := m.DodgePMLostProbability.Serialize(w); err != nil {
		return err
	}

	if err := m.NeutralElementResistPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.EarthElementResistPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.WaterElementResistPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.AirElementResistPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.FireElementResistPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.NeutralElementReduction.Serialize(w); err != nil {
		return err
	}

	if err := m.EarthElementReduction.Serialize(w); err != nil {
		return err
	}

	if err := m.WaterElementReduction.Serialize(w); err != nil {
		return err
	}

	if err := m.AirElementReduction.Serialize(w); err != nil {
		return err
	}

	if err := m.FireElementReduction.Serialize(w); err != nil {
		return err
	}

	if err := m.PushDamageReduction.Serialize(w); err != nil {
		return err
	}

	if err := m.CriticalDamageReduction.Serialize(w); err != nil {
		return err
	}

	if err := m.PvpNeutralElementResistPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.PvpEarthElementResistPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.PvpWaterElementResistPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.PvpAirElementResistPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.PvpFireElementResistPercent.Serialize(w); err != nil {
		return err
	}

	if err := m.PvpNeutralElementReduction.Serialize(w); err != nil {
		return err
	}

	if err := m.PvpEarthElementReduction.Serialize(w); err != nil {
		return err
	}

	if err := m.PvpWaterElementReduction.Serialize(w); err != nil {
		return err
	}

	if err := m.PvpAirElementReduction.Serialize(w); err != nil {
		return err
	}

	if err := m.PvpFireElementReduction.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.SpellModifications))); err != nil {
		return err
	}

	for i := range m.SpellModifications {

		if err := m.SpellModifications[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteUInt32(m.ProbationTime); err != nil {
		return err
	}

	return nil
}

func (m *CharacterCharacteristicsInformations) Deserialize(r Reader) error {

	lexperience, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Experience = lexperience

	lexperienceLevelFloor, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceLevelFloor = lexperienceLevelFloor

	lexperienceNextLevelFloor, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceNextLevelFloor = lexperienceNextLevelFloor

	lexperienceBonusLimit, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceBonusLimit = lexperienceBonusLimit

	lkamas, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Kamas = lkamas

	lstatsPoints, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.StatsPoints = lstatsPoints

	ladditionnalPoints, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.AdditionnalPoints = ladditionnalPoints

	lspellsPoints, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellsPoints = lspellsPoints

	var lalignmentInfos ActorExtendedAlignmentInformations

	lalignmentInfos.Deserialize(r)

	m.AlignmentInfos = &lalignmentInfos

	llifePoints, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.LifePoints = llifePoints

	lmaxLifePoints, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MaxLifePoints = lmaxLifePoints

	lenergyPoints, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.EnergyPoints = lenergyPoints

	lmaxEnergyPoints, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MaxEnergyPoints = lmaxEnergyPoints

	lactionPointsCurrent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.ActionPointsCurrent = lactionPointsCurrent

	lmovementPointsCurrent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.MovementPointsCurrent = lmovementPointsCurrent

	var linitiative CharacterBaseCharacteristic

	linitiative.Deserialize(r)

	m.Initiative = &linitiative

	var lprospecting CharacterBaseCharacteristic

	lprospecting.Deserialize(r)

	m.Prospecting = &lprospecting

	var lactionPoints CharacterBaseCharacteristic

	lactionPoints.Deserialize(r)

	m.ActionPoints = &lactionPoints

	var lmovementPoints CharacterBaseCharacteristic

	lmovementPoints.Deserialize(r)

	m.MovementPoints = &lmovementPoints

	var lstrength CharacterBaseCharacteristic

	lstrength.Deserialize(r)

	m.Strength = &lstrength

	var lvitality CharacterBaseCharacteristic

	lvitality.Deserialize(r)

	m.Vitality = &lvitality

	var lwisdom CharacterBaseCharacteristic

	lwisdom.Deserialize(r)

	m.Wisdom = &lwisdom

	var lchance CharacterBaseCharacteristic

	lchance.Deserialize(r)

	m.Chance = &lchance

	var lagility CharacterBaseCharacteristic

	lagility.Deserialize(r)

	m.Agility = &lagility

	var lintelligence CharacterBaseCharacteristic

	lintelligence.Deserialize(r)

	m.Intelligence = &lintelligence

	var lrange CharacterBaseCharacteristic

	lrange.Deserialize(r)

	m.Range = &lrange

	var lsummonableCreaturesBoost CharacterBaseCharacteristic

	lsummonableCreaturesBoost.Deserialize(r)

	m.SummonableCreaturesBoost = &lsummonableCreaturesBoost

	var lreflect CharacterBaseCharacteristic

	lreflect.Deserialize(r)

	m.Reflect = &lreflect

	var lcriticalHit CharacterBaseCharacteristic

	lcriticalHit.Deserialize(r)

	m.CriticalHit = &lcriticalHit

	lcriticalHitWeapon, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CriticalHitWeapon = lcriticalHitWeapon

	var lcriticalMiss CharacterBaseCharacteristic

	lcriticalMiss.Deserialize(r)

	m.CriticalMiss = &lcriticalMiss

	var lhealBonus CharacterBaseCharacteristic

	lhealBonus.Deserialize(r)

	m.HealBonus = &lhealBonus

	var lallDamagesBonus CharacterBaseCharacteristic

	lallDamagesBonus.Deserialize(r)

	m.AllDamagesBonus = &lallDamagesBonus

	var lweaponDamagesBonusPercent CharacterBaseCharacteristic

	lweaponDamagesBonusPercent.Deserialize(r)

	m.WeaponDamagesBonusPercent = &lweaponDamagesBonusPercent

	var ldamagesBonusPercent CharacterBaseCharacteristic

	ldamagesBonusPercent.Deserialize(r)

	m.DamagesBonusPercent = &ldamagesBonusPercent

	var ltrapBonus CharacterBaseCharacteristic

	ltrapBonus.Deserialize(r)

	m.TrapBonus = &ltrapBonus

	var ltrapBonusPercent CharacterBaseCharacteristic

	ltrapBonusPercent.Deserialize(r)

	m.TrapBonusPercent = &ltrapBonusPercent

	var lglyphBonusPercent CharacterBaseCharacteristic

	lglyphBonusPercent.Deserialize(r)

	m.GlyphBonusPercent = &lglyphBonusPercent

	var lruneBonusPercent CharacterBaseCharacteristic

	lruneBonusPercent.Deserialize(r)

	m.RuneBonusPercent = &lruneBonusPercent

	var lpermanentDamagePercent CharacterBaseCharacteristic

	lpermanentDamagePercent.Deserialize(r)

	m.PermanentDamagePercent = &lpermanentDamagePercent

	var ltackleBlock CharacterBaseCharacteristic

	ltackleBlock.Deserialize(r)

	m.TackleBlock = &ltackleBlock

	var ltackleEvade CharacterBaseCharacteristic

	ltackleEvade.Deserialize(r)

	m.TackleEvade = &ltackleEvade

	var lPAAttack CharacterBaseCharacteristic

	lPAAttack.Deserialize(r)

	m.PAAttack = &lPAAttack

	var lPMAttack CharacterBaseCharacteristic

	lPMAttack.Deserialize(r)

	m.PMAttack = &lPMAttack

	var lpushDamageBonus CharacterBaseCharacteristic

	lpushDamageBonus.Deserialize(r)

	m.PushDamageBonus = &lpushDamageBonus

	var lcriticalDamageBonus CharacterBaseCharacteristic

	lcriticalDamageBonus.Deserialize(r)

	m.CriticalDamageBonus = &lcriticalDamageBonus

	var lneutralDamageBonus CharacterBaseCharacteristic

	lneutralDamageBonus.Deserialize(r)

	m.NeutralDamageBonus = &lneutralDamageBonus

	var learthDamageBonus CharacterBaseCharacteristic

	learthDamageBonus.Deserialize(r)

	m.EarthDamageBonus = &learthDamageBonus

	var lwaterDamageBonus CharacterBaseCharacteristic

	lwaterDamageBonus.Deserialize(r)

	m.WaterDamageBonus = &lwaterDamageBonus

	var lairDamageBonus CharacterBaseCharacteristic

	lairDamageBonus.Deserialize(r)

	m.AirDamageBonus = &lairDamageBonus

	var lfireDamageBonus CharacterBaseCharacteristic

	lfireDamageBonus.Deserialize(r)

	m.FireDamageBonus = &lfireDamageBonus

	var ldodgePALostProbability CharacterBaseCharacteristic

	ldodgePALostProbability.Deserialize(r)

	m.DodgePALostProbability = &ldodgePALostProbability

	var ldodgePMLostProbability CharacterBaseCharacteristic

	ldodgePMLostProbability.Deserialize(r)

	m.DodgePMLostProbability = &ldodgePMLostProbability

	var lneutralElementResistPercent CharacterBaseCharacteristic

	lneutralElementResistPercent.Deserialize(r)

	m.NeutralElementResistPercent = &lneutralElementResistPercent

	var learthElementResistPercent CharacterBaseCharacteristic

	learthElementResistPercent.Deserialize(r)

	m.EarthElementResistPercent = &learthElementResistPercent

	var lwaterElementResistPercent CharacterBaseCharacteristic

	lwaterElementResistPercent.Deserialize(r)

	m.WaterElementResistPercent = &lwaterElementResistPercent

	var lairElementResistPercent CharacterBaseCharacteristic

	lairElementResistPercent.Deserialize(r)

	m.AirElementResistPercent = &lairElementResistPercent

	var lfireElementResistPercent CharacterBaseCharacteristic

	lfireElementResistPercent.Deserialize(r)

	m.FireElementResistPercent = &lfireElementResistPercent

	var lneutralElementReduction CharacterBaseCharacteristic

	lneutralElementReduction.Deserialize(r)

	m.NeutralElementReduction = &lneutralElementReduction

	var learthElementReduction CharacterBaseCharacteristic

	learthElementReduction.Deserialize(r)

	m.EarthElementReduction = &learthElementReduction

	var lwaterElementReduction CharacterBaseCharacteristic

	lwaterElementReduction.Deserialize(r)

	m.WaterElementReduction = &lwaterElementReduction

	var lairElementReduction CharacterBaseCharacteristic

	lairElementReduction.Deserialize(r)

	m.AirElementReduction = &lairElementReduction

	var lfireElementReduction CharacterBaseCharacteristic

	lfireElementReduction.Deserialize(r)

	m.FireElementReduction = &lfireElementReduction

	var lpushDamageReduction CharacterBaseCharacteristic

	lpushDamageReduction.Deserialize(r)

	m.PushDamageReduction = &lpushDamageReduction

	var lcriticalDamageReduction CharacterBaseCharacteristic

	lcriticalDamageReduction.Deserialize(r)

	m.CriticalDamageReduction = &lcriticalDamageReduction

	var lpvpNeutralElementResistPercent CharacterBaseCharacteristic

	lpvpNeutralElementResistPercent.Deserialize(r)

	m.PvpNeutralElementResistPercent = &lpvpNeutralElementResistPercent

	var lpvpEarthElementResistPercent CharacterBaseCharacteristic

	lpvpEarthElementResistPercent.Deserialize(r)

	m.PvpEarthElementResistPercent = &lpvpEarthElementResistPercent

	var lpvpWaterElementResistPercent CharacterBaseCharacteristic

	lpvpWaterElementResistPercent.Deserialize(r)

	m.PvpWaterElementResistPercent = &lpvpWaterElementResistPercent

	var lpvpAirElementResistPercent CharacterBaseCharacteristic

	lpvpAirElementResistPercent.Deserialize(r)

	m.PvpAirElementResistPercent = &lpvpAirElementResistPercent

	var lpvpFireElementResistPercent CharacterBaseCharacteristic

	lpvpFireElementResistPercent.Deserialize(r)

	m.PvpFireElementResistPercent = &lpvpFireElementResistPercent

	var lpvpNeutralElementReduction CharacterBaseCharacteristic

	lpvpNeutralElementReduction.Deserialize(r)

	m.PvpNeutralElementReduction = &lpvpNeutralElementReduction

	var lpvpEarthElementReduction CharacterBaseCharacteristic

	lpvpEarthElementReduction.Deserialize(r)

	m.PvpEarthElementReduction = &lpvpEarthElementReduction

	var lpvpWaterElementReduction CharacterBaseCharacteristic

	lpvpWaterElementReduction.Deserialize(r)

	m.PvpWaterElementReduction = &lpvpWaterElementReduction

	var lpvpAirElementReduction CharacterBaseCharacteristic

	lpvpAirElementReduction.Deserialize(r)

	m.PvpAirElementReduction = &lpvpAirElementReduction

	var lpvpFireElementReduction CharacterBaseCharacteristic

	lpvpFireElementReduction.Deserialize(r)

	m.PvpFireElementReduction = &lpvpFireElementReduction

	lspellModificationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SpellModifications = make([]*CharacterSpellModification, lspellModificationsLen)

	for i := range m.SpellModifications {

		var lspellModifications CharacterSpellModification

		lspellModifications.Deserialize(r)

		m.SpellModifications[i] = &lspellModifications

	}

	lprobationTime, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.ProbationTime = lprobationTime

	return nil
}

type PrismInformation struct {
	TypeId uint8

	State uint8

	NextVulnerabilityDate uint32

	PlacementDate uint32

	RewardTokenCount uint32
}

func (m *PrismInformation) ID() uint16 {
	return 428
}

func (m *PrismInformation) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.TypeId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.State); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.NextVulnerabilityDate); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.PlacementDate); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.RewardTokenCount); err != nil {
		return err
	}

	return nil
}

func (m *PrismInformation) Deserialize(r Reader) error {

	ltypeId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TypeId = ltypeId

	lstate, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.State = lstate

	lnextVulnerabilityDate, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.NextVulnerabilityDate = lnextVulnerabilityDate

	lplacementDate, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.PlacementDate = lplacementDate

	lrewardTokenCount, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.RewardTokenCount = lrewardTokenCount

	return nil
}

type AlliancePrismInformation struct {
	PrismInformation

	Alliance *AllianceInformations
}

func (m *AlliancePrismInformation) ID() uint16 {
	return 427
}

func (m *AlliancePrismInformation) Serialize(w Writer) error {

	if err := m.PrismInformation.Serialize(w); err != nil {
		return err
	}

	if err := m.Alliance.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AlliancePrismInformation) Deserialize(r Reader) error {

	if err := m.PrismInformation.Deserialize(r); err != nil {
		return err
	}

	var lalliance AllianceInformations

	lalliance.Deserialize(r)

	m.Alliance = &lalliance

	return nil
}

type DareReward struct {
	Type uint8

	MonsterId uint16

	Kamas int64

	DareId float64
}

func (m *DareReward) ID() uint16 {
	return 505
}

func (m *DareReward) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MonsterId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Kamas); err != nil {
		return err
	}

	if err := w.WriteDouble(m.DareId); err != nil {
		return err
	}

	return nil
}

func (m *DareReward) Deserialize(r Reader) error {

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	lmonsterId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MonsterId = lmonsterId

	lkamas, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Kamas = lkamas

	ldareId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DareId = ldareId

	return nil
}

type AllianceInsiderPrismInformation struct {
	PrismInformation

	LastTimeSlotModificationDate uint32

	LastTimeSlotModificationAuthorGuildId uint32

	LastTimeSlotModificationAuthorId int64

	LastTimeSlotModificationAuthorName string

	ModulesObjects []*ObjectItem
}

func (m *AllianceInsiderPrismInformation) ID() uint16 {
	return 431
}

func (m *AllianceInsiderPrismInformation) Serialize(w Writer) error {

	if err := m.PrismInformation.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.LastTimeSlotModificationDate); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.LastTimeSlotModificationAuthorGuildId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.LastTimeSlotModificationAuthorId); err != nil {
		return err
	}

	if err := w.WriteString(m.LastTimeSlotModificationAuthorName); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.ModulesObjects))); err != nil {
		return err
	}

	for i := range m.ModulesObjects {

		if err := m.ModulesObjects[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *AllianceInsiderPrismInformation) Deserialize(r Reader) error {

	if err := m.PrismInformation.Deserialize(r); err != nil {
		return err
	}

	llastTimeSlotModificationDate, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.LastTimeSlotModificationDate = llastTimeSlotModificationDate

	llastTimeSlotModificationAuthorGuildId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.LastTimeSlotModificationAuthorGuildId = llastTimeSlotModificationAuthorGuildId

	llastTimeSlotModificationAuthorId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.LastTimeSlotModificationAuthorId = llastTimeSlotModificationAuthorId

	llastTimeSlotModificationAuthorName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.LastTimeSlotModificationAuthorName = llastTimeSlotModificationAuthorName

	lmodulesObjectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ModulesObjects = make([]*ObjectItem, lmodulesObjectsLen)

	for i := range m.ModulesObjects {

		var lmodulesObjects ObjectItem

		lmodulesObjects.Deserialize(r)

		m.ModulesObjects[i] = &lmodulesObjects

	}

	return nil
}

type AbstractCharacterInformation struct {
	Id int64
}

func (m *AbstractCharacterInformation) ID() uint16 {
	return 400
}

func (m *AbstractCharacterInformation) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *AbstractCharacterInformation) Deserialize(r Reader) error {

	lid, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type CharacterBasicMinimalInformations struct {
	AbstractCharacterInformation

	Name string
}

func (m *CharacterBasicMinimalInformations) ID() uint16 {
	return 503
}

func (m *CharacterBasicMinimalInformations) Serialize(w Writer) error {

	if err := m.AbstractCharacterInformation.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CharacterBasicMinimalInformations) Deserialize(r Reader) error {

	if err := m.AbstractCharacterInformation.Deserialize(r); err != nil {
		return err
	}

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	return nil
}

type CharacterMinimalInformations struct {
	CharacterBasicMinimalInformations

	Level uint8
}

func (m *CharacterMinimalInformations) ID() uint16 {
	return 110
}

func (m *CharacterMinimalInformations) Serialize(w Writer) error {

	if err := m.CharacterBasicMinimalInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	return nil
}

func (m *CharacterMinimalInformations) Deserialize(r Reader) error {

	if err := m.CharacterBasicMinimalInformations.Deserialize(r); err != nil {
		return err
	}

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	return nil
}

type GuildMember struct {
	CharacterMinimalInformations

	Breed int8

	Sex bool

	Rank uint16

	GivenExperience int64

	ExperienceGivenPercent uint8

	Rights uint32

	Connected uint8

	AlignmentSide int8

	HoursSinceLastConnection uint16

	MoodSmileyId uint16

	AccountId uint32

	AchievementPoints int32

	Status *PlayerStatus

	HavenBagShared bool
}

func (m *GuildMember) ID() uint16 {
	return 88
}

func (m *GuildMember) Serialize(w Writer) error {

	if err := m.CharacterMinimalInformations.Serialize(w); err != nil {
		return err
	}

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Sex)

	setWrappedFlag(bbw0, 1, m.HavenBagShared)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Rank); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.GivenExperience); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.ExperienceGivenPercent); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Rights); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Connected); err != nil {
		return err
	}

	if err := w.WriteInt8(m.AlignmentSide); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.HoursSinceLastConnection); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MoodSmileyId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	if err := w.WriteInt32(m.AchievementPoints); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Status.ID()); err != nil {
		return err
	}

	if err := m.Status.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildMember) Deserialize(r Reader) error {

	if err := m.CharacterMinimalInformations.Deserialize(r); err != nil {
		return err
	}

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Sex = getWrappedFlag(bbw0, 0)

	m.HavenBagShared = getWrappedFlag(bbw0, 1)

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	lrank, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Rank = lrank

	lgivenExperience, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.GivenExperience = lgivenExperience

	lexperienceGivenPercent, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ExperienceGivenPercent = lexperienceGivenPercent

	lrights, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Rights = lrights

	lconnected, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Connected = lconnected

	lalignmentSide, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.AlignmentSide = lalignmentSide

	lhoursSinceLastConnection, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.HoursSinceLastConnection = lhoursSinceLastConnection

	lmoodSmileyId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MoodSmileyId = lmoodSmileyId

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	lachievementPoints, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.AchievementPoints = lachievementPoints

	typestatusID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lstatus, err := GetType(typestatusID)
	if err != nil {
		return err
	}

	lstatus.Deserialize(r)

	m.Status = lstatus.(*PlayerStatus)

	return nil
}

type EntityDispositionInformations struct {
	CellId int16

	Direction uint8
}

func (m *EntityDispositionInformations) ID() uint16 {
	return 60
}

func (m *EntityDispositionInformations) Serialize(w Writer) error {

	if err := w.WriteInt16(m.CellId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *EntityDispositionInformations) Deserialize(r Reader) error {

	lcellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	ldirection, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Direction = ldirection

	return nil
}

type GameFightMinimalStats struct {
	LifePoints uint32

	MaxLifePoints uint32

	BaseMaxLifePoints uint32

	PermanentDamagePercent uint32

	ShieldPoints uint32

	ActionPoints int16

	MaxActionPoints int16

	MovementPoints int16

	MaxMovementPoints int16

	Summoner float64

	Summoned bool

	NeutralElementResistPercent int16

	EarthElementResistPercent int16

	WaterElementResistPercent int16

	AirElementResistPercent int16

	FireElementResistPercent int16

	NeutralElementReduction int16

	EarthElementReduction int16

	WaterElementReduction int16

	AirElementReduction int16

	FireElementReduction int16

	CriticalDamageFixedResist int16

	PushDamageFixedResist int16

	PvpNeutralElementResistPercent int16

	PvpEarthElementResistPercent int16

	PvpWaterElementResistPercent int16

	PvpAirElementResistPercent int16

	PvpFireElementResistPercent int16

	PvpNeutralElementReduction int16

	PvpEarthElementReduction int16

	PvpWaterElementReduction int16

	PvpAirElementReduction int16

	PvpFireElementReduction int16

	DodgePALostProbability uint16

	DodgePMLostProbability uint16

	TackleBlock int16

	TackleEvade int16

	FixedDamageReflection int16

	InvisibilityState uint8
}

func (m *GameFightMinimalStats) ID() uint16 {
	return 31
}

func (m *GameFightMinimalStats) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.LifePoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MaxLifePoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.BaseMaxLifePoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.PermanentDamagePercent); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ShieldPoints); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.ActionPoints); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.MaxActionPoints); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.MovementPoints); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.MaxMovementPoints); err != nil {
		return err
	}

	if err := w.WriteDouble(m.Summoner); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Summoned); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.NeutralElementResistPercent); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.EarthElementResistPercent); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.WaterElementResistPercent); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.AirElementResistPercent); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.FireElementResistPercent); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.NeutralElementReduction); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.EarthElementReduction); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.WaterElementReduction); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.AirElementReduction); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.FireElementReduction); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.CriticalDamageFixedResist); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.PushDamageFixedResist); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.PvpNeutralElementResistPercent); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.PvpEarthElementResistPercent); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.PvpWaterElementResistPercent); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.PvpAirElementResistPercent); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.PvpFireElementResistPercent); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.PvpNeutralElementReduction); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.PvpEarthElementReduction); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.PvpWaterElementReduction); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.PvpAirElementReduction); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.PvpFireElementReduction); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DodgePALostProbability); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DodgePMLostProbability); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.TackleBlock); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.TackleEvade); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.FixedDamageReflection); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.InvisibilityState); err != nil {
		return err
	}

	return nil
}

func (m *GameFightMinimalStats) Deserialize(r Reader) error {

	llifePoints, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.LifePoints = llifePoints

	lmaxLifePoints, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MaxLifePoints = lmaxLifePoints

	lbaseMaxLifePoints, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.BaseMaxLifePoints = lbaseMaxLifePoints

	lpermanentDamagePercent, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.PermanentDamagePercent = lpermanentDamagePercent

	lshieldPoints, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ShieldPoints = lshieldPoints

	lactionPoints, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.ActionPoints = lactionPoints

	lmaxActionPoints, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.MaxActionPoints = lmaxActionPoints

	lmovementPoints, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.MovementPoints = lmovementPoints

	lmaxMovementPoints, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.MaxMovementPoints = lmaxMovementPoints

	lsummoner, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Summoner = lsummoner

	lsummoned, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Summoned = lsummoned

	lneutralElementResistPercent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.NeutralElementResistPercent = lneutralElementResistPercent

	learthElementResistPercent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.EarthElementResistPercent = learthElementResistPercent

	lwaterElementResistPercent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.WaterElementResistPercent = lwaterElementResistPercent

	lairElementResistPercent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.AirElementResistPercent = lairElementResistPercent

	lfireElementResistPercent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.FireElementResistPercent = lfireElementResistPercent

	lneutralElementReduction, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.NeutralElementReduction = lneutralElementReduction

	learthElementReduction, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.EarthElementReduction = learthElementReduction

	lwaterElementReduction, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.WaterElementReduction = lwaterElementReduction

	lairElementReduction, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.AirElementReduction = lairElementReduction

	lfireElementReduction, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.FireElementReduction = lfireElementReduction

	lcriticalDamageFixedResist, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.CriticalDamageFixedResist = lcriticalDamageFixedResist

	lpushDamageFixedResist, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.PushDamageFixedResist = lpushDamageFixedResist

	lpvpNeutralElementResistPercent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.PvpNeutralElementResistPercent = lpvpNeutralElementResistPercent

	lpvpEarthElementResistPercent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.PvpEarthElementResistPercent = lpvpEarthElementResistPercent

	lpvpWaterElementResistPercent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.PvpWaterElementResistPercent = lpvpWaterElementResistPercent

	lpvpAirElementResistPercent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.PvpAirElementResistPercent = lpvpAirElementResistPercent

	lpvpFireElementResistPercent, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.PvpFireElementResistPercent = lpvpFireElementResistPercent

	lpvpNeutralElementReduction, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.PvpNeutralElementReduction = lpvpNeutralElementReduction

	lpvpEarthElementReduction, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.PvpEarthElementReduction = lpvpEarthElementReduction

	lpvpWaterElementReduction, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.PvpWaterElementReduction = lpvpWaterElementReduction

	lpvpAirElementReduction, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.PvpAirElementReduction = lpvpAirElementReduction

	lpvpFireElementReduction, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.PvpFireElementReduction = lpvpFireElementReduction

	ldodgePALostProbability, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DodgePALostProbability = ldodgePALostProbability

	ldodgePMLostProbability, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DodgePMLostProbability = ldodgePMLostProbability

	ltackleBlock, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.TackleBlock = ltackleBlock

	ltackleEvade, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.TackleEvade = ltackleEvade

	lfixedDamageReflection, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.FixedDamageReflection = lfixedDamageReflection

	linvisibilityState, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.InvisibilityState = linvisibilityState

	return nil
}

type InteractiveElementSkill struct {
	SkillId uint32

	SkillInstanceUid uint32
}

func (m *InteractiveElementSkill) ID() uint16 {
	return 219
}

func (m *InteractiveElementSkill) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.SkillId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.SkillInstanceUid); err != nil {
		return err
	}

	return nil
}

func (m *InteractiveElementSkill) Deserialize(r Reader) error {

	lskillId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SkillId = lskillId

	lskillInstanceUid, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.SkillInstanceUid = lskillInstanceUid

	return nil
}

type ObjectEffect struct {
	ActionId uint16
}

func (m *ObjectEffect) ID() uint16 {
	return 76
}

func (m *ObjectEffect) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ActionId); err != nil {
		return err
	}

	return nil
}

func (m *ObjectEffect) Deserialize(r Reader) error {

	lactionId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ActionId = lactionId

	return nil
}

type ObjectEffectInteger struct {
	ObjectEffect

	Value uint16
}

func (m *ObjectEffectInteger) ID() uint16 {
	return 70
}

func (m *ObjectEffectInteger) Serialize(w Writer) error {

	if err := m.ObjectEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *ObjectEffectInteger) Deserialize(r Reader) error {

	if err := m.ObjectEffect.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type PortalInformation struct {
	PortalId int32

	AreaId int16
}

func (m *PortalInformation) ID() uint16 {
	return 466
}

func (m *PortalInformation) Serialize(w Writer) error {

	if err := w.WriteInt32(m.PortalId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.AreaId); err != nil {
		return err
	}

	return nil
}

func (m *PortalInformation) Deserialize(r Reader) error {

	lportalId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.PortalId = lportalId

	lareaId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AreaId = lareaId

	return nil
}

type GroupMonsterStaticInformations struct {
	MainCreatureLightInfos *MonsterInGroupLightInformations

	Underlings []*MonsterInGroupInformations
}

func (m *GroupMonsterStaticInformations) ID() uint16 {
	return 140
}

func (m *GroupMonsterStaticInformations) Serialize(w Writer) error {

	if err := m.MainCreatureLightInfos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Underlings))); err != nil {
		return err
	}

	for i := range m.Underlings {

		if err := m.Underlings[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GroupMonsterStaticInformations) Deserialize(r Reader) error {

	var lmainCreatureLightInfos MonsterInGroupLightInformations

	lmainCreatureLightInfos.Deserialize(r)

	m.MainCreatureLightInfos = &lmainCreatureLightInfos

	lunderlingsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Underlings = make([]*MonsterInGroupInformations, lunderlingsLen)

	for i := range m.Underlings {

		var lunderlings MonsterInGroupInformations

		lunderlings.Deserialize(r)

		m.Underlings[i] = &lunderlings

	}

	return nil
}

type AchievementRewardable struct {
	Id uint16

	Finishedlevel uint8
}

func (m *AchievementRewardable) ID() uint16 {
	return 412
}

func (m *AchievementRewardable) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Id); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Finishedlevel); err != nil {
		return err
	}

	return nil
}

func (m *AchievementRewardable) Deserialize(r Reader) error {

	lid, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Id = lid

	lfinishedlevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Finishedlevel = lfinishedlevel

	return nil
}

type QuestActiveInformations struct {
	QuestId uint16
}

func (m *QuestActiveInformations) ID() uint16 {
	return 381
}

func (m *QuestActiveInformations) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.QuestId); err != nil {
		return err
	}

	return nil
}

func (m *QuestActiveInformations) Deserialize(r Reader) error {

	lquestId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.QuestId = lquestId

	return nil
}

type JobDescription struct {
	JobId uint8

	Skills []*SkillActionDescription
}

func (m *JobDescription) ID() uint16 {
	return 101
}

func (m *JobDescription) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.JobId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Skills))); err != nil {
		return err
	}

	for i := range m.Skills {

		if err := w.WriteUInt16(m.Skills[i].ID()); err != nil {
			return err
		}

		if err := m.Skills[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *JobDescription) Deserialize(r Reader) error {

	ljobId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.JobId = ljobId

	lskillsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Skills = make([]*SkillActionDescription, lskillsLen)

	for i := range m.Skills {

		typeskillsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lskills, err := GetType(typeskillsID)
		if err != nil {
			return err
		}

		lskills.Deserialize(r)

		m.Skills[i] = lskills.(*SkillActionDescription)

	}

	return nil
}

type SkillActionDescription struct {
	SkillId uint16
}

func (m *SkillActionDescription) ID() uint16 {
	return 102
}

func (m *SkillActionDescription) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SkillId); err != nil {
		return err
	}

	return nil
}

func (m *SkillActionDescription) Deserialize(r Reader) error {

	lskillId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SkillId = lskillId

	return nil
}

type SkillActionDescriptionCraft struct {
	SkillActionDescription

	Probability uint8
}

func (m *SkillActionDescriptionCraft) ID() uint16 {
	return 100
}

func (m *SkillActionDescriptionCraft) Serialize(w Writer) error {

	if err := m.SkillActionDescription.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Probability); err != nil {
		return err
	}

	return nil
}

func (m *SkillActionDescriptionCraft) Deserialize(r Reader) error {

	if err := m.SkillActionDescription.Deserialize(r); err != nil {
		return err
	}

	lprobability, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Probability = lprobability

	return nil
}

type SkillActionDescriptionTimed struct {
	SkillActionDescription

	Time uint8
}

func (m *SkillActionDescriptionTimed) ID() uint16 {
	return 103
}

func (m *SkillActionDescriptionTimed) Serialize(w Writer) error {

	if err := m.SkillActionDescription.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Time); err != nil {
		return err
	}

	return nil
}

func (m *SkillActionDescriptionTimed) Deserialize(r Reader) error {

	if err := m.SkillActionDescription.Deserialize(r); err != nil {
		return err
	}

	ltime, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Time = ltime

	return nil
}

type SkillActionDescriptionCollect struct {
	SkillActionDescriptionTimed

	Min uint16

	Max uint16
}

func (m *SkillActionDescriptionCollect) ID() uint16 {
	return 99
}

func (m *SkillActionDescriptionCollect) Serialize(w Writer) error {

	if err := m.SkillActionDescriptionTimed.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Min); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Max); err != nil {
		return err
	}

	return nil
}

func (m *SkillActionDescriptionCollect) Deserialize(r Reader) error {

	if err := m.SkillActionDescriptionTimed.Deserialize(r); err != nil {
		return err
	}

	lmin, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Min = lmin

	lmax, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Max = lmax

	return nil
}

type ActorRestrictionsInformations struct {
	CantBeAggressed bool

	CantBeChallenged bool

	CantTrade bool

	CantBeAttackedByMutant bool

	CantRun bool

	ForceSlowWalk bool

	CantMinimize bool

	CantMove bool

	CantAggress bool

	CantChallenge bool

	CantExchange bool

	CantAttack bool

	CantChat bool

	CantBeMerchant bool

	CantUseObject bool

	CantUseTaxCollector bool

	CantUseInteractive bool

	CantSpeakToNPC bool

	CantChangeZone bool

	CantAttackMonster bool

	CantWalk8Directions bool
}

func (m *ActorRestrictionsInformations) ID() uint16 {
	return 204
}

func (m *ActorRestrictionsInformations) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.CantBeAggressed)

	setWrappedFlag(bbw0, 1, m.CantBeChallenged)

	setWrappedFlag(bbw0, 2, m.CantTrade)

	setWrappedFlag(bbw0, 3, m.CantBeAttackedByMutant)

	setWrappedFlag(bbw0, 4, m.CantRun)

	setWrappedFlag(bbw0, 5, m.ForceSlowWalk)

	setWrappedFlag(bbw0, 6, m.CantMinimize)

	setWrappedFlag(bbw0, 7, m.CantMove)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	var bbw1 uint8

	setWrappedFlag(bbw1, 0, m.CantAggress)

	setWrappedFlag(bbw1, 1, m.CantChallenge)

	setWrappedFlag(bbw1, 2, m.CantExchange)

	setWrappedFlag(bbw1, 3, m.CantAttack)

	setWrappedFlag(bbw1, 4, m.CantChat)

	setWrappedFlag(bbw1, 5, m.CantBeMerchant)

	setWrappedFlag(bbw1, 6, m.CantUseObject)

	setWrappedFlag(bbw1, 7, m.CantUseTaxCollector)

	if err := w.WriteUInt8(bbw1); err != nil {
		return err
	}

	var bbw2 uint8

	setWrappedFlag(bbw2, 0, m.CantUseInteractive)

	setWrappedFlag(bbw2, 1, m.CantSpeakToNPC)

	setWrappedFlag(bbw2, 2, m.CantChangeZone)

	setWrappedFlag(bbw2, 3, m.CantAttackMonster)

	setWrappedFlag(bbw2, 4, m.CantWalk8Directions)

	if err := w.WriteUInt8(bbw2); err != nil {
		return err
	}

	return nil
}

func (m *ActorRestrictionsInformations) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CantBeAggressed = getWrappedFlag(bbw0, 0)

	m.CantBeChallenged = getWrappedFlag(bbw0, 1)

	m.CantTrade = getWrappedFlag(bbw0, 2)

	m.CantBeAttackedByMutant = getWrappedFlag(bbw0, 3)

	m.CantRun = getWrappedFlag(bbw0, 4)

	m.ForceSlowWalk = getWrappedFlag(bbw0, 5)

	m.CantMinimize = getWrappedFlag(bbw0, 6)

	m.CantMove = getWrappedFlag(bbw0, 7)

	bbw1, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CantAggress = getWrappedFlag(bbw1, 0)

	m.CantChallenge = getWrappedFlag(bbw1, 1)

	m.CantExchange = getWrappedFlag(bbw1, 2)

	m.CantAttack = getWrappedFlag(bbw1, 3)

	m.CantChat = getWrappedFlag(bbw1, 4)

	m.CantBeMerchant = getWrappedFlag(bbw1, 5)

	m.CantUseObject = getWrappedFlag(bbw1, 6)

	m.CantUseTaxCollector = getWrappedFlag(bbw1, 7)

	bbw2, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CantUseInteractive = getWrappedFlag(bbw2, 0)

	m.CantSpeakToNPC = getWrappedFlag(bbw2, 1)

	m.CantChangeZone = getWrappedFlag(bbw2, 2)

	m.CantAttackMonster = getWrappedFlag(bbw2, 3)

	m.CantWalk8Directions = getWrappedFlag(bbw2, 4)

	return nil
}

type HumanOption struct {
}

func (m *HumanOption) ID() uint16 {
	return 406
}

func (m *HumanOption) Serialize(w Writer) error {

	return nil
}

func (m *HumanOption) Deserialize(r Reader) error {

	return nil
}

type HumanOptionTitle struct {
	HumanOption

	TitleId uint16

	TitleParam string
}

func (m *HumanOptionTitle) ID() uint16 {
	return 408
}

func (m *HumanOptionTitle) Serialize(w Writer) error {

	if err := m.HumanOption.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TitleId); err != nil {
		return err
	}

	if err := w.WriteString(m.TitleParam); err != nil {
		return err
	}

	return nil
}

func (m *HumanOptionTitle) Deserialize(r Reader) error {

	if err := m.HumanOption.Deserialize(r); err != nil {
		return err
	}

	ltitleId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TitleId = ltitleId

	ltitleParam, err := r.ReadString()
	if err != nil {
		return err
	}

	m.TitleParam = ltitleParam

	return nil
}

type CharacterMinimalPlusLookInformations struct {
	CharacterMinimalInformations

	EntityLook *EntityLook
}

func (m *CharacterMinimalPlusLookInformations) ID() uint16 {
	return 163
}

func (m *CharacterMinimalPlusLookInformations) Serialize(w Writer) error {

	if err := m.CharacterMinimalInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.EntityLook.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *CharacterMinimalPlusLookInformations) Deserialize(r Reader) error {

	if err := m.CharacterMinimalInformations.Deserialize(r); err != nil {
		return err
	}

	var lentityLook EntityLook

	lentityLook.Deserialize(r)

	m.EntityLook = &lentityLook

	return nil
}

type CharacterBaseInformations struct {
	CharacterMinimalPlusLookInformations

	Breed int8

	Sex bool
}

func (m *CharacterBaseInformations) ID() uint16 {
	return 45
}

func (m *CharacterBaseInformations) Serialize(w Writer) error {

	if err := m.CharacterMinimalPlusLookInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Sex); err != nil {
		return err
	}

	return nil
}

func (m *CharacterBaseInformations) Deserialize(r Reader) error {

	if err := m.CharacterMinimalPlusLookInformations.Deserialize(r); err != nil {
		return err
	}

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	lsex, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Sex = lsex

	return nil
}

type CharacterSpellModification struct {
	ModificationType uint8

	SpellId uint16

	Value *CharacterBaseCharacteristic
}

func (m *CharacterSpellModification) ID() uint16 {
	return 215
}

func (m *CharacterSpellModification) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.ModificationType); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	if err := m.Value.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *CharacterSpellModification) Deserialize(r Reader) error {

	lmodificationType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ModificationType = lmodificationType

	lspellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	var lvalue CharacterBaseCharacteristic

	lvalue.Deserialize(r)

	m.Value = &lvalue

	return nil
}

type HouseInformations struct {
	HouseId uint32

	ModelId uint16
}

func (m *HouseInformations) ID() uint16 {
	return 111
}

func (m *HouseInformations) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HouseId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ModelId); err != nil {
		return err
	}

	return nil
}

func (m *HouseInformations) Deserialize(r Reader) error {

	lhouseId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HouseId = lhouseId

	lmodelId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ModelId = lmodelId

	return nil
}

type AccountHouseInformations struct {
	HouseInformations

	InstanceId uint32

	SecondHand bool

	WorldX int16

	WorldY int16

	MapId int32

	SubAreaId uint16
}

func (m *AccountHouseInformations) ID() uint16 {
	return 390
}

func (m *AccountHouseInformations) Serialize(w Writer) error {

	if err := m.HouseInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.SecondHand); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	return nil
}

func (m *AccountHouseInformations) Deserialize(r Reader) error {

	if err := m.HouseInformations.Deserialize(r); err != nil {
		return err
	}

	linstanceId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.InstanceId = linstanceId

	lsecondHand, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.SecondHand = lsecondHand

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	return nil
}

type TaxCollectorStaticInformations struct {
	FirstNameId uint16

	LastNameId uint16

	GuildIdentity *GuildInformations
}

func (m *TaxCollectorStaticInformations) ID() uint16 {
	return 147
}

func (m *TaxCollectorStaticInformations) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.FirstNameId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.LastNameId); err != nil {
		return err
	}

	if err := m.GuildIdentity.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorStaticInformations) Deserialize(r Reader) error {

	lfirstNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.FirstNameId = lfirstNameId

	llastNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.LastNameId = llastNameId

	var lguildIdentity GuildInformations

	lguildIdentity.Deserialize(r)

	m.GuildIdentity = &lguildIdentity

	return nil
}

type CharacterBaseCharacteristic struct {
	Base int16

	Additionnal int16

	ObjectsAndMountBonus int16

	AlignGiftBonus int16

	ContextModif int16
}

func (m *CharacterBaseCharacteristic) ID() uint16 {
	return 4
}

func (m *CharacterBaseCharacteristic) Serialize(w Writer) error {

	if err := w.WriteVarInt16(m.Base); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.Additionnal); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.ObjectsAndMountBonus); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.AlignGiftBonus); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.ContextModif); err != nil {
		return err
	}

	return nil
}

func (m *CharacterBaseCharacteristic) Deserialize(r Reader) error {

	lbase, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.Base = lbase

	ladditionnal, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.Additionnal = ladditionnal

	lobjectsAndMountBonus, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.ObjectsAndMountBonus = lobjectsAndMountBonus

	lalignGiftBonus, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.AlignGiftBonus = lalignGiftBonus

	lcontextModif, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.ContextModif = lcontextModif

	return nil
}

type GameServerInformations struct {
	Id uint16

	Type int8

	Status uint8

	Completion uint8

	IsSelectable bool

	CharactersCount uint8

	CharactersSlots uint8

	Date float64
}

func (m *GameServerInformations) ID() uint16 {
	return 25
}

func (m *GameServerInformations) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Id); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Type); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Status); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Completion); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.IsSelectable); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.CharactersCount); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.CharactersSlots); err != nil {
		return err
	}

	if err := w.WriteDouble(m.Date); err != nil {
		return err
	}

	return nil
}

func (m *GameServerInformations) Deserialize(r Reader) error {

	lid, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Id = lid

	ltype, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	lstatus, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Status = lstatus

	lcompletion, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Completion = lcompletion

	lisSelectable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsSelectable = lisSelectable

	lcharactersCount, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CharactersCount = lcharactersCount

	lcharactersSlots, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CharactersSlots = lcharactersSlots

	ldate, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Date = ldate

	return nil
}

type PlayerStatus struct {
	StatusId uint8
}

func (m *PlayerStatus) ID() uint16 {
	return 415
}

func (m *PlayerStatus) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.StatusId); err != nil {
		return err
	}

	return nil
}

func (m *PlayerStatus) Deserialize(r Reader) error {

	lstatusId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.StatusId = lstatusId

	return nil
}

type ActorAlignmentInformations struct {
	AlignmentSide int8

	AlignmentValue uint8

	AlignmentGrade uint8

	CharacterPower float64
}

func (m *ActorAlignmentInformations) ID() uint16 {
	return 201
}

func (m *ActorAlignmentInformations) Serialize(w Writer) error {

	if err := w.WriteInt8(m.AlignmentSide); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.AlignmentValue); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.AlignmentGrade); err != nil {
		return err
	}

	if err := w.WriteDouble(m.CharacterPower); err != nil {
		return err
	}

	return nil
}

func (m *ActorAlignmentInformations) Deserialize(r Reader) error {

	lalignmentSide, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.AlignmentSide = lalignmentSide

	lalignmentValue, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.AlignmentValue = lalignmentValue

	lalignmentGrade, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.AlignmentGrade = lalignmentGrade

	lcharacterPower, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.CharacterPower = lcharacterPower

	return nil
}

type HumanOptionOrnament struct {
	HumanOption

	OrnamentId uint16
}

func (m *HumanOptionOrnament) ID() uint16 {
	return 411
}

func (m *HumanOptionOrnament) Serialize(w Writer) error {

	if err := m.HumanOption.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.OrnamentId); err != nil {
		return err
	}

	return nil
}

func (m *HumanOptionOrnament) Deserialize(r Reader) error {

	if err := m.HumanOption.Deserialize(r); err != nil {
		return err
	}

	lornamentId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.OrnamentId = lornamentId

	return nil
}

type ItemDurability struct {
	Durability int16

	DurabilityMax int16
}

func (m *ItemDurability) ID() uint16 {
	return 168
}

func (m *ItemDurability) Serialize(w Writer) error {

	if err := w.WriteInt16(m.Durability); err != nil {
		return err
	}

	if err := w.WriteInt16(m.DurabilityMax); err != nil {
		return err
	}

	return nil
}

func (m *ItemDurability) Deserialize(r Reader) error {

	ldurability, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Durability = ldurability

	ldurabilityMax, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DurabilityMax = ldurabilityMax

	return nil
}

type HumanInformations struct {
	Restrictions *ActorRestrictionsInformations

	Sex bool

	Options []*HumanOption
}

func (m *HumanInformations) ID() uint16 {
	return 157
}

func (m *HumanInformations) Serialize(w Writer) error {

	if err := m.Restrictions.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Sex); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Options))); err != nil {
		return err
	}

	for i := range m.Options {

		if err := w.WriteUInt16(m.Options[i].ID()); err != nil {
			return err
		}

		if err := m.Options[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *HumanInformations) Deserialize(r Reader) error {

	var lrestrictions ActorRestrictionsInformations

	lrestrictions.Deserialize(r)

	m.Restrictions = &lrestrictions

	lsex, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Sex = lsex

	loptionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Options = make([]*HumanOption, loptionsLen)

	for i := range m.Options {

		typeoptionsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		loptions, err := GetType(typeoptionsID)
		if err != nil {
			return err
		}

		loptions.Deserialize(r)

		m.Options[i] = loptions.(*HumanOption)

	}

	return nil
}

type MonsterInGroupLightInformations struct {
	CreatureGenericId int32

	Grade uint8
}

func (m *MonsterInGroupLightInformations) ID() uint16 {
	return 395
}

func (m *MonsterInGroupLightInformations) Serialize(w Writer) error {

	if err := w.WriteInt32(m.CreatureGenericId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Grade); err != nil {
		return err
	}

	return nil
}

func (m *MonsterInGroupLightInformations) Deserialize(r Reader) error {

	lcreatureGenericId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.CreatureGenericId = lcreatureGenericId

	lgrade, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Grade = lgrade

	return nil
}

type GroupMonsterStaticInformationsWithAlternatives struct {
	GroupMonsterStaticInformations

	Alternatives []*AlternativeMonstersInGroupLightInformations
}

func (m *GroupMonsterStaticInformationsWithAlternatives) ID() uint16 {
	return 396
}

func (m *GroupMonsterStaticInformationsWithAlternatives) Serialize(w Writer) error {

	if err := m.GroupMonsterStaticInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Alternatives))); err != nil {
		return err
	}

	for i := range m.Alternatives {

		if err := m.Alternatives[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GroupMonsterStaticInformationsWithAlternatives) Deserialize(r Reader) error {

	if err := m.GroupMonsterStaticInformations.Deserialize(r); err != nil {
		return err
	}

	lalternativesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Alternatives = make([]*AlternativeMonstersInGroupLightInformations, lalternativesLen)

	for i := range m.Alternatives {

		var lalternatives AlternativeMonstersInGroupLightInformations

		lalternatives.Deserialize(r)

		m.Alternatives[i] = &lalternatives

	}

	return nil
}

type MonsterInGroupInformations struct {
	MonsterInGroupLightInformations

	Look *EntityLook
}

func (m *MonsterInGroupInformations) ID() uint16 {
	return 144
}

func (m *MonsterInGroupInformations) Serialize(w Writer) error {

	if err := m.MonsterInGroupLightInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.Look.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *MonsterInGroupInformations) Deserialize(r Reader) error {

	if err := m.MonsterInGroupLightInformations.Deserialize(r); err != nil {
		return err
	}

	var llook EntityLook

	llook.Deserialize(r)

	m.Look = &llook

	return nil
}

type AlternativeMonstersInGroupLightInformations struct {
	PlayerCount int32

	Monsters []*MonsterInGroupLightInformations
}

func (m *AlternativeMonstersInGroupLightInformations) ID() uint16 {
	return 394
}

func (m *AlternativeMonstersInGroupLightInformations) Serialize(w Writer) error {

	if err := w.WriteInt32(m.PlayerCount); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Monsters))); err != nil {
		return err
	}

	for i := range m.Monsters {

		if err := m.Monsters[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *AlternativeMonstersInGroupLightInformations) Deserialize(r Reader) error {

	lplayerCount, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.PlayerCount = lplayerCount

	lmonstersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Monsters = make([]*MonsterInGroupLightInformations, lmonstersLen)

	for i := range m.Monsters {

		var lmonsters MonsterInGroupLightInformations

		lmonsters.Deserialize(r)

		m.Monsters[i] = &lmonsters

	}

	return nil
}

type PaddockInformations struct {
	MaxOutdoorMount uint16

	MaxItems uint16
}

func (m *PaddockInformations) ID() uint16 {
	return 132
}

func (m *PaddockInformations) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.MaxOutdoorMount); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MaxItems); err != nil {
		return err
	}

	return nil
}

func (m *PaddockInformations) Deserialize(r Reader) error {

	lmaxOutdoorMount, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MaxOutdoorMount = lmaxOutdoorMount

	lmaxItems, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MaxItems = lmaxItems

	return nil
}

type PaddockInstancesInformations struct {
	PaddockInformations

	Paddocks []*PaddockBuyableInformations
}

func (m *PaddockInstancesInformations) ID() uint16 {
	return 509
}

func (m *PaddockInstancesInformations) Serialize(w Writer) error {

	if err := m.PaddockInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Paddocks))); err != nil {
		return err
	}

	for i := range m.Paddocks {

		if err := w.WriteUInt16(m.Paddocks[i].ID()); err != nil {
			return err
		}

		if err := m.Paddocks[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PaddockInstancesInformations) Deserialize(r Reader) error {

	if err := m.PaddockInformations.Deserialize(r); err != nil {
		return err
	}

	lpaddocksLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Paddocks = make([]*PaddockBuyableInformations, lpaddocksLen)

	for i := range m.Paddocks {

		typepaddocksID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lpaddocks, err := GetType(typepaddocksID)
		if err != nil {
			return err
		}

		lpaddocks.Deserialize(r)

		m.Paddocks[i] = lpaddocks.(*PaddockBuyableInformations)

	}

	return nil
}

type PaddockBuyableInformations struct {
	Price int64

	Locked bool
}

func (m *PaddockBuyableInformations) ID() uint16 {
	return 130
}

func (m *PaddockBuyableInformations) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.Price); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Locked); err != nil {
		return err
	}

	return nil
}

func (m *PaddockBuyableInformations) Deserialize(r Reader) error {

	lprice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Price = lprice

	llocked, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Locked = llocked

	return nil
}

type HouseInformationsInside struct {
	HouseInformations

	InstanceId uint32

	SecondHand bool

	OwnerId int32

	OwnerName string

	WorldX int16

	WorldY int16

	Price int64

	IsLocked bool
}

func (m *HouseInformationsInside) ID() uint16 {
	return 218
}

func (m *HouseInformationsInside) Serialize(w Writer) error {

	if err := m.HouseInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.SecondHand); err != nil {
		return err
	}

	if err := w.WriteInt32(m.OwnerId); err != nil {
		return err
	}

	if err := w.WriteString(m.OwnerName); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Price); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.IsLocked); err != nil {
		return err
	}

	return nil
}

func (m *HouseInformationsInside) Deserialize(r Reader) error {

	if err := m.HouseInformations.Deserialize(r); err != nil {
		return err
	}

	linstanceId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.InstanceId = linstanceId

	lsecondHand, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.SecondHand = lsecondHand

	lownerId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.OwnerId = lownerId

	lownerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.OwnerName = lownerName

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lprice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Price = lprice

	lisLocked, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsLocked = lisLocked

	return nil
}

type HouseOnMapInformations struct {
	HouseInformations

	DoorsOnMap []uint32

	HouseInstances []*HouseInstanceInformations
}

func (m *HouseOnMapInformations) ID() uint16 {
	return 510
}

func (m *HouseOnMapInformations) Serialize(w Writer) error {

	if err := m.HouseInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.DoorsOnMap))); err != nil {
		return err
	}

	for i := range m.DoorsOnMap {

		if err := w.WriteUInt32(m.DoorsOnMap[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.HouseInstances))); err != nil {
		return err
	}

	for i := range m.HouseInstances {

		if err := m.HouseInstances[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *HouseOnMapInformations) Deserialize(r Reader) error {

	if err := m.HouseInformations.Deserialize(r); err != nil {
		return err
	}

	ldoorsOnMapLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DoorsOnMap = make([]uint32, ldoorsOnMapLen)

	for i := range m.DoorsOnMap {

		ldoorsOnMap, err := r.ReadUInt32()
		if err != nil {
			return err
		}

		m.DoorsOnMap[i] = ldoorsOnMap

	}

	lhouseInstancesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.HouseInstances = make([]*HouseInstanceInformations, lhouseInstancesLen)

	for i := range m.HouseInstances {

		var lhouseInstances HouseInstanceInformations

		lhouseInstances.Deserialize(r)

		m.HouseInstances[i] = &lhouseInstances

	}

	return nil
}

type HouseInstanceInformations struct {
	InstanceId uint32

	SecondHand bool

	OwnerName string

	IsOnSale bool

	IsSaleLocked bool
}

func (m *HouseInstanceInformations) ID() uint16 {
	return 511
}

func (m *HouseInstanceInformations) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.SecondHand)

	setWrappedFlag(bbw0, 1, m.IsOnSale)

	setWrappedFlag(bbw0, 2, m.IsSaleLocked)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteString(m.OwnerName); err != nil {
		return err
	}

	return nil
}

func (m *HouseInstanceInformations) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.SecondHand = getWrappedFlag(bbw0, 0)

	m.IsOnSale = getWrappedFlag(bbw0, 1)

	m.IsSaleLocked = getWrappedFlag(bbw0, 2)

	linstanceId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.InstanceId = linstanceId

	lownerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.OwnerName = lownerName

	return nil
}

type GameRolePlayNpcQuestFlag struct {
	QuestsToValidId []uint16

	QuestsToStartId []uint16
}

func (m *GameRolePlayNpcQuestFlag) ID() uint16 {
	return 384
}

func (m *GameRolePlayNpcQuestFlag) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.QuestsToValidId))); err != nil {
		return err
	}

	for i := range m.QuestsToValidId {

		if err := w.WriteVarUInt16(m.QuestsToValidId[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.QuestsToStartId))); err != nil {
		return err
	}

	for i := range m.QuestsToStartId {

		if err := w.WriteVarUInt16(m.QuestsToStartId[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameRolePlayNpcQuestFlag) Deserialize(r Reader) error {

	lquestsToValidIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.QuestsToValidId = make([]uint16, lquestsToValidIdLen)

	for i := range m.QuestsToValidId {

		lquestsToValidId, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.QuestsToValidId[i] = lquestsToValidId

	}

	lquestsToStartIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.QuestsToStartId = make([]uint16, lquestsToStartIdLen)

	for i := range m.QuestsToStartId {

		lquestsToStartId, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.QuestsToStartId[i] = lquestsToStartId

	}

	return nil
}

type DungeonPartyFinderPlayer struct {
	PlayerId int64

	PlayerName string

	Breed int8

	Sex bool

	Level uint8
}

func (m *DungeonPartyFinderPlayer) ID() uint16 {
	return 373
}

func (m *DungeonPartyFinderPlayer) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	if err := w.WriteString(m.PlayerName); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Sex); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	return nil
}

func (m *DungeonPartyFinderPlayer) Deserialize(r Reader) error {

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	lplayerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PlayerName = lplayerName

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	lsex, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Sex = lsex

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	return nil
}

type AbstractFightTeamInformations struct {
	TeamId uint8

	LeaderId float64

	TeamSide int8

	TeamTypeId uint8

	NbWaves uint8
}

func (m *AbstractFightTeamInformations) ID() uint16 {
	return 116
}

func (m *AbstractFightTeamInformations) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.TeamId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.LeaderId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.TeamSide); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.TeamTypeId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.NbWaves); err != nil {
		return err
	}

	return nil
}

func (m *AbstractFightTeamInformations) Deserialize(r Reader) error {

	lteamId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TeamId = lteamId

	lleaderId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.LeaderId = lleaderId

	lteamSide, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.TeamSide = lteamSide

	lteamTypeId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TeamTypeId = lteamTypeId

	lnbWaves, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NbWaves = lnbWaves

	return nil
}

type FightTeamInformations struct {
	AbstractFightTeamInformations

	TeamMembers []*FightTeamMemberInformations
}

func (m *FightTeamInformations) ID() uint16 {
	return 33
}

func (m *FightTeamInformations) Serialize(w Writer) error {

	if err := m.AbstractFightTeamInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.TeamMembers))); err != nil {
		return err
	}

	for i := range m.TeamMembers {

		if err := w.WriteUInt16(m.TeamMembers[i].ID()); err != nil {
			return err
		}

		if err := m.TeamMembers[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *FightTeamInformations) Deserialize(r Reader) error {

	if err := m.AbstractFightTeamInformations.Deserialize(r); err != nil {
		return err
	}

	lteamMembersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.TeamMembers = make([]*FightTeamMemberInformations, lteamMembersLen)

	for i := range m.TeamMembers {

		typeteamMembersID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lteamMembers, err := GetType(typeteamMembersID)
		if err != nil {
			return err
		}

		lteamMembers.Deserialize(r)

		m.TeamMembers[i] = lteamMembers.(*FightTeamMemberInformations)

	}

	return nil
}

type PartyInvitationMemberInformations struct {
	CharacterBaseInformations

	WorldX int16

	WorldY int16

	MapId int32

	SubAreaId uint16

	Companions []*PartyCompanionBaseInformations
}

func (m *PartyInvitationMemberInformations) ID() uint16 {
	return 376
}

func (m *PartyInvitationMemberInformations) Serialize(w Writer) error {

	if err := m.CharacterBaseInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Companions))); err != nil {
		return err
	}

	for i := range m.Companions {

		if err := m.Companions[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PartyInvitationMemberInformations) Deserialize(r Reader) error {

	if err := m.CharacterBaseInformations.Deserialize(r); err != nil {
		return err
	}

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lcompanionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Companions = make([]*PartyCompanionBaseInformations, lcompanionsLen)

	for i := range m.Companions {

		var lcompanions PartyCompanionBaseInformations

		lcompanions.Deserialize(r)

		m.Companions[i] = &lcompanions

	}

	return nil
}

type PartyMemberInformations struct {
	CharacterBaseInformations

	LifePoints uint32

	MaxLifePoints uint32

	Prospecting uint16

	RegenRate uint8

	Initiative uint16

	AlignmentSide int8

	WorldX int16

	WorldY int16

	MapId int32

	SubAreaId uint16

	Status *PlayerStatus

	Companions []*PartyCompanionMemberInformations
}

func (m *PartyMemberInformations) ID() uint16 {
	return 90
}

func (m *PartyMemberInformations) Serialize(w Writer) error {

	if err := m.CharacterBaseInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.LifePoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MaxLifePoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Prospecting); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.RegenRate); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Initiative); err != nil {
		return err
	}

	if err := w.WriteInt8(m.AlignmentSide); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Status.ID()); err != nil {
		return err
	}

	if err := m.Status.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Companions))); err != nil {
		return err
	}

	for i := range m.Companions {

		if err := m.Companions[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PartyMemberInformations) Deserialize(r Reader) error {

	if err := m.CharacterBaseInformations.Deserialize(r); err != nil {
		return err
	}

	llifePoints, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.LifePoints = llifePoints

	lmaxLifePoints, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MaxLifePoints = lmaxLifePoints

	lprospecting, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Prospecting = lprospecting

	lregenRate, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.RegenRate = lregenRate

	linitiative, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Initiative = linitiative

	lalignmentSide, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.AlignmentSide = lalignmentSide

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	typestatusID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lstatus, err := GetType(typestatusID)
	if err != nil {
		return err
	}

	lstatus.Deserialize(r)

	m.Status = lstatus.(*PlayerStatus)

	lcompanionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Companions = make([]*PartyCompanionMemberInformations, lcompanionsLen)

	for i := range m.Companions {

		var lcompanions PartyCompanionMemberInformations

		lcompanions.Deserialize(r)

		m.Companions[i] = &lcompanions

	}

	return nil
}

type PartyCompanionBaseInformations struct {
	IndexId uint8

	CompanionGenericId uint8

	EntityLook *EntityLook
}

func (m *PartyCompanionBaseInformations) ID() uint16 {
	return 453
}

func (m *PartyCompanionBaseInformations) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.IndexId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.CompanionGenericId); err != nil {
		return err
	}

	if err := m.EntityLook.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyCompanionBaseInformations) Deserialize(r Reader) error {

	lindexId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.IndexId = lindexId

	lcompanionGenericId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CompanionGenericId = lcompanionGenericId

	var lentityLook EntityLook

	lentityLook.Deserialize(r)

	m.EntityLook = &lentityLook

	return nil
}

type FightTeamMemberInformations struct {
	Id float64
}

func (m *FightTeamMemberInformations) ID() uint16 {
	return 44
}

func (m *FightTeamMemberInformations) Serialize(w Writer) error {

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *FightTeamMemberInformations) Deserialize(r Reader) error {

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type PartyGuestInformations struct {
	GuestId int64

	HostId int64

	Name string

	GuestLook *EntityLook

	Breed int8

	Sex bool

	Status *PlayerStatus

	Companions []*PartyCompanionBaseInformations
}

func (m *PartyGuestInformations) ID() uint16 {
	return 374
}

func (m *PartyGuestInformations) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.GuestId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.HostId); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := m.GuestLook.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Sex); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Status.ID()); err != nil {
		return err
	}

	if err := m.Status.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Companions))); err != nil {
		return err
	}

	for i := range m.Companions {

		if err := m.Companions[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PartyGuestInformations) Deserialize(r Reader) error {

	lguestId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.GuestId = lguestId

	lhostId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.HostId = lhostId

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	var lguestLook EntityLook

	lguestLook.Deserialize(r)

	m.GuestLook = &lguestLook

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	lsex, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Sex = lsex

	typestatusID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lstatus, err := GetType(typestatusID)
	if err != nil {
		return err
	}

	lstatus.Deserialize(r)

	m.Status = lstatus.(*PlayerStatus)

	lcompanionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Companions = make([]*PartyCompanionBaseInformations, lcompanionsLen)

	for i := range m.Companions {

		var lcompanions PartyCompanionBaseInformations

		lcompanions.Deserialize(r)

		m.Companions[i] = &lcompanions

	}

	return nil
}

type PartyCompanionMemberInformations struct {
	PartyCompanionBaseInformations

	Initiative uint16

	LifePoints uint32

	MaxLifePoints uint32

	Prospecting uint16

	RegenRate uint8
}

func (m *PartyCompanionMemberInformations) ID() uint16 {
	return 452
}

func (m *PartyCompanionMemberInformations) Serialize(w Writer) error {

	if err := m.PartyCompanionBaseInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Initiative); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.LifePoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MaxLifePoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Prospecting); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.RegenRate); err != nil {
		return err
	}

	return nil
}

func (m *PartyCompanionMemberInformations) Deserialize(r Reader) error {

	if err := m.PartyCompanionBaseInformations.Deserialize(r); err != nil {
		return err
	}

	linitiative, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Initiative = linitiative

	llifePoints, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.LifePoints = llifePoints

	lmaxLifePoints, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MaxLifePoints = lmaxLifePoints

	lprospecting, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Prospecting = lprospecting

	lregenRate, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.RegenRate = lregenRate

	return nil
}

type FightCommonInformations struct {
	FightId int32

	FightType uint8

	FightTeams []*FightTeamInformations

	FightTeamsPositions []uint16

	FightTeamsOptions []*FightOptionsInformations
}

func (m *FightCommonInformations) ID() uint16 {
	return 43
}

func (m *FightCommonInformations) Serialize(w Writer) error {

	if err := w.WriteInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.FightType); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.FightTeams))); err != nil {
		return err
	}

	for i := range m.FightTeams {

		if err := w.WriteUInt16(m.FightTeams[i].ID()); err != nil {
			return err
		}

		if err := m.FightTeams[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.FightTeamsPositions))); err != nil {
		return err
	}

	for i := range m.FightTeamsPositions {

		if err := w.WriteVarUInt16(m.FightTeamsPositions[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.FightTeamsOptions))); err != nil {
		return err
	}

	for i := range m.FightTeamsOptions {

		if err := m.FightTeamsOptions[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *FightCommonInformations) Deserialize(r Reader) error {

	lfightId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lfightType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.FightType = lfightType

	lfightTeamsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FightTeams = make([]*FightTeamInformations, lfightTeamsLen)

	for i := range m.FightTeams {

		typefightTeamsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lfightTeams, err := GetType(typefightTeamsID)
		if err != nil {
			return err
		}

		lfightTeams.Deserialize(r)

		m.FightTeams[i] = lfightTeams.(*FightTeamInformations)

	}

	lfightTeamsPositionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FightTeamsPositions = make([]uint16, lfightTeamsPositionsLen)

	for i := range m.FightTeamsPositions {

		lfightTeamsPositions, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.FightTeamsPositions[i] = lfightTeamsPositions

	}

	lfightTeamsOptionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FightTeamsOptions = make([]*FightOptionsInformations, lfightTeamsOptionsLen)

	for i := range m.FightTeamsOptions {

		var lfightTeamsOptions FightOptionsInformations

		lfightTeamsOptions.Deserialize(r)

		m.FightTeamsOptions[i] = &lfightTeamsOptions

	}

	return nil
}

type PartyMemberArenaInformations struct {
	PartyMemberInformations

	Rank uint16
}

func (m *PartyMemberArenaInformations) ID() uint16 {
	return 391
}

func (m *PartyMemberArenaInformations) Serialize(w Writer) error {

	if err := m.PartyMemberInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Rank); err != nil {
		return err
	}

	return nil
}

func (m *PartyMemberArenaInformations) Deserialize(r Reader) error {

	if err := m.PartyMemberInformations.Deserialize(r); err != nil {
		return err
	}

	lrank, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Rank = lrank

	return nil
}

type HouseInformationsForGuild struct {
	HouseInformations

	InstanceId uint32

	SecondHand bool

	OwnerName string

	WorldX int16

	WorldY int16

	MapId int32

	SubAreaId uint16

	SkillListIds []int32

	GuildshareParams uint32
}

func (m *HouseInformationsForGuild) ID() uint16 {
	return 170
}

func (m *HouseInformationsForGuild) Serialize(w Writer) error {

	if err := m.HouseInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.SecondHand); err != nil {
		return err
	}

	if err := w.WriteString(m.OwnerName); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.SkillListIds))); err != nil {
		return err
	}

	for i := range m.SkillListIds {

		if err := w.WriteInt32(m.SkillListIds[i]); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt32(m.GuildshareParams); err != nil {
		return err
	}

	return nil
}

func (m *HouseInformationsForGuild) Deserialize(r Reader) error {

	if err := m.HouseInformations.Deserialize(r); err != nil {
		return err
	}

	linstanceId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.InstanceId = linstanceId

	lsecondHand, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.SecondHand = lsecondHand

	lownerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.OwnerName = lownerName

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lskillListIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SkillListIds = make([]int32, lskillListIdsLen)

	for i := range m.SkillListIds {

		lskillListIds, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.SkillListIds[i] = lskillListIds

	}

	lguildshareParams, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.GuildshareParams = lguildshareParams

	return nil
}

type Item struct {
}

func (m *Item) ID() uint16 {
	return 7
}

func (m *Item) Serialize(w Writer) error {

	return nil
}

func (m *Item) Deserialize(r Reader) error {

	return nil
}

type ObjectItemGenericQuantity struct {
	Item

	ObjectGID uint16

	Quantity uint32
}

func (m *ObjectItemGenericQuantity) ID() uint16 {
	return 483
}

func (m *ObjectItemGenericQuantity) Serialize(w Writer) error {

	if err := m.Item.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGID); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ObjectItemGenericQuantity) Deserialize(r Reader) error {

	if err := m.Item.Deserialize(r); err != nil {
		return err
	}

	lobjectGID, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGID = lobjectGID

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type AbstractContactInformations struct {
	AccountId uint32

	AccountName string
}

func (m *AbstractContactInformations) ID() uint16 {
	return 380
}

func (m *AbstractContactInformations) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	if err := w.WriteString(m.AccountName); err != nil {
		return err
	}

	return nil
}

func (m *AbstractContactInformations) Deserialize(r Reader) error {

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	laccountName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.AccountName = laccountName

	return nil
}

type IgnoredInformations struct {
	AbstractContactInformations
}

func (m *IgnoredInformations) ID() uint16 {
	return 106
}

func (m *IgnoredInformations) Serialize(w Writer) error {

	if err := m.AbstractContactInformations.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *IgnoredInformations) Deserialize(r Reader) error {

	if err := m.AbstractContactInformations.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type IgnoredOnlineInformations struct {
	IgnoredInformations

	PlayerId int64

	PlayerName string

	Breed int8

	Sex bool
}

func (m *IgnoredOnlineInformations) ID() uint16 {
	return 105
}

func (m *IgnoredOnlineInformations) Serialize(w Writer) error {

	if err := m.IgnoredInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	if err := w.WriteString(m.PlayerName); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Sex); err != nil {
		return err
	}

	return nil
}

func (m *IgnoredOnlineInformations) Deserialize(r Reader) error {

	if err := m.IgnoredInformations.Deserialize(r); err != nil {
		return err
	}

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	lplayerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PlayerName = lplayerName

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	lsex, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Sex = lsex

	return nil
}

type FriendInformations struct {
	AbstractContactInformations

	PlayerState uint8

	LastConnection uint16

	AchievementPoints int32
}

func (m *FriendInformations) ID() uint16 {
	return 78
}

func (m *FriendInformations) Serialize(w Writer) error {

	if err := m.AbstractContactInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.PlayerState); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.LastConnection); err != nil {
		return err
	}

	if err := w.WriteInt32(m.AchievementPoints); err != nil {
		return err
	}

	return nil
}

func (m *FriendInformations) Deserialize(r Reader) error {

	if err := m.AbstractContactInformations.Deserialize(r); err != nil {
		return err
	}

	lplayerState, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PlayerState = lplayerState

	llastConnection, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.LastConnection = llastConnection

	lachievementPoints, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.AchievementPoints = lachievementPoints

	return nil
}

type FriendOnlineInformations struct {
	FriendInformations

	PlayerId int64

	PlayerName string

	Level uint8

	AlignmentSide int8

	Breed int8

	Sex bool

	GuildInfo *GuildInformations

	MoodSmileyId uint16

	Status *PlayerStatus

	HavenBagShared bool
}

func (m *FriendOnlineInformations) ID() uint16 {
	return 92
}

func (m *FriendOnlineInformations) Serialize(w Writer) error {

	if err := m.FriendInformations.Serialize(w); err != nil {
		return err
	}

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Sex)

	setWrappedFlag(bbw0, 1, m.HavenBagShared)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	if err := w.WriteString(m.PlayerName); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	if err := w.WriteInt8(m.AlignmentSide); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := m.GuildInfo.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MoodSmileyId); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Status.ID()); err != nil {
		return err
	}

	if err := m.Status.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *FriendOnlineInformations) Deserialize(r Reader) error {

	if err := m.FriendInformations.Deserialize(r); err != nil {
		return err
	}

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Sex = getWrappedFlag(bbw0, 0)

	m.HavenBagShared = getWrappedFlag(bbw0, 1)

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	lplayerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PlayerName = lplayerName

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	lalignmentSide, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.AlignmentSide = lalignmentSide

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	var lguildInfo GuildInformations

	lguildInfo.Deserialize(r)

	m.GuildInfo = &lguildInfo

	lmoodSmileyId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MoodSmileyId = lmoodSmileyId

	typestatusID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lstatus, err := GetType(typestatusID)
	if err != nil {
		return err
	}

	lstatus.Deserialize(r)

	m.Status = lstatus.(*PlayerStatus)

	return nil
}

type PaddockContentInformations struct {
	PaddockInformations

	PaddockId int32

	WorldX int16

	WorldY int16

	MapId int32

	SubAreaId uint16

	Abandonned bool

	MountsInformations []*MountInformationsForPaddock
}

func (m *PaddockContentInformations) ID() uint16 {
	return 183
}

func (m *PaddockContentInformations) Serialize(w Writer) error {

	if err := m.PaddockInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.PaddockId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Abandonned); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.MountsInformations))); err != nil {
		return err
	}

	for i := range m.MountsInformations {

		if err := m.MountsInformations[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PaddockContentInformations) Deserialize(r Reader) error {

	if err := m.PaddockInformations.Deserialize(r); err != nil {
		return err
	}

	lpaddockId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.PaddockId = lpaddockId

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	labandonned, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Abandonned = labandonned

	lmountsInformationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MountsInformations = make([]*MountInformationsForPaddock, lmountsInformationsLen)

	for i := range m.MountsInformations {

		var lmountsInformations MountInformationsForPaddock

		lmountsInformations.Deserialize(r)

		m.MountsInformations[i] = &lmountsInformations

	}

	return nil
}

type TaxCollectorMovement struct {
	MovementType uint8

	BasicInfos *TaxCollectorBasicInformations

	PlayerId int64

	PlayerName string
}

func (m *TaxCollectorMovement) ID() uint16 {
	return 493
}

func (m *TaxCollectorMovement) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.MovementType); err != nil {
		return err
	}

	if err := m.BasicInfos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	if err := w.WriteString(m.PlayerName); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorMovement) Deserialize(r Reader) error {

	lmovementType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MovementType = lmovementType

	var lbasicInfos TaxCollectorBasicInformations

	lbasicInfos.Deserialize(r)

	m.BasicInfos = &lbasicInfos

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	lplayerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PlayerName = lplayerName

	return nil
}

type TaxCollectorInformations struct {
	UniqueId int32

	FirtNameId uint16

	LastNameId uint16

	AdditionalInfos *AdditionalTaxCollectorInformations

	WorldX int16

	WorldY int16

	SubAreaId uint16

	State uint8

	Look *EntityLook

	Complements []*TaxCollectorComplementaryInformations
}

func (m *TaxCollectorInformations) ID() uint16 {
	return 167
}

func (m *TaxCollectorInformations) Serialize(w Writer) error {

	if err := w.WriteInt32(m.UniqueId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.FirtNameId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.LastNameId); err != nil {
		return err
	}

	if err := m.AdditionalInfos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.State); err != nil {
		return err
	}

	if err := m.Look.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Complements))); err != nil {
		return err
	}

	for i := range m.Complements {

		if err := w.WriteUInt16(m.Complements[i].ID()); err != nil {
			return err
		}

		if err := m.Complements[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *TaxCollectorInformations) Deserialize(r Reader) error {

	luniqueId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.UniqueId = luniqueId

	lfirtNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.FirtNameId = lfirtNameId

	llastNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.LastNameId = llastNameId

	var ladditionalInfos AdditionalTaxCollectorInformations

	ladditionalInfos.Deserialize(r)

	m.AdditionalInfos = &ladditionalInfos

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lstate, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.State = lstate

	var llook EntityLook

	llook.Deserialize(r)

	m.Look = &llook

	lcomplementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Complements = make([]*TaxCollectorComplementaryInformations, lcomplementsLen)

	for i := range m.Complements {

		typecomplementsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lcomplements, err := GetType(typecomplementsID)
		if err != nil {
			return err
		}

		lcomplements.Deserialize(r)

		m.Complements[i] = lcomplements.(*TaxCollectorComplementaryInformations)

	}

	return nil
}

type PlayerStatusExtended struct {
	PlayerStatus

	Message string
}

func (m *PlayerStatusExtended) ID() uint16 {
	return 414
}

func (m *PlayerStatusExtended) Serialize(w Writer) error {

	if err := m.PlayerStatus.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Message); err != nil {
		return err
	}

	return nil
}

func (m *PlayerStatusExtended) Deserialize(r Reader) error {

	if err := m.PlayerStatus.Deserialize(r); err != nil {
		return err
	}

	lmessage, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Message = lmessage

	return nil
}

type AbstractSocialGroupInfos struct {
}

func (m *AbstractSocialGroupInfos) ID() uint16 {
	return 416
}

func (m *AbstractSocialGroupInfos) Serialize(w Writer) error {

	return nil
}

func (m *AbstractSocialGroupInfos) Deserialize(r Reader) error {

	return nil
}

type BasicGuildInformations struct {
	AbstractSocialGroupInfos

	GuildId uint32

	GuildName string

	GuildLevel uint8
}

func (m *BasicGuildInformations) ID() uint16 {
	return 365
}

func (m *BasicGuildInformations) Serialize(w Writer) error {

	if err := m.AbstractSocialGroupInfos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.GuildId); err != nil {
		return err
	}

	if err := w.WriteString(m.GuildName); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.GuildLevel); err != nil {
		return err
	}

	return nil
}

func (m *BasicGuildInformations) Deserialize(r Reader) error {

	if err := m.AbstractSocialGroupInfos.Deserialize(r); err != nil {
		return err
	}

	lguildId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.GuildId = lguildId

	lguildName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.GuildName = lguildName

	lguildLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.GuildLevel = lguildLevel

	return nil
}

type GuildInformations struct {
	BasicGuildInformations

	GuildEmblem *GuildEmblem
}

func (m *GuildInformations) ID() uint16 {
	return 127
}

func (m *GuildInformations) Serialize(w Writer) error {

	if err := m.BasicGuildInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.GuildEmblem.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildInformations) Deserialize(r Reader) error {

	if err := m.BasicGuildInformations.Deserialize(r); err != nil {
		return err
	}

	var lguildEmblem GuildEmblem

	lguildEmblem.Deserialize(r)

	m.GuildEmblem = &lguildEmblem

	return nil
}

type GuildFactSheetInformations struct {
	GuildInformations

	LeaderId int64

	NbMembers uint16
}

func (m *GuildFactSheetInformations) ID() uint16 {
	return 424
}

func (m *GuildFactSheetInformations) Serialize(w Writer) error {

	if err := m.GuildInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.LeaderId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NbMembers); err != nil {
		return err
	}

	return nil
}

func (m *GuildFactSheetInformations) Deserialize(r Reader) error {

	if err := m.GuildInformations.Deserialize(r); err != nil {
		return err
	}

	lleaderId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.LeaderId = lleaderId

	lnbMembers, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NbMembers = lnbMembers

	return nil
}

type GuildInsiderFactSheetInformations struct {
	GuildFactSheetInformations

	LeaderName string

	NbConnectedMembers uint16

	NbTaxCollectors uint8

	LastActivity uint32
}

func (m *GuildInsiderFactSheetInformations) ID() uint16 {
	return 423
}

func (m *GuildInsiderFactSheetInformations) Serialize(w Writer) error {

	if err := m.GuildFactSheetInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.LeaderName); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NbConnectedMembers); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.NbTaxCollectors); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.LastActivity); err != nil {
		return err
	}

	return nil
}

func (m *GuildInsiderFactSheetInformations) Deserialize(r Reader) error {

	if err := m.GuildFactSheetInformations.Deserialize(r); err != nil {
		return err
	}

	lleaderName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.LeaderName = lleaderName

	lnbConnectedMembers, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NbConnectedMembers = lnbConnectedMembers

	lnbTaxCollectors, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NbTaxCollectors = lnbTaxCollectors

	llastActivity, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.LastActivity = llastActivity

	return nil
}

type PrismSubareaEmptyInfo struct {
	SubAreaId uint16

	AllianceId uint32
}

func (m *PrismSubareaEmptyInfo) ID() uint16 {
	return 438
}

func (m *PrismSubareaEmptyInfo) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.AllianceId); err != nil {
		return err
	}

	return nil
}

func (m *PrismSubareaEmptyInfo) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lallianceId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.AllianceId = lallianceId

	return nil
}

type PrismGeolocalizedInformation struct {
	PrismSubareaEmptyInfo

	WorldX int16

	WorldY int16

	MapId int32

	Prism *PrismInformation
}

func (m *PrismGeolocalizedInformation) ID() uint16 {
	return 434
}

func (m *PrismGeolocalizedInformation) Serialize(w Writer) error {

	if err := m.PrismSubareaEmptyInfo.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Prism.ID()); err != nil {
		return err
	}

	if err := m.Prism.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PrismGeolocalizedInformation) Deserialize(r Reader) error {

	if err := m.PrismSubareaEmptyInfo.Deserialize(r); err != nil {
		return err
	}

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	typeprismID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lprism, err := GetType(typeprismID)
	if err != nil {
		return err
	}

	lprism.Deserialize(r)

	m.Prism = lprism.(*PrismInformation)

	return nil
}

type GuildInAllianceInformations struct {
	GuildInformations

	NbMembers uint8
}

func (m *GuildInAllianceInformations) ID() uint16 {
	return 420
}

func (m *GuildInAllianceInformations) Serialize(w Writer) error {

	if err := m.GuildInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.NbMembers); err != nil {
		return err
	}

	return nil
}

func (m *GuildInAllianceInformations) Deserialize(r Reader) error {

	if err := m.GuildInformations.Deserialize(r); err != nil {
		return err
	}

	lnbMembers, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NbMembers = lnbMembers

	return nil
}

type ObjectItemGenericQuantityPrice struct {
	ObjectItemGenericQuantity

	Price int64
}

func (m *ObjectItemGenericQuantityPrice) ID() uint16 {
	return 494
}

func (m *ObjectItemGenericQuantityPrice) Serialize(w Writer) error {

	if err := m.ObjectItemGenericQuantity.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Price); err != nil {
		return err
	}

	return nil
}

func (m *ObjectItemGenericQuantityPrice) Deserialize(r Reader) error {

	if err := m.ObjectItemGenericQuantity.Deserialize(r); err != nil {
		return err
	}

	lprice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Price = lprice

	return nil
}

type ObjectItem struct {
	Item

	Position uint8

	ObjectGID uint16

	Effects []*ObjectEffect

	ObjectUID uint32

	Quantity uint32
}

func (m *ObjectItem) ID() uint16 {
	return 37
}

func (m *ObjectItem) Serialize(w Writer) error {

	if err := m.Item.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Position); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGID); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Effects))); err != nil {
		return err
	}

	for i := range m.Effects {

		if err := w.WriteUInt16(m.Effects[i].ID()); err != nil {
			return err
		}

		if err := m.Effects[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ObjectItem) Deserialize(r Reader) error {

	if err := m.Item.Deserialize(r); err != nil {
		return err
	}

	lposition, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Position = lposition

	lobjectGID, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGID = lobjectGID

	leffectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Effects = make([]*ObjectEffect, leffectsLen)

	for i := range m.Effects {

		typeeffectsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		leffects, err := GetType(typeeffectsID)
		if err != nil {
			return err
		}

		leffects.Deserialize(r)

		m.Effects[i] = leffects.(*ObjectEffect)

	}

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type BasicAllianceInformations struct {
	AbstractSocialGroupInfos

	AllianceId uint32

	AllianceTag string
}

func (m *BasicAllianceInformations) ID() uint16 {
	return 419
}

func (m *BasicAllianceInformations) Serialize(w Writer) error {

	if err := m.AbstractSocialGroupInfos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.AllianceId); err != nil {
		return err
	}

	if err := w.WriteString(m.AllianceTag); err != nil {
		return err
	}

	return nil
}

func (m *BasicAllianceInformations) Deserialize(r Reader) error {

	if err := m.AbstractSocialGroupInfos.Deserialize(r); err != nil {
		return err
	}

	lallianceId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.AllianceId = lallianceId

	lallianceTag, err := r.ReadString()
	if err != nil {
		return err
	}

	m.AllianceTag = lallianceTag

	return nil
}

type FightEntityDispositionInformations struct {
	EntityDispositionInformations

	CarryingCharacterId float64
}

func (m *FightEntityDispositionInformations) ID() uint16 {
	return 217
}

func (m *FightEntityDispositionInformations) Serialize(w Writer) error {

	if err := m.EntityDispositionInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.CarryingCharacterId); err != nil {
		return err
	}

	return nil
}

func (m *FightEntityDispositionInformations) Deserialize(r Reader) error {

	if err := m.EntityDispositionInformations.Deserialize(r); err != nil {
		return err
	}

	lcarryingCharacterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.CarryingCharacterId = lcarryingCharacterId

	return nil
}

type StatedElement struct {
	ElementId uint32

	ElementCellId uint16

	ElementState uint32

	OnCurrentMap bool
}

func (m *StatedElement) ID() uint16 {
	return 108
}

func (m *StatedElement) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.ElementId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ElementCellId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ElementState); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.OnCurrentMap); err != nil {
		return err
	}

	return nil
}

func (m *StatedElement) Deserialize(r Reader) error {

	lelementId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.ElementId = lelementId

	lelementCellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ElementCellId = lelementCellId

	lelementState, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ElementState = lelementState

	lonCurrentMap, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.OnCurrentMap = lonCurrentMap

	return nil
}

type GameFightMinimalStatsPreparation struct {
	GameFightMinimalStats

	Initiative uint32
}

func (m *GameFightMinimalStatsPreparation) ID() uint16 {
	return 360
}

func (m *GameFightMinimalStatsPreparation) Serialize(w Writer) error {

	if err := m.GameFightMinimalStats.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Initiative); err != nil {
		return err
	}

	return nil
}

func (m *GameFightMinimalStatsPreparation) Deserialize(r Reader) error {

	if err := m.GameFightMinimalStats.Deserialize(r); err != nil {
		return err
	}

	linitiative, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Initiative = linitiative

	return nil
}

type MapObstacle struct {
	ObstacleCellId uint16

	State uint8
}

func (m *MapObstacle) ID() uint16 {
	return 200
}

func (m *MapObstacle) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ObstacleCellId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.State); err != nil {
		return err
	}

	return nil
}

func (m *MapObstacle) Deserialize(r Reader) error {

	lobstacleCellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObstacleCellId = lobstacleCellId

	lstate, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.State = lstate

	return nil
}

type IdentifiedEntityDispositionInformations struct {
	EntityDispositionInformations

	Id float64
}

func (m *IdentifiedEntityDispositionInformations) ID() uint16 {
	return 107
}

func (m *IdentifiedEntityDispositionInformations) Serialize(w Writer) error {

	if err := m.EntityDispositionInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *IdentifiedEntityDispositionInformations) Deserialize(r Reader) error {

	if err := m.EntityDispositionInformations.Deserialize(r); err != nil {
		return err
	}

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type ObjectItemInRolePlay struct {
	CellId uint16

	ObjectGID uint16
}

func (m *ObjectItemInRolePlay) ID() uint16 {
	return 198
}

func (m *ObjectItemInRolePlay) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGID); err != nil {
		return err
	}

	return nil
}

func (m *ObjectItemInRolePlay) Deserialize(r Reader) error {

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	lobjectGID, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGID = lobjectGID

	return nil
}

type PaddockItem struct {
	ObjectItemInRolePlay

	Durability *ItemDurability
}

func (m *PaddockItem) ID() uint16 {
	return 185
}

func (m *PaddockItem) Serialize(w Writer) error {

	if err := m.ObjectItemInRolePlay.Serialize(w); err != nil {
		return err
	}

	if err := m.Durability.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PaddockItem) Deserialize(r Reader) error {

	if err := m.ObjectItemInRolePlay.Deserialize(r); err != nil {
		return err
	}

	var ldurability ItemDurability

	ldurability.Deserialize(r)

	m.Durability = &ldurability

	return nil
}

type HumanOptionAlliance struct {
	HumanOption

	AllianceInformations *AllianceInformations

	Aggressable uint8
}

func (m *HumanOptionAlliance) ID() uint16 {
	return 425
}

func (m *HumanOptionAlliance) Serialize(w Writer) error {

	if err := m.HumanOption.Serialize(w); err != nil {
		return err
	}

	if err := m.AllianceInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Aggressable); err != nil {
		return err
	}

	return nil
}

func (m *HumanOptionAlliance) Deserialize(r Reader) error {

	if err := m.HumanOption.Deserialize(r); err != nil {
		return err
	}

	var lallianceInformations AllianceInformations

	lallianceInformations.Deserialize(r)

	m.AllianceInformations = &lallianceInformations

	laggressable, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Aggressable = laggressable

	return nil
}

type IndexedEntityLook struct {
	Look *EntityLook

	Index uint8
}

func (m *IndexedEntityLook) ID() uint16 {
	return 405
}

func (m *IndexedEntityLook) Serialize(w Writer) error {

	if err := m.Look.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Index); err != nil {
		return err
	}

	return nil
}

func (m *IndexedEntityLook) Deserialize(r Reader) error {

	var llook EntityLook

	llook.Deserialize(r)

	m.Look = &llook

	lindex, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Index = lindex

	return nil
}

type HumanOptionSkillUse struct {
	HumanOption

	ElementId uint32

	SkillId uint16

	SkillEndTime float64
}

func (m *HumanOptionSkillUse) ID() uint16 {
	return 495
}

func (m *HumanOptionSkillUse) Serialize(w Writer) error {

	if err := m.HumanOption.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ElementId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SkillId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.SkillEndTime); err != nil {
		return err
	}

	return nil
}

func (m *HumanOptionSkillUse) Deserialize(r Reader) error {

	if err := m.HumanOption.Deserialize(r); err != nil {
		return err
	}

	lelementId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ElementId = lelementId

	lskillId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SkillId = lskillId

	lskillEndTime, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.SkillEndTime = lskillEndTime

	return nil
}

type HumanOptionEmote struct {
	HumanOption

	EmoteId uint8

	EmoteStartTime float64
}

func (m *HumanOptionEmote) ID() uint16 {
	return 407
}

func (m *HumanOptionEmote) Serialize(w Writer) error {

	if err := m.HumanOption.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.EmoteId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.EmoteStartTime); err != nil {
		return err
	}

	return nil
}

func (m *HumanOptionEmote) Deserialize(r Reader) error {

	if err := m.HumanOption.Deserialize(r); err != nil {
		return err
	}

	lemoteId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.EmoteId = lemoteId

	lemoteStartTime, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.EmoteStartTime = lemoteStartTime

	return nil
}

type HumanOptionFollowers struct {
	HumanOption

	FollowingCharactersLook []*IndexedEntityLook
}

func (m *HumanOptionFollowers) ID() uint16 {
	return 410
}

func (m *HumanOptionFollowers) Serialize(w Writer) error {

	if err := m.HumanOption.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.FollowingCharactersLook))); err != nil {
		return err
	}

	for i := range m.FollowingCharactersLook {

		if err := m.FollowingCharactersLook[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *HumanOptionFollowers) Deserialize(r Reader) error {

	if err := m.HumanOption.Deserialize(r); err != nil {
		return err
	}

	lfollowingCharactersLookLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FollowingCharactersLook = make([]*IndexedEntityLook, lfollowingCharactersLookLen)

	for i := range m.FollowingCharactersLook {

		var lfollowingCharactersLook IndexedEntityLook

		lfollowingCharactersLook.Deserialize(r)

		m.FollowingCharactersLook[i] = &lfollowingCharactersLook

	}

	return nil
}

type ActorOrientation struct {
	Id float64

	Direction uint8
}

func (m *ActorOrientation) ID() uint16 {
	return 353
}

func (m *ActorOrientation) Serialize(w Writer) error {

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *ActorOrientation) Deserialize(r Reader) error {

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	ldirection, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Direction = ldirection

	return nil
}

type HumanOptionObjectUse struct {
	HumanOption

	DelayTypeId uint8

	DelayEndTime float64

	ObjectGID uint16
}

func (m *HumanOptionObjectUse) ID() uint16 {
	return 449
}

func (m *HumanOptionObjectUse) Serialize(w Writer) error {

	if err := m.HumanOption.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.DelayTypeId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.DelayEndTime); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGID); err != nil {
		return err
	}

	return nil
}

func (m *HumanOptionObjectUse) Deserialize(r Reader) error {

	if err := m.HumanOption.Deserialize(r); err != nil {
		return err
	}

	ldelayTypeId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.DelayTypeId = ldelayTypeId

	ldelayEndTime, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DelayEndTime = ldelayEndTime

	lobjectGID, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGID = lobjectGID

	return nil
}

type TrustCertificate struct {
	Id uint32

	Hash string
}

func (m *TrustCertificate) ID() uint16 {
	return 377
}

func (m *TrustCertificate) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.Id); err != nil {
		return err
	}

	if err := w.WriteString(m.Hash); err != nil {
		return err
	}

	return nil
}

func (m *TrustCertificate) Deserialize(r Reader) error {

	lid, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.Id = lid

	lhash, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Hash = lhash

	return nil
}

type AbstractFightDispellableEffect struct {
	Uid uint32

	TargetId float64

	TurnDuration int16

	Dispelable uint8

	SpellId uint16

	EffectId uint32

	ParentBoostUid uint32
}

func (m *AbstractFightDispellableEffect) ID() uint16 {
	return 206
}

func (m *AbstractFightDispellableEffect) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.Uid); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.TurnDuration); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Dispelable); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.EffectId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ParentBoostUid); err != nil {
		return err
	}

	return nil
}

func (m *AbstractFightDispellableEffect) Deserialize(r Reader) error {

	luid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Uid = luid

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lturnDuration, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.TurnDuration = lturnDuration

	ldispelable, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Dispelable = ldispelable

	lspellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	leffectId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.EffectId = leffectId

	lparentBoostUid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ParentBoostUid = lparentBoostUid

	return nil
}

type FightTemporaryBoostEffect struct {
	AbstractFightDispellableEffect

	Delta int16
}

func (m *FightTemporaryBoostEffect) ID() uint16 {
	return 209
}

func (m *FightTemporaryBoostEffect) Serialize(w Writer) error {

	if err := m.AbstractFightDispellableEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.Delta); err != nil {
		return err
	}

	return nil
}

func (m *FightTemporaryBoostEffect) Deserialize(r Reader) error {

	if err := m.AbstractFightDispellableEffect.Deserialize(r); err != nil {
		return err
	}

	ldelta, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Delta = ldelta

	return nil
}

type GameFightSpellCooldown struct {
	SpellId int32

	Cooldown uint8
}

func (m *GameFightSpellCooldown) ID() uint16 {
	return 205
}

func (m *GameFightSpellCooldown) Serialize(w Writer) error {

	if err := w.WriteInt32(m.SpellId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Cooldown); err != nil {
		return err
	}

	return nil
}

func (m *GameFightSpellCooldown) Deserialize(r Reader) error {

	lspellId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	lcooldown, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Cooldown = lcooldown

	return nil
}

type GameActionMarkedCell struct {
	CellId uint16

	ZoneSize int8

	CellColor int32

	CellsType int8
}

func (m *GameActionMarkedCell) ID() uint16 {
	return 85
}

func (m *GameActionMarkedCell) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.ZoneSize); err != nil {
		return err
	}

	if err := w.WriteInt32(m.CellColor); err != nil {
		return err
	}

	if err := w.WriteInt8(m.CellsType); err != nil {
		return err
	}

	return nil
}

func (m *GameActionMarkedCell) Deserialize(r Reader) error {

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	lzoneSize, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.ZoneSize = lzoneSize

	lcellColor, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.CellColor = lcellColor

	lcellsType, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.CellsType = lcellsType

	return nil
}

type FightResultListEntry struct {
	Outcome uint16

	Wave uint8

	Rewards *FightLoot
}

func (m *FightResultListEntry) ID() uint16 {
	return 16
}

func (m *FightResultListEntry) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Outcome); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Wave); err != nil {
		return err
	}

	if err := m.Rewards.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *FightResultListEntry) Deserialize(r Reader) error {

	loutcome, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Outcome = loutcome

	lwave, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Wave = lwave

	var lrewards FightLoot

	lrewards.Deserialize(r)

	m.Rewards = &lrewards

	return nil
}

type FightResultFighterListEntry struct {
	FightResultListEntry

	Id float64

	Alive bool
}

func (m *FightResultFighterListEntry) ID() uint16 {
	return 189
}

func (m *FightResultFighterListEntry) Serialize(w Writer) error {

	if err := m.FightResultListEntry.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Alive); err != nil {
		return err
	}

	return nil
}

func (m *FightResultFighterListEntry) Deserialize(r Reader) error {

	if err := m.FightResultListEntry.Deserialize(r); err != nil {
		return err
	}

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	lalive, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Alive = lalive

	return nil
}

type FightDispellableEffectExtendedInformations struct {
	ActionId uint16

	SourceId float64

	Effect *AbstractFightDispellableEffect
}

func (m *FightDispellableEffectExtendedInformations) ID() uint16 {
	return 208
}

func (m *FightDispellableEffectExtendedInformations) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ActionId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.SourceId); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Effect.ID()); err != nil {
		return err
	}

	if err := m.Effect.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *FightDispellableEffectExtendedInformations) Deserialize(r Reader) error {

	lactionId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ActionId = lactionId

	lsourceId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.SourceId = lsourceId

	typeeffectID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	leffect, err := GetType(typeeffectID)
	if err != nil {
		return err
	}

	leffect.Deserialize(r)

	m.Effect = leffect.(*AbstractFightDispellableEffect)

	return nil
}

type FightResultPlayerListEntry struct {
	FightResultFighterListEntry

	Level uint8

	Additional []*FightResultAdditionalData
}

func (m *FightResultPlayerListEntry) ID() uint16 {
	return 24
}

func (m *FightResultPlayerListEntry) Serialize(w Writer) error {

	if err := m.FightResultFighterListEntry.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Additional))); err != nil {
		return err
	}

	for i := range m.Additional {

		if err := w.WriteUInt16(m.Additional[i].ID()); err != nil {
			return err
		}

		if err := m.Additional[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *FightResultPlayerListEntry) Deserialize(r Reader) error {

	if err := m.FightResultFighterListEntry.Deserialize(r); err != nil {
		return err
	}

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	ladditionalLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Additional = make([]*FightResultAdditionalData, ladditionalLen)

	for i := range m.Additional {

		typeadditionalID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		ladditional, err := GetType(typeadditionalID)
		if err != nil {
			return err
		}

		ladditional.Deserialize(r)

		m.Additional[i] = ladditional.(*FightResultAdditionalData)

	}

	return nil
}

type NamedPartyTeam struct {
	TeamId uint8

	PartyName string
}

func (m *NamedPartyTeam) ID() uint16 {
	return 469
}

func (m *NamedPartyTeam) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.TeamId); err != nil {
		return err
	}

	if err := w.WriteString(m.PartyName); err != nil {
		return err
	}

	return nil
}

func (m *NamedPartyTeam) Deserialize(r Reader) error {

	lteamId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TeamId = lteamId

	lpartyName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PartyName = lpartyName

	return nil
}

type Idol struct {
	Id uint16

	XpBonusPercent uint16

	DropBonusPercent uint16
}

func (m *Idol) ID() uint16 {
	return 489
}

func (m *Idol) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Id); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.XpBonusPercent); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DropBonusPercent); err != nil {
		return err
	}

	return nil
}

func (m *Idol) Deserialize(r Reader) error {

	lid, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Id = lid

	lxpBonusPercent, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.XpBonusPercent = lxpBonusPercent

	ldropBonusPercent, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DropBonusPercent = ldropBonusPercent

	return nil
}

type GameActionMark struct {
	MarkAuthorId float64

	MarkTeamId uint8

	MarkSpellId uint32

	MarkSpellLevel int16

	MarkId int16

	MarkType int8

	MarkimpactCell int16

	Cells []*GameActionMarkedCell

	Active bool
}

func (m *GameActionMark) ID() uint16 {
	return 351
}

func (m *GameActionMark) Serialize(w Writer) error {

	if err := w.WriteDouble(m.MarkAuthorId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.MarkTeamId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.MarkSpellId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.MarkSpellLevel); err != nil {
		return err
	}

	if err := w.WriteInt16(m.MarkId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.MarkType); err != nil {
		return err
	}

	if err := w.WriteInt16(m.MarkimpactCell); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Cells))); err != nil {
		return err
	}

	for i := range m.Cells {

		if err := m.Cells[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteBoolean(m.Active); err != nil {
		return err
	}

	return nil
}

func (m *GameActionMark) Deserialize(r Reader) error {

	lmarkAuthorId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.MarkAuthorId = lmarkAuthorId

	lmarkTeamId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MarkTeamId = lmarkTeamId

	lmarkSpellId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MarkSpellId = lmarkSpellId

	lmarkSpellLevel, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MarkSpellLevel = lmarkSpellLevel

	lmarkId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MarkId = lmarkId

	lmarkType, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.MarkType = lmarkType

	lmarkimpactCell, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MarkimpactCell = lmarkimpactCell

	lcellsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Cells = make([]*GameActionMarkedCell, lcellsLen)

	for i := range m.Cells {

		var lcells GameActionMarkedCell

		lcells.Deserialize(r)

		m.Cells[i] = &lcells

	}

	lactive, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Active = lactive

	return nil
}

type NamedPartyTeamWithOutcome struct {
	Team *NamedPartyTeam

	Outcome uint16
}

func (m *NamedPartyTeamWithOutcome) ID() uint16 {
	return 470
}

func (m *NamedPartyTeamWithOutcome) Serialize(w Writer) error {

	if err := m.Team.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Outcome); err != nil {
		return err
	}

	return nil
}

func (m *NamedPartyTeamWithOutcome) Deserialize(r Reader) error {

	var lteam NamedPartyTeam

	lteam.Deserialize(r)

	m.Team = &lteam

	loutcome, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Outcome = loutcome

	return nil
}

type GameFightMutantInformations struct {
	GameFightFighterNamedInformations

	PowerLevel uint8
}

func (m *GameFightMutantInformations) ID() uint16 {
	return 50
}

func (m *GameFightMutantInformations) Serialize(w Writer) error {

	if err := m.GameFightFighterNamedInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.PowerLevel); err != nil {
		return err
	}

	return nil
}

func (m *GameFightMutantInformations) Deserialize(r Reader) error {

	if err := m.GameFightFighterNamedInformations.Deserialize(r); err != nil {
		return err
	}

	lpowerLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PowerLevel = lpowerLevel

	return nil
}

type GameFightResumeSlaveInfo struct {
	SlaveId float64

	SpellCooldowns []*GameFightSpellCooldown

	SummonCount uint8

	BombCount uint8
}

func (m *GameFightResumeSlaveInfo) ID() uint16 {
	return 364
}

func (m *GameFightResumeSlaveInfo) Serialize(w Writer) error {

	if err := w.WriteDouble(m.SlaveId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.SpellCooldowns))); err != nil {
		return err
	}

	for i := range m.SpellCooldowns {

		if err := m.SpellCooldowns[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteUInt8(m.SummonCount); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.BombCount); err != nil {
		return err
	}

	return nil
}

func (m *GameFightResumeSlaveInfo) Deserialize(r Reader) error {

	lslaveId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.SlaveId = lslaveId

	lspellCooldownsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SpellCooldowns = make([]*GameFightSpellCooldown, lspellCooldownsLen)

	for i := range m.SpellCooldowns {

		var lspellCooldowns GameFightSpellCooldown

		lspellCooldowns.Deserialize(r)

		m.SpellCooldowns[i] = &lspellCooldowns

	}

	lsummonCount, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.SummonCount = lsummonCount

	lbombCount, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BombCount = lbombCount

	return nil
}

type FightResultTaxCollectorListEntry struct {
	FightResultFighterListEntry

	Level uint8

	GuildInfo *BasicGuildInformations

	ExperienceForGuild int32
}

func (m *FightResultTaxCollectorListEntry) ID() uint16 {
	return 84
}

func (m *FightResultTaxCollectorListEntry) Serialize(w Writer) error {

	if err := m.FightResultFighterListEntry.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	if err := m.GuildInfo.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.ExperienceForGuild); err != nil {
		return err
	}

	return nil
}

func (m *FightResultTaxCollectorListEntry) Deserialize(r Reader) error {

	if err := m.FightResultFighterListEntry.Deserialize(r); err != nil {
		return err
	}

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	var lguildInfo BasicGuildInformations

	lguildInfo.Deserialize(r)

	m.GuildInfo = &lguildInfo

	lexperienceForGuild, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.ExperienceForGuild = lexperienceForGuild

	return nil
}

type PaddockInformationsForSell struct {
	GuildOwner string

	WorldX int16

	WorldY int16

	SubAreaId uint16

	NbMount int8

	NbObject int8

	Price int64
}

func (m *PaddockInformationsForSell) ID() uint16 {
	return 222
}

func (m *PaddockInformationsForSell) Serialize(w Writer) error {

	if err := w.WriteString(m.GuildOwner); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.NbMount); err != nil {
		return err
	}

	if err := w.WriteInt8(m.NbObject); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Price); err != nil {
		return err
	}

	return nil
}

func (m *PaddockInformationsForSell) Deserialize(r Reader) error {

	lguildOwner, err := r.ReadString()
	if err != nil {
		return err
	}

	m.GuildOwner = lguildOwner

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lnbMount, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.NbMount = lnbMount

	lnbObject, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.NbObject = lnbObject

	lprice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Price = lprice

	return nil
}

type HouseInformationsForSell struct {
	InstanceId uint32

	SecondHand bool

	ModelId uint32

	OwnerName string

	OwnerConnected bool

	WorldX int16

	WorldY int16

	SubAreaId uint16

	NbRoom int8

	NbChest int8

	SkillListIds []int32

	IsLocked bool

	Price int64
}

func (m *HouseInformationsForSell) ID() uint16 {
	return 221
}

func (m *HouseInformationsForSell) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.SecondHand); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ModelId); err != nil {
		return err
	}

	if err := w.WriteString(m.OwnerName); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.OwnerConnected); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.NbRoom); err != nil {
		return err
	}

	if err := w.WriteInt8(m.NbChest); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.SkillListIds))); err != nil {
		return err
	}

	for i := range m.SkillListIds {

		if err := w.WriteInt32(m.SkillListIds[i]); err != nil {
			return err
		}

	}

	if err := w.WriteBoolean(m.IsLocked); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Price); err != nil {
		return err
	}

	return nil
}

func (m *HouseInformationsForSell) Deserialize(r Reader) error {

	linstanceId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.InstanceId = linstanceId

	lsecondHand, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.SecondHand = lsecondHand

	lmodelId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ModelId = lmodelId

	lownerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.OwnerName = lownerName

	lownerConnected, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.OwnerConnected = lownerConnected

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lnbRoom, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.NbRoom = lnbRoom

	lnbChest, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.NbChest = lnbChest

	lskillListIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SkillListIds = make([]int32, lskillListIdsLen)

	for i := range m.SkillListIds {

		lskillListIds, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.SkillListIds[i] = lskillListIds

	}

	lisLocked, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsLocked = lisLocked

	lprice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Price = lprice

	return nil
}

type FightTemporaryBoostWeaponDamagesEffect struct {
	FightTemporaryBoostEffect

	WeaponTypeId int16
}

func (m *FightTemporaryBoostWeaponDamagesEffect) ID() uint16 {
	return 211
}

func (m *FightTemporaryBoostWeaponDamagesEffect) Serialize(w Writer) error {

	if err := m.FightTemporaryBoostEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WeaponTypeId); err != nil {
		return err
	}

	return nil
}

func (m *FightTemporaryBoostWeaponDamagesEffect) Deserialize(r Reader) error {

	if err := m.FightTemporaryBoostEffect.Deserialize(r); err != nil {
		return err
	}

	lweaponTypeId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WeaponTypeId = lweaponTypeId

	return nil
}

type FightTemporarySpellBoostEffect struct {
	FightTemporaryBoostEffect

	BoostedSpellId uint16
}

func (m *FightTemporarySpellBoostEffect) ID() uint16 {
	return 207
}

func (m *FightTemporarySpellBoostEffect) Serialize(w Writer) error {

	if err := m.FightTemporaryBoostEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.BoostedSpellId); err != nil {
		return err
	}

	return nil
}

func (m *FightTemporarySpellBoostEffect) Deserialize(r Reader) error {

	if err := m.FightTemporaryBoostEffect.Deserialize(r); err != nil {
		return err
	}

	lboostedSpellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.BoostedSpellId = lboostedSpellId

	return nil
}

type FightTemporaryBoostStateEffect struct {
	FightTemporaryBoostEffect

	StateId int16
}

func (m *FightTemporaryBoostStateEffect) ID() uint16 {
	return 214
}

func (m *FightTemporaryBoostStateEffect) Serialize(w Writer) error {

	if err := m.FightTemporaryBoostEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.StateId); err != nil {
		return err
	}

	return nil
}

func (m *FightTemporaryBoostStateEffect) Deserialize(r Reader) error {

	if err := m.FightTemporaryBoostEffect.Deserialize(r); err != nil {
		return err
	}

	lstateId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.StateId = lstateId

	return nil
}

type FightTriggeredEffect struct {
	AbstractFightDispellableEffect

	Param1 int32

	Param2 int32

	Param3 int32

	Delay int16
}

func (m *FightTriggeredEffect) ID() uint16 {
	return 210
}

func (m *FightTriggeredEffect) Serialize(w Writer) error {

	if err := m.AbstractFightDispellableEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.Param1); err != nil {
		return err
	}

	if err := w.WriteInt32(m.Param2); err != nil {
		return err
	}

	if err := w.WriteInt32(m.Param3); err != nil {
		return err
	}

	if err := w.WriteInt16(m.Delay); err != nil {
		return err
	}

	return nil
}

func (m *FightTriggeredEffect) Deserialize(r Reader) error {

	if err := m.AbstractFightDispellableEffect.Deserialize(r); err != nil {
		return err
	}

	lparam1, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Param1 = lparam1

	lparam2, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Param2 = lparam2

	lparam3, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Param3 = lparam3

	ldelay, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Delay = ldelay

	return nil
}

type FightTemporarySpellImmunityEffect struct {
	AbstractFightDispellableEffect

	ImmuneSpellId int32
}

func (m *FightTemporarySpellImmunityEffect) ID() uint16 {
	return 366
}

func (m *FightTemporarySpellImmunityEffect) Serialize(w Writer) error {

	if err := m.AbstractFightDispellableEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.ImmuneSpellId); err != nil {
		return err
	}

	return nil
}

func (m *FightTemporarySpellImmunityEffect) Deserialize(r Reader) error {

	if err := m.AbstractFightDispellableEffect.Deserialize(r); err != nil {
		return err
	}

	limmuneSpellId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.ImmuneSpellId = limmuneSpellId

	return nil
}

type ActorExtendedAlignmentInformations struct {
	ActorAlignmentInformations

	Honor uint16

	HonorGradeFloor uint16

	HonorNextGradeFloor uint16

	Aggressable uint8
}

func (m *ActorExtendedAlignmentInformations) ID() uint16 {
	return 202
}

func (m *ActorExtendedAlignmentInformations) Serialize(w Writer) error {

	if err := m.ActorAlignmentInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Honor); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.HonorGradeFloor); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.HonorNextGradeFloor); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Aggressable); err != nil {
		return err
	}

	return nil
}

func (m *ActorExtendedAlignmentInformations) Deserialize(r Reader) error {

	if err := m.ActorAlignmentInformations.Deserialize(r); err != nil {
		return err
	}

	lhonor, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Honor = lhonor

	lhonorGradeFloor, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.HonorGradeFloor = lhonorGradeFloor

	lhonorNextGradeFloor, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.HonorNextGradeFloor = lhonorNextGradeFloor

	laggressable, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Aggressable = laggressable

	return nil
}

type PrismFightersInformation struct {
	SubAreaId uint16

	WaitingForHelpInfo *ProtectedEntityWaitingForHelpInfo

	AllyCharactersInformations []*CharacterMinimalPlusLookInformations

	EnemyCharactersInformations []*CharacterMinimalPlusLookInformations
}

func (m *PrismFightersInformation) ID() uint16 {
	return 443
}

func (m *PrismFightersInformation) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := m.WaitingForHelpInfo.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.AllyCharactersInformations))); err != nil {
		return err
	}

	for i := range m.AllyCharactersInformations {

		if err := w.WriteUInt16(m.AllyCharactersInformations[i].ID()); err != nil {
			return err
		}

		if err := m.AllyCharactersInformations[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.EnemyCharactersInformations))); err != nil {
		return err
	}

	for i := range m.EnemyCharactersInformations {

		if err := w.WriteUInt16(m.EnemyCharactersInformations[i].ID()); err != nil {
			return err
		}

		if err := m.EnemyCharactersInformations[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PrismFightersInformation) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	var lwaitingForHelpInfo ProtectedEntityWaitingForHelpInfo

	lwaitingForHelpInfo.Deserialize(r)

	m.WaitingForHelpInfo = &lwaitingForHelpInfo

	lallyCharactersInformationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AllyCharactersInformations = make([]*CharacterMinimalPlusLookInformations, lallyCharactersInformationsLen)

	for i := range m.AllyCharactersInformations {

		typeallyCharactersInformationsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lallyCharactersInformations, err := GetType(typeallyCharactersInformationsID)
		if err != nil {
			return err
		}

		lallyCharactersInformations.Deserialize(r)

		m.AllyCharactersInformations[i] = lallyCharactersInformations.(*CharacterMinimalPlusLookInformations)

	}

	lenemyCharactersInformationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.EnemyCharactersInformations = make([]*CharacterMinimalPlusLookInformations, lenemyCharactersInformationsLen)

	for i := range m.EnemyCharactersInformations {

		typeenemyCharactersInformationsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lenemyCharactersInformations, err := GetType(typeenemyCharactersInformationsID)
		if err != nil {
			return err
		}

		lenemyCharactersInformations.Deserialize(r)

		m.EnemyCharactersInformations[i] = lenemyCharactersInformations.(*CharacterMinimalPlusLookInformations)

	}

	return nil
}

type TaxCollectorFightersInformation struct {
	CollectorId int32

	AllyCharactersInformations []*CharacterMinimalPlusLookInformations

	EnemyCharactersInformations []*CharacterMinimalPlusLookInformations
}

func (m *TaxCollectorFightersInformation) ID() uint16 {
	return 169
}

func (m *TaxCollectorFightersInformation) Serialize(w Writer) error {

	if err := w.WriteInt32(m.CollectorId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.AllyCharactersInformations))); err != nil {
		return err
	}

	for i := range m.AllyCharactersInformations {

		if err := w.WriteUInt16(m.AllyCharactersInformations[i].ID()); err != nil {
			return err
		}

		if err := m.AllyCharactersInformations[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.EnemyCharactersInformations))); err != nil {
		return err
	}

	for i := range m.EnemyCharactersInformations {

		if err := w.WriteUInt16(m.EnemyCharactersInformations[i].ID()); err != nil {
			return err
		}

		if err := m.EnemyCharactersInformations[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *TaxCollectorFightersInformation) Deserialize(r Reader) error {

	lcollectorId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.CollectorId = lcollectorId

	lallyCharactersInformationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AllyCharactersInformations = make([]*CharacterMinimalPlusLookInformations, lallyCharactersInformationsLen)

	for i := range m.AllyCharactersInformations {

		typeallyCharactersInformationsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lallyCharactersInformations, err := GetType(typeallyCharactersInformationsID)
		if err != nil {
			return err
		}

		lallyCharactersInformations.Deserialize(r)

		m.AllyCharactersInformations[i] = lallyCharactersInformations.(*CharacterMinimalPlusLookInformations)

	}

	lenemyCharactersInformationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.EnemyCharactersInformations = make([]*CharacterMinimalPlusLookInformations, lenemyCharactersInformationsLen)

	for i := range m.EnemyCharactersInformations {

		typeenemyCharactersInformationsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lenemyCharactersInformations, err := GetType(typeenemyCharactersInformationsID)
		if err != nil {
			return err
		}

		lenemyCharactersInformations.Deserialize(r)

		m.EnemyCharactersInformations[i] = lenemyCharactersInformations.(*CharacterMinimalPlusLookInformations)

	}

	return nil
}

type ObjectItemMinimalInformation struct {
	Item

	ObjectGID uint16

	Effects []*ObjectEffect
}

func (m *ObjectItemMinimalInformation) ID() uint16 {
	return 124
}

func (m *ObjectItemMinimalInformation) Serialize(w Writer) error {

	if err := m.Item.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGID); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Effects))); err != nil {
		return err
	}

	for i := range m.Effects {

		if err := w.WriteUInt16(m.Effects[i].ID()); err != nil {
			return err
		}

		if err := m.Effects[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ObjectItemMinimalInformation) Deserialize(r Reader) error {

	if err := m.Item.Deserialize(r); err != nil {
		return err
	}

	lobjectGID, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGID = lobjectGID

	leffectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Effects = make([]*ObjectEffect, leffectsLen)

	for i := range m.Effects {

		typeeffectsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		leffects, err := GetType(typeeffectsID)
		if err != nil {
			return err
		}

		leffects.Deserialize(r)

		m.Effects[i] = leffects.(*ObjectEffect)

	}

	return nil
}

type ObjectItemInformationWithQuantity struct {
	ObjectItemMinimalInformation

	Quantity uint32
}

func (m *ObjectItemInformationWithQuantity) ID() uint16 {
	return 387
}

func (m *ObjectItemInformationWithQuantity) Serialize(w Writer) error {

	if err := m.ObjectItemMinimalInformation.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ObjectItemInformationWithQuantity) Deserialize(r Reader) error {

	if err := m.ObjectItemMinimalInformation.Deserialize(r); err != nil {
		return err
	}

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type StartupActionAddObject struct {
	Uid uint32

	Title string

	Text string

	DescUrl string

	PictureUrl string

	Items []*ObjectItemInformationWithQuantity
}

func (m *StartupActionAddObject) ID() uint16 {
	return 52
}

func (m *StartupActionAddObject) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.Uid); err != nil {
		return err
	}

	if err := w.WriteString(m.Title); err != nil {
		return err
	}

	if err := w.WriteString(m.Text); err != nil {
		return err
	}

	if err := w.WriteString(m.DescUrl); err != nil {
		return err
	}

	if err := w.WriteString(m.PictureUrl); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Items))); err != nil {
		return err
	}

	for i := range m.Items {

		if err := m.Items[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *StartupActionAddObject) Deserialize(r Reader) error {

	luid, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.Uid = luid

	ltitle, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Title = ltitle

	ltext, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Text = ltext

	ldescUrl, err := r.ReadString()
	if err != nil {
		return err
	}

	m.DescUrl = ldescUrl

	lpictureUrl, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PictureUrl = lpictureUrl

	litemsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Items = make([]*ObjectItemInformationWithQuantity, litemsLen)

	for i := range m.Items {

		var litems ObjectItemInformationWithQuantity

		litems.Deserialize(r)

		m.Items[i] = &litems

	}

	return nil
}

type DareInformations struct {
	DareId float64

	Creator *CharacterBasicMinimalInformations

	SubscriptionFee int64

	Jackpot int64

	MaxCountWinners uint16

	EndDate float64

	IsPrivate bool

	GuildId uint32

	AllianceId uint32

	Criterions []*DareCriteria

	StartDate float64
}

func (m *DareInformations) ID() uint16 {
	return 502
}

func (m *DareInformations) Serialize(w Writer) error {

	if err := w.WriteDouble(m.DareId); err != nil {
		return err
	}

	if err := m.Creator.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.SubscriptionFee); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Jackpot); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.MaxCountWinners); err != nil {
		return err
	}

	if err := w.WriteDouble(m.EndDate); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.IsPrivate); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.GuildId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.AllianceId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Criterions))); err != nil {
		return err
	}

	for i := range m.Criterions {

		if err := m.Criterions[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteDouble(m.StartDate); err != nil {
		return err
	}

	return nil
}

func (m *DareInformations) Deserialize(r Reader) error {

	ldareId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DareId = ldareId

	var lcreator CharacterBasicMinimalInformations

	lcreator.Deserialize(r)

	m.Creator = &lcreator

	lsubscriptionFee, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.SubscriptionFee = lsubscriptionFee

	ljackpot, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Jackpot = ljackpot

	lmaxCountWinners, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.MaxCountWinners = lmaxCountWinners

	lendDate, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.EndDate = lendDate

	lisPrivate, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsPrivate = lisPrivate

	lguildId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.GuildId = lguildId

	lallianceId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.AllianceId = lallianceId

	lcriterionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Criterions = make([]*DareCriteria, lcriterionsLen)

	for i := range m.Criterions {

		var lcriterions DareCriteria

		lcriterions.Deserialize(r)

		m.Criterions[i] = &lcriterions

	}

	lstartDate, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.StartDate = lstartDate

	return nil
}

type DareVersatileInformations struct {
	DareId float64

	CountEntrants uint32

	CountWinners uint32
}

func (m *DareVersatileInformations) ID() uint16 {
	return 504
}

func (m *DareVersatileInformations) Serialize(w Writer) error {

	if err := w.WriteDouble(m.DareId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.CountEntrants); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.CountWinners); err != nil {
		return err
	}

	return nil
}

func (m *DareVersatileInformations) Deserialize(r Reader) error {

	ldareId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DareId = ldareId

	lcountEntrants, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.CountEntrants = lcountEntrants

	lcountWinners, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.CountWinners = lcountWinners

	return nil
}

type DareCriteria struct {
	Type uint8

	Params []int32
}

func (m *DareCriteria) ID() uint16 {
	return 501
}

func (m *DareCriteria) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Params))); err != nil {
		return err
	}

	for i := range m.Params {

		if err := w.WriteInt32(m.Params[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *DareCriteria) Deserialize(r Reader) error {

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	lparamsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Params = make([]int32, lparamsLen)

	for i := range m.Params {

		lparams, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.Params[i] = lparams

	}

	return nil
}

type GuildEmblem struct {
	SymbolShape uint16

	SymbolColor int32

	BackgroundShape uint8

	BackgroundColor int32
}

func (m *GuildEmblem) ID() uint16 {
	return 87
}

func (m *GuildEmblem) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SymbolShape); err != nil {
		return err
	}

	if err := w.WriteInt32(m.SymbolColor); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.BackgroundShape); err != nil {
		return err
	}

	if err := w.WriteInt32(m.BackgroundColor); err != nil {
		return err
	}

	return nil
}

func (m *GuildEmblem) Deserialize(r Reader) error {

	lsymbolShape, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SymbolShape = lsymbolShape

	lsymbolColor, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.SymbolColor = lsymbolColor

	lbackgroundShape, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BackgroundShape = lbackgroundShape

	lbackgroundColor, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.BackgroundColor = lbackgroundColor

	return nil
}

type AlliancedGuildFactSheetInformations struct {
	GuildInformations

	AllianceInfos *BasicNamedAllianceInformations
}

func (m *AlliancedGuildFactSheetInformations) ID() uint16 {
	return 422
}

func (m *AlliancedGuildFactSheetInformations) Serialize(w Writer) error {

	if err := m.GuildInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.AllianceInfos.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AlliancedGuildFactSheetInformations) Deserialize(r Reader) error {

	if err := m.GuildInformations.Deserialize(r); err != nil {
		return err
	}

	var lallianceInfos BasicNamedAllianceInformations

	lallianceInfos.Deserialize(r)

	m.AllianceInfos = &lallianceInfos

	return nil
}

type GuildVersatileInformations struct {
	GuildId uint32

	LeaderId int64

	GuildLevel uint8

	NbMembers uint8
}

func (m *GuildVersatileInformations) ID() uint16 {
	return 435
}

func (m *GuildVersatileInformations) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.GuildId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.LeaderId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.GuildLevel); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.NbMembers); err != nil {
		return err
	}

	return nil
}

func (m *GuildVersatileInformations) Deserialize(r Reader) error {

	lguildId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.GuildId = lguildId

	lleaderId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.LeaderId = lleaderId

	lguildLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.GuildLevel = lguildLevel

	lnbMembers, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NbMembers = lnbMembers

	return nil
}

type AllianceVersatileInformations struct {
	AllianceId uint32

	NbGuilds uint16

	NbMembers uint16

	NbSubarea uint16
}

func (m *AllianceVersatileInformations) ID() uint16 {
	return 432
}

func (m *AllianceVersatileInformations) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.AllianceId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NbGuilds); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NbMembers); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NbSubarea); err != nil {
		return err
	}

	return nil
}

func (m *AllianceVersatileInformations) Deserialize(r Reader) error {

	lallianceId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.AllianceId = lallianceId

	lnbGuilds, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NbGuilds = lnbGuilds

	lnbMembers, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NbMembers = lnbMembers

	lnbSubarea, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NbSubarea = lnbSubarea

	return nil
}

type BasicNamedAllianceInformations struct {
	BasicAllianceInformations

	AllianceName string
}

func (m *BasicNamedAllianceInformations) ID() uint16 {
	return 418
}

func (m *BasicNamedAllianceInformations) Serialize(w Writer) error {

	if err := m.BasicAllianceInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.AllianceName); err != nil {
		return err
	}

	return nil
}

func (m *BasicNamedAllianceInformations) Deserialize(r Reader) error {

	if err := m.BasicAllianceInformations.Deserialize(r); err != nil {
		return err
	}

	lallianceName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.AllianceName = lallianceName

	return nil
}

type AllianceInformations struct {
	BasicNamedAllianceInformations

	AllianceEmblem *GuildEmblem
}

func (m *AllianceInformations) ID() uint16 {
	return 417
}

func (m *AllianceInformations) Serialize(w Writer) error {

	if err := m.BasicNamedAllianceInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.AllianceEmblem.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AllianceInformations) Deserialize(r Reader) error {

	if err := m.BasicNamedAllianceInformations.Deserialize(r); err != nil {
		return err
	}

	var lallianceEmblem GuildEmblem

	lallianceEmblem.Deserialize(r)

	m.AllianceEmblem = &lallianceEmblem

	return nil
}

type AllianceFactSheetInformations struct {
	AllianceInformations

	CreationDate uint32
}

func (m *AllianceFactSheetInformations) ID() uint16 {
	return 421
}

func (m *AllianceFactSheetInformations) Serialize(w Writer) error {

	if err := m.AllianceInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.CreationDate); err != nil {
		return err
	}

	return nil
}

func (m *AllianceFactSheetInformations) Deserialize(r Reader) error {

	if err := m.AllianceInformations.Deserialize(r); err != nil {
		return err
	}

	lcreationDate, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.CreationDate = lcreationDate

	return nil
}

type TaxCollectorComplementaryInformations struct {
}

func (m *TaxCollectorComplementaryInformations) ID() uint16 {
	return 448
}

func (m *TaxCollectorComplementaryInformations) Serialize(w Writer) error {

	return nil
}

func (m *TaxCollectorComplementaryInformations) Deserialize(r Reader) error {

	return nil
}

type TaxCollectorGuildInformations struct {
	TaxCollectorComplementaryInformations

	Guild *BasicGuildInformations
}

func (m *TaxCollectorGuildInformations) ID() uint16 {
	return 446
}

func (m *TaxCollectorGuildInformations) Serialize(w Writer) error {

	if err := m.TaxCollectorComplementaryInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.Guild.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorGuildInformations) Deserialize(r Reader) error {

	if err := m.TaxCollectorComplementaryInformations.Deserialize(r); err != nil {
		return err
	}

	var lguild BasicGuildInformations

	lguild.Deserialize(r)

	m.Guild = &lguild

	return nil
}

type AdditionalTaxCollectorInformations struct {
	CollectorCallerName string

	Date uint32
}

func (m *AdditionalTaxCollectorInformations) ID() uint16 {
	return 165
}

func (m *AdditionalTaxCollectorInformations) Serialize(w Writer) error {

	if err := w.WriteString(m.CollectorCallerName); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.Date); err != nil {
		return err
	}

	return nil
}

func (m *AdditionalTaxCollectorInformations) Deserialize(r Reader) error {

	lcollectorCallerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.CollectorCallerName = lcollectorCallerName

	ldate, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.Date = ldate

	return nil
}

type TaxCollectorLootInformations struct {
	TaxCollectorComplementaryInformations

	Kamas int64

	Experience int64

	Pods uint32

	ItemsValue int64
}

func (m *TaxCollectorLootInformations) ID() uint16 {
	return 372
}

func (m *TaxCollectorLootInformations) Serialize(w Writer) error {

	if err := m.TaxCollectorComplementaryInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Kamas); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Experience); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Pods); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ItemsValue); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorLootInformations) Deserialize(r Reader) error {

	if err := m.TaxCollectorComplementaryInformations.Deserialize(r); err != nil {
		return err
	}

	lkamas, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Kamas = lkamas

	lexperience, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Experience = lexperience

	lpods, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Pods = lpods

	litemsValue, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ItemsValue = litemsValue

	return nil
}

type TaxCollectorWaitingForHelpInformations struct {
	TaxCollectorComplementaryInformations

	WaitingForHelpInfo *ProtectedEntityWaitingForHelpInfo
}

func (m *TaxCollectorWaitingForHelpInformations) ID() uint16 {
	return 447
}

func (m *TaxCollectorWaitingForHelpInformations) Serialize(w Writer) error {

	if err := m.TaxCollectorComplementaryInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.WaitingForHelpInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorWaitingForHelpInformations) Deserialize(r Reader) error {

	if err := m.TaxCollectorComplementaryInformations.Deserialize(r); err != nil {
		return err
	}

	var lwaitingForHelpInfo ProtectedEntityWaitingForHelpInfo

	lwaitingForHelpInfo.Deserialize(r)

	m.WaitingForHelpInfo = &lwaitingForHelpInfo

	return nil
}

type FriendSpouseInformations struct {
	SpouseAccountId uint32

	SpouseId int64

	SpouseName string

	SpouseLevel uint8

	Breed int8

	Sex int8

	SpouseEntityLook *EntityLook

	GuildInfo *GuildInformations

	AlignmentSide int8
}

func (m *FriendSpouseInformations) ID() uint16 {
	return 77
}

func (m *FriendSpouseInformations) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.SpouseAccountId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.SpouseId); err != nil {
		return err
	}

	if err := w.WriteString(m.SpouseName); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.SpouseLevel); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Sex); err != nil {
		return err
	}

	if err := m.SpouseEntityLook.Serialize(w); err != nil {
		return err
	}

	if err := m.GuildInfo.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt8(m.AlignmentSide); err != nil {
		return err
	}

	return nil
}

func (m *FriendSpouseInformations) Deserialize(r Reader) error {

	lspouseAccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.SpouseAccountId = lspouseAccountId

	lspouseId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.SpouseId = lspouseId

	lspouseName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.SpouseName = lspouseName

	lspouseLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.SpouseLevel = lspouseLevel

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	lsex, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Sex = lsex

	var lspouseEntityLook EntityLook

	lspouseEntityLook.Deserialize(r)

	m.SpouseEntityLook = &lspouseEntityLook

	var lguildInfo GuildInformations

	lguildInfo.Deserialize(r)

	m.GuildInfo = &lguildInfo

	lalignmentSide, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.AlignmentSide = lalignmentSide

	return nil
}

type FriendSpouseOnlineInformations struct {
	FriendSpouseInformations

	MapId uint32

	SubAreaId uint16

	InFight bool

	FollowSpouse bool
}

func (m *FriendSpouseOnlineInformations) ID() uint16 {
	return 93
}

func (m *FriendSpouseOnlineInformations) Serialize(w Writer) error {

	if err := m.FriendSpouseInformations.Serialize(w); err != nil {
		return err
	}

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.InFight)

	setWrappedFlag(bbw0, 1, m.FollowSpouse)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	return nil
}

func (m *FriendSpouseOnlineInformations) Deserialize(r Reader) error {

	if err := m.FriendSpouseInformations.Deserialize(r); err != nil {
		return err
	}

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.InFight = getWrappedFlag(bbw0, 0)

	m.FollowSpouse = getWrappedFlag(bbw0, 1)

	lmapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	return nil
}

type ObjectItemQuantity struct {
	Item

	ObjectUID uint32

	Quantity uint32
}

func (m *ObjectItemQuantity) ID() uint16 {
	return 119
}

func (m *ObjectItemQuantity) Serialize(w Writer) error {

	if err := m.Item.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ObjectItemQuantity) Deserialize(r Reader) error {

	if err := m.Item.Deserialize(r); err != nil {
		return err
	}

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type Shortcut struct {
	Slot uint8
}

func (m *Shortcut) ID() uint16 {
	return 369
}

func (m *Shortcut) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Slot); err != nil {
		return err
	}

	return nil
}

func (m *Shortcut) Deserialize(r Reader) error {

	lslot, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Slot = lslot

	return nil
}

type ShortcutObject struct {
	Shortcut
}

func (m *ShortcutObject) ID() uint16 {
	return 367
}

func (m *ShortcutObject) Serialize(w Writer) error {

	if err := m.Shortcut.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutObject) Deserialize(r Reader) error {

	if err := m.Shortcut.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type ShortcutObjectIdolsPreset struct {
	ShortcutObject

	PresetId uint8
}

func (m *ShortcutObjectIdolsPreset) ID() uint16 {
	return 492
}

func (m *ShortcutObjectIdolsPreset) Serialize(w Writer) error {

	if err := m.ShortcutObject.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutObjectIdolsPreset) Deserialize(r Reader) error {

	if err := m.ShortcutObject.Deserialize(r); err != nil {
		return err
	}

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	return nil
}

type Preset struct {
	PresetId uint8

	SymbolId uint8

	Mount bool

	Objects []*PresetItem
}

func (m *Preset) ID() uint16 {
	return 355
}

func (m *Preset) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.SymbolId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Mount); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Objects))); err != nil {
		return err
	}

	for i := range m.Objects {

		if err := m.Objects[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *Preset) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	lsymbolId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.SymbolId = lsymbolId

	lmount, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Mount = lmount

	lobjectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Objects = make([]*PresetItem, lobjectsLen)

	for i := range m.Objects {

		var lobjects PresetItem

		lobjects.Deserialize(r)

		m.Objects[i] = &lobjects

	}

	return nil
}

type ShortcutSmiley struct {
	Shortcut

	SmileyId uint16
}

func (m *ShortcutSmiley) ID() uint16 {
	return 388
}

func (m *ShortcutSmiley) Serialize(w Writer) error {

	if err := m.Shortcut.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SmileyId); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutSmiley) Deserialize(r Reader) error {

	if err := m.Shortcut.Deserialize(r); err != nil {
		return err
	}

	lsmileyId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SmileyId = lsmileyId

	return nil
}

type IdolsPreset struct {
	PresetId uint8

	SymbolId uint8

	IdolId []uint16
}

func (m *IdolsPreset) ID() uint16 {
	return 491
}

func (m *IdolsPreset) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.SymbolId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.IdolId))); err != nil {
		return err
	}

	for i := range m.IdolId {

		if err := w.WriteVarUInt16(m.IdolId[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *IdolsPreset) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	lsymbolId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.SymbolId = lsymbolId

	lidolIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.IdolId = make([]uint16, lidolIdLen)

	for i := range m.IdolId {

		lidolId, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.IdolId[i] = lidolId

	}

	return nil
}

type ShortcutEmote struct {
	Shortcut

	EmoteId uint8
}

func (m *ShortcutEmote) ID() uint16 {
	return 389
}

func (m *ShortcutEmote) Serialize(w Writer) error {

	if err := m.Shortcut.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.EmoteId); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutEmote) Deserialize(r Reader) error {

	if err := m.Shortcut.Deserialize(r); err != nil {
		return err
	}

	lemoteId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.EmoteId = lemoteId

	return nil
}

type ShortcutObjectPreset struct {
	ShortcutObject

	PresetId uint8
}

func (m *ShortcutObjectPreset) ID() uint16 {
	return 370
}

func (m *ShortcutObjectPreset) Serialize(w Writer) error {

	if err := m.ShortcutObject.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutObjectPreset) Deserialize(r Reader) error {

	if err := m.ShortcutObject.Deserialize(r); err != nil {
		return err
	}

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	return nil
}

type ShortcutObjectItem struct {
	ShortcutObject

	ItemUID int32

	ItemGID int32
}

func (m *ShortcutObjectItem) ID() uint16 {
	return 371
}

func (m *ShortcutObjectItem) Serialize(w Writer) error {

	if err := m.ShortcutObject.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.ItemUID); err != nil {
		return err
	}

	if err := w.WriteInt32(m.ItemGID); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutObjectItem) Deserialize(r Reader) error {

	if err := m.ShortcutObject.Deserialize(r); err != nil {
		return err
	}

	litemUID, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.ItemUID = litemUID

	litemGID, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.ItemGID = litemGID

	return nil
}

type ShortcutSpell struct {
	Shortcut

	SpellId uint16
}

func (m *ShortcutSpell) ID() uint16 {
	return 368
}

func (m *ShortcutSpell) Serialize(w Writer) error {

	if err := m.Shortcut.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutSpell) Deserialize(r Reader) error {

	if err := m.Shortcut.Deserialize(r); err != nil {
		return err
	}

	lspellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	return nil
}

type MountClientData struct {
	Id float64

	Model uint32

	Ancestor []uint32

	Behaviors []uint32

	Name string

	Sex bool

	OwnerId uint32

	Experience int64

	ExperienceForLevel int64

	ExperienceForNextLevel float64

	Level uint8

	IsRideable bool

	MaxPods uint32

	IsWild bool

	Stamina uint32

	StaminaMax uint32

	Maturity uint32

	MaturityForAdult uint32

	Energy uint32

	EnergyMax uint32

	Serenity int32

	AggressivityMax int32

	SerenityMax uint32

	Love uint32

	LoveMax uint32

	FecondationTime int32

	IsFecondationReady bool

	BoostLimiter uint32

	BoostMax float64

	ReproductionCount int32

	ReproductionCountMax uint32

	HarnessGID uint16

	UseHarnessColors bool

	EffectList []*ObjectEffectInteger
}

func (m *MountClientData) ID() uint16 {
	return 178
}

func (m *MountClientData) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Sex)

	setWrappedFlag(bbw0, 1, m.IsRideable)

	setWrappedFlag(bbw0, 2, m.IsWild)

	setWrappedFlag(bbw0, 3, m.IsFecondationReady)

	setWrappedFlag(bbw0, 4, m.UseHarnessColors)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Model); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Ancestor))); err != nil {
		return err
	}

	for i := range m.Ancestor {

		if err := w.WriteUInt32(m.Ancestor[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Behaviors))); err != nil {
		return err
	}

	for i := range m.Behaviors {

		if err := w.WriteUInt32(m.Behaviors[i]); err != nil {
			return err
		}

	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.OwnerId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Experience); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceForLevel); err != nil {
		return err
	}

	if err := w.WriteDouble(m.ExperienceForNextLevel); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MaxPods); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Stamina); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.StaminaMax); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Maturity); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MaturityForAdult); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Energy); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.EnergyMax); err != nil {
		return err
	}

	if err := w.WriteInt32(m.Serenity); err != nil {
		return err
	}

	if err := w.WriteInt32(m.AggressivityMax); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.SerenityMax); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Love); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.LoveMax); err != nil {
		return err
	}

	if err := w.WriteInt32(m.FecondationTime); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.BoostLimiter); err != nil {
		return err
	}

	if err := w.WriteDouble(m.BoostMax); err != nil {
		return err
	}

	if err := w.WriteInt32(m.ReproductionCount); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ReproductionCountMax); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.HarnessGID); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.EffectList))); err != nil {
		return err
	}

	for i := range m.EffectList {

		if err := m.EffectList[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *MountClientData) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Sex = getWrappedFlag(bbw0, 0)

	m.IsRideable = getWrappedFlag(bbw0, 1)

	m.IsWild = getWrappedFlag(bbw0, 2)

	m.IsFecondationReady = getWrappedFlag(bbw0, 3)

	m.UseHarnessColors = getWrappedFlag(bbw0, 4)

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	lmodel, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Model = lmodel

	lancestorLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Ancestor = make([]uint32, lancestorLen)

	for i := range m.Ancestor {

		lancestor, err := r.ReadUInt32()
		if err != nil {
			return err
		}

		m.Ancestor[i] = lancestor

	}

	lbehaviorsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Behaviors = make([]uint32, lbehaviorsLen)

	for i := range m.Behaviors {

		lbehaviors, err := r.ReadUInt32()
		if err != nil {
			return err
		}

		m.Behaviors[i] = lbehaviors

	}

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	lownerId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.OwnerId = lownerId

	lexperience, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Experience = lexperience

	lexperienceForLevel, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceForLevel = lexperienceForLevel

	lexperienceForNextLevel, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.ExperienceForNextLevel = lexperienceForNextLevel

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	lmaxPods, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MaxPods = lmaxPods

	lstamina, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Stamina = lstamina

	lstaminaMax, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.StaminaMax = lstaminaMax

	lmaturity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Maturity = lmaturity

	lmaturityForAdult, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MaturityForAdult = lmaturityForAdult

	lenergy, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Energy = lenergy

	lenergyMax, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.EnergyMax = lenergyMax

	lserenity, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Serenity = lserenity

	laggressivityMax, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.AggressivityMax = laggressivityMax

	lserenityMax, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SerenityMax = lserenityMax

	llove, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Love = llove

	lloveMax, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.LoveMax = lloveMax

	lfecondationTime, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FecondationTime = lfecondationTime

	lboostLimiter, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.BoostLimiter = lboostLimiter

	lboostMax, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.BoostMax = lboostMax

	lreproductionCount, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.ReproductionCount = lreproductionCount

	lreproductionCountMax, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ReproductionCountMax = lreproductionCountMax

	lharnessGID, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.HarnessGID = lharnessGID

	leffectListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.EffectList = make([]*ObjectEffectInteger, leffectListLen)

	for i := range m.EffectList {

		var leffectList ObjectEffectInteger

		leffectList.Deserialize(r)

		m.EffectList[i] = &leffectList

	}

	return nil
}

type UpdateMountBoost struct {
	Type uint8
}

func (m *UpdateMountBoost) ID() uint16 {
	return 356
}

func (m *UpdateMountBoost) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMountBoost) Deserialize(r Reader) error {

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	return nil
}

type UpdateMountIntBoost struct {
	UpdateMountBoost

	Value int32
}

func (m *UpdateMountIntBoost) ID() uint16 {
	return 357
}

func (m *UpdateMountIntBoost) Serialize(w Writer) error {

	if err := m.UpdateMountBoost.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *UpdateMountIntBoost) Deserialize(r Reader) error {

	if err := m.UpdateMountBoost.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type StatisticData struct {
}

func (m *StatisticData) ID() uint16 {
	return 484
}

func (m *StatisticData) Serialize(w Writer) error {

	return nil
}

func (m *StatisticData) Deserialize(r Reader) error {

	return nil
}

type StatisticDataInt struct {
	StatisticData

	Value int32
}

func (m *StatisticDataInt) ID() uint16 {
	return 485
}

func (m *StatisticDataInt) Serialize(w Writer) error {

	if err := m.StatisticData.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *StatisticDataInt) Deserialize(r Reader) error {

	if err := m.StatisticData.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type CharacterMinimalGuildInformations struct {
	CharacterMinimalPlusLookInformations

	Guild *BasicGuildInformations
}

func (m *CharacterMinimalGuildInformations) ID() uint16 {
	return 445
}

func (m *CharacterMinimalGuildInformations) Serialize(w Writer) error {

	if err := m.CharacterMinimalPlusLookInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.Guild.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *CharacterMinimalGuildInformations) Deserialize(r Reader) error {

	if err := m.CharacterMinimalPlusLookInformations.Deserialize(r); err != nil {
		return err
	}

	var lguild BasicGuildInformations

	lguild.Deserialize(r)

	m.Guild = &lguild

	return nil
}

type StatisticDataByte struct {
	StatisticData

	Value int8
}

func (m *StatisticDataByte) ID() uint16 {
	return 486
}

func (m *StatisticDataByte) Serialize(w Writer) error {

	if err := m.StatisticData.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *StatisticDataByte) Deserialize(r Reader) error {

	if err := m.StatisticData.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type QuestObjectiveInformations struct {
	ObjectiveId uint16

	ObjectiveStatus bool

	DialogParams []string
}

func (m *QuestObjectiveInformations) ID() uint16 {
	return 385
}

func (m *QuestObjectiveInformations) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ObjectiveId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.ObjectiveStatus); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.DialogParams))); err != nil {
		return err
	}

	for i := range m.DialogParams {

		if err := w.WriteString(m.DialogParams[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *QuestObjectiveInformations) Deserialize(r Reader) error {

	lobjectiveId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectiveId = lobjectiveId

	lobjectiveStatus, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.ObjectiveStatus = lobjectiveStatus

	ldialogParamsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DialogParams = make([]string, ldialogParamsLen)

	for i := range m.DialogParams {

		ldialogParams, err := r.ReadString()
		if err != nil {
			return err
		}

		m.DialogParams[i] = ldialogParams

	}

	return nil
}

type FightTeamMemberTaxCollectorInformations struct {
	FightTeamMemberInformations

	FirstNameId uint16

	LastNameId uint16

	Level uint8

	GuildId uint32

	Uid uint32
}

func (m *FightTeamMemberTaxCollectorInformations) ID() uint16 {
	return 177
}

func (m *FightTeamMemberTaxCollectorInformations) Serialize(w Writer) error {

	if err := m.FightTeamMemberInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.FirstNameId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.LastNameId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.GuildId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Uid); err != nil {
		return err
	}

	return nil
}

func (m *FightTeamMemberTaxCollectorInformations) Deserialize(r Reader) error {

	if err := m.FightTeamMemberInformations.Deserialize(r); err != nil {
		return err
	}

	lfirstNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.FirstNameId = lfirstNameId

	llastNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.LastNameId = llastNameId

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	lguildId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.GuildId = lguildId

	luid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Uid = luid

	return nil
}

type ServerSessionConstant struct {
	Id uint16
}

func (m *ServerSessionConstant) ID() uint16 {
	return 430
}

func (m *ServerSessionConstant) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *ServerSessionConstant) Deserialize(r Reader) error {

	lid, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type ServerSessionConstantString struct {
	ServerSessionConstant

	Value string
}

func (m *ServerSessionConstantString) ID() uint16 {
	return 436
}

func (m *ServerSessionConstantString) Serialize(w Writer) error {

	if err := m.ServerSessionConstant.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *ServerSessionConstantString) Deserialize(r Reader) error {

	if err := m.ServerSessionConstant.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type MapCoordinates struct {
	WorldX int16

	WorldY int16
}

func (m *MapCoordinates) ID() uint16 {
	return 174
}

func (m *MapCoordinates) Serialize(w Writer) error {

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	return nil
}

func (m *MapCoordinates) Deserialize(r Reader) error {

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	return nil
}

type MapCoordinatesAndId struct {
	MapCoordinates

	MapId int32
}

func (m *MapCoordinatesAndId) ID() uint16 {
	return 392
}

func (m *MapCoordinatesAndId) Serialize(w Writer) error {

	if err := m.MapCoordinates.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	return nil
}

func (m *MapCoordinatesAndId) Deserialize(r Reader) error {

	if err := m.MapCoordinates.Deserialize(r); err != nil {
		return err
	}

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	return nil
}

type MapCoordinatesExtended struct {
	MapCoordinatesAndId

	SubAreaId uint16
}

func (m *MapCoordinatesExtended) ID() uint16 {
	return 176
}

func (m *MapCoordinatesExtended) Serialize(w Writer) error {

	if err := m.MapCoordinatesAndId.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	return nil
}

func (m *MapCoordinatesExtended) Deserialize(r Reader) error {

	if err := m.MapCoordinatesAndId.Deserialize(r); err != nil {
		return err
	}

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	return nil
}

type TreasureHuntStep struct {
}

func (m *TreasureHuntStep) ID() uint16 {
	return 463
}

func (m *TreasureHuntStep) Serialize(w Writer) error {

	return nil
}

func (m *TreasureHuntStep) Deserialize(r Reader) error {

	return nil
}

type TreasureHuntStepFight struct {
	TreasureHuntStep
}

func (m *TreasureHuntStepFight) ID() uint16 {
	return 462
}

func (m *TreasureHuntStepFight) Serialize(w Writer) error {

	if err := m.TreasureHuntStep.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntStepFight) Deserialize(r Reader) error {

	if err := m.TreasureHuntStep.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type QuestActiveDetailedInformations struct {
	QuestActiveInformations

	StepId uint16

	Objectives []*QuestObjectiveInformations
}

func (m *QuestActiveDetailedInformations) ID() uint16 {
	return 382
}

func (m *QuestActiveDetailedInformations) Serialize(w Writer) error {

	if err := m.QuestActiveInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.StepId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Objectives))); err != nil {
		return err
	}

	for i := range m.Objectives {

		if err := w.WriteUInt16(m.Objectives[i].ID()); err != nil {
			return err
		}

		if err := m.Objectives[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *QuestActiveDetailedInformations) Deserialize(r Reader) error {

	if err := m.QuestActiveInformations.Deserialize(r); err != nil {
		return err
	}

	lstepId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.StepId = lstepId

	lobjectivesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Objectives = make([]*QuestObjectiveInformations, lobjectivesLen)

	for i := range m.Objectives {

		typeobjectivesID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lobjectives, err := GetType(typeobjectivesID)
		if err != nil {
			return err
		}

		lobjectives.Deserialize(r)

		m.Objectives[i] = lobjectives.(*QuestObjectiveInformations)

	}

	return nil
}

type ObjectEffectCreature struct {
	ObjectEffect

	MonsterFamilyId uint16
}

func (m *ObjectEffectCreature) ID() uint16 {
	return 71
}

func (m *ObjectEffectCreature) Serialize(w Writer) error {

	if err := m.ObjectEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MonsterFamilyId); err != nil {
		return err
	}

	return nil
}

func (m *ObjectEffectCreature) Deserialize(r Reader) error {

	if err := m.ObjectEffect.Deserialize(r); err != nil {
		return err
	}

	lmonsterFamilyId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MonsterFamilyId = lmonsterFamilyId

	return nil
}

type ObjectEffectString struct {
	ObjectEffect

	Value string
}

func (m *ObjectEffectString) ID() uint16 {
	return 74
}

func (m *ObjectEffectString) Serialize(w Writer) error {

	if err := m.ObjectEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *ObjectEffectString) Deserialize(r Reader) error {

	if err := m.ObjectEffect.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type StatisticDataBoolean struct {
	StatisticData

	Value bool
}

func (m *StatisticDataBoolean) ID() uint16 {
	return 482
}

func (m *StatisticDataBoolean) Serialize(w Writer) error {

	if err := m.StatisticData.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *StatisticDataBoolean) Deserialize(r Reader) error {

	if err := m.StatisticData.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type TaxCollectorStaticExtendedInformations struct {
	TaxCollectorStaticInformations

	AllianceIdentity *AllianceInformations
}

func (m *TaxCollectorStaticExtendedInformations) ID() uint16 {
	return 440
}

func (m *TaxCollectorStaticExtendedInformations) Serialize(w Writer) error {

	if err := m.TaxCollectorStaticInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.AllianceIdentity.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorStaticExtendedInformations) Deserialize(r Reader) error {

	if err := m.TaxCollectorStaticInformations.Deserialize(r); err != nil {
		return err
	}

	var lallianceIdentity AllianceInformations

	lallianceIdentity.Deserialize(r)

	m.AllianceIdentity = &lallianceIdentity

	return nil
}

type FightResultAdditionalData struct {
}

func (m *FightResultAdditionalData) ID() uint16 {
	return 191
}

func (m *FightResultAdditionalData) Serialize(w Writer) error {

	return nil
}

func (m *FightResultAdditionalData) Deserialize(r Reader) error {

	return nil
}

type FightResultPvpData struct {
	FightResultAdditionalData

	Grade uint8

	MinHonorForGrade uint16

	MaxHonorForGrade uint16

	Honor uint16

	HonorDelta int16
}

func (m *FightResultPvpData) ID() uint16 {
	return 190
}

func (m *FightResultPvpData) Serialize(w Writer) error {

	if err := m.FightResultAdditionalData.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Grade); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MinHonorForGrade); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MaxHonorForGrade); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Honor); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.HonorDelta); err != nil {
		return err
	}

	return nil
}

func (m *FightResultPvpData) Deserialize(r Reader) error {

	if err := m.FightResultAdditionalData.Deserialize(r); err != nil {
		return err
	}

	lgrade, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Grade = lgrade

	lminHonorForGrade, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MinHonorForGrade = lminHonorForGrade

	lmaxHonorForGrade, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MaxHonorForGrade = lmaxHonorForGrade

	lhonor, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Honor = lhonor

	lhonorDelta, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.HonorDelta = lhonorDelta

	return nil
}

type CharacterMinimalAllianceInformations struct {
	CharacterMinimalGuildInformations

	Alliance *BasicAllianceInformations
}

func (m *CharacterMinimalAllianceInformations) ID() uint16 {
	return 444
}

func (m *CharacterMinimalAllianceInformations) Serialize(w Writer) error {

	if err := m.CharacterMinimalGuildInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.Alliance.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *CharacterMinimalAllianceInformations) Deserialize(r Reader) error {

	if err := m.CharacterMinimalGuildInformations.Deserialize(r); err != nil {
		return err
	}

	var lalliance BasicAllianceInformations

	lalliance.Deserialize(r)

	m.Alliance = &lalliance

	return nil
}

type FightTeamMemberMonsterInformations struct {
	FightTeamMemberInformations

	MonsterId int32

	Grade uint8
}

func (m *FightTeamMemberMonsterInformations) ID() uint16 {
	return 6
}

func (m *FightTeamMemberMonsterInformations) Serialize(w Writer) error {

	if err := m.FightTeamMemberInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MonsterId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Grade); err != nil {
		return err
	}

	return nil
}

func (m *FightTeamMemberMonsterInformations) Deserialize(r Reader) error {

	if err := m.FightTeamMemberInformations.Deserialize(r); err != nil {
		return err
	}

	lmonsterId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MonsterId = lmonsterId

	lgrade, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Grade = lgrade

	return nil
}

type ObjectEffectMinMax struct {
	ObjectEffect

	Min uint32

	Max uint32
}

func (m *ObjectEffectMinMax) ID() uint16 {
	return 82
}

func (m *ObjectEffectMinMax) Serialize(w Writer) error {

	if err := m.ObjectEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Min); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Max); err != nil {
		return err
	}

	return nil
}

func (m *ObjectEffectMinMax) Deserialize(r Reader) error {

	if err := m.ObjectEffect.Deserialize(r); err != nil {
		return err
	}

	lmin, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Min = lmin

	lmax, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Max = lmax

	return nil
}

type FightAllianceTeamInformations struct {
	FightTeamInformations

	Relation uint8
}

func (m *FightAllianceTeamInformations) ID() uint16 {
	return 439
}

func (m *FightAllianceTeamInformations) Serialize(w Writer) error {

	if err := m.FightTeamInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Relation); err != nil {
		return err
	}

	return nil
}

func (m *FightAllianceTeamInformations) Deserialize(r Reader) error {

	if err := m.FightTeamInformations.Deserialize(r); err != nil {
		return err
	}

	lrelation, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Relation = lrelation

	return nil
}

type PartyIdol struct {
	Idol

	OwnersIds []int64
}

func (m *PartyIdol) ID() uint16 {
	return 490
}

func (m *PartyIdol) Serialize(w Writer) error {

	if err := m.Idol.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.OwnersIds))); err != nil {
		return err
	}

	for i := range m.OwnersIds {

		if err := w.WriteVarInt64(m.OwnersIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *PartyIdol) Deserialize(r Reader) error {

	if err := m.Idol.Deserialize(r); err != nil {
		return err
	}

	lownersIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.OwnersIds = make([]int64, lownersIdsLen)

	for i := range m.OwnersIds {

		lownersIds, err := r.ReadVarInt64()
		if err != nil {
			return err
		}

		m.OwnersIds[i] = lownersIds

	}

	return nil
}

type GameFightMonsterWithAlignmentInformations struct {
	GameFightMonsterInformations

	AlignmentInfos *ActorAlignmentInformations
}

func (m *GameFightMonsterWithAlignmentInformations) ID() uint16 {
	return 203
}

func (m *GameFightMonsterWithAlignmentInformations) Serialize(w Writer) error {

	if err := m.GameFightMonsterInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.AlignmentInfos.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameFightMonsterWithAlignmentInformations) Deserialize(r Reader) error {

	if err := m.GameFightMonsterInformations.Deserialize(r); err != nil {
		return err
	}

	var lalignmentInfos ActorAlignmentInformations

	lalignmentInfos.Deserialize(r)

	m.AlignmentInfos = &lalignmentInfos

	return nil
}

type InteractiveElementWithAgeBonus struct {
	InteractiveElement

	AgeBonus int16
}

func (m *InteractiveElementWithAgeBonus) ID() uint16 {
	return 398
}

func (m *InteractiveElementWithAgeBonus) Serialize(w Writer) error {

	if err := m.InteractiveElement.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.AgeBonus); err != nil {
		return err
	}

	return nil
}

func (m *InteractiveElementWithAgeBonus) Deserialize(r Reader) error {

	if err := m.InteractiveElement.Deserialize(r); err != nil {
		return err
	}

	lageBonus, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AgeBonus = lageBonus

	return nil
}

type GameFightFighterLightInformations struct {
	Id float64

	Wave uint8

	Level uint16

	Breed int8

	Sex bool

	Alive bool
}

func (m *GameFightFighterLightInformations) ID() uint16 {
	return 413
}

func (m *GameFightFighterLightInformations) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Sex)

	setWrappedFlag(bbw0, 1, m.Alive)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Wave); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Level); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	return nil
}

func (m *GameFightFighterLightInformations) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Sex = getWrappedFlag(bbw0, 0)

	m.Alive = getWrappedFlag(bbw0, 1)

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	lwave, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Wave = lwave

	llevel, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Level = llevel

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	return nil
}

type FightTeamMemberCharacterInformations struct {
	FightTeamMemberInformations

	Name string

	Level uint8
}

func (m *FightTeamMemberCharacterInformations) ID() uint16 {
	return 13
}

func (m *FightTeamMemberCharacterInformations) Serialize(w Writer) error {

	if err := m.FightTeamMemberInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	return nil
}

func (m *FightTeamMemberCharacterInformations) Deserialize(r Reader) error {

	if err := m.FightTeamMemberInformations.Deserialize(r); err != nil {
		return err
	}

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	return nil
}

type ObjectEffectMount struct {
	ObjectEffect

	MountId uint32

	Date float64

	ModelId uint16
}

func (m *ObjectEffectMount) ID() uint16 {
	return 179
}

func (m *ObjectEffectMount) Serialize(w Writer) error {

	if err := m.ObjectEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.MountId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.Date); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ModelId); err != nil {
		return err
	}

	return nil
}

func (m *ObjectEffectMount) Deserialize(r Reader) error {

	if err := m.ObjectEffect.Deserialize(r); err != nil {
		return err
	}

	lmountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MountId = lmountId

	ldate, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Date = ldate

	lmodelId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ModelId = lmodelId

	return nil
}

type StatisticDataString struct {
	StatisticData

	Value string
}

func (m *StatisticDataString) ID() uint16 {
	return 487
}

func (m *StatisticDataString) Serialize(w Writer) error {

	if err := m.StatisticData.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *StatisticDataString) Deserialize(r Reader) error {

	if err := m.StatisticData.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type ServerSessionConstantLong struct {
	ServerSessionConstant

	Value float64
}

func (m *ServerSessionConstantLong) ID() uint16 {
	return 429
}

func (m *ServerSessionConstantLong) Serialize(w Writer) error {

	if err := m.ServerSessionConstant.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *ServerSessionConstantLong) Deserialize(r Reader) error {

	if err := m.ServerSessionConstant.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type GameFightFighterTaxCollectorLightInformations struct {
	GameFightFighterLightInformations

	FirstNameId uint16

	LastNameId uint16
}

func (m *GameFightFighterTaxCollectorLightInformations) ID() uint16 {
	return 457
}

func (m *GameFightFighterTaxCollectorLightInformations) Serialize(w Writer) error {

	if err := m.GameFightFighterLightInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.FirstNameId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.LastNameId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightFighterTaxCollectorLightInformations) Deserialize(r Reader) error {

	if err := m.GameFightFighterLightInformations.Deserialize(r); err != nil {
		return err
	}

	lfirstNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.FirstNameId = lfirstNameId

	llastNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.LastNameId = llastNameId

	return nil
}

type FightTeamMemberWithAllianceCharacterInformations struct {
	FightTeamMemberCharacterInformations

	AllianceInfos *BasicAllianceInformations
}

func (m *FightTeamMemberWithAllianceCharacterInformations) ID() uint16 {
	return 426
}

func (m *FightTeamMemberWithAllianceCharacterInformations) Serialize(w Writer) error {

	if err := m.FightTeamMemberCharacterInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.AllianceInfos.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *FightTeamMemberWithAllianceCharacterInformations) Deserialize(r Reader) error {

	if err := m.FightTeamMemberCharacterInformations.Deserialize(r); err != nil {
		return err
	}

	var lallianceInfos BasicAllianceInformations

	lallianceInfos.Deserialize(r)

	m.AllianceInfos = &lallianceInfos

	return nil
}

type CharacterHardcoreOrEpicInformations struct {
	CharacterBaseInformations

	DeathState uint8

	DeathCount uint16

	DeathMaxLevel uint8
}

func (m *CharacterHardcoreOrEpicInformations) ID() uint16 {
	return 474
}

func (m *CharacterHardcoreOrEpicInformations) Serialize(w Writer) error {

	if err := m.CharacterBaseInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.DeathState); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DeathCount); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.DeathMaxLevel); err != nil {
		return err
	}

	return nil
}

func (m *CharacterHardcoreOrEpicInformations) Deserialize(r Reader) error {

	if err := m.CharacterBaseInformations.Deserialize(r); err != nil {
		return err
	}

	ldeathState, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.DeathState = ldeathState

	ldeathCount, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DeathCount = ldeathCount

	ldeathMaxLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.DeathMaxLevel = ldeathMaxLevel

	return nil
}

type GameRolePlayTreasureHintInformations struct {
	GameRolePlayActorInformations

	NpcId uint16
}

func (m *GameRolePlayTreasureHintInformations) ID() uint16 {
	return 471
}

func (m *GameRolePlayTreasureHintInformations) Serialize(w Writer) error {

	if err := m.GameRolePlayActorInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NpcId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayTreasureHintInformations) Deserialize(r Reader) error {

	if err := m.GameRolePlayActorInformations.Deserialize(r); err != nil {
		return err
	}

	lnpcId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NpcId = lnpcId

	return nil
}

type ObjectEffectDuration struct {
	ObjectEffect

	Days uint16

	Hours uint8

	Minutes uint8
}

func (m *ObjectEffectDuration) ID() uint16 {
	return 75
}

func (m *ObjectEffectDuration) Serialize(w Writer) error {

	if err := m.ObjectEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Days); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Hours); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Minutes); err != nil {
		return err
	}

	return nil
}

func (m *ObjectEffectDuration) Deserialize(r Reader) error {

	if err := m.ObjectEffect.Deserialize(r); err != nil {
		return err
	}

	ldays, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Days = ldays

	lhours, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Hours = lhours

	lminutes, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Minutes = lminutes

	return nil
}

type GameFightFighterMonsterLightInformations struct {
	GameFightFighterLightInformations

	CreatureGenericId uint16
}

func (m *GameFightFighterMonsterLightInformations) ID() uint16 {
	return 455
}

func (m *GameFightFighterMonsterLightInformations) Serialize(w Writer) error {

	if err := m.GameFightFighterLightInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.CreatureGenericId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightFighterMonsterLightInformations) Deserialize(r Reader) error {

	if err := m.GameFightFighterLightInformations.Deserialize(r); err != nil {
		return err
	}

	lcreatureGenericId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CreatureGenericId = lcreatureGenericId

	return nil
}

type HumanOptionGuild struct {
	HumanOption

	GuildInformations *GuildInformations
}

func (m *HumanOptionGuild) ID() uint16 {
	return 409
}

func (m *HumanOptionGuild) Serialize(w Writer) error {

	if err := m.HumanOption.Serialize(w); err != nil {
		return err
	}

	if err := m.GuildInformations.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *HumanOptionGuild) Deserialize(r Reader) error {

	if err := m.HumanOption.Deserialize(r); err != nil {
		return err
	}

	var lguildInformations GuildInformations

	lguildInformations.Deserialize(r)

	m.GuildInformations = &lguildInformations

	return nil
}

type HouseGuildedInformations struct {
	HouseInstanceInformations

	GuildInfo *GuildInformations
}

func (m *HouseGuildedInformations) ID() uint16 {
	return 512
}

func (m *HouseGuildedInformations) Serialize(w Writer) error {

	if err := m.HouseInstanceInformations.Serialize(w); err != nil {
		return err
	}

	if err := m.GuildInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *HouseGuildedInformations) Deserialize(r Reader) error {

	if err := m.HouseInstanceInformations.Deserialize(r); err != nil {
		return err
	}

	var lguildInfo GuildInformations

	lguildInfo.Deserialize(r)

	m.GuildInfo = &lguildInfo

	return nil
}

type ObjectEffectDate struct {
	ObjectEffect

	Year uint16

	Month uint8

	Day uint8

	Hour uint8

	Minute uint8
}

func (m *ObjectEffectDate) ID() uint16 {
	return 72
}

func (m *ObjectEffectDate) Serialize(w Writer) error {

	if err := m.ObjectEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Year); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Month); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Day); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Hour); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Minute); err != nil {
		return err
	}

	return nil
}

func (m *ObjectEffectDate) Deserialize(r Reader) error {

	if err := m.ObjectEffect.Deserialize(r); err != nil {
		return err
	}

	lyear, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Year = lyear

	lmonth, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Month = lmonth

	lday, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Day = lday

	lhour, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Hour = lhour

	lminute, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Minute = lminute

	return nil
}

type ObjectEffectLadder struct {
	ObjectEffectCreature

	MonsterCount uint32
}

func (m *ObjectEffectLadder) ID() uint16 {
	return 81
}

func (m *ObjectEffectLadder) Serialize(w Writer) error {

	if err := m.ObjectEffectCreature.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MonsterCount); err != nil {
		return err
	}

	return nil
}

func (m *ObjectEffectLadder) Deserialize(r Reader) error {

	if err := m.ObjectEffectCreature.Deserialize(r); err != nil {
		return err
	}

	lmonsterCount, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MonsterCount = lmonsterCount

	return nil
}

type GameFightFighterCompanionLightInformations struct {
	GameFightFighterLightInformations

	CompanionId uint8

	MasterId float64
}

func (m *GameFightFighterCompanionLightInformations) ID() uint16 {
	return 454
}

func (m *GameFightFighterCompanionLightInformations) Serialize(w Writer) error {

	if err := m.GameFightFighterLightInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.CompanionId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.MasterId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightFighterCompanionLightInformations) Deserialize(r Reader) error {

	if err := m.GameFightFighterLightInformations.Deserialize(r); err != nil {
		return err
	}

	lcompanionId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CompanionId = lcompanionId

	lmasterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.MasterId = lmasterId

	return nil
}

type ObjectEffectDice struct {
	ObjectEffect

	DiceNum uint16

	DiceSide uint16

	DiceConst uint16
}

func (m *ObjectEffectDice) ID() uint16 {
	return 73
}

func (m *ObjectEffectDice) Serialize(w Writer) error {

	if err := m.ObjectEffect.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DiceNum); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DiceSide); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DiceConst); err != nil {
		return err
	}

	return nil
}

func (m *ObjectEffectDice) Deserialize(r Reader) error {

	if err := m.ObjectEffect.Deserialize(r); err != nil {
		return err
	}

	ldiceNum, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DiceNum = ldiceNum

	ldiceSide, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DiceSide = ldiceSide

	ldiceConst, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DiceConst = ldiceConst

	return nil
}

type TreasureHuntStepDig struct {
	TreasureHuntStep
}

func (m *TreasureHuntStepDig) ID() uint16 {
	return 465
}

func (m *TreasureHuntStepDig) Serialize(w Writer) error {

	if err := m.TreasureHuntStep.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntStepDig) Deserialize(r Reader) error {

	if err := m.TreasureHuntStep.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type TreasureHuntStepFollowDirection struct {
	TreasureHuntStep

	Direction uint8

	MapCount uint16
}

func (m *TreasureHuntStepFollowDirection) ID() uint16 {
	return 468
}

func (m *TreasureHuntStepFollowDirection) Serialize(w Writer) error {

	if err := m.TreasureHuntStep.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Direction); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MapCount); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntStepFollowDirection) Deserialize(r Reader) error {

	if err := m.TreasureHuntStep.Deserialize(r); err != nil {
		return err
	}

	ldirection, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Direction = ldirection

	lmapCount, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MapCount = lmapCount

	return nil
}

type PaddockGuildedInformations struct {
	PaddockBuyableInformations

	Deserted bool

	GuildInfo *GuildInformations
}

func (m *PaddockGuildedInformations) ID() uint16 {
	return 508
}

func (m *PaddockGuildedInformations) Serialize(w Writer) error {

	if err := m.PaddockBuyableInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Deserted); err != nil {
		return err
	}

	if err := m.GuildInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PaddockGuildedInformations) Deserialize(r Reader) error {

	if err := m.PaddockBuyableInformations.Deserialize(r); err != nil {
		return err
	}

	ldeserted, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Deserted = ldeserted

	var lguildInfo GuildInformations

	lguildInfo.Deserialize(r)

	m.GuildInfo = &lguildInfo

	return nil
}

type GuildInAllianceVersatileInformations struct {
	GuildVersatileInformations

	AllianceId uint32
}

func (m *GuildInAllianceVersatileInformations) ID() uint16 {
	return 437
}

func (m *GuildInAllianceVersatileInformations) Serialize(w Writer) error {

	if err := m.GuildVersatileInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.AllianceId); err != nil {
		return err
	}

	return nil
}

func (m *GuildInAllianceVersatileInformations) Deserialize(r Reader) error {

	if err := m.GuildVersatileInformations.Deserialize(r); err != nil {
		return err
	}

	lallianceId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.AllianceId = lallianceId

	return nil
}

type GameFightFighterNamedLightInformations struct {
	GameFightFighterLightInformations

	Name string
}

func (m *GameFightFighterNamedLightInformations) ID() uint16 {
	return 456
}

func (m *GameFightFighterNamedLightInformations) Serialize(w Writer) error {

	if err := m.GameFightFighterLightInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GameFightFighterNamedLightInformations) Deserialize(r Reader) error {

	if err := m.GameFightFighterLightInformations.Deserialize(r); err != nil {
		return err
	}

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	return nil
}

type FightResultExperienceData struct {
	FightResultAdditionalData

	Experience int64

	ShowExperience bool

	ExperienceLevelFloor int64

	ShowExperienceLevelFloor bool

	ExperienceNextLevelFloor int64

	ShowExperienceNextLevelFloor bool

	ExperienceFightDelta int64

	ShowExperienceFightDelta bool

	ExperienceForGuild int64

	ShowExperienceForGuild bool

	ExperienceForMount int64

	ShowExperienceForMount bool

	IsIncarnationExperience bool

	RerollExperienceMul uint8
}

func (m *FightResultExperienceData) ID() uint16 {
	return 192
}

func (m *FightResultExperienceData) Serialize(w Writer) error {

	if err := m.FightResultAdditionalData.Serialize(w); err != nil {
		return err
	}

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.ShowExperience)

	setWrappedFlag(bbw0, 1, m.ShowExperienceLevelFloor)

	setWrappedFlag(bbw0, 2, m.ShowExperienceNextLevelFloor)

	setWrappedFlag(bbw0, 3, m.ShowExperienceFightDelta)

	setWrappedFlag(bbw0, 4, m.ShowExperienceForGuild)

	setWrappedFlag(bbw0, 5, m.ShowExperienceForMount)

	setWrappedFlag(bbw0, 6, m.IsIncarnationExperience)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Experience); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceLevelFloor); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceNextLevelFloor); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceFightDelta); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceForGuild); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceForMount); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.RerollExperienceMul); err != nil {
		return err
	}

	return nil
}

func (m *FightResultExperienceData) Deserialize(r Reader) error {

	if err := m.FightResultAdditionalData.Deserialize(r); err != nil {
		return err
	}

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ShowExperience = getWrappedFlag(bbw0, 0)

	m.ShowExperienceLevelFloor = getWrappedFlag(bbw0, 1)

	m.ShowExperienceNextLevelFloor = getWrappedFlag(bbw0, 2)

	m.ShowExperienceFightDelta = getWrappedFlag(bbw0, 3)

	m.ShowExperienceForGuild = getWrappedFlag(bbw0, 4)

	m.ShowExperienceForMount = getWrappedFlag(bbw0, 5)

	m.IsIncarnationExperience = getWrappedFlag(bbw0, 6)

	lexperience, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Experience = lexperience

	lexperienceLevelFloor, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceLevelFloor = lexperienceLevelFloor

	lexperienceNextLevelFloor, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceNextLevelFloor = lexperienceNextLevelFloor

	lexperienceFightDelta, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceFightDelta = lexperienceFightDelta

	lexperienceForGuild, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceForGuild = lexperienceForGuild

	lexperienceForMount, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceForMount = lexperienceForMount

	lrerollExperienceMul, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.RerollExperienceMul = lrerollExperienceMul

	return nil
}

type FightTeamMemberCompanionInformations struct {
	FightTeamMemberInformations

	CompanionId uint8

	Level uint8

	MasterId float64
}

func (m *FightTeamMemberCompanionInformations) ID() uint16 {
	return 451
}

func (m *FightTeamMemberCompanionInformations) Serialize(w Writer) error {

	if err := m.FightTeamMemberInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.CompanionId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	if err := w.WriteDouble(m.MasterId); err != nil {
		return err
	}

	return nil
}

func (m *FightTeamMemberCompanionInformations) Deserialize(r Reader) error {

	if err := m.FightTeamMemberInformations.Deserialize(r); err != nil {
		return err
	}

	lcompanionId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CompanionId = lcompanionId

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	lmasterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.MasterId = lmasterId

	return nil
}

type InteractiveElementNamedSkill struct {
	InteractiveElementSkill

	NameId uint32
}

func (m *InteractiveElementNamedSkill) ID() uint16 {
	return 220
}

func (m *InteractiveElementNamedSkill) Serialize(w Writer) error {

	if err := m.InteractiveElementSkill.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.NameId); err != nil {
		return err
	}

	return nil
}

func (m *InteractiveElementNamedSkill) Deserialize(r Reader) error {

	if err := m.InteractiveElementSkill.Deserialize(r); err != nil {
		return err
	}

	lnameId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.NameId = lnameId

	return nil
}

type StatisticDataShort struct {
	StatisticData

	Value int16
}

func (m *StatisticDataShort) ID() uint16 {
	return 488
}

func (m *StatisticDataShort) Serialize(w Writer) error {

	if err := m.StatisticData.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *StatisticDataShort) Deserialize(r Reader) error {

	if err := m.StatisticData.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type FightResultMutantListEntry struct {
	FightResultFighterListEntry

	Level uint16
}

func (m *FightResultMutantListEntry) ID() uint16 {
	return 216
}

func (m *FightResultMutantListEntry) Serialize(w Writer) error {

	if err := m.FightResultFighterListEntry.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Level); err != nil {
		return err
	}

	return nil
}

func (m *FightResultMutantListEntry) Deserialize(r Reader) error {

	if err := m.FightResultFighterListEntry.Deserialize(r); err != nil {
		return err
	}

	llevel, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Level = llevel

	return nil
}

type ServerSessionConstantInteger struct {
	ServerSessionConstant

	Value int32
}

func (m *ServerSessionConstantInteger) ID() uint16 {
	return 433
}

func (m *ServerSessionConstantInteger) Serialize(w Writer) error {

	if err := m.ServerSessionConstant.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *ServerSessionConstantInteger) Deserialize(r Reader) error {

	if err := m.ServerSessionConstant.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type TreasureHuntStepFollowDirectionToPOI struct {
	TreasureHuntStep

	Direction uint8

	PoiLabelId uint16
}

func (m *TreasureHuntStepFollowDirectionToPOI) ID() uint16 {
	return 461
}

func (m *TreasureHuntStepFollowDirectionToPOI) Serialize(w Writer) error {

	if err := m.TreasureHuntStep.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Direction); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.PoiLabelId); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntStepFollowDirectionToPOI) Deserialize(r Reader) error {

	if err := m.TreasureHuntStep.Deserialize(r); err != nil {
		return err
	}

	ldirection, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Direction = ldirection

	lpoiLabelId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.PoiLabelId = lpoiLabelId

	return nil
}

type CharacterMinimalPlusLookAndGradeInformations struct {
	CharacterMinimalPlusLookInformations

	Grade uint32
}

func (m *CharacterMinimalPlusLookAndGradeInformations) ID() uint16 {
	return 193
}

func (m *CharacterMinimalPlusLookAndGradeInformations) Serialize(w Writer) error {

	if err := m.CharacterMinimalPlusLookInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Grade); err != nil {
		return err
	}

	return nil
}

func (m *CharacterMinimalPlusLookAndGradeInformations) Deserialize(r Reader) error {

	if err := m.CharacterMinimalPlusLookInformations.Deserialize(r); err != nil {
		return err
	}

	lgrade, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Grade = lgrade

	return nil
}

type TreasureHuntStepFollowDirectionToHint struct {
	TreasureHuntStep

	Direction uint8

	NpcId uint16
}

func (m *TreasureHuntStepFollowDirectionToHint) ID() uint16 {
	return 472
}

func (m *TreasureHuntStepFollowDirectionToHint) Serialize(w Writer) error {

	if err := m.TreasureHuntStep.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Direction); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NpcId); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntStepFollowDirectionToHint) Deserialize(r Reader) error {

	if err := m.TreasureHuntStep.Deserialize(r); err != nil {
		return err
	}

	ldirection, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Direction = ldirection

	lnpcId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NpcId = lnpcId

	return nil
}

type QuestObjectiveInformationsWithCompletion struct {
	QuestObjectiveInformations

	CurCompletion uint16

	MaxCompletion uint16
}

func (m *QuestObjectiveInformationsWithCompletion) ID() uint16 {
	return 386
}

func (m *QuestObjectiveInformationsWithCompletion) Serialize(w Writer) error {

	if err := m.QuestObjectiveInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.CurCompletion); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MaxCompletion); err != nil {
		return err
	}

	return nil
}

func (m *QuestObjectiveInformationsWithCompletion) Deserialize(r Reader) error {

	if err := m.QuestObjectiveInformations.Deserialize(r); err != nil {
		return err
	}

	lcurCompletion, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CurCompletion = lcurCompletion

	lmaxCompletion, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MaxCompletion = lmaxCompletion

	return nil
}

type TreasureHuntFlag struct {
	MapId int32

	State uint8
}

func (m *TreasureHuntFlag) ID() uint16 {
	return 473
}

func (m *TreasureHuntFlag) Serialize(w Writer) error {

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.State); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntFlag) Deserialize(r Reader) error {

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lstate, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.State = lstate

	return nil
}

type JobExperience struct {
	JobId uint8

	JobLevel uint8

	JobXP int64

	JobXpLevelFloor int64

	JobXpNextLevelFloor int64
}

func (m *JobExperience) ID() uint16 {
	return 98
}

func (m *JobExperience) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.JobId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.JobLevel); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.JobXP); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.JobXpLevelFloor); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.JobXpNextLevelFloor); err != nil {
		return err
	}

	return nil
}

func (m *JobExperience) Deserialize(r Reader) error {

	ljobId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.JobId = ljobId

	ljobLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.JobLevel = ljobLevel

	ljobXP, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.JobXP = ljobXP

	ljobXpLevelFloor, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.JobXpLevelFloor = ljobXpLevelFloor

	ljobXpNextLevelFloor, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.JobXpNextLevelFloor = ljobXpNextLevelFloor

	return nil
}

type JobBookSubscription struct {
	JobId uint8

	Subscribed bool
}

func (m *JobBookSubscription) ID() uint16 {
	return 500
}

func (m *JobBookSubscription) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.JobId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Subscribed); err != nil {
		return err
	}

	return nil
}

func (m *JobBookSubscription) Deserialize(r Reader) error {

	ljobId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.JobId = ljobId

	lsubscribed, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Subscribed = lsubscribed

	return nil
}

type JobCrafterDirectorySettings struct {
	JobId uint8

	MinLevel uint8

	Free bool
}

func (m *JobCrafterDirectorySettings) ID() uint16 {
	return 97
}

func (m *JobCrafterDirectorySettings) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.JobId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.MinLevel); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Free); err != nil {
		return err
	}

	return nil
}

func (m *JobCrafterDirectorySettings) Deserialize(r Reader) error {

	ljobId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.JobId = ljobId

	lminLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MinLevel = lminLevel

	lfree, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Free = lfree

	return nil
}

type PresetItem struct {
	Position uint8

	ObjGid uint16

	ObjUid uint32
}

func (m *PresetItem) ID() uint16 {
	return 354
}

func (m *PresetItem) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Position); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjGid); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ObjUid); err != nil {
		return err
	}

	return nil
}

func (m *PresetItem) Deserialize(r Reader) error {

	lposition, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Position = lposition

	lobjGid, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjGid = lobjGid

	lobjUid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjUid = lobjUid

	return nil
}

type CharacterRemodelingInformation struct {
	AbstractCharacterInformation

	Name string

	Breed int8

	Sex bool

	CosmeticId uint16

	Colors []int32
}

func (m *CharacterRemodelingInformation) ID() uint16 {
	return 479
}

func (m *CharacterRemodelingInformation) Serialize(w Writer) error {

	if err := m.AbstractCharacterInformation.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Sex); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.CosmeticId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Colors))); err != nil {
		return err
	}

	for i := range m.Colors {

		if err := w.WriteInt32(m.Colors[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *CharacterRemodelingInformation) Deserialize(r Reader) error {

	if err := m.AbstractCharacterInformation.Deserialize(r); err != nil {
		return err
	}

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	lsex, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Sex = lsex

	lcosmeticId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CosmeticId = lcosmeticId

	lcolorsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Colors = make([]int32, lcolorsLen)

	for i := range m.Colors {

		lcolors, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.Colors[i] = lcolors

	}

	return nil
}

type CharacterToRemodelInformations struct {
	CharacterRemodelingInformation

	PossibleChangeMask uint8

	MandatoryChangeMask uint8
}

func (m *CharacterToRemodelInformations) ID() uint16 {
	return 477
}

func (m *CharacterToRemodelInformations) Serialize(w Writer) error {

	if err := m.CharacterRemodelingInformation.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.PossibleChangeMask); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.MandatoryChangeMask); err != nil {
		return err
	}

	return nil
}

func (m *CharacterToRemodelInformations) Deserialize(r Reader) error {

	if err := m.CharacterRemodelingInformation.Deserialize(r); err != nil {
		return err
	}

	lpossibleChangeMask, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PossibleChangeMask = lpossibleChangeMask

	lmandatoryChangeMask, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MandatoryChangeMask = lmandatoryChangeMask

	return nil
}

type RemodelingInformation struct {
	Name string

	Breed int8

	Sex bool

	CosmeticId uint16

	Colors []int32
}

func (m *RemodelingInformation) ID() uint16 {
	return 480
}

func (m *RemodelingInformation) Serialize(w Writer) error {

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Sex); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.CosmeticId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Colors))); err != nil {
		return err
	}

	for i := range m.Colors {

		if err := w.WriteInt32(m.Colors[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *RemodelingInformation) Deserialize(r Reader) error {

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	lsex, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Sex = lsex

	lcosmeticId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CosmeticId = lcosmeticId

	lcolorsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Colors = make([]int32, lcolorsLen)

	for i := range m.Colors {

		lcolors, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.Colors[i] = lcolors

	}

	return nil
}

type Version struct {
	Major uint8

	Minor uint8

	Release uint8

	Revision uint32

	Patch uint8

	BuildType uint8
}

func (m *Version) ID() uint16 {
	return 11
}

func (m *Version) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Major); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Minor); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Release); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.Revision); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Patch); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.BuildType); err != nil {
		return err
	}

	return nil
}

func (m *Version) Deserialize(r Reader) error {

	lmajor, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Major = lmajor

	lminor, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Minor = lminor

	lrelease, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Release = lrelease

	lrevision, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.Revision = lrevision

	lpatch, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Patch = lpatch

	lbuildType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BuildType = lbuildType

	return nil
}

type VersionExtended struct {
	Version

	Install uint8

	Technology uint8
}

func (m *VersionExtended) ID() uint16 {
	return 393
}

func (m *VersionExtended) Serialize(w Writer) error {

	if err := m.Version.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Install); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Technology); err != nil {
		return err
	}

	return nil
}

func (m *VersionExtended) Deserialize(r Reader) error {

	if err := m.Version.Deserialize(r); err != nil {
		return err
	}

	linstall, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Install = linstall

	ltechnology, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Technology = ltechnology

	return nil
}

type MonsterBoosts struct {
	Id uint32

	XpBoost uint16

	DropBoost uint16
}

func (m *MonsterBoosts) ID() uint16 {
	return 497
}

func (m *MonsterBoosts) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.Id); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.XpBoost); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DropBoost); err != nil {
		return err
	}

	return nil
}

func (m *MonsterBoosts) Deserialize(r Reader) error {

	lid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Id = lid

	lxpBoost, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.XpBoost = lxpBoost

	ldropBoost, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DropBoost = ldropBoost

	return nil
}

type FinishMoveInformations struct {
	FinishMoveId uint32

	FinishMoveState bool
}

func (m *FinishMoveInformations) ID() uint16 {
	return 506
}

func (m *FinishMoveInformations) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.FinishMoveId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.FinishMoveState); err != nil {
		return err
	}

	return nil
}

func (m *FinishMoveInformations) Deserialize(r Reader) error {

	lfinishMoveId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.FinishMoveId = lfinishMoveId

	lfinishMoveState, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.FinishMoveState = lfinishMoveState

	return nil
}

type SpellItem struct {
	Item

	SpellId int32

	SpellLevel int16
}

func (m *SpellItem) ID() uint16 {
	return 49
}

func (m *SpellItem) Serialize(w Writer) error {

	if err := m.Item.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.SpellId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.SpellLevel); err != nil {
		return err
	}

	return nil
}

func (m *SpellItem) Deserialize(r Reader) error {

	if err := m.Item.Deserialize(r); err != nil {
		return err
	}

	lspellId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	lspellLevel, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SpellLevel = lspellLevel

	return nil
}

type ArenaRankInfos struct {
	Rank uint16

	BestRank uint16

	VictoryCount uint16

	Fightcount uint16
}

func (m *ArenaRankInfos) ID() uint16 {
	return 499
}

func (m *ArenaRankInfos) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Rank); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.BestRank); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.VictoryCount); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Fightcount); err != nil {
		return err
	}

	return nil
}

func (m *ArenaRankInfos) Deserialize(r Reader) error {

	lrank, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Rank = lrank

	lbestRank, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.BestRank = lbestRank

	lvictoryCount, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.VictoryCount = lvictoryCount

	lfightcount, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Fightcount = lfightcount

	return nil
}

type PartyMemberGeoPosition struct {
	MemberId uint32

	WorldX int16

	WorldY int16

	MapId int32

	SubAreaId uint16
}

func (m *PartyMemberGeoPosition) ID() uint16 {
	return 378
}

func (m *PartyMemberGeoPosition) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.MemberId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	return nil
}

func (m *PartyMemberGeoPosition) Deserialize(r Reader) error {

	lmemberId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MemberId = lmemberId

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	return nil
}

type FightOptionsInformations struct {
	IsSecret bool

	IsRestrictedToPartyOnly bool

	IsClosed bool

	IsAskingForHelp bool
}

func (m *FightOptionsInformations) ID() uint16 {
	return 20
}

func (m *FightOptionsInformations) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.IsSecret)

	setWrappedFlag(bbw0, 1, m.IsRestrictedToPartyOnly)

	setWrappedFlag(bbw0, 2, m.IsClosed)

	setWrappedFlag(bbw0, 3, m.IsAskingForHelp)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	return nil
}

func (m *FightOptionsInformations) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.IsSecret = getWrappedFlag(bbw0, 0)

	m.IsRestrictedToPartyOnly = getWrappedFlag(bbw0, 1)

	m.IsClosed = getWrappedFlag(bbw0, 2)

	m.IsAskingForHelp = getWrappedFlag(bbw0, 3)

	return nil
}

type FightStartingPositions struct {
	PositionsForChallengers []uint16

	PositionsForDefenders []uint16
}

func (m *FightStartingPositions) ID() uint16 {
	return 513
}

func (m *FightStartingPositions) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.PositionsForChallengers))); err != nil {
		return err
	}

	for i := range m.PositionsForChallengers {

		if err := w.WriteVarUInt16(m.PositionsForChallengers[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.PositionsForDefenders))); err != nil {
		return err
	}

	for i := range m.PositionsForDefenders {

		if err := w.WriteVarUInt16(m.PositionsForDefenders[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *FightStartingPositions) Deserialize(r Reader) error {

	lpositionsForChallengersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PositionsForChallengers = make([]uint16, lpositionsForChallengersLen)

	for i := range m.PositionsForChallengers {

		lpositionsForChallengers, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.PositionsForChallengers[i] = lpositionsForChallengers

	}

	lpositionsForDefendersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PositionsForDefenders = make([]uint16, lpositionsForDefendersLen)

	for i := range m.PositionsForDefenders {

		lpositionsForDefenders, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.PositionsForDefenders[i] = lpositionsForDefenders

	}

	return nil
}

type TaxCollectorBasicInformations struct {
	FirstNameId uint16

	LastNameId uint16

	WorldX int16

	WorldY int16

	MapId int32

	SubAreaId uint16
}

func (m *TaxCollectorBasicInformations) ID() uint16 {
	return 96
}

func (m *TaxCollectorBasicInformations) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.FirstNameId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.LastNameId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorBasicInformations) Deserialize(r Reader) error {

	lfirstNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.FirstNameId = lfirstNameId

	llastNameId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.LastNameId = llastNameId

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	return nil
}

type MountInformationsForPaddock struct {
	ModelId uint16

	Name string

	OwnerName string
}

func (m *MountInformationsForPaddock) ID() uint16 {
	return 184
}

func (m *MountInformationsForPaddock) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ModelId); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteString(m.OwnerName); err != nil {
		return err
	}

	return nil
}

func (m *MountInformationsForPaddock) Deserialize(r Reader) error {

	lmodelId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ModelId = lmodelId

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	lownerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.OwnerName = lownerName

	return nil
}

type FightLoot struct {
	Objects []uint16

	Kamas int64
}

func (m *FightLoot) ID() uint16 {
	return 41
}

func (m *FightLoot) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Objects))); err != nil {
		return err
	}

	for i := range m.Objects {

		if err := w.WriteVarUInt16(m.Objects[i]); err != nil {
			return err
		}

	}

	if err := w.WriteVarInt64(m.Kamas); err != nil {
		return err
	}

	return nil
}

func (m *FightLoot) Deserialize(r Reader) error {

	lobjectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Objects = make([]uint16, lobjectsLen)

	for i := range m.Objects {

		lobjects, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Objects[i] = lobjects

	}

	lkamas, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Kamas = lkamas

	return nil
}

type ObjectItemToSell struct {
	Item

	ObjectGID uint16

	Effects []*ObjectEffect

	ObjectUID uint32

	Quantity uint32

	ObjectPrice int64
}

func (m *ObjectItemToSell) ID() uint16 {
	return 120
}

func (m *ObjectItemToSell) Serialize(w Writer) error {

	if err := m.Item.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGID); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Effects))); err != nil {
		return err
	}

	for i := range m.Effects {

		if err := w.WriteUInt16(m.Effects[i].ID()); err != nil {
			return err
		}

		if err := m.Effects[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ObjectPrice); err != nil {
		return err
	}

	return nil
}

func (m *ObjectItemToSell) Deserialize(r Reader) error {

	if err := m.Item.Deserialize(r); err != nil {
		return err
	}

	lobjectGID, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGID = lobjectGID

	leffectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Effects = make([]*ObjectEffect, leffectsLen)

	for i := range m.Effects {

		typeeffectsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		leffects, err := GetType(typeeffectsID)
		if err != nil {
			return err
		}

		leffects.Deserialize(r)

		m.Effects[i] = leffects.(*ObjectEffect)

	}

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	lobjectPrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ObjectPrice = lobjectPrice

	return nil
}

type ObjectItemToSellInHumanVendorShop struct {
	Item

	ObjectGID uint16

	Effects []*ObjectEffect

	ObjectUID uint32

	Quantity uint32

	ObjectPrice int64

	PublicPrice int64
}

func (m *ObjectItemToSellInHumanVendorShop) ID() uint16 {
	return 359
}

func (m *ObjectItemToSellInHumanVendorShop) Serialize(w Writer) error {

	if err := m.Item.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGID); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Effects))); err != nil {
		return err
	}

	for i := range m.Effects {

		if err := w.WriteUInt16(m.Effects[i].ID()); err != nil {
			return err
		}

		if err := m.Effects[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ObjectPrice); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PublicPrice); err != nil {
		return err
	}

	return nil
}

func (m *ObjectItemToSellInHumanVendorShop) Deserialize(r Reader) error {

	if err := m.Item.Deserialize(r); err != nil {
		return err
	}

	lobjectGID, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGID = lobjectGID

	leffectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Effects = make([]*ObjectEffect, leffectsLen)

	for i := range m.Effects {

		typeeffectsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		leffects, err := GetType(typeeffectsID)
		if err != nil {
			return err
		}

		leffects.Deserialize(r)

		m.Effects[i] = leffects.(*ObjectEffect)

	}

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	lobjectPrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ObjectPrice = lobjectPrice

	lpublicPrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PublicPrice = lpublicPrice

	return nil
}

type ObjectItemToSellInBid struct {
	ObjectItemToSell

	UnsoldDelay uint32
}

func (m *ObjectItemToSellInBid) ID() uint16 {
	return 164
}

func (m *ObjectItemToSellInBid) Serialize(w Writer) error {

	if err := m.ObjectItemToSell.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.UnsoldDelay); err != nil {
		return err
	}

	return nil
}

func (m *ObjectItemToSellInBid) Deserialize(r Reader) error {

	if err := m.ObjectItemToSell.Deserialize(r); err != nil {
		return err
	}

	lunsoldDelay, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.UnsoldDelay = lunsoldDelay

	return nil
}

type BidExchangerObjectInfo struct {
	ObjectUID uint32

	Effects []*ObjectEffect

	Prices []int64
}

func (m *BidExchangerObjectInfo) ID() uint16 {
	return 122
}

func (m *BidExchangerObjectInfo) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Effects))); err != nil {
		return err
	}

	for i := range m.Effects {

		if err := w.WriteUInt16(m.Effects[i].ID()); err != nil {
			return err
		}

		if err := m.Effects[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Prices))); err != nil {
		return err
	}

	for i := range m.Prices {

		if err := w.WriteVarInt64(m.Prices[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *BidExchangerObjectInfo) Deserialize(r Reader) error {

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	leffectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Effects = make([]*ObjectEffect, leffectsLen)

	for i := range m.Effects {

		typeeffectsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		leffects, err := GetType(typeeffectsID)
		if err != nil {
			return err
		}

		leffects.Deserialize(r)

		m.Effects[i] = leffects.(*ObjectEffect)

	}

	lpricesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Prices = make([]int64, lpricesLen)

	for i := range m.Prices {

		lprices, err := r.ReadVarInt64()
		if err != nil {
			return err
		}

		m.Prices[i] = lprices

	}

	return nil
}

type SellerBuyerDescriptor struct {
	Quantities []uint32

	Types []uint32

	TaxPercentage float32

	TaxModificationPercentage float32

	MaxItemLevel uint8

	MaxItemPerAccount uint32

	NpcContextualId int32

	UnsoldDelay uint16
}

func (m *SellerBuyerDescriptor) ID() uint16 {
	return 121
}

func (m *SellerBuyerDescriptor) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Quantities))); err != nil {
		return err
	}

	for i := range m.Quantities {

		if err := w.WriteVarUInt32(m.Quantities[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Types))); err != nil {
		return err
	}

	for i := range m.Types {

		if err := w.WriteVarUInt32(m.Types[i]); err != nil {
			return err
		}

	}

	if err := w.WriteFloat(m.TaxPercentage); err != nil {
		return err
	}

	if err := w.WriteFloat(m.TaxModificationPercentage); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.MaxItemLevel); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MaxItemPerAccount); err != nil {
		return err
	}

	if err := w.WriteInt32(m.NpcContextualId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.UnsoldDelay); err != nil {
		return err
	}

	return nil
}

func (m *SellerBuyerDescriptor) Deserialize(r Reader) error {

	lquantitiesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Quantities = make([]uint32, lquantitiesLen)

	for i := range m.Quantities {

		lquantities, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.Quantities[i] = lquantities

	}

	ltypesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Types = make([]uint32, ltypesLen)

	for i := range m.Types {

		ltypes, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.Types[i] = ltypes

	}

	ltaxPercentage, err := r.ReadFloat()
	if err != nil {
		return err
	}

	m.TaxPercentage = ltaxPercentage

	ltaxModificationPercentage, err := r.ReadFloat()
	if err != nil {
		return err
	}

	m.TaxModificationPercentage = ltaxModificationPercentage

	lmaxItemLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MaxItemLevel = lmaxItemLevel

	lmaxItemPerAccount, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MaxItemPerAccount = lmaxItemPerAccount

	lnpcContextualId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.NpcContextualId = lnpcContextualId

	lunsoldDelay, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.UnsoldDelay = lunsoldDelay

	return nil
}

type FightExternalInformations struct {
	FightId int32

	FightType uint8

	FightStart uint32

	FightSpectatorLocked bool

	FightTeams []*FightTeamLightInformations

	FightTeamsOptions []*FightOptionsInformations
}

func (m *FightExternalInformations) ID() uint16 {
	return 117
}

func (m *FightExternalInformations) Serialize(w Writer) error {

	if err := w.WriteInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.FightType); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.FightStart); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.FightSpectatorLocked); err != nil {
		return err
	}

	for i := range m.FightTeams {

		if err := m.FightTeams[i].Serialize(w); err != nil {
			return err
		}

	}

	for i := range m.FightTeamsOptions {

		if err := m.FightTeamsOptions[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *FightExternalInformations) Deserialize(r Reader) error {

	lfightId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lfightType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.FightType = lfightType

	lfightStart, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.FightStart = lfightStart

	lfightSpectatorLocked, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.FightSpectatorLocked = lfightSpectatorLocked

	m.FightTeams = make([]*FightTeamLightInformations, 2)

	for i := range m.FightTeams {

		var lfightTeams FightTeamLightInformations

		lfightTeams.Deserialize(r)

		m.FightTeams[i] = &lfightTeams

	}

	m.FightTeamsOptions = make([]*FightOptionsInformations, 2)

	for i := range m.FightTeamsOptions {

		var lfightTeamsOptions FightOptionsInformations

		lfightTeamsOptions.Deserialize(r)

		m.FightTeamsOptions[i] = &lfightTeamsOptions

	}

	return nil
}

type ObjectItemToSellInNpcShop struct {
	ObjectItemMinimalInformation

	ObjectPrice int64

	BuyCriterion string
}

func (m *ObjectItemToSellInNpcShop) ID() uint16 {
	return 352
}

func (m *ObjectItemToSellInNpcShop) Serialize(w Writer) error {

	if err := m.ObjectItemMinimalInformation.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ObjectPrice); err != nil {
		return err
	}

	if err := w.WriteString(m.BuyCriterion); err != nil {
		return err
	}

	return nil
}

func (m *ObjectItemToSellInNpcShop) Deserialize(r Reader) error {

	if err := m.ObjectItemMinimalInformation.Deserialize(r); err != nil {
		return err
	}

	lobjectPrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ObjectPrice = lobjectPrice

	lbuyCriterion, err := r.ReadString()
	if err != nil {
		return err
	}

	m.BuyCriterion = lbuyCriterion

	return nil
}

type ProtectedEntityWaitingForHelpInfo struct {
	TimeLeftBeforeFight int32

	WaitTimeForPlacement int32

	NbPositionForDefensors uint8
}

func (m *ProtectedEntityWaitingForHelpInfo) ID() uint16 {
	return 186
}

func (m *ProtectedEntityWaitingForHelpInfo) Serialize(w Writer) error {

	if err := w.WriteInt32(m.TimeLeftBeforeFight); err != nil {
		return err
	}

	if err := w.WriteInt32(m.WaitTimeForPlacement); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.NbPositionForDefensors); err != nil {
		return err
	}

	return nil
}

func (m *ProtectedEntityWaitingForHelpInfo) Deserialize(r Reader) error {

	ltimeLeftBeforeFight, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.TimeLeftBeforeFight = ltimeLeftBeforeFight

	lwaitTimeForPlacement, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.WaitTimeForPlacement = lwaitTimeForPlacement

	lnbPositionForDefensors, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NbPositionForDefensors = lnbPositionForDefensors

	return nil
}

type DecraftedItemStackInfo struct {
	ObjectUID uint32

	BonusMin float32

	BonusMax float32

	RunesId []uint16

	RunesQty []uint32
}

func (m *DecraftedItemStackInfo) ID() uint16 {
	return 481
}

func (m *DecraftedItemStackInfo) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteFloat(m.BonusMin); err != nil {
		return err
	}

	if err := w.WriteFloat(m.BonusMax); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.RunesId))); err != nil {
		return err
	}

	for i := range m.RunesId {

		if err := w.WriteVarUInt16(m.RunesId[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.RunesQty))); err != nil {
		return err
	}

	for i := range m.RunesQty {

		if err := w.WriteVarUInt32(m.RunesQty[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *DecraftedItemStackInfo) Deserialize(r Reader) error {

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	lbonusMin, err := r.ReadFloat()
	if err != nil {
		return err
	}

	m.BonusMin = lbonusMin

	lbonusMax, err := r.ReadFloat()
	if err != nil {
		return err
	}

	m.BonusMax = lbonusMax

	lrunesIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.RunesId = make([]uint16, lrunesIdLen)

	for i := range m.RunesId {

		lrunesId, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.RunesId[i] = lrunesId

	}

	lrunesQtyLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.RunesQty = make([]uint32, lrunesQtyLen)

	for i := range m.RunesQty {

		lrunesQty, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.RunesQty[i] = lrunesQty

	}

	return nil
}

type Achievement struct {
	Id uint16

	FinishedObjective []*AchievementObjective

	StartedObjectives []*AchievementStartedObjective
}

func (m *Achievement) ID() uint16 {
	return 363
}

func (m *Achievement) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Id); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.FinishedObjective))); err != nil {
		return err
	}

	for i := range m.FinishedObjective {

		if err := m.FinishedObjective[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.StartedObjectives))); err != nil {
		return err
	}

	for i := range m.StartedObjectives {

		if err := m.StartedObjectives[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *Achievement) Deserialize(r Reader) error {

	lid, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Id = lid

	lfinishedObjectiveLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FinishedObjective = make([]*AchievementObjective, lfinishedObjectiveLen)

	for i := range m.FinishedObjective {

		var lfinishedObjective AchievementObjective

		lfinishedObjective.Deserialize(r)

		m.FinishedObjective[i] = &lfinishedObjective

	}

	lstartedObjectivesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.StartedObjectives = make([]*AchievementStartedObjective, lstartedObjectivesLen)

	for i := range m.StartedObjectives {

		var lstartedObjectives AchievementStartedObjective

		lstartedObjectives.Deserialize(r)

		m.StartedObjectives[i] = &lstartedObjectives

	}

	return nil
}

type JobCrafterDirectoryListEntry struct {
	PlayerInfo *JobCrafterDirectoryEntryPlayerInfo

	JobInfo *JobCrafterDirectoryEntryJobInfo
}

func (m *JobCrafterDirectoryListEntry) ID() uint16 {
	return 196
}

func (m *JobCrafterDirectoryListEntry) Serialize(w Writer) error {

	if err := m.PlayerInfo.Serialize(w); err != nil {
		return err
	}

	if err := m.JobInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *JobCrafterDirectoryListEntry) Deserialize(r Reader) error {

	var lplayerInfo JobCrafterDirectoryEntryPlayerInfo

	lplayerInfo.Deserialize(r)

	m.PlayerInfo = &lplayerInfo

	var ljobInfo JobCrafterDirectoryEntryJobInfo

	ljobInfo.Deserialize(r)

	m.JobInfo = &ljobInfo

	return nil
}

type HavenBagFurnitureInformation struct {
	CellId uint16

	FunitureId int32

	Orientation uint8
}

func (m *HavenBagFurnitureInformation) ID() uint16 {
	return 498
}

func (m *HavenBagFurnitureInformation) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	if err := w.WriteInt32(m.FunitureId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Orientation); err != nil {
		return err
	}

	return nil
}

func (m *HavenBagFurnitureInformation) Deserialize(r Reader) error {

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	lfunitureId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FunitureId = lfunitureId

	lorientation, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Orientation = lorientation

	return nil
}

type KrosmasterFigure struct {
	Uid string

	Figure uint16

	Pedestal uint16

	Bound bool
}

func (m *KrosmasterFigure) ID() uint16 {
	return 397
}

func (m *KrosmasterFigure) Serialize(w Writer) error {

	if err := w.WriteString(m.Uid); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Figure); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Pedestal); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Bound); err != nil {
		return err
	}

	return nil
}

func (m *KrosmasterFigure) Deserialize(r Reader) error {

	luid, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Uid = luid

	lfigure, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Figure = lfigure

	lpedestal, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Pedestal = lpedestal

	lbound, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Bound = lbound

	return nil
}

type EntityMovementInformations struct {
	Id int32

	Steps []int8
}

func (m *EntityMovementInformations) ID() uint16 {
	return 63
}

func (m *EntityMovementInformations) Serialize(w Writer) error {

	if err := w.WriteInt32(m.Id); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Steps))); err != nil {
		return err
	}

	for i := range m.Steps {

		if err := w.WriteInt8(m.Steps[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *EntityMovementInformations) Deserialize(r Reader) error {

	lid, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Id = lid

	lstepsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Steps = make([]int8, lstepsLen)

	for i := range m.Steps {

		lsteps, err := r.ReadInt8()
		if err != nil {
			return err
		}

		m.Steps[i] = lsteps

	}

	return nil
}

type ObjectItemNotInContainer struct {
	Item

	ObjectGID uint16

	Effects []*ObjectEffect

	ObjectUID uint32

	Quantity uint32
}

func (m *ObjectItemNotInContainer) ID() uint16 {
	return 134
}

func (m *ObjectItemNotInContainer) Serialize(w Writer) error {

	if err := m.Item.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGID); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Effects))); err != nil {
		return err
	}

	for i := range m.Effects {

		if err := w.WriteUInt16(m.Effects[i].ID()); err != nil {
			return err
		}

		if err := m.Effects[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ObjectItemNotInContainer) Deserialize(r Reader) error {

	if err := m.Item.Deserialize(r); err != nil {
		return err
	}

	lobjectGID, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGID = lobjectGID

	leffectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Effects = make([]*ObjectEffect, leffectsLen)

	for i := range m.Effects {

		typeeffectsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		leffects, err := GetType(typeeffectsID)
		if err != nil {
			return err
		}

		leffects.Deserialize(r)

		m.Effects[i] = leffects.(*ObjectEffect)

	}

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type GoldItem struct {
	Item

	Sum int64
}

func (m *GoldItem) ID() uint16 {
	return 123
}

func (m *GoldItem) Serialize(w Writer) error {

	if err := m.Item.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Sum); err != nil {
		return err
	}

	return nil
}

func (m *GoldItem) Deserialize(r Reader) error {

	if err := m.Item.Deserialize(r); err != nil {
		return err
	}

	lsum, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Sum = lsum

	return nil
}

type AbstractCharacterToRefurbishInformation struct {
	AbstractCharacterInformation

	Colors []int32

	CosmeticId uint32
}

func (m *AbstractCharacterToRefurbishInformation) ID() uint16 {
	return 475
}

func (m *AbstractCharacterToRefurbishInformation) Serialize(w Writer) error {

	if err := m.AbstractCharacterInformation.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Colors))); err != nil {
		return err
	}

	for i := range m.Colors {

		if err := w.WriteInt32(m.Colors[i]); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt32(m.CosmeticId); err != nil {
		return err
	}

	return nil
}

func (m *AbstractCharacterToRefurbishInformation) Deserialize(r Reader) error {

	if err := m.AbstractCharacterInformation.Deserialize(r); err != nil {
		return err
	}

	lcolorsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Colors = make([]int32, lcolorsLen)

	for i := range m.Colors {

		lcolors, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.Colors[i] = lcolors

	}

	lcosmeticId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.CosmeticId = lcosmeticId

	return nil
}

type CharacterToRecolorInformation struct {
	AbstractCharacterToRefurbishInformation
}

func (m *CharacterToRecolorInformation) ID() uint16 {
	return 212
}

func (m *CharacterToRecolorInformation) Serialize(w Writer) error {

	if err := m.AbstractCharacterToRefurbishInformation.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *CharacterToRecolorInformation) Deserialize(r Reader) error {

	if err := m.AbstractCharacterToRefurbishInformation.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type CharacterToRelookInformation struct {
	AbstractCharacterToRefurbishInformation
}

func (m *CharacterToRelookInformation) ID() uint16 {
	return 399
}

func (m *CharacterToRelookInformation) Serialize(w Writer) error {

	if err := m.AbstractCharacterToRefurbishInformation.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *CharacterToRelookInformation) Deserialize(r Reader) error {

	if err := m.AbstractCharacterToRefurbishInformation.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type JobCrafterDirectoryEntryJobInfo struct {
	JobId uint8

	JobLevel uint8

	Free bool

	MinLevel uint8
}

func (m *JobCrafterDirectoryEntryJobInfo) ID() uint16 {
	return 195
}

func (m *JobCrafterDirectoryEntryJobInfo) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.JobId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.JobLevel); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Free); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.MinLevel); err != nil {
		return err
	}

	return nil
}

func (m *JobCrafterDirectoryEntryJobInfo) Deserialize(r Reader) error {

	ljobId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.JobId = ljobId

	ljobLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.JobLevel = ljobLevel

	lfree, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Free = lfree

	lminLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MinLevel = lminLevel

	return nil
}

type JobCrafterDirectoryEntryPlayerInfo struct {
	PlayerId int64

	PlayerName string

	AlignmentSide int8

	Breed int8

	Sex bool

	IsInWorkshop bool

	WorldX int16

	WorldY int16

	MapId int32

	SubAreaId uint16

	Status *PlayerStatus
}

func (m *JobCrafterDirectoryEntryPlayerInfo) ID() uint16 {
	return 194
}

func (m *JobCrafterDirectoryEntryPlayerInfo) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	if err := w.WriteString(m.PlayerName); err != nil {
		return err
	}

	if err := w.WriteInt8(m.AlignmentSide); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Sex); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.IsInWorkshop); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Status.ID()); err != nil {
		return err
	}

	if err := m.Status.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *JobCrafterDirectoryEntryPlayerInfo) Deserialize(r Reader) error {

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	lplayerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PlayerName = lplayerName

	lalignmentSide, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.AlignmentSide = lalignmentSide

	lbreed, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Breed = lbreed

	lsex, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Sex = lsex

	lisInWorkshop, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsInWorkshop = lisInWorkshop

	lworldX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldX = lworldX

	lworldY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WorldY = lworldY

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	typestatusID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lstatus, err := GetType(typestatusID)
	if err != nil {
		return err
	}

	lstatus.Deserialize(r)

	m.Status = lstatus.(*PlayerStatus)

	return nil
}

type AtlasPointsInformations struct {
	Type uint8

	Coords []*MapCoordinatesExtended
}

func (m *AtlasPointsInformations) ID() uint16 {
	return 175
}

func (m *AtlasPointsInformations) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Coords))); err != nil {
		return err
	}

	for i := range m.Coords {

		if err := m.Coords[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *AtlasPointsInformations) Deserialize(r Reader) error {

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	lcoordsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Coords = make([]*MapCoordinatesExtended, lcoordsLen)

	for i := range m.Coords {

		var lcoords MapCoordinatesExtended

		lcoords.Deserialize(r)

		m.Coords[i] = &lcoords

	}

	return nil
}

type FightTeamLightInformations struct {
	AbstractFightTeamInformations

	TeamMembersCount uint8

	MeanLevel uint32

	HasFriend bool

	HasGuildMember bool

	HasAllianceMember bool

	HasGroupMember bool

	HasMyTaxCollector bool
}

func (m *FightTeamLightInformations) ID() uint16 {
	return 115
}

func (m *FightTeamLightInformations) Serialize(w Writer) error {

	if err := m.AbstractFightTeamInformations.Serialize(w); err != nil {
		return err
	}

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.HasFriend)

	setWrappedFlag(bbw0, 1, m.HasGuildMember)

	setWrappedFlag(bbw0, 2, m.HasAllianceMember)

	setWrappedFlag(bbw0, 3, m.HasGroupMember)

	setWrappedFlag(bbw0, 4, m.HasMyTaxCollector)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.TeamMembersCount); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MeanLevel); err != nil {
		return err
	}

	return nil
}

func (m *FightTeamLightInformations) Deserialize(r Reader) error {

	if err := m.AbstractFightTeamInformations.Deserialize(r); err != nil {
		return err
	}

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.HasFriend = getWrappedFlag(bbw0, 0)

	m.HasGuildMember = getWrappedFlag(bbw0, 1)

	m.HasAllianceMember = getWrappedFlag(bbw0, 2)

	m.HasGroupMember = getWrappedFlag(bbw0, 3)

	m.HasMyTaxCollector = getWrappedFlag(bbw0, 4)

	lteamMembersCount, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TeamMembersCount = lteamMembersCount

	lmeanLevel, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MeanLevel = lmeanLevel

	return nil
}

type AchievementObjective struct {
	Id uint32

	MaxValue uint16
}

func (m *AchievementObjective) ID() uint16 {
	return 404
}

func (m *AchievementObjective) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.Id); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MaxValue); err != nil {
		return err
	}

	return nil
}

func (m *AchievementObjective) Deserialize(r Reader) error {

	lid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Id = lid

	lmaxValue, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MaxValue = lmaxValue

	return nil
}

type AchievementStartedObjective struct {
	AchievementObjective

	Value uint16
}

func (m *AchievementStartedObjective) ID() uint16 {
	return 402
}

func (m *AchievementStartedObjective) Serialize(w Writer) error {

	if err := m.AchievementObjective.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *AchievementStartedObjective) Deserialize(r Reader) error {

	if err := m.AchievementObjective.Deserialize(r); err != nil {
		return err
	}

	lvalue, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}
