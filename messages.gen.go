package d2protocol

import (
	"reflect"
)

var messages = map[uint16]reflect.Type{

	1: reflect.TypeOf((ProtocolRequired)(ProtocolRequired{})),

	2: reflect.TypeOf((NetworkDataContainerMessage)(NetworkDataContainerMessage{})),

	3: reflect.TypeOf((HelloConnectMessage)(HelloConnectMessage{})),

	4: reflect.TypeOf((IdentificationMessage)(IdentificationMessage{})),

	10: reflect.TypeOf((LoginQueueStatusMessage)(LoginQueueStatusMessage{})),

	20: reflect.TypeOf((IdentificationFailedMessage)(IdentificationFailedMessage{})),

	21: reflect.TypeOf((IdentificationFailedForBadVersionMessage)(IdentificationFailedForBadVersionMessage{})),

	22: reflect.TypeOf((IdentificationSuccessMessage)(IdentificationSuccessMessage{})),

	30: reflect.TypeOf((ServersListMessage)(ServersListMessage{})),

	40: reflect.TypeOf((ServerSelectionMessage)(ServerSelectionMessage{})),

	41: reflect.TypeOf((SelectedServerRefusedMessage)(SelectedServerRefusedMessage{})),

	42: reflect.TypeOf((SelectedServerDataMessage)(SelectedServerDataMessage{})),

	50: reflect.TypeOf((ServerStatusUpdateMessage)(ServerStatusUpdateMessage{})),

	75: reflect.TypeOf((ConsoleMessage)(ConsoleMessage{})),

	76: reflect.TypeOf((AdminCommandMessage)(AdminCommandMessage{})),

	101: reflect.TypeOf((HelloGameMessage)(HelloGameMessage{})),

	109: reflect.TypeOf((AlreadyConnectedMessage)(AlreadyConnectedMessage{})),

	110: reflect.TypeOf((AuthenticationTicketMessage)(AuthenticationTicketMessage{})),

	111: reflect.TypeOf((AuthenticationTicketAcceptedMessage)(AuthenticationTicketAcceptedMessage{})),

	112: reflect.TypeOf((AuthenticationTicketRefusedMessage)(AuthenticationTicketRefusedMessage{})),

	150: reflect.TypeOf((CharactersListRequestMessage)(CharactersListRequestMessage{})),

	151: reflect.TypeOf((CharactersListMessage)(CharactersListMessage{})),

	152: reflect.TypeOf((CharacterSelectionMessage)(CharacterSelectionMessage{})),

	153: reflect.TypeOf((CharacterSelectedSuccessMessage)(CharacterSelectedSuccessMessage{})),

	160: reflect.TypeOf((CharacterCreationRequestMessage)(CharacterCreationRequestMessage{})),

	161: reflect.TypeOf((CharacterCreationResultMessage)(CharacterCreationResultMessage{})),

	162: reflect.TypeOf((CharacterNameSuggestionRequestMessage)(CharacterNameSuggestionRequestMessage{})),

	164: reflect.TypeOf((CharacterNameSuggestionFailureMessage)(CharacterNameSuggestionFailureMessage{})),

	165: reflect.TypeOf((CharacterDeletionRequestMessage)(CharacterDeletionRequestMessage{})),

	166: reflect.TypeOf((CharacterDeletionErrorMessage)(CharacterDeletionErrorMessage{})),

	167: reflect.TypeOf((CharacterReplayRequestMessage)(CharacterReplayRequestMessage{})),

	170: reflect.TypeOf((SetCharacterRestrictionsMessage)(SetCharacterRestrictionsMessage{})),

	175: reflect.TypeOf((BasicTimeMessage)(BasicTimeMessage{})),

	176: reflect.TypeOf((BasicNoOperationMessage)(BasicNoOperationMessage{})),

	177: reflect.TypeOf((BasicDateMessage)(BasicDateMessage{})),

	179: reflect.TypeOf((BasicWhoIsNoMatchMessage)(BasicWhoIsNoMatchMessage{})),

	180: reflect.TypeOf((BasicWhoIsMessage)(BasicWhoIsMessage{})),

	181: reflect.TypeOf((BasicWhoIsRequestMessage)(BasicWhoIsRequestMessage{})),

	182: reflect.TypeOf((BasicPingMessage)(BasicPingMessage{})),

	183: reflect.TypeOf((BasicPongMessage)(BasicPongMessage{})),

	189: reflect.TypeOf((SystemMessageDisplayMessage)(SystemMessageDisplayMessage{})),

	200: reflect.TypeOf((GameContextCreateMessage)(GameContextCreateMessage{})),

	201: reflect.TypeOf((GameContextDestroyMessage)(GameContextDestroyMessage{})),

	210: reflect.TypeOf((MapFightCountMessage)(MapFightCountMessage{})),

	220: reflect.TypeOf((CurrentMapMessage)(CurrentMapMessage{})),

	221: reflect.TypeOf((ChangeMapMessage)(ChangeMapMessage{})),

	225: reflect.TypeOf((MapInformationsRequestMessage)(MapInformationsRequestMessage{})),

	226: reflect.TypeOf((MapComplementaryInformationsDataMessage)(MapComplementaryInformationsDataMessage{})),

	250: reflect.TypeOf((GameContextCreateRequestMessage)(GameContextCreateRequestMessage{})),

	251: reflect.TypeOf((GameContextRemoveElementMessage)(GameContextRemoveElementMessage{})),

	252: reflect.TypeOf((GameContextRemoveMultipleElementsMessage)(GameContextRemoveMultipleElementsMessage{})),

	253: reflect.TypeOf((GameContextMoveElementMessage)(GameContextMoveElementMessage{})),

	254: reflect.TypeOf((GameContextMoveMultipleElementsMessage)(GameContextMoveMultipleElementsMessage{})),

	255: reflect.TypeOf((GameContextQuitMessage)(GameContextQuitMessage{})),

	300: reflect.TypeOf((GameRolePlayRemoveChallengeMessage)(GameRolePlayRemoveChallengeMessage{})),

	301: reflect.TypeOf((GameRolePlayShowChallengeMessage)(GameRolePlayShowChallengeMessage{})),

	500: reflect.TypeOf((CharacterStatsListMessage)(CharacterStatsListMessage{})),

	700: reflect.TypeOf((GameFightStartingMessage)(GameFightStartingMessage{})),

	701: reflect.TypeOf((GameFightJoinRequestMessage)(GameFightJoinRequestMessage{})),

	702: reflect.TypeOf((GameFightJoinMessage)(GameFightJoinMessage{})),

	703: reflect.TypeOf((GameFightPlacementPossiblePositionsMessage)(GameFightPlacementPossiblePositionsMessage{})),

	704: reflect.TypeOf((GameFightPlacementPositionRequestMessage)(GameFightPlacementPositionRequestMessage{})),

	707: reflect.TypeOf((GameFightOptionToggleMessage)(GameFightOptionToggleMessage{})),

	708: reflect.TypeOf((GameFightReadyMessage)(GameFightReadyMessage{})),

	711: reflect.TypeOf((GameFightRemoveTeamMemberMessage)(GameFightRemoveTeamMemberMessage{})),

	712: reflect.TypeOf((GameFightStartMessage)(GameFightStartMessage{})),

	713: reflect.TypeOf((GameFightTurnListMessage)(GameFightTurnListMessage{})),

	714: reflect.TypeOf((GameFightTurnStartMessage)(GameFightTurnStartMessage{})),

	715: reflect.TypeOf((GameFightTurnReadyRequestMessage)(GameFightTurnReadyRequestMessage{})),

	716: reflect.TypeOf((GameFightTurnReadyMessage)(GameFightTurnReadyMessage{})),

	718: reflect.TypeOf((GameFightTurnFinishMessage)(GameFightTurnFinishMessage{})),

	719: reflect.TypeOf((GameFightTurnEndMessage)(GameFightTurnEndMessage{})),

	720: reflect.TypeOf((GameFightEndMessage)(GameFightEndMessage{})),

	721: reflect.TypeOf((GameFightLeaveMessage)(GameFightLeaveMessage{})),

	740: reflect.TypeOf((GameFightHumanReadyStateMessage)(GameFightHumanReadyStateMessage{})),

	745: reflect.TypeOf((GameRolePlayFreeSoulRequestMessage)(GameRolePlayFreeSoulRequestMessage{})),

	746: reflect.TypeOf((GameRolePlayGameOverMessage)(GameRolePlayGameOverMessage{})),

	780: reflect.TypeOf((TextInformationMessage)(TextInformationMessage{})),

	800: reflect.TypeOf((ChatSmileyRequestMessage)(ChatSmileyRequestMessage{})),

	801: reflect.TypeOf((ChatSmileyMessage)(ChatSmileyMessage{})),

	821: reflect.TypeOf((ChatMessageReportMessage)(ChatMessageReportMessage{})),

	850: reflect.TypeOf((ChatAbstractClientMessage)(ChatAbstractClientMessage{})),

	851: reflect.TypeOf((ChatClientPrivateMessage)(ChatClientPrivateMessage{})),

	852: reflect.TypeOf((ChatClientPrivateWithObjectMessage)(ChatClientPrivateWithObjectMessage{})),

	861: reflect.TypeOf((ChatClientMultiMessage)(ChatClientMultiMessage{})),

	862: reflect.TypeOf((ChatClientMultiWithObjectMessage)(ChatClientMultiWithObjectMessage{})),

	870: reflect.TypeOf((ChatErrorMessage)(ChatErrorMessage{})),

	880: reflect.TypeOf((ChatAbstractServerMessage)(ChatAbstractServerMessage{})),

	881: reflect.TypeOf((ChatServerMessage)(ChatServerMessage{})),

	882: reflect.TypeOf((ChatServerCopyMessage)(ChatServerCopyMessage{})),

	883: reflect.TypeOf((ChatServerWithObjectMessage)(ChatServerWithObjectMessage{})),

	884: reflect.TypeOf((ChatServerCopyWithObjectMessage)(ChatServerCopyWithObjectMessage{})),

	890: reflect.TypeOf((ChannelEnablingMessage)(ChannelEnablingMessage{})),

	891: reflect.TypeOf((ChannelEnablingChangeMessage)(ChannelEnablingChangeMessage{})),

	892: reflect.TypeOf((EnabledChannelsMessage)(EnabledChannelsMessage{})),

	945: reflect.TypeOf((GameMapChangeOrientationRequestMessage)(GameMapChangeOrientationRequestMessage{})),

	946: reflect.TypeOf((GameMapChangeOrientationMessage)(GameMapChangeOrientationMessage{})),

	950: reflect.TypeOf((GameMapMovementRequestMessage)(GameMapMovementRequestMessage{})),

	951: reflect.TypeOf((GameMapMovementMessage)(GameMapMovementMessage{})),

	952: reflect.TypeOf((GameMapMovementConfirmMessage)(GameMapMovementConfirmMessage{})),

	953: reflect.TypeOf((GameMapMovementCancelMessage)(GameMapMovementCancelMessage{})),

	954: reflect.TypeOf((GameMapNoMovementMessage)(GameMapNoMovementMessage{})),

	955: reflect.TypeOf((SequenceStartMessage)(SequenceStartMessage{})),

	956: reflect.TypeOf((SequenceEndMessage)(SequenceEndMessage{})),

	957: reflect.TypeOf((GameActionAcknowledgementMessage)(GameActionAcknowledgementMessage{})),

	1000: reflect.TypeOf((AbstractGameActionMessage)(AbstractGameActionMessage{})),

	1001: reflect.TypeOf((AbstractGameActionWithAckMessage)(AbstractGameActionWithAckMessage{})),

	1002: reflect.TypeOf((GameActionNoopMessage)(GameActionNoopMessage{})),

	1004: reflect.TypeOf((GameActionFightTackledMessage)(GameActionFightTackledMessage{})),

	1005: reflect.TypeOf((GameActionFightCastRequestMessage)(GameActionFightCastRequestMessage{})),

	1010: reflect.TypeOf((GameActionFightSpellCastMessage)(GameActionFightSpellCastMessage{})),

	1030: reflect.TypeOf((GameActionFightPointsVariationMessage)(GameActionFightPointsVariationMessage{})),

	1099: reflect.TypeOf((GameActionFightDeathMessage)(GameActionFightDeathMessage{})),

	1200: reflect.TypeOf((SpellListMessage)(SpellListMessage{})),

	1301: reflect.TypeOf((StartupActionsListMessage)(StartupActionsListMessage{})),

	1302: reflect.TypeOf((StartupActionsExecuteMessage)(StartupActionsExecuteMessage{})),

	1303: reflect.TypeOf((StartupActionsObjetAttributionMessage)(StartupActionsObjetAttributionMessage{})),

	1304: reflect.TypeOf((StartupActionFinishedMessage)(StartupActionFinishedMessage{})),

	1501: reflect.TypeOf((GetPartsListMessage)(GetPartsListMessage{})),

	1502: reflect.TypeOf((PartsListMessage)(PartsListMessage{})),

	1503: reflect.TypeOf((DownloadPartMessage)(DownloadPartMessage{})),

	1506: reflect.TypeOf((GetPartInfoMessage)(GetPartInfoMessage{})),

	1508: reflect.TypeOf((PartInfoMessage)(PartInfoMessage{})),

	1510: reflect.TypeOf((DownloadGetCurrentSpeedRequestMessage)(DownloadGetCurrentSpeedRequestMessage{})),

	1511: reflect.TypeOf((DownloadCurrentSpeedMessage)(DownloadCurrentSpeedMessage{})),

	1512: reflect.TypeOf((DownloadSetSpeedRequestMessage)(DownloadSetSpeedRequestMessage{})),

	1513: reflect.TypeOf((DownloadErrorMessage)(DownloadErrorMessage{})),

	1604: reflect.TypeOf((ZaapListMessage)(ZaapListMessage{})),

	1810: reflect.TypeOf((SetEnablePVPRequestMessage)(SetEnablePVPRequestMessage{})),

	2001: reflect.TypeOf((DebugHighlightCellsMessage)(DebugHighlightCellsMessage{})),

	2002: reflect.TypeOf((DebugClearHighlightCellsMessage)(DebugClearHighlightCellsMessage{})),

	3003: reflect.TypeOf((ObjectUseOnCharacterMessage)(ObjectUseOnCharacterMessage{})),

	3004: reflect.TypeOf((ObjectErrorMessage)(ObjectErrorMessage{})),

	3005: reflect.TypeOf((ObjectDropMessage)(ObjectDropMessage{})),

	3009: reflect.TypeOf((InventoryWeightMessage)(InventoryWeightMessage{})),

	3010: reflect.TypeOf((ObjectMovementMessage)(ObjectMovementMessage{})),

	3013: reflect.TypeOf((ObjectUseOnCellMessage)(ObjectUseOnCellMessage{})),

	3014: reflect.TypeOf((ObjectGroundRemovedMessage)(ObjectGroundRemovedMessage{})),

	3016: reflect.TypeOf((InventoryContentMessage)(InventoryContentMessage{})),

	3017: reflect.TypeOf((ObjectGroundAddedMessage)(ObjectGroundAddedMessage{})),

	3019: reflect.TypeOf((ObjectUseMessage)(ObjectUseMessage{})),

	3021: reflect.TypeOf((ObjectSetPositionMessage)(ObjectSetPositionMessage{})),

	3022: reflect.TypeOf((ObjectDeleteMessage)(ObjectDeleteMessage{})),

	3023: reflect.TypeOf((ObjectQuantityMessage)(ObjectQuantityMessage{})),

	3024: reflect.TypeOf((ObjectDeletedMessage)(ObjectDeletedMessage{})),

	3025: reflect.TypeOf((ObjectAddedMessage)(ObjectAddedMessage{})),

	3029: reflect.TypeOf((ObjectModifiedMessage)(ObjectModifiedMessage{})),

	4001: reflect.TypeOf((FriendsGetListMessage)(FriendsGetListMessage{})),

	4002: reflect.TypeOf((FriendsListMessage)(FriendsListMessage{})),

	4004: reflect.TypeOf((FriendAddRequestMessage)(FriendAddRequestMessage{})),

	5001: reflect.TypeOf((InteractiveUseRequestMessage)(InteractiveUseRequestMessage{})),

	5002: reflect.TypeOf((InteractiveMapUpdateMessage)(InteractiveMapUpdateMessage{})),

	5501: reflect.TypeOf((LeaveDialogRequestMessage)(LeaveDialogRequestMessage{})),

	5502: reflect.TypeOf((LeaveDialogMessage)(LeaveDialogMessage{})),

	5503: reflect.TypeOf((SetUpdateMessage)(SetUpdateMessage{})),

	5505: reflect.TypeOf((ExchangeRequestMessage)(ExchangeRequestMessage{})),

	5508: reflect.TypeOf((ExchangeAcceptMessage)(ExchangeAcceptMessage{})),

	5509: reflect.TypeOf((ExchangeIsReadyMessage)(ExchangeIsReadyMessage{})),

	5511: reflect.TypeOf((ExchangeReadyMessage)(ExchangeReadyMessage{})),

	5512: reflect.TypeOf((ExchangeStartedMessage)(ExchangeStartedMessage{})),

	5513: reflect.TypeOf((ExchangeErrorMessage)(ExchangeErrorMessage{})),

	5514: reflect.TypeOf((ExchangeObjectMovePricedMessage)(ExchangeObjectMovePricedMessage{})),

	5515: reflect.TypeOf((ExchangeObjectMessage)(ExchangeObjectMessage{})),

	5516: reflect.TypeOf((ExchangeObjectAddedMessage)(ExchangeObjectAddedMessage{})),

	5517: reflect.TypeOf((ExchangeObjectRemovedMessage)(ExchangeObjectRemovedMessage{})),

	5518: reflect.TypeOf((ExchangeObjectMoveMessage)(ExchangeObjectMoveMessage{})),

	5519: reflect.TypeOf((ExchangeObjectModifiedMessage)(ExchangeObjectModifiedMessage{})),

	5520: reflect.TypeOf((ExchangeObjectMoveKamaMessage)(ExchangeObjectMoveKamaMessage{})),

	5521: reflect.TypeOf((ExchangeKamaModifiedMessage)(ExchangeKamaModifiedMessage{})),

	5522: reflect.TypeOf((ExchangeRequestedMessage)(ExchangeRequestedMessage{})),

	5523: reflect.TypeOf((ExchangeRequestedTradeMessage)(ExchangeRequestedTradeMessage{})),

	5525: reflect.TypeOf((GameActionFightSlideMessage)(GameActionFightSlideMessage{})),

	5526: reflect.TypeOf((GameActionFightReduceDamagesMessage)(GameActionFightReduceDamagesMessage{})),

	5527: reflect.TypeOf((GameActionFightExchangePositionsMessage)(GameActionFightExchangePositionsMessage{})),

	5528: reflect.TypeOf((GameActionFightTeleportOnSameMapMessage)(GameActionFightTeleportOnSameMapMessage{})),

	5530: reflect.TypeOf((GameActionFightReflectDamagesMessage)(GameActionFightReflectDamagesMessage{})),

	5531: reflect.TypeOf((GameActionFightReflectSpellMessage)(GameActionFightReflectSpellMessage{})),

	5532: reflect.TypeOf((GameActionFightChangeLookMessage)(GameActionFightChangeLookMessage{})),

	5533: reflect.TypeOf((GameActionFightDispellMessage)(GameActionFightDispellMessage{})),

	5535: reflect.TypeOf((GameActionFightStealKamaMessage)(GameActionFightStealKamaMessage{})),

	5537: reflect.TypeOf((KamasUpdateMessage)(KamasUpdateMessage{})),

	5540: reflect.TypeOf((GameActionFightMarkCellsMessage)(GameActionFightMarkCellsMessage{})),

	5542: reflect.TypeOf((SubscriptionLimitationMessage)(SubscriptionLimitationMessage{})),

	5544: reflect.TypeOf((CharacterNameSuggestionSuccessMessage)(CharacterNameSuggestionSuccessMessage{})),

	5545: reflect.TypeOf((CharactersListErrorMessage)(CharactersListErrorMessage{})),

	5546: reflect.TypeOf((GuildCreationValidMessage)(GuildCreationValidMessage{})),

	5548: reflect.TypeOf((GuildInvitationStateRecrutedMessage)(GuildInvitationStateRecrutedMessage{})),

	5549: reflect.TypeOf((GuildChangeMemberParametersMessage)(GuildChangeMemberParametersMessage{})),

	5550: reflect.TypeOf((GuildGetInformationsMessage)(GuildGetInformationsMessage{})),

	5551: reflect.TypeOf((GuildInvitationMessage)(GuildInvitationMessage{})),

	5552: reflect.TypeOf((GuildInvitedMessage)(GuildInvitedMessage{})),

	5554: reflect.TypeOf((GuildCreationResultMessage)(GuildCreationResultMessage{})),

	5556: reflect.TypeOf((GuildInvitationAnswerMessage)(GuildInvitationAnswerMessage{})),

	5557: reflect.TypeOf((GuildInformationsGeneralMessage)(GuildInformationsGeneralMessage{})),

	5558: reflect.TypeOf((GuildInformationsMembersMessage)(GuildInformationsMembersMessage{})),

	5562: reflect.TypeOf((GuildLeftMessage)(GuildLeftMessage{})),

	5563: reflect.TypeOf((GuildInvitationStateRecruterMessage)(GuildInvitationStateRecruterMessage{})),

	5564: reflect.TypeOf((GuildJoinedMessage)(GuildJoinedMessage{})),

	5570: reflect.TypeOf((GameActionFightUnmarkCellsMessage)(GameActionFightUnmarkCellsMessage{})),

	5571: reflect.TypeOf((GameActionFightKillMessage)(GameActionFightKillMessage{})),

	5572: reflect.TypeOf((GameFightUpdateTeamMessage)(GameFightUpdateTeamMessage{})),

	5573: reflect.TypeOf((SubscriptionZoneMessage)(SubscriptionZoneMessage{})),

	5574: reflect.TypeOf((PartyStopFollowRequestMessage)(PartyStopFollowRequestMessage{})),

	5575: reflect.TypeOf((PartyUpdateMessage)(PartyUpdateMessage{})),

	5576: reflect.TypeOf((PartyJoinMessage)(PartyJoinMessage{})),

	5577: reflect.TypeOf((PartyFollowMemberRequestMessage)(PartyFollowMemberRequestMessage{})),

	5578: reflect.TypeOf((PartyLeaderUpdateMessage)(PartyLeaderUpdateMessage{})),

	5579: reflect.TypeOf((PartyMemberRemoveMessage)(PartyMemberRemoveMessage{})),

	5580: reflect.TypeOf((PartyAcceptInvitationMessage)(PartyAcceptInvitationMessage{})),

	5581: reflect.TypeOf((PartyFollowStatusUpdateMessage)(PartyFollowStatusUpdateMessage{})),

	5582: reflect.TypeOf((PartyRefuseInvitationMessage)(PartyRefuseInvitationMessage{})),

	5583: reflect.TypeOf((PartyCannotJoinErrorMessage)(PartyCannotJoinErrorMessage{})),

	5584: reflect.TypeOf((CompassResetMessage)(CompassResetMessage{})),

	5585: reflect.TypeOf((PartyInvitationRequestMessage)(PartyInvitationRequestMessage{})),

	5586: reflect.TypeOf((PartyInvitationMessage)(PartyInvitationMessage{})),

	5587: reflect.TypeOf((PartyLocateMembersRequestMessage)(PartyLocateMembersRequestMessage{})),

	5588: reflect.TypeOf((PartyFollowThisMemberRequestMessage)(PartyFollowThisMemberRequestMessage{})),

	5589: reflect.TypeOf((CompassUpdatePartyMemberMessage)(CompassUpdatePartyMemberMessage{})),

	5590: reflect.TypeOf((PartyKickedByMessage)(PartyKickedByMessage{})),

	5591: reflect.TypeOf((CompassUpdateMessage)(CompassUpdateMessage{})),

	5592: reflect.TypeOf((PartyKickRequestMessage)(PartyKickRequestMessage{})),

	5593: reflect.TypeOf((PartyLeaveRequestMessage)(PartyLeaveRequestMessage{})),

	5594: reflect.TypeOf((PartyLeaveMessage)(PartyLeaveMessage{})),

	5595: reflect.TypeOf((PartyLocateMembersMessage)(PartyLocateMembersMessage{})),

	5596: reflect.TypeOf((PartyRefuseInvitationNotificationMessage)(PartyRefuseInvitationNotificationMessage{})),

	5597: reflect.TypeOf((GuildInformationsMemberUpdateMessage)(GuildInformationsMemberUpdateMessage{})),

	5599: reflect.TypeOf((FriendAddedMessage)(FriendAddedMessage{})),

	5600: reflect.TypeOf((FriendAddFailureMessage)(FriendAddFailureMessage{})),

	5601: reflect.TypeOf((FriendDeleteResultMessage)(FriendDeleteResultMessage{})),

	5602: reflect.TypeOf((FriendSetWarnOnConnectionMessage)(FriendSetWarnOnConnectionMessage{})),

	5603: reflect.TypeOf((FriendDeleteRequestMessage)(FriendDeleteRequestMessage{})),

	5604: reflect.TypeOf((FriendSpouseJoinRequestMessage)(FriendSpouseJoinRequestMessage{})),

	5605: reflect.TypeOf((FriendJoinRequestMessage)(FriendJoinRequestMessage{})),

	5606: reflect.TypeOf((FriendSpouseFollowWithCompassRequestMessage)(FriendSpouseFollowWithCompassRequestMessage{})),

	5607: reflect.TypeOf((ClientKeyMessage)(ClientKeyMessage{})),

	5609: reflect.TypeOf((StatsUpgradeResultMessage)(StatsUpgradeResultMessage{})),

	5610: reflect.TypeOf((StatsUpgradeRequestMessage)(StatsUpgradeRequestMessage{})),

	5611: reflect.TypeOf((ShowCellRequestMessage)(ShowCellRequestMessage{})),

	5612: reflect.TypeOf((ShowCellMessage)(ShowCellMessage{})),

	5613: reflect.TypeOf((ChallengeTargetsListMessage)(ChallengeTargetsListMessage{})),

	5614: reflect.TypeOf((ChallengeTargetsListRequestMessage)(ChallengeTargetsListRequestMessage{})),

	5615: reflect.TypeOf((TaxCollectorDialogQuestionExtendedMessage)(TaxCollectorDialogQuestionExtendedMessage{})),

	5616: reflect.TypeOf((NpcDialogReplyMessage)(NpcDialogReplyMessage{})),

	5617: reflect.TypeOf((NpcDialogQuestionMessage)(NpcDialogQuestionMessage{})),

	5618: reflect.TypeOf((NpcDialogCreationMessage)(NpcDialogCreationMessage{})),

	5619: reflect.TypeOf((TaxCollectorDialogQuestionBasicMessage)(TaxCollectorDialogQuestionBasicMessage{})),

	5622: reflect.TypeOf((QuestStepInfoRequestMessage)(QuestStepInfoRequestMessage{})),

	5623: reflect.TypeOf((QuestListRequestMessage)(QuestListRequestMessage{})),

	5625: reflect.TypeOf((QuestStepInfoMessage)(QuestStepInfoMessage{})),

	5626: reflect.TypeOf((QuestListMessage)(QuestListMessage{})),

	5628: reflect.TypeOf((ExchangeLeaveMessage)(ExchangeLeaveMessage{})),

	5630: reflect.TypeOf((FriendWarnOnConnectionStateMessage)(FriendWarnOnConnectionStateMessage{})),

	5632: reflect.TypeOf((GameRolePlayShowActorMessage)(GameRolePlayShowActorMessage{})),

	5633: reflect.TypeOf((TaxCollectorMovementMessage)(TaxCollectorMovementMessage{})),

	5634: reflect.TypeOf((TaxCollectorErrorMessage)(TaxCollectorErrorMessage{})),

	5635: reflect.TypeOf((TaxCollectorAttackedResultMessage)(TaxCollectorAttackedResultMessage{})),

	5636: reflect.TypeOf((GuildInfosUpgradeMessage)(GuildInfosUpgradeMessage{})),

	5637: reflect.TypeOf((GameContextRefreshEntityLookMessage)(GameContextRefreshEntityLookMessage{})),

	5638: reflect.TypeOf((NicknameRefusedMessage)(NicknameRefusedMessage{})),

	5639: reflect.TypeOf((NicknameChoiceRequestMessage)(NicknameChoiceRequestMessage{})),

	5640: reflect.TypeOf((NicknameRegistrationMessage)(NicknameRegistrationMessage{})),

	5641: reflect.TypeOf((NicknameAcceptedMessage)(NicknameAcceptedMessage{})),

	5642: reflect.TypeOf((MapNpcsQuestStatusUpdateMessage)(MapNpcsQuestStatusUpdateMessage{})),

	5643: reflect.TypeOf((QuestStartRequestMessage)(QuestStartRequestMessage{})),

	5644: reflect.TypeOf((EmoteAddMessage)(EmoteAddMessage{})),

	5645: reflect.TypeOf((StorageKamasUpdateMessage)(StorageKamasUpdateMessage{})),

	5646: reflect.TypeOf((StorageInventoryContentMessage)(StorageInventoryContentMessage{})),

	5647: reflect.TypeOf((StorageObjectUpdateMessage)(StorageObjectUpdateMessage{})),

	5648: reflect.TypeOf((StorageObjectRemoveMessage)(StorageObjectRemoveMessage{})),

	5649: reflect.TypeOf((JobCrafterDirectoryDefineSettingsMessage)(JobCrafterDirectoryDefineSettingsMessage{})),

	5651: reflect.TypeOf((JobCrafterDirectoryAddMessage)(JobCrafterDirectoryAddMessage{})),

	5652: reflect.TypeOf((JobCrafterDirectorySettingsMessage)(JobCrafterDirectorySettingsMessage{})),

	5653: reflect.TypeOf((JobCrafterDirectoryRemoveMessage)(JobCrafterDirectoryRemoveMessage{})),

	5654: reflect.TypeOf((JobExperienceUpdateMessage)(JobExperienceUpdateMessage{})),

	5655: reflect.TypeOf((JobDescriptionMessage)(JobDescriptionMessage{})),

	5656: reflect.TypeOf((JobLevelUpMessage)(JobLevelUpMessage{})),

	5658: reflect.TypeOf((UpdateLifePointsMessage)(UpdateLifePointsMessage{})),

	5661: reflect.TypeOf((HouseKickIndoorMerchantRequestMessage)(HouseKickIndoorMerchantRequestMessage{})),

	5662: reflect.TypeOf((AdminQuietCommandMessage)(AdminQuietCommandMessage{})),

	5663: reflect.TypeOf((BasicLatencyStatsMessage)(BasicLatencyStatsMessage{})),

	5664: reflect.TypeOf((BasicWhoAmIRequestMessage)(BasicWhoAmIRequestMessage{})),

	5666: reflect.TypeOf((LockableChangeCodeMessage)(LockableChangeCodeMessage{})),

	5667: reflect.TypeOf((LockableUseCodeMessage)(LockableUseCodeMessage{})),

	5668: reflect.TypeOf((LockableStateUpdateHouseDoorMessage)(LockableStateUpdateHouseDoorMessage{})),

	5669: reflect.TypeOf((LockableStateUpdateStorageMessage)(LockableStateUpdateStorageMessage{})),

	5670: reflect.TypeOf((CharacterLevelUpMessage)(CharacterLevelUpMessage{})),

	5671: reflect.TypeOf((LockableStateUpdateAbstractMessage)(LockableStateUpdateAbstractMessage{})),

	5672: reflect.TypeOf((LockableCodeResultMessage)(LockableCodeResultMessage{})),

	5673: reflect.TypeOf((IgnoredAddRequestMessage)(IgnoredAddRequestMessage{})),

	5674: reflect.TypeOf((IgnoredListMessage)(IgnoredListMessage{})),

	5675: reflect.TypeOf((DocumentReadingBeginMessage)(DocumentReadingBeginMessage{})),

	5676: reflect.TypeOf((IgnoredGetListMessage)(IgnoredGetListMessage{})),

	5677: reflect.TypeOf((IgnoredDeleteResultMessage)(IgnoredDeleteResultMessage{})),

	5678: reflect.TypeOf((IgnoredAddedMessage)(IgnoredAddedMessage{})),

	5679: reflect.TypeOf((IgnoredAddFailureMessage)(IgnoredAddFailureMessage{})),

	5680: reflect.TypeOf((IgnoredDeleteRequestMessage)(IgnoredDeleteRequestMessage{})),

	5683: reflect.TypeOf((EmotePlayMessage)(EmotePlayMessage{})),

	5684: reflect.TypeOf((LifePointsRegenBeginMessage)(LifePointsRegenBeginMessage{})),

	5685: reflect.TypeOf((EmotePlayRequestMessage)(EmotePlayRequestMessage{})),

	5686: reflect.TypeOf((LifePointsRegenEndMessage)(LifePointsRegenEndMessage{})),

	5687: reflect.TypeOf((EmoteRemoveMessage)(EmoteRemoveMessage{})),

	5688: reflect.TypeOf((EmotePlayErrorMessage)(EmotePlayErrorMessage{})),

	5689: reflect.TypeOf((EmoteListMessage)(EmoteListMessage{})),

	5690: reflect.TypeOf((EmotePlayAbstractMessage)(EmotePlayAbstractMessage{})),

	5691: reflect.TypeOf((EmotePlayMassiveMessage)(EmotePlayMassiveMessage{})),

	5693: reflect.TypeOf((GameEntityDispositionMessage)(GameEntityDispositionMessage{})),

	5695: reflect.TypeOf((GameEntityDispositionErrorMessage)(GameEntityDispositionErrorMessage{})),

	5696: reflect.TypeOf((GameEntitiesDispositionMessage)(GameEntitiesDispositionMessage{})),

	5697: reflect.TypeOf((HouseSellRequestMessage)(HouseSellRequestMessage{})),

	5698: reflect.TypeOf((HouseKickRequestMessage)(HouseKickRequestMessage{})),

	5699: reflect.TypeOf((GuildSpellUpgradeRequestMessage)(GuildSpellUpgradeRequestMessage{})),

	5700: reflect.TypeOf((HouseGuildRightsViewMessage)(HouseGuildRightsViewMessage{})),

	5701: reflect.TypeOf((HouseGuildNoneMessage)(HouseGuildNoneMessage{})),

	5703: reflect.TypeOf((HouseGuildRightsMessage)(HouseGuildRightsMessage{})),

	5704: reflect.TypeOf((HouseGuildShareRequestMessage)(HouseGuildShareRequestMessage{})),

	5706: reflect.TypeOf((GuildCharacsUpgradeRequestMessage)(GuildCharacsUpgradeRequestMessage{})),

	5708: reflect.TypeOf((InteractiveElementUpdatedMessage)(InteractiveElementUpdatedMessage{})),

	5709: reflect.TypeOf((StatedElementUpdatedMessage)(StatedElementUpdatedMessage{})),

	5715: reflect.TypeOf((GuildFightLeaveRequestMessage)(GuildFightLeaveRequestMessage{})),

	5716: reflect.TypeOf((StatedMapUpdateMessage)(StatedMapUpdateMessage{})),

	5717: reflect.TypeOf((GuildFightJoinRequestMessage)(GuildFightJoinRequestMessage{})),

	5719: reflect.TypeOf((GuildFightPlayersHelpersLeaveMessage)(GuildFightPlayersHelpersLeaveMessage{})),

	5720: reflect.TypeOf((GuildFightPlayersHelpersJoinMessage)(GuildFightPlayersHelpersJoinMessage{})),

	5723: reflect.TypeOf((LivingObjectDissociateMessage)(LivingObjectDissociateMessage{})),

	5725: reflect.TypeOf((LivingObjectChangeSkinRequestMessage)(LivingObjectChangeSkinRequestMessage{})),

	5726: reflect.TypeOf((OnConnectionEventMessage)(OnConnectionEventMessage{})),

	5731: reflect.TypeOf((GameRolePlayPlayerFightRequestMessage)(GameRolePlayPlayerFightRequestMessage{})),

	5732: reflect.TypeOf((GameRolePlayPlayerFightFriendlyAnswerMessage)(GameRolePlayPlayerFightFriendlyAnswerMessage{})),

	5733: reflect.TypeOf((GameRolePlayPlayerFightFriendlyAnsweredMessage)(GameRolePlayPlayerFightFriendlyAnsweredMessage{})),

	5734: reflect.TypeOf((HousePropertiesMessage)(HousePropertiesMessage{})),

	5735: reflect.TypeOf((HouseBuyResultMessage)(HouseBuyResultMessage{})),

	5738: reflect.TypeOf((HouseBuyRequestMessage)(HouseBuyRequestMessage{})),

	5739: reflect.TypeOf((PurchasableDialogMessage)(PurchasableDialogMessage{})),

	5740: reflect.TypeOf((LockableShowCodeDialogMessage)(LockableShowCodeDialogMessage{})),

	5741: reflect.TypeOf((GameActionFightTriggerGlyphTrapMessage)(GameActionFightTriggerGlyphTrapMessage{})),

	5742: reflect.TypeOf((MapRunningFightListRequestMessage)(MapRunningFightListRequestMessage{})),

	5743: reflect.TypeOf((MapRunningFightListMessage)(MapRunningFightListMessage{})),

	5745: reflect.TypeOf((InteractiveUsedMessage)(InteractiveUsedMessage{})),

	5747: reflect.TypeOf((JobMultiCraftAvailableSkillsMessage)(JobMultiCraftAvailableSkillsMessage{})),

	5748: reflect.TypeOf((JobAllowMultiCraftRequestMessage)(JobAllowMultiCraftRequestMessage{})),

	5750: reflect.TypeOf((MapRunningFightDetailsRequestMessage)(MapRunningFightDetailsRequestMessage{})),

	5751: reflect.TypeOf((MapRunningFightDetailsMessage)(MapRunningFightDetailsMessage{})),

	5752: reflect.TypeOf((ExchangeTypesItemsExchangerDescriptionForUserMessage)(ExchangeTypesItemsExchangerDescriptionForUserMessage{})),

	5753: reflect.TypeOf((ExchangeRequestOnShopStockMessage)(ExchangeRequestOnShopStockMessage{})),

	5755: reflect.TypeOf((ExchangeBidPriceMessage)(ExchangeBidPriceMessage{})),

	5759: reflect.TypeOf((ExchangeBuyOkMessage)(ExchangeBuyOkMessage{})),

	5761: reflect.TypeOf((ExchangeStartOkNpcShopMessage)(ExchangeStartOkNpcShopMessage{})),

	5762: reflect.TypeOf((ExchangeGuildTaxCollectorGetMessage)(ExchangeGuildTaxCollectorGetMessage{})),

	5765: reflect.TypeOf((ExchangeTypesExchangerDescriptionForUserMessage)(ExchangeTypesExchangerDescriptionForUserMessage{})),

	5767: reflect.TypeOf((ExchangeStartOkHumanVendorMessage)(ExchangeStartOkHumanVendorMessage{})),

	5768: reflect.TypeOf((ExchangeOkMultiCraftMessage)(ExchangeOkMultiCraftMessage{})),

	5769: reflect.TypeOf((ItemNoMoreAvailableMessage)(ItemNoMoreAvailableMessage{})),

	5772: reflect.TypeOf((ExchangeOnHumanVendorRequestMessage)(ExchangeOnHumanVendorRequestMessage{})),

	5773: reflect.TypeOf((ExchangePlayerRequestMessage)(ExchangePlayerRequestMessage{})),

	5774: reflect.TypeOf((ExchangeBuyMessage)(ExchangeBuyMessage{})),

	5775: reflect.TypeOf((ExchangeStartAsVendorMessage)(ExchangeStartAsVendorMessage{})),

	5778: reflect.TypeOf((ExchangeSellMessage)(ExchangeSellMessage{})),

	5779: reflect.TypeOf((ExchangeRequestOnTaxCollectorMessage)(ExchangeRequestOnTaxCollectorMessage{})),

	5783: reflect.TypeOf((ExchangeShowVendorTaxMessage)(ExchangeShowVendorTaxMessage{})),

	5784: reflect.TypeOf((ExchangePlayerMultiCraftRequestMessage)(ExchangePlayerMultiCraftRequestMessage{})),

	5785: reflect.TypeOf((ExchangeStartOkNpcTradeMessage)(ExchangeStartOkNpcTradeMessage{})),

	5786: reflect.TypeOf((ExchangeWaitingResultMessage)(ExchangeWaitingResultMessage{})),

	5787: reflect.TypeOf((ExchangeReplyTaxVendorMessage)(ExchangeReplyTaxVendorMessage{})),

	5790: reflect.TypeOf((ExchangeCraftResultMessage)(ExchangeCraftResultMessage{})),

	5792: reflect.TypeOf((ExchangeSellOkMessage)(ExchangeSellOkMessage{})),

	5793: reflect.TypeOf((ExchangeWeightMessage)(ExchangeWeightMessage{})),

	5794: reflect.TypeOf((ExchangeCraftInformationObjectMessage)(ExchangeCraftInformationObjectMessage{})),

	5802: reflect.TypeOf((ExchangeBidSearchOkMessage)(ExchangeBidSearchOkMessage{})),

	5803: reflect.TypeOf((ExchangeBidHouseTypeMessage)(ExchangeBidHouseTypeMessage{})),

	5804: reflect.TypeOf((ExchangeBidHouseBuyMessage)(ExchangeBidHouseBuyMessage{})),

	5805: reflect.TypeOf((ExchangeBidHousePriceMessage)(ExchangeBidHousePriceMessage{})),

	5806: reflect.TypeOf((ExchangeBidHouseSearchMessage)(ExchangeBidHouseSearchMessage{})),

	5807: reflect.TypeOf((ExchangeBidHouseListMessage)(ExchangeBidHouseListMessage{})),

	5809: reflect.TypeOf((JobExperienceMultiUpdateMessage)(JobExperienceMultiUpdateMessage{})),

	5810: reflect.TypeOf((ExchangeItemAutoCraftStopedMessage)(ExchangeItemAutoCraftStopedMessage{})),

	5813: reflect.TypeOf((ExchangeStartOkCraftMessage)(ExchangeStartOkCraftMessage{})),

	5816: reflect.TypeOf((BasicLatencyStatsRequestMessage)(BasicLatencyStatsRequestMessage{})),

	5817: reflect.TypeOf((ExchangeStartOkMulticraftCustomerMessage)(ExchangeStartOkMulticraftCustomerMessage{})),

	5818: reflect.TypeOf((ExchangeStartOkMulticraftCrafterMessage)(ExchangeStartOkMulticraftCrafterMessage{})),

	5819: reflect.TypeOf((ExchangeStartOkJobIndexMessage)(ExchangeStartOkJobIndexMessage{})),

	5821: reflect.TypeOf((GameActionFightInvisibilityMessage)(GameActionFightInvisibilityMessage{})),

	5822: reflect.TypeOf((GameRolePlayFightRequestCanceledMessage)(GameRolePlayFightRequestCanceledMessage{})),

	5824: reflect.TypeOf((PaddockPropertiesMessage)(PaddockPropertiesMessage{})),

	5825: reflect.TypeOf((GameActionFightSummonMessage)(GameActionFightSummonMessage{})),

	5826: reflect.TypeOf((GameActionFightDropCharacterMessage)(GameActionFightDropCharacterMessage{})),

	5828: reflect.TypeOf((GameActionFightDodgePointLossMessage)(GameActionFightDodgePointLossMessage{})),

	5829: reflect.TypeOf((GameActionFightThrowCharacterMessage)(GameActionFightThrowCharacterMessage{})),

	5830: reflect.TypeOf((GameActionFightCarryCharacterMessage)(GameActionFightCarryCharacterMessage{})),

	5835: reflect.TypeOf((GuildMembershipMessage)(GuildMembershipMessage{})),

	5836: reflect.TypeOf((CharacterSelectedErrorMessage)(CharacterSelectedErrorMessage{})),

	5843: reflect.TypeOf((PrismFightJoinLeaveRequestMessage)(PrismFightJoinLeaveRequestMessage{})),

	5844: reflect.TypeOf((PrismInfoJoinLeaveRequestMessage)(PrismInfoJoinLeaveRequestMessage{})),

	5853: reflect.TypeOf((PrismInfoCloseMessage)(PrismInfoCloseMessage{})),

	5859: reflect.TypeOf((PrismInfoInValidMessage)(PrismInfoInValidMessage{})),

	5864: reflect.TypeOf((GameFightShowFighterMessage)(GameFightShowFighterMessage{})),

	5884: reflect.TypeOf((HouseSellFromInsideRequestMessage)(HouseSellFromInsideRequestMessage{})),

	5885: reflect.TypeOf((HouseLockFromInsideRequestMessage)(HouseLockFromInsideRequestMessage{})),

	5887: reflect.TypeOf((GuildKickRequestMessage)(GuildKickRequestMessage{})),

	5892: reflect.TypeOf((PrismFightDefenderLeaveMessage)(PrismFightDefenderLeaveMessage{})),

	5893: reflect.TypeOf((PrismFightAttackerAddMessage)(PrismFightAttackerAddMessage{})),

	5895: reflect.TypeOf((PrismFightDefenderAddMessage)(PrismFightDefenderAddMessage{})),

	5897: reflect.TypeOf((PrismFightAttackerRemoveMessage)(PrismFightAttackerRemoveMessage{})),

	5898: reflect.TypeOf((NpcGenericActionRequestMessage)(NpcGenericActionRequestMessage{})),

	5900: reflect.TypeOf((NpcGenericActionFailureMessage)(NpcGenericActionFailureMessage{})),

	5901: reflect.TypeOf((PrismFightSwapRequestMessage)(PrismFightSwapRequestMessage{})),

	5904: reflect.TypeOf((ExchangeStartedBidBuyerMessage)(ExchangeStartedBidBuyerMessage{})),

	5905: reflect.TypeOf((ExchangeStartedBidSellerMessage)(ExchangeStartedBidSellerMessage{})),

	5907: reflect.TypeOf((ExchangeShopStockMovementRemovedMessage)(ExchangeShopStockMovementRemovedMessage{})),

	5908: reflect.TypeOf((ChallengeFightJoinRefusedMessage)(ChallengeFightJoinRefusedMessage{})),

	5909: reflect.TypeOf((ExchangeShopStockMovementUpdatedMessage)(ExchangeShopStockMovementUpdatedMessage{})),

	5910: reflect.TypeOf((ExchangeShopStockStartedMessage)(ExchangeShopStockStartedMessage{})),

	5915: reflect.TypeOf((TaxCollectorMovementRemoveMessage)(TaxCollectorMovementRemoveMessage{})),

	5917: reflect.TypeOf((TaxCollectorMovementAddMessage)(TaxCollectorMovementAddMessage{})),

	5918: reflect.TypeOf((TaxCollectorAttackedMessage)(TaxCollectorAttackedMessage{})),

	5919: reflect.TypeOf((GuildHousesInformationMessage)(GuildHousesInformationMessage{})),

	5920: reflect.TypeOf((GuildCreationStartedMessage)(GuildCreationStartedMessage{})),

	5921: reflect.TypeOf((GameFightSynchronizeMessage)(GameFightSynchronizeMessage{})),

	5923: reflect.TypeOf((GuildMemberLeavingMessage)(GuildMemberLeavingMessage{})),

	5924: reflect.TypeOf((FriendUpdateMessage)(FriendUpdateMessage{})),

	5925: reflect.TypeOf((ObjectGroundListAddedMessage)(ObjectGroundListAddedMessage{})),

	5927: reflect.TypeOf((GameFightOptionStateUpdateMessage)(GameFightOptionStateUpdateMessage{})),

	5928: reflect.TypeOf((GuildFightPlayersEnemiesListMessage)(GuildFightPlayersEnemiesListMessage{})),

	5929: reflect.TypeOf((GuildFightPlayersEnemyRemoveMessage)(GuildFightPlayersEnemyRemoveMessage{})),

	5930: reflect.TypeOf((TaxCollectorListMessage)(TaxCollectorListMessage{})),

	5932: reflect.TypeOf((ContactLookRequestMessage)(ContactLookRequestMessage{})),

	5933: reflect.TypeOf((ContactLookRequestByNameMessage)(ContactLookRequestByNameMessage{})),

	5934: reflect.TypeOf((ContactLookMessage)(ContactLookMessage{})),

	5935: reflect.TypeOf((ContactLookRequestByIdMessage)(ContactLookRequestByIdMessage{})),

	5937: reflect.TypeOf((GameRolePlayPlayerFightFriendlyRequestedMessage)(GameRolePlayPlayerFightFriendlyRequestedMessage{})),

	5941: reflect.TypeOf((ExchangeStartOkCraftWithInformationMessage)(ExchangeStartOkCraftWithInformationMessage{})),

	5944: reflect.TypeOf((ObjectGroundRemovedMultipleMessage)(ObjectGroundRemovedMultipleMessage{})),

	5945: reflect.TypeOf((ExchangeBidHouseItemAddOkMessage)(ExchangeBidHouseItemAddOkMessage{})),

	5946: reflect.TypeOf((ExchangeBidHouseItemRemoveOkMessage)(ExchangeBidHouseItemRemoveOkMessage{})),

	5947: reflect.TypeOf((ExchangeBidHouseGenericItemAddedMessage)(ExchangeBidHouseGenericItemAddedMessage{})),

	5948: reflect.TypeOf((ExchangeBidHouseGenericItemRemovedMessage)(ExchangeBidHouseGenericItemRemovedMessage{})),

	5949: reflect.TypeOf((ExchangeBidHouseInListAddedMessage)(ExchangeBidHouseInListAddedMessage{})),

	5950: reflect.TypeOf((ExchangeBidHouseInListRemovedMessage)(ExchangeBidHouseInListRemovedMessage{})),

	5951: reflect.TypeOf((PaddockBuyRequestMessage)(PaddockBuyRequestMessage{})),

	5952: reflect.TypeOf((GuildPaddockBoughtMessage)(GuildPaddockBoughtMessage{})),

	5953: reflect.TypeOf((PaddockSellRequestMessage)(PaddockSellRequestMessage{})),

	5954: reflect.TypeOf((GameRolePlayTaxCollectorFightRequestMessage)(GameRolePlayTaxCollectorFightRequestMessage{})),

	5955: reflect.TypeOf((GuildPaddockRemovedMessage)(GuildPaddockRemovedMessage{})),

	5956: reflect.TypeOf((AtlasPointInformationsMessage)(AtlasPointInformationsMessage{})),

	5957: reflect.TypeOf((GuildPaddockTeleportRequestMessage)(GuildPaddockTeleportRequestMessage{})),

	5958: reflect.TypeOf((PaddockRemoveItemRequestMessage)(PaddockRemoveItemRequestMessage{})),

	5959: reflect.TypeOf((GuildInformationsPaddocksMessage)(GuildInformationsPaddocksMessage{})),

	5960: reflect.TypeOf((TeleportDestinationsListMessage)(TeleportDestinationsListMessage{})),

	5961: reflect.TypeOf((TeleportRequestMessage)(TeleportRequestMessage{})),

	5962: reflect.TypeOf((MountSterilizeRequestMessage)(MountSterilizeRequestMessage{})),

	5963: reflect.TypeOf((MountEquipedErrorMessage)(MountEquipedErrorMessage{})),

	5967: reflect.TypeOf((MountRidingMessage)(MountRidingMessage{})),

	5968: reflect.TypeOf((MountSetMessage)(MountSetMessage{})),

	5970: reflect.TypeOf((MountXpRatioMessage)(MountXpRatioMessage{})),

	5972: reflect.TypeOf((MountInformationRequestMessage)(MountInformationRequestMessage{})),

	5973: reflect.TypeOf((MountDataMessage)(MountDataMessage{})),

	5975: reflect.TypeOf((MountInformationInPaddockRequestMessage)(MountInformationInPaddockRequestMessage{})),

	5976: reflect.TypeOf((MountToggleRidingRequestMessage)(MountToggleRidingRequestMessage{})),

	5977: reflect.TypeOf((MountSterilizedMessage)(MountSterilizedMessage{})),

	5978: reflect.TypeOf((MountEmoteIconUsedOkMessage)(MountEmoteIconUsedOkMessage{})),

	5979: reflect.TypeOf((ExchangeStartOkMountMessage)(ExchangeStartOkMountMessage{})),

	5980: reflect.TypeOf((MountReleaseRequestMessage)(MountReleaseRequestMessage{})),

	5981: reflect.TypeOf((ExchangeMountStableErrorMessage)(ExchangeMountStableErrorMessage{})),

	5982: reflect.TypeOf((MountUnSetMessage)(MountUnSetMessage{})),

	5983: reflect.TypeOf((MountRenamedMessage)(MountRenamedMessage{})),

	5984: reflect.TypeOf((ExchangeStartedMountStockMessage)(ExchangeStartedMountStockMessage{})),

	5986: reflect.TypeOf((ExchangeRequestOnMountStockMessage)(ExchangeRequestOnMountStockMessage{})),

	5987: reflect.TypeOf((MountRenameRequestMessage)(MountRenameRequestMessage{})),

	5989: reflect.TypeOf((MountSetXpRatioRequestMessage)(MountSetXpRatioRequestMessage{})),

	5990: reflect.TypeOf((GameDataPaddockObjectAddMessage)(GameDataPaddockObjectAddMessage{})),

	5991: reflect.TypeOf((ExchangeStartOkMountWithOutPaddockMessage)(ExchangeStartOkMountWithOutPaddockMessage{})),

	5992: reflect.TypeOf((GameDataPaddockObjectListAddMessage)(GameDataPaddockObjectListAddMessage{})),

	5993: reflect.TypeOf((GameDataPaddockObjectRemoveMessage)(GameDataPaddockObjectRemoveMessage{})),

	5996: reflect.TypeOf((GameRolePlayPlayerLifeStatusMessage)(GameRolePlayPlayerLifeStatusMessage{})),

	5999: reflect.TypeOf((ExchangeCraftResultWithObjectDescMessage)(ExchangeCraftResultWithObjectDescMessage{})),

	6000: reflect.TypeOf((ExchangeCraftResultWithObjectIdMessage)(ExchangeCraftResultWithObjectIdMessage{})),

	6001: reflect.TypeOf((ExchangeReplayStopMessage)(ExchangeReplayStopMessage{})),

	6004: reflect.TypeOf((ExchangeObjectUseInWorkshopMessage)(ExchangeObjectUseInWorkshopMessage{})),

	6008: reflect.TypeOf((ExchangeObjectModifiedInBagMessage)(ExchangeObjectModifiedInBagMessage{})),

	6009: reflect.TypeOf((ExchangeObjectPutInBagMessage)(ExchangeObjectPutInBagMessage{})),

	6010: reflect.TypeOf((ExchangeObjectRemovedFromBagMessage)(ExchangeObjectRemovedFromBagMessage{})),

	6011: reflect.TypeOf((SpellItemBoostMessage)(SpellItemBoostMessage{})),

	6012: reflect.TypeOf((PauseDialogMessage)(PauseDialogMessage{})),

	6013: reflect.TypeOf((CompassUpdatePvpSeekMessage)(CompassUpdatePvpSeekMessage{})),

	6014: reflect.TypeOf((ObjectJobAddedMessage)(ObjectJobAddedMessage{})),

	6018: reflect.TypeOf((PaddockSellBuyDialogMessage)(PaddockSellBuyDialogMessage{})),

	6019: reflect.TypeOf((ChallengeResultMessage)(ChallengeResultMessage{})),

	6020: reflect.TypeOf((ExchangeMultiCraftCrafterCanUseHisRessourcesMessage)(ExchangeMultiCraftCrafterCanUseHisRessourcesMessage{})),

	6021: reflect.TypeOf((ExchangeMultiCraftSetCrafterCanUseHisRessourcesMessage)(ExchangeMultiCraftSetCrafterCanUseHisRessourcesMessage{})),

	6022: reflect.TypeOf((ChallengeInfoMessage)(ChallengeInfoMessage{})),

	6024: reflect.TypeOf((GameContextCreateErrorMessage)(GameContextCreateErrorMessage{})),

	6026: reflect.TypeOf((GameDataPlayFarmObjectAnimationMessage)(GameDataPlayFarmObjectAnimationMessage{})),

	6028: reflect.TypeOf((DebugInClientMessage)(DebugInClientMessage{})),

	6029: reflect.TypeOf((AccountLoggingKickedMessage)(AccountLoggingKickedMessage{})),

	6030: reflect.TypeOf((GoldAddedMessage)(GoldAddedMessage{})),

	6032: reflect.TypeOf((ExchangeObjectTransfertAllToInvMessage)(ExchangeObjectTransfertAllToInvMessage{})),

	6033: reflect.TypeOf((ObjectsAddedMessage)(ObjectsAddedMessage{})),

	6034: reflect.TypeOf((ObjectsDeletedMessage)(ObjectsDeletedMessage{})),

	6035: reflect.TypeOf((StorageObjectsRemoveMessage)(StorageObjectsRemoveMessage{})),

	6036: reflect.TypeOf((StorageObjectsUpdateMessage)(StorageObjectsUpdateMessage{})),

	6037: reflect.TypeOf((ExchangeShopStockMultiMovementRemovedMessage)(ExchangeShopStockMultiMovementRemovedMessage{})),

	6038: reflect.TypeOf((ExchangeShopStockMultiMovementUpdatedMessage)(ExchangeShopStockMultiMovementUpdatedMessage{})),

	6039: reflect.TypeOf((ExchangeObjectTransfertListToInvMessage)(ExchangeObjectTransfertListToInvMessage{})),

	6040: reflect.TypeOf((PrismFightStateUpdateMessage)(PrismFightStateUpdateMessage{})),

	6041: reflect.TypeOf((PrismUseRequestMessage)(PrismUseRequestMessage{})),

	6042: reflect.TypeOf((PrismAttackRequestMessage)(PrismAttackRequestMessage{})),

	6043: reflect.TypeOf((JobCrafterDirectoryEntryRequestMessage)(JobCrafterDirectoryEntryRequestMessage{})),

	6044: reflect.TypeOf((JobCrafterDirectoryEntryMessage)(JobCrafterDirectoryEntryMessage{})),

	6045: reflect.TypeOf((ContactLookErrorMessage)(ContactLookErrorMessage{})),

	6046: reflect.TypeOf((JobCrafterDirectoryListMessage)(JobCrafterDirectoryListMessage{})),

	6047: reflect.TypeOf((JobCrafterDirectoryListRequestMessage)(JobCrafterDirectoryListRequestMessage{})),

	6048: reflect.TypeOf((TeleportOnSameMapMessage)(TeleportOnSameMapMessage{})),

	6051: reflect.TypeOf((MapObstacleUpdateMessage)(MapObstacleUpdateMessage{})),

	6052: reflect.TypeOf((PaddockMoveItemRequestMessage)(PaddockMoveItemRequestMessage{})),

	6053: reflect.TypeOf((CinematicMessage)(CinematicMessage{})),

	6054: reflect.TypeOf((PartyUpdateLightMessage)(PartyUpdateLightMessage{})),

	6055: reflect.TypeOf((ExchangeMountFreeFromPaddockMessage)(ExchangeMountFreeFromPaddockMessage{})),

	6056: reflect.TypeOf((ExchangeMountSterilizeFromPaddockMessage)(ExchangeMountSterilizeFromPaddockMessage{})),

	6058: reflect.TypeOf((AlignmentRankUpdateMessage)(AlignmentRankUpdateMessage{})),

	6061: reflect.TypeOf((GuildMemberOnlineStatusMessage)(GuildMemberOnlineStatusMessage{})),

	6062: reflect.TypeOf((GuildLevelUpMessage)(GuildLevelUpMessage{})),

	6065: reflect.TypeOf((LivingObjectMessageMessage)(LivingObjectMessageMessage{})),

	6066: reflect.TypeOf((LivingObjectMessageRequestMessage)(LivingObjectMessageRequestMessage{})),

	6067: reflect.TypeOf((GameFightResumeMessage)(GameFightResumeMessage{})),

	6068: reflect.TypeOf((CharacterSelectedForceMessage)(CharacterSelectedForceMessage{})),

	6069: reflect.TypeOf((GameFightSpectateMessage)(GameFightSpectateMessage{})),

	6070: reflect.TypeOf((GameActionFightDispellableEffectMessage)(GameActionFightDispellableEffectMessage{})),

	6071: reflect.TypeOf((GameContextReadyMessage)(GameContextReadyMessage{})),

	6072: reflect.TypeOf((CharacterSelectedForceReadyMessage)(CharacterSelectedForceReadyMessage{})),

	6073: reflect.TypeOf((GameRolePlayAggressionMessage)(GameRolePlayAggressionMessage{})),

	6076: reflect.TypeOf((CharacterLevelUpInformationMessage)(CharacterLevelUpInformationMessage{})),

	6077: reflect.TypeOf((FriendSetWarnOnLevelGainMessage)(FriendSetWarnOnLevelGainMessage{})),

	6078: reflect.TypeOf((FriendWarnOnLevelGainStateMessage)(FriendWarnOnLevelGainStateMessage{})),

	6079: reflect.TypeOf((CharacterReportMessage)(CharacterReportMessage{})),

	6080: reflect.TypeOf((PartyAbdicateThroneMessage)(PartyAbdicateThroneMessage{})),

	6081: reflect.TypeOf((GameContextKickMessage)(GameContextKickMessage{})),

	6084: reflect.TypeOf((CharacterFirstSelectionMessage)(CharacterFirstSelectionMessage{})),

	6085: reflect.TypeOf((QuestObjectiveValidationMessage)(QuestObjectiveValidationMessage{})),

	6087: reflect.TypeOf((NotificationListMessage)(NotificationListMessage{})),

	6088: reflect.TypeOf((GuidedModeReturnRequestMessage)(GuidedModeReturnRequestMessage{})),

	6089: reflect.TypeOf((NotificationResetMessage)(NotificationResetMessage{})),

	6090: reflect.TypeOf((NotificationUpdateFlagMessage)(NotificationUpdateFlagMessage{})),

	6091: reflect.TypeOf((QuestStartedMessage)(QuestStartedMessage{})),

	6092: reflect.TypeOf((GuidedModeQuitRequestMessage)(GuidedModeQuitRequestMessage{})),

	6096: reflect.TypeOf((QuestStepStartedMessage)(QuestStepStartedMessage{})),

	6097: reflect.TypeOf((QuestValidatedMessage)(QuestValidatedMessage{})),

	6098: reflect.TypeOf((QuestObjectiveValidatedMessage)(QuestObjectiveValidatedMessage{})),

	6099: reflect.TypeOf((QuestStepValidatedMessage)(QuestStepValidatedMessage{})),

	6100: reflect.TypeOf((QueueStatusMessage)(QueueStatusMessage{})),

	6103: reflect.TypeOf((NotificationByServerMessage)(NotificationByServerMessage{})),

	6110: reflect.TypeOf((EntityTalkMessage)(EntityTalkMessage{})),

	6112: reflect.TypeOf((InteractiveUseEndedMessage)(InteractiveUseEndedMessage{})),

	6113: reflect.TypeOf((GameActionFightDispellEffectMessage)(GameActionFightDispellEffectMessage{})),

	6114: reflect.TypeOf((GameRolePlaySpellAnimMessage)(GameRolePlaySpellAnimMessage{})),

	6115: reflect.TypeOf((GuildInvitationByNameMessage)(GuildInvitationByNameMessage{})),

	6116: reflect.TypeOf((GameActionFightCloseCombatMessage)(GameActionFightCloseCombatMessage{})),

	6118: reflect.TypeOf((AbstractGameActionFightTargetedAbilityMessage)(AbstractGameActionFightTargetedAbilityMessage{})),

	6119: reflect.TypeOf((IdentificationAccountForceMessage)(IdentificationAccountForceMessage{})),

	6120: reflect.TypeOf((CharactersListWithModificationsMessage)(CharactersListWithModificationsMessage{})),

	6123: reflect.TypeOf((ChallengeTargetUpdateMessage)(ChallengeTargetUpdateMessage{})),

	6124: reflect.TypeOf((StopToListenRunningFightRequestMessage)(StopToListenRunningFightRequestMessage{})),

	6127: reflect.TypeOf((ConsoleCommandsListMessage)(ConsoleCommandsListMessage{})),

	6129: reflect.TypeOf((ExchangeStartedWithPodsMessage)(ExchangeStartedWithPodsMessage{})),

	6130: reflect.TypeOf((MapComplementaryInformationsDataInHouseMessage)(MapComplementaryInformationsDataInHouseMessage{})),

	6132: reflect.TypeOf((GameActionFightNoSpellCastMessage)(GameActionFightNoSpellCastMessage{})),

	6134: reflect.TypeOf((PopupWarningMessage)(PopupWarningMessage{})),

	6135: reflect.TypeOf((ChatAdminServerMessage)(ChatAdminServerMessage{})),

	6137: reflect.TypeOf((HouseToSellFilterMessage)(HouseToSellFilterMessage{})),

	6138: reflect.TypeOf((PaddockToSellListMessage)(PaddockToSellListMessage{})),

	6139: reflect.TypeOf((HouseToSellListRequestMessage)(HouseToSellListRequestMessage{})),

	6140: reflect.TypeOf((HouseToSellListMessage)(HouseToSellListMessage{})),

	6141: reflect.TypeOf((PaddockToSellListRequestMessage)(PaddockToSellListRequestMessage{})),

	6142: reflect.TypeOf((AcquaintanceServerListMessage)(AcquaintanceServerListMessage{})),

	6143: reflect.TypeOf((AcquaintanceSearchErrorMessage)(AcquaintanceSearchErrorMessage{})),

	6144: reflect.TypeOf((AcquaintanceSearchMessage)(AcquaintanceSearchMessage{})),

	6147: reflect.TypeOf((GameActionFightTriggerEffectMessage)(GameActionFightTriggerEffectMessage{})),

	6150: reflect.TypeOf((GameRolePlayDelayedActionFinishedMessage)(GameRolePlayDelayedActionFinishedMessage{})),

	6153: reflect.TypeOf((GameRolePlayDelayedActionMessage)(GameRolePlayDelayedActionMessage{})),

	6154: reflect.TypeOf((CheckFileRequestMessage)(CheckFileRequestMessage{})),

	6155: reflect.TypeOf((GameMapChangeOrientationsMessage)(GameMapChangeOrientationsMessage{})),

	6156: reflect.TypeOf((CheckFileMessage)(CheckFileMessage{})),

	6158: reflect.TypeOf((ShowCellSpectatorMessage)(ShowCellSpectatorMessage{})),

	6159: reflect.TypeOf((GuildMemberSetWarnOnConnectionMessage)(GuildMemberSetWarnOnConnectionMessage{})),

	6160: reflect.TypeOf((GuildMemberWarnOnConnectionStateMessage)(GuildMemberWarnOnConnectionStateMessage{})),

	6161: reflect.TypeOf((PaddockToSellFilterMessage)(PaddockToSellFilterMessage{})),

	6162: reflect.TypeOf((InventoryContentAndPresetMessage)(InventoryContentAndPresetMessage{})),

	6163: reflect.TypeOf((InventoryPresetUseResultMessage)(InventoryPresetUseResultMessage{})),

	6165: reflect.TypeOf((InventoryPresetSaveMessage)(InventoryPresetSaveMessage{})),

	6167: reflect.TypeOf((InventoryPresetUseMessage)(InventoryPresetUseMessage{})),

	6168: reflect.TypeOf((InventoryPresetItemUpdateMessage)(InventoryPresetItemUpdateMessage{})),

	6169: reflect.TypeOf((InventoryPresetDeleteMessage)(InventoryPresetDeleteMessage{})),

	6170: reflect.TypeOf((InventoryPresetSaveResultMessage)(InventoryPresetSaveResultMessage{})),

	6171: reflect.TypeOf((InventoryPresetUpdateMessage)(InventoryPresetUpdateMessage{})),

	6172: reflect.TypeOf((MountDataErrorMessage)(MountDataErrorMessage{})),

	6173: reflect.TypeOf((InventoryPresetDeleteResultMessage)(InventoryPresetDeleteResultMessage{})),

	6174: reflect.TypeOf((IdentificationFailedBannedMessage)(IdentificationFailedBannedMessage{})),

	6175: reflect.TypeOf((PartyRestrictedMessage)(PartyRestrictedMessage{})),

	6176: reflect.TypeOf((GameActionFightDispellSpellMessage)(GameActionFightDispellSpellMessage{})),

	6179: reflect.TypeOf((UpdateMountBoostMessage)(UpdateMountBoostMessage{})),

	6180: reflect.TypeOf((GuildHouseRemoveMessage)(GuildHouseRemoveMessage{})),

	6181: reflect.TypeOf((GuildHouseUpdateInformationMessage)(GuildHouseUpdateInformationMessage{})),

	6183: reflect.TypeOf((ExchangeObjectTransfertListFromInvMessage)(ExchangeObjectTransfertListFromInvMessage{})),

	6184: reflect.TypeOf((ExchangeObjectTransfertAllFromInvMessage)(ExchangeObjectTransfertAllFromInvMessage{})),

	6185: reflect.TypeOf((LocalizedChatSmileyMessage)(LocalizedChatSmileyMessage{})),

	6188: reflect.TypeOf((ExchangeCraftResultMagicWithObjectDescMessage)(ExchangeCraftResultMagicWithObjectDescMessage{})),

	6189: reflect.TypeOf((MountFeedRequestMessage)(MountFeedRequestMessage{})),

	6191: reflect.TypeOf((GameRolePlayAttackMonsterRequestMessage)(GameRolePlayAttackMonsterRequestMessage{})),

	6192: reflect.TypeOf((MoodSmileyRequestMessage)(MoodSmileyRequestMessage{})),

	6196: reflect.TypeOf((MoodSmileyResultMessage)(MoodSmileyResultMessage{})),

	6197: reflect.TypeOf((ErrorMapNotFoundMessage)(ErrorMapNotFoundMessage{})),

	6205: reflect.TypeOf((AchievementListMessage)(AchievementListMessage{})),

	6206: reflect.TypeOf((ObjectsQuantityMessage)(ObjectsQuantityMessage{})),

	6208: reflect.TypeOf((AchievementFinishedMessage)(AchievementFinishedMessage{})),

	6209: reflect.TypeOf((IdentificationSuccessWithLoginTokenMessage)(IdentificationSuccessWithLoginTokenMessage{})),

	6210: reflect.TypeOf((InventoryPresetItemUpdateRequestMessage)(InventoryPresetItemUpdateRequestMessage{})),

	6211: reflect.TypeOf((InventoryPresetItemUpdateErrorMessage)(InventoryPresetItemUpdateErrorMessage{})),

	6214: reflect.TypeOf((SlaveSwitchContextMessage)(SlaveSwitchContextMessage{})),

	6215: reflect.TypeOf((GameFightResumeWithSlavesMessage)(GameFightResumeWithSlavesMessage{})),

	6216: reflect.TypeOf((AccountCapabilitiesMessage)(AccountCapabilitiesMessage{})),

	6217: reflect.TypeOf((GameActionFightVanishMessage)(GameActionFightVanishMessage{})),

	6218: reflect.TypeOf((GameFightShowFighterRandomStaticPoseMessage)(GameFightShowFighterRandomStaticPoseMessage{})),

	6219: reflect.TypeOf((GameActionFightSpellCooldownVariationMessage)(GameActionFightSpellCooldownVariationMessage{})),

	6221: reflect.TypeOf((GameActionFightSpellImmunityMessage)(GameActionFightSpellImmunityMessage{})),

	6222: reflect.TypeOf((ShortcutBarRemoveErrorMessage)(ShortcutBarRemoveErrorMessage{})),

	6224: reflect.TypeOf((ShortcutBarRemovedMessage)(ShortcutBarRemovedMessage{})),

	6225: reflect.TypeOf((ShortcutBarAddRequestMessage)(ShortcutBarAddRequestMessage{})),

	6226: reflect.TypeOf((ShortcutBarSwapErrorMessage)(ShortcutBarSwapErrorMessage{})),

	6227: reflect.TypeOf((ShortcutBarAddErrorMessage)(ShortcutBarAddErrorMessage{})),

	6228: reflect.TypeOf((ShortcutBarRemoveRequestMessage)(ShortcutBarRemoveRequestMessage{})),

	6229: reflect.TypeOf((ShortcutBarRefreshMessage)(ShortcutBarRefreshMessage{})),

	6230: reflect.TypeOf((ShortcutBarSwapRequestMessage)(ShortcutBarSwapRequestMessage{})),

	6231: reflect.TypeOf((ShortcutBarContentMessage)(ShortcutBarContentMessage{})),

	6234: reflect.TypeOf((ObjectUseMultipleMessage)(ObjectUseMultipleMessage{})),

	6235: reflect.TypeOf((GuildFightTakePlaceRequestMessage)(GuildFightTakePlaceRequestMessage{})),

	6236: reflect.TypeOf((ExchangeStartedWithStorageMessage)(ExchangeStartedWithStorageMessage{})),

	6237: reflect.TypeOf((ServerExperienceModificatorMessage)(ServerExperienceModificatorMessage{})),

	6238: reflect.TypeOf((ExchangeObjectModifyPricedMessage)(ExchangeObjectModifyPricedMessage{})),

	6239: reflect.TypeOf((GameFightNewRoundMessage)(GameFightNewRoundMessage{})),

	6240: reflect.TypeOf((DungeonPartyFinderAvailableDungeonsRequestMessage)(DungeonPartyFinderAvailableDungeonsRequestMessage{})),

	6241: reflect.TypeOf((DungeonPartyFinderRegisterSuccessMessage)(DungeonPartyFinderRegisterSuccessMessage{})),

	6242: reflect.TypeOf((DungeonPartyFinderAvailableDungeonsMessage)(DungeonPartyFinderAvailableDungeonsMessage{})),

	6243: reflect.TypeOf((DungeonPartyFinderRegisterErrorMessage)(DungeonPartyFinderRegisterErrorMessage{})),

	6244: reflect.TypeOf((PartyInvitationDungeonMessage)(PartyInvitationDungeonMessage{})),

	6245: reflect.TypeOf((PartyInvitationDungeonRequestMessage)(PartyInvitationDungeonRequestMessage{})),

	6246: reflect.TypeOf((DungeonPartyFinderListenRequestMessage)(DungeonPartyFinderListenRequestMessage{})),

	6247: reflect.TypeOf((DungeonPartyFinderRoomContentMessage)(DungeonPartyFinderRoomContentMessage{})),

	6248: reflect.TypeOf((DungeonPartyFinderListenErrorMessage)(DungeonPartyFinderListenErrorMessage{})),

	6249: reflect.TypeOf((DungeonPartyFinderRegisterRequestMessage)(DungeonPartyFinderRegisterRequestMessage{})),

	6250: reflect.TypeOf((DungeonPartyFinderRoomContentUpdateMessage)(DungeonPartyFinderRoomContentUpdateMessage{})),

	6251: reflect.TypeOf((PartyCancelInvitationNotificationMessage)(PartyCancelInvitationNotificationMessage{})),

	6252: reflect.TypeOf((PartyMemberEjectedMessage)(PartyMemberEjectedMessage{})),

	6253: reflect.TypeOf((RawDataMessage)(RawDataMessage{})),

	6254: reflect.TypeOf((PartyCancelInvitationMessage)(PartyCancelInvitationMessage{})),

	6256: reflect.TypeOf((PartyInvitationCancelledForGuestMessage)(PartyInvitationCancelledForGuestMessage{})),

	6260: reflect.TypeOf((PartyNewGuestMessage)(PartyNewGuestMessage{})),

	6261: reflect.TypeOf((PartyDeletedMessage)(PartyDeletedMessage{})),

	6262: reflect.TypeOf((PartyInvitationDungeonDetailsMessage)(PartyInvitationDungeonDetailsMessage{})),

	6263: reflect.TypeOf((PartyInvitationDetailsMessage)(PartyInvitationDetailsMessage{})),

	6264: reflect.TypeOf((PartyInvitationDetailsRequestMessage)(PartyInvitationDetailsRequestMessage{})),

	6265: reflect.TypeOf((SpouseStatusMessage)(SpouseStatusMessage{})),

	6266: reflect.TypeOf((URLOpenMessage)(URLOpenMessage{})),

	6267: reflect.TypeOf((TrustStatusMessage)(TrustStatusMessage{})),

	6268: reflect.TypeOf((MapComplementaryInformationsWithCoordsMessage)(MapComplementaryInformationsWithCoordsMessage{})),

	6269: reflect.TypeOf((PartyPledgeLoyaltyRequestMessage)(PartyPledgeLoyaltyRequestMessage{})),

	6270: reflect.TypeOf((PartyLoyaltyStatusMessage)(PartyLoyaltyStatusMessage{})),

	6272: reflect.TypeOf((ExchangeBidHouseBuyResultMessage)(ExchangeBidHouseBuyResultMessage{})),

	6273: reflect.TypeOf((AbstractPartyEventMessage)(AbstractPartyEventMessage{})),

	6274: reflect.TypeOf((AbstractPartyMessage)(AbstractPartyMessage{})),

	6275: reflect.TypeOf((MailStatusMessage)(MailStatusMessage{})),

	6276: reflect.TypeOf((GameRolePlayArenaFightPropositionMessage)(GameRolePlayArenaFightPropositionMessage{})),

	6277: reflect.TypeOf((PartyModifiableStatusMessage)(PartyModifiableStatusMessage{})),

	6279: reflect.TypeOf((GameRolePlayArenaFightAnswerMessage)(GameRolePlayArenaFightAnswerMessage{})),

	6280: reflect.TypeOf((GameRolePlayArenaRegisterMessage)(GameRolePlayArenaRegisterMessage{})),

	6281: reflect.TypeOf((GameRolePlayArenaFighterStatusMessage)(GameRolePlayArenaFighterStatusMessage{})),

	6282: reflect.TypeOf((GameRolePlayArenaUnregisterMessage)(GameRolePlayArenaUnregisterMessage{})),

	6283: reflect.TypeOf((PartyInvitationArenaRequestMessage)(PartyInvitationArenaRequestMessage{})),

	6284: reflect.TypeOf((GameRolePlayArenaRegistrationStatusMessage)(GameRolePlayArenaRegistrationStatusMessage{})),

	6287: reflect.TypeOf((TeleportToBuddyOfferMessage)(TeleportToBuddyOfferMessage{})),

	6289: reflect.TypeOf((TeleportBuddiesMessage)(TeleportBuddiesMessage{})),

	6290: reflect.TypeOf((ObjectFeedMessage)(ObjectFeedMessage{})),

	6292: reflect.TypeOf((NewMailMessage)(NewMailMessage{})),

	6293: reflect.TypeOf((TeleportToBuddyAnswerMessage)(TeleportToBuddyAnswerMessage{})),

	6294: reflect.TypeOf((TeleportBuddiesAnswerMessage)(TeleportBuddiesAnswerMessage{})),

	6296: reflect.TypeOf((DungeonKeyRingUpdateMessage)(DungeonKeyRingUpdateMessage{})),

	6297: reflect.TypeOf((NumericWhoIsMessage)(NumericWhoIsMessage{})),

	6298: reflect.TypeOf((NumericWhoIsRequestMessage)(NumericWhoIsRequestMessage{})),

	6299: reflect.TypeOf((DungeonKeyRingMessage)(DungeonKeyRingMessage{})),

	6301: reflect.TypeOf((GameRolePlayArenaUpdatePlayerInfosMessage)(GameRolePlayArenaUpdatePlayerInfosMessage{})),

	6302: reflect.TypeOf((TeleportBuddiesRequestedMessage)(TeleportBuddiesRequestedMessage{})),

	6303: reflect.TypeOf((TeleportToBuddyCloseMessage)(TeleportToBuddyCloseMessage{})),

	6304: reflect.TypeOf((GameActionFightModifyEffectsDurationMessage)(GameActionFightModifyEffectsDurationMessage{})),

	6305: reflect.TypeOf((ServerOptionalFeaturesMessage)(ServerOptionalFeaturesMessage{})),

	6306: reflect.TypeOf((PartyNewMemberMessage)(PartyNewMemberMessage{})),

	6307: reflect.TypeOf((GameFightTurnResumeMessage)(GameFightTurnResumeMessage{})),

	6308: reflect.TypeOf((MountReleasedMessage)(MountReleasedMessage{})),

	6309: reflect.TypeOf((GameFightRefreshFighterMessage)(GameFightRefreshFighterMessage{})),

	6310: reflect.TypeOf((GameActionFightLifeAndShieldPointsLostMessage)(GameActionFightLifeAndShieldPointsLostMessage{})),

	6311: reflect.TypeOf((GameActionFightLifePointsGainMessage)(GameActionFightLifePointsGainMessage{})),

	6312: reflect.TypeOf((GameActionFightLifePointsLostMessage)(GameActionFightLifePointsLostMessage{})),

	6314: reflect.TypeOf((CredentialsAcknowledgementMessage)(CredentialsAcknowledgementMessage{})),

	6315: reflect.TypeOf((AccountHouseMessage)(AccountHouseMessage{})),

	6316: reflect.TypeOf((SequenceNumberRequestMessage)(SequenceNumberRequestMessage{})),

	6317: reflect.TypeOf((SequenceNumberMessage)(SequenceNumberMessage{})),

	6320: reflect.TypeOf((GameActionFightInvisibleDetectedMessage)(GameActionFightInvisibleDetectedMessage{})),

	6321: reflect.TypeOf((CharacterExperienceGainMessage)(CharacterExperienceGainMessage{})),

	6322: reflect.TypeOf((FighterStatsListMessage)(FighterStatsListMessage{})),

	6323: reflect.TypeOf((GuildModificationValidMessage)(GuildModificationValidMessage{})),

	6324: reflect.TypeOf((GuildModificationStartedMessage)(GuildModificationStartedMessage{})),

	6325: reflect.TypeOf((ExchangeObjectTransfertExistingFromInvMessage)(ExchangeObjectTransfertExistingFromInvMessage{})),

	6326: reflect.TypeOf((ExchangeObjectTransfertExistingToInvMessage)(ExchangeObjectTransfertExistingToInvMessage{})),

	6327: reflect.TypeOf((GuildModificationNameValidMessage)(GuildModificationNameValidMessage{})),

	6328: reflect.TypeOf((GuildModificationEmblemValidMessage)(GuildModificationEmblemValidMessage{})),

	6329: reflect.TypeOf((InventoryPresetSaveCustomMessage)(InventoryPresetSaveCustomMessage{})),

	6330: reflect.TypeOf((GameActionFightCastOnTargetRequestMessage)(GameActionFightCastOnTargetRequestMessage{})),

	6334: reflect.TypeOf((ObjectAveragePricesGetMessage)(ObjectAveragePricesGetMessage{})),

	6335: reflect.TypeOf((ObjectAveragePricesMessage)(ObjectAveragePricesMessage{})),

	6336: reflect.TypeOf((ObjectAveragePricesErrorMessage)(ObjectAveragePricesErrorMessage{})),

	6337: reflect.TypeOf((ExchangeBidHouseInListUpdatedMessage)(ExchangeBidHouseInListUpdatedMessage{})),

	6339: reflect.TypeOf((CharacterCapabilitiesMessage)(CharacterCapabilitiesMessage{})),

	6340: reflect.TypeOf((ServerSettingsMessage)(ServerSettingsMessage{})),

	6341: reflect.TypeOf((AlmanachCalendarDateMessage)(AlmanachCalendarDateMessage{})),

	6342: reflect.TypeOf((PartyMemberInFightMessage)(PartyMemberInFightMessage{})),

	6343: reflect.TypeOf((KrosmasterInventoryErrorMessage)(KrosmasterInventoryErrorMessage{})),

	6344: reflect.TypeOf((KrosmasterInventoryRequestMessage)(KrosmasterInventoryRequestMessage{})),

	6345: reflect.TypeOf((KrosmasterAuthTokenErrorMessage)(KrosmasterAuthTokenErrorMessage{})),

	6346: reflect.TypeOf((KrosmasterAuthTokenRequestMessage)(KrosmasterAuthTokenRequestMessage{})),

	6347: reflect.TypeOf((KrosmasterPlayingStatusMessage)(KrosmasterPlayingStatusMessage{})),

	6348: reflect.TypeOf((KrosmasterTransferMessage)(KrosmasterTransferMessage{})),

	6349: reflect.TypeOf((KrosmasterTransferRequestMessage)(KrosmasterTransferRequestMessage{})),

	6350: reflect.TypeOf((KrosmasterInventoryMessage)(KrosmasterInventoryMessage{})),

	6351: reflect.TypeOf((KrosmasterAuthTokenMessage)(KrosmasterAuthTokenMessage{})),

	6355: reflect.TypeOf((SpouseGetInformationsMessage)(SpouseGetInformationsMessage{})),

	6356: reflect.TypeOf((SpouseInformationsMessage)(SpouseInformationsMessage{})),

	6357: reflect.TypeOf((AchievementDetailedListRequestMessage)(AchievementDetailedListRequestMessage{})),

	6358: reflect.TypeOf((AchievementDetailedListMessage)(AchievementDetailedListMessage{})),

	6362: reflect.TypeOf((BasicAckMessage)(BasicAckMessage{})),

	6363: reflect.TypeOf((TitlesAndOrnamentsListRequestMessage)(TitlesAndOrnamentsListRequestMessage{})),

	6364: reflect.TypeOf((TitleGainedMessage)(TitleGainedMessage{})),

	6365: reflect.TypeOf((TitleSelectRequestMessage)(TitleSelectRequestMessage{})),

	6366: reflect.TypeOf((TitleSelectedMessage)(TitleSelectedMessage{})),

	6367: reflect.TypeOf((TitlesAndOrnamentsListMessage)(TitlesAndOrnamentsListMessage{})),

	6368: reflect.TypeOf((OrnamentGainedMessage)(OrnamentGainedMessage{})),

	6369: reflect.TypeOf((OrnamentSelectedMessage)(OrnamentSelectedMessage{})),

	6370: reflect.TypeOf((OrnamentSelectErrorMessage)(OrnamentSelectErrorMessage{})),

	6371: reflect.TypeOf((TitleLostMessage)(TitleLostMessage{})),

	6372: reflect.TypeOf((CheckIntegrityMessage)(CheckIntegrityMessage{})),

	6373: reflect.TypeOf((TitleSelectErrorMessage)(TitleSelectErrorMessage{})),

	6374: reflect.TypeOf((OrnamentSelectRequestMessage)(OrnamentSelectRequestMessage{})),

	6375: reflect.TypeOf((AchievementRewardErrorMessage)(AchievementRewardErrorMessage{})),

	6376: reflect.TypeOf((AchievementRewardSuccessMessage)(AchievementRewardSuccessMessage{})),

	6377: reflect.TypeOf((AchievementRewardRequestMessage)(AchievementRewardRequestMessage{})),

	6378: reflect.TypeOf((AchievementDetailsMessage)(AchievementDetailsMessage{})),

	6380: reflect.TypeOf((AchievementDetailsRequestMessage)(AchievementDetailsRequestMessage{})),

	6381: reflect.TypeOf((AchievementFinishedInformationMessage)(AchievementFinishedInformationMessage{})),

	6382: reflect.TypeOf((FriendGuildSetWarnOnAchievementCompleteMessage)(FriendGuildSetWarnOnAchievementCompleteMessage{})),

	6383: reflect.TypeOf((FriendGuildWarnOnAchievementCompleteStateMessage)(FriendGuildWarnOnAchievementCompleteStateMessage{})),

	6384: reflect.TypeOf((InteractiveUseErrorMessage)(InteractiveUseErrorMessage{})),

	6385: reflect.TypeOf((PlayerStatusUpdateErrorMessage)(PlayerStatusUpdateErrorMessage{})),

	6386: reflect.TypeOf((PlayerStatusUpdateMessage)(PlayerStatusUpdateMessage{})),

	6387: reflect.TypeOf((PlayerStatusUpdateRequestMessage)(PlayerStatusUpdateRequestMessage{})),

	6388: reflect.TypeOf((MoodSmileyUpdateMessage)(MoodSmileyUpdateMessage{})),

	6389: reflect.TypeOf((ExchangeSetCraftRecipeMessage)(ExchangeSetCraftRecipeMessage{})),

	6390: reflect.TypeOf((AllianceMembershipMessage)(AllianceMembershipMessage{})),

	6391: reflect.TypeOf((AllianceCreationResultMessage)(AllianceCreationResultMessage{})),

	6392: reflect.TypeOf((AllianceInvitationStateRecrutedMessage)(AllianceInvitationStateRecrutedMessage{})),

	6393: reflect.TypeOf((AllianceCreationValidMessage)(AllianceCreationValidMessage{})),

	6394: reflect.TypeOf((AllianceCreationStartedMessage)(AllianceCreationStartedMessage{})),

	6395: reflect.TypeOf((AllianceInvitationMessage)(AllianceInvitationMessage{})),

	6396: reflect.TypeOf((AllianceInvitationStateRecruterMessage)(AllianceInvitationStateRecruterMessage{})),

	6397: reflect.TypeOf((AllianceInvitedMessage)(AllianceInvitedMessage{})),

	6398: reflect.TypeOf((AllianceLeftMessage)(AllianceLeftMessage{})),

	6399: reflect.TypeOf((AllianceGuildLeavingMessage)(AllianceGuildLeavingMessage{})),

	6400: reflect.TypeOf((AllianceKickRequestMessage)(AllianceKickRequestMessage{})),

	6401: reflect.TypeOf((AllianceInvitationAnswerMessage)(AllianceInvitationAnswerMessage{})),

	6402: reflect.TypeOf((AllianceJoinedMessage)(AllianceJoinedMessage{})),

	6403: reflect.TypeOf((AllianceInsiderInfoMessage)(AllianceInsiderInfoMessage{})),

	6404: reflect.TypeOf((GuildFactsRequestMessage)(GuildFactsRequestMessage{})),

	6407: reflect.TypeOf((GameRolePlayShowActorWithEventMessage)(GameRolePlayShowActorWithEventMessage{})),

	6408: reflect.TypeOf((AllianceListMessage)(AllianceListMessage{})),

	6409: reflect.TypeOf((AllianceFactsRequestMessage)(AllianceFactsRequestMessage{})),

	6412: reflect.TypeOf((GameContextRemoveElementWithEventMessage)(GameContextRemoveElementWithEventMessage{})),

	6413: reflect.TypeOf((GuildListMessage)(GuildListMessage{})),

	6414: reflect.TypeOf((AllianceFactsMessage)(AllianceFactsMessage{})),

	6415: reflect.TypeOf((GuildFactsMessage)(GuildFactsMessage{})),

	6416: reflect.TypeOf((GameContextRemoveMultipleElementsWithEventsMessage)(GameContextRemoveMultipleElementsWithEventsMessage{})),

	6417: reflect.TypeOf((AllianceInsiderInfoRequestMessage)(AllianceInsiderInfoRequestMessage{})),

	6422: reflect.TypeOf((GuildInAllianceFactsMessage)(GuildInAllianceFactsMessage{})),

	6423: reflect.TypeOf((AllianceFactsErrorMessage)(AllianceFactsErrorMessage{})),

	6424: reflect.TypeOf((GuildFactsErrorMessage)(GuildFactsErrorMessage{})),

	6425: reflect.TypeOf((GameRolePlayDelayedObjectUseMessage)(GameRolePlayDelayedObjectUseMessage{})),

	6426: reflect.TypeOf((AllianceChangeGuildRightsMessage)(AllianceChangeGuildRightsMessage{})),

	6427: reflect.TypeOf((AlliancePartialListMessage)(AlliancePartialListMessage{})),

	6434: reflect.TypeOf((ServerSessionConstantsMessage)(ServerSessionConstantsMessage{})),

	6435: reflect.TypeOf((GuildVersatileInfoListMessage)(GuildVersatileInfoListMessage{})),

	6436: reflect.TypeOf((AllianceVersatileInfoListMessage)(AllianceVersatileInfoListMessage{})),

	6437: reflect.TypeOf((PrismSettingsRequestMessage)(PrismSettingsRequestMessage{})),

	6438: reflect.TypeOf((PrismsListUpdateMessage)(PrismsListUpdateMessage{})),

	6439: reflect.TypeOf((KohUpdateMessage)(KohUpdateMessage{})),

	6440: reflect.TypeOf((PrismsListMessage)(PrismsListMessage{})),

	6441: reflect.TypeOf((PrismsListRegisterMessage)(PrismsListRegisterMessage{})),

	6442: reflect.TypeOf((PrismSettingsErrorMessage)(PrismSettingsErrorMessage{})),

	6443: reflect.TypeOf((SetEnableAVARequestMessage)(SetEnableAVARequestMessage{})),

	6444: reflect.TypeOf((AllianceModificationStartedMessage)(AllianceModificationStartedMessage{})),

	6445: reflect.TypeOf((AllianceTaxCollectorDialogQuestionExtendedMessage)(AllianceTaxCollectorDialogQuestionExtendedMessage{})),

	6447: reflect.TypeOf((AllianceModificationEmblemValidMessage)(AllianceModificationEmblemValidMessage{})),

	6448: reflect.TypeOf((AlliancePrismDialogQuestionMessage)(AlliancePrismDialogQuestionMessage{})),

	6449: reflect.TypeOf((AllianceModificationNameAndTagValidMessage)(AllianceModificationNameAndTagValidMessage{})),

	6450: reflect.TypeOf((AllianceModificationValidMessage)(AllianceModificationValidMessage{})),

	6451: reflect.TypeOf((PrismsInfoValidMessage)(PrismsInfoValidMessage{})),

	6452: reflect.TypeOf((PrismFightAddedMessage)(PrismFightAddedMessage{})),

	6453: reflect.TypeOf((PrismFightRemovedMessage)(PrismFightRemovedMessage{})),

	6454: reflect.TypeOf((UpdateMapPlayersAgressableStatusMessage)(UpdateMapPlayersAgressableStatusMessage{})),

	6455: reflect.TypeOf((TaxCollectorStateUpdateMessage)(TaxCollectorStateUpdateMessage{})),

	6456: reflect.TypeOf((UpdateSelfAgressableStatusMessage)(UpdateSelfAgressableStatusMessage{})),

	6457: reflect.TypeOf((MimicryObjectEraseRequestMessage)(MimicryObjectEraseRequestMessage{})),

	6458: reflect.TypeOf((MimicryObjectPreviewMessage)(MimicryObjectPreviewMessage{})),

	6459: reflect.TypeOf((ClientUIOpenedMessage)(ClientUIOpenedMessage{})),

	6460: reflect.TypeOf((MimicryObjectFeedAndAssociateRequestMessage)(MimicryObjectFeedAndAssociateRequestMessage{})),

	6461: reflect.TypeOf((MimicryObjectErrorMessage)(MimicryObjectErrorMessage{})),

	6462: reflect.TypeOf((MimicryObjectAssociatedMessage)(MimicryObjectAssociatedMessage{})),

	6463: reflect.TypeOf((ClientUIOpenedByObjectMessage)(ClientUIOpenedByObjectMessage{})),

	6464: reflect.TypeOf((ExchangeBidPriceForSellerMessage)(ExchangeBidPriceForSellerMessage{})),

	6465: reflect.TypeOf((GameFightTurnStartPlayingMessage)(GameFightTurnStartPlayingMessage{})),

	6466: reflect.TypeOf((PrismSetSabotagedRefusedMessage)(PrismSetSabotagedRefusedMessage{})),

	6468: reflect.TypeOf((PrismSetSabotagedRequestMessage)(PrismSetSabotagedRequestMessage{})),

	6469: reflect.TypeOf((SelectedServerDataExtendedMessage)(SelectedServerDataExtendedMessage{})),

	6470: reflect.TypeOf((ExchangeObjectTransfertListWithQuantityToInvMessage)(ExchangeObjectTransfertListWithQuantityToInvMessage{})),

	6471: reflect.TypeOf((CharacterLoadingCompleteMessage)(CharacterLoadingCompleteMessage{})),

	6472: reflect.TypeOf((PartyCompanionUpdateLightMessage)(PartyCompanionUpdateLightMessage{})),

	6474: reflect.TypeOf((GameFightSpectatePlayerRequestMessage)(GameFightSpectatePlayerRequestMessage{})),

	6475: reflect.TypeOf((BasicCharactersListMessage)(BasicCharactersListMessage{})),

	6483: reflect.TypeOf((TreasureHuntFinishedMessage)(TreasureHuntFinishedMessage{})),

	6484: reflect.TypeOf((TreasureHuntDigRequestAnswerMessage)(TreasureHuntDigRequestAnswerMessage{})),

	6485: reflect.TypeOf((TreasureHuntDigRequestMessage)(TreasureHuntDigRequestMessage{})),

	6486: reflect.TypeOf((TreasureHuntMessage)(TreasureHuntMessage{})),

	6487: reflect.TypeOf((TreasureHuntGiveUpRequestMessage)(TreasureHuntGiveUpRequestMessage{})),

	6488: reflect.TypeOf((TreasureHuntRequestMessage)(TreasureHuntRequestMessage{})),

	6489: reflect.TypeOf((TreasureHuntRequestAnswerMessage)(TreasureHuntRequestAnswerMessage{})),

	6490: reflect.TypeOf((GameFightNewWaveMessage)(GameFightNewWaveMessage{})),

	6491: reflect.TypeOf((TreasureHuntAvailableRetryCountUpdateMessage)(TreasureHuntAvailableRetryCountUpdateMessage{})),

	6492: reflect.TypeOf((PortalUseRequestMessage)(PortalUseRequestMessage{})),

	6493: reflect.TypeOf((AreaFightModificatorUpdateMessage)(AreaFightModificatorUpdateMessage{})),

	6496: reflect.TypeOf((GameCautiousMapMovementRequestMessage)(GameCautiousMapMovementRequestMessage{})),

	6497: reflect.TypeOf((GameCautiousMapMovementMessage)(GameCautiousMapMovementMessage{})),

	6498: reflect.TypeOf((TreasureHuntShowLegendaryUIMessage)(TreasureHuntShowLegendaryUIMessage{})),

	6499: reflect.TypeOf((TreasureHuntLegendaryRequestMessage)(TreasureHuntLegendaryRequestMessage{})),

	6500: reflect.TypeOf((MapRunningFightDetailsExtendedMessage)(MapRunningFightDetailsExtendedMessage{})),

	6501: reflect.TypeOf((PartyNameSetErrorMessage)(PartyNameSetErrorMessage{})),

	6502: reflect.TypeOf((PartyNameUpdateMessage)(PartyNameUpdateMessage{})),

	6503: reflect.TypeOf((PartyNameSetRequestMessage)(PartyNameSetRequestMessage{})),

	6504: reflect.TypeOf((GameFightSpectatorJoinMessage)(GameFightSpectatorJoinMessage{})),

	6505: reflect.TypeOf((GuestModeMessage)(GuestModeMessage{})),

	6506: reflect.TypeOf((GuestLimitationMessage)(GuestLimitationMessage{})),

	6507: reflect.TypeOf((TreasureHuntFlagRequestAnswerMessage)(TreasureHuntFlagRequestAnswerMessage{})),

	6508: reflect.TypeOf((TreasureHuntFlagRequestMessage)(TreasureHuntFlagRequestMessage{})),

	6509: reflect.TypeOf((TreasureHuntDigRequestAnswerFailedMessage)(TreasureHuntDigRequestAnswerFailedMessage{})),

	6510: reflect.TypeOf((TreasureHuntFlagRemoveRequestMessage)(TreasureHuntFlagRemoveRequestMessage{})),

	6512: reflect.TypeOf((WarnOnPermaDeathMessage)(WarnOnPermaDeathMessage{})),

	6513: reflect.TypeOf((WarnOnPermaDeathStateMessage)(WarnOnPermaDeathStateMessage{})),

	6516: reflect.TypeOf((PaddockBuyResultMessage)(PaddockBuyResultMessage{})),

	6517: reflect.TypeOf((AccessoryPreviewMessage)(AccessoryPreviewMessage{})),

	6518: reflect.TypeOf((AccessoryPreviewRequestMessage)(AccessoryPreviewRequestMessage{})),

	6519: reflect.TypeOf((ObtainedItemMessage)(ObtainedItemMessage{})),

	6520: reflect.TypeOf((ObtainedItemWithBonusMessage)(ObtainedItemWithBonusMessage{})),

	6521: reflect.TypeOf((AccessoryPreviewErrorMessage)(AccessoryPreviewErrorMessage{})),

	6522: reflect.TypeOf((SymbioticObjectAssociateRequestMessage)(SymbioticObjectAssociateRequestMessage{})),

	6523: reflect.TypeOf((WrapperObjectAssociatedMessage)(WrapperObjectAssociatedMessage{})),

	6524: reflect.TypeOf((WrapperObjectDissociateRequestMessage)(WrapperObjectDissociateRequestMessage{})),

	6525: reflect.TypeOf((CurrentServerStatusUpdateMessage)(CurrentServerStatusUpdateMessage{})),

	6526: reflect.TypeOf((SymbioticObjectErrorMessage)(SymbioticObjectErrorMessage{})),

	6527: reflect.TypeOf((SymbioticObjectAssociatedMessage)(SymbioticObjectAssociatedMessage{})),

	6529: reflect.TypeOf((WrapperObjectErrorMessage)(WrapperObjectErrorMessage{})),

	6530: reflect.TypeOf((BasicStatMessage)(BasicStatMessage{})),

	6531: reflect.TypeOf((PrismModuleExchangeRequestMessage)(PrismModuleExchangeRequestMessage{})),

	6532: reflect.TypeOf((ExchangeObjectsRemovedMessage)(ExchangeObjectsRemovedMessage{})),

	6533: reflect.TypeOf((ExchangeObjectsModifiedMessage)(ExchangeObjectsModifiedMessage{})),

	6535: reflect.TypeOf((ExchangeObjectsAddedMessage)(ExchangeObjectsAddedMessage{})),

	6536: reflect.TypeOf((ComicReadingBeginMessage)(ComicReadingBeginMessage{})),

	6537: reflect.TypeOf((StartupActionsAllAttributionMessage)(StartupActionsAllAttributionMessage{})),

	6538: reflect.TypeOf((StartupActionAddMessage)(StartupActionAddMessage{})),

	6539: reflect.TypeOf((ReloginTokenStatusMessage)(ReloginTokenStatusMessage{})),

	6540: reflect.TypeOf((ReloginTokenRequestMessage)(ReloginTokenRequestMessage{})),

	6541: reflect.TypeOf((GameFightPlacementSwapPositionsRequestMessage)(GameFightPlacementSwapPositionsRequestMessage{})),

	6542: reflect.TypeOf((GameFightPlacementSwapPositionsOfferMessage)(GameFightPlacementSwapPositionsOfferMessage{})),

	6543: reflect.TypeOf((GameFightPlacementSwapPositionsCancelMessage)(GameFightPlacementSwapPositionsCancelMessage{})),

	6544: reflect.TypeOf((GameFightPlacementSwapPositionsMessage)(GameFightPlacementSwapPositionsMessage{})),

	6545: reflect.TypeOf((GameActionFightActivateGlyphTrapMessage)(GameActionFightActivateGlyphTrapMessage{})),

	6546: reflect.TypeOf((GameFightPlacementSwapPositionsCancelledMessage)(GameFightPlacementSwapPositionsCancelledMessage{})),

	6547: reflect.TypeOf((GameFightPlacementSwapPositionsAcceptMessage)(GameFightPlacementSwapPositionsAcceptMessage{})),

	6548: reflect.TypeOf((GameFightPlacementSwapPositionsErrorMessage)(GameFightPlacementSwapPositionsErrorMessage{})),

	6549: reflect.TypeOf((CharacterSelectionWithRemodelMessage)(CharacterSelectionWithRemodelMessage{})),

	6550: reflect.TypeOf((CharactersListWithRemodelingMessage)(CharactersListWithRemodelingMessage{})),

	6551: reflect.TypeOf((CharacterReplayWithRemodelRequestMessage)(CharacterReplayWithRemodelRequestMessage{})),

	6554: reflect.TypeOf((ExchangeMountsTakenFromPaddockMessage)(ExchangeMountsTakenFromPaddockMessage{})),

	6555: reflect.TypeOf((ExchangeMountsStableAddMessage)(ExchangeMountsStableAddMessage{})),

	6556: reflect.TypeOf((ExchangeMountsStableRemoveMessage)(ExchangeMountsStableRemoveMessage{})),

	6557: reflect.TypeOf((ExchangeMountsStableBornAddMessage)(ExchangeMountsStableBornAddMessage{})),

	6559: reflect.TypeOf((ExchangeMountsPaddockRemoveMessage)(ExchangeMountsPaddockRemoveMessage{})),

	6561: reflect.TypeOf((ExchangeMountsPaddockAddMessage)(ExchangeMountsPaddockAddMessage{})),

	6562: reflect.TypeOf((ExchangeHandleMountsStableMessage)(ExchangeHandleMountsStableMessage{})),

	6563: reflect.TypeOf((DisplayNumericalValuePaddockMessage)(DisplayNumericalValuePaddockMessage{})),

	6565: reflect.TypeOf((TopTaxCollectorListMessage)(TopTaxCollectorListMessage{})),

	6567: reflect.TypeOf((ExchangeStartOkRunesTradeMessage)(ExchangeStartOkRunesTradeMessage{})),

	6568: reflect.TypeOf((AbstractTaxCollectorListMessage)(AbstractTaxCollectorListMessage{})),

	6569: reflect.TypeOf((DecraftResultMessage)(DecraftResultMessage{})),

	6571: reflect.TypeOf((ZaapRespawnUpdatedMessage)(ZaapRespawnUpdatedMessage{})),

	6572: reflect.TypeOf((ZaapRespawnSaveRequestMessage)(ZaapRespawnSaveRequestMessage{})),

	6573: reflect.TypeOf((BasicStatWithDataMessage)(BasicStatWithDataMessage{})),

	6574: reflect.TypeOf((GameRolePlayArenaSwitchToGameServerMessage)(GameRolePlayArenaSwitchToGameServerMessage{})),

	6575: reflect.TypeOf((GameRolePlayArenaSwitchToFightServerMessage)(GameRolePlayArenaSwitchToFightServerMessage{})),

	6578: reflect.TypeOf((ExchangeCraftPaymentModifiedMessage)(ExchangeCraftPaymentModifiedMessage{})),

	6579: reflect.TypeOf((ExchangeCraftPaymentModificationRequestMessage)(ExchangeCraftPaymentModificationRequestMessage{})),

	6580: reflect.TypeOf((IdolPartyLostMessage)(IdolPartyLostMessage{})),

	6581: reflect.TypeOf((IdolSelectedMessage)(IdolSelectedMessage{})),

	6582: reflect.TypeOf((IdolPartyRegisterRequestMessage)(IdolPartyRegisterRequestMessage{})),

	6583: reflect.TypeOf((IdolPartyRefreshMessage)(IdolPartyRefreshMessage{})),

	6584: reflect.TypeOf((IdolSelectErrorMessage)(IdolSelectErrorMessage{})),

	6585: reflect.TypeOf((IdolListMessage)(IdolListMessage{})),

	6586: reflect.TypeOf((IdolFightPreparationUpdateMessage)(IdolFightPreparationUpdateMessage{})),

	6587: reflect.TypeOf((IdolSelectRequestMessage)(IdolSelectRequestMessage{})),

	6588: reflect.TypeOf((GuildMotdSetRequestMessage)(GuildMotdSetRequestMessage{})),

	6589: reflect.TypeOf((ExchangeStoppedMessage)(ExchangeStoppedMessage{})),

	6590: reflect.TypeOf((GuildMotdMessage)(GuildMotdMessage{})),

	6591: reflect.TypeOf((GuildMotdSetErrorMessage)(GuildMotdSetErrorMessage{})),

	6592: reflect.TypeOf((JobBookSubscribeRequestMessage)(JobBookSubscribeRequestMessage{})),

	6593: reflect.TypeOf((JobBookSubscriptionMessage)(JobBookSubscriptionMessage{})),

	6594: reflect.TypeOf((ClientYouAreDrunkMessage)(ClientYouAreDrunkMessage{})),

	6595: reflect.TypeOf((ExchangeCraftCountModifiedMessage)(ExchangeCraftCountModifiedMessage{})),

	6596: reflect.TypeOf((ChatSmileyExtraPackListMessage)(ChatSmileyExtraPackListMessage{})),

	6597: reflect.TypeOf((ExchangeCraftCountRequestMessage)(ExchangeCraftCountRequestMessage{})),

	6598: reflect.TypeOf((ExchangeCrafterJobLevelupMessage)(ExchangeCrafterJobLevelupMessage{})),

	6599: reflect.TypeOf((JobExperienceOtherPlayerUpdateMessage)(JobExperienceOtherPlayerUpdateMessage{})),

	6600: reflect.TypeOf((ExchangeStartOkRecycleTradeMessage)(ExchangeStartOkRecycleTradeMessage{})),

	6601: reflect.TypeOf((RecycleResultMessage)(RecycleResultMessage{})),

	6602: reflect.TypeOf((IdolsPresetDeleteMessage)(IdolsPresetDeleteMessage{})),

	6603: reflect.TypeOf((IdolsPresetSaveMessage)(IdolsPresetSaveMessage{})),

	6604: reflect.TypeOf((IdolsPresetSaveResultMessage)(IdolsPresetSaveResultMessage{})),

	6605: reflect.TypeOf((IdolsPresetDeleteResultMessage)(IdolsPresetDeleteResultMessage{})),

	6606: reflect.TypeOf((IdolsPresetUpdateMessage)(IdolsPresetUpdateMessage{})),

	6607: reflect.TypeOf((AccountLinkRequiredMessage)(AccountLinkRequiredMessage{})),

	6611: reflect.TypeOf((TaxCollectorMovementsOfflineMessage)(TaxCollectorMovementsOfflineMessage{})),

	6612: reflect.TypeOf((ExchangeBidHouseUnsoldItemsMessage)(ExchangeBidHouseUnsoldItemsMessage{})),

	6613: reflect.TypeOf((ExchangeOfflineSoldItemsMessage)(ExchangeOfflineSoldItemsMessage{})),

	6614: reflect.TypeOf((IdolsPresetUseResultMessage)(IdolsPresetUseResultMessage{})),

	6615: reflect.TypeOf((IdolsPresetUseMessage)(IdolsPresetUseMessage{})),

	6616: reflect.TypeOf((SubscriptionUpdateMessage)(SubscriptionUpdateMessage{})),

	6618: reflect.TypeOf((GameRefreshMonsterBoostsMessage)(GameRefreshMonsterBoostsMessage{})),

	6619: reflect.TypeOf((EditHavenBagCancelRequestMessage)(EditHavenBagCancelRequestMessage{})),

	6620: reflect.TypeOf((HavenBagPackListMessage)(HavenBagPackListMessage{})),

	6621: reflect.TypeOf((CloseHavenBagFurnitureSequenceRequestMessage)(CloseHavenBagFurnitureSequenceRequestMessage{})),

	6622: reflect.TypeOf((MapComplementaryInformationsDataInHavenBagMessage)(MapComplementaryInformationsDataInHavenBagMessage{})),

	6626: reflect.TypeOf((EditHavenBagRequestMessage)(EditHavenBagRequestMessage{})),

	6628: reflect.TypeOf((EditHavenBagFinishedMessage)(EditHavenBagFinishedMessage{})),

	6630: reflect.TypeOf((RoomAvailableUpdateMessage)(RoomAvailableUpdateMessage{})),

	6631: reflect.TypeOf((ExitHavenBagRequestMessage)(ExitHavenBagRequestMessage{})),

	6632: reflect.TypeOf((EditHavenBagStartMessage)(EditHavenBagStartMessage{})),

	6634: reflect.TypeOf((HavenBagFurnituresMessage)(HavenBagFurnituresMessage{})),

	6635: reflect.TypeOf((OpenHavenBagFurnitureSequenceRequestMessage)(OpenHavenBagFurnitureSequenceRequestMessage{})),

	6636: reflect.TypeOf((EnterHavenBagRequestMessage)(EnterHavenBagRequestMessage{})),

	6637: reflect.TypeOf((HavenBagFurnituresRequestMessage)(HavenBagFurnituresRequestMessage{})),

	6638: reflect.TypeOf((ChangeHavenBagRoomRequestMessage)(ChangeHavenBagRoomRequestMessage{})),

	6639: reflect.TypeOf((ChangeThemeRequestMessage)(ChangeThemeRequestMessage{})),

	6642: reflect.TypeOf((InviteInHavenBagMessage)(InviteInHavenBagMessage{})),

	6643: reflect.TypeOf((InviteInHavenBagOfferMessage)(InviteInHavenBagOfferMessage{})),

	6644: reflect.TypeOf((HavenBagDailyLoteryMessage)(HavenBagDailyLoteryMessage{})),

	6645: reflect.TypeOf((InviteInHavenBagClosedMessage)(InviteInHavenBagClosedMessage{})),

	6646: reflect.TypeOf((TeleportHavenBagAnswerMessage)(TeleportHavenBagAnswerMessage{})),

	6647: reflect.TypeOf((TeleportHavenBagRequestMessage)(TeleportHavenBagRequestMessage{})),

	6648: reflect.TypeOf((HaapiApiKeyRequestMessage)(HaapiApiKeyRequestMessage{})),

	6649: reflect.TypeOf((HaapiApiKeyMessage)(HaapiApiKeyMessage{})),

	6652: reflect.TypeOf((KickHavenBagRequestMessage)(KickHavenBagRequestMessage{})),

	6653: reflect.TypeOf((SpellModifyFailureMessage)(SpellModifyFailureMessage{})),

	6654: reflect.TypeOf((SpellModifySuccessMessage)(SpellModifySuccessMessage{})),

	6655: reflect.TypeOf((SpellModifyRequestMessage)(SpellModifyRequestMessage{})),

	6656: reflect.TypeOf((DareInformationsMessage)(DareInformationsMessage{})),

	6657: reflect.TypeOf((DareVersatileListMessage)(DareVersatileListMessage{})),

	6658: reflect.TypeOf((DareSubscribedListMessage)(DareSubscribedListMessage{})),

	6659: reflect.TypeOf((DareInformationsRequestMessage)(DareInformationsRequestMessage{})),

	6660: reflect.TypeOf((DareSubscribedMessage)(DareSubscribedMessage{})),

	6661: reflect.TypeOf((DareListMessage)(DareListMessage{})),

	6662: reflect.TypeOf((AggregateStatWithDataMessage)(AggregateStatWithDataMessage{})),

	6663: reflect.TypeOf((DareCreatedListMessage)(DareCreatedListMessage{})),

	6664: reflect.TypeOf((ExchangeStartedTaxCollectorShopMessage)(ExchangeStartedTaxCollectorShopMessage{})),

	6665: reflect.TypeOf((DareCreationRequestMessage)(DareCreationRequestMessage{})),

	6666: reflect.TypeOf((DareSubscribeRequestMessage)(DareSubscribeRequestMessage{})),

	6667: reflect.TypeOf((DareErrorMessage)(DareErrorMessage{})),

	6668: reflect.TypeOf((DareCreatedMessage)(DareCreatedMessage{})),

	6669: reflect.TypeOf((AggregateStatMessage)(AggregateStatMessage{})),

	6670: reflect.TypeOf((ExchangePodsModifiedMessage)(ExchangePodsModifiedMessage{})),

	6675: reflect.TypeOf((DareRewardConsumeValidationMessage)(DareRewardConsumeValidationMessage{})),

	6676: reflect.TypeOf((DareRewardConsumeRequestMessage)(DareRewardConsumeRequestMessage{})),

	6677: reflect.TypeOf((DareRewardsListMessage)(DareRewardsListMessage{})),

	6678: reflect.TypeOf((DareRewardWonMessage)(DareRewardWonMessage{})),

	6679: reflect.TypeOf((DareCanceledMessage)(DareCanceledMessage{})),

	6680: reflect.TypeOf((DareCancelRequestMessage)(DareCancelRequestMessage{})),

	6681: reflect.TypeOf((DareWonMessage)(DareWonMessage{})),

	6682: reflect.TypeOf((DareWonListMessage)(DareWonListMessage{})),

	6683: reflect.TypeOf((AllianceMotdSetErrorMessage)(AllianceMotdSetErrorMessage{})),

	6684: reflect.TypeOf((SocialNoticeSetErrorMessage)(SocialNoticeSetErrorMessage{})),

	6685: reflect.TypeOf((AllianceMotdMessage)(AllianceMotdMessage{})),

	6686: reflect.TypeOf((SocialNoticeSetRequestMessage)(SocialNoticeSetRequestMessage{})),

	6687: reflect.TypeOf((AllianceMotdSetRequestMessage)(AllianceMotdSetRequestMessage{})),

	6688: reflect.TypeOf((SocialNoticeMessage)(SocialNoticeMessage{})),

	6689: reflect.TypeOf((GuildBulletinMessage)(GuildBulletinMessage{})),

	6690: reflect.TypeOf((AllianceBulletinMessage)(AllianceBulletinMessage{})),

	6691: reflect.TypeOf((GuildBulletinSetErrorMessage)(GuildBulletinSetErrorMessage{})),

	6692: reflect.TypeOf((AllianceBulletinSetErrorMessage)(AllianceBulletinSetErrorMessage{})),

	6693: reflect.TypeOf((AllianceBulletinSetRequestMessage)(AllianceBulletinSetRequestMessage{})),

	6694: reflect.TypeOf((GuildBulletinSetRequestMessage)(GuildBulletinSetRequestMessage{})),

	6695: reflect.TypeOf((BulletinMessage)(BulletinMessage{})),

	6696: reflect.TypeOf((MountHarnessDissociateRequestMessage)(MountHarnessDissociateRequestMessage{})),

	6697: reflect.TypeOf((MountHarnessColorsUpdateRequestMessage)(MountHarnessColorsUpdateRequestMessage{})),

	6699: reflect.TypeOf((RefreshCharacterStatsMessage)(RefreshCharacterStatsMessage{})),

	6700: reflect.TypeOf((ArenaFighterLeaveMessage)(ArenaFighterLeaveMessage{})),

	6701: reflect.TypeOf((FocusedExchangeReadyMessage)(FocusedExchangeReadyMessage{})),

	6702: reflect.TypeOf((FinishMoveListRequestMessage)(FinishMoveListRequestMessage{})),

	6703: reflect.TypeOf((FinishMoveSetRequestMessage)(FinishMoveSetRequestMessage{})),

	6704: reflect.TypeOf((FinishMoveListMessage)(FinishMoveListMessage{})),

	6705: reflect.TypeOf((SpellVariantActivationMessage)(SpellVariantActivationMessage{})),

	6706: reflect.TypeOf((ShortcutBarReplacedMessage)(ShortcutBarReplacedMessage{})),

	6707: reflect.TypeOf((SpellVariantActivationRequestMessage)(SpellVariantActivationRequestMessage{})),

	6712: reflect.TypeOf((GameRolePlayShowMultipleActorsMessage)(GameRolePlayShowMultipleActorsMessage{})),

	6713: reflect.TypeOf((HavenBagPermissionsUpdateMessage)(HavenBagPermissionsUpdateMessage{})),

	6714: reflect.TypeOf((HavenBagPermissionsUpdateRequestMessage)(HavenBagPermissionsUpdateRequestMessage{})),

	6715: reflect.TypeOf((InteractiveUseWithParamRequestMessage)(InteractiveUseWithParamRequestMessage{})),

	6716: reflect.TypeOf((MapFightStartPositionsUpdateMessage)(MapFightStartPositionsUpdateMessage{})),

	6717: reflect.TypeOf((FollowedQuestsMessage)(FollowedQuestsMessage{})),

	6722: reflect.TypeOf((RefreshFollowedQuestsOrderRequestMessage)(RefreshFollowedQuestsOrderRequestMessage{})),

	6723: reflect.TypeOf((UnfollowQuestObjectiveRequestMessage)(UnfollowQuestObjectiveRequestMessage{})),

	6724: reflect.TypeOf((FollowQuestObjectiveRequestMessage)(FollowQuestObjectiveRequestMessage{})),

	6726: reflect.TypeOf((HouseTeleportRequestMessage)(HouseTeleportRequestMessage{})),

	6727: reflect.TypeOf((HouseSellingUpdateMessage)(HouseSellingUpdateMessage{})),

	6728: reflect.TypeOf((GameRolePlayArenaUpdatePlayerInfosAllQueuesMessage)(GameRolePlayArenaUpdatePlayerInfosAllQueuesMessage{})),

	6729: reflect.TypeOf((ChatCommunityChannelSetCommunityRequestMessage)(ChatCommunityChannelSetCommunityRequestMessage{})),

	6730: reflect.TypeOf((ChatCommunityChannelCommunityMessage)(ChatCommunityChannelCommunityMessage{})),

	6731: reflect.TypeOf((MigratedServerListMessage)(MigratedServerListMessage{})),
}

type ProtocolRequired struct {
	RequiredVersion uint32

	CurrentVersion uint32
}

func (m *ProtocolRequired) ID() uint16 {
	return 1
}

func (m *ProtocolRequired) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.RequiredVersion); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.CurrentVersion); err != nil {
		return err
	}

	return nil
}

func (m *ProtocolRequired) Deserialize(r Reader) error {

	lrequiredVersion, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.RequiredVersion = lrequiredVersion

	lcurrentVersion, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.CurrentVersion = lcurrentVersion

	return nil
}

type NetworkDataContainerMessage struct {
	Content []uint8
}

func (m *NetworkDataContainerMessage) ID() uint16 {
	return 2
}

func (m *NetworkDataContainerMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt32(int32(len(m.Content))); err != nil {
		return err
	}

	for i := range m.Content {

		if err := w.WriteUInt8(m.Content[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *NetworkDataContainerMessage) Deserialize(r Reader) error {

	lcontentLen, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Content = make([]uint8, lcontentLen)

	for i := range m.Content {

		lcontent, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.Content[i] = lcontent

	}

	return nil
}

type HelloConnectMessage struct {
	Salt string

	Key []int8
}

func (m *HelloConnectMessage) ID() uint16 {
	return 3
}

func (m *HelloConnectMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Salt); err != nil {
		return err
	}

	if err := w.WriteVarInt32(int32(len(m.Key))); err != nil {
		return err
	}

	for i := range m.Key {

		if err := w.WriteInt8(m.Key[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *HelloConnectMessage) Deserialize(r Reader) error {

	lsalt, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Salt = lsalt

	lkeyLen, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Key = make([]int8, lkeyLen)

	for i := range m.Key {

		lkey, err := r.ReadInt8()
		if err != nil {
			return err
		}

		m.Key[i] = lkey

	}

	return nil
}

type IdentificationMessage struct {
	Version VersionExtended

	Lang string

	Credentials []int8

	ServerId int16

	Autoconnect bool

	UseCertificate bool

	UseLoginToken bool

	SessionOptionalSalt int64

	FailedAttempts []uint16
}

func (m *IdentificationMessage) ID() uint16 {
	return 4
}

func (m *IdentificationMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Autoconnect)

	setWrappedFlag(bbw0, 1, m.UseCertificate)

	setWrappedFlag(bbw0, 2, m.UseLoginToken)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := m.Version.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Lang); err != nil {
		return err
	}

	if err := w.WriteVarInt32(int32(len(m.Credentials))); err != nil {
		return err
	}

	for i := range m.Credentials {

		if err := w.WriteInt8(m.Credentials[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(m.ServerId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.SessionOptionalSalt); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.FailedAttempts))); err != nil {
		return err
	}

	for i := range m.FailedAttempts {

		if err := w.WriteVarUInt16(m.FailedAttempts[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *IdentificationMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Autoconnect = getWrappedFlag(bbw0, 0)

	m.UseCertificate = getWrappedFlag(bbw0, 1)

	m.UseLoginToken = getWrappedFlag(bbw0, 2)

	var lversion VersionExtended

	lversion.Deserialize(r)

	m.Version = lversion

	llang, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Lang = llang

	lcredentialsLen, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Credentials = make([]int8, lcredentialsLen)

	for i := range m.Credentials {

		lcredentials, err := r.ReadInt8()
		if err != nil {
			return err
		}

		m.Credentials[i] = lcredentials

	}

	lserverId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ServerId = lserverId

	lsessionOptionalSalt, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.SessionOptionalSalt = lsessionOptionalSalt

	lfailedAttemptsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FailedAttempts = make([]uint16, lfailedAttemptsLen)

	for i := range m.FailedAttempts {

		lfailedAttempts, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.FailedAttempts[i] = lfailedAttempts

	}

	return nil
}

type LoginQueueStatusMessage struct {
	Position uint16

	Total uint16
}

func (m *LoginQueueStatusMessage) ID() uint16 {
	return 10
}

func (m *LoginQueueStatusMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Position); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Total); err != nil {
		return err
	}

	return nil
}

func (m *LoginQueueStatusMessage) Deserialize(r Reader) error {

	lposition, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.Position = lposition

	ltotal, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.Total = ltotal

	return nil
}

type IdentificationFailedMessage struct {
	Reason uint8
}

func (m *IdentificationFailedMessage) ID() uint16 {
	return 20
}

func (m *IdentificationFailedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *IdentificationFailedMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type IdentificationFailedForBadVersionMessage struct {
	IdentificationFailedMessage

	RequiredVersion Version
}

func (m *IdentificationFailedForBadVersionMessage) ID() uint16 {
	return 21
}

func (m *IdentificationFailedForBadVersionMessage) Serialize(w Writer) error {

	if err := m.IdentificationFailedMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.RequiredVersion.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *IdentificationFailedForBadVersionMessage) Deserialize(r Reader) error {

	if err := m.IdentificationFailedMessage.Deserialize(r); err != nil {
		return err
	}

	var lrequiredVersion Version

	lrequiredVersion.Deserialize(r)

	m.RequiredVersion = lrequiredVersion

	return nil
}

type IdentificationSuccessMessage struct {
	Login string

	Nickname string

	AccountId uint32

	CommunityId uint8

	HasRights bool

	SecretQuestion string

	AccountCreation float64

	SubscriptionElapsedDuration float64

	SubscriptionEndDate float64

	WasAlreadyConnected bool

	HavenbagAvailableRoom uint8
}

func (m *IdentificationSuccessMessage) ID() uint16 {
	return 22
}

func (m *IdentificationSuccessMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.HasRights)

	setWrappedFlag(bbw0, 1, m.WasAlreadyConnected)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteString(m.Login); err != nil {
		return err
	}

	if err := w.WriteString(m.Nickname); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.CommunityId); err != nil {
		return err
	}

	if err := w.WriteString(m.SecretQuestion); err != nil {
		return err
	}

	if err := w.WriteDouble(m.AccountCreation); err != nil {
		return err
	}

	if err := w.WriteDouble(m.SubscriptionElapsedDuration); err != nil {
		return err
	}

	if err := w.WriteDouble(m.SubscriptionEndDate); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.HavenbagAvailableRoom); err != nil {
		return err
	}

	return nil
}

func (m *IdentificationSuccessMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.HasRights = getWrappedFlag(bbw0, 0)

	m.WasAlreadyConnected = getWrappedFlag(bbw0, 1)

	llogin, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Login = llogin

	lnickname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Nickname = lnickname

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	lcommunityId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CommunityId = lcommunityId

	lsecretQuestion, err := r.ReadString()
	if err != nil {
		return err
	}

	m.SecretQuestion = lsecretQuestion

	laccountCreation, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.AccountCreation = laccountCreation

	lsubscriptionElapsedDuration, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.SubscriptionElapsedDuration = lsubscriptionElapsedDuration

	lsubscriptionEndDate, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.SubscriptionEndDate = lsubscriptionEndDate

	lhavenbagAvailableRoom, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.HavenbagAvailableRoom = lhavenbagAvailableRoom

	return nil
}

type ServersListMessage struct {
	Servers []GameServerInformations

	AlreadyConnectedToServerId uint16

	CanCreateNewCharacter bool
}

func (m *ServersListMessage) ID() uint16 {
	return 30
}

func (m *ServersListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Servers))); err != nil {
		return err
	}

	for i := range m.Servers {

		if err := m.Servers[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt16(m.AlreadyConnectedToServerId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.CanCreateNewCharacter); err != nil {
		return err
	}

	return nil
}

func (m *ServersListMessage) Deserialize(r Reader) error {

	lserversLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Servers = make([]GameServerInformations, lserversLen)

	for i := range m.Servers {

		var lservers GameServerInformations

		lservers.Deserialize(r)

		m.Servers[i] = lservers

	}

	lalreadyConnectedToServerId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.AlreadyConnectedToServerId = lalreadyConnectedToServerId

	lcanCreateNewCharacter, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.CanCreateNewCharacter = lcanCreateNewCharacter

	return nil
}

type ServerSelectionMessage struct {
	ServerId uint16
}

func (m *ServerSelectionMessage) ID() uint16 {
	return 40
}

func (m *ServerSelectionMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ServerId); err != nil {
		return err
	}

	return nil
}

func (m *ServerSelectionMessage) Deserialize(r Reader) error {

	lserverId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ServerId = lserverId

	return nil
}

type SelectedServerRefusedMessage struct {
	ServerId uint16

	Error uint8

	ServerStatus uint8
}

func (m *SelectedServerRefusedMessage) ID() uint16 {
	return 41
}

func (m *SelectedServerRefusedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ServerId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Error); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.ServerStatus); err != nil {
		return err
	}

	return nil
}

func (m *SelectedServerRefusedMessage) Deserialize(r Reader) error {

	lserverId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ServerId = lserverId

	lerror, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Error = lerror

	lserverStatus, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ServerStatus = lserverStatus

	return nil
}

type SelectedServerDataMessage struct {
	ServerId uint16

	Address string

	Port uint16

	CanCreateNewCharacter bool

	Ticket []int8
}

func (m *SelectedServerDataMessage) ID() uint16 {
	return 42
}

func (m *SelectedServerDataMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ServerId); err != nil {
		return err
	}

	if err := w.WriteString(m.Address); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Port); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.CanCreateNewCharacter); err != nil {
		return err
	}

	if err := w.WriteVarInt32(int32(len(m.Ticket))); err != nil {
		return err
	}

	for i := range m.Ticket {

		if err := w.WriteInt8(m.Ticket[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *SelectedServerDataMessage) Deserialize(r Reader) error {

	lserverId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ServerId = lserverId

	laddress, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Address = laddress

	lport, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.Port = lport

	lcanCreateNewCharacter, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.CanCreateNewCharacter = lcanCreateNewCharacter

	lticketLen, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Ticket = make([]int8, lticketLen)

	for i := range m.Ticket {

		lticket, err := r.ReadInt8()
		if err != nil {
			return err
		}

		m.Ticket[i] = lticket

	}

	return nil
}

type ServerStatusUpdateMessage struct {
	Server GameServerInformations
}

func (m *ServerStatusUpdateMessage) ID() uint16 {
	return 50
}

func (m *ServerStatusUpdateMessage) Serialize(w Writer) error {

	if err := m.Server.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ServerStatusUpdateMessage) Deserialize(r Reader) error {

	var lserver GameServerInformations

	lserver.Deserialize(r)

	m.Server = lserver

	return nil
}

type ConsoleMessage struct {
	Type uint8

	Content string
}

func (m *ConsoleMessage) ID() uint16 {
	return 75
}

func (m *ConsoleMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	if err := w.WriteString(m.Content); err != nil {
		return err
	}

	return nil
}

func (m *ConsoleMessage) Deserialize(r Reader) error {

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	lcontent, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Content = lcontent

	return nil
}

type AdminCommandMessage struct {
	Content string
}

func (m *AdminCommandMessage) ID() uint16 {
	return 76
}

func (m *AdminCommandMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Content); err != nil {
		return err
	}

	return nil
}

func (m *AdminCommandMessage) Deserialize(r Reader) error {

	lcontent, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Content = lcontent

	return nil
}

type HelloGameMessage struct {
}

func (m *HelloGameMessage) ID() uint16 {
	return 101
}

func (m *HelloGameMessage) Serialize(w Writer) error {

	return nil
}

func (m *HelloGameMessage) Deserialize(r Reader) error {

	return nil
}

type AlreadyConnectedMessage struct {
}

func (m *AlreadyConnectedMessage) ID() uint16 {
	return 109
}

func (m *AlreadyConnectedMessage) Serialize(w Writer) error {

	return nil
}

func (m *AlreadyConnectedMessage) Deserialize(r Reader) error {

	return nil
}

type AuthenticationTicketMessage struct {
	Lang string

	Ticket string
}

func (m *AuthenticationTicketMessage) ID() uint16 {
	return 110
}

func (m *AuthenticationTicketMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Lang); err != nil {
		return err
	}

	if err := w.WriteString(m.Ticket); err != nil {
		return err
	}

	return nil
}

func (m *AuthenticationTicketMessage) Deserialize(r Reader) error {

	llang, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Lang = llang

	lticket, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Ticket = lticket

	return nil
}

type AuthenticationTicketAcceptedMessage struct {
}

func (m *AuthenticationTicketAcceptedMessage) ID() uint16 {
	return 111
}

func (m *AuthenticationTicketAcceptedMessage) Serialize(w Writer) error {

	return nil
}

func (m *AuthenticationTicketAcceptedMessage) Deserialize(r Reader) error {

	return nil
}

type AuthenticationTicketRefusedMessage struct {
}

func (m *AuthenticationTicketRefusedMessage) ID() uint16 {
	return 112
}

func (m *AuthenticationTicketRefusedMessage) Serialize(w Writer) error {

	return nil
}

func (m *AuthenticationTicketRefusedMessage) Deserialize(r Reader) error {

	return nil
}

type CharactersListRequestMessage struct {
}

func (m *CharactersListRequestMessage) ID() uint16 {
	return 150
}

func (m *CharactersListRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *CharactersListRequestMessage) Deserialize(r Reader) error {

	return nil
}

type CharactersListMessage struct {
	BasicCharactersListMessage

	HasStartupActions bool
}

func (m *CharactersListMessage) ID() uint16 {
	return 151
}

func (m *CharactersListMessage) Serialize(w Writer) error {

	if err := m.BasicCharactersListMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.HasStartupActions); err != nil {
		return err
	}

	return nil
}

func (m *CharactersListMessage) Deserialize(r Reader) error {

	if err := m.BasicCharactersListMessage.Deserialize(r); err != nil {
		return err
	}

	lhasStartupActions, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.HasStartupActions = lhasStartupActions

	return nil
}

type CharacterSelectionMessage struct {
	Id int64
}

func (m *CharacterSelectionMessage) ID() uint16 {
	return 152
}

func (m *CharacterSelectionMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *CharacterSelectionMessage) Deserialize(r Reader) error {

	lid, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type CharacterSelectedSuccessMessage struct {
	Infos CharacterBaseInformations

	IsCollectingStats bool
}

func (m *CharacterSelectedSuccessMessage) ID() uint16 {
	return 153
}

func (m *CharacterSelectedSuccessMessage) Serialize(w Writer) error {

	if err := m.Infos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.IsCollectingStats); err != nil {
		return err
	}

	return nil
}

func (m *CharacterSelectedSuccessMessage) Deserialize(r Reader) error {

	var linfos CharacterBaseInformations

	linfos.Deserialize(r)

	m.Infos = linfos

	lisCollectingStats, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsCollectingStats = lisCollectingStats

	return nil
}

type CharacterCreationRequestMessage struct {
	Name string

	Breed int8

	Sex bool

	Colors []int32

	CosmeticId uint16
}

func (m *CharacterCreationRequestMessage) ID() uint16 {
	return 160
}

func (m *CharacterCreationRequestMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Breed); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Sex); err != nil {
		return err
	}

	for i := range m.Colors {

		if err := w.WriteInt32(m.Colors[i]); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt16(m.CosmeticId); err != nil {
		return err
	}

	return nil
}

func (m *CharacterCreationRequestMessage) Deserialize(r Reader) error {

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

	m.Colors = make([]int32, 5)

	for i := range m.Colors {

		lcolors, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.Colors[i] = lcolors

	}

	lcosmeticId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CosmeticId = lcosmeticId

	return nil
}

type CharacterCreationResultMessage struct {
	Result uint8
}

func (m *CharacterCreationResultMessage) ID() uint16 {
	return 161
}

func (m *CharacterCreationResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Result); err != nil {
		return err
	}

	return nil
}

func (m *CharacterCreationResultMessage) Deserialize(r Reader) error {

	lresult, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Result = lresult

	return nil
}

type CharacterNameSuggestionRequestMessage struct {
}

func (m *CharacterNameSuggestionRequestMessage) ID() uint16 {
	return 162
}

func (m *CharacterNameSuggestionRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *CharacterNameSuggestionRequestMessage) Deserialize(r Reader) error {

	return nil
}

type CharacterNameSuggestionFailureMessage struct {
	Reason uint8
}

func (m *CharacterNameSuggestionFailureMessage) ID() uint16 {
	return 164
}

func (m *CharacterNameSuggestionFailureMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *CharacterNameSuggestionFailureMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type CharacterDeletionRequestMessage struct {
	CharacterId int64

	SecretAnswerHash string
}

func (m *CharacterDeletionRequestMessage) ID() uint16 {
	return 165
}

func (m *CharacterDeletionRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.CharacterId); err != nil {
		return err
	}

	if err := w.WriteString(m.SecretAnswerHash); err != nil {
		return err
	}

	return nil
}

func (m *CharacterDeletionRequestMessage) Deserialize(r Reader) error {

	lcharacterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.CharacterId = lcharacterId

	lsecretAnswerHash, err := r.ReadString()
	if err != nil {
		return err
	}

	m.SecretAnswerHash = lsecretAnswerHash

	return nil
}

type CharacterDeletionErrorMessage struct {
	Reason uint8
}

func (m *CharacterDeletionErrorMessage) ID() uint16 {
	return 166
}

func (m *CharacterDeletionErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *CharacterDeletionErrorMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type CharacterReplayRequestMessage struct {
	CharacterId int64
}

func (m *CharacterReplayRequestMessage) ID() uint16 {
	return 167
}

func (m *CharacterReplayRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.CharacterId); err != nil {
		return err
	}

	return nil
}

func (m *CharacterReplayRequestMessage) Deserialize(r Reader) error {

	lcharacterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.CharacterId = lcharacterId

	return nil
}

type SetCharacterRestrictionsMessage struct {
	ActorId float64

	Restrictions ActorRestrictionsInformations
}

func (m *SetCharacterRestrictionsMessage) ID() uint16 {
	return 170
}

func (m *SetCharacterRestrictionsMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.ActorId); err != nil {
		return err
	}

	if err := m.Restrictions.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *SetCharacterRestrictionsMessage) Deserialize(r Reader) error {

	lactorId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.ActorId = lactorId

	var lrestrictions ActorRestrictionsInformations

	lrestrictions.Deserialize(r)

	m.Restrictions = lrestrictions

	return nil
}

type BasicTimeMessage struct {
	Timestamp float64

	TimezoneOffset int16
}

func (m *BasicTimeMessage) ID() uint16 {
	return 175
}

func (m *BasicTimeMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.Timestamp); err != nil {
		return err
	}

	if err := w.WriteInt16(m.TimezoneOffset); err != nil {
		return err
	}

	return nil
}

func (m *BasicTimeMessage) Deserialize(r Reader) error {

	ltimestamp, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Timestamp = ltimestamp

	ltimezoneOffset, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.TimezoneOffset = ltimezoneOffset

	return nil
}

type BasicNoOperationMessage struct {
}

func (m *BasicNoOperationMessage) ID() uint16 {
	return 176
}

func (m *BasicNoOperationMessage) Serialize(w Writer) error {

	return nil
}

func (m *BasicNoOperationMessage) Deserialize(r Reader) error {

	return nil
}

type BasicDateMessage struct {
	Day uint8

	Month uint8

	Year uint16
}

func (m *BasicDateMessage) ID() uint16 {
	return 177
}

func (m *BasicDateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Day); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Month); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Year); err != nil {
		return err
	}

	return nil
}

func (m *BasicDateMessage) Deserialize(r Reader) error {

	lday, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Day = lday

	lmonth, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Month = lmonth

	lyear, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.Year = lyear

	return nil
}

type BasicWhoIsNoMatchMessage struct {
	Search string
}

func (m *BasicWhoIsNoMatchMessage) ID() uint16 {
	return 179
}

func (m *BasicWhoIsNoMatchMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Search); err != nil {
		return err
	}

	return nil
}

func (m *BasicWhoIsNoMatchMessage) Deserialize(r Reader) error {

	lsearch, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Search = lsearch

	return nil
}

type BasicWhoIsMessage struct {
	Self bool

	Position int8

	AccountNickname string

	AccountId uint32

	PlayerName string

	PlayerId int64

	AreaId int16

	ServerId int16

	OriginServerId int16

	SocialGroups []AbstractSocialGroupInfosIntrf

	Verbose bool

	PlayerState uint8
}

func (m *BasicWhoIsMessage) ID() uint16 {
	return 180
}

func (m *BasicWhoIsMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Self)

	setWrappedFlag(bbw0, 1, m.Verbose)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Position); err != nil {
		return err
	}

	if err := w.WriteString(m.AccountNickname); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	if err := w.WriteString(m.PlayerName); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.AreaId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.ServerId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.OriginServerId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.SocialGroups))); err != nil {
		return err
	}

	for i := range m.SocialGroups {

		if err := w.WriteUInt16(m.SocialGroups[i].ID()); err != nil {
			return err
		}

		if err := m.SocialGroups[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteUInt8(m.PlayerState); err != nil {
		return err
	}

	return nil
}

func (m *BasicWhoIsMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Self = getWrappedFlag(bbw0, 0)

	m.Verbose = getWrappedFlag(bbw0, 1)

	lposition, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Position = lposition

	laccountNickname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.AccountNickname = laccountNickname

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	lplayerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PlayerName = lplayerName

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	lareaId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AreaId = lareaId

	lserverId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ServerId = lserverId

	loriginServerId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.OriginServerId = loriginServerId

	lsocialGroupsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SocialGroups = make([]AbstractSocialGroupInfosIntrf, lsocialGroupsLen)

	for i := range m.SocialGroups {

		typesocialGroupsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lsocialGroups, err := GetType(typesocialGroupsID)
		if err != nil {
			return err
		}

		lsocialGroups.Deserialize(r)

		m.SocialGroups[i] = lsocialGroups.(AbstractSocialGroupInfosIntrf)

	}

	lplayerState, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PlayerState = lplayerState

	return nil
}

type BasicWhoIsRequestMessage struct {
	Verbose bool

	Search string
}

func (m *BasicWhoIsRequestMessage) ID() uint16 {
	return 181
}

func (m *BasicWhoIsRequestMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Verbose); err != nil {
		return err
	}

	if err := w.WriteString(m.Search); err != nil {
		return err
	}

	return nil
}

func (m *BasicWhoIsRequestMessage) Deserialize(r Reader) error {

	lverbose, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Verbose = lverbose

	lsearch, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Search = lsearch

	return nil
}

type BasicPingMessage struct {
	Quiet bool
}

func (m *BasicPingMessage) ID() uint16 {
	return 182
}

func (m *BasicPingMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Quiet); err != nil {
		return err
	}

	return nil
}

func (m *BasicPingMessage) Deserialize(r Reader) error {

	lquiet, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Quiet = lquiet

	return nil
}

type BasicPongMessage struct {
	Quiet bool
}

func (m *BasicPongMessage) ID() uint16 {
	return 183
}

func (m *BasicPongMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Quiet); err != nil {
		return err
	}

	return nil
}

func (m *BasicPongMessage) Deserialize(r Reader) error {

	lquiet, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Quiet = lquiet

	return nil
}

type SystemMessageDisplayMessage struct {
	HangUp bool

	MsgId uint16

	Parameters []string
}

func (m *SystemMessageDisplayMessage) ID() uint16 {
	return 189
}

func (m *SystemMessageDisplayMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.HangUp); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MsgId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Parameters))); err != nil {
		return err
	}

	for i := range m.Parameters {

		if err := w.WriteString(m.Parameters[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *SystemMessageDisplayMessage) Deserialize(r Reader) error {

	lhangUp, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.HangUp = lhangUp

	lmsgId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MsgId = lmsgId

	lparametersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Parameters = make([]string, lparametersLen)

	for i := range m.Parameters {

		lparameters, err := r.ReadString()
		if err != nil {
			return err
		}

		m.Parameters[i] = lparameters

	}

	return nil
}

type GameContextCreateMessage struct {
	Context uint8
}

func (m *GameContextCreateMessage) ID() uint16 {
	return 200
}

func (m *GameContextCreateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Context); err != nil {
		return err
	}

	return nil
}

func (m *GameContextCreateMessage) Deserialize(r Reader) error {

	lcontext, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Context = lcontext

	return nil
}

type GameContextDestroyMessage struct {
}

func (m *GameContextDestroyMessage) ID() uint16 {
	return 201
}

func (m *GameContextDestroyMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameContextDestroyMessage) Deserialize(r Reader) error {

	return nil
}

type MapFightCountMessage struct {
	FightCount uint16
}

func (m *MapFightCountMessage) ID() uint16 {
	return 210
}

func (m *MapFightCountMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.FightCount); err != nil {
		return err
	}

	return nil
}

func (m *MapFightCountMessage) Deserialize(r Reader) error {

	lfightCount, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.FightCount = lfightCount

	return nil
}

type CurrentMapMessage struct {
	MapId uint32

	MapKey string
}

func (m *CurrentMapMessage) ID() uint16 {
	return 220
}

func (m *CurrentMapMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteString(m.MapKey); err != nil {
		return err
	}

	return nil
}

func (m *CurrentMapMessage) Deserialize(r Reader) error {

	lmapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lmapKey, err := r.ReadString()
	if err != nil {
		return err
	}

	m.MapKey = lmapKey

	return nil
}

type ChangeMapMessage struct {
	MapId uint32
}

func (m *ChangeMapMessage) ID() uint16 {
	return 221
}

func (m *ChangeMapMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.MapId); err != nil {
		return err
	}

	return nil
}

func (m *ChangeMapMessage) Deserialize(r Reader) error {

	lmapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	return nil
}

type MapInformationsRequestMessage struct {
	MapId uint32
}

func (m *MapInformationsRequestMessage) ID() uint16 {
	return 225
}

func (m *MapInformationsRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.MapId); err != nil {
		return err
	}

	return nil
}

func (m *MapInformationsRequestMessage) Deserialize(r Reader) error {

	lmapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	return nil
}

type MapComplementaryInformationsDataMessage struct {
	SubAreaId uint16

	MapId uint32

	Houses []HouseInformationsIntrf

	Actors []GameRolePlayActorInformationsIntrf

	InteractiveElements []InteractiveElementIntrf

	StatedElements []StatedElement

	Obstacles []MapObstacle

	Fights []FightCommonInformations

	HasAggressiveMonsters bool

	FightStartPositions FightStartingPositions
}

func (m *MapComplementaryInformationsDataMessage) ID() uint16 {
	return 226
}

func (m *MapComplementaryInformationsDataMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Houses))); err != nil {
		return err
	}

	for i := range m.Houses {

		if err := w.WriteUInt16(m.Houses[i].ID()); err != nil {
			return err
		}

		if err := m.Houses[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Actors))); err != nil {
		return err
	}

	for i := range m.Actors {

		if err := w.WriteUInt16(m.Actors[i].ID()); err != nil {
			return err
		}

		if err := m.Actors[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.InteractiveElements))); err != nil {
		return err
	}

	for i := range m.InteractiveElements {

		if err := w.WriteUInt16(m.InteractiveElements[i].ID()); err != nil {
			return err
		}

		if err := m.InteractiveElements[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.StatedElements))); err != nil {
		return err
	}

	for i := range m.StatedElements {

		if err := m.StatedElements[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Obstacles))); err != nil {
		return err
	}

	for i := range m.Obstacles {

		if err := m.Obstacles[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Fights))); err != nil {
		return err
	}

	for i := range m.Fights {

		if err := m.Fights[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteBoolean(m.HasAggressiveMonsters); err != nil {
		return err
	}

	if err := m.FightStartPositions.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *MapComplementaryInformationsDataMessage) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lmapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lhousesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Houses = make([]HouseInformationsIntrf, lhousesLen)

	for i := range m.Houses {

		typehousesID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lhouses, err := GetType(typehousesID)
		if err != nil {
			return err
		}

		lhouses.Deserialize(r)

		m.Houses[i] = lhouses.(HouseInformationsIntrf)

	}

	lactorsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Actors = make([]GameRolePlayActorInformationsIntrf, lactorsLen)

	for i := range m.Actors {

		typeactorsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lactors, err := GetType(typeactorsID)
		if err != nil {
			return err
		}

		lactors.Deserialize(r)

		m.Actors[i] = lactors.(GameRolePlayActorInformationsIntrf)

	}

	linteractiveElementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.InteractiveElements = make([]InteractiveElementIntrf, linteractiveElementsLen)

	for i := range m.InteractiveElements {

		typeinteractiveElementsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		linteractiveElements, err := GetType(typeinteractiveElementsID)
		if err != nil {
			return err
		}

		linteractiveElements.Deserialize(r)

		m.InteractiveElements[i] = linteractiveElements.(InteractiveElementIntrf)

	}

	lstatedElementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.StatedElements = make([]StatedElement, lstatedElementsLen)

	for i := range m.StatedElements {

		var lstatedElements StatedElement

		lstatedElements.Deserialize(r)

		m.StatedElements[i] = lstatedElements

	}

	lobstaclesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Obstacles = make([]MapObstacle, lobstaclesLen)

	for i := range m.Obstacles {

		var lobstacles MapObstacle

		lobstacles.Deserialize(r)

		m.Obstacles[i] = lobstacles

	}

	lfightsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Fights = make([]FightCommonInformations, lfightsLen)

	for i := range m.Fights {

		var lfights FightCommonInformations

		lfights.Deserialize(r)

		m.Fights[i] = lfights

	}

	lhasAggressiveMonsters, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.HasAggressiveMonsters = lhasAggressiveMonsters

	var lfightStartPositions FightStartingPositions

	lfightStartPositions.Deserialize(r)

	m.FightStartPositions = lfightStartPositions

	return nil
}

type GameContextCreateRequestMessage struct {
}

func (m *GameContextCreateRequestMessage) ID() uint16 {
	return 250
}

func (m *GameContextCreateRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameContextCreateRequestMessage) Deserialize(r Reader) error {

	return nil
}

type GameContextRemoveElementMessage struct {
	Id float64
}

func (m *GameContextRemoveElementMessage) ID() uint16 {
	return 251
}

func (m *GameContextRemoveElementMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *GameContextRemoveElementMessage) Deserialize(r Reader) error {

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type GameContextRemoveMultipleElementsMessage struct {
	ElementsIds []float64
}

func (m *GameContextRemoveMultipleElementsMessage) ID() uint16 {
	return 252
}

func (m *GameContextRemoveMultipleElementsMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ElementsIds))); err != nil {
		return err
	}

	for i := range m.ElementsIds {

		if err := w.WriteDouble(m.ElementsIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameContextRemoveMultipleElementsMessage) Deserialize(r Reader) error {

	lelementsIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ElementsIds = make([]float64, lelementsIdsLen)

	for i := range m.ElementsIds {

		lelementsIds, err := r.ReadDouble()
		if err != nil {
			return err
		}

		m.ElementsIds[i] = lelementsIds

	}

	return nil
}

type GameContextMoveElementMessage struct {
	Movement EntityMovementInformations
}

func (m *GameContextMoveElementMessage) ID() uint16 {
	return 253
}

func (m *GameContextMoveElementMessage) Serialize(w Writer) error {

	if err := m.Movement.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameContextMoveElementMessage) Deserialize(r Reader) error {

	var lmovement EntityMovementInformations

	lmovement.Deserialize(r)

	m.Movement = lmovement

	return nil
}

type GameContextMoveMultipleElementsMessage struct {
	Movements []EntityMovementInformations
}

func (m *GameContextMoveMultipleElementsMessage) ID() uint16 {
	return 254
}

func (m *GameContextMoveMultipleElementsMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Movements))); err != nil {
		return err
	}

	for i := range m.Movements {

		if err := m.Movements[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameContextMoveMultipleElementsMessage) Deserialize(r Reader) error {

	lmovementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Movements = make([]EntityMovementInformations, lmovementsLen)

	for i := range m.Movements {

		var lmovements EntityMovementInformations

		lmovements.Deserialize(r)

		m.Movements[i] = lmovements

	}

	return nil
}

type GameContextQuitMessage struct {
}

func (m *GameContextQuitMessage) ID() uint16 {
	return 255
}

func (m *GameContextQuitMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameContextQuitMessage) Deserialize(r Reader) error {

	return nil
}

type GameRolePlayRemoveChallengeMessage struct {
	FightId int32
}

func (m *GameRolePlayRemoveChallengeMessage) ID() uint16 {
	return 300
}

func (m *GameRolePlayRemoveChallengeMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.FightId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayRemoveChallengeMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	return nil
}

type GameRolePlayShowChallengeMessage struct {
	CommonsInfos FightCommonInformations
}

func (m *GameRolePlayShowChallengeMessage) ID() uint16 {
	return 301
}

func (m *GameRolePlayShowChallengeMessage) Serialize(w Writer) error {

	if err := m.CommonsInfos.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayShowChallengeMessage) Deserialize(r Reader) error {

	var lcommonsInfos FightCommonInformations

	lcommonsInfos.Deserialize(r)

	m.CommonsInfos = lcommonsInfos

	return nil
}

type CharacterStatsListMessage struct {
	Stats CharacterCharacteristicsInformations
}

func (m *CharacterStatsListMessage) ID() uint16 {
	return 500
}

func (m *CharacterStatsListMessage) Serialize(w Writer) error {

	if err := m.Stats.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *CharacterStatsListMessage) Deserialize(r Reader) error {

	var lstats CharacterCharacteristicsInformations

	lstats.Deserialize(r)

	m.Stats = lstats

	return nil
}

type GameFightStartingMessage struct {
	FightType uint8

	AttackerId float64

	DefenderId float64
}

func (m *GameFightStartingMessage) ID() uint16 {
	return 700
}

func (m *GameFightStartingMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.FightType); err != nil {
		return err
	}

	if err := w.WriteDouble(m.AttackerId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.DefenderId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightStartingMessage) Deserialize(r Reader) error {

	lfightType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.FightType = lfightType

	lattackerId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.AttackerId = lattackerId

	ldefenderId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DefenderId = ldefenderId

	return nil
}

type GameFightJoinRequestMessage struct {
	FighterId float64

	FightId int32
}

func (m *GameFightJoinRequestMessage) ID() uint16 {
	return 701
}

func (m *GameFightJoinRequestMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.FighterId); err != nil {
		return err
	}

	if err := w.WriteInt32(m.FightId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightJoinRequestMessage) Deserialize(r Reader) error {

	lfighterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.FighterId = lfighterId

	lfightId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	return nil
}

type GameFightJoinMessage struct {
	IsTeamPhase bool

	CanBeCancelled bool

	CanSayReady bool

	IsFightStarted bool

	TimeMaxBeforeFightStart uint16

	FightType uint8
}

func (m *GameFightJoinMessage) ID() uint16 {
	return 702
}

func (m *GameFightJoinMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.IsTeamPhase)

	setWrappedFlag(bbw0, 1, m.CanBeCancelled)

	setWrappedFlag(bbw0, 2, m.CanSayReady)

	setWrappedFlag(bbw0, 3, m.IsFightStarted)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.TimeMaxBeforeFightStart); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.FightType); err != nil {
		return err
	}

	return nil
}

func (m *GameFightJoinMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.IsTeamPhase = getWrappedFlag(bbw0, 0)

	m.CanBeCancelled = getWrappedFlag(bbw0, 1)

	m.CanSayReady = getWrappedFlag(bbw0, 2)

	m.IsFightStarted = getWrappedFlag(bbw0, 3)

	ltimeMaxBeforeFightStart, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.TimeMaxBeforeFightStart = ltimeMaxBeforeFightStart

	lfightType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.FightType = lfightType

	return nil
}

type GameFightPlacementPossiblePositionsMessage struct {
	PositionsForChallengers []uint16

	PositionsForDefenders []uint16

	TeamNumber uint8
}

func (m *GameFightPlacementPossiblePositionsMessage) ID() uint16 {
	return 703
}

func (m *GameFightPlacementPossiblePositionsMessage) Serialize(w Writer) error {

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

	if err := w.WriteUInt8(m.TeamNumber); err != nil {
		return err
	}

	return nil
}

func (m *GameFightPlacementPossiblePositionsMessage) Deserialize(r Reader) error {

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

	lteamNumber, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TeamNumber = lteamNumber

	return nil
}

type GameFightPlacementPositionRequestMessage struct {
	CellId uint16
}

func (m *GameFightPlacementPositionRequestMessage) ID() uint16 {
	return 704
}

func (m *GameFightPlacementPositionRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightPlacementPositionRequestMessage) Deserialize(r Reader) error {

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type GameFightOptionToggleMessage struct {
	Option uint8
}

func (m *GameFightOptionToggleMessage) ID() uint16 {
	return 707
}

func (m *GameFightOptionToggleMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Option); err != nil {
		return err
	}

	return nil
}

func (m *GameFightOptionToggleMessage) Deserialize(r Reader) error {

	loption, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Option = loption

	return nil
}

type GameFightReadyMessage struct {
	IsReady bool
}

func (m *GameFightReadyMessage) ID() uint16 {
	return 708
}

func (m *GameFightReadyMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.IsReady); err != nil {
		return err
	}

	return nil
}

func (m *GameFightReadyMessage) Deserialize(r Reader) error {

	lisReady, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsReady = lisReady

	return nil
}

type GameFightRemoveTeamMemberMessage struct {
	FightId uint16

	TeamId uint8

	CharId float64
}

func (m *GameFightRemoveTeamMemberMessage) ID() uint16 {
	return 711
}

func (m *GameFightRemoveTeamMemberMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.FightId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.TeamId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.CharId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightRemoveTeamMemberMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lteamId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TeamId = lteamId

	lcharId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.CharId = lcharId

	return nil
}

type GameFightStartMessage struct {
	Idols []Idol
}

func (m *GameFightStartMessage) ID() uint16 {
	return 712
}

func (m *GameFightStartMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Idols))); err != nil {
		return err
	}

	for i := range m.Idols {

		if err := m.Idols[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameFightStartMessage) Deserialize(r Reader) error {

	lidolsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Idols = make([]Idol, lidolsLen)

	for i := range m.Idols {

		var lidols Idol

		lidols.Deserialize(r)

		m.Idols[i] = lidols

	}

	return nil
}

type GameFightTurnListMessage struct {
	Ids []float64

	DeadsIds []float64
}

func (m *GameFightTurnListMessage) ID() uint16 {
	return 713
}

func (m *GameFightTurnListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Ids))); err != nil {
		return err
	}

	for i := range m.Ids {

		if err := w.WriteDouble(m.Ids[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.DeadsIds))); err != nil {
		return err
	}

	for i := range m.DeadsIds {

		if err := w.WriteDouble(m.DeadsIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameFightTurnListMessage) Deserialize(r Reader) error {

	lidsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Ids = make([]float64, lidsLen)

	for i := range m.Ids {

		lids, err := r.ReadDouble()
		if err != nil {
			return err
		}

		m.Ids[i] = lids

	}

	ldeadsIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DeadsIds = make([]float64, ldeadsIdsLen)

	for i := range m.DeadsIds {

		ldeadsIds, err := r.ReadDouble()
		if err != nil {
			return err
		}

		m.DeadsIds[i] = ldeadsIds

	}

	return nil
}

type GameFightTurnStartMessage struct {
	Id float64

	WaitTime uint32
}

func (m *GameFightTurnStartMessage) ID() uint16 {
	return 714
}

func (m *GameFightTurnStartMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.WaitTime); err != nil {
		return err
	}

	return nil
}

func (m *GameFightTurnStartMessage) Deserialize(r Reader) error {

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	lwaitTime, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.WaitTime = lwaitTime

	return nil
}

type GameFightTurnReadyRequestMessage struct {
	Id float64
}

func (m *GameFightTurnReadyRequestMessage) ID() uint16 {
	return 715
}

func (m *GameFightTurnReadyRequestMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *GameFightTurnReadyRequestMessage) Deserialize(r Reader) error {

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type GameFightTurnReadyMessage struct {
	IsReady bool
}

func (m *GameFightTurnReadyMessage) ID() uint16 {
	return 716
}

func (m *GameFightTurnReadyMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.IsReady); err != nil {
		return err
	}

	return nil
}

func (m *GameFightTurnReadyMessage) Deserialize(r Reader) error {

	lisReady, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsReady = lisReady

	return nil
}

type GameFightTurnFinishMessage struct {
	IsAfk bool
}

func (m *GameFightTurnFinishMessage) ID() uint16 {
	return 718
}

func (m *GameFightTurnFinishMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.IsAfk); err != nil {
		return err
	}

	return nil
}

func (m *GameFightTurnFinishMessage) Deserialize(r Reader) error {

	lisAfk, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsAfk = lisAfk

	return nil
}

type GameFightTurnEndMessage struct {
	Id float64
}

func (m *GameFightTurnEndMessage) ID() uint16 {
	return 719
}

func (m *GameFightTurnEndMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *GameFightTurnEndMessage) Deserialize(r Reader) error {

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type GameFightEndMessage struct {
	Duration uint32

	AgeBonus int16

	LootShareLimitMalus int16

	Results []FightResultListEntryIntrf

	NamedPartyTeamsOutcomes []NamedPartyTeamWithOutcome
}

func (m *GameFightEndMessage) ID() uint16 {
	return 720
}

func (m *GameFightEndMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.Duration); err != nil {
		return err
	}

	if err := w.WriteInt16(m.AgeBonus); err != nil {
		return err
	}

	if err := w.WriteInt16(m.LootShareLimitMalus); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Results))); err != nil {
		return err
	}

	for i := range m.Results {

		if err := w.WriteUInt16(m.Results[i].ID()); err != nil {
			return err
		}

		if err := m.Results[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.NamedPartyTeamsOutcomes))); err != nil {
		return err
	}

	for i := range m.NamedPartyTeamsOutcomes {

		if err := m.NamedPartyTeamsOutcomes[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameFightEndMessage) Deserialize(r Reader) error {

	lduration, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.Duration = lduration

	lageBonus, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AgeBonus = lageBonus

	llootShareLimitMalus, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.LootShareLimitMalus = llootShareLimitMalus

	lresultsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Results = make([]FightResultListEntryIntrf, lresultsLen)

	for i := range m.Results {

		typeresultsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lresults, err := GetType(typeresultsID)
		if err != nil {
			return err
		}

		lresults.Deserialize(r)

		m.Results[i] = lresults.(FightResultListEntryIntrf)

	}

	lnamedPartyTeamsOutcomesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.NamedPartyTeamsOutcomes = make([]NamedPartyTeamWithOutcome, lnamedPartyTeamsOutcomesLen)

	for i := range m.NamedPartyTeamsOutcomes {

		var lnamedPartyTeamsOutcomes NamedPartyTeamWithOutcome

		lnamedPartyTeamsOutcomes.Deserialize(r)

		m.NamedPartyTeamsOutcomes[i] = lnamedPartyTeamsOutcomes

	}

	return nil
}

type GameFightLeaveMessage struct {
	CharId float64
}

func (m *GameFightLeaveMessage) ID() uint16 {
	return 721
}

func (m *GameFightLeaveMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.CharId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightLeaveMessage) Deserialize(r Reader) error {

	lcharId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.CharId = lcharId

	return nil
}

type GameFightHumanReadyStateMessage struct {
	CharacterId int64

	IsReady bool
}

func (m *GameFightHumanReadyStateMessage) ID() uint16 {
	return 740
}

func (m *GameFightHumanReadyStateMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.CharacterId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.IsReady); err != nil {
		return err
	}

	return nil
}

func (m *GameFightHumanReadyStateMessage) Deserialize(r Reader) error {

	lcharacterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.CharacterId = lcharacterId

	lisReady, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsReady = lisReady

	return nil
}

type GameRolePlayFreeSoulRequestMessage struct {
}

func (m *GameRolePlayFreeSoulRequestMessage) ID() uint16 {
	return 745
}

func (m *GameRolePlayFreeSoulRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameRolePlayFreeSoulRequestMessage) Deserialize(r Reader) error {

	return nil
}

type GameRolePlayGameOverMessage struct {
}

func (m *GameRolePlayGameOverMessage) ID() uint16 {
	return 746
}

func (m *GameRolePlayGameOverMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameRolePlayGameOverMessage) Deserialize(r Reader) error {

	return nil
}

type TextInformationMessage struct {
	MsgType uint8

	MsgId uint16

	Parameters []string
}

func (m *TextInformationMessage) ID() uint16 {
	return 780
}

func (m *TextInformationMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.MsgType); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MsgId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Parameters))); err != nil {
		return err
	}

	for i := range m.Parameters {

		if err := w.WriteString(m.Parameters[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *TextInformationMessage) Deserialize(r Reader) error {

	lmsgType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MsgType = lmsgType

	lmsgId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MsgId = lmsgId

	lparametersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Parameters = make([]string, lparametersLen)

	for i := range m.Parameters {

		lparameters, err := r.ReadString()
		if err != nil {
			return err
		}

		m.Parameters[i] = lparameters

	}

	return nil
}

type ChatSmileyRequestMessage struct {
	SmileyId uint16
}

func (m *ChatSmileyRequestMessage) ID() uint16 {
	return 800
}

func (m *ChatSmileyRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SmileyId); err != nil {
		return err
	}

	return nil
}

func (m *ChatSmileyRequestMessage) Deserialize(r Reader) error {

	lsmileyId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SmileyId = lsmileyId

	return nil
}

type ChatSmileyMessage struct {
	EntityId float64

	SmileyId uint16

	AccountId uint32
}

func (m *ChatSmileyMessage) ID() uint16 {
	return 801
}

func (m *ChatSmileyMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.EntityId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SmileyId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	return nil
}

func (m *ChatSmileyMessage) Deserialize(r Reader) error {

	lentityId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.EntityId = lentityId

	lsmileyId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SmileyId = lsmileyId

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	return nil
}

type ChatMessageReportMessage struct {
	SenderName string

	Content string

	Timestamp uint32

	Channel uint8

	Fingerprint string

	Reason uint8
}

func (m *ChatMessageReportMessage) ID() uint16 {
	return 821
}

func (m *ChatMessageReportMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.SenderName); err != nil {
		return err
	}

	if err := w.WriteString(m.Content); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.Timestamp); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Channel); err != nil {
		return err
	}

	if err := w.WriteString(m.Fingerprint); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *ChatMessageReportMessage) Deserialize(r Reader) error {

	lsenderName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.SenderName = lsenderName

	lcontent, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Content = lcontent

	ltimestamp, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.Timestamp = ltimestamp

	lchannel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Channel = lchannel

	lfingerprint, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Fingerprint = lfingerprint

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type ChatAbstractClientMessage struct {
	Content string
}

func (m *ChatAbstractClientMessage) ID() uint16 {
	return 850
}

func (m *ChatAbstractClientMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Content); err != nil {
		return err
	}

	return nil
}

func (m *ChatAbstractClientMessage) Deserialize(r Reader) error {

	lcontent, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Content = lcontent

	return nil
}

type ChatClientPrivateMessage struct {
	ChatAbstractClientMessage

	Receiver string
}

func (m *ChatClientPrivateMessage) ID() uint16 {
	return 851
}

func (m *ChatClientPrivateMessage) Serialize(w Writer) error {

	if err := m.ChatAbstractClientMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Receiver); err != nil {
		return err
	}

	return nil
}

func (m *ChatClientPrivateMessage) Deserialize(r Reader) error {

	if err := m.ChatAbstractClientMessage.Deserialize(r); err != nil {
		return err
	}

	lreceiver, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Receiver = lreceiver

	return nil
}

type ChatClientPrivateWithObjectMessage struct {
	ChatClientPrivateMessage

	Objects []ObjectItem
}

func (m *ChatClientPrivateWithObjectMessage) ID() uint16 {
	return 852
}

func (m *ChatClientPrivateWithObjectMessage) Serialize(w Writer) error {

	if err := m.ChatClientPrivateMessage.Serialize(w); err != nil {
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

func (m *ChatClientPrivateWithObjectMessage) Deserialize(r Reader) error {

	if err := m.ChatClientPrivateMessage.Deserialize(r); err != nil {
		return err
	}

	lobjectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Objects = make([]ObjectItem, lobjectsLen)

	for i := range m.Objects {

		var lobjects ObjectItem

		lobjects.Deserialize(r)

		m.Objects[i] = lobjects

	}

	return nil
}

type ChatClientMultiMessage struct {
	ChatAbstractClientMessage

	Channel uint8
}

func (m *ChatClientMultiMessage) ID() uint16 {
	return 861
}

func (m *ChatClientMultiMessage) Serialize(w Writer) error {

	if err := m.ChatAbstractClientMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Channel); err != nil {
		return err
	}

	return nil
}

func (m *ChatClientMultiMessage) Deserialize(r Reader) error {

	if err := m.ChatAbstractClientMessage.Deserialize(r); err != nil {
		return err
	}

	lchannel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Channel = lchannel

	return nil
}

type ChatClientMultiWithObjectMessage struct {
	ChatClientMultiMessage

	Objects []ObjectItem
}

func (m *ChatClientMultiWithObjectMessage) ID() uint16 {
	return 862
}

func (m *ChatClientMultiWithObjectMessage) Serialize(w Writer) error {

	if err := m.ChatClientMultiMessage.Serialize(w); err != nil {
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

func (m *ChatClientMultiWithObjectMessage) Deserialize(r Reader) error {

	if err := m.ChatClientMultiMessage.Deserialize(r); err != nil {
		return err
	}

	lobjectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Objects = make([]ObjectItem, lobjectsLen)

	for i := range m.Objects {

		var lobjects ObjectItem

		lobjects.Deserialize(r)

		m.Objects[i] = lobjects

	}

	return nil
}

type ChatErrorMessage struct {
	Reason uint8
}

func (m *ChatErrorMessage) ID() uint16 {
	return 870
}

func (m *ChatErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *ChatErrorMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type ChatAbstractServerMessage struct {
	Channel uint8

	Content string

	Timestamp uint32

	Fingerprint string
}

func (m *ChatAbstractServerMessage) ID() uint16 {
	return 880
}

func (m *ChatAbstractServerMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Channel); err != nil {
		return err
	}

	if err := w.WriteString(m.Content); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.Timestamp); err != nil {
		return err
	}

	if err := w.WriteString(m.Fingerprint); err != nil {
		return err
	}

	return nil
}

func (m *ChatAbstractServerMessage) Deserialize(r Reader) error {

	lchannel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Channel = lchannel

	lcontent, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Content = lcontent

	ltimestamp, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.Timestamp = ltimestamp

	lfingerprint, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Fingerprint = lfingerprint

	return nil
}

type ChatServerMessage struct {
	ChatAbstractServerMessage

	SenderId float64

	SenderName string

	SenderAccountId uint32
}

func (m *ChatServerMessage) ID() uint16 {
	return 881
}

func (m *ChatServerMessage) Serialize(w Writer) error {

	if err := m.ChatAbstractServerMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.SenderId); err != nil {
		return err
	}

	if err := w.WriteString(m.SenderName); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.SenderAccountId); err != nil {
		return err
	}

	return nil
}

func (m *ChatServerMessage) Deserialize(r Reader) error {

	if err := m.ChatAbstractServerMessage.Deserialize(r); err != nil {
		return err
	}

	lsenderId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.SenderId = lsenderId

	lsenderName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.SenderName = lsenderName

	lsenderAccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.SenderAccountId = lsenderAccountId

	return nil
}

type ChatServerCopyMessage struct {
	ChatAbstractServerMessage

	ReceiverId int64

	ReceiverName string
}

func (m *ChatServerCopyMessage) ID() uint16 {
	return 882
}

func (m *ChatServerCopyMessage) Serialize(w Writer) error {

	if err := m.ChatAbstractServerMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ReceiverId); err != nil {
		return err
	}

	if err := w.WriteString(m.ReceiverName); err != nil {
		return err
	}

	return nil
}

func (m *ChatServerCopyMessage) Deserialize(r Reader) error {

	if err := m.ChatAbstractServerMessage.Deserialize(r); err != nil {
		return err
	}

	lreceiverId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ReceiverId = lreceiverId

	lreceiverName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.ReceiverName = lreceiverName

	return nil
}

type ChatServerWithObjectMessage struct {
	ChatServerMessage

	Objects []ObjectItem
}

func (m *ChatServerWithObjectMessage) ID() uint16 {
	return 883
}

func (m *ChatServerWithObjectMessage) Serialize(w Writer) error {

	if err := m.ChatServerMessage.Serialize(w); err != nil {
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

func (m *ChatServerWithObjectMessage) Deserialize(r Reader) error {

	if err := m.ChatServerMessage.Deserialize(r); err != nil {
		return err
	}

	lobjectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Objects = make([]ObjectItem, lobjectsLen)

	for i := range m.Objects {

		var lobjects ObjectItem

		lobjects.Deserialize(r)

		m.Objects[i] = lobjects

	}

	return nil
}

type ChatServerCopyWithObjectMessage struct {
	ChatServerCopyMessage

	Objects []ObjectItem
}

func (m *ChatServerCopyWithObjectMessage) ID() uint16 {
	return 884
}

func (m *ChatServerCopyWithObjectMessage) Serialize(w Writer) error {

	if err := m.ChatServerCopyMessage.Serialize(w); err != nil {
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

func (m *ChatServerCopyWithObjectMessage) Deserialize(r Reader) error {

	if err := m.ChatServerCopyMessage.Deserialize(r); err != nil {
		return err
	}

	lobjectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Objects = make([]ObjectItem, lobjectsLen)

	for i := range m.Objects {

		var lobjects ObjectItem

		lobjects.Deserialize(r)

		m.Objects[i] = lobjects

	}

	return nil
}

type ChannelEnablingMessage struct {
	Channel uint8

	Enable bool
}

func (m *ChannelEnablingMessage) ID() uint16 {
	return 890
}

func (m *ChannelEnablingMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Channel); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *ChannelEnablingMessage) Deserialize(r Reader) error {

	lchannel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Channel = lchannel

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type ChannelEnablingChangeMessage struct {
	Channel uint8

	Enable bool
}

func (m *ChannelEnablingChangeMessage) ID() uint16 {
	return 891
}

func (m *ChannelEnablingChangeMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Channel); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *ChannelEnablingChangeMessage) Deserialize(r Reader) error {

	lchannel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Channel = lchannel

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type EnabledChannelsMessage struct {
	Channels []uint8

	Disallowed []uint8
}

func (m *EnabledChannelsMessage) ID() uint16 {
	return 892
}

func (m *EnabledChannelsMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Channels))); err != nil {
		return err
	}

	for i := range m.Channels {

		if err := w.WriteUInt8(m.Channels[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Disallowed))); err != nil {
		return err
	}

	for i := range m.Disallowed {

		if err := w.WriteUInt8(m.Disallowed[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *EnabledChannelsMessage) Deserialize(r Reader) error {

	lchannelsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Channels = make([]uint8, lchannelsLen)

	for i := range m.Channels {

		lchannels, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.Channels[i] = lchannels

	}

	ldisallowedLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Disallowed = make([]uint8, ldisallowedLen)

	for i := range m.Disallowed {

		ldisallowed, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.Disallowed[i] = ldisallowed

	}

	return nil
}

type GameMapChangeOrientationRequestMessage struct {
	Direction uint8
}

func (m *GameMapChangeOrientationRequestMessage) ID() uint16 {
	return 945
}

func (m *GameMapChangeOrientationRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Direction); err != nil {
		return err
	}

	return nil
}

func (m *GameMapChangeOrientationRequestMessage) Deserialize(r Reader) error {

	ldirection, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Direction = ldirection

	return nil
}

type GameMapChangeOrientationMessage struct {
	Orientation ActorOrientation
}

func (m *GameMapChangeOrientationMessage) ID() uint16 {
	return 946
}

func (m *GameMapChangeOrientationMessage) Serialize(w Writer) error {

	if err := m.Orientation.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameMapChangeOrientationMessage) Deserialize(r Reader) error {

	var lorientation ActorOrientation

	lorientation.Deserialize(r)

	m.Orientation = lorientation

	return nil
}

type GameMapMovementRequestMessage struct {
	KeyMovements []uint16

	MapId uint32
}

func (m *GameMapMovementRequestMessage) ID() uint16 {
	return 950
}

func (m *GameMapMovementRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.KeyMovements))); err != nil {
		return err
	}

	for i := range m.KeyMovements {

		if err := w.WriteUInt16(m.KeyMovements[i]); err != nil {
			return err
		}

	}

	if err := w.WriteUInt32(m.MapId); err != nil {
		return err
	}

	return nil
}

func (m *GameMapMovementRequestMessage) Deserialize(r Reader) error {

	lkeyMovementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.KeyMovements = make([]uint16, lkeyMovementsLen)

	for i := range m.KeyMovements {

		lkeyMovements, err := r.ReadUInt16()
		if err != nil {
			return err
		}

		m.KeyMovements[i] = lkeyMovements

	}

	lmapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	return nil
}

type GameMapMovementMessage struct {
	KeyMovements []uint16

	ForcedDirection int16

	ActorId float64
}

func (m *GameMapMovementMessage) ID() uint16 {
	return 951
}

func (m *GameMapMovementMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.KeyMovements))); err != nil {
		return err
	}

	for i := range m.KeyMovements {

		if err := w.WriteUInt16(m.KeyMovements[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(m.ForcedDirection); err != nil {
		return err
	}

	if err := w.WriteDouble(m.ActorId); err != nil {
		return err
	}

	return nil
}

func (m *GameMapMovementMessage) Deserialize(r Reader) error {

	lkeyMovementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.KeyMovements = make([]uint16, lkeyMovementsLen)

	for i := range m.KeyMovements {

		lkeyMovements, err := r.ReadUInt16()
		if err != nil {
			return err
		}

		m.KeyMovements[i] = lkeyMovements

	}

	lforcedDirection, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ForcedDirection = lforcedDirection

	lactorId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.ActorId = lactorId

	return nil
}

type GameMapMovementConfirmMessage struct {
}

func (m *GameMapMovementConfirmMessage) ID() uint16 {
	return 952
}

func (m *GameMapMovementConfirmMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameMapMovementConfirmMessage) Deserialize(r Reader) error {

	return nil
}

type GameMapMovementCancelMessage struct {
	CellId uint16
}

func (m *GameMapMovementCancelMessage) ID() uint16 {
	return 953
}

func (m *GameMapMovementCancelMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *GameMapMovementCancelMessage) Deserialize(r Reader) error {

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type GameMapNoMovementMessage struct {
	CellX int16

	CellY int16
}

func (m *GameMapNoMovementMessage) ID() uint16 {
	return 954
}

func (m *GameMapNoMovementMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(m.CellX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.CellY); err != nil {
		return err
	}

	return nil
}

func (m *GameMapNoMovementMessage) Deserialize(r Reader) error {

	lcellX, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CellX = lcellX

	lcellY, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CellY = lcellY

	return nil
}

type SequenceStartMessage struct {
	SequenceType int8

	AuthorId float64
}

func (m *SequenceStartMessage) ID() uint16 {
	return 955
}

func (m *SequenceStartMessage) Serialize(w Writer) error {

	if err := w.WriteInt8(m.SequenceType); err != nil {
		return err
	}

	if err := w.WriteDouble(m.AuthorId); err != nil {
		return err
	}

	return nil
}

func (m *SequenceStartMessage) Deserialize(r Reader) error {

	lsequenceType, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.SequenceType = lsequenceType

	lauthorId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.AuthorId = lauthorId

	return nil
}

type SequenceEndMessage struct {
	ActionId uint16

	AuthorId float64

	SequenceType int8
}

func (m *SequenceEndMessage) ID() uint16 {
	return 956
}

func (m *SequenceEndMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ActionId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.AuthorId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.SequenceType); err != nil {
		return err
	}

	return nil
}

func (m *SequenceEndMessage) Deserialize(r Reader) error {

	lactionId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ActionId = lactionId

	lauthorId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.AuthorId = lauthorId

	lsequenceType, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.SequenceType = lsequenceType

	return nil
}

type GameActionAcknowledgementMessage struct {
	Valid bool

	ActionId int8
}

func (m *GameActionAcknowledgementMessage) ID() uint16 {
	return 957
}

func (m *GameActionAcknowledgementMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Valid); err != nil {
		return err
	}

	if err := w.WriteInt8(m.ActionId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionAcknowledgementMessage) Deserialize(r Reader) error {

	lvalid, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Valid = lvalid

	lactionId, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.ActionId = lactionId

	return nil
}

type AbstractGameActionMessage struct {
	ActionId uint16

	SourceId float64
}

func (m *AbstractGameActionMessage) ID() uint16 {
	return 1000
}

func (m *AbstractGameActionMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ActionId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.SourceId); err != nil {
		return err
	}

	return nil
}

func (m *AbstractGameActionMessage) Deserialize(r Reader) error {

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

	return nil
}

type AbstractGameActionWithAckMessage struct {
	AbstractGameActionMessage

	WaitAckId int16
}

func (m *AbstractGameActionWithAckMessage) ID() uint16 {
	return 1001
}

func (m *AbstractGameActionWithAckMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WaitAckId); err != nil {
		return err
	}

	return nil
}

func (m *AbstractGameActionWithAckMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	lwaitAckId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.WaitAckId = lwaitAckId

	return nil
}

type GameActionNoopMessage struct {
}

func (m *GameActionNoopMessage) ID() uint16 {
	return 1002
}

func (m *GameActionNoopMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameActionNoopMessage) Deserialize(r Reader) error {

	return nil
}

type GameActionFightTackledMessage struct {
	AbstractGameActionMessage

	TacklersIds []float64
}

func (m *GameActionFightTackledMessage) ID() uint16 {
	return 1004
}

func (m *GameActionFightTackledMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.TacklersIds))); err != nil {
		return err
	}

	for i := range m.TacklersIds {

		if err := w.WriteDouble(m.TacklersIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameActionFightTackledMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltacklersIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.TacklersIds = make([]float64, ltacklersIdsLen)

	for i := range m.TacklersIds {

		ltacklersIds, err := r.ReadDouble()
		if err != nil {
			return err
		}

		m.TacklersIds[i] = ltacklersIds

	}

	return nil
}

type GameActionFightCastRequestMessage struct {
	SpellId uint16

	CellId int16
}

func (m *GameActionFightCastRequestMessage) ID() uint16 {
	return 1005
}

func (m *GameActionFightCastRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightCastRequestMessage) Deserialize(r Reader) error {

	lspellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	lcellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type GameActionFightSpellCastMessage struct {
	AbstractGameActionFightTargetedAbilityMessage

	SpellId uint16

	SpellLevel int16

	PortalsIds []int16
}

func (m *GameActionFightSpellCastMessage) ID() uint16 {
	return 1010
}

func (m *GameActionFightSpellCastMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionFightTargetedAbilityMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.SpellLevel); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.PortalsIds))); err != nil {
		return err
	}

	for i := range m.PortalsIds {

		if err := w.WriteInt16(m.PortalsIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameActionFightSpellCastMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionFightTargetedAbilityMessage.Deserialize(r); err != nil {
		return err
	}

	lspellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	lspellLevel, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SpellLevel = lspellLevel

	lportalsIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PortalsIds = make([]int16, lportalsIdsLen)

	for i := range m.PortalsIds {

		lportalsIds, err := r.ReadInt16()
		if err != nil {
			return err
		}

		m.PortalsIds[i] = lportalsIds

	}

	return nil
}

type GameActionFightPointsVariationMessage struct {
	AbstractGameActionMessage

	TargetId float64

	Delta int16
}

func (m *GameActionFightPointsVariationMessage) ID() uint16 {
	return 1030
}

func (m *GameActionFightPointsVariationMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.Delta); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightPointsVariationMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	ldelta, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Delta = ldelta

	return nil
}

type GameActionFightDeathMessage struct {
	AbstractGameActionMessage

	TargetId float64
}

func (m *GameActionFightDeathMessage) ID() uint16 {
	return 1099
}

func (m *GameActionFightDeathMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightDeathMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type SpellListMessage struct {
	SpellPrevisualization bool

	Spells []SpellItem
}

func (m *SpellListMessage) ID() uint16 {
	return 1200
}

func (m *SpellListMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.SpellPrevisualization); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Spells))); err != nil {
		return err
	}

	for i := range m.Spells {

		if err := m.Spells[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *SpellListMessage) Deserialize(r Reader) error {

	lspellPrevisualization, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.SpellPrevisualization = lspellPrevisualization

	lspellsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Spells = make([]SpellItem, lspellsLen)

	for i := range m.Spells {

		var lspells SpellItem

		lspells.Deserialize(r)

		m.Spells[i] = lspells

	}

	return nil
}

type StartupActionsListMessage struct {
	Actions []StartupActionAddObject
}

func (m *StartupActionsListMessage) ID() uint16 {
	return 1301
}

func (m *StartupActionsListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Actions))); err != nil {
		return err
	}

	for i := range m.Actions {

		if err := m.Actions[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *StartupActionsListMessage) Deserialize(r Reader) error {

	lactionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Actions = make([]StartupActionAddObject, lactionsLen)

	for i := range m.Actions {

		var lactions StartupActionAddObject

		lactions.Deserialize(r)

		m.Actions[i] = lactions

	}

	return nil
}

type StartupActionsExecuteMessage struct {
}

func (m *StartupActionsExecuteMessage) ID() uint16 {
	return 1302
}

func (m *StartupActionsExecuteMessage) Serialize(w Writer) error {

	return nil
}

func (m *StartupActionsExecuteMessage) Deserialize(r Reader) error {

	return nil
}

type StartupActionsObjetAttributionMessage struct {
	ActionId uint32

	CharacterId int64
}

func (m *StartupActionsObjetAttributionMessage) ID() uint16 {
	return 1303
}

func (m *StartupActionsObjetAttributionMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.ActionId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.CharacterId); err != nil {
		return err
	}

	return nil
}

func (m *StartupActionsObjetAttributionMessage) Deserialize(r Reader) error {

	lactionId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.ActionId = lactionId

	lcharacterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.CharacterId = lcharacterId

	return nil
}

type StartupActionFinishedMessage struct {
	Success bool

	ActionId uint32

	AutomaticAction bool
}

func (m *StartupActionFinishedMessage) ID() uint16 {
	return 1304
}

func (m *StartupActionFinishedMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Success)

	setWrappedFlag(bbw0, 1, m.AutomaticAction)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.ActionId); err != nil {
		return err
	}

	return nil
}

func (m *StartupActionFinishedMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Success = getWrappedFlag(bbw0, 0)

	m.AutomaticAction = getWrappedFlag(bbw0, 1)

	lactionId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.ActionId = lactionId

	return nil
}

type GetPartsListMessage struct {
}

func (m *GetPartsListMessage) ID() uint16 {
	return 1501
}

func (m *GetPartsListMessage) Serialize(w Writer) error {

	return nil
}

func (m *GetPartsListMessage) Deserialize(r Reader) error {

	return nil
}

type PartsListMessage struct {
	Parts []ContentPart
}

func (m *PartsListMessage) ID() uint16 {
	return 1502
}

func (m *PartsListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Parts))); err != nil {
		return err
	}

	for i := range m.Parts {

		if err := m.Parts[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PartsListMessage) Deserialize(r Reader) error {

	lpartsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Parts = make([]ContentPart, lpartsLen)

	for i := range m.Parts {

		var lparts ContentPart

		lparts.Deserialize(r)

		m.Parts[i] = lparts

	}

	return nil
}

type DownloadPartMessage struct {
	Id string
}

func (m *DownloadPartMessage) ID() uint16 {
	return 1503
}

func (m *DownloadPartMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *DownloadPartMessage) Deserialize(r Reader) error {

	lid, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type GetPartInfoMessage struct {
	Id string
}

func (m *GetPartInfoMessage) ID() uint16 {
	return 1506
}

func (m *GetPartInfoMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *GetPartInfoMessage) Deserialize(r Reader) error {

	lid, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type PartInfoMessage struct {
	Part ContentPart

	InstallationPercent float32
}

func (m *PartInfoMessage) ID() uint16 {
	return 1508
}

func (m *PartInfoMessage) Serialize(w Writer) error {

	if err := m.Part.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteFloat(m.InstallationPercent); err != nil {
		return err
	}

	return nil
}

func (m *PartInfoMessage) Deserialize(r Reader) error {

	var lpart ContentPart

	lpart.Deserialize(r)

	m.Part = lpart

	linstallationPercent, err := r.ReadFloat()
	if err != nil {
		return err
	}

	m.InstallationPercent = linstallationPercent

	return nil
}

type DownloadGetCurrentSpeedRequestMessage struct {
}

func (m *DownloadGetCurrentSpeedRequestMessage) ID() uint16 {
	return 1510
}

func (m *DownloadGetCurrentSpeedRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *DownloadGetCurrentSpeedRequestMessage) Deserialize(r Reader) error {

	return nil
}

type DownloadCurrentSpeedMessage struct {
	DownloadSpeed uint8
}

func (m *DownloadCurrentSpeedMessage) ID() uint16 {
	return 1511
}

func (m *DownloadCurrentSpeedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.DownloadSpeed); err != nil {
		return err
	}

	return nil
}

func (m *DownloadCurrentSpeedMessage) Deserialize(r Reader) error {

	ldownloadSpeed, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.DownloadSpeed = ldownloadSpeed

	return nil
}

type DownloadSetSpeedRequestMessage struct {
	DownloadSpeed uint8
}

func (m *DownloadSetSpeedRequestMessage) ID() uint16 {
	return 1512
}

func (m *DownloadSetSpeedRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.DownloadSpeed); err != nil {
		return err
	}

	return nil
}

func (m *DownloadSetSpeedRequestMessage) Deserialize(r Reader) error {

	ldownloadSpeed, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.DownloadSpeed = ldownloadSpeed

	return nil
}

type DownloadErrorMessage struct {
	ErrorId uint8

	Message string

	HelpUrl string
}

func (m *DownloadErrorMessage) ID() uint16 {
	return 1513
}

func (m *DownloadErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.ErrorId); err != nil {
		return err
	}

	if err := w.WriteString(m.Message); err != nil {
		return err
	}

	if err := w.WriteString(m.HelpUrl); err != nil {
		return err
	}

	return nil
}

func (m *DownloadErrorMessage) Deserialize(r Reader) error {

	lerrorId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ErrorId = lerrorId

	lmessage, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Message = lmessage

	lhelpUrl, err := r.ReadString()
	if err != nil {
		return err
	}

	m.HelpUrl = lhelpUrl

	return nil
}

type ZaapListMessage struct {
	TeleportDestinationsListMessage

	SpawnMapId uint32
}

func (m *ZaapListMessage) ID() uint16 {
	return 1604
}

func (m *ZaapListMessage) Serialize(w Writer) error {

	if err := m.TeleportDestinationsListMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.SpawnMapId); err != nil {
		return err
	}

	return nil
}

func (m *ZaapListMessage) Deserialize(r Reader) error {

	if err := m.TeleportDestinationsListMessage.Deserialize(r); err != nil {
		return err
	}

	lspawnMapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.SpawnMapId = lspawnMapId

	return nil
}

type SetEnablePVPRequestMessage struct {
	Enable bool
}

func (m *SetEnablePVPRequestMessage) ID() uint16 {
	return 1810
}

func (m *SetEnablePVPRequestMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *SetEnablePVPRequestMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type DebugHighlightCellsMessage struct {
	Color int32

	Cells []uint16
}

func (m *DebugHighlightCellsMessage) ID() uint16 {
	return 2001
}

func (m *DebugHighlightCellsMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.Color); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Cells))); err != nil {
		return err
	}

	for i := range m.Cells {

		if err := w.WriteVarUInt16(m.Cells[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *DebugHighlightCellsMessage) Deserialize(r Reader) error {

	lcolor, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Color = lcolor

	lcellsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Cells = make([]uint16, lcellsLen)

	for i := range m.Cells {

		lcells, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Cells[i] = lcells

	}

	return nil
}

type DebugClearHighlightCellsMessage struct {
}

func (m *DebugClearHighlightCellsMessage) ID() uint16 {
	return 2002
}

func (m *DebugClearHighlightCellsMessage) Serialize(w Writer) error {

	return nil
}

func (m *DebugClearHighlightCellsMessage) Deserialize(r Reader) error {

	return nil
}

type ObjectUseOnCharacterMessage struct {
	ObjectUseMessage

	CharacterId int64
}

func (m *ObjectUseOnCharacterMessage) ID() uint16 {
	return 3003
}

func (m *ObjectUseOnCharacterMessage) Serialize(w Writer) error {

	if err := m.ObjectUseMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.CharacterId); err != nil {
		return err
	}

	return nil
}

func (m *ObjectUseOnCharacterMessage) Deserialize(r Reader) error {

	if err := m.ObjectUseMessage.Deserialize(r); err != nil {
		return err
	}

	lcharacterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.CharacterId = lcharacterId

	return nil
}

type ObjectErrorMessage struct {
	Reason int8
}

func (m *ObjectErrorMessage) ID() uint16 {
	return 3004
}

func (m *ObjectErrorMessage) Serialize(w Writer) error {

	if err := w.WriteInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *ObjectErrorMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type ObjectDropMessage struct {
	ObjectUID uint32

	Quantity uint32
}

func (m *ObjectDropMessage) ID() uint16 {
	return 3005
}

func (m *ObjectDropMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ObjectDropMessage) Deserialize(r Reader) error {

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

type InventoryWeightMessage struct {
	Weight uint32

	WeightMax uint32
}

func (m *InventoryWeightMessage) ID() uint16 {
	return 3009
}

func (m *InventoryWeightMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.Weight); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.WeightMax); err != nil {
		return err
	}

	return nil
}

func (m *InventoryWeightMessage) Deserialize(r Reader) error {

	lweight, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Weight = lweight

	lweightMax, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.WeightMax = lweightMax

	return nil
}

type ObjectMovementMessage struct {
	ObjectUID uint32

	Position uint8
}

func (m *ObjectMovementMessage) ID() uint16 {
	return 3010
}

func (m *ObjectMovementMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Position); err != nil {
		return err
	}

	return nil
}

func (m *ObjectMovementMessage) Deserialize(r Reader) error {

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	lposition, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Position = lposition

	return nil
}

type ObjectUseOnCellMessage struct {
	ObjectUseMessage

	Cells uint16
}

func (m *ObjectUseOnCellMessage) ID() uint16 {
	return 3013
}

func (m *ObjectUseOnCellMessage) Serialize(w Writer) error {

	if err := m.ObjectUseMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Cells); err != nil {
		return err
	}

	return nil
}

func (m *ObjectUseOnCellMessage) Deserialize(r Reader) error {

	if err := m.ObjectUseMessage.Deserialize(r); err != nil {
		return err
	}

	lcells, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Cells = lcells

	return nil
}

type ObjectGroundRemovedMessage struct {
	Cell uint16
}

func (m *ObjectGroundRemovedMessage) ID() uint16 {
	return 3014
}

func (m *ObjectGroundRemovedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Cell); err != nil {
		return err
	}

	return nil
}

func (m *ObjectGroundRemovedMessage) Deserialize(r Reader) error {

	lcell, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Cell = lcell

	return nil
}

type InventoryContentMessage struct {
	Objects []ObjectItem

	Kamas int64
}

func (m *InventoryContentMessage) ID() uint16 {
	return 3016
}

func (m *InventoryContentMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Objects))); err != nil {
		return err
	}

	for i := range m.Objects {

		if err := m.Objects[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteVarInt64(m.Kamas); err != nil {
		return err
	}

	return nil
}

func (m *InventoryContentMessage) Deserialize(r Reader) error {

	lobjectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Objects = make([]ObjectItem, lobjectsLen)

	for i := range m.Objects {

		var lobjects ObjectItem

		lobjects.Deserialize(r)

		m.Objects[i] = lobjects

	}

	lkamas, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Kamas = lkamas

	return nil
}

type ObjectGroundAddedMessage struct {
	CellId uint16

	ObjectGID uint16
}

func (m *ObjectGroundAddedMessage) ID() uint16 {
	return 3017
}

func (m *ObjectGroundAddedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGID); err != nil {
		return err
	}

	return nil
}

func (m *ObjectGroundAddedMessage) Deserialize(r Reader) error {

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

type ObjectUseMessage struct {
	ObjectUID uint32
}

func (m *ObjectUseMessage) ID() uint16 {
	return 3019
}

func (m *ObjectUseMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	return nil
}

func (m *ObjectUseMessage) Deserialize(r Reader) error {

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	return nil
}

type ObjectSetPositionMessage struct {
	ObjectUID uint32

	Position uint8

	Quantity uint32
}

func (m *ObjectSetPositionMessage) ID() uint16 {
	return 3021
}

func (m *ObjectSetPositionMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Position); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ObjectSetPositionMessage) Deserialize(r Reader) error {

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	lposition, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Position = lposition

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type ObjectDeleteMessage struct {
	ObjectUID uint32

	Quantity uint32
}

func (m *ObjectDeleteMessage) ID() uint16 {
	return 3022
}

func (m *ObjectDeleteMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ObjectDeleteMessage) Deserialize(r Reader) error {

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

type ObjectQuantityMessage struct {
	ObjectUID uint32

	Quantity uint32
}

func (m *ObjectQuantityMessage) ID() uint16 {
	return 3023
}

func (m *ObjectQuantityMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ObjectQuantityMessage) Deserialize(r Reader) error {

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

type ObjectDeletedMessage struct {
	ObjectUID uint32
}

func (m *ObjectDeletedMessage) ID() uint16 {
	return 3024
}

func (m *ObjectDeletedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	return nil
}

func (m *ObjectDeletedMessage) Deserialize(r Reader) error {

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	return nil
}

type ObjectAddedMessage struct {
	Object ObjectItem
}

func (m *ObjectAddedMessage) ID() uint16 {
	return 3025
}

func (m *ObjectAddedMessage) Serialize(w Writer) error {

	if err := m.Object.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ObjectAddedMessage) Deserialize(r Reader) error {

	var lobject ObjectItem

	lobject.Deserialize(r)

	m.Object = lobject

	return nil
}

type ObjectModifiedMessage struct {
	Object ObjectItem
}

func (m *ObjectModifiedMessage) ID() uint16 {
	return 3029
}

func (m *ObjectModifiedMessage) Serialize(w Writer) error {

	if err := m.Object.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ObjectModifiedMessage) Deserialize(r Reader) error {

	var lobject ObjectItem

	lobject.Deserialize(r)

	m.Object = lobject

	return nil
}

type FriendsGetListMessage struct {
}

func (m *FriendsGetListMessage) ID() uint16 {
	return 4001
}

func (m *FriendsGetListMessage) Serialize(w Writer) error {

	return nil
}

func (m *FriendsGetListMessage) Deserialize(r Reader) error {

	return nil
}

type FriendsListMessage struct {
	FriendsList []FriendInformationsIntrf
}

func (m *FriendsListMessage) ID() uint16 {
	return 4002
}

func (m *FriendsListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.FriendsList))); err != nil {
		return err
	}

	for i := range m.FriendsList {

		if err := w.WriteUInt16(m.FriendsList[i].ID()); err != nil {
			return err
		}

		if err := m.FriendsList[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *FriendsListMessage) Deserialize(r Reader) error {

	lfriendsListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FriendsList = make([]FriendInformationsIntrf, lfriendsListLen)

	for i := range m.FriendsList {

		typefriendsListID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lfriendsList, err := GetType(typefriendsListID)
		if err != nil {
			return err
		}

		lfriendsList.Deserialize(r)

		m.FriendsList[i] = lfriendsList.(FriendInformationsIntrf)

	}

	return nil
}

type FriendAddRequestMessage struct {
	Name string
}

func (m *FriendAddRequestMessage) ID() uint16 {
	return 4004
}

func (m *FriendAddRequestMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FriendAddRequestMessage) Deserialize(r Reader) error {

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	return nil
}

type InteractiveUseRequestMessage struct {
	ElemId uint32

	SkillInstanceUid uint32
}

func (m *InteractiveUseRequestMessage) ID() uint16 {
	return 5001
}

func (m *InteractiveUseRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ElemId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.SkillInstanceUid); err != nil {
		return err
	}

	return nil
}

func (m *InteractiveUseRequestMessage) Deserialize(r Reader) error {

	lelemId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ElemId = lelemId

	lskillInstanceUid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SkillInstanceUid = lskillInstanceUid

	return nil
}

type InteractiveMapUpdateMessage struct {
	InteractiveElements []InteractiveElementIntrf
}

func (m *InteractiveMapUpdateMessage) ID() uint16 {
	return 5002
}

func (m *InteractiveMapUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.InteractiveElements))); err != nil {
		return err
	}

	for i := range m.InteractiveElements {

		if err := w.WriteUInt16(m.InteractiveElements[i].ID()); err != nil {
			return err
		}

		if err := m.InteractiveElements[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *InteractiveMapUpdateMessage) Deserialize(r Reader) error {

	linteractiveElementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.InteractiveElements = make([]InteractiveElementIntrf, linteractiveElementsLen)

	for i := range m.InteractiveElements {

		typeinteractiveElementsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		linteractiveElements, err := GetType(typeinteractiveElementsID)
		if err != nil {
			return err
		}

		linteractiveElements.Deserialize(r)

		m.InteractiveElements[i] = linteractiveElements.(InteractiveElementIntrf)

	}

	return nil
}

type LeaveDialogRequestMessage struct {
}

func (m *LeaveDialogRequestMessage) ID() uint16 {
	return 5501
}

func (m *LeaveDialogRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *LeaveDialogRequestMessage) Deserialize(r Reader) error {

	return nil
}

type LeaveDialogMessage struct {
	DialogType uint8
}

func (m *LeaveDialogMessage) ID() uint16 {
	return 5502
}

func (m *LeaveDialogMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.DialogType); err != nil {
		return err
	}

	return nil
}

func (m *LeaveDialogMessage) Deserialize(r Reader) error {

	ldialogType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.DialogType = ldialogType

	return nil
}

type SetUpdateMessage struct {
	SetId uint16

	SetObjects []uint16

	SetEffects []ObjectEffectIntrf
}

func (m *SetUpdateMessage) ID() uint16 {
	return 5503
}

func (m *SetUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SetId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.SetObjects))); err != nil {
		return err
	}

	for i := range m.SetObjects {

		if err := w.WriteVarUInt16(m.SetObjects[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.SetEffects))); err != nil {
		return err
	}

	for i := range m.SetEffects {

		if err := w.WriteUInt16(m.SetEffects[i].ID()); err != nil {
			return err
		}

		if err := m.SetEffects[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *SetUpdateMessage) Deserialize(r Reader) error {

	lsetId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SetId = lsetId

	lsetObjectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SetObjects = make([]uint16, lsetObjectsLen)

	for i := range m.SetObjects {

		lsetObjects, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.SetObjects[i] = lsetObjects

	}

	lsetEffectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SetEffects = make([]ObjectEffectIntrf, lsetEffectsLen)

	for i := range m.SetEffects {

		typesetEffectsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lsetEffects, err := GetType(typesetEffectsID)
		if err != nil {
			return err
		}

		lsetEffects.Deserialize(r)

		m.SetEffects[i] = lsetEffects.(ObjectEffectIntrf)

	}

	return nil
}

type ExchangeRequestMessage struct {
	ExchangeType int8
}

func (m *ExchangeRequestMessage) ID() uint16 {
	return 5505
}

func (m *ExchangeRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt8(m.ExchangeType); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeRequestMessage) Deserialize(r Reader) error {

	lexchangeType, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.ExchangeType = lexchangeType

	return nil
}

type ExchangeAcceptMessage struct {
}

func (m *ExchangeAcceptMessage) ID() uint16 {
	return 5508
}

func (m *ExchangeAcceptMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeAcceptMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeIsReadyMessage struct {
	Id float64

	Ready bool
}

func (m *ExchangeIsReadyMessage) ID() uint16 {
	return 5509
}

func (m *ExchangeIsReadyMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Ready); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeIsReadyMessage) Deserialize(r Reader) error {

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	lready, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Ready = lready

	return nil
}

type ExchangeReadyMessage struct {
	Ready bool

	Step uint16
}

func (m *ExchangeReadyMessage) ID() uint16 {
	return 5511
}

func (m *ExchangeReadyMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Ready); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Step); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeReadyMessage) Deserialize(r Reader) error {

	lready, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Ready = lready

	lstep, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Step = lstep

	return nil
}

type ExchangeStartedMessage struct {
	ExchangeType int8
}

func (m *ExchangeStartedMessage) ID() uint16 {
	return 5512
}

func (m *ExchangeStartedMessage) Serialize(w Writer) error {

	if err := w.WriteInt8(m.ExchangeType); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeStartedMessage) Deserialize(r Reader) error {

	lexchangeType, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.ExchangeType = lexchangeType

	return nil
}

type ExchangeErrorMessage struct {
	ErrorType int8
}

func (m *ExchangeErrorMessage) ID() uint16 {
	return 5513
}

func (m *ExchangeErrorMessage) Serialize(w Writer) error {

	if err := w.WriteInt8(m.ErrorType); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeErrorMessage) Deserialize(r Reader) error {

	lerrorType, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.ErrorType = lerrorType

	return nil
}

type ExchangeObjectMovePricedMessage struct {
	ExchangeObjectMoveMessage

	Price int64
}

func (m *ExchangeObjectMovePricedMessage) ID() uint16 {
	return 5514
}

func (m *ExchangeObjectMovePricedMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMoveMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Price); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectMovePricedMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMoveMessage.Deserialize(r); err != nil {
		return err
	}

	lprice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Price = lprice

	return nil
}

type ExchangeObjectMessage struct {
	Remote bool
}

func (m *ExchangeObjectMessage) ID() uint16 {
	return 5515
}

func (m *ExchangeObjectMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Remote); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectMessage) Deserialize(r Reader) error {

	lremote, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Remote = lremote

	return nil
}

type ExchangeObjectAddedMessage struct {
	ExchangeObjectMessage

	Object ObjectItem
}

func (m *ExchangeObjectAddedMessage) ID() uint16 {
	return 5516
}

func (m *ExchangeObjectAddedMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.Object.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectAddedMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMessage.Deserialize(r); err != nil {
		return err
	}

	var lobject ObjectItem

	lobject.Deserialize(r)

	m.Object = lobject

	return nil
}

type ExchangeObjectRemovedMessage struct {
	ExchangeObjectMessage

	ObjectUID uint32
}

func (m *ExchangeObjectRemovedMessage) ID() uint16 {
	return 5517
}

func (m *ExchangeObjectRemovedMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectRemovedMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMessage.Deserialize(r); err != nil {
		return err
	}

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	return nil
}

type ExchangeObjectMoveMessage struct {
	ObjectUID uint32

	Quantity int32
}

func (m *ExchangeObjectMoveMessage) ID() uint16 {
	return 5518
}

func (m *ExchangeObjectMoveMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteVarInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectMoveMessage) Deserialize(r Reader) error {

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	lquantity, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type ExchangeObjectModifiedMessage struct {
	ExchangeObjectMessage

	Object ObjectItem
}

func (m *ExchangeObjectModifiedMessage) ID() uint16 {
	return 5519
}

func (m *ExchangeObjectModifiedMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.Object.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectModifiedMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMessage.Deserialize(r); err != nil {
		return err
	}

	var lobject ObjectItem

	lobject.Deserialize(r)

	m.Object = lobject

	return nil
}

type ExchangeObjectMoveKamaMessage struct {
	Quantity int64
}

func (m *ExchangeObjectMoveKamaMessage) ID() uint16 {
	return 5520
}

func (m *ExchangeObjectMoveKamaMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectMoveKamaMessage) Deserialize(r Reader) error {

	lquantity, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type ExchangeKamaModifiedMessage struct {
	ExchangeObjectMessage

	Quantity int64
}

func (m *ExchangeKamaModifiedMessage) ID() uint16 {
	return 5521
}

func (m *ExchangeKamaModifiedMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeKamaModifiedMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMessage.Deserialize(r); err != nil {
		return err
	}

	lquantity, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type ExchangeRequestedMessage struct {
	ExchangeType int8
}

func (m *ExchangeRequestedMessage) ID() uint16 {
	return 5522
}

func (m *ExchangeRequestedMessage) Serialize(w Writer) error {

	if err := w.WriteInt8(m.ExchangeType); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeRequestedMessage) Deserialize(r Reader) error {

	lexchangeType, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.ExchangeType = lexchangeType

	return nil
}

type ExchangeRequestedTradeMessage struct {
	ExchangeRequestedMessage

	Source int64

	Target int64
}

func (m *ExchangeRequestedTradeMessage) ID() uint16 {
	return 5523
}

func (m *ExchangeRequestedTradeMessage) Serialize(w Writer) error {

	if err := m.ExchangeRequestedMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Source); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Target); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeRequestedTradeMessage) Deserialize(r Reader) error {

	if err := m.ExchangeRequestedMessage.Deserialize(r); err != nil {
		return err
	}

	lsource, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Source = lsource

	ltarget, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Target = ltarget

	return nil
}

type GameActionFightSlideMessage struct {
	AbstractGameActionMessage

	TargetId float64

	StartCellId int16

	EndCellId int16
}

func (m *GameActionFightSlideMessage) ID() uint16 {
	return 5525
}

func (m *GameActionFightSlideMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.StartCellId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.EndCellId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightSlideMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lstartCellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.StartCellId = lstartCellId

	lendCellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.EndCellId = lendCellId

	return nil
}

type GameActionFightReduceDamagesMessage struct {
	AbstractGameActionMessage

	TargetId float64

	Amount uint32
}

func (m *GameActionFightReduceDamagesMessage) ID() uint16 {
	return 5526
}

func (m *GameActionFightReduceDamagesMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightReduceDamagesMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lamount, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Amount = lamount

	return nil
}

type GameActionFightExchangePositionsMessage struct {
	AbstractGameActionMessage

	TargetId float64

	CasterCellId int16

	TargetCellId int16
}

func (m *GameActionFightExchangePositionsMessage) ID() uint16 {
	return 5527
}

func (m *GameActionFightExchangePositionsMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.CasterCellId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.TargetCellId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightExchangePositionsMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lcasterCellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CasterCellId = lcasterCellId

	ltargetCellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.TargetCellId = ltargetCellId

	return nil
}

type GameActionFightTeleportOnSameMapMessage struct {
	AbstractGameActionMessage

	TargetId float64

	CellId int16
}

func (m *GameActionFightTeleportOnSameMapMessage) ID() uint16 {
	return 5528
}

func (m *GameActionFightTeleportOnSameMapMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightTeleportOnSameMapMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lcellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type GameActionFightReflectDamagesMessage struct {
	AbstractGameActionMessage

	TargetId float64
}

func (m *GameActionFightReflectDamagesMessage) ID() uint16 {
	return 5530
}

func (m *GameActionFightReflectDamagesMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightReflectDamagesMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type GameActionFightReflectSpellMessage struct {
	AbstractGameActionMessage

	TargetId float64
}

func (m *GameActionFightReflectSpellMessage) ID() uint16 {
	return 5531
}

func (m *GameActionFightReflectSpellMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightReflectSpellMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type GameActionFightChangeLookMessage struct {
	AbstractGameActionMessage

	TargetId float64

	EntityLook EntityLook
}

func (m *GameActionFightChangeLookMessage) ID() uint16 {
	return 5532
}

func (m *GameActionFightChangeLookMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := m.EntityLook.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightChangeLookMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	var lentityLook EntityLook

	lentityLook.Deserialize(r)

	m.EntityLook = lentityLook

	return nil
}

type GameActionFightDispellMessage struct {
	AbstractGameActionMessage

	TargetId float64
}

func (m *GameActionFightDispellMessage) ID() uint16 {
	return 5533
}

func (m *GameActionFightDispellMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightDispellMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type GameActionFightStealKamaMessage struct {
	AbstractGameActionMessage

	TargetId float64

	Amount int64
}

func (m *GameActionFightStealKamaMessage) ID() uint16 {
	return 5535
}

func (m *GameActionFightStealKamaMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightStealKamaMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lamount, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Amount = lamount

	return nil
}

type KamasUpdateMessage struct {
	KamasTotal int64
}

func (m *KamasUpdateMessage) ID() uint16 {
	return 5537
}

func (m *KamasUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.KamasTotal); err != nil {
		return err
	}

	return nil
}

func (m *KamasUpdateMessage) Deserialize(r Reader) error {

	lkamasTotal, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.KamasTotal = lkamasTotal

	return nil
}

type GameActionFightMarkCellsMessage struct {
	AbstractGameActionMessage

	Mark GameActionMark
}

func (m *GameActionFightMarkCellsMessage) ID() uint16 {
	return 5540
}

func (m *GameActionFightMarkCellsMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.Mark.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightMarkCellsMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	var lmark GameActionMark

	lmark.Deserialize(r)

	m.Mark = lmark

	return nil
}

type SubscriptionLimitationMessage struct {
	Reason uint8
}

func (m *SubscriptionLimitationMessage) ID() uint16 {
	return 5542
}

func (m *SubscriptionLimitationMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionLimitationMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type CharacterNameSuggestionSuccessMessage struct {
	Suggestion string
}

func (m *CharacterNameSuggestionSuccessMessage) ID() uint16 {
	return 5544
}

func (m *CharacterNameSuggestionSuccessMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Suggestion); err != nil {
		return err
	}

	return nil
}

func (m *CharacterNameSuggestionSuccessMessage) Deserialize(r Reader) error {

	lsuggestion, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Suggestion = lsuggestion

	return nil
}

type CharactersListErrorMessage struct {
}

func (m *CharactersListErrorMessage) ID() uint16 {
	return 5545
}

func (m *CharactersListErrorMessage) Serialize(w Writer) error {

	return nil
}

func (m *CharactersListErrorMessage) Deserialize(r Reader) error {

	return nil
}

type GuildCreationValidMessage struct {
	GuildName string

	GuildEmblem GuildEmblem
}

func (m *GuildCreationValidMessage) ID() uint16 {
	return 5546
}

func (m *GuildCreationValidMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.GuildName); err != nil {
		return err
	}

	if err := m.GuildEmblem.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildCreationValidMessage) Deserialize(r Reader) error {

	lguildName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.GuildName = lguildName

	var lguildEmblem GuildEmblem

	lguildEmblem.Deserialize(r)

	m.GuildEmblem = lguildEmblem

	return nil
}

type GuildInvitationStateRecrutedMessage struct {
	InvitationState uint8
}

func (m *GuildInvitationStateRecrutedMessage) ID() uint16 {
	return 5548
}

func (m *GuildInvitationStateRecrutedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.InvitationState); err != nil {
		return err
	}

	return nil
}

func (m *GuildInvitationStateRecrutedMessage) Deserialize(r Reader) error {

	linvitationState, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.InvitationState = linvitationState

	return nil
}

type GuildChangeMemberParametersMessage struct {
	MemberId int64

	Rank uint16

	ExperienceGivenPercent uint8

	Rights uint32
}

func (m *GuildChangeMemberParametersMessage) ID() uint16 {
	return 5549
}

func (m *GuildChangeMemberParametersMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.MemberId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Rank); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.ExperienceGivenPercent); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Rights); err != nil {
		return err
	}

	return nil
}

func (m *GuildChangeMemberParametersMessage) Deserialize(r Reader) error {

	lmemberId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.MemberId = lmemberId

	lrank, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Rank = lrank

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

	return nil
}

type GuildGetInformationsMessage struct {
	InfoType uint8
}

func (m *GuildGetInformationsMessage) ID() uint16 {
	return 5550
}

func (m *GuildGetInformationsMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.InfoType); err != nil {
		return err
	}

	return nil
}

func (m *GuildGetInformationsMessage) Deserialize(r Reader) error {

	linfoType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.InfoType = linfoType

	return nil
}

type GuildInvitationMessage struct {
	TargetId int64
}

func (m *GuildInvitationMessage) ID() uint16 {
	return 5551
}

func (m *GuildInvitationMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *GuildInvitationMessage) Deserialize(r Reader) error {

	ltargetId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type GuildInvitedMessage struct {
	RecruterId int64

	RecruterName string

	GuildInfo BasicGuildInformations
}

func (m *GuildInvitedMessage) ID() uint16 {
	return 5552
}

func (m *GuildInvitedMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.RecruterId); err != nil {
		return err
	}

	if err := w.WriteString(m.RecruterName); err != nil {
		return err
	}

	if err := m.GuildInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildInvitedMessage) Deserialize(r Reader) error {

	lrecruterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.RecruterId = lrecruterId

	lrecruterName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.RecruterName = lrecruterName

	var lguildInfo BasicGuildInformations

	lguildInfo.Deserialize(r)

	m.GuildInfo = lguildInfo

	return nil
}

type GuildCreationResultMessage struct {
	Result uint8
}

func (m *GuildCreationResultMessage) ID() uint16 {
	return 5554
}

func (m *GuildCreationResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Result); err != nil {
		return err
	}

	return nil
}

func (m *GuildCreationResultMessage) Deserialize(r Reader) error {

	lresult, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Result = lresult

	return nil
}

type GuildInvitationAnswerMessage struct {
	Accept bool
}

func (m *GuildInvitationAnswerMessage) ID() uint16 {
	return 5556
}

func (m *GuildInvitationAnswerMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Accept); err != nil {
		return err
	}

	return nil
}

func (m *GuildInvitationAnswerMessage) Deserialize(r Reader) error {

	laccept, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Accept = laccept

	return nil
}

type GuildInformationsGeneralMessage struct {
	AbandonnedPaddock bool

	Level uint8

	ExpLevelFloor int64

	Experience int64

	ExpNextLevelFloor int64

	CreationDate uint32

	NbTotalMembers uint16

	NbConnectedMembers uint16
}

func (m *GuildInformationsGeneralMessage) ID() uint16 {
	return 5557
}

func (m *GuildInformationsGeneralMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.AbandonnedPaddock); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExpLevelFloor); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Experience); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExpNextLevelFloor); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.CreationDate); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NbTotalMembers); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NbConnectedMembers); err != nil {
		return err
	}

	return nil
}

func (m *GuildInformationsGeneralMessage) Deserialize(r Reader) error {

	labandonnedPaddock, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.AbandonnedPaddock = labandonnedPaddock

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	lexpLevelFloor, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExpLevelFloor = lexpLevelFloor

	lexperience, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Experience = lexperience

	lexpNextLevelFloor, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExpNextLevelFloor = lexpNextLevelFloor

	lcreationDate, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.CreationDate = lcreationDate

	lnbTotalMembers, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NbTotalMembers = lnbTotalMembers

	lnbConnectedMembers, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NbConnectedMembers = lnbConnectedMembers

	return nil
}

type GuildInformationsMembersMessage struct {
	Members []GuildMember
}

func (m *GuildInformationsMembersMessage) ID() uint16 {
	return 5558
}

func (m *GuildInformationsMembersMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Members))); err != nil {
		return err
	}

	for i := range m.Members {

		if err := m.Members[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GuildInformationsMembersMessage) Deserialize(r Reader) error {

	lmembersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Members = make([]GuildMember, lmembersLen)

	for i := range m.Members {

		var lmembers GuildMember

		lmembers.Deserialize(r)

		m.Members[i] = lmembers

	}

	return nil
}

type GuildLeftMessage struct {
}

func (m *GuildLeftMessage) ID() uint16 {
	return 5562
}

func (m *GuildLeftMessage) Serialize(w Writer) error {

	return nil
}

func (m *GuildLeftMessage) Deserialize(r Reader) error {

	return nil
}

type GuildInvitationStateRecruterMessage struct {
	RecrutedName string

	InvitationState uint8
}

func (m *GuildInvitationStateRecruterMessage) ID() uint16 {
	return 5563
}

func (m *GuildInvitationStateRecruterMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.RecrutedName); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.InvitationState); err != nil {
		return err
	}

	return nil
}

func (m *GuildInvitationStateRecruterMessage) Deserialize(r Reader) error {

	lrecrutedName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.RecrutedName = lrecrutedName

	linvitationState, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.InvitationState = linvitationState

	return nil
}

type GuildJoinedMessage struct {
	GuildInfo GuildInformations

	MemberRights uint32
}

func (m *GuildJoinedMessage) ID() uint16 {
	return 5564
}

func (m *GuildJoinedMessage) Serialize(w Writer) error {

	if err := m.GuildInfo.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MemberRights); err != nil {
		return err
	}

	return nil
}

func (m *GuildJoinedMessage) Deserialize(r Reader) error {

	var lguildInfo GuildInformations

	lguildInfo.Deserialize(r)

	m.GuildInfo = lguildInfo

	lmemberRights, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MemberRights = lmemberRights

	return nil
}

type GameActionFightUnmarkCellsMessage struct {
	AbstractGameActionMessage

	MarkId int16
}

func (m *GameActionFightUnmarkCellsMessage) ID() uint16 {
	return 5570
}

func (m *GameActionFightUnmarkCellsMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.MarkId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightUnmarkCellsMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	lmarkId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MarkId = lmarkId

	return nil
}

type GameActionFightKillMessage struct {
	AbstractGameActionMessage

	TargetId float64
}

func (m *GameActionFightKillMessage) ID() uint16 {
	return 5571
}

func (m *GameActionFightKillMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightKillMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type GameFightUpdateTeamMessage struct {
	FightId uint16

	Team FightTeamInformations
}

func (m *GameFightUpdateTeamMessage) ID() uint16 {
	return 5572
}

func (m *GameFightUpdateTeamMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.FightId); err != nil {
		return err
	}

	if err := m.Team.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameFightUpdateTeamMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	var lteam FightTeamInformations

	lteam.Deserialize(r)

	m.Team = lteam

	return nil
}

type SubscriptionZoneMessage struct {
	Active bool
}

func (m *SubscriptionZoneMessage) ID() uint16 {
	return 5573
}

func (m *SubscriptionZoneMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Active); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionZoneMessage) Deserialize(r Reader) error {

	lactive, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Active = lactive

	return nil
}

type PartyStopFollowRequestMessage struct {
	AbstractPartyMessage

	PlayerId int64
}

func (m *PartyStopFollowRequestMessage) ID() uint16 {
	return 5574
}

func (m *PartyStopFollowRequestMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *PartyStopFollowRequestMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type PartyUpdateMessage struct {
	AbstractPartyEventMessage

	MemberInformations PartyMemberInformationsIntrf
}

func (m *PartyUpdateMessage) ID() uint16 {
	return 5575
}

func (m *PartyUpdateMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyEventMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.MemberInformations.ID()); err != nil {
		return err
	}

	if err := m.MemberInformations.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyUpdateMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyEventMessage.Deserialize(r); err != nil {
		return err
	}

	typememberInformationsID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lmemberInformations, err := GetType(typememberInformationsID)
	if err != nil {
		return err
	}

	lmemberInformations.Deserialize(r)

	m.MemberInformations = lmemberInformations.(PartyMemberInformationsIntrf)

	return nil
}

type PartyJoinMessage struct {
	AbstractPartyMessage

	PartyType uint8

	PartyLeaderId int64

	MaxParticipants uint8

	Members []PartyMemberInformationsIntrf

	Guests []PartyGuestInformations

	Restricted bool

	PartyName string
}

func (m *PartyJoinMessage) ID() uint16 {
	return 5576
}

func (m *PartyJoinMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.PartyType); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PartyLeaderId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.MaxParticipants); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Members))); err != nil {
		return err
	}

	for i := range m.Members {

		if err := w.WriteUInt16(m.Members[i].ID()); err != nil {
			return err
		}

		if err := m.Members[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Guests))); err != nil {
		return err
	}

	for i := range m.Guests {

		if err := m.Guests[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteBoolean(m.Restricted); err != nil {
		return err
	}

	if err := w.WriteString(m.PartyName); err != nil {
		return err
	}

	return nil
}

func (m *PartyJoinMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lpartyType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PartyType = lpartyType

	lpartyLeaderId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PartyLeaderId = lpartyLeaderId

	lmaxParticipants, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MaxParticipants = lmaxParticipants

	lmembersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Members = make([]PartyMemberInformationsIntrf, lmembersLen)

	for i := range m.Members {

		typemembersID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lmembers, err := GetType(typemembersID)
		if err != nil {
			return err
		}

		lmembers.Deserialize(r)

		m.Members[i] = lmembers.(PartyMemberInformationsIntrf)

	}

	lguestsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Guests = make([]PartyGuestInformations, lguestsLen)

	for i := range m.Guests {

		var lguests PartyGuestInformations

		lguests.Deserialize(r)

		m.Guests[i] = lguests

	}

	lrestricted, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Restricted = lrestricted

	lpartyName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PartyName = lpartyName

	return nil
}

type PartyFollowMemberRequestMessage struct {
	AbstractPartyMessage

	PlayerId int64
}

func (m *PartyFollowMemberRequestMessage) ID() uint16 {
	return 5577
}

func (m *PartyFollowMemberRequestMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *PartyFollowMemberRequestMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type PartyLeaderUpdateMessage struct {
	AbstractPartyEventMessage

	PartyLeaderId int64
}

func (m *PartyLeaderUpdateMessage) ID() uint16 {
	return 5578
}

func (m *PartyLeaderUpdateMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyEventMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PartyLeaderId); err != nil {
		return err
	}

	return nil
}

func (m *PartyLeaderUpdateMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyEventMessage.Deserialize(r); err != nil {
		return err
	}

	lpartyLeaderId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PartyLeaderId = lpartyLeaderId

	return nil
}

type PartyMemberRemoveMessage struct {
	AbstractPartyEventMessage

	LeavingPlayerId int64
}

func (m *PartyMemberRemoveMessage) ID() uint16 {
	return 5579
}

func (m *PartyMemberRemoveMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyEventMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.LeavingPlayerId); err != nil {
		return err
	}

	return nil
}

func (m *PartyMemberRemoveMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyEventMessage.Deserialize(r); err != nil {
		return err
	}

	lleavingPlayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.LeavingPlayerId = lleavingPlayerId

	return nil
}

type PartyAcceptInvitationMessage struct {
	AbstractPartyMessage
}

func (m *PartyAcceptInvitationMessage) ID() uint16 {
	return 5580
}

func (m *PartyAcceptInvitationMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyAcceptInvitationMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type PartyFollowStatusUpdateMessage struct {
	AbstractPartyMessage

	Success bool

	IsFollowed bool

	FollowedId int64
}

func (m *PartyFollowStatusUpdateMessage) ID() uint16 {
	return 5581
}

func (m *PartyFollowStatusUpdateMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Success)

	setWrappedFlag(bbw0, 1, m.IsFollowed)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.FollowedId); err != nil {
		return err
	}

	return nil
}

func (m *PartyFollowStatusUpdateMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Success = getWrappedFlag(bbw0, 0)

	m.IsFollowed = getWrappedFlag(bbw0, 1)

	lfollowedId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.FollowedId = lfollowedId

	return nil
}

type PartyRefuseInvitationMessage struct {
	AbstractPartyMessage
}

func (m *PartyRefuseInvitationMessage) ID() uint16 {
	return 5582
}

func (m *PartyRefuseInvitationMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyRefuseInvitationMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type PartyCannotJoinErrorMessage struct {
	AbstractPartyMessage

	Reason uint8
}

func (m *PartyCannotJoinErrorMessage) ID() uint16 {
	return 5583
}

func (m *PartyCannotJoinErrorMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *PartyCannotJoinErrorMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type CompassResetMessage struct {
	Type uint8
}

func (m *CompassResetMessage) ID() uint16 {
	return 5584
}

func (m *CompassResetMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CompassResetMessage) Deserialize(r Reader) error {

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	return nil
}

type PartyInvitationRequestMessage struct {
	Name string
}

func (m *PartyInvitationRequestMessage) ID() uint16 {
	return 5585
}

func (m *PartyInvitationRequestMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PartyInvitationRequestMessage) Deserialize(r Reader) error {

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	return nil
}

type PartyInvitationMessage struct {
	AbstractPartyMessage

	PartyType uint8

	PartyName string

	MaxParticipants uint8

	FromId int64

	FromName string

	ToId int64
}

func (m *PartyInvitationMessage) ID() uint16 {
	return 5586
}

func (m *PartyInvitationMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.PartyType); err != nil {
		return err
	}

	if err := w.WriteString(m.PartyName); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.MaxParticipants); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.FromId); err != nil {
		return err
	}

	if err := w.WriteString(m.FromName); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ToId); err != nil {
		return err
	}

	return nil
}

func (m *PartyInvitationMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lpartyType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PartyType = lpartyType

	lpartyName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PartyName = lpartyName

	lmaxParticipants, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MaxParticipants = lmaxParticipants

	lfromId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.FromId = lfromId

	lfromName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.FromName = lfromName

	ltoId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ToId = ltoId

	return nil
}

type PartyLocateMembersRequestMessage struct {
	AbstractPartyMessage
}

func (m *PartyLocateMembersRequestMessage) ID() uint16 {
	return 5587
}

func (m *PartyLocateMembersRequestMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyLocateMembersRequestMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type PartyFollowThisMemberRequestMessage struct {
	PartyFollowMemberRequestMessage

	Enabled bool
}

func (m *PartyFollowThisMemberRequestMessage) ID() uint16 {
	return 5588
}

func (m *PartyFollowThisMemberRequestMessage) Serialize(w Writer) error {

	if err := m.PartyFollowMemberRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *PartyFollowThisMemberRequestMessage) Deserialize(r Reader) error {

	if err := m.PartyFollowMemberRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lenabled, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enabled = lenabled

	return nil
}

type CompassUpdatePartyMemberMessage struct {
	CompassUpdateMessage

	MemberId int64

	Active bool
}

func (m *CompassUpdatePartyMemberMessage) ID() uint16 {
	return 5589
}

func (m *CompassUpdatePartyMemberMessage) Serialize(w Writer) error {

	if err := m.CompassUpdateMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.MemberId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Active); err != nil {
		return err
	}

	return nil
}

func (m *CompassUpdatePartyMemberMessage) Deserialize(r Reader) error {

	if err := m.CompassUpdateMessage.Deserialize(r); err != nil {
		return err
	}

	lmemberId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.MemberId = lmemberId

	lactive, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Active = lactive

	return nil
}

type PartyKickedByMessage struct {
	AbstractPartyMessage

	KickerId int64
}

func (m *PartyKickedByMessage) ID() uint16 {
	return 5590
}

func (m *PartyKickedByMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.KickerId); err != nil {
		return err
	}

	return nil
}

func (m *PartyKickedByMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lkickerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.KickerId = lkickerId

	return nil
}

type CompassUpdateMessage struct {
	Type uint8

	Coords MapCoordinatesIntrf
}

func (m *CompassUpdateMessage) ID() uint16 {
	return 5591
}

func (m *CompassUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Coords.ID()); err != nil {
		return err
	}

	if err := m.Coords.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *CompassUpdateMessage) Deserialize(r Reader) error {

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	typecoordsID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lcoords, err := GetType(typecoordsID)
	if err != nil {
		return err
	}

	lcoords.Deserialize(r)

	m.Coords = lcoords.(MapCoordinatesIntrf)

	return nil
}

type PartyKickRequestMessage struct {
	AbstractPartyMessage

	PlayerId int64
}

func (m *PartyKickRequestMessage) ID() uint16 {
	return 5592
}

func (m *PartyKickRequestMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *PartyKickRequestMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type PartyLeaveRequestMessage struct {
	AbstractPartyMessage
}

func (m *PartyLeaveRequestMessage) ID() uint16 {
	return 5593
}

func (m *PartyLeaveRequestMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyLeaveRequestMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type PartyLeaveMessage struct {
	AbstractPartyMessage
}

func (m *PartyLeaveMessage) ID() uint16 {
	return 5594
}

func (m *PartyLeaveMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyLeaveMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type PartyLocateMembersMessage struct {
	AbstractPartyMessage

	Geopositions []PartyMemberGeoPosition
}

func (m *PartyLocateMembersMessage) ID() uint16 {
	return 5595
}

func (m *PartyLocateMembersMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Geopositions))); err != nil {
		return err
	}

	for i := range m.Geopositions {

		if err := m.Geopositions[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PartyLocateMembersMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lgeopositionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Geopositions = make([]PartyMemberGeoPosition, lgeopositionsLen)

	for i := range m.Geopositions {

		var lgeopositions PartyMemberGeoPosition

		lgeopositions.Deserialize(r)

		m.Geopositions[i] = lgeopositions

	}

	return nil
}

type PartyRefuseInvitationNotificationMessage struct {
	AbstractPartyEventMessage

	GuestId int64
}

func (m *PartyRefuseInvitationNotificationMessage) ID() uint16 {
	return 5596
}

func (m *PartyRefuseInvitationNotificationMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyEventMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.GuestId); err != nil {
		return err
	}

	return nil
}

func (m *PartyRefuseInvitationNotificationMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyEventMessage.Deserialize(r); err != nil {
		return err
	}

	lguestId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.GuestId = lguestId

	return nil
}

type GuildInformationsMemberUpdateMessage struct {
	Member GuildMember
}

func (m *GuildInformationsMemberUpdateMessage) ID() uint16 {
	return 5597
}

func (m *GuildInformationsMemberUpdateMessage) Serialize(w Writer) error {

	if err := m.Member.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildInformationsMemberUpdateMessage) Deserialize(r Reader) error {

	var lmember GuildMember

	lmember.Deserialize(r)

	m.Member = lmember

	return nil
}

type FriendAddedMessage struct {
	FriendAdded FriendInformationsIntrf
}

func (m *FriendAddedMessage) ID() uint16 {
	return 5599
}

func (m *FriendAddedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.FriendAdded.ID()); err != nil {
		return err
	}

	if err := m.FriendAdded.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *FriendAddedMessage) Deserialize(r Reader) error {

	typefriendAddedID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lfriendAdded, err := GetType(typefriendAddedID)
	if err != nil {
		return err
	}

	lfriendAdded.Deserialize(r)

	m.FriendAdded = lfriendAdded.(FriendInformationsIntrf)

	return nil
}

type FriendAddFailureMessage struct {
	Reason uint8
}

func (m *FriendAddFailureMessage) ID() uint16 {
	return 5600
}

func (m *FriendAddFailureMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *FriendAddFailureMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type FriendDeleteResultMessage struct {
	Success bool

	Name string
}

func (m *FriendDeleteResultMessage) ID() uint16 {
	return 5601
}

func (m *FriendDeleteResultMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Success); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FriendDeleteResultMessage) Deserialize(r Reader) error {

	lsuccess, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Success = lsuccess

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	return nil
}

type FriendSetWarnOnConnectionMessage struct {
	Enable bool
}

func (m *FriendSetWarnOnConnectionMessage) ID() uint16 {
	return 5602
}

func (m *FriendSetWarnOnConnectionMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *FriendSetWarnOnConnectionMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type FriendDeleteRequestMessage struct {
	AccountId uint32
}

func (m *FriendDeleteRequestMessage) ID() uint16 {
	return 5603
}

func (m *FriendDeleteRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	return nil
}

func (m *FriendDeleteRequestMessage) Deserialize(r Reader) error {

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	return nil
}

type FriendSpouseJoinRequestMessage struct {
}

func (m *FriendSpouseJoinRequestMessage) ID() uint16 {
	return 5604
}

func (m *FriendSpouseJoinRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *FriendSpouseJoinRequestMessage) Deserialize(r Reader) error {

	return nil
}

type FriendJoinRequestMessage struct {
	Name string
}

func (m *FriendJoinRequestMessage) ID() uint16 {
	return 5605
}

func (m *FriendJoinRequestMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FriendJoinRequestMessage) Deserialize(r Reader) error {

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	return nil
}

type FriendSpouseFollowWithCompassRequestMessage struct {
	Enable bool
}

func (m *FriendSpouseFollowWithCompassRequestMessage) ID() uint16 {
	return 5606
}

func (m *FriendSpouseFollowWithCompassRequestMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *FriendSpouseFollowWithCompassRequestMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type ClientKeyMessage struct {
	Key string
}

func (m *ClientKeyMessage) ID() uint16 {
	return 5607
}

func (m *ClientKeyMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Key); err != nil {
		return err
	}

	return nil
}

func (m *ClientKeyMessage) Deserialize(r Reader) error {

	lkey, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Key = lkey

	return nil
}

type StatsUpgradeResultMessage struct {
	Result int8

	NbCharacBoost uint16
}

func (m *StatsUpgradeResultMessage) ID() uint16 {
	return 5609
}

func (m *StatsUpgradeResultMessage) Serialize(w Writer) error {

	if err := w.WriteInt8(m.Result); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NbCharacBoost); err != nil {
		return err
	}

	return nil
}

func (m *StatsUpgradeResultMessage) Deserialize(r Reader) error {

	lresult, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Result = lresult

	lnbCharacBoost, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NbCharacBoost = lnbCharacBoost

	return nil
}

type StatsUpgradeRequestMessage struct {
	UseAdditionnal bool

	StatId uint8

	BoostPoint uint16
}

func (m *StatsUpgradeRequestMessage) ID() uint16 {
	return 5610
}

func (m *StatsUpgradeRequestMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.UseAdditionnal); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.StatId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.BoostPoint); err != nil {
		return err
	}

	return nil
}

func (m *StatsUpgradeRequestMessage) Deserialize(r Reader) error {

	luseAdditionnal, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.UseAdditionnal = luseAdditionnal

	lstatId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.StatId = lstatId

	lboostPoint, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.BoostPoint = lboostPoint

	return nil
}

type ShowCellRequestMessage struct {
	CellId uint16
}

func (m *ShowCellRequestMessage) ID() uint16 {
	return 5611
}

func (m *ShowCellRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *ShowCellRequestMessage) Deserialize(r Reader) error {

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type ShowCellMessage struct {
	SourceId float64

	CellId uint16
}

func (m *ShowCellMessage) ID() uint16 {
	return 5612
}

func (m *ShowCellMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.SourceId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *ShowCellMessage) Deserialize(r Reader) error {

	lsourceId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.SourceId = lsourceId

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type ChallengeTargetsListMessage struct {
	TargetIds []float64

	TargetCells []int16
}

func (m *ChallengeTargetsListMessage) ID() uint16 {
	return 5613
}

func (m *ChallengeTargetsListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.TargetIds))); err != nil {
		return err
	}

	for i := range m.TargetIds {

		if err := w.WriteDouble(m.TargetIds[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.TargetCells))); err != nil {
		return err
	}

	for i := range m.TargetCells {

		if err := w.WriteInt16(m.TargetCells[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ChallengeTargetsListMessage) Deserialize(r Reader) error {

	ltargetIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.TargetIds = make([]float64, ltargetIdsLen)

	for i := range m.TargetIds {

		ltargetIds, err := r.ReadDouble()
		if err != nil {
			return err
		}

		m.TargetIds[i] = ltargetIds

	}

	ltargetCellsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.TargetCells = make([]int16, ltargetCellsLen)

	for i := range m.TargetCells {

		ltargetCells, err := r.ReadInt16()
		if err != nil {
			return err
		}

		m.TargetCells[i] = ltargetCells

	}

	return nil
}

type ChallengeTargetsListRequestMessage struct {
	ChallengeId uint16
}

func (m *ChallengeTargetsListRequestMessage) ID() uint16 {
	return 5614
}

func (m *ChallengeTargetsListRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ChallengeId); err != nil {
		return err
	}

	return nil
}

func (m *ChallengeTargetsListRequestMessage) Deserialize(r Reader) error {

	lchallengeId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ChallengeId = lchallengeId

	return nil
}

type TaxCollectorDialogQuestionExtendedMessage struct {
	TaxCollectorDialogQuestionBasicMessage

	MaxPods uint16

	Prospecting uint16

	Wisdom uint16

	TaxCollectorsCount uint8

	TaxCollectorAttack int32

	Kamas int64

	Experience int64

	Pods uint32

	ItemsValue int64
}

func (m *TaxCollectorDialogQuestionExtendedMessage) ID() uint16 {
	return 5615
}

func (m *TaxCollectorDialogQuestionExtendedMessage) Serialize(w Writer) error {

	if err := m.TaxCollectorDialogQuestionBasicMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.MaxPods); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Prospecting); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Wisdom); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.TaxCollectorsCount); err != nil {
		return err
	}

	if err := w.WriteInt32(m.TaxCollectorAttack); err != nil {
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

func (m *TaxCollectorDialogQuestionExtendedMessage) Deserialize(r Reader) error {

	if err := m.TaxCollectorDialogQuestionBasicMessage.Deserialize(r); err != nil {
		return err
	}

	lmaxPods, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MaxPods = lmaxPods

	lprospecting, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Prospecting = lprospecting

	lwisdom, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Wisdom = lwisdom

	ltaxCollectorsCount, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TaxCollectorsCount = ltaxCollectorsCount

	ltaxCollectorAttack, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.TaxCollectorAttack = ltaxCollectorAttack

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

type NpcDialogReplyMessage struct {
	ReplyId uint32
}

func (m *NpcDialogReplyMessage) ID() uint16 {
	return 5616
}

func (m *NpcDialogReplyMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ReplyId); err != nil {
		return err
	}

	return nil
}

func (m *NpcDialogReplyMessage) Deserialize(r Reader) error {

	lreplyId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ReplyId = lreplyId

	return nil
}

type NpcDialogQuestionMessage struct {
	MessageId uint32

	DialogParams []string

	VisibleReplies []uint32
}

func (m *NpcDialogQuestionMessage) ID() uint16 {
	return 5617
}

func (m *NpcDialogQuestionMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.MessageId); err != nil {
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

	if err := w.WriteInt16(int16(len(m.VisibleReplies))); err != nil {
		return err
	}

	for i := range m.VisibleReplies {

		if err := w.WriteVarUInt32(m.VisibleReplies[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *NpcDialogQuestionMessage) Deserialize(r Reader) error {

	lmessageId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MessageId = lmessageId

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

	lvisibleRepliesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.VisibleReplies = make([]uint32, lvisibleRepliesLen)

	for i := range m.VisibleReplies {

		lvisibleReplies, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.VisibleReplies[i] = lvisibleReplies

	}

	return nil
}

type NpcDialogCreationMessage struct {
	MapId int32

	NpcId int32
}

func (m *NpcDialogCreationMessage) ID() uint16 {
	return 5618
}

func (m *NpcDialogCreationMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteInt32(m.NpcId); err != nil {
		return err
	}

	return nil
}

func (m *NpcDialogCreationMessage) Deserialize(r Reader) error {

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lnpcId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.NpcId = lnpcId

	return nil
}

type TaxCollectorDialogQuestionBasicMessage struct {
	GuildInfo BasicGuildInformations
}

func (m *TaxCollectorDialogQuestionBasicMessage) ID() uint16 {
	return 5619
}

func (m *TaxCollectorDialogQuestionBasicMessage) Serialize(w Writer) error {

	if err := m.GuildInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorDialogQuestionBasicMessage) Deserialize(r Reader) error {

	var lguildInfo BasicGuildInformations

	lguildInfo.Deserialize(r)

	m.GuildInfo = lguildInfo

	return nil
}

type QuestStepInfoRequestMessage struct {
	QuestId uint16
}

func (m *QuestStepInfoRequestMessage) ID() uint16 {
	return 5622
}

func (m *QuestStepInfoRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.QuestId); err != nil {
		return err
	}

	return nil
}

func (m *QuestStepInfoRequestMessage) Deserialize(r Reader) error {

	lquestId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.QuestId = lquestId

	return nil
}

type QuestListRequestMessage struct {
}

func (m *QuestListRequestMessage) ID() uint16 {
	return 5623
}

func (m *QuestListRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *QuestListRequestMessage) Deserialize(r Reader) error {

	return nil
}

type QuestStepInfoMessage struct {
	Infos QuestActiveInformationsIntrf
}

func (m *QuestStepInfoMessage) ID() uint16 {
	return 5625
}

func (m *QuestStepInfoMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Infos.ID()); err != nil {
		return err
	}

	if err := m.Infos.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *QuestStepInfoMessage) Deserialize(r Reader) error {

	typeinfosID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	linfos, err := GetType(typeinfosID)
	if err != nil {
		return err
	}

	linfos.Deserialize(r)

	m.Infos = linfos.(QuestActiveInformationsIntrf)

	return nil
}

type QuestListMessage struct {
	FinishedQuestsIds []uint16

	FinishedQuestsCounts []uint16

	ActiveQuests []QuestActiveInformationsIntrf

	ReinitDoneQuestsIds []uint16
}

func (m *QuestListMessage) ID() uint16 {
	return 5626
}

func (m *QuestListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.FinishedQuestsIds))); err != nil {
		return err
	}

	for i := range m.FinishedQuestsIds {

		if err := w.WriteVarUInt16(m.FinishedQuestsIds[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.FinishedQuestsCounts))); err != nil {
		return err
	}

	for i := range m.FinishedQuestsCounts {

		if err := w.WriteVarUInt16(m.FinishedQuestsCounts[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.ActiveQuests))); err != nil {
		return err
	}

	for i := range m.ActiveQuests {

		if err := w.WriteUInt16(m.ActiveQuests[i].ID()); err != nil {
			return err
		}

		if err := m.ActiveQuests[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.ReinitDoneQuestsIds))); err != nil {
		return err
	}

	for i := range m.ReinitDoneQuestsIds {

		if err := w.WriteVarUInt16(m.ReinitDoneQuestsIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *QuestListMessage) Deserialize(r Reader) error {

	lfinishedQuestsIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FinishedQuestsIds = make([]uint16, lfinishedQuestsIdsLen)

	for i := range m.FinishedQuestsIds {

		lfinishedQuestsIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.FinishedQuestsIds[i] = lfinishedQuestsIds

	}

	lfinishedQuestsCountsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FinishedQuestsCounts = make([]uint16, lfinishedQuestsCountsLen)

	for i := range m.FinishedQuestsCounts {

		lfinishedQuestsCounts, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.FinishedQuestsCounts[i] = lfinishedQuestsCounts

	}

	lactiveQuestsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ActiveQuests = make([]QuestActiveInformationsIntrf, lactiveQuestsLen)

	for i := range m.ActiveQuests {

		typeactiveQuestsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lactiveQuests, err := GetType(typeactiveQuestsID)
		if err != nil {
			return err
		}

		lactiveQuests.Deserialize(r)

		m.ActiveQuests[i] = lactiveQuests.(QuestActiveInformationsIntrf)

	}

	lreinitDoneQuestsIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ReinitDoneQuestsIds = make([]uint16, lreinitDoneQuestsIdsLen)

	for i := range m.ReinitDoneQuestsIds {

		lreinitDoneQuestsIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.ReinitDoneQuestsIds[i] = lreinitDoneQuestsIds

	}

	return nil
}

type ExchangeLeaveMessage struct {
	LeaveDialogMessage

	Success bool
}

func (m *ExchangeLeaveMessage) ID() uint16 {
	return 5628
}

func (m *ExchangeLeaveMessage) Serialize(w Writer) error {

	if err := m.LeaveDialogMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Success); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeLeaveMessage) Deserialize(r Reader) error {

	if err := m.LeaveDialogMessage.Deserialize(r); err != nil {
		return err
	}

	lsuccess, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Success = lsuccess

	return nil
}

type FriendWarnOnConnectionStateMessage struct {
	Enable bool
}

func (m *FriendWarnOnConnectionStateMessage) ID() uint16 {
	return 5630
}

func (m *FriendWarnOnConnectionStateMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *FriendWarnOnConnectionStateMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type GameRolePlayShowActorMessage struct {
	Informations GameRolePlayActorInformationsIntrf
}

func (m *GameRolePlayShowActorMessage) ID() uint16 {
	return 5632
}

func (m *GameRolePlayShowActorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Informations.ID()); err != nil {
		return err
	}

	if err := m.Informations.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayShowActorMessage) Deserialize(r Reader) error {

	typeinformationsID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	linformations, err := GetType(typeinformationsID)
	if err != nil {
		return err
	}

	linformations.Deserialize(r)

	m.Informations = linformations.(GameRolePlayActorInformationsIntrf)

	return nil
}

type TaxCollectorMovementMessage struct {
	MovementType uint8

	BasicInfos TaxCollectorBasicInformations

	PlayerId int64

	PlayerName string
}

func (m *TaxCollectorMovementMessage) ID() uint16 {
	return 5633
}

func (m *TaxCollectorMovementMessage) Serialize(w Writer) error {

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

func (m *TaxCollectorMovementMessage) Deserialize(r Reader) error {

	lmovementType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MovementType = lmovementType

	var lbasicInfos TaxCollectorBasicInformations

	lbasicInfos.Deserialize(r)

	m.BasicInfos = lbasicInfos

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

type TaxCollectorErrorMessage struct {
	Reason int8
}

func (m *TaxCollectorErrorMessage) ID() uint16 {
	return 5634
}

func (m *TaxCollectorErrorMessage) Serialize(w Writer) error {

	if err := w.WriteInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorErrorMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type TaxCollectorAttackedResultMessage struct {
	DeadOrAlive bool

	BasicInfos TaxCollectorBasicInformations

	Guild BasicGuildInformations
}

func (m *TaxCollectorAttackedResultMessage) ID() uint16 {
	return 5635
}

func (m *TaxCollectorAttackedResultMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.DeadOrAlive); err != nil {
		return err
	}

	if err := m.BasicInfos.Serialize(w); err != nil {
		return err
	}

	if err := m.Guild.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorAttackedResultMessage) Deserialize(r Reader) error {

	ldeadOrAlive, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.DeadOrAlive = ldeadOrAlive

	var lbasicInfos TaxCollectorBasicInformations

	lbasicInfos.Deserialize(r)

	m.BasicInfos = lbasicInfos

	var lguild BasicGuildInformations

	lguild.Deserialize(r)

	m.Guild = lguild

	return nil
}

type GuildInfosUpgradeMessage struct {
	MaxTaxCollectorsCount uint8

	TaxCollectorsCount uint8

	TaxCollectorLifePoints uint16

	TaxCollectorDamagesBonuses uint16

	TaxCollectorPods uint16

	TaxCollectorProspecting uint16

	TaxCollectorWisdom uint16

	BoostPoints uint16

	SpellId []uint16

	SpellLevel []int16
}

func (m *GuildInfosUpgradeMessage) ID() uint16 {
	return 5636
}

func (m *GuildInfosUpgradeMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.MaxTaxCollectorsCount); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.TaxCollectorsCount); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TaxCollectorLifePoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TaxCollectorDamagesBonuses); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TaxCollectorPods); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TaxCollectorProspecting); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TaxCollectorWisdom); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.BoostPoints); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.SpellId))); err != nil {
		return err
	}

	for i := range m.SpellId {

		if err := w.WriteVarUInt16(m.SpellId[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.SpellLevel))); err != nil {
		return err
	}

	for i := range m.SpellLevel {

		if err := w.WriteInt16(m.SpellLevel[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GuildInfosUpgradeMessage) Deserialize(r Reader) error {

	lmaxTaxCollectorsCount, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MaxTaxCollectorsCount = lmaxTaxCollectorsCount

	ltaxCollectorsCount, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TaxCollectorsCount = ltaxCollectorsCount

	ltaxCollectorLifePoints, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TaxCollectorLifePoints = ltaxCollectorLifePoints

	ltaxCollectorDamagesBonuses, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TaxCollectorDamagesBonuses = ltaxCollectorDamagesBonuses

	ltaxCollectorPods, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TaxCollectorPods = ltaxCollectorPods

	ltaxCollectorProspecting, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TaxCollectorProspecting = ltaxCollectorProspecting

	ltaxCollectorWisdom, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TaxCollectorWisdom = ltaxCollectorWisdom

	lboostPoints, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.BoostPoints = lboostPoints

	lspellIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SpellId = make([]uint16, lspellIdLen)

	for i := range m.SpellId {

		lspellId, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.SpellId[i] = lspellId

	}

	lspellLevelLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SpellLevel = make([]int16, lspellLevelLen)

	for i := range m.SpellLevel {

		lspellLevel, err := r.ReadInt16()
		if err != nil {
			return err
		}

		m.SpellLevel[i] = lspellLevel

	}

	return nil
}

type GameContextRefreshEntityLookMessage struct {
	Id float64

	Look EntityLook
}

func (m *GameContextRefreshEntityLookMessage) ID() uint16 {
	return 5637
}

func (m *GameContextRefreshEntityLookMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	if err := m.Look.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameContextRefreshEntityLookMessage) Deserialize(r Reader) error {

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	var llook EntityLook

	llook.Deserialize(r)

	m.Look = llook

	return nil
}

type NicknameRefusedMessage struct {
	Reason uint8
}

func (m *NicknameRefusedMessage) ID() uint16 {
	return 5638
}

func (m *NicknameRefusedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *NicknameRefusedMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type NicknameChoiceRequestMessage struct {
	Nickname string
}

func (m *NicknameChoiceRequestMessage) ID() uint16 {
	return 5639
}

func (m *NicknameChoiceRequestMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Nickname); err != nil {
		return err
	}

	return nil
}

func (m *NicknameChoiceRequestMessage) Deserialize(r Reader) error {

	lnickname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Nickname = lnickname

	return nil
}

type NicknameRegistrationMessage struct {
}

func (m *NicknameRegistrationMessage) ID() uint16 {
	return 5640
}

func (m *NicknameRegistrationMessage) Serialize(w Writer) error {

	return nil
}

func (m *NicknameRegistrationMessage) Deserialize(r Reader) error {

	return nil
}

type NicknameAcceptedMessage struct {
}

func (m *NicknameAcceptedMessage) ID() uint16 {
	return 5641
}

func (m *NicknameAcceptedMessage) Serialize(w Writer) error {

	return nil
}

func (m *NicknameAcceptedMessage) Deserialize(r Reader) error {

	return nil
}

type MapNpcsQuestStatusUpdateMessage struct {
	MapId int32

	NpcsIdsWithQuest []int32

	QuestFlags []GameRolePlayNpcQuestFlag

	NpcsIdsWithoutQuest []int32
}

func (m *MapNpcsQuestStatusUpdateMessage) ID() uint16 {
	return 5642
}

func (m *MapNpcsQuestStatusUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.NpcsIdsWithQuest))); err != nil {
		return err
	}

	for i := range m.NpcsIdsWithQuest {

		if err := w.WriteInt32(m.NpcsIdsWithQuest[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.QuestFlags))); err != nil {
		return err
	}

	for i := range m.QuestFlags {

		if err := m.QuestFlags[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.NpcsIdsWithoutQuest))); err != nil {
		return err
	}

	for i := range m.NpcsIdsWithoutQuest {

		if err := w.WriteInt32(m.NpcsIdsWithoutQuest[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *MapNpcsQuestStatusUpdateMessage) Deserialize(r Reader) error {

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lnpcsIdsWithQuestLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.NpcsIdsWithQuest = make([]int32, lnpcsIdsWithQuestLen)

	for i := range m.NpcsIdsWithQuest {

		lnpcsIdsWithQuest, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.NpcsIdsWithQuest[i] = lnpcsIdsWithQuest

	}

	lquestFlagsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.QuestFlags = make([]GameRolePlayNpcQuestFlag, lquestFlagsLen)

	for i := range m.QuestFlags {

		var lquestFlags GameRolePlayNpcQuestFlag

		lquestFlags.Deserialize(r)

		m.QuestFlags[i] = lquestFlags

	}

	lnpcsIdsWithoutQuestLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.NpcsIdsWithoutQuest = make([]int32, lnpcsIdsWithoutQuestLen)

	for i := range m.NpcsIdsWithoutQuest {

		lnpcsIdsWithoutQuest, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.NpcsIdsWithoutQuest[i] = lnpcsIdsWithoutQuest

	}

	return nil
}

type QuestStartRequestMessage struct {
	QuestId uint16
}

func (m *QuestStartRequestMessage) ID() uint16 {
	return 5643
}

func (m *QuestStartRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.QuestId); err != nil {
		return err
	}

	return nil
}

func (m *QuestStartRequestMessage) Deserialize(r Reader) error {

	lquestId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.QuestId = lquestId

	return nil
}

type EmoteAddMessage struct {
	EmoteId uint8
}

func (m *EmoteAddMessage) ID() uint16 {
	return 5644
}

func (m *EmoteAddMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.EmoteId); err != nil {
		return err
	}

	return nil
}

func (m *EmoteAddMessage) Deserialize(r Reader) error {

	lemoteId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.EmoteId = lemoteId

	return nil
}

type StorageKamasUpdateMessage struct {
	KamasTotal int64
}

func (m *StorageKamasUpdateMessage) ID() uint16 {
	return 5645
}

func (m *StorageKamasUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.KamasTotal); err != nil {
		return err
	}

	return nil
}

func (m *StorageKamasUpdateMessage) Deserialize(r Reader) error {

	lkamasTotal, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.KamasTotal = lkamasTotal

	return nil
}

type StorageInventoryContentMessage struct {
	InventoryContentMessage
}

func (m *StorageInventoryContentMessage) ID() uint16 {
	return 5646
}

func (m *StorageInventoryContentMessage) Serialize(w Writer) error {

	if err := m.InventoryContentMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *StorageInventoryContentMessage) Deserialize(r Reader) error {

	if err := m.InventoryContentMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type StorageObjectUpdateMessage struct {
	Object ObjectItem
}

func (m *StorageObjectUpdateMessage) ID() uint16 {
	return 5647
}

func (m *StorageObjectUpdateMessage) Serialize(w Writer) error {

	if err := m.Object.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *StorageObjectUpdateMessage) Deserialize(r Reader) error {

	var lobject ObjectItem

	lobject.Deserialize(r)

	m.Object = lobject

	return nil
}

type StorageObjectRemoveMessage struct {
	ObjectUID uint32
}

func (m *StorageObjectRemoveMessage) ID() uint16 {
	return 5648
}

func (m *StorageObjectRemoveMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	return nil
}

func (m *StorageObjectRemoveMessage) Deserialize(r Reader) error {

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	return nil
}

type JobCrafterDirectoryDefineSettingsMessage struct {
	Settings JobCrafterDirectorySettings
}

func (m *JobCrafterDirectoryDefineSettingsMessage) ID() uint16 {
	return 5649
}

func (m *JobCrafterDirectoryDefineSettingsMessage) Serialize(w Writer) error {

	if err := m.Settings.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *JobCrafterDirectoryDefineSettingsMessage) Deserialize(r Reader) error {

	var lsettings JobCrafterDirectorySettings

	lsettings.Deserialize(r)

	m.Settings = lsettings

	return nil
}

type JobCrafterDirectoryAddMessage struct {
	ListEntry JobCrafterDirectoryListEntry
}

func (m *JobCrafterDirectoryAddMessage) ID() uint16 {
	return 5651
}

func (m *JobCrafterDirectoryAddMessage) Serialize(w Writer) error {

	if err := m.ListEntry.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *JobCrafterDirectoryAddMessage) Deserialize(r Reader) error {

	var llistEntry JobCrafterDirectoryListEntry

	llistEntry.Deserialize(r)

	m.ListEntry = llistEntry

	return nil
}

type JobCrafterDirectorySettingsMessage struct {
	CraftersSettings []JobCrafterDirectorySettings
}

func (m *JobCrafterDirectorySettingsMessage) ID() uint16 {
	return 5652
}

func (m *JobCrafterDirectorySettingsMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.CraftersSettings))); err != nil {
		return err
	}

	for i := range m.CraftersSettings {

		if err := m.CraftersSettings[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *JobCrafterDirectorySettingsMessage) Deserialize(r Reader) error {

	lcraftersSettingsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CraftersSettings = make([]JobCrafterDirectorySettings, lcraftersSettingsLen)

	for i := range m.CraftersSettings {

		var lcraftersSettings JobCrafterDirectorySettings

		lcraftersSettings.Deserialize(r)

		m.CraftersSettings[i] = lcraftersSettings

	}

	return nil
}

type JobCrafterDirectoryRemoveMessage struct {
	JobId uint8

	PlayerId int64
}

func (m *JobCrafterDirectoryRemoveMessage) ID() uint16 {
	return 5653
}

func (m *JobCrafterDirectoryRemoveMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.JobId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *JobCrafterDirectoryRemoveMessage) Deserialize(r Reader) error {

	ljobId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.JobId = ljobId

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type JobExperienceUpdateMessage struct {
	ExperiencesUpdate JobExperience
}

func (m *JobExperienceUpdateMessage) ID() uint16 {
	return 5654
}

func (m *JobExperienceUpdateMessage) Serialize(w Writer) error {

	if err := m.ExperiencesUpdate.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *JobExperienceUpdateMessage) Deserialize(r Reader) error {

	var lexperiencesUpdate JobExperience

	lexperiencesUpdate.Deserialize(r)

	m.ExperiencesUpdate = lexperiencesUpdate

	return nil
}

type JobDescriptionMessage struct {
	JobsDescription []JobDescription
}

func (m *JobDescriptionMessage) ID() uint16 {
	return 5655
}

func (m *JobDescriptionMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.JobsDescription))); err != nil {
		return err
	}

	for i := range m.JobsDescription {

		if err := m.JobsDescription[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *JobDescriptionMessage) Deserialize(r Reader) error {

	ljobsDescriptionLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.JobsDescription = make([]JobDescription, ljobsDescriptionLen)

	for i := range m.JobsDescription {

		var ljobsDescription JobDescription

		ljobsDescription.Deserialize(r)

		m.JobsDescription[i] = ljobsDescription

	}

	return nil
}

type JobLevelUpMessage struct {
	NewLevel uint8

	JobsDescription JobDescription
}

func (m *JobLevelUpMessage) ID() uint16 {
	return 5656
}

func (m *JobLevelUpMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.NewLevel); err != nil {
		return err
	}

	if err := m.JobsDescription.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *JobLevelUpMessage) Deserialize(r Reader) error {

	lnewLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NewLevel = lnewLevel

	var ljobsDescription JobDescription

	ljobsDescription.Deserialize(r)

	m.JobsDescription = ljobsDescription

	return nil
}

type UpdateLifePointsMessage struct {
	LifePoints uint32

	MaxLifePoints uint32
}

func (m *UpdateLifePointsMessage) ID() uint16 {
	return 5658
}

func (m *UpdateLifePointsMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.LifePoints); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MaxLifePoints); err != nil {
		return err
	}

	return nil
}

func (m *UpdateLifePointsMessage) Deserialize(r Reader) error {

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

	return nil
}

type HouseKickIndoorMerchantRequestMessage struct {
	CellId uint16
}

func (m *HouseKickIndoorMerchantRequestMessage) ID() uint16 {
	return 5661
}

func (m *HouseKickIndoorMerchantRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *HouseKickIndoorMerchantRequestMessage) Deserialize(r Reader) error {

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type AdminQuietCommandMessage struct {
	AdminCommandMessage
}

func (m *AdminQuietCommandMessage) ID() uint16 {
	return 5662
}

func (m *AdminQuietCommandMessage) Serialize(w Writer) error {

	if err := m.AdminCommandMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AdminQuietCommandMessage) Deserialize(r Reader) error {

	if err := m.AdminCommandMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type BasicLatencyStatsMessage struct {
	Latency uint16

	SampleCount uint16

	Max uint16
}

func (m *BasicLatencyStatsMessage) ID() uint16 {
	return 5663
}

func (m *BasicLatencyStatsMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Latency); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SampleCount); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Max); err != nil {
		return err
	}

	return nil
}

func (m *BasicLatencyStatsMessage) Deserialize(r Reader) error {

	llatency, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.Latency = llatency

	lsampleCount, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SampleCount = lsampleCount

	lmax, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Max = lmax

	return nil
}

type BasicWhoAmIRequestMessage struct {
	Verbose bool
}

func (m *BasicWhoAmIRequestMessage) ID() uint16 {
	return 5664
}

func (m *BasicWhoAmIRequestMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Verbose); err != nil {
		return err
	}

	return nil
}

func (m *BasicWhoAmIRequestMessage) Deserialize(r Reader) error {

	lverbose, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Verbose = lverbose

	return nil
}

type LockableChangeCodeMessage struct {
	Code string
}

func (m *LockableChangeCodeMessage) ID() uint16 {
	return 5666
}

func (m *LockableChangeCodeMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Code); err != nil {
		return err
	}

	return nil
}

func (m *LockableChangeCodeMessage) Deserialize(r Reader) error {

	lcode, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Code = lcode

	return nil
}

type LockableUseCodeMessage struct {
	Code string
}

func (m *LockableUseCodeMessage) ID() uint16 {
	return 5667
}

func (m *LockableUseCodeMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Code); err != nil {
		return err
	}

	return nil
}

func (m *LockableUseCodeMessage) Deserialize(r Reader) error {

	lcode, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Code = lcode

	return nil
}

type LockableStateUpdateHouseDoorMessage struct {
	LockableStateUpdateAbstractMessage

	HouseId uint32

	InstanceId uint32

	SecondHand bool
}

func (m *LockableStateUpdateHouseDoorMessage) ID() uint16 {
	return 5668
}

func (m *LockableStateUpdateHouseDoorMessage) Serialize(w Writer) error {

	if err := m.LockableStateUpdateAbstractMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.HouseId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.SecondHand); err != nil {
		return err
	}

	return nil
}

func (m *LockableStateUpdateHouseDoorMessage) Deserialize(r Reader) error {

	if err := m.LockableStateUpdateAbstractMessage.Deserialize(r); err != nil {
		return err
	}

	lhouseId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HouseId = lhouseId

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

	return nil
}

type LockableStateUpdateStorageMessage struct {
	LockableStateUpdateAbstractMessage

	MapId int32

	ElementId uint32
}

func (m *LockableStateUpdateStorageMessage) ID() uint16 {
	return 5669
}

func (m *LockableStateUpdateStorageMessage) Serialize(w Writer) error {

	if err := m.LockableStateUpdateAbstractMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.MapId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ElementId); err != nil {
		return err
	}

	return nil
}

func (m *LockableStateUpdateStorageMessage) Deserialize(r Reader) error {

	if err := m.LockableStateUpdateAbstractMessage.Deserialize(r); err != nil {
		return err
	}

	lmapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	lelementId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ElementId = lelementId

	return nil
}

type CharacterLevelUpMessage struct {
	NewLevel uint8
}

func (m *CharacterLevelUpMessage) ID() uint16 {
	return 5670
}

func (m *CharacterLevelUpMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.NewLevel); err != nil {
		return err
	}

	return nil
}

func (m *CharacterLevelUpMessage) Deserialize(r Reader) error {

	lnewLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NewLevel = lnewLevel

	return nil
}

type LockableStateUpdateAbstractMessage struct {
	Locked bool
}

func (m *LockableStateUpdateAbstractMessage) ID() uint16 {
	return 5671
}

func (m *LockableStateUpdateAbstractMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Locked); err != nil {
		return err
	}

	return nil
}

func (m *LockableStateUpdateAbstractMessage) Deserialize(r Reader) error {

	llocked, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Locked = llocked

	return nil
}

type LockableCodeResultMessage struct {
	Result uint8
}

func (m *LockableCodeResultMessage) ID() uint16 {
	return 5672
}

func (m *LockableCodeResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Result); err != nil {
		return err
	}

	return nil
}

func (m *LockableCodeResultMessage) Deserialize(r Reader) error {

	lresult, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Result = lresult

	return nil
}

type IgnoredAddRequestMessage struct {
	Name string

	Session bool
}

func (m *IgnoredAddRequestMessage) ID() uint16 {
	return 5673
}

func (m *IgnoredAddRequestMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Session); err != nil {
		return err
	}

	return nil
}

func (m *IgnoredAddRequestMessage) Deserialize(r Reader) error {

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	lsession, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Session = lsession

	return nil
}

type IgnoredListMessage struct {
	IgnoredList []IgnoredInformationsIntrf
}

func (m *IgnoredListMessage) ID() uint16 {
	return 5674
}

func (m *IgnoredListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.IgnoredList))); err != nil {
		return err
	}

	for i := range m.IgnoredList {

		if err := w.WriteUInt16(m.IgnoredList[i].ID()); err != nil {
			return err
		}

		if err := m.IgnoredList[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *IgnoredListMessage) Deserialize(r Reader) error {

	lignoredListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.IgnoredList = make([]IgnoredInformationsIntrf, lignoredListLen)

	for i := range m.IgnoredList {

		typeignoredListID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lignoredList, err := GetType(typeignoredListID)
		if err != nil {
			return err
		}

		lignoredList.Deserialize(r)

		m.IgnoredList[i] = lignoredList.(IgnoredInformationsIntrf)

	}

	return nil
}

type DocumentReadingBeginMessage struct {
	DocumentId uint16
}

func (m *DocumentReadingBeginMessage) ID() uint16 {
	return 5675
}

func (m *DocumentReadingBeginMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.DocumentId); err != nil {
		return err
	}

	return nil
}

func (m *DocumentReadingBeginMessage) Deserialize(r Reader) error {

	ldocumentId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DocumentId = ldocumentId

	return nil
}

type IgnoredGetListMessage struct {
}

func (m *IgnoredGetListMessage) ID() uint16 {
	return 5676
}

func (m *IgnoredGetListMessage) Serialize(w Writer) error {

	return nil
}

func (m *IgnoredGetListMessage) Deserialize(r Reader) error {

	return nil
}

type IgnoredDeleteResultMessage struct {
	Success bool

	Name string

	Session bool
}

func (m *IgnoredDeleteResultMessage) ID() uint16 {
	return 5677
}

func (m *IgnoredDeleteResultMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Success)

	setWrappedFlag(bbw0, 1, m.Session)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	return nil
}

func (m *IgnoredDeleteResultMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Success = getWrappedFlag(bbw0, 0)

	m.Session = getWrappedFlag(bbw0, 1)

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	return nil
}

type IgnoredAddedMessage struct {
	IgnoreAdded IgnoredInformationsIntrf

	Session bool
}

func (m *IgnoredAddedMessage) ID() uint16 {
	return 5678
}

func (m *IgnoredAddedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.IgnoreAdded.ID()); err != nil {
		return err
	}

	if err := m.IgnoreAdded.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Session); err != nil {
		return err
	}

	return nil
}

func (m *IgnoredAddedMessage) Deserialize(r Reader) error {

	typeignoreAddedID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lignoreAdded, err := GetType(typeignoreAddedID)
	if err != nil {
		return err
	}

	lignoreAdded.Deserialize(r)

	m.IgnoreAdded = lignoreAdded.(IgnoredInformationsIntrf)

	lsession, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Session = lsession

	return nil
}

type IgnoredAddFailureMessage struct {
	Reason uint8
}

func (m *IgnoredAddFailureMessage) ID() uint16 {
	return 5679
}

func (m *IgnoredAddFailureMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *IgnoredAddFailureMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type IgnoredDeleteRequestMessage struct {
	AccountId uint32

	Session bool
}

func (m *IgnoredDeleteRequestMessage) ID() uint16 {
	return 5680
}

func (m *IgnoredDeleteRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Session); err != nil {
		return err
	}

	return nil
}

func (m *IgnoredDeleteRequestMessage) Deserialize(r Reader) error {

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	lsession, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Session = lsession

	return nil
}

type EmotePlayMessage struct {
	EmotePlayAbstractMessage

	ActorId float64

	AccountId uint32
}

func (m *EmotePlayMessage) ID() uint16 {
	return 5683
}

func (m *EmotePlayMessage) Serialize(w Writer) error {

	if err := m.EmotePlayAbstractMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.ActorId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	return nil
}

func (m *EmotePlayMessage) Deserialize(r Reader) error {

	if err := m.EmotePlayAbstractMessage.Deserialize(r); err != nil {
		return err
	}

	lactorId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.ActorId = lactorId

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	return nil
}

type LifePointsRegenBeginMessage struct {
	RegenRate uint8
}

func (m *LifePointsRegenBeginMessage) ID() uint16 {
	return 5684
}

func (m *LifePointsRegenBeginMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.RegenRate); err != nil {
		return err
	}

	return nil
}

func (m *LifePointsRegenBeginMessage) Deserialize(r Reader) error {

	lregenRate, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.RegenRate = lregenRate

	return nil
}

type EmotePlayRequestMessage struct {
	EmoteId uint8
}

func (m *EmotePlayRequestMessage) ID() uint16 {
	return 5685
}

func (m *EmotePlayRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.EmoteId); err != nil {
		return err
	}

	return nil
}

func (m *EmotePlayRequestMessage) Deserialize(r Reader) error {

	lemoteId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.EmoteId = lemoteId

	return nil
}

type LifePointsRegenEndMessage struct {
	UpdateLifePointsMessage

	LifePointsGained uint32
}

func (m *LifePointsRegenEndMessage) ID() uint16 {
	return 5686
}

func (m *LifePointsRegenEndMessage) Serialize(w Writer) error {

	if err := m.UpdateLifePointsMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.LifePointsGained); err != nil {
		return err
	}

	return nil
}

func (m *LifePointsRegenEndMessage) Deserialize(r Reader) error {

	if err := m.UpdateLifePointsMessage.Deserialize(r); err != nil {
		return err
	}

	llifePointsGained, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.LifePointsGained = llifePointsGained

	return nil
}

type EmoteRemoveMessage struct {
	EmoteId uint8
}

func (m *EmoteRemoveMessage) ID() uint16 {
	return 5687
}

func (m *EmoteRemoveMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.EmoteId); err != nil {
		return err
	}

	return nil
}

func (m *EmoteRemoveMessage) Deserialize(r Reader) error {

	lemoteId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.EmoteId = lemoteId

	return nil
}

type EmotePlayErrorMessage struct {
	EmoteId uint8
}

func (m *EmotePlayErrorMessage) ID() uint16 {
	return 5688
}

func (m *EmotePlayErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.EmoteId); err != nil {
		return err
	}

	return nil
}

func (m *EmotePlayErrorMessage) Deserialize(r Reader) error {

	lemoteId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.EmoteId = lemoteId

	return nil
}

type EmoteListMessage struct {
	EmoteIds []uint8
}

func (m *EmoteListMessage) ID() uint16 {
	return 5689
}

func (m *EmoteListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.EmoteIds))); err != nil {
		return err
	}

	for i := range m.EmoteIds {

		if err := w.WriteUInt8(m.EmoteIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *EmoteListMessage) Deserialize(r Reader) error {

	lemoteIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.EmoteIds = make([]uint8, lemoteIdsLen)

	for i := range m.EmoteIds {

		lemoteIds, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.EmoteIds[i] = lemoteIds

	}

	return nil
}

type EmotePlayAbstractMessage struct {
	EmoteId uint8

	EmoteStartTime float64
}

func (m *EmotePlayAbstractMessage) ID() uint16 {
	return 5690
}

func (m *EmotePlayAbstractMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.EmoteId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.EmoteStartTime); err != nil {
		return err
	}

	return nil
}

func (m *EmotePlayAbstractMessage) Deserialize(r Reader) error {

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

type EmotePlayMassiveMessage struct {
	EmotePlayAbstractMessage

	ActorIds []float64
}

func (m *EmotePlayMassiveMessage) ID() uint16 {
	return 5691
}

func (m *EmotePlayMassiveMessage) Serialize(w Writer) error {

	if err := m.EmotePlayAbstractMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.ActorIds))); err != nil {
		return err
	}

	for i := range m.ActorIds {

		if err := w.WriteDouble(m.ActorIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *EmotePlayMassiveMessage) Deserialize(r Reader) error {

	if err := m.EmotePlayAbstractMessage.Deserialize(r); err != nil {
		return err
	}

	lactorIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ActorIds = make([]float64, lactorIdsLen)

	for i := range m.ActorIds {

		lactorIds, err := r.ReadDouble()
		if err != nil {
			return err
		}

		m.ActorIds[i] = lactorIds

	}

	return nil
}

type GameEntityDispositionMessage struct {
	Disposition IdentifiedEntityDispositionInformations
}

func (m *GameEntityDispositionMessage) ID() uint16 {
	return 5693
}

func (m *GameEntityDispositionMessage) Serialize(w Writer) error {

	if err := m.Disposition.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameEntityDispositionMessage) Deserialize(r Reader) error {

	var ldisposition IdentifiedEntityDispositionInformations

	ldisposition.Deserialize(r)

	m.Disposition = ldisposition

	return nil
}

type GameEntityDispositionErrorMessage struct {
}

func (m *GameEntityDispositionErrorMessage) ID() uint16 {
	return 5695
}

func (m *GameEntityDispositionErrorMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameEntityDispositionErrorMessage) Deserialize(r Reader) error {

	return nil
}

type GameEntitiesDispositionMessage struct {
	Dispositions []IdentifiedEntityDispositionInformations
}

func (m *GameEntitiesDispositionMessage) ID() uint16 {
	return 5696
}

func (m *GameEntitiesDispositionMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Dispositions))); err != nil {
		return err
	}

	for i := range m.Dispositions {

		if err := m.Dispositions[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameEntitiesDispositionMessage) Deserialize(r Reader) error {

	ldispositionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Dispositions = make([]IdentifiedEntityDispositionInformations, ldispositionsLen)

	for i := range m.Dispositions {

		var ldispositions IdentifiedEntityDispositionInformations

		ldispositions.Deserialize(r)

		m.Dispositions[i] = ldispositions

	}

	return nil
}

type HouseSellRequestMessage struct {
	InstanceId uint32

	Amount int64

	ForSale bool
}

func (m *HouseSellRequestMessage) ID() uint16 {
	return 5697
}

func (m *HouseSellRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Amount); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.ForSale); err != nil {
		return err
	}

	return nil
}

func (m *HouseSellRequestMessage) Deserialize(r Reader) error {

	linstanceId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.InstanceId = linstanceId

	lamount, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Amount = lamount

	lforSale, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.ForSale = lforSale

	return nil
}

type HouseKickRequestMessage struct {
	Id int64
}

func (m *HouseKickRequestMessage) ID() uint16 {
	return 5698
}

func (m *HouseKickRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *HouseKickRequestMessage) Deserialize(r Reader) error {

	lid, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type GuildSpellUpgradeRequestMessage struct {
	SpellId uint32
}

func (m *GuildSpellUpgradeRequestMessage) ID() uint16 {
	return 5699
}

func (m *GuildSpellUpgradeRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.SpellId); err != nil {
		return err
	}

	return nil
}

func (m *GuildSpellUpgradeRequestMessage) Deserialize(r Reader) error {

	lspellId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	return nil
}

type HouseGuildRightsViewMessage struct {
	HouseId uint32

	InstanceId uint32
}

func (m *HouseGuildRightsViewMessage) ID() uint16 {
	return 5700
}

func (m *HouseGuildRightsViewMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HouseId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	return nil
}

func (m *HouseGuildRightsViewMessage) Deserialize(r Reader) error {

	lhouseId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HouseId = lhouseId

	linstanceId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.InstanceId = linstanceId

	return nil
}

type HouseGuildNoneMessage struct {
	HouseId uint32

	InstanceId uint32

	SecondHand bool
}

func (m *HouseGuildNoneMessage) ID() uint16 {
	return 5701
}

func (m *HouseGuildNoneMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HouseId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.SecondHand); err != nil {
		return err
	}

	return nil
}

func (m *HouseGuildNoneMessage) Deserialize(r Reader) error {

	lhouseId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HouseId = lhouseId

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

	return nil
}

type HouseGuildRightsMessage struct {
	HouseId uint32

	InstanceId uint32

	SecondHand bool

	GuildInfo GuildInformations

	Rights uint32
}

func (m *HouseGuildRightsMessage) ID() uint16 {
	return 5703
}

func (m *HouseGuildRightsMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HouseId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.SecondHand); err != nil {
		return err
	}

	if err := m.GuildInfo.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Rights); err != nil {
		return err
	}

	return nil
}

func (m *HouseGuildRightsMessage) Deserialize(r Reader) error {

	lhouseId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HouseId = lhouseId

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

	var lguildInfo GuildInformations

	lguildInfo.Deserialize(r)

	m.GuildInfo = lguildInfo

	lrights, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Rights = lrights

	return nil
}

type HouseGuildShareRequestMessage struct {
	HouseId uint32

	InstanceId uint32

	Enable bool

	Rights uint32
}

func (m *HouseGuildShareRequestMessage) ID() uint16 {
	return 5704
}

func (m *HouseGuildShareRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HouseId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Rights); err != nil {
		return err
	}

	return nil
}

func (m *HouseGuildShareRequestMessage) Deserialize(r Reader) error {

	lhouseId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HouseId = lhouseId

	linstanceId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.InstanceId = linstanceId

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	lrights, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Rights = lrights

	return nil
}

type GuildCharacsUpgradeRequestMessage struct {
	CharaTypeTarget uint8
}

func (m *GuildCharacsUpgradeRequestMessage) ID() uint16 {
	return 5706
}

func (m *GuildCharacsUpgradeRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.CharaTypeTarget); err != nil {
		return err
	}

	return nil
}

func (m *GuildCharacsUpgradeRequestMessage) Deserialize(r Reader) error {

	lcharaTypeTarget, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CharaTypeTarget = lcharaTypeTarget

	return nil
}

type InteractiveElementUpdatedMessage struct {
	InteractiveElement InteractiveElement
}

func (m *InteractiveElementUpdatedMessage) ID() uint16 {
	return 5708
}

func (m *InteractiveElementUpdatedMessage) Serialize(w Writer) error {

	if err := m.InteractiveElement.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *InteractiveElementUpdatedMessage) Deserialize(r Reader) error {

	var linteractiveElement InteractiveElement

	linteractiveElement.Deserialize(r)

	m.InteractiveElement = linteractiveElement

	return nil
}

type StatedElementUpdatedMessage struct {
	StatedElement StatedElement
}

func (m *StatedElementUpdatedMessage) ID() uint16 {
	return 5709
}

func (m *StatedElementUpdatedMessage) Serialize(w Writer) error {

	if err := m.StatedElement.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *StatedElementUpdatedMessage) Deserialize(r Reader) error {

	var lstatedElement StatedElement

	lstatedElement.Deserialize(r)

	m.StatedElement = lstatedElement

	return nil
}

type GuildFightLeaveRequestMessage struct {
	TaxCollectorId uint32

	CharacterId int64
}

func (m *GuildFightLeaveRequestMessage) ID() uint16 {
	return 5715
}

func (m *GuildFightLeaveRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.TaxCollectorId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.CharacterId); err != nil {
		return err
	}

	return nil
}

func (m *GuildFightLeaveRequestMessage) Deserialize(r Reader) error {

	ltaxCollectorId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.TaxCollectorId = ltaxCollectorId

	lcharacterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.CharacterId = lcharacterId

	return nil
}

type StatedMapUpdateMessage struct {
	StatedElements []StatedElement
}

func (m *StatedMapUpdateMessage) ID() uint16 {
	return 5716
}

func (m *StatedMapUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.StatedElements))); err != nil {
		return err
	}

	for i := range m.StatedElements {

		if err := m.StatedElements[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *StatedMapUpdateMessage) Deserialize(r Reader) error {

	lstatedElementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.StatedElements = make([]StatedElement, lstatedElementsLen)

	for i := range m.StatedElements {

		var lstatedElements StatedElement

		lstatedElements.Deserialize(r)

		m.StatedElements[i] = lstatedElements

	}

	return nil
}

type GuildFightJoinRequestMessage struct {
	TaxCollectorId int32
}

func (m *GuildFightJoinRequestMessage) ID() uint16 {
	return 5717
}

func (m *GuildFightJoinRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.TaxCollectorId); err != nil {
		return err
	}

	return nil
}

func (m *GuildFightJoinRequestMessage) Deserialize(r Reader) error {

	ltaxCollectorId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.TaxCollectorId = ltaxCollectorId

	return nil
}

type GuildFightPlayersHelpersLeaveMessage struct {
	FightId uint32

	PlayerId int64
}

func (m *GuildFightPlayersHelpersLeaveMessage) ID() uint16 {
	return 5719
}

func (m *GuildFightPlayersHelpersLeaveMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *GuildFightPlayersHelpersLeaveMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type GuildFightPlayersHelpersJoinMessage struct {
	FightId uint32

	PlayerInfo CharacterMinimalPlusLookInformations
}

func (m *GuildFightPlayersHelpersJoinMessage) ID() uint16 {
	return 5720
}

func (m *GuildFightPlayersHelpersJoinMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.FightId); err != nil {
		return err
	}

	if err := m.PlayerInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildFightPlayersHelpersJoinMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	var lplayerInfo CharacterMinimalPlusLookInformations

	lplayerInfo.Deserialize(r)

	m.PlayerInfo = lplayerInfo

	return nil
}

type LivingObjectDissociateMessage struct {
	LivingUID uint32

	LivingPosition uint8
}

func (m *LivingObjectDissociateMessage) ID() uint16 {
	return 5723
}

func (m *LivingObjectDissociateMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.LivingUID); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.LivingPosition); err != nil {
		return err
	}

	return nil
}

func (m *LivingObjectDissociateMessage) Deserialize(r Reader) error {

	llivingUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.LivingUID = llivingUID

	llivingPosition, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.LivingPosition = llivingPosition

	return nil
}

type LivingObjectChangeSkinRequestMessage struct {
	LivingUID uint32

	LivingPosition uint8

	SkinId uint32
}

func (m *LivingObjectChangeSkinRequestMessage) ID() uint16 {
	return 5725
}

func (m *LivingObjectChangeSkinRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.LivingUID); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.LivingPosition); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.SkinId); err != nil {
		return err
	}

	return nil
}

func (m *LivingObjectChangeSkinRequestMessage) Deserialize(r Reader) error {

	llivingUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.LivingUID = llivingUID

	llivingPosition, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.LivingPosition = llivingPosition

	lskinId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SkinId = lskinId

	return nil
}

type OnConnectionEventMessage struct {
	EventType uint8
}

func (m *OnConnectionEventMessage) ID() uint16 {
	return 5726
}

func (m *OnConnectionEventMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.EventType); err != nil {
		return err
	}

	return nil
}

func (m *OnConnectionEventMessage) Deserialize(r Reader) error {

	leventType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.EventType = leventType

	return nil
}

type GameRolePlayPlayerFightRequestMessage struct {
	TargetId int64

	TargetCellId int16

	Friendly bool
}

func (m *GameRolePlayPlayerFightRequestMessage) ID() uint16 {
	return 5731
}

func (m *GameRolePlayPlayerFightRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.TargetCellId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Friendly); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayPlayerFightRequestMessage) Deserialize(r Reader) error {

	ltargetId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	ltargetCellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.TargetCellId = ltargetCellId

	lfriendly, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Friendly = lfriendly

	return nil
}

type GameRolePlayPlayerFightFriendlyAnswerMessage struct {
	FightId int32

	Accept bool
}

func (m *GameRolePlayPlayerFightFriendlyAnswerMessage) ID() uint16 {
	return 5732
}

func (m *GameRolePlayPlayerFightFriendlyAnswerMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Accept); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayPlayerFightFriendlyAnswerMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	laccept, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Accept = laccept

	return nil
}

type GameRolePlayPlayerFightFriendlyAnsweredMessage struct {
	FightId int32

	SourceId int64

	TargetId int64

	Accept bool
}

func (m *GameRolePlayPlayerFightFriendlyAnsweredMessage) ID() uint16 {
	return 5733
}

func (m *GameRolePlayPlayerFightFriendlyAnsweredMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.SourceId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Accept); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayPlayerFightFriendlyAnsweredMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lsourceId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.SourceId = lsourceId

	ltargetId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	laccept, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Accept = laccept

	return nil
}

type HousePropertiesMessage struct {
	HouseId uint32

	DoorsOnMap []uint32

	Properties HouseInstanceInformationsIntrf
}

func (m *HousePropertiesMessage) ID() uint16 {
	return 5734
}

func (m *HousePropertiesMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HouseId); err != nil {
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

	if err := w.WriteUInt16(m.Properties.ID()); err != nil {
		return err
	}

	if err := m.Properties.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *HousePropertiesMessage) Deserialize(r Reader) error {

	lhouseId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HouseId = lhouseId

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

	typepropertiesID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lproperties, err := GetType(typepropertiesID)
	if err != nil {
		return err
	}

	lproperties.Deserialize(r)

	m.Properties = lproperties.(HouseInstanceInformationsIntrf)

	return nil
}

type HouseBuyResultMessage struct {
	HouseId uint32

	InstanceId uint32

	SecondHand bool

	Bought bool

	RealPrice int64
}

func (m *HouseBuyResultMessage) ID() uint16 {
	return 5735
}

func (m *HouseBuyResultMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.SecondHand)

	setWrappedFlag(bbw0, 1, m.Bought)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.HouseId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.RealPrice); err != nil {
		return err
	}

	return nil
}

func (m *HouseBuyResultMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.SecondHand = getWrappedFlag(bbw0, 0)

	m.Bought = getWrappedFlag(bbw0, 1)

	lhouseId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HouseId = lhouseId

	linstanceId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.InstanceId = linstanceId

	lrealPrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.RealPrice = lrealPrice

	return nil
}

type HouseBuyRequestMessage struct {
	ProposedPrice int64
}

func (m *HouseBuyRequestMessage) ID() uint16 {
	return 5738
}

func (m *HouseBuyRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.ProposedPrice); err != nil {
		return err
	}

	return nil
}

func (m *HouseBuyRequestMessage) Deserialize(r Reader) error {

	lproposedPrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ProposedPrice = lproposedPrice

	return nil
}

type PurchasableDialogMessage struct {
	BuyOrSell bool

	PurchasableId uint32

	PurchasableInstanceId uint32

	SecondHand bool

	Price int64
}

func (m *PurchasableDialogMessage) ID() uint16 {
	return 5739
}

func (m *PurchasableDialogMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.BuyOrSell)

	setWrappedFlag(bbw0, 1, m.SecondHand)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.PurchasableId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.PurchasableInstanceId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Price); err != nil {
		return err
	}

	return nil
}

func (m *PurchasableDialogMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BuyOrSell = getWrappedFlag(bbw0, 0)

	m.SecondHand = getWrappedFlag(bbw0, 1)

	lpurchasableId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.PurchasableId = lpurchasableId

	lpurchasableInstanceId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.PurchasableInstanceId = lpurchasableInstanceId

	lprice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Price = lprice

	return nil
}

type LockableShowCodeDialogMessage struct {
	ChangeOrUse bool

	CodeSize uint8
}

func (m *LockableShowCodeDialogMessage) ID() uint16 {
	return 5740
}

func (m *LockableShowCodeDialogMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.ChangeOrUse); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.CodeSize); err != nil {
		return err
	}

	return nil
}

func (m *LockableShowCodeDialogMessage) Deserialize(r Reader) error {

	lchangeOrUse, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.ChangeOrUse = lchangeOrUse

	lcodeSize, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CodeSize = lcodeSize

	return nil
}

type GameActionFightTriggerGlyphTrapMessage struct {
	AbstractGameActionMessage

	MarkId int16

	TriggeringCharacterId float64

	TriggeredSpellId uint16
}

func (m *GameActionFightTriggerGlyphTrapMessage) ID() uint16 {
	return 5741
}

func (m *GameActionFightTriggerGlyphTrapMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.MarkId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TriggeringCharacterId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TriggeredSpellId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightTriggerGlyphTrapMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	lmarkId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MarkId = lmarkId

	ltriggeringCharacterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TriggeringCharacterId = ltriggeringCharacterId

	ltriggeredSpellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TriggeredSpellId = ltriggeredSpellId

	return nil
}

type MapRunningFightListRequestMessage struct {
}

func (m *MapRunningFightListRequestMessage) ID() uint16 {
	return 5742
}

func (m *MapRunningFightListRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *MapRunningFightListRequestMessage) Deserialize(r Reader) error {

	return nil
}

type MapRunningFightListMessage struct {
	Fights []FightExternalInformations
}

func (m *MapRunningFightListMessage) ID() uint16 {
	return 5743
}

func (m *MapRunningFightListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Fights))); err != nil {
		return err
	}

	for i := range m.Fights {

		if err := m.Fights[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *MapRunningFightListMessage) Deserialize(r Reader) error {

	lfightsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Fights = make([]FightExternalInformations, lfightsLen)

	for i := range m.Fights {

		var lfights FightExternalInformations

		lfights.Deserialize(r)

		m.Fights[i] = lfights

	}

	return nil
}

type InteractiveUsedMessage struct {
	EntityId int64

	ElemId uint32

	SkillId uint16

	Duration uint16

	CanMove bool
}

func (m *InteractiveUsedMessage) ID() uint16 {
	return 5745
}

func (m *InteractiveUsedMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.EntityId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ElemId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SkillId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Duration); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.CanMove); err != nil {
		return err
	}

	return nil
}

func (m *InteractiveUsedMessage) Deserialize(r Reader) error {

	lentityId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.EntityId = lentityId

	lelemId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ElemId = lelemId

	lskillId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SkillId = lskillId

	lduration, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Duration = lduration

	lcanMove, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.CanMove = lcanMove

	return nil
}

type JobMultiCraftAvailableSkillsMessage struct {
	JobAllowMultiCraftRequestMessage

	PlayerId int64

	Skills []uint16
}

func (m *JobMultiCraftAvailableSkillsMessage) ID() uint16 {
	return 5747
}

func (m *JobMultiCraftAvailableSkillsMessage) Serialize(w Writer) error {

	if err := m.JobAllowMultiCraftRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Skills))); err != nil {
		return err
	}

	for i := range m.Skills {

		if err := w.WriteVarUInt16(m.Skills[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *JobMultiCraftAvailableSkillsMessage) Deserialize(r Reader) error {

	if err := m.JobAllowMultiCraftRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	lskillsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Skills = make([]uint16, lskillsLen)

	for i := range m.Skills {

		lskills, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Skills[i] = lskills

	}

	return nil
}

type JobAllowMultiCraftRequestMessage struct {
	Enabled bool
}

func (m *JobAllowMultiCraftRequestMessage) ID() uint16 {
	return 5748
}

func (m *JobAllowMultiCraftRequestMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *JobAllowMultiCraftRequestMessage) Deserialize(r Reader) error {

	lenabled, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enabled = lenabled

	return nil
}

type MapRunningFightDetailsRequestMessage struct {
	FightId uint32
}

func (m *MapRunningFightDetailsRequestMessage) ID() uint16 {
	return 5750
}

func (m *MapRunningFightDetailsRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.FightId); err != nil {
		return err
	}

	return nil
}

func (m *MapRunningFightDetailsRequestMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	return nil
}

type MapRunningFightDetailsMessage struct {
	FightId uint32

	Attackers []GameFightFighterLightInformationsIntrf

	Defenders []GameFightFighterLightInformationsIntrf
}

func (m *MapRunningFightDetailsMessage) ID() uint16 {
	return 5751
}

func (m *MapRunningFightDetailsMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Attackers))); err != nil {
		return err
	}

	for i := range m.Attackers {

		if err := w.WriteUInt16(m.Attackers[i].ID()); err != nil {
			return err
		}

		if err := m.Attackers[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Defenders))); err != nil {
		return err
	}

	for i := range m.Defenders {

		if err := w.WriteUInt16(m.Defenders[i].ID()); err != nil {
			return err
		}

		if err := m.Defenders[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *MapRunningFightDetailsMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lattackersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Attackers = make([]GameFightFighterLightInformationsIntrf, lattackersLen)

	for i := range m.Attackers {

		typeattackersID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lattackers, err := GetType(typeattackersID)
		if err != nil {
			return err
		}

		lattackers.Deserialize(r)

		m.Attackers[i] = lattackers.(GameFightFighterLightInformationsIntrf)

	}

	ldefendersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Defenders = make([]GameFightFighterLightInformationsIntrf, ldefendersLen)

	for i := range m.Defenders {

		typedefendersID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		ldefenders, err := GetType(typedefendersID)
		if err != nil {
			return err
		}

		ldefenders.Deserialize(r)

		m.Defenders[i] = ldefenders.(GameFightFighterLightInformationsIntrf)

	}

	return nil
}

type ExchangeTypesItemsExchangerDescriptionForUserMessage struct {
	ItemTypeDescriptions []BidExchangerObjectInfo
}

func (m *ExchangeTypesItemsExchangerDescriptionForUserMessage) ID() uint16 {
	return 5752
}

func (m *ExchangeTypesItemsExchangerDescriptionForUserMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ItemTypeDescriptions))); err != nil {
		return err
	}

	for i := range m.ItemTypeDescriptions {

		if err := m.ItemTypeDescriptions[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeTypesItemsExchangerDescriptionForUserMessage) Deserialize(r Reader) error {

	litemTypeDescriptionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ItemTypeDescriptions = make([]BidExchangerObjectInfo, litemTypeDescriptionsLen)

	for i := range m.ItemTypeDescriptions {

		var litemTypeDescriptions BidExchangerObjectInfo

		litemTypeDescriptions.Deserialize(r)

		m.ItemTypeDescriptions[i] = litemTypeDescriptions

	}

	return nil
}

type ExchangeRequestOnShopStockMessage struct {
}

func (m *ExchangeRequestOnShopStockMessage) ID() uint16 {
	return 5753
}

func (m *ExchangeRequestOnShopStockMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeRequestOnShopStockMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeBidPriceMessage struct {
	GenericId uint16

	AveragePrice int64
}

func (m *ExchangeBidPriceMessage) ID() uint16 {
	return 5755
}

func (m *ExchangeBidPriceMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.GenericId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.AveragePrice); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidPriceMessage) Deserialize(r Reader) error {

	lgenericId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.GenericId = lgenericId

	laveragePrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.AveragePrice = laveragePrice

	return nil
}

type ExchangeBuyOkMessage struct {
}

func (m *ExchangeBuyOkMessage) ID() uint16 {
	return 5759
}

func (m *ExchangeBuyOkMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeBuyOkMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeStartOkNpcShopMessage struct {
	NpcSellerId float64

	TokenId uint16

	ObjectsInfos []ObjectItemToSellInNpcShop
}

func (m *ExchangeStartOkNpcShopMessage) ID() uint16 {
	return 5761
}

func (m *ExchangeStartOkNpcShopMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.NpcSellerId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TokenId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.ObjectsInfos))); err != nil {
		return err
	}

	for i := range m.ObjectsInfos {

		if err := m.ObjectsInfos[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeStartOkNpcShopMessage) Deserialize(r Reader) error {

	lnpcSellerId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.NpcSellerId = lnpcSellerId

	ltokenId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TokenId = ltokenId

	lobjectsInfosLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectsInfos = make([]ObjectItemToSellInNpcShop, lobjectsInfosLen)

	for i := range m.ObjectsInfos {

		var lobjectsInfos ObjectItemToSellInNpcShop

		lobjectsInfos.Deserialize(r)

		m.ObjectsInfos[i] = lobjectsInfos

	}

	return nil
}

type ExchangeGuildTaxCollectorGetMessage struct {
	CollectorName string

	WorldX int16

	WorldY int16

	MapId int32

	SubAreaId uint16

	UserName string

	CallerId int64

	CallerName string

	Experience float64

	Pods uint16

	ObjectsInfos []ObjectItemGenericQuantity
}

func (m *ExchangeGuildTaxCollectorGetMessage) ID() uint16 {
	return 5762
}

func (m *ExchangeGuildTaxCollectorGetMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.CollectorName); err != nil {
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

	if err := w.WriteString(m.UserName); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.CallerId); err != nil {
		return err
	}

	if err := w.WriteString(m.CallerName); err != nil {
		return err
	}

	if err := w.WriteDouble(m.Experience); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Pods); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.ObjectsInfos))); err != nil {
		return err
	}

	for i := range m.ObjectsInfos {

		if err := m.ObjectsInfos[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeGuildTaxCollectorGetMessage) Deserialize(r Reader) error {

	lcollectorName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.CollectorName = lcollectorName

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

	luserName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.UserName = luserName

	lcallerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.CallerId = lcallerId

	lcallerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.CallerName = lcallerName

	lexperience, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Experience = lexperience

	lpods, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Pods = lpods

	lobjectsInfosLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectsInfos = make([]ObjectItemGenericQuantity, lobjectsInfosLen)

	for i := range m.ObjectsInfos {

		var lobjectsInfos ObjectItemGenericQuantity

		lobjectsInfos.Deserialize(r)

		m.ObjectsInfos[i] = lobjectsInfos

	}

	return nil
}

type ExchangeTypesExchangerDescriptionForUserMessage struct {
	TypeDescription []uint32
}

func (m *ExchangeTypesExchangerDescriptionForUserMessage) ID() uint16 {
	return 5765
}

func (m *ExchangeTypesExchangerDescriptionForUserMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.TypeDescription))); err != nil {
		return err
	}

	for i := range m.TypeDescription {

		if err := w.WriteVarUInt32(m.TypeDescription[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeTypesExchangerDescriptionForUserMessage) Deserialize(r Reader) error {

	ltypeDescriptionLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.TypeDescription = make([]uint32, ltypeDescriptionLen)

	for i := range m.TypeDescription {

		ltypeDescription, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.TypeDescription[i] = ltypeDescription

	}

	return nil
}

type ExchangeStartOkHumanVendorMessage struct {
	SellerId float64

	ObjectsInfos []ObjectItemToSellInHumanVendorShop
}

func (m *ExchangeStartOkHumanVendorMessage) ID() uint16 {
	return 5767
}

func (m *ExchangeStartOkHumanVendorMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.SellerId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.ObjectsInfos))); err != nil {
		return err
	}

	for i := range m.ObjectsInfos {

		if err := m.ObjectsInfos[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeStartOkHumanVendorMessage) Deserialize(r Reader) error {

	lsellerId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.SellerId = lsellerId

	lobjectsInfosLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectsInfos = make([]ObjectItemToSellInHumanVendorShop, lobjectsInfosLen)

	for i := range m.ObjectsInfos {

		var lobjectsInfos ObjectItemToSellInHumanVendorShop

		lobjectsInfos.Deserialize(r)

		m.ObjectsInfos[i] = lobjectsInfos

	}

	return nil
}

type ExchangeOkMultiCraftMessage struct {
	InitiatorId int64

	OtherId int64

	Role int8
}

func (m *ExchangeOkMultiCraftMessage) ID() uint16 {
	return 5768
}

func (m *ExchangeOkMultiCraftMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.InitiatorId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.OtherId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Role); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeOkMultiCraftMessage) Deserialize(r Reader) error {

	linitiatorId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.InitiatorId = linitiatorId

	lotherId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.OtherId = lotherId

	lrole, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Role = lrole

	return nil
}

type ItemNoMoreAvailableMessage struct {
}

func (m *ItemNoMoreAvailableMessage) ID() uint16 {
	return 5769
}

func (m *ItemNoMoreAvailableMessage) Serialize(w Writer) error {

	return nil
}

func (m *ItemNoMoreAvailableMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeOnHumanVendorRequestMessage struct {
	HumanVendorId int64

	HumanVendorCell uint16
}

func (m *ExchangeOnHumanVendorRequestMessage) ID() uint16 {
	return 5772
}

func (m *ExchangeOnHumanVendorRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.HumanVendorId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.HumanVendorCell); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeOnHumanVendorRequestMessage) Deserialize(r Reader) error {

	lhumanVendorId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.HumanVendorId = lhumanVendorId

	lhumanVendorCell, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.HumanVendorCell = lhumanVendorCell

	return nil
}

type ExchangePlayerRequestMessage struct {
	ExchangeRequestMessage

	Target int64
}

func (m *ExchangePlayerRequestMessage) ID() uint16 {
	return 5773
}

func (m *ExchangePlayerRequestMessage) Serialize(w Writer) error {

	if err := m.ExchangeRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Target); err != nil {
		return err
	}

	return nil
}

func (m *ExchangePlayerRequestMessage) Deserialize(r Reader) error {

	if err := m.ExchangeRequestMessage.Deserialize(r); err != nil {
		return err
	}

	ltarget, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Target = ltarget

	return nil
}

type ExchangeBuyMessage struct {
	ObjectToBuyId uint32

	Quantity uint32
}

func (m *ExchangeBuyMessage) ID() uint16 {
	return 5774
}

func (m *ExchangeBuyMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectToBuyId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBuyMessage) Deserialize(r Reader) error {

	lobjectToBuyId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectToBuyId = lobjectToBuyId

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type ExchangeStartAsVendorMessage struct {
}

func (m *ExchangeStartAsVendorMessage) ID() uint16 {
	return 5775
}

func (m *ExchangeStartAsVendorMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeStartAsVendorMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeSellMessage struct {
	ObjectToSellId uint32

	Quantity uint32
}

func (m *ExchangeSellMessage) ID() uint16 {
	return 5778
}

func (m *ExchangeSellMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectToSellId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeSellMessage) Deserialize(r Reader) error {

	lobjectToSellId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectToSellId = lobjectToSellId

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type ExchangeRequestOnTaxCollectorMessage struct {
	TaxCollectorId int32
}

func (m *ExchangeRequestOnTaxCollectorMessage) ID() uint16 {
	return 5779
}

func (m *ExchangeRequestOnTaxCollectorMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.TaxCollectorId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeRequestOnTaxCollectorMessage) Deserialize(r Reader) error {

	ltaxCollectorId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.TaxCollectorId = ltaxCollectorId

	return nil
}

type ExchangeShowVendorTaxMessage struct {
}

func (m *ExchangeShowVendorTaxMessage) ID() uint16 {
	return 5783
}

func (m *ExchangeShowVendorTaxMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeShowVendorTaxMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangePlayerMultiCraftRequestMessage struct {
	ExchangeRequestMessage

	Target int64

	SkillId uint32
}

func (m *ExchangePlayerMultiCraftRequestMessage) ID() uint16 {
	return 5784
}

func (m *ExchangePlayerMultiCraftRequestMessage) Serialize(w Writer) error {

	if err := m.ExchangeRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Target); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.SkillId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangePlayerMultiCraftRequestMessage) Deserialize(r Reader) error {

	if err := m.ExchangeRequestMessage.Deserialize(r); err != nil {
		return err
	}

	ltarget, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Target = ltarget

	lskillId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SkillId = lskillId

	return nil
}

type ExchangeStartOkNpcTradeMessage struct {
	NpcId float64
}

func (m *ExchangeStartOkNpcTradeMessage) ID() uint16 {
	return 5785
}

func (m *ExchangeStartOkNpcTradeMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.NpcId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeStartOkNpcTradeMessage) Deserialize(r Reader) error {

	lnpcId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.NpcId = lnpcId

	return nil
}

type ExchangeWaitingResultMessage struct {
	Bwait bool
}

func (m *ExchangeWaitingResultMessage) ID() uint16 {
	return 5786
}

func (m *ExchangeWaitingResultMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Bwait); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeWaitingResultMessage) Deserialize(r Reader) error {

	lbwait, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Bwait = lbwait

	return nil
}

type ExchangeReplyTaxVendorMessage struct {
	ObjectValue int64

	TotalTaxValue int64
}

func (m *ExchangeReplyTaxVendorMessage) ID() uint16 {
	return 5787
}

func (m *ExchangeReplyTaxVendorMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.ObjectValue); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.TotalTaxValue); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeReplyTaxVendorMessage) Deserialize(r Reader) error {

	lobjectValue, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ObjectValue = lobjectValue

	ltotalTaxValue, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.TotalTaxValue = ltotalTaxValue

	return nil
}

type ExchangeCraftResultMessage struct {
	CraftResult uint8
}

func (m *ExchangeCraftResultMessage) ID() uint16 {
	return 5790
}

func (m *ExchangeCraftResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.CraftResult); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeCraftResultMessage) Deserialize(r Reader) error {

	lcraftResult, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CraftResult = lcraftResult

	return nil
}

type ExchangeSellOkMessage struct {
}

func (m *ExchangeSellOkMessage) ID() uint16 {
	return 5792
}

func (m *ExchangeSellOkMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeSellOkMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeWeightMessage struct {
	CurrentWeight uint32

	MaxWeight uint32
}

func (m *ExchangeWeightMessage) ID() uint16 {
	return 5793
}

func (m *ExchangeWeightMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.CurrentWeight); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MaxWeight); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeWeightMessage) Deserialize(r Reader) error {

	lcurrentWeight, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.CurrentWeight = lcurrentWeight

	lmaxWeight, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MaxWeight = lmaxWeight

	return nil
}

type ExchangeCraftInformationObjectMessage struct {
	ExchangeCraftResultWithObjectIdMessage

	PlayerId int64
}

func (m *ExchangeCraftInformationObjectMessage) ID() uint16 {
	return 5794
}

func (m *ExchangeCraftInformationObjectMessage) Serialize(w Writer) error {

	if err := m.ExchangeCraftResultWithObjectIdMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeCraftInformationObjectMessage) Deserialize(r Reader) error {

	if err := m.ExchangeCraftResultWithObjectIdMessage.Deserialize(r); err != nil {
		return err
	}

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type ExchangeBidSearchOkMessage struct {
}

func (m *ExchangeBidSearchOkMessage) ID() uint16 {
	return 5802
}

func (m *ExchangeBidSearchOkMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeBidSearchOkMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeBidHouseTypeMessage struct {
	Type uint32
}

func (m *ExchangeBidHouseTypeMessage) ID() uint16 {
	return 5803
}

func (m *ExchangeBidHouseTypeMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHouseTypeMessage) Deserialize(r Reader) error {

	ltype, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Type = ltype

	return nil
}

type ExchangeBidHouseBuyMessage struct {
	Uid uint32

	Qty uint32

	Price int64
}

func (m *ExchangeBidHouseBuyMessage) ID() uint16 {
	return 5804
}

func (m *ExchangeBidHouseBuyMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.Uid); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Qty); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Price); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHouseBuyMessage) Deserialize(r Reader) error {

	luid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Uid = luid

	lqty, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Qty = lqty

	lprice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Price = lprice

	return nil
}

type ExchangeBidHousePriceMessage struct {
	GenId uint16
}

func (m *ExchangeBidHousePriceMessage) ID() uint16 {
	return 5805
}

func (m *ExchangeBidHousePriceMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.GenId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHousePriceMessage) Deserialize(r Reader) error {

	lgenId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.GenId = lgenId

	return nil
}

type ExchangeBidHouseSearchMessage struct {
	Type uint32

	GenId uint16
}

func (m *ExchangeBidHouseSearchMessage) ID() uint16 {
	return 5806
}

func (m *ExchangeBidHouseSearchMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.Type); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.GenId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHouseSearchMessage) Deserialize(r Reader) error {

	ltype, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Type = ltype

	lgenId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.GenId = lgenId

	return nil
}

type ExchangeBidHouseListMessage struct {
	Id uint16
}

func (m *ExchangeBidHouseListMessage) ID() uint16 {
	return 5807
}

func (m *ExchangeBidHouseListMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHouseListMessage) Deserialize(r Reader) error {

	lid, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type JobExperienceMultiUpdateMessage struct {
	ExperiencesUpdate []JobExperience
}

func (m *JobExperienceMultiUpdateMessage) ID() uint16 {
	return 5809
}

func (m *JobExperienceMultiUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ExperiencesUpdate))); err != nil {
		return err
	}

	for i := range m.ExperiencesUpdate {

		if err := m.ExperiencesUpdate[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *JobExperienceMultiUpdateMessage) Deserialize(r Reader) error {

	lexperiencesUpdateLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ExperiencesUpdate = make([]JobExperience, lexperiencesUpdateLen)

	for i := range m.ExperiencesUpdate {

		var lexperiencesUpdate JobExperience

		lexperiencesUpdate.Deserialize(r)

		m.ExperiencesUpdate[i] = lexperiencesUpdate

	}

	return nil
}

type ExchangeItemAutoCraftStopedMessage struct {
	Reason int8
}

func (m *ExchangeItemAutoCraftStopedMessage) ID() uint16 {
	return 5810
}

func (m *ExchangeItemAutoCraftStopedMessage) Serialize(w Writer) error {

	if err := w.WriteInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeItemAutoCraftStopedMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type ExchangeStartOkCraftMessage struct {
}

func (m *ExchangeStartOkCraftMessage) ID() uint16 {
	return 5813
}

func (m *ExchangeStartOkCraftMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeStartOkCraftMessage) Deserialize(r Reader) error {

	return nil
}

type BasicLatencyStatsRequestMessage struct {
}

func (m *BasicLatencyStatsRequestMessage) ID() uint16 {
	return 5816
}

func (m *BasicLatencyStatsRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *BasicLatencyStatsRequestMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeStartOkMulticraftCustomerMessage struct {
	SkillId uint32

	CrafterJobLevel uint8
}

func (m *ExchangeStartOkMulticraftCustomerMessage) ID() uint16 {
	return 5817
}

func (m *ExchangeStartOkMulticraftCustomerMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.SkillId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.CrafterJobLevel); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeStartOkMulticraftCustomerMessage) Deserialize(r Reader) error {

	lskillId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SkillId = lskillId

	lcrafterJobLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CrafterJobLevel = lcrafterJobLevel

	return nil
}

type ExchangeStartOkMulticraftCrafterMessage struct {
	SkillId uint32
}

func (m *ExchangeStartOkMulticraftCrafterMessage) ID() uint16 {
	return 5818
}

func (m *ExchangeStartOkMulticraftCrafterMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.SkillId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeStartOkMulticraftCrafterMessage) Deserialize(r Reader) error {

	lskillId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SkillId = lskillId

	return nil
}

type ExchangeStartOkJobIndexMessage struct {
	Jobs []uint32
}

func (m *ExchangeStartOkJobIndexMessage) ID() uint16 {
	return 5819
}

func (m *ExchangeStartOkJobIndexMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Jobs))); err != nil {
		return err
	}

	for i := range m.Jobs {

		if err := w.WriteVarUInt32(m.Jobs[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeStartOkJobIndexMessage) Deserialize(r Reader) error {

	ljobsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Jobs = make([]uint32, ljobsLen)

	for i := range m.Jobs {

		ljobs, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.Jobs[i] = ljobs

	}

	return nil
}

type GameActionFightInvisibilityMessage struct {
	AbstractGameActionMessage

	TargetId float64

	State uint8
}

func (m *GameActionFightInvisibilityMessage) ID() uint16 {
	return 5821
}

func (m *GameActionFightInvisibilityMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.State); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightInvisibilityMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lstate, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.State = lstate

	return nil
}

type GameRolePlayFightRequestCanceledMessage struct {
	FightId int32

	SourceId float64

	TargetId float64
}

func (m *GameRolePlayFightRequestCanceledMessage) ID() uint16 {
	return 5822
}

func (m *GameRolePlayFightRequestCanceledMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.SourceId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayFightRequestCanceledMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lsourceId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.SourceId = lsourceId

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type PaddockPropertiesMessage struct {
	Properties PaddockInstancesInformations
}

func (m *PaddockPropertiesMessage) ID() uint16 {
	return 5824
}

func (m *PaddockPropertiesMessage) Serialize(w Writer) error {

	if err := m.Properties.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PaddockPropertiesMessage) Deserialize(r Reader) error {

	var lproperties PaddockInstancesInformations

	lproperties.Deserialize(r)

	m.Properties = lproperties

	return nil
}

type GameActionFightSummonMessage struct {
	AbstractGameActionMessage

	Summons []GameFightFighterInformationsIntrf
}

func (m *GameActionFightSummonMessage) ID() uint16 {
	return 5825
}

func (m *GameActionFightSummonMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Summons))); err != nil {
		return err
	}

	for i := range m.Summons {

		if err := w.WriteUInt16(m.Summons[i].ID()); err != nil {
			return err
		}

		if err := m.Summons[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameActionFightSummonMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	lsummonsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Summons = make([]GameFightFighterInformationsIntrf, lsummonsLen)

	for i := range m.Summons {

		typesummonsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lsummons, err := GetType(typesummonsID)
		if err != nil {
			return err
		}

		lsummons.Deserialize(r)

		m.Summons[i] = lsummons.(GameFightFighterInformationsIntrf)

	}

	return nil
}

type GameActionFightDropCharacterMessage struct {
	AbstractGameActionMessage

	TargetId float64

	CellId int16
}

func (m *GameActionFightDropCharacterMessage) ID() uint16 {
	return 5826
}

func (m *GameActionFightDropCharacterMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightDropCharacterMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lcellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type GameActionFightDodgePointLossMessage struct {
	AbstractGameActionMessage

	TargetId float64

	Amount uint16
}

func (m *GameActionFightDodgePointLossMessage) ID() uint16 {
	return 5828
}

func (m *GameActionFightDodgePointLossMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightDodgePointLossMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lamount, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Amount = lamount

	return nil
}

type GameActionFightThrowCharacterMessage struct {
	AbstractGameActionMessage

	TargetId float64

	CellId int16
}

func (m *GameActionFightThrowCharacterMessage) ID() uint16 {
	return 5829
}

func (m *GameActionFightThrowCharacterMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightThrowCharacterMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lcellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type GameActionFightCarryCharacterMessage struct {
	AbstractGameActionMessage

	TargetId float64

	CellId int16
}

func (m *GameActionFightCarryCharacterMessage) ID() uint16 {
	return 5830
}

func (m *GameActionFightCarryCharacterMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightCarryCharacterMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lcellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type GuildMembershipMessage struct {
	GuildJoinedMessage
}

func (m *GuildMembershipMessage) ID() uint16 {
	return 5835
}

func (m *GuildMembershipMessage) Serialize(w Writer) error {

	if err := m.GuildJoinedMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildMembershipMessage) Deserialize(r Reader) error {

	if err := m.GuildJoinedMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type CharacterSelectedErrorMessage struct {
}

func (m *CharacterSelectedErrorMessage) ID() uint16 {
	return 5836
}

func (m *CharacterSelectedErrorMessage) Serialize(w Writer) error {

	return nil
}

func (m *CharacterSelectedErrorMessage) Deserialize(r Reader) error {

	return nil
}

type PrismFightJoinLeaveRequestMessage struct {
	SubAreaId uint16

	Join bool
}

func (m *PrismFightJoinLeaveRequestMessage) ID() uint16 {
	return 5843
}

func (m *PrismFightJoinLeaveRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Join); err != nil {
		return err
	}

	return nil
}

func (m *PrismFightJoinLeaveRequestMessage) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	ljoin, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Join = ljoin

	return nil
}

type PrismInfoJoinLeaveRequestMessage struct {
	Join bool
}

func (m *PrismInfoJoinLeaveRequestMessage) ID() uint16 {
	return 5844
}

func (m *PrismInfoJoinLeaveRequestMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Join); err != nil {
		return err
	}

	return nil
}

func (m *PrismInfoJoinLeaveRequestMessage) Deserialize(r Reader) error {

	ljoin, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Join = ljoin

	return nil
}

type PrismInfoCloseMessage struct {
}

func (m *PrismInfoCloseMessage) ID() uint16 {
	return 5853
}

func (m *PrismInfoCloseMessage) Serialize(w Writer) error {

	return nil
}

func (m *PrismInfoCloseMessage) Deserialize(r Reader) error {

	return nil
}

type PrismInfoInValidMessage struct {
	Reason uint8
}

func (m *PrismInfoInValidMessage) ID() uint16 {
	return 5859
}

func (m *PrismInfoInValidMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *PrismInfoInValidMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type GameFightShowFighterMessage struct {
	Informations GameFightFighterInformationsIntrf
}

func (m *GameFightShowFighterMessage) ID() uint16 {
	return 5864
}

func (m *GameFightShowFighterMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Informations.ID()); err != nil {
		return err
	}

	if err := m.Informations.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameFightShowFighterMessage) Deserialize(r Reader) error {

	typeinformationsID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	linformations, err := GetType(typeinformationsID)
	if err != nil {
		return err
	}

	linformations.Deserialize(r)

	m.Informations = linformations.(GameFightFighterInformationsIntrf)

	return nil
}

type HouseSellFromInsideRequestMessage struct {
	HouseSellRequestMessage
}

func (m *HouseSellFromInsideRequestMessage) ID() uint16 {
	return 5884
}

func (m *HouseSellFromInsideRequestMessage) Serialize(w Writer) error {

	if err := m.HouseSellRequestMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *HouseSellFromInsideRequestMessage) Deserialize(r Reader) error {

	if err := m.HouseSellRequestMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type HouseLockFromInsideRequestMessage struct {
	LockableChangeCodeMessage
}

func (m *HouseLockFromInsideRequestMessage) ID() uint16 {
	return 5885
}

func (m *HouseLockFromInsideRequestMessage) Serialize(w Writer) error {

	if err := m.LockableChangeCodeMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *HouseLockFromInsideRequestMessage) Deserialize(r Reader) error {

	if err := m.LockableChangeCodeMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type GuildKickRequestMessage struct {
	KickedId int64
}

func (m *GuildKickRequestMessage) ID() uint16 {
	return 5887
}

func (m *GuildKickRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.KickedId); err != nil {
		return err
	}

	return nil
}

func (m *GuildKickRequestMessage) Deserialize(r Reader) error {

	lkickedId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.KickedId = lkickedId

	return nil
}

type PrismFightDefenderLeaveMessage struct {
	SubAreaId uint16

	FightId uint16

	FighterToRemoveId int64
}

func (m *PrismFightDefenderLeaveMessage) ID() uint16 {
	return 5892
}

func (m *PrismFightDefenderLeaveMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.FightId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.FighterToRemoveId); err != nil {
		return err
	}

	return nil
}

func (m *PrismFightDefenderLeaveMessage) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lfightId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lfighterToRemoveId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.FighterToRemoveId = lfighterToRemoveId

	return nil
}

type PrismFightAttackerAddMessage struct {
	SubAreaId uint16

	FightId uint16

	Attacker CharacterMinimalPlusLookInformationsIntrf
}

func (m *PrismFightAttackerAddMessage) ID() uint16 {
	return 5893
}

func (m *PrismFightAttackerAddMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.FightId); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Attacker.ID()); err != nil {
		return err
	}

	if err := m.Attacker.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PrismFightAttackerAddMessage) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lfightId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	typeattackerID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lattacker, err := GetType(typeattackerID)
	if err != nil {
		return err
	}

	lattacker.Deserialize(r)

	m.Attacker = lattacker.(CharacterMinimalPlusLookInformationsIntrf)

	return nil
}

type PrismFightDefenderAddMessage struct {
	SubAreaId uint16

	FightId uint16

	Defender CharacterMinimalPlusLookInformationsIntrf
}

func (m *PrismFightDefenderAddMessage) ID() uint16 {
	return 5895
}

func (m *PrismFightDefenderAddMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.FightId); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Defender.ID()); err != nil {
		return err
	}

	if err := m.Defender.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PrismFightDefenderAddMessage) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lfightId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	typedefenderID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	ldefender, err := GetType(typedefenderID)
	if err != nil {
		return err
	}

	ldefender.Deserialize(r)

	m.Defender = ldefender.(CharacterMinimalPlusLookInformationsIntrf)

	return nil
}

type PrismFightAttackerRemoveMessage struct {
	SubAreaId uint16

	FightId uint16

	FighterToRemoveId int64
}

func (m *PrismFightAttackerRemoveMessage) ID() uint16 {
	return 5897
}

func (m *PrismFightAttackerRemoveMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.FightId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.FighterToRemoveId); err != nil {
		return err
	}

	return nil
}

func (m *PrismFightAttackerRemoveMessage) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lfightId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lfighterToRemoveId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.FighterToRemoveId = lfighterToRemoveId

	return nil
}

type NpcGenericActionRequestMessage struct {
	NpcId int32

	NpcActionId uint8

	NpcMapId int32
}

func (m *NpcGenericActionRequestMessage) ID() uint16 {
	return 5898
}

func (m *NpcGenericActionRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.NpcId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.NpcActionId); err != nil {
		return err
	}

	if err := w.WriteInt32(m.NpcMapId); err != nil {
		return err
	}

	return nil
}

func (m *NpcGenericActionRequestMessage) Deserialize(r Reader) error {

	lnpcId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.NpcId = lnpcId

	lnpcActionId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NpcActionId = lnpcActionId

	lnpcMapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.NpcMapId = lnpcMapId

	return nil
}

type NpcGenericActionFailureMessage struct {
}

func (m *NpcGenericActionFailureMessage) ID() uint16 {
	return 5900
}

func (m *NpcGenericActionFailureMessage) Serialize(w Writer) error {

	return nil
}

func (m *NpcGenericActionFailureMessage) Deserialize(r Reader) error {

	return nil
}

type PrismFightSwapRequestMessage struct {
	SubAreaId uint16

	TargetId int64
}

func (m *PrismFightSwapRequestMessage) ID() uint16 {
	return 5901
}

func (m *PrismFightSwapRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *PrismFightSwapRequestMessage) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	ltargetId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type ExchangeStartedBidBuyerMessage struct {
	BuyerDescriptor SellerBuyerDescriptor
}

func (m *ExchangeStartedBidBuyerMessage) ID() uint16 {
	return 5904
}

func (m *ExchangeStartedBidBuyerMessage) Serialize(w Writer) error {

	if err := m.BuyerDescriptor.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeStartedBidBuyerMessage) Deserialize(r Reader) error {

	var lbuyerDescriptor SellerBuyerDescriptor

	lbuyerDescriptor.Deserialize(r)

	m.BuyerDescriptor = lbuyerDescriptor

	return nil
}

type ExchangeStartedBidSellerMessage struct {
	SellerDescriptor SellerBuyerDescriptor

	ObjectsInfos []ObjectItemToSellInBid
}

func (m *ExchangeStartedBidSellerMessage) ID() uint16 {
	return 5905
}

func (m *ExchangeStartedBidSellerMessage) Serialize(w Writer) error {

	if err := m.SellerDescriptor.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.ObjectsInfos))); err != nil {
		return err
	}

	for i := range m.ObjectsInfos {

		if err := m.ObjectsInfos[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeStartedBidSellerMessage) Deserialize(r Reader) error {

	var lsellerDescriptor SellerBuyerDescriptor

	lsellerDescriptor.Deserialize(r)

	m.SellerDescriptor = lsellerDescriptor

	lobjectsInfosLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectsInfos = make([]ObjectItemToSellInBid, lobjectsInfosLen)

	for i := range m.ObjectsInfos {

		var lobjectsInfos ObjectItemToSellInBid

		lobjectsInfos.Deserialize(r)

		m.ObjectsInfos[i] = lobjectsInfos

	}

	return nil
}

type ExchangeShopStockMovementRemovedMessage struct {
	ObjectId uint32
}

func (m *ExchangeShopStockMovementRemovedMessage) ID() uint16 {
	return 5907
}

func (m *ExchangeShopStockMovementRemovedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeShopStockMovementRemovedMessage) Deserialize(r Reader) error {

	lobjectId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectId = lobjectId

	return nil
}

type ChallengeFightJoinRefusedMessage struct {
	PlayerId int64

	Reason int8
}

func (m *ChallengeFightJoinRefusedMessage) ID() uint16 {
	return 5908
}

func (m *ChallengeFightJoinRefusedMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *ChallengeFightJoinRefusedMessage) Deserialize(r Reader) error {

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	lreason, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type ExchangeShopStockMovementUpdatedMessage struct {
	ObjectInfo ObjectItemToSell
}

func (m *ExchangeShopStockMovementUpdatedMessage) ID() uint16 {
	return 5909
}

func (m *ExchangeShopStockMovementUpdatedMessage) Serialize(w Writer) error {

	if err := m.ObjectInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeShopStockMovementUpdatedMessage) Deserialize(r Reader) error {

	var lobjectInfo ObjectItemToSell

	lobjectInfo.Deserialize(r)

	m.ObjectInfo = lobjectInfo

	return nil
}

type ExchangeShopStockStartedMessage struct {
	ObjectsInfos []ObjectItemToSell
}

func (m *ExchangeShopStockStartedMessage) ID() uint16 {
	return 5910
}

func (m *ExchangeShopStockStartedMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ObjectsInfos))); err != nil {
		return err
	}

	for i := range m.ObjectsInfos {

		if err := m.ObjectsInfos[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeShopStockStartedMessage) Deserialize(r Reader) error {

	lobjectsInfosLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectsInfos = make([]ObjectItemToSell, lobjectsInfosLen)

	for i := range m.ObjectsInfos {

		var lobjectsInfos ObjectItemToSell

		lobjectsInfos.Deserialize(r)

		m.ObjectsInfos[i] = lobjectsInfos

	}

	return nil
}

type TaxCollectorMovementRemoveMessage struct {
	CollectorId int32
}

func (m *TaxCollectorMovementRemoveMessage) ID() uint16 {
	return 5915
}

func (m *TaxCollectorMovementRemoveMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.CollectorId); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorMovementRemoveMessage) Deserialize(r Reader) error {

	lcollectorId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.CollectorId = lcollectorId

	return nil
}

type TaxCollectorMovementAddMessage struct {
	Informations TaxCollectorInformationsIntrf
}

func (m *TaxCollectorMovementAddMessage) ID() uint16 {
	return 5917
}

func (m *TaxCollectorMovementAddMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Informations.ID()); err != nil {
		return err
	}

	if err := m.Informations.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorMovementAddMessage) Deserialize(r Reader) error {

	typeinformationsID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	linformations, err := GetType(typeinformationsID)
	if err != nil {
		return err
	}

	linformations.Deserialize(r)

	m.Informations = linformations.(TaxCollectorInformationsIntrf)

	return nil
}

type TaxCollectorAttackedMessage struct {
	FirstNameId uint16

	LastNameId uint16

	WorldX int16

	WorldY int16

	MapId int32

	SubAreaId uint16

	Guild BasicGuildInformations
}

func (m *TaxCollectorAttackedMessage) ID() uint16 {
	return 5918
}

func (m *TaxCollectorAttackedMessage) Serialize(w Writer) error {

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

	if err := m.Guild.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorAttackedMessage) Deserialize(r Reader) error {

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

	var lguild BasicGuildInformations

	lguild.Deserialize(r)

	m.Guild = lguild

	return nil
}

type GuildHousesInformationMessage struct {
	HousesInformations []HouseInformationsForGuild
}

func (m *GuildHousesInformationMessage) ID() uint16 {
	return 5919
}

func (m *GuildHousesInformationMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.HousesInformations))); err != nil {
		return err
	}

	for i := range m.HousesInformations {

		if err := m.HousesInformations[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GuildHousesInformationMessage) Deserialize(r Reader) error {

	lhousesInformationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.HousesInformations = make([]HouseInformationsForGuild, lhousesInformationsLen)

	for i := range m.HousesInformations {

		var lhousesInformations HouseInformationsForGuild

		lhousesInformations.Deserialize(r)

		m.HousesInformations[i] = lhousesInformations

	}

	return nil
}

type GuildCreationStartedMessage struct {
}

func (m *GuildCreationStartedMessage) ID() uint16 {
	return 5920
}

func (m *GuildCreationStartedMessage) Serialize(w Writer) error {

	return nil
}

func (m *GuildCreationStartedMessage) Deserialize(r Reader) error {

	return nil
}

type GameFightSynchronizeMessage struct {
	Fighters []GameFightFighterInformationsIntrf
}

func (m *GameFightSynchronizeMessage) ID() uint16 {
	return 5921
}

func (m *GameFightSynchronizeMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Fighters))); err != nil {
		return err
	}

	for i := range m.Fighters {

		if err := w.WriteUInt16(m.Fighters[i].ID()); err != nil {
			return err
		}

		if err := m.Fighters[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameFightSynchronizeMessage) Deserialize(r Reader) error {

	lfightersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Fighters = make([]GameFightFighterInformationsIntrf, lfightersLen)

	for i := range m.Fighters {

		typefightersID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lfighters, err := GetType(typefightersID)
		if err != nil {
			return err
		}

		lfighters.Deserialize(r)

		m.Fighters[i] = lfighters.(GameFightFighterInformationsIntrf)

	}

	return nil
}

type GuildMemberLeavingMessage struct {
	Kicked bool

	MemberId int64
}

func (m *GuildMemberLeavingMessage) ID() uint16 {
	return 5923
}

func (m *GuildMemberLeavingMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Kicked); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.MemberId); err != nil {
		return err
	}

	return nil
}

func (m *GuildMemberLeavingMessage) Deserialize(r Reader) error {

	lkicked, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Kicked = lkicked

	lmemberId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.MemberId = lmemberId

	return nil
}

type FriendUpdateMessage struct {
	FriendUpdated FriendInformationsIntrf
}

func (m *FriendUpdateMessage) ID() uint16 {
	return 5924
}

func (m *FriendUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.FriendUpdated.ID()); err != nil {
		return err
	}

	if err := m.FriendUpdated.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *FriendUpdateMessage) Deserialize(r Reader) error {

	typefriendUpdatedID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lfriendUpdated, err := GetType(typefriendUpdatedID)
	if err != nil {
		return err
	}

	lfriendUpdated.Deserialize(r)

	m.FriendUpdated = lfriendUpdated.(FriendInformationsIntrf)

	return nil
}

type ObjectGroundListAddedMessage struct {
	Cells []uint16

	ReferenceIds []uint16
}

func (m *ObjectGroundListAddedMessage) ID() uint16 {
	return 5925
}

func (m *ObjectGroundListAddedMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Cells))); err != nil {
		return err
	}

	for i := range m.Cells {

		if err := w.WriteVarUInt16(m.Cells[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.ReferenceIds))); err != nil {
		return err
	}

	for i := range m.ReferenceIds {

		if err := w.WriteVarUInt16(m.ReferenceIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ObjectGroundListAddedMessage) Deserialize(r Reader) error {

	lcellsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Cells = make([]uint16, lcellsLen)

	for i := range m.Cells {

		lcells, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Cells[i] = lcells

	}

	lreferenceIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ReferenceIds = make([]uint16, lreferenceIdsLen)

	for i := range m.ReferenceIds {

		lreferenceIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.ReferenceIds[i] = lreferenceIds

	}

	return nil
}

type GameFightOptionStateUpdateMessage struct {
	FightId uint16

	TeamId uint8

	Option uint8

	State bool
}

func (m *GameFightOptionStateUpdateMessage) ID() uint16 {
	return 5927
}

func (m *GameFightOptionStateUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.FightId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.TeamId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Option); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.State); err != nil {
		return err
	}

	return nil
}

func (m *GameFightOptionStateUpdateMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lteamId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TeamId = lteamId

	loption, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Option = loption

	lstate, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.State = lstate

	return nil
}

type GuildFightPlayersEnemiesListMessage struct {
	FightId uint32

	PlayerInfo []CharacterMinimalPlusLookInformations
}

func (m *GuildFightPlayersEnemiesListMessage) ID() uint16 {
	return 5928
}

func (m *GuildFightPlayersEnemiesListMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.PlayerInfo))); err != nil {
		return err
	}

	for i := range m.PlayerInfo {

		if err := m.PlayerInfo[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GuildFightPlayersEnemiesListMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lplayerInfoLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PlayerInfo = make([]CharacterMinimalPlusLookInformations, lplayerInfoLen)

	for i := range m.PlayerInfo {

		var lplayerInfo CharacterMinimalPlusLookInformations

		lplayerInfo.Deserialize(r)

		m.PlayerInfo[i] = lplayerInfo

	}

	return nil
}

type GuildFightPlayersEnemyRemoveMessage struct {
	FightId uint32

	PlayerId int64
}

func (m *GuildFightPlayersEnemyRemoveMessage) ID() uint16 {
	return 5929
}

func (m *GuildFightPlayersEnemyRemoveMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *GuildFightPlayersEnemyRemoveMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type TaxCollectorListMessage struct {
	AbstractTaxCollectorListMessage

	NbcollectorMax uint8

	FightersInformations []TaxCollectorFightersInformation
}

func (m *TaxCollectorListMessage) ID() uint16 {
	return 5930
}

func (m *TaxCollectorListMessage) Serialize(w Writer) error {

	if err := m.AbstractTaxCollectorListMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.NbcollectorMax); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.FightersInformations))); err != nil {
		return err
	}

	for i := range m.FightersInformations {

		if err := m.FightersInformations[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *TaxCollectorListMessage) Deserialize(r Reader) error {

	if err := m.AbstractTaxCollectorListMessage.Deserialize(r); err != nil {
		return err
	}

	lnbcollectorMax, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NbcollectorMax = lnbcollectorMax

	lfightersInformationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FightersInformations = make([]TaxCollectorFightersInformation, lfightersInformationsLen)

	for i := range m.FightersInformations {

		var lfightersInformations TaxCollectorFightersInformation

		lfightersInformations.Deserialize(r)

		m.FightersInformations[i] = lfightersInformations

	}

	return nil
}

type ContactLookRequestMessage struct {
	RequestId uint8

	ContactType uint8
}

func (m *ContactLookRequestMessage) ID() uint16 {
	return 5932
}

func (m *ContactLookRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.RequestId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.ContactType); err != nil {
		return err
	}

	return nil
}

func (m *ContactLookRequestMessage) Deserialize(r Reader) error {

	lrequestId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.RequestId = lrequestId

	lcontactType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ContactType = lcontactType

	return nil
}

type ContactLookRequestByNameMessage struct {
	ContactLookRequestMessage

	PlayerName string
}

func (m *ContactLookRequestByNameMessage) ID() uint16 {
	return 5933
}

func (m *ContactLookRequestByNameMessage) Serialize(w Writer) error {

	if err := m.ContactLookRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.PlayerName); err != nil {
		return err
	}

	return nil
}

func (m *ContactLookRequestByNameMessage) Deserialize(r Reader) error {

	if err := m.ContactLookRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lplayerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PlayerName = lplayerName

	return nil
}

type ContactLookMessage struct {
	RequestId uint32

	PlayerName string

	PlayerId int64

	Look EntityLook
}

func (m *ContactLookMessage) ID() uint16 {
	return 5934
}

func (m *ContactLookMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.RequestId); err != nil {
		return err
	}

	if err := w.WriteString(m.PlayerName); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	if err := m.Look.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ContactLookMessage) Deserialize(r Reader) error {

	lrequestId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.RequestId = lrequestId

	lplayerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PlayerName = lplayerName

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	var llook EntityLook

	llook.Deserialize(r)

	m.Look = llook

	return nil
}

type ContactLookRequestByIdMessage struct {
	ContactLookRequestMessage

	PlayerId int64
}

func (m *ContactLookRequestByIdMessage) ID() uint16 {
	return 5935
}

func (m *ContactLookRequestByIdMessage) Serialize(w Writer) error {

	if err := m.ContactLookRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *ContactLookRequestByIdMessage) Deserialize(r Reader) error {

	if err := m.ContactLookRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type GameRolePlayPlayerFightFriendlyRequestedMessage struct {
	FightId uint32

	SourceId int64

	TargetId int64
}

func (m *GameRolePlayPlayerFightFriendlyRequestedMessage) ID() uint16 {
	return 5937
}

func (m *GameRolePlayPlayerFightFriendlyRequestedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.SourceId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayPlayerFightFriendlyRequestedMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lsourceId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.SourceId = lsourceId

	ltargetId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type ExchangeStartOkCraftWithInformationMessage struct {
	ExchangeStartOkCraftMessage

	SkillId uint32
}

func (m *ExchangeStartOkCraftWithInformationMessage) ID() uint16 {
	return 5941
}

func (m *ExchangeStartOkCraftWithInformationMessage) Serialize(w Writer) error {

	if err := m.ExchangeStartOkCraftMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.SkillId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeStartOkCraftWithInformationMessage) Deserialize(r Reader) error {

	if err := m.ExchangeStartOkCraftMessage.Deserialize(r); err != nil {
		return err
	}

	lskillId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SkillId = lskillId

	return nil
}

type ObjectGroundRemovedMultipleMessage struct {
	Cells []uint16
}

func (m *ObjectGroundRemovedMultipleMessage) ID() uint16 {
	return 5944
}

func (m *ObjectGroundRemovedMultipleMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Cells))); err != nil {
		return err
	}

	for i := range m.Cells {

		if err := w.WriteVarUInt16(m.Cells[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ObjectGroundRemovedMultipleMessage) Deserialize(r Reader) error {

	lcellsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Cells = make([]uint16, lcellsLen)

	for i := range m.Cells {

		lcells, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Cells[i] = lcells

	}

	return nil
}

type ExchangeBidHouseItemAddOkMessage struct {
	ItemInfo ObjectItemToSellInBid
}

func (m *ExchangeBidHouseItemAddOkMessage) ID() uint16 {
	return 5945
}

func (m *ExchangeBidHouseItemAddOkMessage) Serialize(w Writer) error {

	if err := m.ItemInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHouseItemAddOkMessage) Deserialize(r Reader) error {

	var litemInfo ObjectItemToSellInBid

	litemInfo.Deserialize(r)

	m.ItemInfo = litemInfo

	return nil
}

type ExchangeBidHouseItemRemoveOkMessage struct {
	SellerId int32
}

func (m *ExchangeBidHouseItemRemoveOkMessage) ID() uint16 {
	return 5946
}

func (m *ExchangeBidHouseItemRemoveOkMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.SellerId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHouseItemRemoveOkMessage) Deserialize(r Reader) error {

	lsellerId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.SellerId = lsellerId

	return nil
}

type ExchangeBidHouseGenericItemAddedMessage struct {
	ObjGenericId uint16
}

func (m *ExchangeBidHouseGenericItemAddedMessage) ID() uint16 {
	return 5947
}

func (m *ExchangeBidHouseGenericItemAddedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ObjGenericId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHouseGenericItemAddedMessage) Deserialize(r Reader) error {

	lobjGenericId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjGenericId = lobjGenericId

	return nil
}

type ExchangeBidHouseGenericItemRemovedMessage struct {
	ObjGenericId uint16
}

func (m *ExchangeBidHouseGenericItemRemovedMessage) ID() uint16 {
	return 5948
}

func (m *ExchangeBidHouseGenericItemRemovedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ObjGenericId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHouseGenericItemRemovedMessage) Deserialize(r Reader) error {

	lobjGenericId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjGenericId = lobjGenericId

	return nil
}

type ExchangeBidHouseInListAddedMessage struct {
	ItemUID int32

	ObjGenericId int32

	Effects []ObjectEffectIntrf

	Prices []int64
}

func (m *ExchangeBidHouseInListAddedMessage) ID() uint16 {
	return 5949
}

func (m *ExchangeBidHouseInListAddedMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.ItemUID); err != nil {
		return err
	}

	if err := w.WriteInt32(m.ObjGenericId); err != nil {
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

func (m *ExchangeBidHouseInListAddedMessage) Deserialize(r Reader) error {

	litemUID, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.ItemUID = litemUID

	lobjGenericId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.ObjGenericId = lobjGenericId

	leffectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Effects = make([]ObjectEffectIntrf, leffectsLen)

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

		m.Effects[i] = leffects.(ObjectEffectIntrf)

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

type ExchangeBidHouseInListRemovedMessage struct {
	ItemUID int32
}

func (m *ExchangeBidHouseInListRemovedMessage) ID() uint16 {
	return 5950
}

func (m *ExchangeBidHouseInListRemovedMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.ItemUID); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHouseInListRemovedMessage) Deserialize(r Reader) error {

	litemUID, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.ItemUID = litemUID

	return nil
}

type PaddockBuyRequestMessage struct {
	ProposedPrice int64
}

func (m *PaddockBuyRequestMessage) ID() uint16 {
	return 5951
}

func (m *PaddockBuyRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.ProposedPrice); err != nil {
		return err
	}

	return nil
}

func (m *PaddockBuyRequestMessage) Deserialize(r Reader) error {

	lproposedPrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ProposedPrice = lproposedPrice

	return nil
}

type GuildPaddockBoughtMessage struct {
	PaddockInfo PaddockContentInformations
}

func (m *GuildPaddockBoughtMessage) ID() uint16 {
	return 5952
}

func (m *GuildPaddockBoughtMessage) Serialize(w Writer) error {

	if err := m.PaddockInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildPaddockBoughtMessage) Deserialize(r Reader) error {

	var lpaddockInfo PaddockContentInformations

	lpaddockInfo.Deserialize(r)

	m.PaddockInfo = lpaddockInfo

	return nil
}

type PaddockSellRequestMessage struct {
	Price int64

	ForSale bool
}

func (m *PaddockSellRequestMessage) ID() uint16 {
	return 5953
}

func (m *PaddockSellRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.Price); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.ForSale); err != nil {
		return err
	}

	return nil
}

func (m *PaddockSellRequestMessage) Deserialize(r Reader) error {

	lprice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Price = lprice

	lforSale, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.ForSale = lforSale

	return nil
}

type GameRolePlayTaxCollectorFightRequestMessage struct {
	TaxCollectorId int32
}

func (m *GameRolePlayTaxCollectorFightRequestMessage) ID() uint16 {
	return 5954
}

func (m *GameRolePlayTaxCollectorFightRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.TaxCollectorId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayTaxCollectorFightRequestMessage) Deserialize(r Reader) error {

	ltaxCollectorId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.TaxCollectorId = ltaxCollectorId

	return nil
}

type GuildPaddockRemovedMessage struct {
	PaddockId int32
}

func (m *GuildPaddockRemovedMessage) ID() uint16 {
	return 5955
}

func (m *GuildPaddockRemovedMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.PaddockId); err != nil {
		return err
	}

	return nil
}

func (m *GuildPaddockRemovedMessage) Deserialize(r Reader) error {

	lpaddockId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.PaddockId = lpaddockId

	return nil
}

type AtlasPointInformationsMessage struct {
	Type AtlasPointsInformations
}

func (m *AtlasPointInformationsMessage) ID() uint16 {
	return 5956
}

func (m *AtlasPointInformationsMessage) Serialize(w Writer) error {

	if err := m.Type.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AtlasPointInformationsMessage) Deserialize(r Reader) error {

	var ltype AtlasPointsInformations

	ltype.Deserialize(r)

	m.Type = ltype

	return nil
}

type GuildPaddockTeleportRequestMessage struct {
	PaddockId int32
}

func (m *GuildPaddockTeleportRequestMessage) ID() uint16 {
	return 5957
}

func (m *GuildPaddockTeleportRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.PaddockId); err != nil {
		return err
	}

	return nil
}

func (m *GuildPaddockTeleportRequestMessage) Deserialize(r Reader) error {

	lpaddockId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.PaddockId = lpaddockId

	return nil
}

type PaddockRemoveItemRequestMessage struct {
	CellId uint16
}

func (m *PaddockRemoveItemRequestMessage) ID() uint16 {
	return 5958
}

func (m *PaddockRemoveItemRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *PaddockRemoveItemRequestMessage) Deserialize(r Reader) error {

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type GuildInformationsPaddocksMessage struct {
	NbPaddockMax uint8

	PaddocksInformations []PaddockContentInformations
}

func (m *GuildInformationsPaddocksMessage) ID() uint16 {
	return 5959
}

func (m *GuildInformationsPaddocksMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.NbPaddockMax); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.PaddocksInformations))); err != nil {
		return err
	}

	for i := range m.PaddocksInformations {

		if err := m.PaddocksInformations[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GuildInformationsPaddocksMessage) Deserialize(r Reader) error {

	lnbPaddockMax, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NbPaddockMax = lnbPaddockMax

	lpaddocksInformationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PaddocksInformations = make([]PaddockContentInformations, lpaddocksInformationsLen)

	for i := range m.PaddocksInformations {

		var lpaddocksInformations PaddockContentInformations

		lpaddocksInformations.Deserialize(r)

		m.PaddocksInformations[i] = lpaddocksInformations

	}

	return nil
}

type TeleportDestinationsListMessage struct {
	TeleporterType uint8

	MapIds []uint32

	SubAreaIds []uint16

	Costs []uint16

	DestTeleporterType []uint8
}

func (m *TeleportDestinationsListMessage) ID() uint16 {
	return 5960
}

func (m *TeleportDestinationsListMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.TeleporterType); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.MapIds))); err != nil {
		return err
	}

	for i := range m.MapIds {

		if err := w.WriteUInt32(m.MapIds[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.SubAreaIds))); err != nil {
		return err
	}

	for i := range m.SubAreaIds {

		if err := w.WriteVarUInt16(m.SubAreaIds[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Costs))); err != nil {
		return err
	}

	for i := range m.Costs {

		if err := w.WriteVarUInt16(m.Costs[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.DestTeleporterType))); err != nil {
		return err
	}

	for i := range m.DestTeleporterType {

		if err := w.WriteUInt8(m.DestTeleporterType[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *TeleportDestinationsListMessage) Deserialize(r Reader) error {

	lteleporterType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TeleporterType = lteleporterType

	lmapIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MapIds = make([]uint32, lmapIdsLen)

	for i := range m.MapIds {

		lmapIds, err := r.ReadUInt32()
		if err != nil {
			return err
		}

		m.MapIds[i] = lmapIds

	}

	lsubAreaIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SubAreaIds = make([]uint16, lsubAreaIdsLen)

	for i := range m.SubAreaIds {

		lsubAreaIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.SubAreaIds[i] = lsubAreaIds

	}

	lcostsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Costs = make([]uint16, lcostsLen)

	for i := range m.Costs {

		lcosts, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Costs[i] = lcosts

	}

	ldestTeleporterTypeLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DestTeleporterType = make([]uint8, ldestTeleporterTypeLen)

	for i := range m.DestTeleporterType {

		ldestTeleporterType, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.DestTeleporterType[i] = ldestTeleporterType

	}

	return nil
}

type TeleportRequestMessage struct {
	TeleporterType uint8

	MapId uint32
}

func (m *TeleportRequestMessage) ID() uint16 {
	return 5961
}

func (m *TeleportRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.TeleporterType); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.MapId); err != nil {
		return err
	}

	return nil
}

func (m *TeleportRequestMessage) Deserialize(r Reader) error {

	lteleporterType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TeleporterType = lteleporterType

	lmapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	return nil
}

type MountSterilizeRequestMessage struct {
}

func (m *MountSterilizeRequestMessage) ID() uint16 {
	return 5962
}

func (m *MountSterilizeRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *MountSterilizeRequestMessage) Deserialize(r Reader) error {

	return nil
}

type MountEquipedErrorMessage struct {
	ErrorType uint8
}

func (m *MountEquipedErrorMessage) ID() uint16 {
	return 5963
}

func (m *MountEquipedErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.ErrorType); err != nil {
		return err
	}

	return nil
}

func (m *MountEquipedErrorMessage) Deserialize(r Reader) error {

	lerrorType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ErrorType = lerrorType

	return nil
}

type MountRidingMessage struct {
	IsRiding bool
}

func (m *MountRidingMessage) ID() uint16 {
	return 5967
}

func (m *MountRidingMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.IsRiding); err != nil {
		return err
	}

	return nil
}

func (m *MountRidingMessage) Deserialize(r Reader) error {

	lisRiding, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsRiding = lisRiding

	return nil
}

type MountSetMessage struct {
	MountData MountClientData
}

func (m *MountSetMessage) ID() uint16 {
	return 5968
}

func (m *MountSetMessage) Serialize(w Writer) error {

	if err := m.MountData.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *MountSetMessage) Deserialize(r Reader) error {

	var lmountData MountClientData

	lmountData.Deserialize(r)

	m.MountData = lmountData

	return nil
}

type MountXpRatioMessage struct {
	Ratio uint8
}

func (m *MountXpRatioMessage) ID() uint16 {
	return 5970
}

func (m *MountXpRatioMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Ratio); err != nil {
		return err
	}

	return nil
}

func (m *MountXpRatioMessage) Deserialize(r Reader) error {

	lratio, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Ratio = lratio

	return nil
}

type MountInformationRequestMessage struct {
	Id float64

	Time float64
}

func (m *MountInformationRequestMessage) ID() uint16 {
	return 5972
}

func (m *MountInformationRequestMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.Id); err != nil {
		return err
	}

	if err := w.WriteDouble(m.Time); err != nil {
		return err
	}

	return nil
}

func (m *MountInformationRequestMessage) Deserialize(r Reader) error {

	lid, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Id = lid

	ltime, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Time = ltime

	return nil
}

type MountDataMessage struct {
	MountData MountClientData
}

func (m *MountDataMessage) ID() uint16 {
	return 5973
}

func (m *MountDataMessage) Serialize(w Writer) error {

	if err := m.MountData.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *MountDataMessage) Deserialize(r Reader) error {

	var lmountData MountClientData

	lmountData.Deserialize(r)

	m.MountData = lmountData

	return nil
}

type MountInformationInPaddockRequestMessage struct {
	MapRideId int32
}

func (m *MountInformationInPaddockRequestMessage) ID() uint16 {
	return 5975
}

func (m *MountInformationInPaddockRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt32(m.MapRideId); err != nil {
		return err
	}

	return nil
}

func (m *MountInformationInPaddockRequestMessage) Deserialize(r Reader) error {

	lmapRideId, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.MapRideId = lmapRideId

	return nil
}

type MountToggleRidingRequestMessage struct {
}

func (m *MountToggleRidingRequestMessage) ID() uint16 {
	return 5976
}

func (m *MountToggleRidingRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *MountToggleRidingRequestMessage) Deserialize(r Reader) error {

	return nil
}

type MountSterilizedMessage struct {
	MountId int32
}

func (m *MountSterilizedMessage) ID() uint16 {
	return 5977
}

func (m *MountSterilizedMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt32(m.MountId); err != nil {
		return err
	}

	return nil
}

func (m *MountSterilizedMessage) Deserialize(r Reader) error {

	lmountId, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.MountId = lmountId

	return nil
}

type MountEmoteIconUsedOkMessage struct {
	MountId int32

	ReactionType uint8
}

func (m *MountEmoteIconUsedOkMessage) ID() uint16 {
	return 5978
}

func (m *MountEmoteIconUsedOkMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt32(m.MountId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.ReactionType); err != nil {
		return err
	}

	return nil
}

func (m *MountEmoteIconUsedOkMessage) Deserialize(r Reader) error {

	lmountId, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.MountId = lmountId

	lreactionType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ReactionType = lreactionType

	return nil
}

type ExchangeStartOkMountMessage struct {
	ExchangeStartOkMountWithOutPaddockMessage

	PaddockedMountsDescription []MountClientData
}

func (m *ExchangeStartOkMountMessage) ID() uint16 {
	return 5979
}

func (m *ExchangeStartOkMountMessage) Serialize(w Writer) error {

	if err := m.ExchangeStartOkMountWithOutPaddockMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.PaddockedMountsDescription))); err != nil {
		return err
	}

	for i := range m.PaddockedMountsDescription {

		if err := m.PaddockedMountsDescription[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeStartOkMountMessage) Deserialize(r Reader) error {

	if err := m.ExchangeStartOkMountWithOutPaddockMessage.Deserialize(r); err != nil {
		return err
	}

	lpaddockedMountsDescriptionLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PaddockedMountsDescription = make([]MountClientData, lpaddockedMountsDescriptionLen)

	for i := range m.PaddockedMountsDescription {

		var lpaddockedMountsDescription MountClientData

		lpaddockedMountsDescription.Deserialize(r)

		m.PaddockedMountsDescription[i] = lpaddockedMountsDescription

	}

	return nil
}

type MountReleaseRequestMessage struct {
}

func (m *MountReleaseRequestMessage) ID() uint16 {
	return 5980
}

func (m *MountReleaseRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *MountReleaseRequestMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeMountStableErrorMessage struct {
}

func (m *ExchangeMountStableErrorMessage) ID() uint16 {
	return 5981
}

func (m *ExchangeMountStableErrorMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeMountStableErrorMessage) Deserialize(r Reader) error {

	return nil
}

type MountUnSetMessage struct {
}

func (m *MountUnSetMessage) ID() uint16 {
	return 5982
}

func (m *MountUnSetMessage) Serialize(w Writer) error {

	return nil
}

func (m *MountUnSetMessage) Deserialize(r Reader) error {

	return nil
}

type MountRenamedMessage struct {
	MountId int32

	Name string
}

func (m *MountRenamedMessage) ID() uint16 {
	return 5983
}

func (m *MountRenamedMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt32(m.MountId); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MountRenamedMessage) Deserialize(r Reader) error {

	lmountId, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.MountId = lmountId

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	return nil
}

type ExchangeStartedMountStockMessage struct {
	ObjectsInfos []ObjectItem
}

func (m *ExchangeStartedMountStockMessage) ID() uint16 {
	return 5984
}

func (m *ExchangeStartedMountStockMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ObjectsInfos))); err != nil {
		return err
	}

	for i := range m.ObjectsInfos {

		if err := m.ObjectsInfos[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeStartedMountStockMessage) Deserialize(r Reader) error {

	lobjectsInfosLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectsInfos = make([]ObjectItem, lobjectsInfosLen)

	for i := range m.ObjectsInfos {

		var lobjectsInfos ObjectItem

		lobjectsInfos.Deserialize(r)

		m.ObjectsInfos[i] = lobjectsInfos

	}

	return nil
}

type ExchangeRequestOnMountStockMessage struct {
}

func (m *ExchangeRequestOnMountStockMessage) ID() uint16 {
	return 5986
}

func (m *ExchangeRequestOnMountStockMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeRequestOnMountStockMessage) Deserialize(r Reader) error {

	return nil
}

type MountRenameRequestMessage struct {
	Name string

	MountId int32
}

func (m *MountRenameRequestMessage) ID() uint16 {
	return 5987
}

func (m *MountRenameRequestMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteVarInt32(m.MountId); err != nil {
		return err
	}

	return nil
}

func (m *MountRenameRequestMessage) Deserialize(r Reader) error {

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	lmountId, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.MountId = lmountId

	return nil
}

type MountSetXpRatioRequestMessage struct {
	XpRatio uint8
}

func (m *MountSetXpRatioRequestMessage) ID() uint16 {
	return 5989
}

func (m *MountSetXpRatioRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.XpRatio); err != nil {
		return err
	}

	return nil
}

func (m *MountSetXpRatioRequestMessage) Deserialize(r Reader) error {

	lxpRatio, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.XpRatio = lxpRatio

	return nil
}

type GameDataPaddockObjectAddMessage struct {
	PaddockItemDescription PaddockItem
}

func (m *GameDataPaddockObjectAddMessage) ID() uint16 {
	return 5990
}

func (m *GameDataPaddockObjectAddMessage) Serialize(w Writer) error {

	if err := m.PaddockItemDescription.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameDataPaddockObjectAddMessage) Deserialize(r Reader) error {

	var lpaddockItemDescription PaddockItem

	lpaddockItemDescription.Deserialize(r)

	m.PaddockItemDescription = lpaddockItemDescription

	return nil
}

type ExchangeStartOkMountWithOutPaddockMessage struct {
	StabledMountsDescription []MountClientData
}

func (m *ExchangeStartOkMountWithOutPaddockMessage) ID() uint16 {
	return 5991
}

func (m *ExchangeStartOkMountWithOutPaddockMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.StabledMountsDescription))); err != nil {
		return err
	}

	for i := range m.StabledMountsDescription {

		if err := m.StabledMountsDescription[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeStartOkMountWithOutPaddockMessage) Deserialize(r Reader) error {

	lstabledMountsDescriptionLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.StabledMountsDescription = make([]MountClientData, lstabledMountsDescriptionLen)

	for i := range m.StabledMountsDescription {

		var lstabledMountsDescription MountClientData

		lstabledMountsDescription.Deserialize(r)

		m.StabledMountsDescription[i] = lstabledMountsDescription

	}

	return nil
}

type GameDataPaddockObjectListAddMessage struct {
	PaddockItemDescription []PaddockItem
}

func (m *GameDataPaddockObjectListAddMessage) ID() uint16 {
	return 5992
}

func (m *GameDataPaddockObjectListAddMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.PaddockItemDescription))); err != nil {
		return err
	}

	for i := range m.PaddockItemDescription {

		if err := m.PaddockItemDescription[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameDataPaddockObjectListAddMessage) Deserialize(r Reader) error {

	lpaddockItemDescriptionLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PaddockItemDescription = make([]PaddockItem, lpaddockItemDescriptionLen)

	for i := range m.PaddockItemDescription {

		var lpaddockItemDescription PaddockItem

		lpaddockItemDescription.Deserialize(r)

		m.PaddockItemDescription[i] = lpaddockItemDescription

	}

	return nil
}

type GameDataPaddockObjectRemoveMessage struct {
	CellId uint16
}

func (m *GameDataPaddockObjectRemoveMessage) ID() uint16 {
	return 5993
}

func (m *GameDataPaddockObjectRemoveMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *GameDataPaddockObjectRemoveMessage) Deserialize(r Reader) error {

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type GameRolePlayPlayerLifeStatusMessage struct {
	State uint8

	PhenixMapId uint32
}

func (m *GameRolePlayPlayerLifeStatusMessage) ID() uint16 {
	return 5996
}

func (m *GameRolePlayPlayerLifeStatusMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.State); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.PhenixMapId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayPlayerLifeStatusMessage) Deserialize(r Reader) error {

	lstate, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.State = lstate

	lphenixMapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.PhenixMapId = lphenixMapId

	return nil
}

type ExchangeCraftResultWithObjectDescMessage struct {
	ExchangeCraftResultMessage

	ObjectInfo ObjectItemNotInContainer
}

func (m *ExchangeCraftResultWithObjectDescMessage) ID() uint16 {
	return 5999
}

func (m *ExchangeCraftResultWithObjectDescMessage) Serialize(w Writer) error {

	if err := m.ExchangeCraftResultMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.ObjectInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeCraftResultWithObjectDescMessage) Deserialize(r Reader) error {

	if err := m.ExchangeCraftResultMessage.Deserialize(r); err != nil {
		return err
	}

	var lobjectInfo ObjectItemNotInContainer

	lobjectInfo.Deserialize(r)

	m.ObjectInfo = lobjectInfo

	return nil
}

type ExchangeCraftResultWithObjectIdMessage struct {
	ExchangeCraftResultMessage

	ObjectGenericId uint16
}

func (m *ExchangeCraftResultWithObjectIdMessage) ID() uint16 {
	return 6000
}

func (m *ExchangeCraftResultWithObjectIdMessage) Serialize(w Writer) error {

	if err := m.ExchangeCraftResultMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGenericId); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeCraftResultWithObjectIdMessage) Deserialize(r Reader) error {

	if err := m.ExchangeCraftResultMessage.Deserialize(r); err != nil {
		return err
	}

	lobjectGenericId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGenericId = lobjectGenericId

	return nil
}

type ExchangeReplayStopMessage struct {
}

func (m *ExchangeReplayStopMessage) ID() uint16 {
	return 6001
}

func (m *ExchangeReplayStopMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeReplayStopMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeObjectUseInWorkshopMessage struct {
	ObjectUID uint32

	Quantity int32
}

func (m *ExchangeObjectUseInWorkshopMessage) ID() uint16 {
	return 6004
}

func (m *ExchangeObjectUseInWorkshopMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteVarInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectUseInWorkshopMessage) Deserialize(r Reader) error {

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	lquantity, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type ExchangeObjectModifiedInBagMessage struct {
	ExchangeObjectMessage

	Object ObjectItem
}

func (m *ExchangeObjectModifiedInBagMessage) ID() uint16 {
	return 6008
}

func (m *ExchangeObjectModifiedInBagMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.Object.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectModifiedInBagMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMessage.Deserialize(r); err != nil {
		return err
	}

	var lobject ObjectItem

	lobject.Deserialize(r)

	m.Object = lobject

	return nil
}

type ExchangeObjectPutInBagMessage struct {
	ExchangeObjectMessage

	Object ObjectItem
}

func (m *ExchangeObjectPutInBagMessage) ID() uint16 {
	return 6009
}

func (m *ExchangeObjectPutInBagMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.Object.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectPutInBagMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMessage.Deserialize(r); err != nil {
		return err
	}

	var lobject ObjectItem

	lobject.Deserialize(r)

	m.Object = lobject

	return nil
}

type ExchangeObjectRemovedFromBagMessage struct {
	ExchangeObjectMessage

	ObjectUID uint32
}

func (m *ExchangeObjectRemovedFromBagMessage) ID() uint16 {
	return 6010
}

func (m *ExchangeObjectRemovedFromBagMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectRemovedFromBagMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMessage.Deserialize(r); err != nil {
		return err
	}

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	return nil
}

type SpellItemBoostMessage struct {
	StatId uint32

	SpellId uint16

	Value int16
}

func (m *SpellItemBoostMessage) ID() uint16 {
	return 6011
}

func (m *SpellItemBoostMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.StatId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *SpellItemBoostMessage) Deserialize(r Reader) error {

	lstatId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.StatId = lstatId

	lspellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	lvalue, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type PauseDialogMessage struct {
	DialogType uint8
}

func (m *PauseDialogMessage) ID() uint16 {
	return 6012
}

func (m *PauseDialogMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.DialogType); err != nil {
		return err
	}

	return nil
}

func (m *PauseDialogMessage) Deserialize(r Reader) error {

	ldialogType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.DialogType = ldialogType

	return nil
}

type CompassUpdatePvpSeekMessage struct {
	CompassUpdateMessage

	MemberId int64

	MemberName string
}

func (m *CompassUpdatePvpSeekMessage) ID() uint16 {
	return 6013
}

func (m *CompassUpdatePvpSeekMessage) Serialize(w Writer) error {

	if err := m.CompassUpdateMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.MemberId); err != nil {
		return err
	}

	if err := w.WriteString(m.MemberName); err != nil {
		return err
	}

	return nil
}

func (m *CompassUpdatePvpSeekMessage) Deserialize(r Reader) error {

	if err := m.CompassUpdateMessage.Deserialize(r); err != nil {
		return err
	}

	lmemberId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.MemberId = lmemberId

	lmemberName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.MemberName = lmemberName

	return nil
}

type ObjectJobAddedMessage struct {
	JobId uint8
}

func (m *ObjectJobAddedMessage) ID() uint16 {
	return 6014
}

func (m *ObjectJobAddedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.JobId); err != nil {
		return err
	}

	return nil
}

func (m *ObjectJobAddedMessage) Deserialize(r Reader) error {

	ljobId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.JobId = ljobId

	return nil
}

type PaddockSellBuyDialogMessage struct {
	Bsell bool

	OwnerId uint32

	Price int64
}

func (m *PaddockSellBuyDialogMessage) ID() uint16 {
	return 6018
}

func (m *PaddockSellBuyDialogMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Bsell); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.OwnerId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Price); err != nil {
		return err
	}

	return nil
}

func (m *PaddockSellBuyDialogMessage) Deserialize(r Reader) error {

	lbsell, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Bsell = lbsell

	lownerId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.OwnerId = lownerId

	lprice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Price = lprice

	return nil
}

type ChallengeResultMessage struct {
	ChallengeId uint16

	Success bool
}

func (m *ChallengeResultMessage) ID() uint16 {
	return 6019
}

func (m *ChallengeResultMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ChallengeId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Success); err != nil {
		return err
	}

	return nil
}

func (m *ChallengeResultMessage) Deserialize(r Reader) error {

	lchallengeId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ChallengeId = lchallengeId

	lsuccess, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Success = lsuccess

	return nil
}

type ExchangeMultiCraftCrafterCanUseHisRessourcesMessage struct {
	Allowed bool
}

func (m *ExchangeMultiCraftCrafterCanUseHisRessourcesMessage) ID() uint16 {
	return 6020
}

func (m *ExchangeMultiCraftCrafterCanUseHisRessourcesMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Allowed); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeMultiCraftCrafterCanUseHisRessourcesMessage) Deserialize(r Reader) error {

	lallowed, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Allowed = lallowed

	return nil
}

type ExchangeMultiCraftSetCrafterCanUseHisRessourcesMessage struct {
	Allow bool
}

func (m *ExchangeMultiCraftSetCrafterCanUseHisRessourcesMessage) ID() uint16 {
	return 6021
}

func (m *ExchangeMultiCraftSetCrafterCanUseHisRessourcesMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Allow); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeMultiCraftSetCrafterCanUseHisRessourcesMessage) Deserialize(r Reader) error {

	lallow, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Allow = lallow

	return nil
}

type ChallengeInfoMessage struct {
	ChallengeId uint16

	TargetId float64

	XpBonus uint32

	DropBonus uint32
}

func (m *ChallengeInfoMessage) ID() uint16 {
	return 6022
}

func (m *ChallengeInfoMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ChallengeId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.XpBonus); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.DropBonus); err != nil {
		return err
	}

	return nil
}

func (m *ChallengeInfoMessage) Deserialize(r Reader) error {

	lchallengeId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ChallengeId = lchallengeId

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lxpBonus, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.XpBonus = lxpBonus

	ldropBonus, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.DropBonus = ldropBonus

	return nil
}

type GameContextCreateErrorMessage struct {
}

func (m *GameContextCreateErrorMessage) ID() uint16 {
	return 6024
}

func (m *GameContextCreateErrorMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameContextCreateErrorMessage) Deserialize(r Reader) error {

	return nil
}

type GameDataPlayFarmObjectAnimationMessage struct {
	CellId []uint16
}

func (m *GameDataPlayFarmObjectAnimationMessage) ID() uint16 {
	return 6026
}

func (m *GameDataPlayFarmObjectAnimationMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.CellId))); err != nil {
		return err
	}

	for i := range m.CellId {

		if err := w.WriteVarUInt16(m.CellId[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameDataPlayFarmObjectAnimationMessage) Deserialize(r Reader) error {

	lcellIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CellId = make([]uint16, lcellIdLen)

	for i := range m.CellId {

		lcellId, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.CellId[i] = lcellId

	}

	return nil
}

type DebugInClientMessage struct {
	Level uint8

	Message string
}

func (m *DebugInClientMessage) ID() uint16 {
	return 6028
}

func (m *DebugInClientMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Level); err != nil {
		return err
	}

	if err := w.WriteString(m.Message); err != nil {
		return err
	}

	return nil
}

func (m *DebugInClientMessage) Deserialize(r Reader) error {

	llevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Level = llevel

	lmessage, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Message = lmessage

	return nil
}

type AccountLoggingKickedMessage struct {
	Days uint16

	Hours uint8

	Minutes uint8
}

func (m *AccountLoggingKickedMessage) ID() uint16 {
	return 6029
}

func (m *AccountLoggingKickedMessage) Serialize(w Writer) error {

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

func (m *AccountLoggingKickedMessage) Deserialize(r Reader) error {

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

type GoldAddedMessage struct {
	Gold GoldItem
}

func (m *GoldAddedMessage) ID() uint16 {
	return 6030
}

func (m *GoldAddedMessage) Serialize(w Writer) error {

	if err := m.Gold.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GoldAddedMessage) Deserialize(r Reader) error {

	var lgold GoldItem

	lgold.Deserialize(r)

	m.Gold = lgold

	return nil
}

type ExchangeObjectTransfertAllToInvMessage struct {
}

func (m *ExchangeObjectTransfertAllToInvMessage) ID() uint16 {
	return 6032
}

func (m *ExchangeObjectTransfertAllToInvMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeObjectTransfertAllToInvMessage) Deserialize(r Reader) error {

	return nil
}

type ObjectsAddedMessage struct {
	Object []ObjectItem
}

func (m *ObjectsAddedMessage) ID() uint16 {
	return 6033
}

func (m *ObjectsAddedMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Object))); err != nil {
		return err
	}

	for i := range m.Object {

		if err := m.Object[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ObjectsAddedMessage) Deserialize(r Reader) error {

	lobjectLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Object = make([]ObjectItem, lobjectLen)

	for i := range m.Object {

		var lobject ObjectItem

		lobject.Deserialize(r)

		m.Object[i] = lobject

	}

	return nil
}

type ObjectsDeletedMessage struct {
	ObjectUID []uint32
}

func (m *ObjectsDeletedMessage) ID() uint16 {
	return 6034
}

func (m *ObjectsDeletedMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ObjectUID))); err != nil {
		return err
	}

	for i := range m.ObjectUID {

		if err := w.WriteVarUInt32(m.ObjectUID[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ObjectsDeletedMessage) Deserialize(r Reader) error {

	lobjectUIDLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectUID = make([]uint32, lobjectUIDLen)

	for i := range m.ObjectUID {

		lobjectUID, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.ObjectUID[i] = lobjectUID

	}

	return nil
}

type StorageObjectsRemoveMessage struct {
	ObjectUIDList []uint32
}

func (m *StorageObjectsRemoveMessage) ID() uint16 {
	return 6035
}

func (m *StorageObjectsRemoveMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ObjectUIDList))); err != nil {
		return err
	}

	for i := range m.ObjectUIDList {

		if err := w.WriteVarUInt32(m.ObjectUIDList[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *StorageObjectsRemoveMessage) Deserialize(r Reader) error {

	lobjectUIDListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectUIDList = make([]uint32, lobjectUIDListLen)

	for i := range m.ObjectUIDList {

		lobjectUIDList, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.ObjectUIDList[i] = lobjectUIDList

	}

	return nil
}

type StorageObjectsUpdateMessage struct {
	ObjectList []ObjectItem
}

func (m *StorageObjectsUpdateMessage) ID() uint16 {
	return 6036
}

func (m *StorageObjectsUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ObjectList))); err != nil {
		return err
	}

	for i := range m.ObjectList {

		if err := m.ObjectList[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *StorageObjectsUpdateMessage) Deserialize(r Reader) error {

	lobjectListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectList = make([]ObjectItem, lobjectListLen)

	for i := range m.ObjectList {

		var lobjectList ObjectItem

		lobjectList.Deserialize(r)

		m.ObjectList[i] = lobjectList

	}

	return nil
}

type ExchangeShopStockMultiMovementRemovedMessage struct {
	ObjectIdList []uint32
}

func (m *ExchangeShopStockMultiMovementRemovedMessage) ID() uint16 {
	return 6037
}

func (m *ExchangeShopStockMultiMovementRemovedMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ObjectIdList))); err != nil {
		return err
	}

	for i := range m.ObjectIdList {

		if err := w.WriteVarUInt32(m.ObjectIdList[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeShopStockMultiMovementRemovedMessage) Deserialize(r Reader) error {

	lobjectIdListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectIdList = make([]uint32, lobjectIdListLen)

	for i := range m.ObjectIdList {

		lobjectIdList, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.ObjectIdList[i] = lobjectIdList

	}

	return nil
}

type ExchangeShopStockMultiMovementUpdatedMessage struct {
	ObjectInfoList []ObjectItemToSell
}

func (m *ExchangeShopStockMultiMovementUpdatedMessage) ID() uint16 {
	return 6038
}

func (m *ExchangeShopStockMultiMovementUpdatedMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ObjectInfoList))); err != nil {
		return err
	}

	for i := range m.ObjectInfoList {

		if err := m.ObjectInfoList[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeShopStockMultiMovementUpdatedMessage) Deserialize(r Reader) error {

	lobjectInfoListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectInfoList = make([]ObjectItemToSell, lobjectInfoListLen)

	for i := range m.ObjectInfoList {

		var lobjectInfoList ObjectItemToSell

		lobjectInfoList.Deserialize(r)

		m.ObjectInfoList[i] = lobjectInfoList

	}

	return nil
}

type ExchangeObjectTransfertListToInvMessage struct {
	Ids []uint32
}

func (m *ExchangeObjectTransfertListToInvMessage) ID() uint16 {
	return 6039
}

func (m *ExchangeObjectTransfertListToInvMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Ids))); err != nil {
		return err
	}

	for i := range m.Ids {

		if err := w.WriteVarUInt32(m.Ids[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeObjectTransfertListToInvMessage) Deserialize(r Reader) error {

	lidsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Ids = make([]uint32, lidsLen)

	for i := range m.Ids {

		lids, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.Ids[i] = lids

	}

	return nil
}

type PrismFightStateUpdateMessage struct {
	State uint8
}

func (m *PrismFightStateUpdateMessage) ID() uint16 {
	return 6040
}

func (m *PrismFightStateUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.State); err != nil {
		return err
	}

	return nil
}

func (m *PrismFightStateUpdateMessage) Deserialize(r Reader) error {

	lstate, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.State = lstate

	return nil
}

type PrismUseRequestMessage struct {
	ModuleToUse uint8
}

func (m *PrismUseRequestMessage) ID() uint16 {
	return 6041
}

func (m *PrismUseRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.ModuleToUse); err != nil {
		return err
	}

	return nil
}

func (m *PrismUseRequestMessage) Deserialize(r Reader) error {

	lmoduleToUse, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ModuleToUse = lmoduleToUse

	return nil
}

type PrismAttackRequestMessage struct {
}

func (m *PrismAttackRequestMessage) ID() uint16 {
	return 6042
}

func (m *PrismAttackRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *PrismAttackRequestMessage) Deserialize(r Reader) error {

	return nil
}

type JobCrafterDirectoryEntryRequestMessage struct {
	PlayerId int64
}

func (m *JobCrafterDirectoryEntryRequestMessage) ID() uint16 {
	return 6043
}

func (m *JobCrafterDirectoryEntryRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *JobCrafterDirectoryEntryRequestMessage) Deserialize(r Reader) error {

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type JobCrafterDirectoryEntryMessage struct {
	PlayerInfo JobCrafterDirectoryEntryPlayerInfo

	JobInfoList []JobCrafterDirectoryEntryJobInfo

	PlayerLook EntityLook
}

func (m *JobCrafterDirectoryEntryMessage) ID() uint16 {
	return 6044
}

func (m *JobCrafterDirectoryEntryMessage) Serialize(w Writer) error {

	if err := m.PlayerInfo.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.JobInfoList))); err != nil {
		return err
	}

	for i := range m.JobInfoList {

		if err := m.JobInfoList[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := m.PlayerLook.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *JobCrafterDirectoryEntryMessage) Deserialize(r Reader) error {

	var lplayerInfo JobCrafterDirectoryEntryPlayerInfo

	lplayerInfo.Deserialize(r)

	m.PlayerInfo = lplayerInfo

	ljobInfoListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.JobInfoList = make([]JobCrafterDirectoryEntryJobInfo, ljobInfoListLen)

	for i := range m.JobInfoList {

		var ljobInfoList JobCrafterDirectoryEntryJobInfo

		ljobInfoList.Deserialize(r)

		m.JobInfoList[i] = ljobInfoList

	}

	var lplayerLook EntityLook

	lplayerLook.Deserialize(r)

	m.PlayerLook = lplayerLook

	return nil
}

type ContactLookErrorMessage struct {
	RequestId uint32
}

func (m *ContactLookErrorMessage) ID() uint16 {
	return 6045
}

func (m *ContactLookErrorMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.RequestId); err != nil {
		return err
	}

	return nil
}

func (m *ContactLookErrorMessage) Deserialize(r Reader) error {

	lrequestId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.RequestId = lrequestId

	return nil
}

type JobCrafterDirectoryListMessage struct {
	ListEntries []JobCrafterDirectoryListEntry
}

func (m *JobCrafterDirectoryListMessage) ID() uint16 {
	return 6046
}

func (m *JobCrafterDirectoryListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ListEntries))); err != nil {
		return err
	}

	for i := range m.ListEntries {

		if err := m.ListEntries[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *JobCrafterDirectoryListMessage) Deserialize(r Reader) error {

	llistEntriesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ListEntries = make([]JobCrafterDirectoryListEntry, llistEntriesLen)

	for i := range m.ListEntries {

		var llistEntries JobCrafterDirectoryListEntry

		llistEntries.Deserialize(r)

		m.ListEntries[i] = llistEntries

	}

	return nil
}

type JobCrafterDirectoryListRequestMessage struct {
	JobId uint8
}

func (m *JobCrafterDirectoryListRequestMessage) ID() uint16 {
	return 6047
}

func (m *JobCrafterDirectoryListRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.JobId); err != nil {
		return err
	}

	return nil
}

func (m *JobCrafterDirectoryListRequestMessage) Deserialize(r Reader) error {

	ljobId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.JobId = ljobId

	return nil
}

type TeleportOnSameMapMessage struct {
	TargetId float64

	CellId uint16
}

func (m *TeleportOnSameMapMessage) ID() uint16 {
	return 6048
}

func (m *TeleportOnSameMapMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *TeleportOnSameMapMessage) Deserialize(r Reader) error {

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type MapObstacleUpdateMessage struct {
	Obstacles []MapObstacle
}

func (m *MapObstacleUpdateMessage) ID() uint16 {
	return 6051
}

func (m *MapObstacleUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Obstacles))); err != nil {
		return err
	}

	for i := range m.Obstacles {

		if err := m.Obstacles[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *MapObstacleUpdateMessage) Deserialize(r Reader) error {

	lobstaclesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Obstacles = make([]MapObstacle, lobstaclesLen)

	for i := range m.Obstacles {

		var lobstacles MapObstacle

		lobstacles.Deserialize(r)

		m.Obstacles[i] = lobstacles

	}

	return nil
}

type PaddockMoveItemRequestMessage struct {
	OldCellId uint16

	NewCellId uint16
}

func (m *PaddockMoveItemRequestMessage) ID() uint16 {
	return 6052
}

func (m *PaddockMoveItemRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.OldCellId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NewCellId); err != nil {
		return err
	}

	return nil
}

func (m *PaddockMoveItemRequestMessage) Deserialize(r Reader) error {

	loldCellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.OldCellId = loldCellId

	lnewCellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NewCellId = lnewCellId

	return nil
}

type CinematicMessage struct {
	CinematicId uint16
}

func (m *CinematicMessage) ID() uint16 {
	return 6053
}

func (m *CinematicMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CinematicId); err != nil {
		return err
	}

	return nil
}

func (m *CinematicMessage) Deserialize(r Reader) error {

	lcinematicId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CinematicId = lcinematicId

	return nil
}

type PartyUpdateLightMessage struct {
	AbstractPartyEventMessage

	Id int64

	LifePoints uint32

	MaxLifePoints uint32

	Prospecting uint16

	RegenRate uint8
}

func (m *PartyUpdateLightMessage) ID() uint16 {
	return 6054
}

func (m *PartyUpdateLightMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyEventMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Id); err != nil {
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

func (m *PartyUpdateLightMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyEventMessage.Deserialize(r); err != nil {
		return err
	}

	lid, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Id = lid

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

type ExchangeMountFreeFromPaddockMessage struct {
	Name string

	WorldX int16

	WorldY int16

	Liberator string
}

func (m *ExchangeMountFreeFromPaddockMessage) ID() uint16 {
	return 6055
}

func (m *ExchangeMountFreeFromPaddockMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteString(m.Liberator); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeMountFreeFromPaddockMessage) Deserialize(r Reader) error {

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

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

	lliberator, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Liberator = lliberator

	return nil
}

type ExchangeMountSterilizeFromPaddockMessage struct {
	Name string

	WorldX int16

	WorldY int16

	Sterilizator string
}

func (m *ExchangeMountSterilizeFromPaddockMessage) ID() uint16 {
	return 6056
}

func (m *ExchangeMountSterilizeFromPaddockMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteString(m.Sterilizator); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeMountSterilizeFromPaddockMessage) Deserialize(r Reader) error {

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

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

	lsterilizator, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Sterilizator = lsterilizator

	return nil
}

type AlignmentRankUpdateMessage struct {
	AlignmentRank uint8

	Verbose bool
}

func (m *AlignmentRankUpdateMessage) ID() uint16 {
	return 6058
}

func (m *AlignmentRankUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.AlignmentRank); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Verbose); err != nil {
		return err
	}

	return nil
}

func (m *AlignmentRankUpdateMessage) Deserialize(r Reader) error {

	lalignmentRank, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.AlignmentRank = lalignmentRank

	lverbose, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Verbose = lverbose

	return nil
}

type GuildMemberOnlineStatusMessage struct {
	MemberId int64

	Online bool
}

func (m *GuildMemberOnlineStatusMessage) ID() uint16 {
	return 6061
}

func (m *GuildMemberOnlineStatusMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.MemberId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Online); err != nil {
		return err
	}

	return nil
}

func (m *GuildMemberOnlineStatusMessage) Deserialize(r Reader) error {

	lmemberId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.MemberId = lmemberId

	lonline, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Online = lonline

	return nil
}

type GuildLevelUpMessage struct {
	NewLevel uint8
}

func (m *GuildLevelUpMessage) ID() uint16 {
	return 6062
}

func (m *GuildLevelUpMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.NewLevel); err != nil {
		return err
	}

	return nil
}

func (m *GuildLevelUpMessage) Deserialize(r Reader) error {

	lnewLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NewLevel = lnewLevel

	return nil
}

type LivingObjectMessageMessage struct {
	MsgId uint16

	TimeStamp uint32

	Owner string

	ObjectGenericId uint16
}

func (m *LivingObjectMessageMessage) ID() uint16 {
	return 6065
}

func (m *LivingObjectMessageMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.MsgId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.TimeStamp); err != nil {
		return err
	}

	if err := w.WriteString(m.Owner); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGenericId); err != nil {
		return err
	}

	return nil
}

func (m *LivingObjectMessageMessage) Deserialize(r Reader) error {

	lmsgId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MsgId = lmsgId

	ltimeStamp, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.TimeStamp = ltimeStamp

	lowner, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Owner = lowner

	lobjectGenericId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGenericId = lobjectGenericId

	return nil
}

type LivingObjectMessageRequestMessage struct {
	MsgId uint16

	Parameters []string

	LivingObject uint32
}

func (m *LivingObjectMessageRequestMessage) ID() uint16 {
	return 6066
}

func (m *LivingObjectMessageRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.MsgId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Parameters))); err != nil {
		return err
	}

	for i := range m.Parameters {

		if err := w.WriteString(m.Parameters[i]); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt32(m.LivingObject); err != nil {
		return err
	}

	return nil
}

func (m *LivingObjectMessageRequestMessage) Deserialize(r Reader) error {

	lmsgId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.MsgId = lmsgId

	lparametersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Parameters = make([]string, lparametersLen)

	for i := range m.Parameters {

		lparameters, err := r.ReadString()
		if err != nil {
			return err
		}

		m.Parameters[i] = lparameters

	}

	llivingObject, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.LivingObject = llivingObject

	return nil
}

type GameFightResumeMessage struct {
	GameFightSpectateMessage

	SpellCooldowns []GameFightSpellCooldown

	SummonCount uint8

	BombCount uint8
}

func (m *GameFightResumeMessage) ID() uint16 {
	return 6067
}

func (m *GameFightResumeMessage) Serialize(w Writer) error {

	if err := m.GameFightSpectateMessage.Serialize(w); err != nil {
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

func (m *GameFightResumeMessage) Deserialize(r Reader) error {

	if err := m.GameFightSpectateMessage.Deserialize(r); err != nil {
		return err
	}

	lspellCooldownsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SpellCooldowns = make([]GameFightSpellCooldown, lspellCooldownsLen)

	for i := range m.SpellCooldowns {

		var lspellCooldowns GameFightSpellCooldown

		lspellCooldowns.Deserialize(r)

		m.SpellCooldowns[i] = lspellCooldowns

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

type CharacterSelectedForceMessage struct {
	Id int32
}

func (m *CharacterSelectedForceMessage) ID() uint16 {
	return 6068
}

func (m *CharacterSelectedForceMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *CharacterSelectedForceMessage) Deserialize(r Reader) error {

	lid, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type GameFightSpectateMessage struct {
	Effects []FightDispellableEffectExtendedInformations

	Marks []GameActionMark

	GameTurn uint16

	FightStart uint32

	Idols []Idol
}

func (m *GameFightSpectateMessage) ID() uint16 {
	return 6069
}

func (m *GameFightSpectateMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Effects))); err != nil {
		return err
	}

	for i := range m.Effects {

		if err := m.Effects[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Marks))); err != nil {
		return err
	}

	for i := range m.Marks {

		if err := m.Marks[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt16(m.GameTurn); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.FightStart); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Idols))); err != nil {
		return err
	}

	for i := range m.Idols {

		if err := m.Idols[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameFightSpectateMessage) Deserialize(r Reader) error {

	leffectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Effects = make([]FightDispellableEffectExtendedInformations, leffectsLen)

	for i := range m.Effects {

		var leffects FightDispellableEffectExtendedInformations

		leffects.Deserialize(r)

		m.Effects[i] = leffects

	}

	lmarksLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Marks = make([]GameActionMark, lmarksLen)

	for i := range m.Marks {

		var lmarks GameActionMark

		lmarks.Deserialize(r)

		m.Marks[i] = lmarks

	}

	lgameTurn, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.GameTurn = lgameTurn

	lfightStart, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.FightStart = lfightStart

	lidolsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Idols = make([]Idol, lidolsLen)

	for i := range m.Idols {

		var lidols Idol

		lidols.Deserialize(r)

		m.Idols[i] = lidols

	}

	return nil
}

type GameActionFightDispellableEffectMessage struct {
	AbstractGameActionMessage

	Effect AbstractFightDispellableEffectIntrf
}

func (m *GameActionFightDispellableEffectMessage) ID() uint16 {
	return 6070
}

func (m *GameActionFightDispellableEffectMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
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

func (m *GameActionFightDispellableEffectMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	typeeffectID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	leffect, err := GetType(typeeffectID)
	if err != nil {
		return err
	}

	leffect.Deserialize(r)

	m.Effect = leffect.(AbstractFightDispellableEffectIntrf)

	return nil
}

type GameContextReadyMessage struct {
	MapId uint32
}

func (m *GameContextReadyMessage) ID() uint16 {
	return 6071
}

func (m *GameContextReadyMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.MapId); err != nil {
		return err
	}

	return nil
}

func (m *GameContextReadyMessage) Deserialize(r Reader) error {

	lmapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	return nil
}

type CharacterSelectedForceReadyMessage struct {
}

func (m *CharacterSelectedForceReadyMessage) ID() uint16 {
	return 6072
}

func (m *CharacterSelectedForceReadyMessage) Serialize(w Writer) error {

	return nil
}

func (m *CharacterSelectedForceReadyMessage) Deserialize(r Reader) error {

	return nil
}

type GameRolePlayAggressionMessage struct {
	AttackerId int64

	DefenderId int64
}

func (m *GameRolePlayAggressionMessage) ID() uint16 {
	return 6073
}

func (m *GameRolePlayAggressionMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.AttackerId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.DefenderId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayAggressionMessage) Deserialize(r Reader) error {

	lattackerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.AttackerId = lattackerId

	ldefenderId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.DefenderId = ldefenderId

	return nil
}

type CharacterLevelUpInformationMessage struct {
	CharacterLevelUpMessage

	Name string

	Id int64
}

func (m *CharacterLevelUpInformationMessage) ID() uint16 {
	return 6076
}

func (m *CharacterLevelUpInformationMessage) Serialize(w Writer) error {

	if err := m.CharacterLevelUpMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *CharacterLevelUpInformationMessage) Deserialize(r Reader) error {

	if err := m.CharacterLevelUpMessage.Deserialize(r); err != nil {
		return err
	}

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	lid, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type FriendSetWarnOnLevelGainMessage struct {
	Enable bool
}

func (m *FriendSetWarnOnLevelGainMessage) ID() uint16 {
	return 6077
}

func (m *FriendSetWarnOnLevelGainMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *FriendSetWarnOnLevelGainMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type FriendWarnOnLevelGainStateMessage struct {
	Enable bool
}

func (m *FriendWarnOnLevelGainStateMessage) ID() uint16 {
	return 6078
}

func (m *FriendWarnOnLevelGainStateMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *FriendWarnOnLevelGainStateMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type CharacterReportMessage struct {
	ReportedId int64

	Reason uint8
}

func (m *CharacterReportMessage) ID() uint16 {
	return 6079
}

func (m *CharacterReportMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.ReportedId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *CharacterReportMessage) Deserialize(r Reader) error {

	lreportedId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ReportedId = lreportedId

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type PartyAbdicateThroneMessage struct {
	AbstractPartyMessage

	PlayerId int64
}

func (m *PartyAbdicateThroneMessage) ID() uint16 {
	return 6080
}

func (m *PartyAbdicateThroneMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *PartyAbdicateThroneMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type GameContextKickMessage struct {
	TargetId float64
}

func (m *GameContextKickMessage) ID() uint16 {
	return 6081
}

func (m *GameContextKickMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *GameContextKickMessage) Deserialize(r Reader) error {

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type CharacterFirstSelectionMessage struct {
	CharacterSelectionMessage

	DoTutorial bool
}

func (m *CharacterFirstSelectionMessage) ID() uint16 {
	return 6084
}

func (m *CharacterFirstSelectionMessage) Serialize(w Writer) error {

	if err := m.CharacterSelectionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.DoTutorial); err != nil {
		return err
	}

	return nil
}

func (m *CharacterFirstSelectionMessage) Deserialize(r Reader) error {

	if err := m.CharacterSelectionMessage.Deserialize(r); err != nil {
		return err
	}

	ldoTutorial, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.DoTutorial = ldoTutorial

	return nil
}

type QuestObjectiveValidationMessage struct {
	QuestId uint16

	ObjectiveId uint16
}

func (m *QuestObjectiveValidationMessage) ID() uint16 {
	return 6085
}

func (m *QuestObjectiveValidationMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.QuestId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectiveId); err != nil {
		return err
	}

	return nil
}

func (m *QuestObjectiveValidationMessage) Deserialize(r Reader) error {

	lquestId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.QuestId = lquestId

	lobjectiveId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectiveId = lobjectiveId

	return nil
}

type NotificationListMessage struct {
	Flags []int32
}

func (m *NotificationListMessage) ID() uint16 {
	return 6087
}

func (m *NotificationListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Flags))); err != nil {
		return err
	}

	for i := range m.Flags {

		if err := w.WriteVarInt32(m.Flags[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *NotificationListMessage) Deserialize(r Reader) error {

	lflagsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Flags = make([]int32, lflagsLen)

	for i := range m.Flags {

		lflags, err := r.ReadVarInt32()
		if err != nil {
			return err
		}

		m.Flags[i] = lflags

	}

	return nil
}

type GuidedModeReturnRequestMessage struct {
}

func (m *GuidedModeReturnRequestMessage) ID() uint16 {
	return 6088
}

func (m *GuidedModeReturnRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *GuidedModeReturnRequestMessage) Deserialize(r Reader) error {

	return nil
}

type NotificationResetMessage struct {
}

func (m *NotificationResetMessage) ID() uint16 {
	return 6089
}

func (m *NotificationResetMessage) Serialize(w Writer) error {

	return nil
}

func (m *NotificationResetMessage) Deserialize(r Reader) error {

	return nil
}

type NotificationUpdateFlagMessage struct {
	Index uint16
}

func (m *NotificationUpdateFlagMessage) ID() uint16 {
	return 6090
}

func (m *NotificationUpdateFlagMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Index); err != nil {
		return err
	}

	return nil
}

func (m *NotificationUpdateFlagMessage) Deserialize(r Reader) error {

	lindex, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Index = lindex

	return nil
}

type QuestStartedMessage struct {
	QuestId uint16
}

func (m *QuestStartedMessage) ID() uint16 {
	return 6091
}

func (m *QuestStartedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.QuestId); err != nil {
		return err
	}

	return nil
}

func (m *QuestStartedMessage) Deserialize(r Reader) error {

	lquestId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.QuestId = lquestId

	return nil
}

type GuidedModeQuitRequestMessage struct {
}

func (m *GuidedModeQuitRequestMessage) ID() uint16 {
	return 6092
}

func (m *GuidedModeQuitRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *GuidedModeQuitRequestMessage) Deserialize(r Reader) error {

	return nil
}

type QuestStepStartedMessage struct {
	QuestId uint16

	StepId uint16
}

func (m *QuestStepStartedMessage) ID() uint16 {
	return 6096
}

func (m *QuestStepStartedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.QuestId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.StepId); err != nil {
		return err
	}

	return nil
}

func (m *QuestStepStartedMessage) Deserialize(r Reader) error {

	lquestId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.QuestId = lquestId

	lstepId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.StepId = lstepId

	return nil
}

type QuestValidatedMessage struct {
	QuestId uint16
}

func (m *QuestValidatedMessage) ID() uint16 {
	return 6097
}

func (m *QuestValidatedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.QuestId); err != nil {
		return err
	}

	return nil
}

func (m *QuestValidatedMessage) Deserialize(r Reader) error {

	lquestId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.QuestId = lquestId

	return nil
}

type QuestObjectiveValidatedMessage struct {
	QuestId uint16

	ObjectiveId uint16
}

func (m *QuestObjectiveValidatedMessage) ID() uint16 {
	return 6098
}

func (m *QuestObjectiveValidatedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.QuestId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectiveId); err != nil {
		return err
	}

	return nil
}

func (m *QuestObjectiveValidatedMessage) Deserialize(r Reader) error {

	lquestId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.QuestId = lquestId

	lobjectiveId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectiveId = lobjectiveId

	return nil
}

type QuestStepValidatedMessage struct {
	QuestId uint16

	StepId uint16
}

func (m *QuestStepValidatedMessage) ID() uint16 {
	return 6099
}

func (m *QuestStepValidatedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.QuestId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.StepId); err != nil {
		return err
	}

	return nil
}

func (m *QuestStepValidatedMessage) Deserialize(r Reader) error {

	lquestId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.QuestId = lquestId

	lstepId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.StepId = lstepId

	return nil
}

type QueueStatusMessage struct {
	Position uint16

	Total uint16
}

func (m *QueueStatusMessage) ID() uint16 {
	return 6100
}

func (m *QueueStatusMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Position); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Total); err != nil {
		return err
	}

	return nil
}

func (m *QueueStatusMessage) Deserialize(r Reader) error {

	lposition, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.Position = lposition

	ltotal, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.Total = ltotal

	return nil
}

type NotificationByServerMessage struct {
	Id uint16

	Parameters []string

	ForceOpen bool
}

func (m *NotificationByServerMessage) ID() uint16 {
	return 6103
}

func (m *NotificationByServerMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Id); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Parameters))); err != nil {
		return err
	}

	for i := range m.Parameters {

		if err := w.WriteString(m.Parameters[i]); err != nil {
			return err
		}

	}

	if err := w.WriteBoolean(m.ForceOpen); err != nil {
		return err
	}

	return nil
}

func (m *NotificationByServerMessage) Deserialize(r Reader) error {

	lid, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Id = lid

	lparametersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Parameters = make([]string, lparametersLen)

	for i := range m.Parameters {

		lparameters, err := r.ReadString()
		if err != nil {
			return err
		}

		m.Parameters[i] = lparameters

	}

	lforceOpen, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.ForceOpen = lforceOpen

	return nil
}

type EntityTalkMessage struct {
	EntityId float64

	TextId uint16

	Parameters []string
}

func (m *EntityTalkMessage) ID() uint16 {
	return 6110
}

func (m *EntityTalkMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.EntityId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TextId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Parameters))); err != nil {
		return err
	}

	for i := range m.Parameters {

		if err := w.WriteString(m.Parameters[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *EntityTalkMessage) Deserialize(r Reader) error {

	lentityId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.EntityId = lentityId

	ltextId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TextId = ltextId

	lparametersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Parameters = make([]string, lparametersLen)

	for i := range m.Parameters {

		lparameters, err := r.ReadString()
		if err != nil {
			return err
		}

		m.Parameters[i] = lparameters

	}

	return nil
}

type InteractiveUseEndedMessage struct {
	ElemId uint32

	SkillId uint16
}

func (m *InteractiveUseEndedMessage) ID() uint16 {
	return 6112
}

func (m *InteractiveUseEndedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ElemId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SkillId); err != nil {
		return err
	}

	return nil
}

func (m *InteractiveUseEndedMessage) Deserialize(r Reader) error {

	lelemId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ElemId = lelemId

	lskillId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SkillId = lskillId

	return nil
}

type GameActionFightDispellEffectMessage struct {
	GameActionFightDispellMessage

	BoostUID uint32
}

func (m *GameActionFightDispellEffectMessage) ID() uint16 {
	return 6113
}

func (m *GameActionFightDispellEffectMessage) Serialize(w Writer) error {

	if err := m.GameActionFightDispellMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.BoostUID); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightDispellEffectMessage) Deserialize(r Reader) error {

	if err := m.GameActionFightDispellMessage.Deserialize(r); err != nil {
		return err
	}

	lboostUID, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.BoostUID = lboostUID

	return nil
}

type GameRolePlaySpellAnimMessage struct {
	CasterId int64

	TargetCellId uint16

	SpellId uint16

	SpellLevel int16
}

func (m *GameRolePlaySpellAnimMessage) ID() uint16 {
	return 6114
}

func (m *GameRolePlaySpellAnimMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.CasterId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TargetCellId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.SpellLevel); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlaySpellAnimMessage) Deserialize(r Reader) error {

	lcasterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.CasterId = lcasterId

	ltargetCellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TargetCellId = ltargetCellId

	lspellId, err := r.ReadVarUInt16()
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

type GuildInvitationByNameMessage struct {
	Name string
}

func (m *GuildInvitationByNameMessage) ID() uint16 {
	return 6115
}

func (m *GuildInvitationByNameMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GuildInvitationByNameMessage) Deserialize(r Reader) error {

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	return nil
}

type GameActionFightCloseCombatMessage struct {
	AbstractGameActionFightTargetedAbilityMessage

	WeaponGenericId uint16
}

func (m *GameActionFightCloseCombatMessage) ID() uint16 {
	return 6116
}

func (m *GameActionFightCloseCombatMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionFightTargetedAbilityMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.WeaponGenericId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightCloseCombatMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionFightTargetedAbilityMessage.Deserialize(r); err != nil {
		return err
	}

	lweaponGenericId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.WeaponGenericId = lweaponGenericId

	return nil
}

type AbstractGameActionFightTargetedAbilityMessage struct {
	AbstractGameActionMessage

	TargetId float64

	DestinationCellId int16

	Critical uint8

	SilentCast bool

	VerboseCast bool
}

func (m *AbstractGameActionFightTargetedAbilityMessage) ID() uint16 {
	return 6118
}

func (m *AbstractGameActionFightTargetedAbilityMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.SilentCast)

	setWrappedFlag(bbw0, 1, m.VerboseCast)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.DestinationCellId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Critical); err != nil {
		return err
	}

	return nil
}

func (m *AbstractGameActionFightTargetedAbilityMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.SilentCast = getWrappedFlag(bbw0, 0)

	m.VerboseCast = getWrappedFlag(bbw0, 1)

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	ldestinationCellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DestinationCellId = ldestinationCellId

	lcritical, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Critical = lcritical

	return nil
}

type IdentificationAccountForceMessage struct {
	IdentificationMessage

	ForcedAccountLogin string
}

func (m *IdentificationAccountForceMessage) ID() uint16 {
	return 6119
}

func (m *IdentificationAccountForceMessage) Serialize(w Writer) error {

	if err := m.IdentificationMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.ForcedAccountLogin); err != nil {
		return err
	}

	return nil
}

func (m *IdentificationAccountForceMessage) Deserialize(r Reader) error {

	if err := m.IdentificationMessage.Deserialize(r); err != nil {
		return err
	}

	lforcedAccountLogin, err := r.ReadString()
	if err != nil {
		return err
	}

	m.ForcedAccountLogin = lforcedAccountLogin

	return nil
}

type CharactersListWithModificationsMessage struct {
	CharactersListMessage

	CharactersToRecolor []CharacterToRecolorInformation

	CharactersToRename []int32

	UnusableCharacters []int32

	CharactersToRelook []CharacterToRelookInformation
}

func (m *CharactersListWithModificationsMessage) ID() uint16 {
	return 6120
}

func (m *CharactersListWithModificationsMessage) Serialize(w Writer) error {

	if err := m.CharactersListMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.CharactersToRecolor))); err != nil {
		return err
	}

	for i := range m.CharactersToRecolor {

		if err := m.CharactersToRecolor[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.CharactersToRename))); err != nil {
		return err
	}

	for i := range m.CharactersToRename {

		if err := w.WriteInt32(m.CharactersToRename[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.UnusableCharacters))); err != nil {
		return err
	}

	for i := range m.UnusableCharacters {

		if err := w.WriteInt32(m.UnusableCharacters[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.CharactersToRelook))); err != nil {
		return err
	}

	for i := range m.CharactersToRelook {

		if err := m.CharactersToRelook[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *CharactersListWithModificationsMessage) Deserialize(r Reader) error {

	if err := m.CharactersListMessage.Deserialize(r); err != nil {
		return err
	}

	lcharactersToRecolorLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CharactersToRecolor = make([]CharacterToRecolorInformation, lcharactersToRecolorLen)

	for i := range m.CharactersToRecolor {

		var lcharactersToRecolor CharacterToRecolorInformation

		lcharactersToRecolor.Deserialize(r)

		m.CharactersToRecolor[i] = lcharactersToRecolor

	}

	lcharactersToRenameLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CharactersToRename = make([]int32, lcharactersToRenameLen)

	for i := range m.CharactersToRename {

		lcharactersToRename, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.CharactersToRename[i] = lcharactersToRename

	}

	lunusableCharactersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.UnusableCharacters = make([]int32, lunusableCharactersLen)

	for i := range m.UnusableCharacters {

		lunusableCharacters, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.UnusableCharacters[i] = lunusableCharacters

	}

	lcharactersToRelookLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CharactersToRelook = make([]CharacterToRelookInformation, lcharactersToRelookLen)

	for i := range m.CharactersToRelook {

		var lcharactersToRelook CharacterToRelookInformation

		lcharactersToRelook.Deserialize(r)

		m.CharactersToRelook[i] = lcharactersToRelook

	}

	return nil
}

type ChallengeTargetUpdateMessage struct {
	ChallengeId uint16

	TargetId float64
}

func (m *ChallengeTargetUpdateMessage) ID() uint16 {
	return 6123
}

func (m *ChallengeTargetUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ChallengeId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *ChallengeTargetUpdateMessage) Deserialize(r Reader) error {

	lchallengeId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ChallengeId = lchallengeId

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type StopToListenRunningFightRequestMessage struct {
}

func (m *StopToListenRunningFightRequestMessage) ID() uint16 {
	return 6124
}

func (m *StopToListenRunningFightRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *StopToListenRunningFightRequestMessage) Deserialize(r Reader) error {

	return nil
}

type ConsoleCommandsListMessage struct {
	Aliases []string

	Args []string

	Descriptions []string
}

func (m *ConsoleCommandsListMessage) ID() uint16 {
	return 6127
}

func (m *ConsoleCommandsListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Aliases))); err != nil {
		return err
	}

	for i := range m.Aliases {

		if err := w.WriteString(m.Aliases[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Args))); err != nil {
		return err
	}

	for i := range m.Args {

		if err := w.WriteString(m.Args[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Descriptions))); err != nil {
		return err
	}

	for i := range m.Descriptions {

		if err := w.WriteString(m.Descriptions[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ConsoleCommandsListMessage) Deserialize(r Reader) error {

	laliasesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Aliases = make([]string, laliasesLen)

	for i := range m.Aliases {

		laliases, err := r.ReadString()
		if err != nil {
			return err
		}

		m.Aliases[i] = laliases

	}

	largsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Args = make([]string, largsLen)

	for i := range m.Args {

		largs, err := r.ReadString()
		if err != nil {
			return err
		}

		m.Args[i] = largs

	}

	ldescriptionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Descriptions = make([]string, ldescriptionsLen)

	for i := range m.Descriptions {

		ldescriptions, err := r.ReadString()
		if err != nil {
			return err
		}

		m.Descriptions[i] = ldescriptions

	}

	return nil
}

type ExchangeStartedWithPodsMessage struct {
	ExchangeStartedMessage

	FirstCharacterId float64

	FirstCharacterCurrentWeight uint32

	FirstCharacterMaxWeight uint32

	SecondCharacterId float64

	SecondCharacterCurrentWeight uint32

	SecondCharacterMaxWeight uint32
}

func (m *ExchangeStartedWithPodsMessage) ID() uint16 {
	return 6129
}

func (m *ExchangeStartedWithPodsMessage) Serialize(w Writer) error {

	if err := m.ExchangeStartedMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.FirstCharacterId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.FirstCharacterCurrentWeight); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.FirstCharacterMaxWeight); err != nil {
		return err
	}

	if err := w.WriteDouble(m.SecondCharacterId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.SecondCharacterCurrentWeight); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.SecondCharacterMaxWeight); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeStartedWithPodsMessage) Deserialize(r Reader) error {

	if err := m.ExchangeStartedMessage.Deserialize(r); err != nil {
		return err
	}

	lfirstCharacterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.FirstCharacterId = lfirstCharacterId

	lfirstCharacterCurrentWeight, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.FirstCharacterCurrentWeight = lfirstCharacterCurrentWeight

	lfirstCharacterMaxWeight, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.FirstCharacterMaxWeight = lfirstCharacterMaxWeight

	lsecondCharacterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.SecondCharacterId = lsecondCharacterId

	lsecondCharacterCurrentWeight, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SecondCharacterCurrentWeight = lsecondCharacterCurrentWeight

	lsecondCharacterMaxWeight, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SecondCharacterMaxWeight = lsecondCharacterMaxWeight

	return nil
}

type MapComplementaryInformationsDataInHouseMessage struct {
	MapComplementaryInformationsDataMessage

	CurrentHouse HouseInformationsInside
}

func (m *MapComplementaryInformationsDataInHouseMessage) ID() uint16 {
	return 6130
}

func (m *MapComplementaryInformationsDataInHouseMessage) Serialize(w Writer) error {

	if err := m.MapComplementaryInformationsDataMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.CurrentHouse.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *MapComplementaryInformationsDataInHouseMessage) Deserialize(r Reader) error {

	if err := m.MapComplementaryInformationsDataMessage.Deserialize(r); err != nil {
		return err
	}

	var lcurrentHouse HouseInformationsInside

	lcurrentHouse.Deserialize(r)

	m.CurrentHouse = lcurrentHouse

	return nil
}

type GameActionFightNoSpellCastMessage struct {
	SpellLevelId uint32
}

func (m *GameActionFightNoSpellCastMessage) ID() uint16 {
	return 6132
}

func (m *GameActionFightNoSpellCastMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.SpellLevelId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightNoSpellCastMessage) Deserialize(r Reader) error {

	lspellLevelId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SpellLevelId = lspellLevelId

	return nil
}

type PopupWarningMessage struct {
	LockDuration uint8

	Author string

	Content string
}

func (m *PopupWarningMessage) ID() uint16 {
	return 6134
}

func (m *PopupWarningMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.LockDuration); err != nil {
		return err
	}

	if err := w.WriteString(m.Author); err != nil {
		return err
	}

	if err := w.WriteString(m.Content); err != nil {
		return err
	}

	return nil
}

func (m *PopupWarningMessage) Deserialize(r Reader) error {

	llockDuration, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.LockDuration = llockDuration

	lauthor, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Author = lauthor

	lcontent, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Content = lcontent

	return nil
}

type ChatAdminServerMessage struct {
	ChatServerMessage
}

func (m *ChatAdminServerMessage) ID() uint16 {
	return 6135
}

func (m *ChatAdminServerMessage) Serialize(w Writer) error {

	if err := m.ChatServerMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ChatAdminServerMessage) Deserialize(r Reader) error {

	if err := m.ChatServerMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type HouseToSellFilterMessage struct {
	AreaId int32

	AtLeastNbRoom uint8

	AtLeastNbChest uint8

	SkillRequested uint16

	MaxPrice int64
}

func (m *HouseToSellFilterMessage) ID() uint16 {
	return 6137
}

func (m *HouseToSellFilterMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.AreaId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.AtLeastNbRoom); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.AtLeastNbChest); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SkillRequested); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.MaxPrice); err != nil {
		return err
	}

	return nil
}

func (m *HouseToSellFilterMessage) Deserialize(r Reader) error {

	lareaId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.AreaId = lareaId

	latLeastNbRoom, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.AtLeastNbRoom = latLeastNbRoom

	latLeastNbChest, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.AtLeastNbChest = latLeastNbChest

	lskillRequested, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SkillRequested = lskillRequested

	lmaxPrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.MaxPrice = lmaxPrice

	return nil
}

type PaddockToSellListMessage struct {
	PageIndex uint16

	TotalPage uint16

	PaddockList []PaddockInformationsForSell
}

func (m *PaddockToSellListMessage) ID() uint16 {
	return 6138
}

func (m *PaddockToSellListMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.PageIndex); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TotalPage); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.PaddockList))); err != nil {
		return err
	}

	for i := range m.PaddockList {

		if err := m.PaddockList[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PaddockToSellListMessage) Deserialize(r Reader) error {

	lpageIndex, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.PageIndex = lpageIndex

	ltotalPage, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TotalPage = ltotalPage

	lpaddockListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PaddockList = make([]PaddockInformationsForSell, lpaddockListLen)

	for i := range m.PaddockList {

		var lpaddockList PaddockInformationsForSell

		lpaddockList.Deserialize(r)

		m.PaddockList[i] = lpaddockList

	}

	return nil
}

type HouseToSellListRequestMessage struct {
	PageIndex uint16
}

func (m *HouseToSellListRequestMessage) ID() uint16 {
	return 6139
}

func (m *HouseToSellListRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.PageIndex); err != nil {
		return err
	}

	return nil
}

func (m *HouseToSellListRequestMessage) Deserialize(r Reader) error {

	lpageIndex, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.PageIndex = lpageIndex

	return nil
}

type HouseToSellListMessage struct {
	PageIndex uint16

	TotalPage uint16

	HouseList []HouseInformationsForSell
}

func (m *HouseToSellListMessage) ID() uint16 {
	return 6140
}

func (m *HouseToSellListMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.PageIndex); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.TotalPage); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.HouseList))); err != nil {
		return err
	}

	for i := range m.HouseList {

		if err := m.HouseList[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *HouseToSellListMessage) Deserialize(r Reader) error {

	lpageIndex, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.PageIndex = lpageIndex

	ltotalPage, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TotalPage = ltotalPage

	lhouseListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.HouseList = make([]HouseInformationsForSell, lhouseListLen)

	for i := range m.HouseList {

		var lhouseList HouseInformationsForSell

		lhouseList.Deserialize(r)

		m.HouseList[i] = lhouseList

	}

	return nil
}

type PaddockToSellListRequestMessage struct {
	PageIndex uint16
}

func (m *PaddockToSellListRequestMessage) ID() uint16 {
	return 6141
}

func (m *PaddockToSellListRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.PageIndex); err != nil {
		return err
	}

	return nil
}

func (m *PaddockToSellListRequestMessage) Deserialize(r Reader) error {

	lpageIndex, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.PageIndex = lpageIndex

	return nil
}

type AcquaintanceServerListMessage struct {
	Servers []uint16
}

func (m *AcquaintanceServerListMessage) ID() uint16 {
	return 6142
}

func (m *AcquaintanceServerListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Servers))); err != nil {
		return err
	}

	for i := range m.Servers {

		if err := w.WriteVarUInt16(m.Servers[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *AcquaintanceServerListMessage) Deserialize(r Reader) error {

	lserversLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Servers = make([]uint16, lserversLen)

	for i := range m.Servers {

		lservers, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Servers[i] = lservers

	}

	return nil
}

type AcquaintanceSearchErrorMessage struct {
	Reason uint8
}

func (m *AcquaintanceSearchErrorMessage) ID() uint16 {
	return 6143
}

func (m *AcquaintanceSearchErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *AcquaintanceSearchErrorMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type AcquaintanceSearchMessage struct {
	Nickname string
}

func (m *AcquaintanceSearchMessage) ID() uint16 {
	return 6144
}

func (m *AcquaintanceSearchMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Nickname); err != nil {
		return err
	}

	return nil
}

func (m *AcquaintanceSearchMessage) Deserialize(r Reader) error {

	lnickname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Nickname = lnickname

	return nil
}

type GameActionFightTriggerEffectMessage struct {
	GameActionFightDispellEffectMessage
}

func (m *GameActionFightTriggerEffectMessage) ID() uint16 {
	return 6147
}

func (m *GameActionFightTriggerEffectMessage) Serialize(w Writer) error {

	if err := m.GameActionFightDispellEffectMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightTriggerEffectMessage) Deserialize(r Reader) error {

	if err := m.GameActionFightDispellEffectMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type GameRolePlayDelayedActionFinishedMessage struct {
	DelayedCharacterId float64

	DelayTypeId uint8
}

func (m *GameRolePlayDelayedActionFinishedMessage) ID() uint16 {
	return 6150
}

func (m *GameRolePlayDelayedActionFinishedMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.DelayedCharacterId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.DelayTypeId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayDelayedActionFinishedMessage) Deserialize(r Reader) error {

	ldelayedCharacterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DelayedCharacterId = ldelayedCharacterId

	ldelayTypeId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.DelayTypeId = ldelayTypeId

	return nil
}

type GameRolePlayDelayedActionMessage struct {
	DelayedCharacterId float64

	DelayTypeId uint8

	DelayEndTime float64
}

func (m *GameRolePlayDelayedActionMessage) ID() uint16 {
	return 6153
}

func (m *GameRolePlayDelayedActionMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.DelayedCharacterId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.DelayTypeId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.DelayEndTime); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayDelayedActionMessage) Deserialize(r Reader) error {

	ldelayedCharacterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DelayedCharacterId = ldelayedCharacterId

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

	return nil
}

type CheckFileRequestMessage struct {
	Filename string

	Type uint8
}

func (m *CheckFileRequestMessage) ID() uint16 {
	return 6154
}

func (m *CheckFileRequestMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Filename); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	return nil
}

func (m *CheckFileRequestMessage) Deserialize(r Reader) error {

	lfilename, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Filename = lfilename

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	return nil
}

type GameMapChangeOrientationsMessage struct {
	Orientations []ActorOrientation
}

func (m *GameMapChangeOrientationsMessage) ID() uint16 {
	return 6155
}

func (m *GameMapChangeOrientationsMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Orientations))); err != nil {
		return err
	}

	for i := range m.Orientations {

		if err := m.Orientations[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameMapChangeOrientationsMessage) Deserialize(r Reader) error {

	lorientationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Orientations = make([]ActorOrientation, lorientationsLen)

	for i := range m.Orientations {

		var lorientations ActorOrientation

		lorientations.Deserialize(r)

		m.Orientations[i] = lorientations

	}

	return nil
}

type CheckFileMessage struct {
	FilenameHash string

	Type uint8

	Value string
}

func (m *CheckFileMessage) ID() uint16 {
	return 6156
}

func (m *CheckFileMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.FilenameHash); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	if err := w.WriteString(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *CheckFileMessage) Deserialize(r Reader) error {

	lfilenameHash, err := r.ReadString()
	if err != nil {
		return err
	}

	m.FilenameHash = lfilenameHash

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	lvalue, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type ShowCellSpectatorMessage struct {
	ShowCellMessage

	PlayerName string
}

func (m *ShowCellSpectatorMessage) ID() uint16 {
	return 6158
}

func (m *ShowCellSpectatorMessage) Serialize(w Writer) error {

	if err := m.ShowCellMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.PlayerName); err != nil {
		return err
	}

	return nil
}

func (m *ShowCellSpectatorMessage) Deserialize(r Reader) error {

	if err := m.ShowCellMessage.Deserialize(r); err != nil {
		return err
	}

	lplayerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PlayerName = lplayerName

	return nil
}

type GuildMemberSetWarnOnConnectionMessage struct {
	Enable bool
}

func (m *GuildMemberSetWarnOnConnectionMessage) ID() uint16 {
	return 6159
}

func (m *GuildMemberSetWarnOnConnectionMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *GuildMemberSetWarnOnConnectionMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type GuildMemberWarnOnConnectionStateMessage struct {
	Enable bool
}

func (m *GuildMemberWarnOnConnectionStateMessage) ID() uint16 {
	return 6160
}

func (m *GuildMemberWarnOnConnectionStateMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *GuildMemberWarnOnConnectionStateMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type PaddockToSellFilterMessage struct {
	AreaId int32

	AtLeastNbMount int8

	AtLeastNbMachine int8

	MaxPrice int64
}

func (m *PaddockToSellFilterMessage) ID() uint16 {
	return 6161
}

func (m *PaddockToSellFilterMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.AreaId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.AtLeastNbMount); err != nil {
		return err
	}

	if err := w.WriteInt8(m.AtLeastNbMachine); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.MaxPrice); err != nil {
		return err
	}

	return nil
}

func (m *PaddockToSellFilterMessage) Deserialize(r Reader) error {

	lareaId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.AreaId = lareaId

	latLeastNbMount, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.AtLeastNbMount = latLeastNbMount

	latLeastNbMachine, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.AtLeastNbMachine = latLeastNbMachine

	lmaxPrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.MaxPrice = lmaxPrice

	return nil
}

type InventoryContentAndPresetMessage struct {
	InventoryContentMessage

	Presets []Preset

	IdolsPresets []IdolsPreset
}

func (m *InventoryContentAndPresetMessage) ID() uint16 {
	return 6162
}

func (m *InventoryContentAndPresetMessage) Serialize(w Writer) error {

	if err := m.InventoryContentMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Presets))); err != nil {
		return err
	}

	for i := range m.Presets {

		if err := m.Presets[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.IdolsPresets))); err != nil {
		return err
	}

	for i := range m.IdolsPresets {

		if err := m.IdolsPresets[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *InventoryContentAndPresetMessage) Deserialize(r Reader) error {

	if err := m.InventoryContentMessage.Deserialize(r); err != nil {
		return err
	}

	lpresetsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Presets = make([]Preset, lpresetsLen)

	for i := range m.Presets {

		var lpresets Preset

		lpresets.Deserialize(r)

		m.Presets[i] = lpresets

	}

	lidolsPresetsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.IdolsPresets = make([]IdolsPreset, lidolsPresetsLen)

	for i := range m.IdolsPresets {

		var lidolsPresets IdolsPreset

		lidolsPresets.Deserialize(r)

		m.IdolsPresets[i] = lidolsPresets

	}

	return nil
}

type InventoryPresetUseResultMessage struct {
	PresetId uint8

	Code uint8

	UnlinkedPosition []uint8
}

func (m *InventoryPresetUseResultMessage) ID() uint16 {
	return 6163
}

func (m *InventoryPresetUseResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Code); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.UnlinkedPosition))); err != nil {
		return err
	}

	for i := range m.UnlinkedPosition {

		if err := w.WriteUInt8(m.UnlinkedPosition[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *InventoryPresetUseResultMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	lcode, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Code = lcode

	lunlinkedPositionLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.UnlinkedPosition = make([]uint8, lunlinkedPositionLen)

	for i := range m.UnlinkedPosition {

		lunlinkedPosition, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.UnlinkedPosition[i] = lunlinkedPosition

	}

	return nil
}

type InventoryPresetSaveMessage struct {
	PresetId uint8

	SymbolId uint8

	SaveEquipment bool
}

func (m *InventoryPresetSaveMessage) ID() uint16 {
	return 6165
}

func (m *InventoryPresetSaveMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.SymbolId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.SaveEquipment); err != nil {
		return err
	}

	return nil
}

func (m *InventoryPresetSaveMessage) Deserialize(r Reader) error {

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

	lsaveEquipment, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.SaveEquipment = lsaveEquipment

	return nil
}

type InventoryPresetUseMessage struct {
	PresetId uint8
}

func (m *InventoryPresetUseMessage) ID() uint16 {
	return 6167
}

func (m *InventoryPresetUseMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	return nil
}

func (m *InventoryPresetUseMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	return nil
}

type InventoryPresetItemUpdateMessage struct {
	PresetId uint8

	PresetItem PresetItem
}

func (m *InventoryPresetItemUpdateMessage) ID() uint16 {
	return 6168
}

func (m *InventoryPresetItemUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := m.PresetItem.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *InventoryPresetItemUpdateMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	var lpresetItem PresetItem

	lpresetItem.Deserialize(r)

	m.PresetItem = lpresetItem

	return nil
}

type InventoryPresetDeleteMessage struct {
	PresetId uint8
}

func (m *InventoryPresetDeleteMessage) ID() uint16 {
	return 6169
}

func (m *InventoryPresetDeleteMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	return nil
}

func (m *InventoryPresetDeleteMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	return nil
}

type InventoryPresetSaveResultMessage struct {
	PresetId uint8

	Code uint8
}

func (m *InventoryPresetSaveResultMessage) ID() uint16 {
	return 6170
}

func (m *InventoryPresetSaveResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Code); err != nil {
		return err
	}

	return nil
}

func (m *InventoryPresetSaveResultMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	lcode, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Code = lcode

	return nil
}

type InventoryPresetUpdateMessage struct {
	Preset Preset
}

func (m *InventoryPresetUpdateMessage) ID() uint16 {
	return 6171
}

func (m *InventoryPresetUpdateMessage) Serialize(w Writer) error {

	if err := m.Preset.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *InventoryPresetUpdateMessage) Deserialize(r Reader) error {

	var lpreset Preset

	lpreset.Deserialize(r)

	m.Preset = lpreset

	return nil
}

type MountDataErrorMessage struct {
	Reason uint8
}

func (m *MountDataErrorMessage) ID() uint16 {
	return 6172
}

func (m *MountDataErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *MountDataErrorMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type InventoryPresetDeleteResultMessage struct {
	PresetId uint8

	Code uint8
}

func (m *InventoryPresetDeleteResultMessage) ID() uint16 {
	return 6173
}

func (m *InventoryPresetDeleteResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Code); err != nil {
		return err
	}

	return nil
}

func (m *InventoryPresetDeleteResultMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	lcode, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Code = lcode

	return nil
}

type IdentificationFailedBannedMessage struct {
	IdentificationFailedMessage

	BanEndDate float64
}

func (m *IdentificationFailedBannedMessage) ID() uint16 {
	return 6174
}

func (m *IdentificationFailedBannedMessage) Serialize(w Writer) error {

	if err := m.IdentificationFailedMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.BanEndDate); err != nil {
		return err
	}

	return nil
}

func (m *IdentificationFailedBannedMessage) Deserialize(r Reader) error {

	if err := m.IdentificationFailedMessage.Deserialize(r); err != nil {
		return err
	}

	lbanEndDate, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.BanEndDate = lbanEndDate

	return nil
}

type PartyRestrictedMessage struct {
	AbstractPartyMessage

	Restricted bool
}

func (m *PartyRestrictedMessage) ID() uint16 {
	return 6175
}

func (m *PartyRestrictedMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Restricted); err != nil {
		return err
	}

	return nil
}

func (m *PartyRestrictedMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lrestricted, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Restricted = lrestricted

	return nil
}

type GameActionFightDispellSpellMessage struct {
	GameActionFightDispellMessage

	SpellId uint16
}

func (m *GameActionFightDispellSpellMessage) ID() uint16 {
	return 6176
}

func (m *GameActionFightDispellSpellMessage) Serialize(w Writer) error {

	if err := m.GameActionFightDispellMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightDispellSpellMessage) Deserialize(r Reader) error {

	if err := m.GameActionFightDispellMessage.Deserialize(r); err != nil {
		return err
	}

	lspellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	return nil
}

type UpdateMountBoostMessage struct {
	RideId int32

	BoostToUpdateList []UpdateMountBoostIntrf
}

func (m *UpdateMountBoostMessage) ID() uint16 {
	return 6179
}

func (m *UpdateMountBoostMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt32(m.RideId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.BoostToUpdateList))); err != nil {
		return err
	}

	for i := range m.BoostToUpdateList {

		if err := w.WriteUInt16(m.BoostToUpdateList[i].ID()); err != nil {
			return err
		}

		if err := m.BoostToUpdateList[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *UpdateMountBoostMessage) Deserialize(r Reader) error {

	lrideId, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.RideId = lrideId

	lboostToUpdateListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.BoostToUpdateList = make([]UpdateMountBoostIntrf, lboostToUpdateListLen)

	for i := range m.BoostToUpdateList {

		typeboostToUpdateListID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lboostToUpdateList, err := GetType(typeboostToUpdateListID)
		if err != nil {
			return err
		}

		lboostToUpdateList.Deserialize(r)

		m.BoostToUpdateList[i] = lboostToUpdateList.(UpdateMountBoostIntrf)

	}

	return nil
}

type GuildHouseRemoveMessage struct {
	HouseId uint32

	InstanceId uint32

	SecondHand bool
}

func (m *GuildHouseRemoveMessage) ID() uint16 {
	return 6180
}

func (m *GuildHouseRemoveMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HouseId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.SecondHand); err != nil {
		return err
	}

	return nil
}

func (m *GuildHouseRemoveMessage) Deserialize(r Reader) error {

	lhouseId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HouseId = lhouseId

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

	return nil
}

type GuildHouseUpdateInformationMessage struct {
	HousesInformations HouseInformationsForGuild
}

func (m *GuildHouseUpdateInformationMessage) ID() uint16 {
	return 6181
}

func (m *GuildHouseUpdateInformationMessage) Serialize(w Writer) error {

	if err := m.HousesInformations.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildHouseUpdateInformationMessage) Deserialize(r Reader) error {

	var lhousesInformations HouseInformationsForGuild

	lhousesInformations.Deserialize(r)

	m.HousesInformations = lhousesInformations

	return nil
}

type ExchangeObjectTransfertListFromInvMessage struct {
	Ids []uint32
}

func (m *ExchangeObjectTransfertListFromInvMessage) ID() uint16 {
	return 6183
}

func (m *ExchangeObjectTransfertListFromInvMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Ids))); err != nil {
		return err
	}

	for i := range m.Ids {

		if err := w.WriteVarUInt32(m.Ids[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeObjectTransfertListFromInvMessage) Deserialize(r Reader) error {

	lidsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Ids = make([]uint32, lidsLen)

	for i := range m.Ids {

		lids, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.Ids[i] = lids

	}

	return nil
}

type ExchangeObjectTransfertAllFromInvMessage struct {
}

func (m *ExchangeObjectTransfertAllFromInvMessage) ID() uint16 {
	return 6184
}

func (m *ExchangeObjectTransfertAllFromInvMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeObjectTransfertAllFromInvMessage) Deserialize(r Reader) error {

	return nil
}

type LocalizedChatSmileyMessage struct {
	ChatSmileyMessage

	CellId uint16
}

func (m *LocalizedChatSmileyMessage) ID() uint16 {
	return 6185
}

func (m *LocalizedChatSmileyMessage) Serialize(w Writer) error {

	if err := m.ChatSmileyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *LocalizedChatSmileyMessage) Deserialize(r Reader) error {

	if err := m.ChatSmileyMessage.Deserialize(r); err != nil {
		return err
	}

	lcellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type ExchangeCraftResultMagicWithObjectDescMessage struct {
	ExchangeCraftResultWithObjectDescMessage

	MagicPoolStatus int8
}

func (m *ExchangeCraftResultMagicWithObjectDescMessage) ID() uint16 {
	return 6188
}

func (m *ExchangeCraftResultMagicWithObjectDescMessage) Serialize(w Writer) error {

	if err := m.ExchangeCraftResultWithObjectDescMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt8(m.MagicPoolStatus); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeCraftResultMagicWithObjectDescMessage) Deserialize(r Reader) error {

	if err := m.ExchangeCraftResultWithObjectDescMessage.Deserialize(r); err != nil {
		return err
	}

	lmagicPoolStatus, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.MagicPoolStatus = lmagicPoolStatus

	return nil
}

type MountFeedRequestMessage struct {
	MountUid uint32

	MountLocation int8

	MountFoodUid uint32

	Quantity uint32
}

func (m *MountFeedRequestMessage) ID() uint16 {
	return 6189
}

func (m *MountFeedRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.MountUid); err != nil {
		return err
	}

	if err := w.WriteInt8(m.MountLocation); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MountFoodUid); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *MountFeedRequestMessage) Deserialize(r Reader) error {

	lmountUid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MountUid = lmountUid

	lmountLocation, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.MountLocation = lmountLocation

	lmountFoodUid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MountFoodUid = lmountFoodUid

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type GameRolePlayAttackMonsterRequestMessage struct {
	MonsterGroupId float64
}

func (m *GameRolePlayAttackMonsterRequestMessage) ID() uint16 {
	return 6191
}

func (m *GameRolePlayAttackMonsterRequestMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.MonsterGroupId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayAttackMonsterRequestMessage) Deserialize(r Reader) error {

	lmonsterGroupId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.MonsterGroupId = lmonsterGroupId

	return nil
}

type MoodSmileyRequestMessage struct {
	SmileyId uint16
}

func (m *MoodSmileyRequestMessage) ID() uint16 {
	return 6192
}

func (m *MoodSmileyRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SmileyId); err != nil {
		return err
	}

	return nil
}

func (m *MoodSmileyRequestMessage) Deserialize(r Reader) error {

	lsmileyId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SmileyId = lsmileyId

	return nil
}

type MoodSmileyResultMessage struct {
	ResultCode uint8

	SmileyId uint16
}

func (m *MoodSmileyResultMessage) ID() uint16 {
	return 6196
}

func (m *MoodSmileyResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.ResultCode); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SmileyId); err != nil {
		return err
	}

	return nil
}

func (m *MoodSmileyResultMessage) Deserialize(r Reader) error {

	lresultCode, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ResultCode = lresultCode

	lsmileyId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SmileyId = lsmileyId

	return nil
}

type ErrorMapNotFoundMessage struct {
	MapId uint32
}

func (m *ErrorMapNotFoundMessage) ID() uint16 {
	return 6197
}

func (m *ErrorMapNotFoundMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.MapId); err != nil {
		return err
	}

	return nil
}

func (m *ErrorMapNotFoundMessage) Deserialize(r Reader) error {

	lmapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	return nil
}

type AchievementListMessage struct {
	FinishedAchievementsIds []uint16

	RewardableAchievements []AchievementRewardable
}

func (m *AchievementListMessage) ID() uint16 {
	return 6205
}

func (m *AchievementListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.FinishedAchievementsIds))); err != nil {
		return err
	}

	for i := range m.FinishedAchievementsIds {

		if err := w.WriteVarUInt16(m.FinishedAchievementsIds[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.RewardableAchievements))); err != nil {
		return err
	}

	for i := range m.RewardableAchievements {

		if err := m.RewardableAchievements[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *AchievementListMessage) Deserialize(r Reader) error {

	lfinishedAchievementsIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FinishedAchievementsIds = make([]uint16, lfinishedAchievementsIdsLen)

	for i := range m.FinishedAchievementsIds {

		lfinishedAchievementsIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.FinishedAchievementsIds[i] = lfinishedAchievementsIds

	}

	lrewardableAchievementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.RewardableAchievements = make([]AchievementRewardable, lrewardableAchievementsLen)

	for i := range m.RewardableAchievements {

		var lrewardableAchievements AchievementRewardable

		lrewardableAchievements.Deserialize(r)

		m.RewardableAchievements[i] = lrewardableAchievements

	}

	return nil
}

type ObjectsQuantityMessage struct {
	ObjectsUIDAndQty []ObjectItemQuantity
}

func (m *ObjectsQuantityMessage) ID() uint16 {
	return 6206
}

func (m *ObjectsQuantityMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ObjectsUIDAndQty))); err != nil {
		return err
	}

	for i := range m.ObjectsUIDAndQty {

		if err := m.ObjectsUIDAndQty[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ObjectsQuantityMessage) Deserialize(r Reader) error {

	lobjectsUIDAndQtyLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectsUIDAndQty = make([]ObjectItemQuantity, lobjectsUIDAndQtyLen)

	for i := range m.ObjectsUIDAndQty {

		var lobjectsUIDAndQty ObjectItemQuantity

		lobjectsUIDAndQty.Deserialize(r)

		m.ObjectsUIDAndQty[i] = lobjectsUIDAndQty

	}

	return nil
}

type AchievementFinishedMessage struct {
	Id uint16

	Finishedlevel uint8
}

func (m *AchievementFinishedMessage) ID() uint16 {
	return 6208
}

func (m *AchievementFinishedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Id); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Finishedlevel); err != nil {
		return err
	}

	return nil
}

func (m *AchievementFinishedMessage) Deserialize(r Reader) error {

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

type IdentificationSuccessWithLoginTokenMessage struct {
	IdentificationSuccessMessage

	LoginToken string
}

func (m *IdentificationSuccessWithLoginTokenMessage) ID() uint16 {
	return 6209
}

func (m *IdentificationSuccessWithLoginTokenMessage) Serialize(w Writer) error {

	if err := m.IdentificationSuccessMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.LoginToken); err != nil {
		return err
	}

	return nil
}

func (m *IdentificationSuccessWithLoginTokenMessage) Deserialize(r Reader) error {

	if err := m.IdentificationSuccessMessage.Deserialize(r); err != nil {
		return err
	}

	lloginToken, err := r.ReadString()
	if err != nil {
		return err
	}

	m.LoginToken = lloginToken

	return nil
}

type InventoryPresetItemUpdateRequestMessage struct {
	PresetId uint8

	Position uint8

	ObjUid uint32
}

func (m *InventoryPresetItemUpdateRequestMessage) ID() uint16 {
	return 6210
}

func (m *InventoryPresetItemUpdateRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Position); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.ObjUid); err != nil {
		return err
	}

	return nil
}

func (m *InventoryPresetItemUpdateRequestMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	lposition, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Position = lposition

	lobjUid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjUid = lobjUid

	return nil
}

type InventoryPresetItemUpdateErrorMessage struct {
	Code uint8
}

func (m *InventoryPresetItemUpdateErrorMessage) ID() uint16 {
	return 6211
}

func (m *InventoryPresetItemUpdateErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Code); err != nil {
		return err
	}

	return nil
}

func (m *InventoryPresetItemUpdateErrorMessage) Deserialize(r Reader) error {

	lcode, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Code = lcode

	return nil
}

type SlaveSwitchContextMessage struct {
	MasterId float64

	SlaveId float64

	SlaveSpells []SpellItem

	SlaveStats CharacterCharacteristicsInformations

	Shortcuts []ShortcutIntrf
}

func (m *SlaveSwitchContextMessage) ID() uint16 {
	return 6214
}

func (m *SlaveSwitchContextMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.MasterId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.SlaveId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.SlaveSpells))); err != nil {
		return err
	}

	for i := range m.SlaveSpells {

		if err := m.SlaveSpells[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := m.SlaveStats.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Shortcuts))); err != nil {
		return err
	}

	for i := range m.Shortcuts {

		if err := w.WriteUInt16(m.Shortcuts[i].ID()); err != nil {
			return err
		}

		if err := m.Shortcuts[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *SlaveSwitchContextMessage) Deserialize(r Reader) error {

	lmasterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.MasterId = lmasterId

	lslaveId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.SlaveId = lslaveId

	lslaveSpellsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SlaveSpells = make([]SpellItem, lslaveSpellsLen)

	for i := range m.SlaveSpells {

		var lslaveSpells SpellItem

		lslaveSpells.Deserialize(r)

		m.SlaveSpells[i] = lslaveSpells

	}

	var lslaveStats CharacterCharacteristicsInformations

	lslaveStats.Deserialize(r)

	m.SlaveStats = lslaveStats

	lshortcutsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Shortcuts = make([]ShortcutIntrf, lshortcutsLen)

	for i := range m.Shortcuts {

		typeshortcutsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lshortcuts, err := GetType(typeshortcutsID)
		if err != nil {
			return err
		}

		lshortcuts.Deserialize(r)

		m.Shortcuts[i] = lshortcuts.(ShortcutIntrf)

	}

	return nil
}

type GameFightResumeWithSlavesMessage struct {
	GameFightResumeMessage

	SlavesInfo []GameFightResumeSlaveInfo
}

func (m *GameFightResumeWithSlavesMessage) ID() uint16 {
	return 6215
}

func (m *GameFightResumeWithSlavesMessage) Serialize(w Writer) error {

	if err := m.GameFightResumeMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.SlavesInfo))); err != nil {
		return err
	}

	for i := range m.SlavesInfo {

		if err := m.SlavesInfo[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameFightResumeWithSlavesMessage) Deserialize(r Reader) error {

	if err := m.GameFightResumeMessage.Deserialize(r); err != nil {
		return err
	}

	lslavesInfoLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SlavesInfo = make([]GameFightResumeSlaveInfo, lslavesInfoLen)

	for i := range m.SlavesInfo {

		var lslavesInfo GameFightResumeSlaveInfo

		lslavesInfo.Deserialize(r)

		m.SlavesInfo[i] = lslavesInfo

	}

	return nil
}

type AccountCapabilitiesMessage struct {
	AccountId uint32

	TutorialAvailable bool

	BreedsVisible uint32

	BreedsAvailable uint32

	Status int8

	CanCreateNewCharacter bool
}

func (m *AccountCapabilitiesMessage) ID() uint16 {
	return 6216
}

func (m *AccountCapabilitiesMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.TutorialAvailable)

	setWrappedFlag(bbw0, 1, m.CanCreateNewCharacter)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.BreedsVisible); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.BreedsAvailable); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Status); err != nil {
		return err
	}

	return nil
}

func (m *AccountCapabilitiesMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TutorialAvailable = getWrappedFlag(bbw0, 0)

	m.CanCreateNewCharacter = getWrappedFlag(bbw0, 1)

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	lbreedsVisible, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.BreedsVisible = lbreedsVisible

	lbreedsAvailable, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.BreedsAvailable = lbreedsAvailable

	lstatus, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Status = lstatus

	return nil
}

type GameActionFightVanishMessage struct {
	AbstractGameActionMessage

	TargetId float64
}

func (m *GameActionFightVanishMessage) ID() uint16 {
	return 6217
}

func (m *GameActionFightVanishMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightVanishMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type GameFightShowFighterRandomStaticPoseMessage struct {
	GameFightShowFighterMessage
}

func (m *GameFightShowFighterRandomStaticPoseMessage) ID() uint16 {
	return 6218
}

func (m *GameFightShowFighterRandomStaticPoseMessage) Serialize(w Writer) error {

	if err := m.GameFightShowFighterMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameFightShowFighterRandomStaticPoseMessage) Deserialize(r Reader) error {

	if err := m.GameFightShowFighterMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type GameActionFightSpellCooldownVariationMessage struct {
	AbstractGameActionMessage

	TargetId float64

	SpellId uint16

	Value int16
}

func (m *GameActionFightSpellCooldownVariationMessage) ID() uint16 {
	return 6219
}

func (m *GameActionFightSpellCooldownVariationMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.Value); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightSpellCooldownVariationMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lspellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	lvalue, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.Value = lvalue

	return nil
}

type GameActionFightSpellImmunityMessage struct {
	AbstractGameActionMessage

	TargetId float64

	SpellId uint16
}

func (m *GameActionFightSpellImmunityMessage) ID() uint16 {
	return 6221
}

func (m *GameActionFightSpellImmunityMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightSpellImmunityMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lspellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	return nil
}

type ShortcutBarRemoveErrorMessage struct {
	Error uint8
}

func (m *ShortcutBarRemoveErrorMessage) ID() uint16 {
	return 6222
}

func (m *ShortcutBarRemoveErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Error); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutBarRemoveErrorMessage) Deserialize(r Reader) error {

	lerror, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Error = lerror

	return nil
}

type ShortcutBarRemovedMessage struct {
	BarType uint8

	Slot uint8
}

func (m *ShortcutBarRemovedMessage) ID() uint16 {
	return 6224
}

func (m *ShortcutBarRemovedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.BarType); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Slot); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutBarRemovedMessage) Deserialize(r Reader) error {

	lbarType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BarType = lbarType

	lslot, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Slot = lslot

	return nil
}

type ShortcutBarAddRequestMessage struct {
	BarType uint8

	Shortcut ShortcutIntrf
}

func (m *ShortcutBarAddRequestMessage) ID() uint16 {
	return 6225
}

func (m *ShortcutBarAddRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.BarType); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Shortcut.ID()); err != nil {
		return err
	}

	if err := m.Shortcut.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutBarAddRequestMessage) Deserialize(r Reader) error {

	lbarType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BarType = lbarType

	typeshortcutID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lshortcut, err := GetType(typeshortcutID)
	if err != nil {
		return err
	}

	lshortcut.Deserialize(r)

	m.Shortcut = lshortcut.(ShortcutIntrf)

	return nil
}

type ShortcutBarSwapErrorMessage struct {
	Error uint8
}

func (m *ShortcutBarSwapErrorMessage) ID() uint16 {
	return 6226
}

func (m *ShortcutBarSwapErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Error); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutBarSwapErrorMessage) Deserialize(r Reader) error {

	lerror, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Error = lerror

	return nil
}

type ShortcutBarAddErrorMessage struct {
	Error uint8
}

func (m *ShortcutBarAddErrorMessage) ID() uint16 {
	return 6227
}

func (m *ShortcutBarAddErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Error); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutBarAddErrorMessage) Deserialize(r Reader) error {

	lerror, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Error = lerror

	return nil
}

type ShortcutBarRemoveRequestMessage struct {
	BarType uint8

	Slot uint8
}

func (m *ShortcutBarRemoveRequestMessage) ID() uint16 {
	return 6228
}

func (m *ShortcutBarRemoveRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.BarType); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Slot); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutBarRemoveRequestMessage) Deserialize(r Reader) error {

	lbarType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BarType = lbarType

	lslot, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Slot = lslot

	return nil
}

type ShortcutBarRefreshMessage struct {
	BarType uint8

	Shortcut ShortcutIntrf
}

func (m *ShortcutBarRefreshMessage) ID() uint16 {
	return 6229
}

func (m *ShortcutBarRefreshMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.BarType); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Shortcut.ID()); err != nil {
		return err
	}

	if err := m.Shortcut.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutBarRefreshMessage) Deserialize(r Reader) error {

	lbarType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BarType = lbarType

	typeshortcutID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lshortcut, err := GetType(typeshortcutID)
	if err != nil {
		return err
	}

	lshortcut.Deserialize(r)

	m.Shortcut = lshortcut.(ShortcutIntrf)

	return nil
}

type ShortcutBarSwapRequestMessage struct {
	BarType uint8

	FirstSlot uint8

	SecondSlot uint8
}

func (m *ShortcutBarSwapRequestMessage) ID() uint16 {
	return 6230
}

func (m *ShortcutBarSwapRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.BarType); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.FirstSlot); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.SecondSlot); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutBarSwapRequestMessage) Deserialize(r Reader) error {

	lbarType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BarType = lbarType

	lfirstSlot, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.FirstSlot = lfirstSlot

	lsecondSlot, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.SecondSlot = lsecondSlot

	return nil
}

type ShortcutBarContentMessage struct {
	BarType uint8

	Shortcuts []ShortcutIntrf
}

func (m *ShortcutBarContentMessage) ID() uint16 {
	return 6231
}

func (m *ShortcutBarContentMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.BarType); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Shortcuts))); err != nil {
		return err
	}

	for i := range m.Shortcuts {

		if err := w.WriteUInt16(m.Shortcuts[i].ID()); err != nil {
			return err
		}

		if err := m.Shortcuts[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ShortcutBarContentMessage) Deserialize(r Reader) error {

	lbarType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BarType = lbarType

	lshortcutsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Shortcuts = make([]ShortcutIntrf, lshortcutsLen)

	for i := range m.Shortcuts {

		typeshortcutsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lshortcuts, err := GetType(typeshortcutsID)
		if err != nil {
			return err
		}

		lshortcuts.Deserialize(r)

		m.Shortcuts[i] = lshortcuts.(ShortcutIntrf)

	}

	return nil
}

type ObjectUseMultipleMessage struct {
	ObjectUseMessage

	Quantity uint32
}

func (m *ObjectUseMultipleMessage) ID() uint16 {
	return 6234
}

func (m *ObjectUseMultipleMessage) Serialize(w Writer) error {

	if err := m.ObjectUseMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ObjectUseMultipleMessage) Deserialize(r Reader) error {

	if err := m.ObjectUseMessage.Deserialize(r); err != nil {
		return err
	}

	lquantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type GuildFightTakePlaceRequestMessage struct {
	GuildFightJoinRequestMessage

	ReplacedCharacterId int32
}

func (m *GuildFightTakePlaceRequestMessage) ID() uint16 {
	return 6235
}

func (m *GuildFightTakePlaceRequestMessage) Serialize(w Writer) error {

	if err := m.GuildFightJoinRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.ReplacedCharacterId); err != nil {
		return err
	}

	return nil
}

func (m *GuildFightTakePlaceRequestMessage) Deserialize(r Reader) error {

	if err := m.GuildFightJoinRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lreplacedCharacterId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.ReplacedCharacterId = lreplacedCharacterId

	return nil
}

type ExchangeStartedWithStorageMessage struct {
	ExchangeStartedMessage

	StorageMaxSlot uint32
}

func (m *ExchangeStartedWithStorageMessage) ID() uint16 {
	return 6236
}

func (m *ExchangeStartedWithStorageMessage) Serialize(w Writer) error {

	if err := m.ExchangeStartedMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.StorageMaxSlot); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeStartedWithStorageMessage) Deserialize(r Reader) error {

	if err := m.ExchangeStartedMessage.Deserialize(r); err != nil {
		return err
	}

	lstorageMaxSlot, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.StorageMaxSlot = lstorageMaxSlot

	return nil
}

type ServerExperienceModificatorMessage struct {
	ExperiencePercent uint16
}

func (m *ServerExperienceModificatorMessage) ID() uint16 {
	return 6237
}

func (m *ServerExperienceModificatorMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ExperiencePercent); err != nil {
		return err
	}

	return nil
}

func (m *ServerExperienceModificatorMessage) Deserialize(r Reader) error {

	lexperiencePercent, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ExperiencePercent = lexperiencePercent

	return nil
}

type ExchangeObjectModifyPricedMessage struct {
	ExchangeObjectMovePricedMessage
}

func (m *ExchangeObjectModifyPricedMessage) ID() uint16 {
	return 6238
}

func (m *ExchangeObjectModifyPricedMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMovePricedMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeObjectModifyPricedMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMovePricedMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type GameFightNewRoundMessage struct {
	RoundNumber uint32
}

func (m *GameFightNewRoundMessage) ID() uint16 {
	return 6239
}

func (m *GameFightNewRoundMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.RoundNumber); err != nil {
		return err
	}

	return nil
}

func (m *GameFightNewRoundMessage) Deserialize(r Reader) error {

	lroundNumber, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.RoundNumber = lroundNumber

	return nil
}

type DungeonPartyFinderAvailableDungeonsRequestMessage struct {
}

func (m *DungeonPartyFinderAvailableDungeonsRequestMessage) ID() uint16 {
	return 6240
}

func (m *DungeonPartyFinderAvailableDungeonsRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *DungeonPartyFinderAvailableDungeonsRequestMessage) Deserialize(r Reader) error {

	return nil
}

type DungeonPartyFinderRegisterSuccessMessage struct {
	DungeonIds []uint16
}

func (m *DungeonPartyFinderRegisterSuccessMessage) ID() uint16 {
	return 6241
}

func (m *DungeonPartyFinderRegisterSuccessMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.DungeonIds))); err != nil {
		return err
	}

	for i := range m.DungeonIds {

		if err := w.WriteVarUInt16(m.DungeonIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *DungeonPartyFinderRegisterSuccessMessage) Deserialize(r Reader) error {

	ldungeonIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DungeonIds = make([]uint16, ldungeonIdsLen)

	for i := range m.DungeonIds {

		ldungeonIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.DungeonIds[i] = ldungeonIds

	}

	return nil
}

type DungeonPartyFinderAvailableDungeonsMessage struct {
	DungeonIds []uint16
}

func (m *DungeonPartyFinderAvailableDungeonsMessage) ID() uint16 {
	return 6242
}

func (m *DungeonPartyFinderAvailableDungeonsMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.DungeonIds))); err != nil {
		return err
	}

	for i := range m.DungeonIds {

		if err := w.WriteVarUInt16(m.DungeonIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *DungeonPartyFinderAvailableDungeonsMessage) Deserialize(r Reader) error {

	ldungeonIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DungeonIds = make([]uint16, ldungeonIdsLen)

	for i := range m.DungeonIds {

		ldungeonIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.DungeonIds[i] = ldungeonIds

	}

	return nil
}

type DungeonPartyFinderRegisterErrorMessage struct {
}

func (m *DungeonPartyFinderRegisterErrorMessage) ID() uint16 {
	return 6243
}

func (m *DungeonPartyFinderRegisterErrorMessage) Serialize(w Writer) error {

	return nil
}

func (m *DungeonPartyFinderRegisterErrorMessage) Deserialize(r Reader) error {

	return nil
}

type PartyInvitationDungeonMessage struct {
	PartyInvitationMessage

	DungeonId uint16
}

func (m *PartyInvitationDungeonMessage) ID() uint16 {
	return 6244
}

func (m *PartyInvitationDungeonMessage) Serialize(w Writer) error {

	if err := m.PartyInvitationMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	return nil
}

func (m *PartyInvitationDungeonMessage) Deserialize(r Reader) error {

	if err := m.PartyInvitationMessage.Deserialize(r); err != nil {
		return err
	}

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	return nil
}

type PartyInvitationDungeonRequestMessage struct {
	PartyInvitationRequestMessage

	DungeonId uint16
}

func (m *PartyInvitationDungeonRequestMessage) ID() uint16 {
	return 6245
}

func (m *PartyInvitationDungeonRequestMessage) Serialize(w Writer) error {

	if err := m.PartyInvitationRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	return nil
}

func (m *PartyInvitationDungeonRequestMessage) Deserialize(r Reader) error {

	if err := m.PartyInvitationRequestMessage.Deserialize(r); err != nil {
		return err
	}

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	return nil
}

type DungeonPartyFinderListenRequestMessage struct {
	DungeonId uint16
}

func (m *DungeonPartyFinderListenRequestMessage) ID() uint16 {
	return 6246
}

func (m *DungeonPartyFinderListenRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	return nil
}

func (m *DungeonPartyFinderListenRequestMessage) Deserialize(r Reader) error {

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	return nil
}

type DungeonPartyFinderRoomContentMessage struct {
	DungeonId uint16

	Players []DungeonPartyFinderPlayer
}

func (m *DungeonPartyFinderRoomContentMessage) ID() uint16 {
	return 6247
}

func (m *DungeonPartyFinderRoomContentMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Players))); err != nil {
		return err
	}

	for i := range m.Players {

		if err := m.Players[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *DungeonPartyFinderRoomContentMessage) Deserialize(r Reader) error {

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	lplayersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Players = make([]DungeonPartyFinderPlayer, lplayersLen)

	for i := range m.Players {

		var lplayers DungeonPartyFinderPlayer

		lplayers.Deserialize(r)

		m.Players[i] = lplayers

	}

	return nil
}

type DungeonPartyFinderListenErrorMessage struct {
	DungeonId uint16
}

func (m *DungeonPartyFinderListenErrorMessage) ID() uint16 {
	return 6248
}

func (m *DungeonPartyFinderListenErrorMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	return nil
}

func (m *DungeonPartyFinderListenErrorMessage) Deserialize(r Reader) error {

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	return nil
}

type DungeonPartyFinderRegisterRequestMessage struct {
	DungeonIds []uint16
}

func (m *DungeonPartyFinderRegisterRequestMessage) ID() uint16 {
	return 6249
}

func (m *DungeonPartyFinderRegisterRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.DungeonIds))); err != nil {
		return err
	}

	for i := range m.DungeonIds {

		if err := w.WriteVarUInt16(m.DungeonIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *DungeonPartyFinderRegisterRequestMessage) Deserialize(r Reader) error {

	ldungeonIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DungeonIds = make([]uint16, ldungeonIdsLen)

	for i := range m.DungeonIds {

		ldungeonIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.DungeonIds[i] = ldungeonIds

	}

	return nil
}

type DungeonPartyFinderRoomContentUpdateMessage struct {
	DungeonId uint16

	AddedPlayers []DungeonPartyFinderPlayer

	RemovedPlayersIds []int64
}

func (m *DungeonPartyFinderRoomContentUpdateMessage) ID() uint16 {
	return 6250
}

func (m *DungeonPartyFinderRoomContentUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.AddedPlayers))); err != nil {
		return err
	}

	for i := range m.AddedPlayers {

		if err := m.AddedPlayers[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.RemovedPlayersIds))); err != nil {
		return err
	}

	for i := range m.RemovedPlayersIds {

		if err := w.WriteVarInt64(m.RemovedPlayersIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *DungeonPartyFinderRoomContentUpdateMessage) Deserialize(r Reader) error {

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	laddedPlayersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AddedPlayers = make([]DungeonPartyFinderPlayer, laddedPlayersLen)

	for i := range m.AddedPlayers {

		var laddedPlayers DungeonPartyFinderPlayer

		laddedPlayers.Deserialize(r)

		m.AddedPlayers[i] = laddedPlayers

	}

	lremovedPlayersIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.RemovedPlayersIds = make([]int64, lremovedPlayersIdsLen)

	for i := range m.RemovedPlayersIds {

		lremovedPlayersIds, err := r.ReadVarInt64()
		if err != nil {
			return err
		}

		m.RemovedPlayersIds[i] = lremovedPlayersIds

	}

	return nil
}

type PartyCancelInvitationNotificationMessage struct {
	AbstractPartyEventMessage

	CancelerId int64

	GuestId int64
}

func (m *PartyCancelInvitationNotificationMessage) ID() uint16 {
	return 6251
}

func (m *PartyCancelInvitationNotificationMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyEventMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.CancelerId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.GuestId); err != nil {
		return err
	}

	return nil
}

func (m *PartyCancelInvitationNotificationMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyEventMessage.Deserialize(r); err != nil {
		return err
	}

	lcancelerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.CancelerId = lcancelerId

	lguestId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.GuestId = lguestId

	return nil
}

type PartyMemberEjectedMessage struct {
	PartyMemberRemoveMessage

	KickerId int64
}

func (m *PartyMemberEjectedMessage) ID() uint16 {
	return 6252
}

func (m *PartyMemberEjectedMessage) Serialize(w Writer) error {

	if err := m.PartyMemberRemoveMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.KickerId); err != nil {
		return err
	}

	return nil
}

func (m *PartyMemberEjectedMessage) Deserialize(r Reader) error {

	if err := m.PartyMemberRemoveMessage.Deserialize(r); err != nil {
		return err
	}

	lkickerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.KickerId = lkickerId

	return nil
}

type RawDataMessage struct {
	Content []uint8
}

func (m *RawDataMessage) ID() uint16 {
	return 6253
}

func (m *RawDataMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt32(int32(len(m.Content))); err != nil {
		return err
	}

	for i := range m.Content {

		if err := w.WriteUInt8(m.Content[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *RawDataMessage) Deserialize(r Reader) error {

	lcontentLen, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Content = make([]uint8, lcontentLen)

	for i := range m.Content {

		lcontent, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.Content[i] = lcontent

	}

	return nil
}

type PartyCancelInvitationMessage struct {
	AbstractPartyMessage

	GuestId int64
}

func (m *PartyCancelInvitationMessage) ID() uint16 {
	return 6254
}

func (m *PartyCancelInvitationMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.GuestId); err != nil {
		return err
	}

	return nil
}

func (m *PartyCancelInvitationMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lguestId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.GuestId = lguestId

	return nil
}

type PartyInvitationCancelledForGuestMessage struct {
	AbstractPartyMessage

	CancelerId int64
}

func (m *PartyInvitationCancelledForGuestMessage) ID() uint16 {
	return 6256
}

func (m *PartyInvitationCancelledForGuestMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.CancelerId); err != nil {
		return err
	}

	return nil
}

func (m *PartyInvitationCancelledForGuestMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lcancelerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.CancelerId = lcancelerId

	return nil
}

type PartyNewGuestMessage struct {
	AbstractPartyEventMessage

	Guest PartyGuestInformations
}

func (m *PartyNewGuestMessage) ID() uint16 {
	return 6260
}

func (m *PartyNewGuestMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyEventMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.Guest.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyNewGuestMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyEventMessage.Deserialize(r); err != nil {
		return err
	}

	var lguest PartyGuestInformations

	lguest.Deserialize(r)

	m.Guest = lguest

	return nil
}

type PartyDeletedMessage struct {
	AbstractPartyMessage
}

func (m *PartyDeletedMessage) ID() uint16 {
	return 6261
}

func (m *PartyDeletedMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyDeletedMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type PartyInvitationDungeonDetailsMessage struct {
	PartyInvitationDetailsMessage

	DungeonId uint16

	PlayersDungeonReady []bool
}

func (m *PartyInvitationDungeonDetailsMessage) ID() uint16 {
	return 6262
}

func (m *PartyInvitationDungeonDetailsMessage) Serialize(w Writer) error {

	if err := m.PartyInvitationDetailsMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.PlayersDungeonReady))); err != nil {
		return err
	}

	for i := range m.PlayersDungeonReady {

		if err := w.WriteBoolean(m.PlayersDungeonReady[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *PartyInvitationDungeonDetailsMessage) Deserialize(r Reader) error {

	if err := m.PartyInvitationDetailsMessage.Deserialize(r); err != nil {
		return err
	}

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	lplayersDungeonReadyLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PlayersDungeonReady = make([]bool, lplayersDungeonReadyLen)

	for i := range m.PlayersDungeonReady {

		lplayersDungeonReady, err := r.ReadBoolean()
		if err != nil {
			return err
		}

		m.PlayersDungeonReady[i] = lplayersDungeonReady

	}

	return nil
}

type PartyInvitationDetailsMessage struct {
	AbstractPartyMessage

	PartyType uint8

	PartyName string

	FromId int64

	FromName string

	LeaderId int64

	Members []PartyInvitationMemberInformations

	Guests []PartyGuestInformations
}

func (m *PartyInvitationDetailsMessage) ID() uint16 {
	return 6263
}

func (m *PartyInvitationDetailsMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.PartyType); err != nil {
		return err
	}

	if err := w.WriteString(m.PartyName); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.FromId); err != nil {
		return err
	}

	if err := w.WriteString(m.FromName); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.LeaderId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Members))); err != nil {
		return err
	}

	for i := range m.Members {

		if err := m.Members[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Guests))); err != nil {
		return err
	}

	for i := range m.Guests {

		if err := m.Guests[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PartyInvitationDetailsMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lpartyType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PartyType = lpartyType

	lpartyName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PartyName = lpartyName

	lfromId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.FromId = lfromId

	lfromName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.FromName = lfromName

	lleaderId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.LeaderId = lleaderId

	lmembersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Members = make([]PartyInvitationMemberInformations, lmembersLen)

	for i := range m.Members {

		var lmembers PartyInvitationMemberInformations

		lmembers.Deserialize(r)

		m.Members[i] = lmembers

	}

	lguestsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Guests = make([]PartyGuestInformations, lguestsLen)

	for i := range m.Guests {

		var lguests PartyGuestInformations

		lguests.Deserialize(r)

		m.Guests[i] = lguests

	}

	return nil
}

type PartyInvitationDetailsRequestMessage struct {
	AbstractPartyMessage
}

func (m *PartyInvitationDetailsRequestMessage) ID() uint16 {
	return 6264
}

func (m *PartyInvitationDetailsRequestMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyInvitationDetailsRequestMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type SpouseStatusMessage struct {
	HasSpouse bool
}

func (m *SpouseStatusMessage) ID() uint16 {
	return 6265
}

func (m *SpouseStatusMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.HasSpouse); err != nil {
		return err
	}

	return nil
}

func (m *SpouseStatusMessage) Deserialize(r Reader) error {

	lhasSpouse, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.HasSpouse = lhasSpouse

	return nil
}

type URLOpenMessage struct {
	UrlId uint8
}

func (m *URLOpenMessage) ID() uint16 {
	return 6266
}

func (m *URLOpenMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.UrlId); err != nil {
		return err
	}

	return nil
}

func (m *URLOpenMessage) Deserialize(r Reader) error {

	lurlId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.UrlId = lurlId

	return nil
}

type TrustStatusMessage struct {
	Trusted bool

	Certified bool
}

func (m *TrustStatusMessage) ID() uint16 {
	return 6267
}

func (m *TrustStatusMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Trusted)

	setWrappedFlag(bbw0, 1, m.Certified)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	return nil
}

func (m *TrustStatusMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Trusted = getWrappedFlag(bbw0, 0)

	m.Certified = getWrappedFlag(bbw0, 1)

	return nil
}

type MapComplementaryInformationsWithCoordsMessage struct {
	MapComplementaryInformationsDataMessage

	WorldX int16

	WorldY int16
}

func (m *MapComplementaryInformationsWithCoordsMessage) ID() uint16 {
	return 6268
}

func (m *MapComplementaryInformationsWithCoordsMessage) Serialize(w Writer) error {

	if err := m.MapComplementaryInformationsDataMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	return nil
}

func (m *MapComplementaryInformationsWithCoordsMessage) Deserialize(r Reader) error {

	if err := m.MapComplementaryInformationsDataMessage.Deserialize(r); err != nil {
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

	return nil
}

type PartyPledgeLoyaltyRequestMessage struct {
	AbstractPartyMessage

	Loyal bool
}

func (m *PartyPledgeLoyaltyRequestMessage) ID() uint16 {
	return 6269
}

func (m *PartyPledgeLoyaltyRequestMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Loyal); err != nil {
		return err
	}

	return nil
}

func (m *PartyPledgeLoyaltyRequestMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lloyal, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Loyal = lloyal

	return nil
}

type PartyLoyaltyStatusMessage struct {
	AbstractPartyMessage

	Loyal bool
}

func (m *PartyLoyaltyStatusMessage) ID() uint16 {
	return 6270
}

func (m *PartyLoyaltyStatusMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Loyal); err != nil {
		return err
	}

	return nil
}

func (m *PartyLoyaltyStatusMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lloyal, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Loyal = lloyal

	return nil
}

type ExchangeBidHouseBuyResultMessage struct {
	Uid uint32

	Bought bool
}

func (m *ExchangeBidHouseBuyResultMessage) ID() uint16 {
	return 6272
}

func (m *ExchangeBidHouseBuyResultMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.Uid); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Bought); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHouseBuyResultMessage) Deserialize(r Reader) error {

	luid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Uid = luid

	lbought, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Bought = lbought

	return nil
}

type AbstractPartyEventMessage struct {
	AbstractPartyMessage
}

func (m *AbstractPartyEventMessage) ID() uint16 {
	return 6273
}

func (m *AbstractPartyEventMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AbstractPartyEventMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type AbstractPartyMessage struct {
	PartyId uint32
}

func (m *AbstractPartyMessage) ID() uint16 {
	return 6274
}

func (m *AbstractPartyMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.PartyId); err != nil {
		return err
	}

	return nil
}

func (m *AbstractPartyMessage) Deserialize(r Reader) error {

	lpartyId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.PartyId = lpartyId

	return nil
}

type MailStatusMessage struct {
	Unread uint16

	Total uint16
}

func (m *MailStatusMessage) ID() uint16 {
	return 6275
}

func (m *MailStatusMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.Unread); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.Total); err != nil {
		return err
	}

	return nil
}

func (m *MailStatusMessage) Deserialize(r Reader) error {

	lunread, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Unread = lunread

	ltotal, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Total = ltotal

	return nil
}

type GameRolePlayArenaFightPropositionMessage struct {
	FightId uint32

	AlliesId []float64

	Duration uint16
}

func (m *GameRolePlayArenaFightPropositionMessage) ID() uint16 {
	return 6276
}

func (m *GameRolePlayArenaFightPropositionMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.AlliesId))); err != nil {
		return err
	}

	for i := range m.AlliesId {

		if err := w.WriteDouble(m.AlliesId[i]); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt16(m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayArenaFightPropositionMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lalliesIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AlliesId = make([]float64, lalliesIdLen)

	for i := range m.AlliesId {

		lalliesId, err := r.ReadDouble()
		if err != nil {
			return err
		}

		m.AlliesId[i] = lalliesId

	}

	lduration, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.Duration = lduration

	return nil
}

type PartyModifiableStatusMessage struct {
	AbstractPartyMessage

	Enabled bool
}

func (m *PartyModifiableStatusMessage) ID() uint16 {
	return 6277
}

func (m *PartyModifiableStatusMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *PartyModifiableStatusMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lenabled, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enabled = lenabled

	return nil
}

type GameRolePlayArenaFightAnswerMessage struct {
	FightId int32

	Accept bool
}

func (m *GameRolePlayArenaFightAnswerMessage) ID() uint16 {
	return 6279
}

func (m *GameRolePlayArenaFightAnswerMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Accept); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayArenaFightAnswerMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	laccept, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Accept = laccept

	return nil
}

type GameRolePlayArenaRegisterMessage struct {
	BattleMode uint32
}

func (m *GameRolePlayArenaRegisterMessage) ID() uint16 {
	return 6280
}

func (m *GameRolePlayArenaRegisterMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.BattleMode); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayArenaRegisterMessage) Deserialize(r Reader) error {

	lbattleMode, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.BattleMode = lbattleMode

	return nil
}

type GameRolePlayArenaFighterStatusMessage struct {
	FightId int32

	PlayerId int32

	Accepted bool
}

func (m *GameRolePlayArenaFighterStatusMessage) ID() uint16 {
	return 6281
}

func (m *GameRolePlayArenaFighterStatusMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.FightId); err != nil {
		return err
	}

	if err := w.WriteInt32(m.PlayerId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Accepted); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayArenaFighterStatusMessage) Deserialize(r Reader) error {

	lfightId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	lplayerId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	laccepted, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Accepted = laccepted

	return nil
}

type GameRolePlayArenaUnregisterMessage struct {
}

func (m *GameRolePlayArenaUnregisterMessage) ID() uint16 {
	return 6282
}

func (m *GameRolePlayArenaUnregisterMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameRolePlayArenaUnregisterMessage) Deserialize(r Reader) error {

	return nil
}

type PartyInvitationArenaRequestMessage struct {
	PartyInvitationRequestMessage
}

func (m *PartyInvitationArenaRequestMessage) ID() uint16 {
	return 6283
}

func (m *PartyInvitationArenaRequestMessage) Serialize(w Writer) error {

	if err := m.PartyInvitationRequestMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyInvitationArenaRequestMessage) Deserialize(r Reader) error {

	if err := m.PartyInvitationRequestMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type GameRolePlayArenaRegistrationStatusMessage struct {
	Registered bool

	Step uint8

	BattleMode uint32
}

func (m *GameRolePlayArenaRegistrationStatusMessage) ID() uint16 {
	return 6284
}

func (m *GameRolePlayArenaRegistrationStatusMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Registered); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Step); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.BattleMode); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayArenaRegistrationStatusMessage) Deserialize(r Reader) error {

	lregistered, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Registered = lregistered

	lstep, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Step = lstep

	lbattleMode, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.BattleMode = lbattleMode

	return nil
}

type TeleportToBuddyOfferMessage struct {
	DungeonId uint16

	BuddyId int64

	TimeLeft uint32
}

func (m *TeleportToBuddyOfferMessage) ID() uint16 {
	return 6287
}

func (m *TeleportToBuddyOfferMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.BuddyId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.TimeLeft); err != nil {
		return err
	}

	return nil
}

func (m *TeleportToBuddyOfferMessage) Deserialize(r Reader) error {

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	lbuddyId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.BuddyId = lbuddyId

	ltimeLeft, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.TimeLeft = ltimeLeft

	return nil
}

type TeleportBuddiesMessage struct {
	DungeonId uint16
}

func (m *TeleportBuddiesMessage) ID() uint16 {
	return 6289
}

func (m *TeleportBuddiesMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	return nil
}

func (m *TeleportBuddiesMessage) Deserialize(r Reader) error {

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	return nil
}

type ObjectFeedMessage struct {
	ObjectUID uint32

	FoodUID uint32

	FoodQuantity uint32
}

func (m *ObjectFeedMessage) ID() uint16 {
	return 6290
}

func (m *ObjectFeedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ObjectUID); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.FoodUID); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.FoodQuantity); err != nil {
		return err
	}

	return nil
}

func (m *ObjectFeedMessage) Deserialize(r Reader) error {

	lobjectUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ObjectUID = lobjectUID

	lfoodUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.FoodUID = lfoodUID

	lfoodQuantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.FoodQuantity = lfoodQuantity

	return nil
}

type NewMailMessage struct {
	MailStatusMessage

	SendersAccountId []uint32
}

func (m *NewMailMessage) ID() uint16 {
	return 6292
}

func (m *NewMailMessage) Serialize(w Writer) error {

	if err := m.MailStatusMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.SendersAccountId))); err != nil {
		return err
	}

	for i := range m.SendersAccountId {

		if err := w.WriteUInt32(m.SendersAccountId[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *NewMailMessage) Deserialize(r Reader) error {

	if err := m.MailStatusMessage.Deserialize(r); err != nil {
		return err
	}

	lsendersAccountIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.SendersAccountId = make([]uint32, lsendersAccountIdLen)

	for i := range m.SendersAccountId {

		lsendersAccountId, err := r.ReadUInt32()
		if err != nil {
			return err
		}

		m.SendersAccountId[i] = lsendersAccountId

	}

	return nil
}

type TeleportToBuddyAnswerMessage struct {
	DungeonId uint16

	BuddyId int64

	Accept bool
}

func (m *TeleportToBuddyAnswerMessage) ID() uint16 {
	return 6293
}

func (m *TeleportToBuddyAnswerMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.BuddyId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Accept); err != nil {
		return err
	}

	return nil
}

func (m *TeleportToBuddyAnswerMessage) Deserialize(r Reader) error {

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	lbuddyId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.BuddyId = lbuddyId

	laccept, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Accept = laccept

	return nil
}

type TeleportBuddiesAnswerMessage struct {
	Accept bool
}

func (m *TeleportBuddiesAnswerMessage) ID() uint16 {
	return 6294
}

func (m *TeleportBuddiesAnswerMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Accept); err != nil {
		return err
	}

	return nil
}

func (m *TeleportBuddiesAnswerMessage) Deserialize(r Reader) error {

	laccept, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Accept = laccept

	return nil
}

type DungeonKeyRingUpdateMessage struct {
	DungeonId uint16

	Available bool
}

func (m *DungeonKeyRingUpdateMessage) ID() uint16 {
	return 6296
}

func (m *DungeonKeyRingUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Available); err != nil {
		return err
	}

	return nil
}

func (m *DungeonKeyRingUpdateMessage) Deserialize(r Reader) error {

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	lavailable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Available = lavailable

	return nil
}

type NumericWhoIsMessage struct {
	PlayerId int64

	AccountId uint32
}

func (m *NumericWhoIsMessage) ID() uint16 {
	return 6297
}

func (m *NumericWhoIsMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	return nil
}

func (m *NumericWhoIsMessage) Deserialize(r Reader) error {

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	return nil
}

type NumericWhoIsRequestMessage struct {
	PlayerId int64
}

func (m *NumericWhoIsRequestMessage) ID() uint16 {
	return 6298
}

func (m *NumericWhoIsRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *NumericWhoIsRequestMessage) Deserialize(r Reader) error {

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type DungeonKeyRingMessage struct {
	Availables []uint16

	Unavailables []uint16
}

func (m *DungeonKeyRingMessage) ID() uint16 {
	return 6299
}

func (m *DungeonKeyRingMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Availables))); err != nil {
		return err
	}

	for i := range m.Availables {

		if err := w.WriteVarUInt16(m.Availables[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Unavailables))); err != nil {
		return err
	}

	for i := range m.Unavailables {

		if err := w.WriteVarUInt16(m.Unavailables[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *DungeonKeyRingMessage) Deserialize(r Reader) error {

	lavailablesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Availables = make([]uint16, lavailablesLen)

	for i := range m.Availables {

		lavailables, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Availables[i] = lavailables

	}

	lunavailablesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Unavailables = make([]uint16, lunavailablesLen)

	for i := range m.Unavailables {

		lunavailables, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Unavailables[i] = lunavailables

	}

	return nil
}

type GameRolePlayArenaUpdatePlayerInfosMessage struct {
	Solo ArenaRankInfos
}

func (m *GameRolePlayArenaUpdatePlayerInfosMessage) ID() uint16 {
	return 6301
}

func (m *GameRolePlayArenaUpdatePlayerInfosMessage) Serialize(w Writer) error {

	if err := m.Solo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayArenaUpdatePlayerInfosMessage) Deserialize(r Reader) error {

	var lsolo ArenaRankInfos

	lsolo.Deserialize(r)

	m.Solo = lsolo

	return nil
}

type TeleportBuddiesRequestedMessage struct {
	DungeonId uint16

	InviterId int64

	InvalidBuddiesIds []int64
}

func (m *TeleportBuddiesRequestedMessage) ID() uint16 {
	return 6302
}

func (m *TeleportBuddiesRequestedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.InviterId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.InvalidBuddiesIds))); err != nil {
		return err
	}

	for i := range m.InvalidBuddiesIds {

		if err := w.WriteVarInt64(m.InvalidBuddiesIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *TeleportBuddiesRequestedMessage) Deserialize(r Reader) error {

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	linviterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.InviterId = linviterId

	linvalidBuddiesIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.InvalidBuddiesIds = make([]int64, linvalidBuddiesIdsLen)

	for i := range m.InvalidBuddiesIds {

		linvalidBuddiesIds, err := r.ReadVarInt64()
		if err != nil {
			return err
		}

		m.InvalidBuddiesIds[i] = linvalidBuddiesIds

	}

	return nil
}

type TeleportToBuddyCloseMessage struct {
	DungeonId uint16

	BuddyId int64
}

func (m *TeleportToBuddyCloseMessage) ID() uint16 {
	return 6303
}

func (m *TeleportToBuddyCloseMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.DungeonId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.BuddyId); err != nil {
		return err
	}

	return nil
}

func (m *TeleportToBuddyCloseMessage) Deserialize(r Reader) error {

	ldungeonId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DungeonId = ldungeonId

	lbuddyId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.BuddyId = lbuddyId

	return nil
}

type GameActionFightModifyEffectsDurationMessage struct {
	AbstractGameActionMessage

	TargetId float64

	Delta int16
}

func (m *GameActionFightModifyEffectsDurationMessage) ID() uint16 {
	return 6304
}

func (m *GameActionFightModifyEffectsDurationMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.Delta); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightModifyEffectsDurationMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	ldelta, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Delta = ldelta

	return nil
}

type ServerOptionalFeaturesMessage struct {
	Features []uint8
}

func (m *ServerOptionalFeaturesMessage) ID() uint16 {
	return 6305
}

func (m *ServerOptionalFeaturesMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Features))); err != nil {
		return err
	}

	for i := range m.Features {

		if err := w.WriteUInt8(m.Features[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ServerOptionalFeaturesMessage) Deserialize(r Reader) error {

	lfeaturesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Features = make([]uint8, lfeaturesLen)

	for i := range m.Features {

		lfeatures, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.Features[i] = lfeatures

	}

	return nil
}

type PartyNewMemberMessage struct {
	PartyUpdateMessage
}

func (m *PartyNewMemberMessage) ID() uint16 {
	return 6306
}

func (m *PartyNewMemberMessage) Serialize(w Writer) error {

	if err := m.PartyUpdateMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PartyNewMemberMessage) Deserialize(r Reader) error {

	if err := m.PartyUpdateMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type GameFightTurnResumeMessage struct {
	GameFightTurnStartMessage

	RemainingTime uint32
}

func (m *GameFightTurnResumeMessage) ID() uint16 {
	return 6307
}

func (m *GameFightTurnResumeMessage) Serialize(w Writer) error {

	if err := m.GameFightTurnStartMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.RemainingTime); err != nil {
		return err
	}

	return nil
}

func (m *GameFightTurnResumeMessage) Deserialize(r Reader) error {

	if err := m.GameFightTurnStartMessage.Deserialize(r); err != nil {
		return err
	}

	lremainingTime, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.RemainingTime = lremainingTime

	return nil
}

type MountReleasedMessage struct {
	MountId int32
}

func (m *MountReleasedMessage) ID() uint16 {
	return 6308
}

func (m *MountReleasedMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt32(m.MountId); err != nil {
		return err
	}

	return nil
}

func (m *MountReleasedMessage) Deserialize(r Reader) error {

	lmountId, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.MountId = lmountId

	return nil
}

type GameFightRefreshFighterMessage struct {
	Informations GameContextActorInformationsIntrf
}

func (m *GameFightRefreshFighterMessage) ID() uint16 {
	return 6309
}

func (m *GameFightRefreshFighterMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Informations.ID()); err != nil {
		return err
	}

	if err := m.Informations.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameFightRefreshFighterMessage) Deserialize(r Reader) error {

	typeinformationsID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	linformations, err := GetType(typeinformationsID)
	if err != nil {
		return err
	}

	linformations.Deserialize(r)

	m.Informations = linformations.(GameContextActorInformationsIntrf)

	return nil
}

type GameActionFightLifeAndShieldPointsLostMessage struct {
	GameActionFightLifePointsLostMessage

	ShieldLoss uint16
}

func (m *GameActionFightLifeAndShieldPointsLostMessage) ID() uint16 {
	return 6310
}

func (m *GameActionFightLifeAndShieldPointsLostMessage) Serialize(w Writer) error {

	if err := m.GameActionFightLifePointsLostMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ShieldLoss); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightLifeAndShieldPointsLostMessage) Deserialize(r Reader) error {

	if err := m.GameActionFightLifePointsLostMessage.Deserialize(r); err != nil {
		return err
	}

	lshieldLoss, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ShieldLoss = lshieldLoss

	return nil
}

type GameActionFightLifePointsGainMessage struct {
	AbstractGameActionMessage

	TargetId float64

	Delta uint32
}

func (m *GameActionFightLifePointsGainMessage) ID() uint16 {
	return 6311
}

func (m *GameActionFightLifePointsGainMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Delta); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightLifePointsGainMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	ldelta, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Delta = ldelta

	return nil
}

type GameActionFightLifePointsLostMessage struct {
	AbstractGameActionMessage

	TargetId float64

	Loss uint32

	PermanentDamages uint32
}

func (m *GameActionFightLifePointsLostMessage) ID() uint16 {
	return 6312
}

func (m *GameActionFightLifePointsLostMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Loss); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.PermanentDamages); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightLifePointsLostMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lloss, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Loss = lloss

	lpermanentDamages, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.PermanentDamages = lpermanentDamages

	return nil
}

type CredentialsAcknowledgementMessage struct {
}

func (m *CredentialsAcknowledgementMessage) ID() uint16 {
	return 6314
}

func (m *CredentialsAcknowledgementMessage) Serialize(w Writer) error {

	return nil
}

func (m *CredentialsAcknowledgementMessage) Deserialize(r Reader) error {

	return nil
}

type AccountHouseMessage struct {
	Houses []AccountHouseInformations
}

func (m *AccountHouseMessage) ID() uint16 {
	return 6315
}

func (m *AccountHouseMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Houses))); err != nil {
		return err
	}

	for i := range m.Houses {

		if err := m.Houses[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *AccountHouseMessage) Deserialize(r Reader) error {

	lhousesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Houses = make([]AccountHouseInformations, lhousesLen)

	for i := range m.Houses {

		var lhouses AccountHouseInformations

		lhouses.Deserialize(r)

		m.Houses[i] = lhouses

	}

	return nil
}

type SequenceNumberRequestMessage struct {
}

func (m *SequenceNumberRequestMessage) ID() uint16 {
	return 6316
}

func (m *SequenceNumberRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *SequenceNumberRequestMessage) Deserialize(r Reader) error {

	return nil
}

type SequenceNumberMessage struct {
	Number uint16
}

func (m *SequenceNumberMessage) ID() uint16 {
	return 6317
}

func (m *SequenceNumberMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Number); err != nil {
		return err
	}

	return nil
}

func (m *SequenceNumberMessage) Deserialize(r Reader) error {

	lnumber, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.Number = lnumber

	return nil
}

type GameActionFightInvisibleDetectedMessage struct {
	AbstractGameActionMessage

	TargetId float64

	CellId int16
}

func (m *GameActionFightInvisibleDetectedMessage) ID() uint16 {
	return 6320
}

func (m *GameActionFightInvisibleDetectedMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.CellId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightInvisibleDetectedMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	lcellId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CellId = lcellId

	return nil
}

type CharacterExperienceGainMessage struct {
	ExperienceCharacter int64

	ExperienceMount int64

	ExperienceGuild int64

	ExperienceIncarnation int64
}

func (m *CharacterExperienceGainMessage) ID() uint16 {
	return 6321
}

func (m *CharacterExperienceGainMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.ExperienceCharacter); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceMount); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceGuild); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.ExperienceIncarnation); err != nil {
		return err
	}

	return nil
}

func (m *CharacterExperienceGainMessage) Deserialize(r Reader) error {

	lexperienceCharacter, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceCharacter = lexperienceCharacter

	lexperienceMount, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceMount = lexperienceMount

	lexperienceGuild, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceGuild = lexperienceGuild

	lexperienceIncarnation, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.ExperienceIncarnation = lexperienceIncarnation

	return nil
}

type FighterStatsListMessage struct {
	Stats CharacterCharacteristicsInformations
}

func (m *FighterStatsListMessage) ID() uint16 {
	return 6322
}

func (m *FighterStatsListMessage) Serialize(w Writer) error {

	if err := m.Stats.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *FighterStatsListMessage) Deserialize(r Reader) error {

	var lstats CharacterCharacteristicsInformations

	lstats.Deserialize(r)

	m.Stats = lstats

	return nil
}

type GuildModificationValidMessage struct {
	GuildName string

	GuildEmblem GuildEmblem
}

func (m *GuildModificationValidMessage) ID() uint16 {
	return 6323
}

func (m *GuildModificationValidMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.GuildName); err != nil {
		return err
	}

	if err := m.GuildEmblem.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildModificationValidMessage) Deserialize(r Reader) error {

	lguildName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.GuildName = lguildName

	var lguildEmblem GuildEmblem

	lguildEmblem.Deserialize(r)

	m.GuildEmblem = lguildEmblem

	return nil
}

type GuildModificationStartedMessage struct {
	CanChangeName bool

	CanChangeEmblem bool
}

func (m *GuildModificationStartedMessage) ID() uint16 {
	return 6324
}

func (m *GuildModificationStartedMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.CanChangeName)

	setWrappedFlag(bbw0, 1, m.CanChangeEmblem)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	return nil
}

func (m *GuildModificationStartedMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CanChangeName = getWrappedFlag(bbw0, 0)

	m.CanChangeEmblem = getWrappedFlag(bbw0, 1)

	return nil
}

type ExchangeObjectTransfertExistingFromInvMessage struct {
}

func (m *ExchangeObjectTransfertExistingFromInvMessage) ID() uint16 {
	return 6325
}

func (m *ExchangeObjectTransfertExistingFromInvMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeObjectTransfertExistingFromInvMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeObjectTransfertExistingToInvMessage struct {
}

func (m *ExchangeObjectTransfertExistingToInvMessage) ID() uint16 {
	return 6326
}

func (m *ExchangeObjectTransfertExistingToInvMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeObjectTransfertExistingToInvMessage) Deserialize(r Reader) error {

	return nil
}

type GuildModificationNameValidMessage struct {
	GuildName string
}

func (m *GuildModificationNameValidMessage) ID() uint16 {
	return 6327
}

func (m *GuildModificationNameValidMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.GuildName); err != nil {
		return err
	}

	return nil
}

func (m *GuildModificationNameValidMessage) Deserialize(r Reader) error {

	lguildName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.GuildName = lguildName

	return nil
}

type GuildModificationEmblemValidMessage struct {
	GuildEmblem GuildEmblem
}

func (m *GuildModificationEmblemValidMessage) ID() uint16 {
	return 6328
}

func (m *GuildModificationEmblemValidMessage) Serialize(w Writer) error {

	if err := m.GuildEmblem.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildModificationEmblemValidMessage) Deserialize(r Reader) error {

	var lguildEmblem GuildEmblem

	lguildEmblem.Deserialize(r)

	m.GuildEmblem = lguildEmblem

	return nil
}

type InventoryPresetSaveCustomMessage struct {
	PresetId uint8

	SymbolId uint8

	ItemsPositions []uint8

	ItemsUids []uint32
}

func (m *InventoryPresetSaveCustomMessage) ID() uint16 {
	return 6329
}

func (m *InventoryPresetSaveCustomMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.SymbolId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.ItemsPositions))); err != nil {
		return err
	}

	for i := range m.ItemsPositions {

		if err := w.WriteUInt8(m.ItemsPositions[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.ItemsUids))); err != nil {
		return err
	}

	for i := range m.ItemsUids {

		if err := w.WriteVarUInt32(m.ItemsUids[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *InventoryPresetSaveCustomMessage) Deserialize(r Reader) error {

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

	litemsPositionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ItemsPositions = make([]uint8, litemsPositionsLen)

	for i := range m.ItemsPositions {

		litemsPositions, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.ItemsPositions[i] = litemsPositions

	}

	litemsUidsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ItemsUids = make([]uint32, litemsUidsLen)

	for i := range m.ItemsUids {

		litemsUids, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.ItemsUids[i] = litemsUids

	}

	return nil
}

type GameActionFightCastOnTargetRequestMessage struct {
	SpellId uint16

	TargetId float64
}

func (m *GameActionFightCastOnTargetRequestMessage) ID() uint16 {
	return 6330
}

func (m *GameActionFightCastOnTargetRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightCastOnTargetRequestMessage) Deserialize(r Reader) error {

	lspellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	ltargetId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type ObjectAveragePricesGetMessage struct {
}

func (m *ObjectAveragePricesGetMessage) ID() uint16 {
	return 6334
}

func (m *ObjectAveragePricesGetMessage) Serialize(w Writer) error {

	return nil
}

func (m *ObjectAveragePricesGetMessage) Deserialize(r Reader) error {

	return nil
}

type ObjectAveragePricesMessage struct {
	Ids []uint16

	AvgPrices []int64
}

func (m *ObjectAveragePricesMessage) ID() uint16 {
	return 6335
}

func (m *ObjectAveragePricesMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Ids))); err != nil {
		return err
	}

	for i := range m.Ids {

		if err := w.WriteVarUInt16(m.Ids[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.AvgPrices))); err != nil {
		return err
	}

	for i := range m.AvgPrices {

		if err := w.WriteVarInt64(m.AvgPrices[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ObjectAveragePricesMessage) Deserialize(r Reader) error {

	lidsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Ids = make([]uint16, lidsLen)

	for i := range m.Ids {

		lids, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Ids[i] = lids

	}

	lavgPricesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AvgPrices = make([]int64, lavgPricesLen)

	for i := range m.AvgPrices {

		lavgPrices, err := r.ReadVarInt64()
		if err != nil {
			return err
		}

		m.AvgPrices[i] = lavgPrices

	}

	return nil
}

type ObjectAveragePricesErrorMessage struct {
}

func (m *ObjectAveragePricesErrorMessage) ID() uint16 {
	return 6336
}

func (m *ObjectAveragePricesErrorMessage) Serialize(w Writer) error {

	return nil
}

func (m *ObjectAveragePricesErrorMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeBidHouseInListUpdatedMessage struct {
	ExchangeBidHouseInListAddedMessage
}

func (m *ExchangeBidHouseInListUpdatedMessage) ID() uint16 {
	return 6337
}

func (m *ExchangeBidHouseInListUpdatedMessage) Serialize(w Writer) error {

	if err := m.ExchangeBidHouseInListAddedMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeBidHouseInListUpdatedMessage) Deserialize(r Reader) error {

	if err := m.ExchangeBidHouseInListAddedMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type CharacterCapabilitiesMessage struct {
	GuildEmblemSymbolCategories uint32
}

func (m *CharacterCapabilitiesMessage) ID() uint16 {
	return 6339
}

func (m *CharacterCapabilitiesMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.GuildEmblemSymbolCategories); err != nil {
		return err
	}

	return nil
}

func (m *CharacterCapabilitiesMessage) Deserialize(r Reader) error {

	lguildEmblemSymbolCategories, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.GuildEmblemSymbolCategories = lguildEmblemSymbolCategories

	return nil
}

type ServerSettingsMessage struct {
	Lang string

	Community uint8

	GameType int8

	ArenaLeaveBanTime uint16
}

func (m *ServerSettingsMessage) ID() uint16 {
	return 6340
}

func (m *ServerSettingsMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Lang); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Community); err != nil {
		return err
	}

	if err := w.WriteInt8(m.GameType); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ArenaLeaveBanTime); err != nil {
		return err
	}

	return nil
}

func (m *ServerSettingsMessage) Deserialize(r Reader) error {

	llang, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Lang = llang

	lcommunity, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Community = lcommunity

	lgameType, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.GameType = lgameType

	larenaLeaveBanTime, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ArenaLeaveBanTime = larenaLeaveBanTime

	return nil
}

type AlmanachCalendarDateMessage struct {
	Date int32
}

func (m *AlmanachCalendarDateMessage) ID() uint16 {
	return 6341
}

func (m *AlmanachCalendarDateMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.Date); err != nil {
		return err
	}

	return nil
}

func (m *AlmanachCalendarDateMessage) Deserialize(r Reader) error {

	ldate, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Date = ldate

	return nil
}

type PartyMemberInFightMessage struct {
	AbstractPartyMessage

	Reason uint8

	MemberId int64

	MemberAccountId uint32

	MemberName string

	FightId int32

	FightMap MapCoordinatesExtended

	TimeBeforeFightStart int16
}

func (m *PartyMemberInFightMessage) ID() uint16 {
	return 6342
}

func (m *PartyMemberInFightMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.MemberId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.MemberAccountId); err != nil {
		return err
	}

	if err := w.WriteString(m.MemberName); err != nil {
		return err
	}

	if err := w.WriteInt32(m.FightId); err != nil {
		return err
	}

	if err := m.FightMap.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt16(m.TimeBeforeFightStart); err != nil {
		return err
	}

	return nil
}

func (m *PartyMemberInFightMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	lmemberId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.MemberId = lmemberId

	lmemberAccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MemberAccountId = lmemberAccountId

	lmemberName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.MemberName = lmemberName

	lfightId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.FightId = lfightId

	var lfightMap MapCoordinatesExtended

	lfightMap.Deserialize(r)

	m.FightMap = lfightMap

	ltimeBeforeFightStart, err := r.ReadVarInt16()
	if err != nil {
		return err
	}

	m.TimeBeforeFightStart = ltimeBeforeFightStart

	return nil
}

type KrosmasterInventoryErrorMessage struct {
	Reason uint8
}

func (m *KrosmasterInventoryErrorMessage) ID() uint16 {
	return 6343
}

func (m *KrosmasterInventoryErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *KrosmasterInventoryErrorMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type KrosmasterInventoryRequestMessage struct {
}

func (m *KrosmasterInventoryRequestMessage) ID() uint16 {
	return 6344
}

func (m *KrosmasterInventoryRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *KrosmasterInventoryRequestMessage) Deserialize(r Reader) error {

	return nil
}

type KrosmasterAuthTokenErrorMessage struct {
	Reason uint8
}

func (m *KrosmasterAuthTokenErrorMessage) ID() uint16 {
	return 6345
}

func (m *KrosmasterAuthTokenErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *KrosmasterAuthTokenErrorMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type KrosmasterAuthTokenRequestMessage struct {
}

func (m *KrosmasterAuthTokenRequestMessage) ID() uint16 {
	return 6346
}

func (m *KrosmasterAuthTokenRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *KrosmasterAuthTokenRequestMessage) Deserialize(r Reader) error {

	return nil
}

type KrosmasterPlayingStatusMessage struct {
	Playing bool
}

func (m *KrosmasterPlayingStatusMessage) ID() uint16 {
	return 6347
}

func (m *KrosmasterPlayingStatusMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Playing); err != nil {
		return err
	}

	return nil
}

func (m *KrosmasterPlayingStatusMessage) Deserialize(r Reader) error {

	lplaying, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Playing = lplaying

	return nil
}

type KrosmasterTransferMessage struct {
	Uid string

	Failure uint8
}

func (m *KrosmasterTransferMessage) ID() uint16 {
	return 6348
}

func (m *KrosmasterTransferMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Uid); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Failure); err != nil {
		return err
	}

	return nil
}

func (m *KrosmasterTransferMessage) Deserialize(r Reader) error {

	luid, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Uid = luid

	lfailure, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Failure = lfailure

	return nil
}

type KrosmasterTransferRequestMessage struct {
	Uid string
}

func (m *KrosmasterTransferRequestMessage) ID() uint16 {
	return 6349
}

func (m *KrosmasterTransferRequestMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Uid); err != nil {
		return err
	}

	return nil
}

func (m *KrosmasterTransferRequestMessage) Deserialize(r Reader) error {

	luid, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Uid = luid

	return nil
}

type KrosmasterInventoryMessage struct {
	Figures []KrosmasterFigure
}

func (m *KrosmasterInventoryMessage) ID() uint16 {
	return 6350
}

func (m *KrosmasterInventoryMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Figures))); err != nil {
		return err
	}

	for i := range m.Figures {

		if err := m.Figures[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *KrosmasterInventoryMessage) Deserialize(r Reader) error {

	lfiguresLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Figures = make([]KrosmasterFigure, lfiguresLen)

	for i := range m.Figures {

		var lfigures KrosmasterFigure

		lfigures.Deserialize(r)

		m.Figures[i] = lfigures

	}

	return nil
}

type KrosmasterAuthTokenMessage struct {
	Token string
}

func (m *KrosmasterAuthTokenMessage) ID() uint16 {
	return 6351
}

func (m *KrosmasterAuthTokenMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Token); err != nil {
		return err
	}

	return nil
}

func (m *KrosmasterAuthTokenMessage) Deserialize(r Reader) error {

	ltoken, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Token = ltoken

	return nil
}

type SpouseGetInformationsMessage struct {
}

func (m *SpouseGetInformationsMessage) ID() uint16 {
	return 6355
}

func (m *SpouseGetInformationsMessage) Serialize(w Writer) error {

	return nil
}

func (m *SpouseGetInformationsMessage) Deserialize(r Reader) error {

	return nil
}

type SpouseInformationsMessage struct {
	Spouse FriendSpouseInformationsIntrf
}

func (m *SpouseInformationsMessage) ID() uint16 {
	return 6356
}

func (m *SpouseInformationsMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Spouse.ID()); err != nil {
		return err
	}

	if err := m.Spouse.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *SpouseInformationsMessage) Deserialize(r Reader) error {

	typespouseID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lspouse, err := GetType(typespouseID)
	if err != nil {
		return err
	}

	lspouse.Deserialize(r)

	m.Spouse = lspouse.(FriendSpouseInformationsIntrf)

	return nil
}

type AchievementDetailedListRequestMessage struct {
	CategoryId uint16
}

func (m *AchievementDetailedListRequestMessage) ID() uint16 {
	return 6357
}

func (m *AchievementDetailedListRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.CategoryId); err != nil {
		return err
	}

	return nil
}

func (m *AchievementDetailedListRequestMessage) Deserialize(r Reader) error {

	lcategoryId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.CategoryId = lcategoryId

	return nil
}

type AchievementDetailedListMessage struct {
	StartedAchievements []Achievement

	FinishedAchievements []Achievement
}

func (m *AchievementDetailedListMessage) ID() uint16 {
	return 6358
}

func (m *AchievementDetailedListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.StartedAchievements))); err != nil {
		return err
	}

	for i := range m.StartedAchievements {

		if err := m.StartedAchievements[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.FinishedAchievements))); err != nil {
		return err
	}

	for i := range m.FinishedAchievements {

		if err := m.FinishedAchievements[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *AchievementDetailedListMessage) Deserialize(r Reader) error {

	lstartedAchievementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.StartedAchievements = make([]Achievement, lstartedAchievementsLen)

	for i := range m.StartedAchievements {

		var lstartedAchievements Achievement

		lstartedAchievements.Deserialize(r)

		m.StartedAchievements[i] = lstartedAchievements

	}

	lfinishedAchievementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FinishedAchievements = make([]Achievement, lfinishedAchievementsLen)

	for i := range m.FinishedAchievements {

		var lfinishedAchievements Achievement

		lfinishedAchievements.Deserialize(r)

		m.FinishedAchievements[i] = lfinishedAchievements

	}

	return nil
}

type BasicAckMessage struct {
	Seq uint32

	LastPacketId uint16
}

func (m *BasicAckMessage) ID() uint16 {
	return 6362
}

func (m *BasicAckMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.Seq); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.LastPacketId); err != nil {
		return err
	}

	return nil
}

func (m *BasicAckMessage) Deserialize(r Reader) error {

	lseq, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Seq = lseq

	llastPacketId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.LastPacketId = llastPacketId

	return nil
}

type TitlesAndOrnamentsListRequestMessage struct {
}

func (m *TitlesAndOrnamentsListRequestMessage) ID() uint16 {
	return 6363
}

func (m *TitlesAndOrnamentsListRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *TitlesAndOrnamentsListRequestMessage) Deserialize(r Reader) error {

	return nil
}

type TitleGainedMessage struct {
	TitleId uint16
}

func (m *TitleGainedMessage) ID() uint16 {
	return 6364
}

func (m *TitleGainedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.TitleId); err != nil {
		return err
	}

	return nil
}

func (m *TitleGainedMessage) Deserialize(r Reader) error {

	ltitleId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TitleId = ltitleId

	return nil
}

type TitleSelectRequestMessage struct {
	TitleId uint16
}

func (m *TitleSelectRequestMessage) ID() uint16 {
	return 6365
}

func (m *TitleSelectRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.TitleId); err != nil {
		return err
	}

	return nil
}

func (m *TitleSelectRequestMessage) Deserialize(r Reader) error {

	ltitleId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TitleId = ltitleId

	return nil
}

type TitleSelectedMessage struct {
	TitleId uint16
}

func (m *TitleSelectedMessage) ID() uint16 {
	return 6366
}

func (m *TitleSelectedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.TitleId); err != nil {
		return err
	}

	return nil
}

func (m *TitleSelectedMessage) Deserialize(r Reader) error {

	ltitleId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TitleId = ltitleId

	return nil
}

type TitlesAndOrnamentsListMessage struct {
	Titles []uint16

	Ornaments []uint16

	ActiveTitle uint16

	ActiveOrnament uint16
}

func (m *TitlesAndOrnamentsListMessage) ID() uint16 {
	return 6367
}

func (m *TitlesAndOrnamentsListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Titles))); err != nil {
		return err
	}

	for i := range m.Titles {

		if err := w.WriteVarUInt16(m.Titles[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Ornaments))); err != nil {
		return err
	}

	for i := range m.Ornaments {

		if err := w.WriteVarUInt16(m.Ornaments[i]); err != nil {
			return err
		}

	}

	if err := w.WriteVarUInt16(m.ActiveTitle); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ActiveOrnament); err != nil {
		return err
	}

	return nil
}

func (m *TitlesAndOrnamentsListMessage) Deserialize(r Reader) error {

	ltitlesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Titles = make([]uint16, ltitlesLen)

	for i := range m.Titles {

		ltitles, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Titles[i] = ltitles

	}

	lornamentsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Ornaments = make([]uint16, lornamentsLen)

	for i := range m.Ornaments {

		lornaments, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Ornaments[i] = lornaments

	}

	lactiveTitle, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ActiveTitle = lactiveTitle

	lactiveOrnament, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ActiveOrnament = lactiveOrnament

	return nil
}

type OrnamentGainedMessage struct {
	OrnamentId uint16
}

func (m *OrnamentGainedMessage) ID() uint16 {
	return 6368
}

func (m *OrnamentGainedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.OrnamentId); err != nil {
		return err
	}

	return nil
}

func (m *OrnamentGainedMessage) Deserialize(r Reader) error {

	lornamentId, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.OrnamentId = lornamentId

	return nil
}

type OrnamentSelectedMessage struct {
	OrnamentId uint16
}

func (m *OrnamentSelectedMessage) ID() uint16 {
	return 6369
}

func (m *OrnamentSelectedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.OrnamentId); err != nil {
		return err
	}

	return nil
}

func (m *OrnamentSelectedMessage) Deserialize(r Reader) error {

	lornamentId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.OrnamentId = lornamentId

	return nil
}

type OrnamentSelectErrorMessage struct {
	Reason uint8
}

func (m *OrnamentSelectErrorMessage) ID() uint16 {
	return 6370
}

func (m *OrnamentSelectErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *OrnamentSelectErrorMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type TitleLostMessage struct {
	TitleId uint16
}

func (m *TitleLostMessage) ID() uint16 {
	return 6371
}

func (m *TitleLostMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.TitleId); err != nil {
		return err
	}

	return nil
}

func (m *TitleLostMessage) Deserialize(r Reader) error {

	ltitleId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.TitleId = ltitleId

	return nil
}

type CheckIntegrityMessage struct {
	Data []int8
}

func (m *CheckIntegrityMessage) ID() uint16 {
	return 6372
}

func (m *CheckIntegrityMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt32(int32(len(m.Data))); err != nil {
		return err
	}

	for i := range m.Data {

		if err := w.WriteInt8(m.Data[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *CheckIntegrityMessage) Deserialize(r Reader) error {

	ldataLen, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Data = make([]int8, ldataLen)

	for i := range m.Data {

		ldata, err := r.ReadInt8()
		if err != nil {
			return err
		}

		m.Data[i] = ldata

	}

	return nil
}

type TitleSelectErrorMessage struct {
	Reason uint8
}

func (m *TitleSelectErrorMessage) ID() uint16 {
	return 6373
}

func (m *TitleSelectErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *TitleSelectErrorMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type OrnamentSelectRequestMessage struct {
	OrnamentId uint16
}

func (m *OrnamentSelectRequestMessage) ID() uint16 {
	return 6374
}

func (m *OrnamentSelectRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.OrnamentId); err != nil {
		return err
	}

	return nil
}

func (m *OrnamentSelectRequestMessage) Deserialize(r Reader) error {

	lornamentId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.OrnamentId = lornamentId

	return nil
}

type AchievementRewardErrorMessage struct {
	AchievementId int16
}

func (m *AchievementRewardErrorMessage) ID() uint16 {
	return 6375
}

func (m *AchievementRewardErrorMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(m.AchievementId); err != nil {
		return err
	}

	return nil
}

func (m *AchievementRewardErrorMessage) Deserialize(r Reader) error {

	lachievementId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AchievementId = lachievementId

	return nil
}

type AchievementRewardSuccessMessage struct {
	AchievementId int16
}

func (m *AchievementRewardSuccessMessage) ID() uint16 {
	return 6376
}

func (m *AchievementRewardSuccessMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(m.AchievementId); err != nil {
		return err
	}

	return nil
}

func (m *AchievementRewardSuccessMessage) Deserialize(r Reader) error {

	lachievementId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AchievementId = lachievementId

	return nil
}

type AchievementRewardRequestMessage struct {
	AchievementId int16
}

func (m *AchievementRewardRequestMessage) ID() uint16 {
	return 6377
}

func (m *AchievementRewardRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(m.AchievementId); err != nil {
		return err
	}

	return nil
}

func (m *AchievementRewardRequestMessage) Deserialize(r Reader) error {

	lachievementId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AchievementId = lachievementId

	return nil
}

type AchievementDetailsMessage struct {
	Achievement Achievement
}

func (m *AchievementDetailsMessage) ID() uint16 {
	return 6378
}

func (m *AchievementDetailsMessage) Serialize(w Writer) error {

	if err := m.Achievement.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AchievementDetailsMessage) Deserialize(r Reader) error {

	var lachievement Achievement

	lachievement.Deserialize(r)

	m.Achievement = lachievement

	return nil
}

type AchievementDetailsRequestMessage struct {
	AchievementId uint16
}

func (m *AchievementDetailsRequestMessage) ID() uint16 {
	return 6380
}

func (m *AchievementDetailsRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.AchievementId); err != nil {
		return err
	}

	return nil
}

func (m *AchievementDetailsRequestMessage) Deserialize(r Reader) error {

	lachievementId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.AchievementId = lachievementId

	return nil
}

type AchievementFinishedInformationMessage struct {
	AchievementFinishedMessage

	Name string

	PlayerId int64
}

func (m *AchievementFinishedInformationMessage) ID() uint16 {
	return 6381
}

func (m *AchievementFinishedInformationMessage) Serialize(w Writer) error {

	if err := m.AchievementFinishedMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *AchievementFinishedInformationMessage) Deserialize(r Reader) error {

	if err := m.AchievementFinishedMessage.Deserialize(r); err != nil {
		return err
	}

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type FriendGuildSetWarnOnAchievementCompleteMessage struct {
	Enable bool
}

func (m *FriendGuildSetWarnOnAchievementCompleteMessage) ID() uint16 {
	return 6382
}

func (m *FriendGuildSetWarnOnAchievementCompleteMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *FriendGuildSetWarnOnAchievementCompleteMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type FriendGuildWarnOnAchievementCompleteStateMessage struct {
	Enable bool
}

func (m *FriendGuildWarnOnAchievementCompleteStateMessage) ID() uint16 {
	return 6383
}

func (m *FriendGuildWarnOnAchievementCompleteStateMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *FriendGuildWarnOnAchievementCompleteStateMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type InteractiveUseErrorMessage struct {
	ElemId uint32

	SkillInstanceUid uint32
}

func (m *InteractiveUseErrorMessage) ID() uint16 {
	return 6384
}

func (m *InteractiveUseErrorMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.ElemId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.SkillInstanceUid); err != nil {
		return err
	}

	return nil
}

func (m *InteractiveUseErrorMessage) Deserialize(r Reader) error {

	lelemId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.ElemId = lelemId

	lskillInstanceUid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SkillInstanceUid = lskillInstanceUid

	return nil
}

type PlayerStatusUpdateErrorMessage struct {
}

func (m *PlayerStatusUpdateErrorMessage) ID() uint16 {
	return 6385
}

func (m *PlayerStatusUpdateErrorMessage) Serialize(w Writer) error {

	return nil
}

func (m *PlayerStatusUpdateErrorMessage) Deserialize(r Reader) error {

	return nil
}

type PlayerStatusUpdateMessage struct {
	AccountId uint32

	PlayerId int64

	Status PlayerStatusIntrf
}

func (m *PlayerStatusUpdateMessage) ID() uint16 {
	return 6386
}

func (m *PlayerStatusUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
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

func (m *PlayerStatusUpdateMessage) Deserialize(r Reader) error {

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	typestatusID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lstatus, err := GetType(typestatusID)
	if err != nil {
		return err
	}

	lstatus.Deserialize(r)

	m.Status = lstatus.(PlayerStatusIntrf)

	return nil
}

type PlayerStatusUpdateRequestMessage struct {
	Status PlayerStatusIntrf
}

func (m *PlayerStatusUpdateRequestMessage) ID() uint16 {
	return 6387
}

func (m *PlayerStatusUpdateRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Status.ID()); err != nil {
		return err
	}

	if err := m.Status.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PlayerStatusUpdateRequestMessage) Deserialize(r Reader) error {

	typestatusID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lstatus, err := GetType(typestatusID)
	if err != nil {
		return err
	}

	lstatus.Deserialize(r)

	m.Status = lstatus.(PlayerStatusIntrf)

	return nil
}

type MoodSmileyUpdateMessage struct {
	AccountId uint32

	PlayerId int64

	SmileyId uint16
}

func (m *MoodSmileyUpdateMessage) ID() uint16 {
	return 6388
}

func (m *MoodSmileyUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.AccountId); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.SmileyId); err != nil {
		return err
	}

	return nil
}

func (m *MoodSmileyUpdateMessage) Deserialize(r Reader) error {

	laccountId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.AccountId = laccountId

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	lsmileyId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SmileyId = lsmileyId

	return nil
}

type ExchangeSetCraftRecipeMessage struct {
	ObjectGID uint16
}

func (m *ExchangeSetCraftRecipeMessage) ID() uint16 {
	return 6389
}

func (m *ExchangeSetCraftRecipeMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ObjectGID); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeSetCraftRecipeMessage) Deserialize(r Reader) error {

	lobjectGID, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGID = lobjectGID

	return nil
}

type AllianceMembershipMessage struct {
	AllianceJoinedMessage
}

func (m *AllianceMembershipMessage) ID() uint16 {
	return 6390
}

func (m *AllianceMembershipMessage) Serialize(w Writer) error {

	if err := m.AllianceJoinedMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AllianceMembershipMessage) Deserialize(r Reader) error {

	if err := m.AllianceJoinedMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type AllianceCreationResultMessage struct {
	Result uint8
}

func (m *AllianceCreationResultMessage) ID() uint16 {
	return 6391
}

func (m *AllianceCreationResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Result); err != nil {
		return err
	}

	return nil
}

func (m *AllianceCreationResultMessage) Deserialize(r Reader) error {

	lresult, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Result = lresult

	return nil
}

type AllianceInvitationStateRecrutedMessage struct {
	InvitationState uint8
}

func (m *AllianceInvitationStateRecrutedMessage) ID() uint16 {
	return 6392
}

func (m *AllianceInvitationStateRecrutedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.InvitationState); err != nil {
		return err
	}

	return nil
}

func (m *AllianceInvitationStateRecrutedMessage) Deserialize(r Reader) error {

	linvitationState, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.InvitationState = linvitationState

	return nil
}

type AllianceCreationValidMessage struct {
	AllianceName string

	AllianceTag string

	AllianceEmblem GuildEmblem
}

func (m *AllianceCreationValidMessage) ID() uint16 {
	return 6393
}

func (m *AllianceCreationValidMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.AllianceName); err != nil {
		return err
	}

	if err := w.WriteString(m.AllianceTag); err != nil {
		return err
	}

	if err := m.AllianceEmblem.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AllianceCreationValidMessage) Deserialize(r Reader) error {

	lallianceName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.AllianceName = lallianceName

	lallianceTag, err := r.ReadString()
	if err != nil {
		return err
	}

	m.AllianceTag = lallianceTag

	var lallianceEmblem GuildEmblem

	lallianceEmblem.Deserialize(r)

	m.AllianceEmblem = lallianceEmblem

	return nil
}

type AllianceCreationStartedMessage struct {
}

func (m *AllianceCreationStartedMessage) ID() uint16 {
	return 6394
}

func (m *AllianceCreationStartedMessage) Serialize(w Writer) error {

	return nil
}

func (m *AllianceCreationStartedMessage) Deserialize(r Reader) error {

	return nil
}

type AllianceInvitationMessage struct {
	TargetId int64
}

func (m *AllianceInvitationMessage) ID() uint16 {
	return 6395
}

func (m *AllianceInvitationMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.TargetId); err != nil {
		return err
	}

	return nil
}

func (m *AllianceInvitationMessage) Deserialize(r Reader) error {

	ltargetId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.TargetId = ltargetId

	return nil
}

type AllianceInvitationStateRecruterMessage struct {
	RecrutedName string

	InvitationState uint8
}

func (m *AllianceInvitationStateRecruterMessage) ID() uint16 {
	return 6396
}

func (m *AllianceInvitationStateRecruterMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.RecrutedName); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.InvitationState); err != nil {
		return err
	}

	return nil
}

func (m *AllianceInvitationStateRecruterMessage) Deserialize(r Reader) error {

	lrecrutedName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.RecrutedName = lrecrutedName

	linvitationState, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.InvitationState = linvitationState

	return nil
}

type AllianceInvitedMessage struct {
	RecruterId int64

	RecruterName string

	AllianceInfo BasicNamedAllianceInformations
}

func (m *AllianceInvitedMessage) ID() uint16 {
	return 6397
}

func (m *AllianceInvitedMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.RecruterId); err != nil {
		return err
	}

	if err := w.WriteString(m.RecruterName); err != nil {
		return err
	}

	if err := m.AllianceInfo.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AllianceInvitedMessage) Deserialize(r Reader) error {

	lrecruterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.RecruterId = lrecruterId

	lrecruterName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.RecruterName = lrecruterName

	var lallianceInfo BasicNamedAllianceInformations

	lallianceInfo.Deserialize(r)

	m.AllianceInfo = lallianceInfo

	return nil
}

type AllianceLeftMessage struct {
}

func (m *AllianceLeftMessage) ID() uint16 {
	return 6398
}

func (m *AllianceLeftMessage) Serialize(w Writer) error {

	return nil
}

func (m *AllianceLeftMessage) Deserialize(r Reader) error {

	return nil
}

type AllianceGuildLeavingMessage struct {
	Kicked bool

	GuildId uint32
}

func (m *AllianceGuildLeavingMessage) ID() uint16 {
	return 6399
}

func (m *AllianceGuildLeavingMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Kicked); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.GuildId); err != nil {
		return err
	}

	return nil
}

func (m *AllianceGuildLeavingMessage) Deserialize(r Reader) error {

	lkicked, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Kicked = lkicked

	lguildId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.GuildId = lguildId

	return nil
}

type AllianceKickRequestMessage struct {
	KickedId uint32
}

func (m *AllianceKickRequestMessage) ID() uint16 {
	return 6400
}

func (m *AllianceKickRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.KickedId); err != nil {
		return err
	}

	return nil
}

func (m *AllianceKickRequestMessage) Deserialize(r Reader) error {

	lkickedId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.KickedId = lkickedId

	return nil
}

type AllianceInvitationAnswerMessage struct {
	Accept bool
}

func (m *AllianceInvitationAnswerMessage) ID() uint16 {
	return 6401
}

func (m *AllianceInvitationAnswerMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Accept); err != nil {
		return err
	}

	return nil
}

func (m *AllianceInvitationAnswerMessage) Deserialize(r Reader) error {

	laccept, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Accept = laccept

	return nil
}

type AllianceJoinedMessage struct {
	AllianceInfo AllianceInformations

	Enabled bool

	LeadingGuildId uint32
}

func (m *AllianceJoinedMessage) ID() uint16 {
	return 6402
}

func (m *AllianceJoinedMessage) Serialize(w Writer) error {

	if err := m.AllianceInfo.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Enabled); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.LeadingGuildId); err != nil {
		return err
	}

	return nil
}

func (m *AllianceJoinedMessage) Deserialize(r Reader) error {

	var lallianceInfo AllianceInformations

	lallianceInfo.Deserialize(r)

	m.AllianceInfo = lallianceInfo

	lenabled, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enabled = lenabled

	lleadingGuildId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.LeadingGuildId = lleadingGuildId

	return nil
}

type AllianceInsiderInfoMessage struct {
	AllianceInfos AllianceFactSheetInformations

	Guilds []GuildInsiderFactSheetInformations

	Prisms []PrismSubareaEmptyInfoIntrf
}

func (m *AllianceInsiderInfoMessage) ID() uint16 {
	return 6403
}

func (m *AllianceInsiderInfoMessage) Serialize(w Writer) error {

	if err := m.AllianceInfos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Guilds))); err != nil {
		return err
	}

	for i := range m.Guilds {

		if err := m.Guilds[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Prisms))); err != nil {
		return err
	}

	for i := range m.Prisms {

		if err := w.WriteUInt16(m.Prisms[i].ID()); err != nil {
			return err
		}

		if err := m.Prisms[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *AllianceInsiderInfoMessage) Deserialize(r Reader) error {

	var lallianceInfos AllianceFactSheetInformations

	lallianceInfos.Deserialize(r)

	m.AllianceInfos = lallianceInfos

	lguildsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Guilds = make([]GuildInsiderFactSheetInformations, lguildsLen)

	for i := range m.Guilds {

		var lguilds GuildInsiderFactSheetInformations

		lguilds.Deserialize(r)

		m.Guilds[i] = lguilds

	}

	lprismsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Prisms = make([]PrismSubareaEmptyInfoIntrf, lprismsLen)

	for i := range m.Prisms {

		typeprismsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lprisms, err := GetType(typeprismsID)
		if err != nil {
			return err
		}

		lprisms.Deserialize(r)

		m.Prisms[i] = lprisms.(PrismSubareaEmptyInfoIntrf)

	}

	return nil
}

type GuildFactsRequestMessage struct {
	GuildId uint32
}

func (m *GuildFactsRequestMessage) ID() uint16 {
	return 6404
}

func (m *GuildFactsRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.GuildId); err != nil {
		return err
	}

	return nil
}

func (m *GuildFactsRequestMessage) Deserialize(r Reader) error {

	lguildId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.GuildId = lguildId

	return nil
}

type GameRolePlayShowActorWithEventMessage struct {
	GameRolePlayShowActorMessage

	ActorEventId uint8
}

func (m *GameRolePlayShowActorWithEventMessage) ID() uint16 {
	return 6407
}

func (m *GameRolePlayShowActorWithEventMessage) Serialize(w Writer) error {

	if err := m.GameRolePlayShowActorMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.ActorEventId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayShowActorWithEventMessage) Deserialize(r Reader) error {

	if err := m.GameRolePlayShowActorMessage.Deserialize(r); err != nil {
		return err
	}

	lactorEventId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ActorEventId = lactorEventId

	return nil
}

type AllianceListMessage struct {
	Alliances []AllianceFactSheetInformations
}

func (m *AllianceListMessage) ID() uint16 {
	return 6408
}

func (m *AllianceListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Alliances))); err != nil {
		return err
	}

	for i := range m.Alliances {

		if err := m.Alliances[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *AllianceListMessage) Deserialize(r Reader) error {

	lalliancesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Alliances = make([]AllianceFactSheetInformations, lalliancesLen)

	for i := range m.Alliances {

		var lalliances AllianceFactSheetInformations

		lalliances.Deserialize(r)

		m.Alliances[i] = lalliances

	}

	return nil
}

type AllianceFactsRequestMessage struct {
	AllianceId uint32
}

func (m *AllianceFactsRequestMessage) ID() uint16 {
	return 6409
}

func (m *AllianceFactsRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.AllianceId); err != nil {
		return err
	}

	return nil
}

func (m *AllianceFactsRequestMessage) Deserialize(r Reader) error {

	lallianceId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.AllianceId = lallianceId

	return nil
}

type GameContextRemoveElementWithEventMessage struct {
	GameContextRemoveElementMessage

	ElementEventId uint8
}

func (m *GameContextRemoveElementWithEventMessage) ID() uint16 {
	return 6412
}

func (m *GameContextRemoveElementWithEventMessage) Serialize(w Writer) error {

	if err := m.GameContextRemoveElementMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.ElementEventId); err != nil {
		return err
	}

	return nil
}

func (m *GameContextRemoveElementWithEventMessage) Deserialize(r Reader) error {

	if err := m.GameContextRemoveElementMessage.Deserialize(r); err != nil {
		return err
	}

	lelementEventId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ElementEventId = lelementEventId

	return nil
}

type GuildListMessage struct {
	Guilds []GuildInformations
}

func (m *GuildListMessage) ID() uint16 {
	return 6413
}

func (m *GuildListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Guilds))); err != nil {
		return err
	}

	for i := range m.Guilds {

		if err := m.Guilds[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GuildListMessage) Deserialize(r Reader) error {

	lguildsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Guilds = make([]GuildInformations, lguildsLen)

	for i := range m.Guilds {

		var lguilds GuildInformations

		lguilds.Deserialize(r)

		m.Guilds[i] = lguilds

	}

	return nil
}

type AllianceFactsMessage struct {
	Infos AllianceFactSheetInformationsIntrf

	Guilds []GuildInAllianceInformations

	ControlledSubareaIds []uint16

	LeaderCharacterId int64

	LeaderCharacterName string
}

func (m *AllianceFactsMessage) ID() uint16 {
	return 6414
}

func (m *AllianceFactsMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Infos.ID()); err != nil {
		return err
	}

	if err := m.Infos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Guilds))); err != nil {
		return err
	}

	for i := range m.Guilds {

		if err := m.Guilds[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.ControlledSubareaIds))); err != nil {
		return err
	}

	for i := range m.ControlledSubareaIds {

		if err := w.WriteVarUInt16(m.ControlledSubareaIds[i]); err != nil {
			return err
		}

	}

	if err := w.WriteVarInt64(m.LeaderCharacterId); err != nil {
		return err
	}

	if err := w.WriteString(m.LeaderCharacterName); err != nil {
		return err
	}

	return nil
}

func (m *AllianceFactsMessage) Deserialize(r Reader) error {

	typeinfosID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	linfos, err := GetType(typeinfosID)
	if err != nil {
		return err
	}

	linfos.Deserialize(r)

	m.Infos = linfos.(AllianceFactSheetInformationsIntrf)

	lguildsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Guilds = make([]GuildInAllianceInformations, lguildsLen)

	for i := range m.Guilds {

		var lguilds GuildInAllianceInformations

		lguilds.Deserialize(r)

		m.Guilds[i] = lguilds

	}

	lcontrolledSubareaIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ControlledSubareaIds = make([]uint16, lcontrolledSubareaIdsLen)

	for i := range m.ControlledSubareaIds {

		lcontrolledSubareaIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.ControlledSubareaIds[i] = lcontrolledSubareaIds

	}

	lleaderCharacterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.LeaderCharacterId = lleaderCharacterId

	lleaderCharacterName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.LeaderCharacterName = lleaderCharacterName

	return nil
}

type GuildFactsMessage struct {
	Infos GuildFactSheetInformationsIntrf

	CreationDate uint32

	NbTaxCollectors uint16

	Members []CharacterMinimalInformations
}

func (m *GuildFactsMessage) ID() uint16 {
	return 6415
}

func (m *GuildFactsMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.Infos.ID()); err != nil {
		return err
	}

	if err := m.Infos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.CreationDate); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.NbTaxCollectors); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Members))); err != nil {
		return err
	}

	for i := range m.Members {

		if err := m.Members[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GuildFactsMessage) Deserialize(r Reader) error {

	typeinfosID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	linfos, err := GetType(typeinfosID)
	if err != nil {
		return err
	}

	linfos.Deserialize(r)

	m.Infos = linfos.(GuildFactSheetInformationsIntrf)

	lcreationDate, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.CreationDate = lcreationDate

	lnbTaxCollectors, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.NbTaxCollectors = lnbTaxCollectors

	lmembersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Members = make([]CharacterMinimalInformations, lmembersLen)

	for i := range m.Members {

		var lmembers CharacterMinimalInformations

		lmembers.Deserialize(r)

		m.Members[i] = lmembers

	}

	return nil
}

type GameContextRemoveMultipleElementsWithEventsMessage struct {
	GameContextRemoveMultipleElementsMessage

	ElementEventIds []uint8
}

func (m *GameContextRemoveMultipleElementsWithEventsMessage) ID() uint16 {
	return 6416
}

func (m *GameContextRemoveMultipleElementsWithEventsMessage) Serialize(w Writer) error {

	if err := m.GameContextRemoveMultipleElementsMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.ElementEventIds))); err != nil {
		return err
	}

	for i := range m.ElementEventIds {

		if err := w.WriteUInt8(m.ElementEventIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameContextRemoveMultipleElementsWithEventsMessage) Deserialize(r Reader) error {

	if err := m.GameContextRemoveMultipleElementsMessage.Deserialize(r); err != nil {
		return err
	}

	lelementEventIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ElementEventIds = make([]uint8, lelementEventIdsLen)

	for i := range m.ElementEventIds {

		lelementEventIds, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.ElementEventIds[i] = lelementEventIds

	}

	return nil
}

type AllianceInsiderInfoRequestMessage struct {
}

func (m *AllianceInsiderInfoRequestMessage) ID() uint16 {
	return 6417
}

func (m *AllianceInsiderInfoRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *AllianceInsiderInfoRequestMessage) Deserialize(r Reader) error {

	return nil
}

type GuildInAllianceFactsMessage struct {
	GuildFactsMessage

	AllianceInfos BasicNamedAllianceInformations
}

func (m *GuildInAllianceFactsMessage) ID() uint16 {
	return 6422
}

func (m *GuildInAllianceFactsMessage) Serialize(w Writer) error {

	if err := m.GuildFactsMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.AllianceInfos.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildInAllianceFactsMessage) Deserialize(r Reader) error {

	if err := m.GuildFactsMessage.Deserialize(r); err != nil {
		return err
	}

	var lallianceInfos BasicNamedAllianceInformations

	lallianceInfos.Deserialize(r)

	m.AllianceInfos = lallianceInfos

	return nil
}

type AllianceFactsErrorMessage struct {
	AllianceId uint32
}

func (m *AllianceFactsErrorMessage) ID() uint16 {
	return 6423
}

func (m *AllianceFactsErrorMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.AllianceId); err != nil {
		return err
	}

	return nil
}

func (m *AllianceFactsErrorMessage) Deserialize(r Reader) error {

	lallianceId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.AllianceId = lallianceId

	return nil
}

type GuildFactsErrorMessage struct {
	GuildId uint32
}

func (m *GuildFactsErrorMessage) ID() uint16 {
	return 6424
}

func (m *GuildFactsErrorMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.GuildId); err != nil {
		return err
	}

	return nil
}

func (m *GuildFactsErrorMessage) Deserialize(r Reader) error {

	lguildId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.GuildId = lguildId

	return nil
}

type GameRolePlayDelayedObjectUseMessage struct {
	GameRolePlayDelayedActionMessage

	ObjectGID uint16
}

func (m *GameRolePlayDelayedObjectUseMessage) ID() uint16 {
	return 6425
}

func (m *GameRolePlayDelayedObjectUseMessage) Serialize(w Writer) error {

	if err := m.GameRolePlayDelayedActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ObjectGID); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayDelayedObjectUseMessage) Deserialize(r Reader) error {

	if err := m.GameRolePlayDelayedActionMessage.Deserialize(r); err != nil {
		return err
	}

	lobjectGID, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ObjectGID = lobjectGID

	return nil
}

type AllianceChangeGuildRightsMessage struct {
	GuildId uint32

	Rights uint8
}

func (m *AllianceChangeGuildRightsMessage) ID() uint16 {
	return 6426
}

func (m *AllianceChangeGuildRightsMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.GuildId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Rights); err != nil {
		return err
	}

	return nil
}

func (m *AllianceChangeGuildRightsMessage) Deserialize(r Reader) error {

	lguildId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.GuildId = lguildId

	lrights, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Rights = lrights

	return nil
}

type AlliancePartialListMessage struct {
	AllianceListMessage
}

func (m *AlliancePartialListMessage) ID() uint16 {
	return 6427
}

func (m *AlliancePartialListMessage) Serialize(w Writer) error {

	if err := m.AllianceListMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AlliancePartialListMessage) Deserialize(r Reader) error {

	if err := m.AllianceListMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type ServerSessionConstantsMessage struct {
	Variables []ServerSessionConstantIntrf
}

func (m *ServerSessionConstantsMessage) ID() uint16 {
	return 6434
}

func (m *ServerSessionConstantsMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Variables))); err != nil {
		return err
	}

	for i := range m.Variables {

		if err := w.WriteUInt16(m.Variables[i].ID()); err != nil {
			return err
		}

		if err := m.Variables[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ServerSessionConstantsMessage) Deserialize(r Reader) error {

	lvariablesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Variables = make([]ServerSessionConstantIntrf, lvariablesLen)

	for i := range m.Variables {

		typevariablesID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lvariables, err := GetType(typevariablesID)
		if err != nil {
			return err
		}

		lvariables.Deserialize(r)

		m.Variables[i] = lvariables.(ServerSessionConstantIntrf)

	}

	return nil
}

type GuildVersatileInfoListMessage struct {
	Guilds []GuildVersatileInformationsIntrf
}

func (m *GuildVersatileInfoListMessage) ID() uint16 {
	return 6435
}

func (m *GuildVersatileInfoListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Guilds))); err != nil {
		return err
	}

	for i := range m.Guilds {

		if err := w.WriteUInt16(m.Guilds[i].ID()); err != nil {
			return err
		}

		if err := m.Guilds[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GuildVersatileInfoListMessage) Deserialize(r Reader) error {

	lguildsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Guilds = make([]GuildVersatileInformationsIntrf, lguildsLen)

	for i := range m.Guilds {

		typeguildsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lguilds, err := GetType(typeguildsID)
		if err != nil {
			return err
		}

		lguilds.Deserialize(r)

		m.Guilds[i] = lguilds.(GuildVersatileInformationsIntrf)

	}

	return nil
}

type AllianceVersatileInfoListMessage struct {
	Alliances []AllianceVersatileInformations
}

func (m *AllianceVersatileInfoListMessage) ID() uint16 {
	return 6436
}

func (m *AllianceVersatileInfoListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Alliances))); err != nil {
		return err
	}

	for i := range m.Alliances {

		if err := m.Alliances[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *AllianceVersatileInfoListMessage) Deserialize(r Reader) error {

	lalliancesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Alliances = make([]AllianceVersatileInformations, lalliancesLen)

	for i := range m.Alliances {

		var lalliances AllianceVersatileInformations

		lalliances.Deserialize(r)

		m.Alliances[i] = lalliances

	}

	return nil
}

type PrismSettingsRequestMessage struct {
	SubAreaId uint16

	StartDefenseTime uint8
}

func (m *PrismSettingsRequestMessage) ID() uint16 {
	return 6437
}

func (m *PrismSettingsRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.StartDefenseTime); err != nil {
		return err
	}

	return nil
}

func (m *PrismSettingsRequestMessage) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lstartDefenseTime, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.StartDefenseTime = lstartDefenseTime

	return nil
}

type PrismsListUpdateMessage struct {
	PrismsListMessage
}

func (m *PrismsListUpdateMessage) ID() uint16 {
	return 6438
}

func (m *PrismsListUpdateMessage) Serialize(w Writer) error {

	if err := m.PrismsListMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PrismsListUpdateMessage) Deserialize(r Reader) error {

	if err := m.PrismsListMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type KohUpdateMessage struct {
	Alliances []AllianceInformations

	AllianceNbMembers []uint16

	AllianceRoundWeigth []uint32

	AllianceMatchScore []uint8

	AllianceMapWinner BasicAllianceInformations

	AllianceMapWinnerScore uint32

	AllianceMapMyAllianceScore uint32

	NextTickTime float64
}

func (m *KohUpdateMessage) ID() uint16 {
	return 6439
}

func (m *KohUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Alliances))); err != nil {
		return err
	}

	for i := range m.Alliances {

		if err := m.Alliances[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.AllianceNbMembers))); err != nil {
		return err
	}

	for i := range m.AllianceNbMembers {

		if err := w.WriteVarUInt16(m.AllianceNbMembers[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.AllianceRoundWeigth))); err != nil {
		return err
	}

	for i := range m.AllianceRoundWeigth {

		if err := w.WriteVarUInt32(m.AllianceRoundWeigth[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.AllianceMatchScore))); err != nil {
		return err
	}

	for i := range m.AllianceMatchScore {

		if err := w.WriteUInt8(m.AllianceMatchScore[i]); err != nil {
			return err
		}

	}

	if err := m.AllianceMapWinner.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.AllianceMapWinnerScore); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.AllianceMapMyAllianceScore); err != nil {
		return err
	}

	if err := w.WriteDouble(m.NextTickTime); err != nil {
		return err
	}

	return nil
}

func (m *KohUpdateMessage) Deserialize(r Reader) error {

	lalliancesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Alliances = make([]AllianceInformations, lalliancesLen)

	for i := range m.Alliances {

		var lalliances AllianceInformations

		lalliances.Deserialize(r)

		m.Alliances[i] = lalliances

	}

	lallianceNbMembersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AllianceNbMembers = make([]uint16, lallianceNbMembersLen)

	for i := range m.AllianceNbMembers {

		lallianceNbMembers, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.AllianceNbMembers[i] = lallianceNbMembers

	}

	lallianceRoundWeigthLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AllianceRoundWeigth = make([]uint32, lallianceRoundWeigthLen)

	for i := range m.AllianceRoundWeigth {

		lallianceRoundWeigth, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.AllianceRoundWeigth[i] = lallianceRoundWeigth

	}

	lallianceMatchScoreLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AllianceMatchScore = make([]uint8, lallianceMatchScoreLen)

	for i := range m.AllianceMatchScore {

		lallianceMatchScore, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.AllianceMatchScore[i] = lallianceMatchScore

	}

	var lallianceMapWinner BasicAllianceInformations

	lallianceMapWinner.Deserialize(r)

	m.AllianceMapWinner = lallianceMapWinner

	lallianceMapWinnerScore, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.AllianceMapWinnerScore = lallianceMapWinnerScore

	lallianceMapMyAllianceScore, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.AllianceMapMyAllianceScore = lallianceMapMyAllianceScore

	lnextTickTime, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.NextTickTime = lnextTickTime

	return nil
}

type PrismsListMessage struct {
	Prisms []PrismSubareaEmptyInfoIntrf
}

func (m *PrismsListMessage) ID() uint16 {
	return 6440
}

func (m *PrismsListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Prisms))); err != nil {
		return err
	}

	for i := range m.Prisms {

		if err := w.WriteUInt16(m.Prisms[i].ID()); err != nil {
			return err
		}

		if err := m.Prisms[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PrismsListMessage) Deserialize(r Reader) error {

	lprismsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Prisms = make([]PrismSubareaEmptyInfoIntrf, lprismsLen)

	for i := range m.Prisms {

		typeprismsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lprisms, err := GetType(typeprismsID)
		if err != nil {
			return err
		}

		lprisms.Deserialize(r)

		m.Prisms[i] = lprisms.(PrismSubareaEmptyInfoIntrf)

	}

	return nil
}

type PrismsListRegisterMessage struct {
	Listen uint8
}

func (m *PrismsListRegisterMessage) ID() uint16 {
	return 6441
}

func (m *PrismsListRegisterMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Listen); err != nil {
		return err
	}

	return nil
}

func (m *PrismsListRegisterMessage) Deserialize(r Reader) error {

	llisten, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Listen = llisten

	return nil
}

type PrismSettingsErrorMessage struct {
}

func (m *PrismSettingsErrorMessage) ID() uint16 {
	return 6442
}

func (m *PrismSettingsErrorMessage) Serialize(w Writer) error {

	return nil
}

func (m *PrismSettingsErrorMessage) Deserialize(r Reader) error {

	return nil
}

type SetEnableAVARequestMessage struct {
	Enable bool
}

func (m *SetEnableAVARequestMessage) ID() uint16 {
	return 6443
}

func (m *SetEnableAVARequestMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *SetEnableAVARequestMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type AllianceModificationStartedMessage struct {
	CanChangeName bool

	CanChangeTag bool

	CanChangeEmblem bool
}

func (m *AllianceModificationStartedMessage) ID() uint16 {
	return 6444
}

func (m *AllianceModificationStartedMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.CanChangeName)

	setWrappedFlag(bbw0, 1, m.CanChangeTag)

	setWrappedFlag(bbw0, 2, m.CanChangeEmblem)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	return nil
}

func (m *AllianceModificationStartedMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CanChangeName = getWrappedFlag(bbw0, 0)

	m.CanChangeTag = getWrappedFlag(bbw0, 1)

	m.CanChangeEmblem = getWrappedFlag(bbw0, 2)

	return nil
}

type AllianceTaxCollectorDialogQuestionExtendedMessage struct {
	TaxCollectorDialogQuestionExtendedMessage

	Alliance BasicNamedAllianceInformations
}

func (m *AllianceTaxCollectorDialogQuestionExtendedMessage) ID() uint16 {
	return 6445
}

func (m *AllianceTaxCollectorDialogQuestionExtendedMessage) Serialize(w Writer) error {

	if err := m.TaxCollectorDialogQuestionExtendedMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.Alliance.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AllianceTaxCollectorDialogQuestionExtendedMessage) Deserialize(r Reader) error {

	if err := m.TaxCollectorDialogQuestionExtendedMessage.Deserialize(r); err != nil {
		return err
	}

	var lalliance BasicNamedAllianceInformations

	lalliance.Deserialize(r)

	m.Alliance = lalliance

	return nil
}

type AllianceModificationEmblemValidMessage struct {
	Alliancemblem GuildEmblem
}

func (m *AllianceModificationEmblemValidMessage) ID() uint16 {
	return 6447
}

func (m *AllianceModificationEmblemValidMessage) Serialize(w Writer) error {

	if err := m.Alliancemblem.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AllianceModificationEmblemValidMessage) Deserialize(r Reader) error {

	var lAlliancemblem GuildEmblem

	lAlliancemblem.Deserialize(r)

	m.Alliancemblem = lAlliancemblem

	return nil
}

type AlliancePrismDialogQuestionMessage struct {
}

func (m *AlliancePrismDialogQuestionMessage) ID() uint16 {
	return 6448
}

func (m *AlliancePrismDialogQuestionMessage) Serialize(w Writer) error {

	return nil
}

func (m *AlliancePrismDialogQuestionMessage) Deserialize(r Reader) error {

	return nil
}

type AllianceModificationNameAndTagValidMessage struct {
	AllianceName string

	AllianceTag string
}

func (m *AllianceModificationNameAndTagValidMessage) ID() uint16 {
	return 6449
}

func (m *AllianceModificationNameAndTagValidMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.AllianceName); err != nil {
		return err
	}

	if err := w.WriteString(m.AllianceTag); err != nil {
		return err
	}

	return nil
}

func (m *AllianceModificationNameAndTagValidMessage) Deserialize(r Reader) error {

	lallianceName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.AllianceName = lallianceName

	lallianceTag, err := r.ReadString()
	if err != nil {
		return err
	}

	m.AllianceTag = lallianceTag

	return nil
}

type AllianceModificationValidMessage struct {
	AllianceName string

	AllianceTag string

	Alliancemblem GuildEmblem
}

func (m *AllianceModificationValidMessage) ID() uint16 {
	return 6450
}

func (m *AllianceModificationValidMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.AllianceName); err != nil {
		return err
	}

	if err := w.WriteString(m.AllianceTag); err != nil {
		return err
	}

	if err := m.Alliancemblem.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AllianceModificationValidMessage) Deserialize(r Reader) error {

	lallianceName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.AllianceName = lallianceName

	lallianceTag, err := r.ReadString()
	if err != nil {
		return err
	}

	m.AllianceTag = lallianceTag

	var lAlliancemblem GuildEmblem

	lAlliancemblem.Deserialize(r)

	m.Alliancemblem = lAlliancemblem

	return nil
}

type PrismsInfoValidMessage struct {
	Fights []PrismFightersInformation
}

func (m *PrismsInfoValidMessage) ID() uint16 {
	return 6451
}

func (m *PrismsInfoValidMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Fights))); err != nil {
		return err
	}

	for i := range m.Fights {

		if err := m.Fights[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *PrismsInfoValidMessage) Deserialize(r Reader) error {

	lfightsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Fights = make([]PrismFightersInformation, lfightsLen)

	for i := range m.Fights {

		var lfights PrismFightersInformation

		lfights.Deserialize(r)

		m.Fights[i] = lfights

	}

	return nil
}

type PrismFightAddedMessage struct {
	Fight PrismFightersInformation
}

func (m *PrismFightAddedMessage) ID() uint16 {
	return 6452
}

func (m *PrismFightAddedMessage) Serialize(w Writer) error {

	if err := m.Fight.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *PrismFightAddedMessage) Deserialize(r Reader) error {

	var lfight PrismFightersInformation

	lfight.Deserialize(r)

	m.Fight = lfight

	return nil
}

type PrismFightRemovedMessage struct {
	SubAreaId uint16
}

func (m *PrismFightRemovedMessage) ID() uint16 {
	return 6453
}

func (m *PrismFightRemovedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	return nil
}

func (m *PrismFightRemovedMessage) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	return nil
}

type UpdateMapPlayersAgressableStatusMessage struct {
	PlayerIds []int64

	Enable []uint8
}

func (m *UpdateMapPlayersAgressableStatusMessage) ID() uint16 {
	return 6454
}

func (m *UpdateMapPlayersAgressableStatusMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.PlayerIds))); err != nil {
		return err
	}

	for i := range m.PlayerIds {

		if err := w.WriteVarInt64(m.PlayerIds[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Enable))); err != nil {
		return err
	}

	for i := range m.Enable {

		if err := w.WriteUInt8(m.Enable[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *UpdateMapPlayersAgressableStatusMessage) Deserialize(r Reader) error {

	lplayerIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PlayerIds = make([]int64, lplayerIdsLen)

	for i := range m.PlayerIds {

		lplayerIds, err := r.ReadVarInt64()
		if err != nil {
			return err
		}

		m.PlayerIds[i] = lplayerIds

	}

	lenableLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Enable = make([]uint8, lenableLen)

	for i := range m.Enable {

		lenable, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.Enable[i] = lenable

	}

	return nil
}

type TaxCollectorStateUpdateMessage struct {
	UniqueId int32

	State int8
}

func (m *TaxCollectorStateUpdateMessage) ID() uint16 {
	return 6455
}

func (m *TaxCollectorStateUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.UniqueId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.State); err != nil {
		return err
	}

	return nil
}

func (m *TaxCollectorStateUpdateMessage) Deserialize(r Reader) error {

	luniqueId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.UniqueId = luniqueId

	lstate, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.State = lstate

	return nil
}

type UpdateSelfAgressableStatusMessage struct {
	Status uint8

	ProbationTime uint32
}

func (m *UpdateSelfAgressableStatusMessage) ID() uint16 {
	return 6456
}

func (m *UpdateSelfAgressableStatusMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Status); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.ProbationTime); err != nil {
		return err
	}

	return nil
}

func (m *UpdateSelfAgressableStatusMessage) Deserialize(r Reader) error {

	lstatus, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Status = lstatus

	lprobationTime, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.ProbationTime = lprobationTime

	return nil
}

type MimicryObjectEraseRequestMessage struct {
	HostUID uint32

	HostPos uint8
}

func (m *MimicryObjectEraseRequestMessage) ID() uint16 {
	return 6457
}

func (m *MimicryObjectEraseRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HostUID); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.HostPos); err != nil {
		return err
	}

	return nil
}

func (m *MimicryObjectEraseRequestMessage) Deserialize(r Reader) error {

	lhostUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HostUID = lhostUID

	lhostPos, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.HostPos = lhostPos

	return nil
}

type MimicryObjectPreviewMessage struct {
	Result ObjectItem
}

func (m *MimicryObjectPreviewMessage) ID() uint16 {
	return 6458
}

func (m *MimicryObjectPreviewMessage) Serialize(w Writer) error {

	if err := m.Result.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *MimicryObjectPreviewMessage) Deserialize(r Reader) error {

	var lresult ObjectItem

	lresult.Deserialize(r)

	m.Result = lresult

	return nil
}

type ClientUIOpenedMessage struct {
	Type uint8
}

func (m *ClientUIOpenedMessage) ID() uint16 {
	return 6459
}

func (m *ClientUIOpenedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	return nil
}

func (m *ClientUIOpenedMessage) Deserialize(r Reader) error {

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	return nil
}

type MimicryObjectFeedAndAssociateRequestMessage struct {
	SymbioticObjectAssociateRequestMessage

	FoodUID uint32

	FoodPos uint8

	Preview bool
}

func (m *MimicryObjectFeedAndAssociateRequestMessage) ID() uint16 {
	return 6460
}

func (m *MimicryObjectFeedAndAssociateRequestMessage) Serialize(w Writer) error {

	if err := m.SymbioticObjectAssociateRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.FoodUID); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.FoodPos); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Preview); err != nil {
		return err
	}

	return nil
}

func (m *MimicryObjectFeedAndAssociateRequestMessage) Deserialize(r Reader) error {

	if err := m.SymbioticObjectAssociateRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lfoodUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.FoodUID = lfoodUID

	lfoodPos, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.FoodPos = lfoodPos

	lpreview, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Preview = lpreview

	return nil
}

type MimicryObjectErrorMessage struct {
	SymbioticObjectErrorMessage

	Preview bool
}

func (m *MimicryObjectErrorMessage) ID() uint16 {
	return 6461
}

func (m *MimicryObjectErrorMessage) Serialize(w Writer) error {

	if err := m.SymbioticObjectErrorMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Preview); err != nil {
		return err
	}

	return nil
}

func (m *MimicryObjectErrorMessage) Deserialize(r Reader) error {

	if err := m.SymbioticObjectErrorMessage.Deserialize(r); err != nil {
		return err
	}

	lpreview, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Preview = lpreview

	return nil
}

type MimicryObjectAssociatedMessage struct {
	SymbioticObjectAssociatedMessage
}

func (m *MimicryObjectAssociatedMessage) ID() uint16 {
	return 6462
}

func (m *MimicryObjectAssociatedMessage) Serialize(w Writer) error {

	if err := m.SymbioticObjectAssociatedMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *MimicryObjectAssociatedMessage) Deserialize(r Reader) error {

	if err := m.SymbioticObjectAssociatedMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type ClientUIOpenedByObjectMessage struct {
	ClientUIOpenedMessage

	Uid uint32
}

func (m *ClientUIOpenedByObjectMessage) ID() uint16 {
	return 6463
}

func (m *ClientUIOpenedByObjectMessage) Serialize(w Writer) error {

	if err := m.ClientUIOpenedMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.Uid); err != nil {
		return err
	}

	return nil
}

func (m *ClientUIOpenedByObjectMessage) Deserialize(r Reader) error {

	if err := m.ClientUIOpenedMessage.Deserialize(r); err != nil {
		return err
	}

	luid, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.Uid = luid

	return nil
}

type ExchangeBidPriceForSellerMessage struct {
	ExchangeBidPriceMessage

	AllIdentical bool

	MinimalPrices []int64
}

func (m *ExchangeBidPriceForSellerMessage) ID() uint16 {
	return 6464
}

func (m *ExchangeBidPriceForSellerMessage) Serialize(w Writer) error {

	if err := m.ExchangeBidPriceMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.AllIdentical); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.MinimalPrices))); err != nil {
		return err
	}

	for i := range m.MinimalPrices {

		if err := w.WriteVarInt64(m.MinimalPrices[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeBidPriceForSellerMessage) Deserialize(r Reader) error {

	if err := m.ExchangeBidPriceMessage.Deserialize(r); err != nil {
		return err
	}

	lallIdentical, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.AllIdentical = lallIdentical

	lminimalPricesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MinimalPrices = make([]int64, lminimalPricesLen)

	for i := range m.MinimalPrices {

		lminimalPrices, err := r.ReadVarInt64()
		if err != nil {
			return err
		}

		m.MinimalPrices[i] = lminimalPrices

	}

	return nil
}

type GameFightTurnStartPlayingMessage struct {
}

func (m *GameFightTurnStartPlayingMessage) ID() uint16 {
	return 6465
}

func (m *GameFightTurnStartPlayingMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameFightTurnStartPlayingMessage) Deserialize(r Reader) error {

	return nil
}

type PrismSetSabotagedRefusedMessage struct {
	SubAreaId uint16

	Reason int8
}

func (m *PrismSetSabotagedRefusedMessage) ID() uint16 {
	return 6466
}

func (m *PrismSetSabotagedRefusedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *PrismSetSabotagedRefusedMessage) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	lreason, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type PrismSetSabotagedRequestMessage struct {
	SubAreaId uint16
}

func (m *PrismSetSabotagedRequestMessage) ID() uint16 {
	return 6468
}

func (m *PrismSetSabotagedRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SubAreaId); err != nil {
		return err
	}

	return nil
}

func (m *PrismSetSabotagedRequestMessage) Deserialize(r Reader) error {

	lsubAreaId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SubAreaId = lsubAreaId

	return nil
}

type SelectedServerDataExtendedMessage struct {
	SelectedServerDataMessage

	ServerIds []uint16
}

func (m *SelectedServerDataExtendedMessage) ID() uint16 {
	return 6469
}

func (m *SelectedServerDataExtendedMessage) Serialize(w Writer) error {

	if err := m.SelectedServerDataMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.ServerIds))); err != nil {
		return err
	}

	for i := range m.ServerIds {

		if err := w.WriteVarUInt16(m.ServerIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *SelectedServerDataExtendedMessage) Deserialize(r Reader) error {

	if err := m.SelectedServerDataMessage.Deserialize(r); err != nil {
		return err
	}

	lserverIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ServerIds = make([]uint16, lserverIdsLen)

	for i := range m.ServerIds {

		lserverIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.ServerIds[i] = lserverIds

	}

	return nil
}

type ExchangeObjectTransfertListWithQuantityToInvMessage struct {
	Ids []uint32

	Qtys []uint32
}

func (m *ExchangeObjectTransfertListWithQuantityToInvMessage) ID() uint16 {
	return 6470
}

func (m *ExchangeObjectTransfertListWithQuantityToInvMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Ids))); err != nil {
		return err
	}

	for i := range m.Ids {

		if err := w.WriteVarUInt32(m.Ids[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Qtys))); err != nil {
		return err
	}

	for i := range m.Qtys {

		if err := w.WriteVarUInt32(m.Qtys[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeObjectTransfertListWithQuantityToInvMessage) Deserialize(r Reader) error {

	lidsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Ids = make([]uint32, lidsLen)

	for i := range m.Ids {

		lids, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.Ids[i] = lids

	}

	lqtysLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Qtys = make([]uint32, lqtysLen)

	for i := range m.Qtys {

		lqtys, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.Qtys[i] = lqtys

	}

	return nil
}

type CharacterLoadingCompleteMessage struct {
}

func (m *CharacterLoadingCompleteMessage) ID() uint16 {
	return 6471
}

func (m *CharacterLoadingCompleteMessage) Serialize(w Writer) error {

	return nil
}

func (m *CharacterLoadingCompleteMessage) Deserialize(r Reader) error {

	return nil
}

type PartyCompanionUpdateLightMessage struct {
	PartyUpdateLightMessage

	IndexId uint8
}

func (m *PartyCompanionUpdateLightMessage) ID() uint16 {
	return 6472
}

func (m *PartyCompanionUpdateLightMessage) Serialize(w Writer) error {

	if err := m.PartyUpdateLightMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.IndexId); err != nil {
		return err
	}

	return nil
}

func (m *PartyCompanionUpdateLightMessage) Deserialize(r Reader) error {

	if err := m.PartyUpdateLightMessage.Deserialize(r); err != nil {
		return err
	}

	lindexId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.IndexId = lindexId

	return nil
}

type GameFightSpectatePlayerRequestMessage struct {
	PlayerId int64
}

func (m *GameFightSpectatePlayerRequestMessage) ID() uint16 {
	return 6474
}

func (m *GameFightSpectatePlayerRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightSpectatePlayerRequestMessage) Deserialize(r Reader) error {

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type BasicCharactersListMessage struct {
	Characters []CharacterBaseInformationsIntrf
}

func (m *BasicCharactersListMessage) ID() uint16 {
	return 6475
}

func (m *BasicCharactersListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Characters))); err != nil {
		return err
	}

	for i := range m.Characters {

		if err := w.WriteUInt16(m.Characters[i].ID()); err != nil {
			return err
		}

		if err := m.Characters[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *BasicCharactersListMessage) Deserialize(r Reader) error {

	lcharactersLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Characters = make([]CharacterBaseInformationsIntrf, lcharactersLen)

	for i := range m.Characters {

		typecharactersID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lcharacters, err := GetType(typecharactersID)
		if err != nil {
			return err
		}

		lcharacters.Deserialize(r)

		m.Characters[i] = lcharacters.(CharacterBaseInformationsIntrf)

	}

	return nil
}

type TreasureHuntFinishedMessage struct {
	QuestType uint8
}

func (m *TreasureHuntFinishedMessage) ID() uint16 {
	return 6483
}

func (m *TreasureHuntFinishedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.QuestType); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntFinishedMessage) Deserialize(r Reader) error {

	lquestType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestType = lquestType

	return nil
}

type TreasureHuntDigRequestAnswerMessage struct {
	QuestType uint8

	Result uint8
}

func (m *TreasureHuntDigRequestAnswerMessage) ID() uint16 {
	return 6484
}

func (m *TreasureHuntDigRequestAnswerMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.QuestType); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Result); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntDigRequestAnswerMessage) Deserialize(r Reader) error {

	lquestType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestType = lquestType

	lresult, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Result = lresult

	return nil
}

type TreasureHuntDigRequestMessage struct {
	QuestType uint8
}

func (m *TreasureHuntDigRequestMessage) ID() uint16 {
	return 6485
}

func (m *TreasureHuntDigRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.QuestType); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntDigRequestMessage) Deserialize(r Reader) error {

	lquestType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestType = lquestType

	return nil
}

type TreasureHuntMessage struct {
	QuestType uint8

	StartMapId int32

	KnownStepsList []TreasureHuntStepIntrf

	TotalStepCount uint8

	CheckPointCurrent uint32

	CheckPointTotal uint32

	AvailableRetryCount int32

	Flags []TreasureHuntFlag
}

func (m *TreasureHuntMessage) ID() uint16 {
	return 6486
}

func (m *TreasureHuntMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.QuestType); err != nil {
		return err
	}

	if err := w.WriteInt32(m.StartMapId); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.KnownStepsList))); err != nil {
		return err
	}

	for i := range m.KnownStepsList {

		if err := w.WriteUInt16(m.KnownStepsList[i].ID()); err != nil {
			return err
		}

		if err := m.KnownStepsList[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteUInt8(m.TotalStepCount); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.CheckPointCurrent); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.CheckPointTotal); err != nil {
		return err
	}

	if err := w.WriteInt32(m.AvailableRetryCount); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Flags))); err != nil {
		return err
	}

	for i := range m.Flags {

		if err := m.Flags[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *TreasureHuntMessage) Deserialize(r Reader) error {

	lquestType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestType = lquestType

	lstartMapId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.StartMapId = lstartMapId

	lknownStepsListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.KnownStepsList = make([]TreasureHuntStepIntrf, lknownStepsListLen)

	for i := range m.KnownStepsList {

		typeknownStepsListID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lknownStepsList, err := GetType(typeknownStepsListID)
		if err != nil {
			return err
		}

		lknownStepsList.Deserialize(r)

		m.KnownStepsList[i] = lknownStepsList.(TreasureHuntStepIntrf)

	}

	ltotalStepCount, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TotalStepCount = ltotalStepCount

	lcheckPointCurrent, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.CheckPointCurrent = lcheckPointCurrent

	lcheckPointTotal, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.CheckPointTotal = lcheckPointTotal

	lavailableRetryCount, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.AvailableRetryCount = lavailableRetryCount

	lflagsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Flags = make([]TreasureHuntFlag, lflagsLen)

	for i := range m.Flags {

		var lflags TreasureHuntFlag

		lflags.Deserialize(r)

		m.Flags[i] = lflags

	}

	return nil
}

type TreasureHuntGiveUpRequestMessage struct {
	QuestType uint8
}

func (m *TreasureHuntGiveUpRequestMessage) ID() uint16 {
	return 6487
}

func (m *TreasureHuntGiveUpRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.QuestType); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntGiveUpRequestMessage) Deserialize(r Reader) error {

	lquestType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestType = lquestType

	return nil
}

type TreasureHuntRequestMessage struct {
	QuestLevel uint8

	QuestType uint8
}

func (m *TreasureHuntRequestMessage) ID() uint16 {
	return 6488
}

func (m *TreasureHuntRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.QuestLevel); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.QuestType); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntRequestMessage) Deserialize(r Reader) error {

	lquestLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestLevel = lquestLevel

	lquestType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestType = lquestType

	return nil
}

type TreasureHuntRequestAnswerMessage struct {
	QuestType uint8

	Result uint8
}

func (m *TreasureHuntRequestAnswerMessage) ID() uint16 {
	return 6489
}

func (m *TreasureHuntRequestAnswerMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.QuestType); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Result); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntRequestAnswerMessage) Deserialize(r Reader) error {

	lquestType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestType = lquestType

	lresult, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Result = lresult

	return nil
}

type GameFightNewWaveMessage struct {
	Id uint8

	TeamId uint8

	NbTurnBeforeNextWave int16
}

func (m *GameFightNewWaveMessage) ID() uint16 {
	return 6490
}

func (m *GameFightNewWaveMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Id); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.TeamId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.NbTurnBeforeNextWave); err != nil {
		return err
	}

	return nil
}

func (m *GameFightNewWaveMessage) Deserialize(r Reader) error {

	lid, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Id = lid

	lteamId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.TeamId = lteamId

	lnbTurnBeforeNextWave, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.NbTurnBeforeNextWave = lnbTurnBeforeNextWave

	return nil
}

type TreasureHuntAvailableRetryCountUpdateMessage struct {
	QuestType uint8

	AvailableRetryCount int32
}

func (m *TreasureHuntAvailableRetryCountUpdateMessage) ID() uint16 {
	return 6491
}

func (m *TreasureHuntAvailableRetryCountUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.QuestType); err != nil {
		return err
	}

	if err := w.WriteInt32(m.AvailableRetryCount); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntAvailableRetryCountUpdateMessage) Deserialize(r Reader) error {

	lquestType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestType = lquestType

	lavailableRetryCount, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.AvailableRetryCount = lavailableRetryCount

	return nil
}

type PortalUseRequestMessage struct {
	PortalId uint32
}

func (m *PortalUseRequestMessage) ID() uint16 {
	return 6492
}

func (m *PortalUseRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.PortalId); err != nil {
		return err
	}

	return nil
}

func (m *PortalUseRequestMessage) Deserialize(r Reader) error {

	lportalId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.PortalId = lportalId

	return nil
}

type AreaFightModificatorUpdateMessage struct {
	SpellPairId int32
}

func (m *AreaFightModificatorUpdateMessage) ID() uint16 {
	return 6493
}

func (m *AreaFightModificatorUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.SpellPairId); err != nil {
		return err
	}

	return nil
}

func (m *AreaFightModificatorUpdateMessage) Deserialize(r Reader) error {

	lspellPairId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.SpellPairId = lspellPairId

	return nil
}

type GameCautiousMapMovementRequestMessage struct {
	GameMapMovementRequestMessage
}

func (m *GameCautiousMapMovementRequestMessage) ID() uint16 {
	return 6496
}

func (m *GameCautiousMapMovementRequestMessage) Serialize(w Writer) error {

	if err := m.GameMapMovementRequestMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameCautiousMapMovementRequestMessage) Deserialize(r Reader) error {

	if err := m.GameMapMovementRequestMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type GameCautiousMapMovementMessage struct {
	GameMapMovementMessage
}

func (m *GameCautiousMapMovementMessage) ID() uint16 {
	return 6497
}

func (m *GameCautiousMapMovementMessage) Serialize(w Writer) error {

	if err := m.GameMapMovementMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameCautiousMapMovementMessage) Deserialize(r Reader) error {

	if err := m.GameMapMovementMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type TreasureHuntShowLegendaryUIMessage struct {
	AvailableLegendaryIds []uint16
}

func (m *TreasureHuntShowLegendaryUIMessage) ID() uint16 {
	return 6498
}

func (m *TreasureHuntShowLegendaryUIMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.AvailableLegendaryIds))); err != nil {
		return err
	}

	for i := range m.AvailableLegendaryIds {

		if err := w.WriteVarUInt16(m.AvailableLegendaryIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *TreasureHuntShowLegendaryUIMessage) Deserialize(r Reader) error {

	lavailableLegendaryIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.AvailableLegendaryIds = make([]uint16, lavailableLegendaryIdsLen)

	for i := range m.AvailableLegendaryIds {

		lavailableLegendaryIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.AvailableLegendaryIds[i] = lavailableLegendaryIds

	}

	return nil
}

type TreasureHuntLegendaryRequestMessage struct {
	LegendaryId uint16
}

func (m *TreasureHuntLegendaryRequestMessage) ID() uint16 {
	return 6499
}

func (m *TreasureHuntLegendaryRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.LegendaryId); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntLegendaryRequestMessage) Deserialize(r Reader) error {

	llegendaryId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.LegendaryId = llegendaryId

	return nil
}

type MapRunningFightDetailsExtendedMessage struct {
	MapRunningFightDetailsMessage

	NamedPartyTeams []NamedPartyTeam
}

func (m *MapRunningFightDetailsExtendedMessage) ID() uint16 {
	return 6500
}

func (m *MapRunningFightDetailsExtendedMessage) Serialize(w Writer) error {

	if err := m.MapRunningFightDetailsMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.NamedPartyTeams))); err != nil {
		return err
	}

	for i := range m.NamedPartyTeams {

		if err := m.NamedPartyTeams[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *MapRunningFightDetailsExtendedMessage) Deserialize(r Reader) error {

	if err := m.MapRunningFightDetailsMessage.Deserialize(r); err != nil {
		return err
	}

	lnamedPartyTeamsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.NamedPartyTeams = make([]NamedPartyTeam, lnamedPartyTeamsLen)

	for i := range m.NamedPartyTeams {

		var lnamedPartyTeams NamedPartyTeam

		lnamedPartyTeams.Deserialize(r)

		m.NamedPartyTeams[i] = lnamedPartyTeams

	}

	return nil
}

type PartyNameSetErrorMessage struct {
	AbstractPartyMessage

	Result uint8
}

func (m *PartyNameSetErrorMessage) ID() uint16 {
	return 6501
}

func (m *PartyNameSetErrorMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Result); err != nil {
		return err
	}

	return nil
}

func (m *PartyNameSetErrorMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lresult, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Result = lresult

	return nil
}

type PartyNameUpdateMessage struct {
	AbstractPartyMessage

	PartyName string
}

func (m *PartyNameUpdateMessage) ID() uint16 {
	return 6502
}

func (m *PartyNameUpdateMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.PartyName); err != nil {
		return err
	}

	return nil
}

func (m *PartyNameUpdateMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lpartyName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PartyName = lpartyName

	return nil
}

type PartyNameSetRequestMessage struct {
	AbstractPartyMessage

	PartyName string
}

func (m *PartyNameSetRequestMessage) ID() uint16 {
	return 6503
}

func (m *PartyNameSetRequestMessage) Serialize(w Writer) error {

	if err := m.AbstractPartyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.PartyName); err != nil {
		return err
	}

	return nil
}

func (m *PartyNameSetRequestMessage) Deserialize(r Reader) error {

	if err := m.AbstractPartyMessage.Deserialize(r); err != nil {
		return err
	}

	lpartyName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.PartyName = lpartyName

	return nil
}

type GameFightSpectatorJoinMessage struct {
	GameFightJoinMessage

	NamedPartyTeams []NamedPartyTeam
}

func (m *GameFightSpectatorJoinMessage) ID() uint16 {
	return 6504
}

func (m *GameFightSpectatorJoinMessage) Serialize(w Writer) error {

	if err := m.GameFightJoinMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.NamedPartyTeams))); err != nil {
		return err
	}

	for i := range m.NamedPartyTeams {

		if err := m.NamedPartyTeams[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameFightSpectatorJoinMessage) Deserialize(r Reader) error {

	if err := m.GameFightJoinMessage.Deserialize(r); err != nil {
		return err
	}

	lnamedPartyTeamsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.NamedPartyTeams = make([]NamedPartyTeam, lnamedPartyTeamsLen)

	for i := range m.NamedPartyTeams {

		var lnamedPartyTeams NamedPartyTeam

		lnamedPartyTeams.Deserialize(r)

		m.NamedPartyTeams[i] = lnamedPartyTeams

	}

	return nil
}

type GuestModeMessage struct {
	Active bool
}

func (m *GuestModeMessage) ID() uint16 {
	return 6505
}

func (m *GuestModeMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Active); err != nil {
		return err
	}

	return nil
}

func (m *GuestModeMessage) Deserialize(r Reader) error {

	lactive, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Active = lactive

	return nil
}

type GuestLimitationMessage struct {
	Reason uint8
}

func (m *GuestLimitationMessage) ID() uint16 {
	return 6506
}

func (m *GuestLimitationMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *GuestLimitationMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type TreasureHuntFlagRequestAnswerMessage struct {
	QuestType uint8

	Result uint8

	Index uint8
}

func (m *TreasureHuntFlagRequestAnswerMessage) ID() uint16 {
	return 6507
}

func (m *TreasureHuntFlagRequestAnswerMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.QuestType); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Result); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Index); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntFlagRequestAnswerMessage) Deserialize(r Reader) error {

	lquestType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestType = lquestType

	lresult, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Result = lresult

	lindex, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Index = lindex

	return nil
}

type TreasureHuntFlagRequestMessage struct {
	QuestType uint8

	Index uint8
}

func (m *TreasureHuntFlagRequestMessage) ID() uint16 {
	return 6508
}

func (m *TreasureHuntFlagRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.QuestType); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Index); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntFlagRequestMessage) Deserialize(r Reader) error {

	lquestType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestType = lquestType

	lindex, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Index = lindex

	return nil
}

type TreasureHuntDigRequestAnswerFailedMessage struct {
	TreasureHuntDigRequestAnswerMessage

	WrongFlagCount uint8
}

func (m *TreasureHuntDigRequestAnswerFailedMessage) ID() uint16 {
	return 6509
}

func (m *TreasureHuntDigRequestAnswerFailedMessage) Serialize(w Writer) error {

	if err := m.TreasureHuntDigRequestAnswerMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.WrongFlagCount); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntDigRequestAnswerFailedMessage) Deserialize(r Reader) error {

	if err := m.TreasureHuntDigRequestAnswerMessage.Deserialize(r); err != nil {
		return err
	}

	lwrongFlagCount, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.WrongFlagCount = lwrongFlagCount

	return nil
}

type TreasureHuntFlagRemoveRequestMessage struct {
	QuestType uint8

	Index uint8
}

func (m *TreasureHuntFlagRemoveRequestMessage) ID() uint16 {
	return 6510
}

func (m *TreasureHuntFlagRemoveRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.QuestType); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Index); err != nil {
		return err
	}

	return nil
}

func (m *TreasureHuntFlagRemoveRequestMessage) Deserialize(r Reader) error {

	lquestType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.QuestType = lquestType

	lindex, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Index = lindex

	return nil
}

type WarnOnPermaDeathMessage struct {
	Enable bool
}

func (m *WarnOnPermaDeathMessage) ID() uint16 {
	return 6512
}

func (m *WarnOnPermaDeathMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *WarnOnPermaDeathMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type WarnOnPermaDeathStateMessage struct {
	Enable bool
}

func (m *WarnOnPermaDeathStateMessage) ID() uint16 {
	return 6513
}

func (m *WarnOnPermaDeathStateMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Enable); err != nil {
		return err
	}

	return nil
}

func (m *WarnOnPermaDeathStateMessage) Deserialize(r Reader) error {

	lenable, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Enable = lenable

	return nil
}

type PaddockBuyResultMessage struct {
	PaddockId int32

	Bought bool

	RealPrice int64
}

func (m *PaddockBuyResultMessage) ID() uint16 {
	return 6516
}

func (m *PaddockBuyResultMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.PaddockId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Bought); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.RealPrice); err != nil {
		return err
	}

	return nil
}

func (m *PaddockBuyResultMessage) Deserialize(r Reader) error {

	lpaddockId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.PaddockId = lpaddockId

	lbought, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Bought = lbought

	lrealPrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.RealPrice = lrealPrice

	return nil
}

type AccessoryPreviewMessage struct {
	Look EntityLook
}

func (m *AccessoryPreviewMessage) ID() uint16 {
	return 6517
}

func (m *AccessoryPreviewMessage) Serialize(w Writer) error {

	if err := m.Look.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AccessoryPreviewMessage) Deserialize(r Reader) error {

	var llook EntityLook

	llook.Deserialize(r)

	m.Look = llook

	return nil
}

type AccessoryPreviewRequestMessage struct {
	GenericId []uint16
}

func (m *AccessoryPreviewRequestMessage) ID() uint16 {
	return 6518
}

func (m *AccessoryPreviewRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.GenericId))); err != nil {
		return err
	}

	for i := range m.GenericId {

		if err := w.WriteVarUInt16(m.GenericId[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *AccessoryPreviewRequestMessage) Deserialize(r Reader) error {

	lgenericIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.GenericId = make([]uint16, lgenericIdLen)

	for i := range m.GenericId {

		lgenericId, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.GenericId[i] = lgenericId

	}

	return nil
}

type ObtainedItemMessage struct {
	GenericId uint16

	BaseQuantity uint32
}

func (m *ObtainedItemMessage) ID() uint16 {
	return 6519
}

func (m *ObtainedItemMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.GenericId); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.BaseQuantity); err != nil {
		return err
	}

	return nil
}

func (m *ObtainedItemMessage) Deserialize(r Reader) error {

	lgenericId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.GenericId = lgenericId

	lbaseQuantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.BaseQuantity = lbaseQuantity

	return nil
}

type ObtainedItemWithBonusMessage struct {
	ObtainedItemMessage

	BonusQuantity uint32
}

func (m *ObtainedItemWithBonusMessage) ID() uint16 {
	return 6520
}

func (m *ObtainedItemWithBonusMessage) Serialize(w Writer) error {

	if err := m.ObtainedItemMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.BonusQuantity); err != nil {
		return err
	}

	return nil
}

func (m *ObtainedItemWithBonusMessage) Deserialize(r Reader) error {

	if err := m.ObtainedItemMessage.Deserialize(r); err != nil {
		return err
	}

	lbonusQuantity, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.BonusQuantity = lbonusQuantity

	return nil
}

type AccessoryPreviewErrorMessage struct {
	Error uint8
}

func (m *AccessoryPreviewErrorMessage) ID() uint16 {
	return 6521
}

func (m *AccessoryPreviewErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Error); err != nil {
		return err
	}

	return nil
}

func (m *AccessoryPreviewErrorMessage) Deserialize(r Reader) error {

	lerror, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Error = lerror

	return nil
}

type SymbioticObjectAssociateRequestMessage struct {
	SymbioteUID uint32

	SymbiotePos uint8

	HostUID uint32

	HostPos uint8
}

func (m *SymbioticObjectAssociateRequestMessage) ID() uint16 {
	return 6522
}

func (m *SymbioticObjectAssociateRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.SymbioteUID); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.SymbiotePos); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.HostUID); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.HostPos); err != nil {
		return err
	}

	return nil
}

func (m *SymbioticObjectAssociateRequestMessage) Deserialize(r Reader) error {

	lsymbioteUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.SymbioteUID = lsymbioteUID

	lsymbiotePos, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.SymbiotePos = lsymbiotePos

	lhostUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HostUID = lhostUID

	lhostPos, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.HostPos = lhostPos

	return nil
}

type WrapperObjectAssociatedMessage struct {
	SymbioticObjectAssociatedMessage
}

func (m *WrapperObjectAssociatedMessage) ID() uint16 {
	return 6523
}

func (m *WrapperObjectAssociatedMessage) Serialize(w Writer) error {

	if err := m.SymbioticObjectAssociatedMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *WrapperObjectAssociatedMessage) Deserialize(r Reader) error {

	if err := m.SymbioticObjectAssociatedMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type WrapperObjectDissociateRequestMessage struct {
	HostUID uint32

	HostPos uint8
}

func (m *WrapperObjectDissociateRequestMessage) ID() uint16 {
	return 6524
}

func (m *WrapperObjectDissociateRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HostUID); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.HostPos); err != nil {
		return err
	}

	return nil
}

func (m *WrapperObjectDissociateRequestMessage) Deserialize(r Reader) error {

	lhostUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HostUID = lhostUID

	lhostPos, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.HostPos = lhostPos

	return nil
}

type CurrentServerStatusUpdateMessage struct {
	Status uint8
}

func (m *CurrentServerStatusUpdateMessage) ID() uint16 {
	return 6525
}

func (m *CurrentServerStatusUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CurrentServerStatusUpdateMessage) Deserialize(r Reader) error {

	lstatus, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Status = lstatus

	return nil
}

type SymbioticObjectErrorMessage struct {
	ObjectErrorMessage

	ErrorCode int8
}

func (m *SymbioticObjectErrorMessage) ID() uint16 {
	return 6526
}

func (m *SymbioticObjectErrorMessage) Serialize(w Writer) error {

	if err := m.ObjectErrorMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt8(m.ErrorCode); err != nil {
		return err
	}

	return nil
}

func (m *SymbioticObjectErrorMessage) Deserialize(r Reader) error {

	if err := m.ObjectErrorMessage.Deserialize(r); err != nil {
		return err
	}

	lerrorCode, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.ErrorCode = lerrorCode

	return nil
}

type SymbioticObjectAssociatedMessage struct {
	HostUID uint32
}

func (m *SymbioticObjectAssociatedMessage) ID() uint16 {
	return 6527
}

func (m *SymbioticObjectAssociatedMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HostUID); err != nil {
		return err
	}

	return nil
}

func (m *SymbioticObjectAssociatedMessage) Deserialize(r Reader) error {

	lhostUID, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HostUID = lhostUID

	return nil
}

type WrapperObjectErrorMessage struct {
	SymbioticObjectErrorMessage
}

func (m *WrapperObjectErrorMessage) ID() uint16 {
	return 6529
}

func (m *WrapperObjectErrorMessage) Serialize(w Writer) error {

	if err := m.SymbioticObjectErrorMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *WrapperObjectErrorMessage) Deserialize(r Reader) error {

	if err := m.SymbioticObjectErrorMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type BasicStatMessage struct {
	TimeSpent float64

	StatId uint16
}

func (m *BasicStatMessage) ID() uint16 {
	return 6530
}

func (m *BasicStatMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.TimeSpent); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.StatId); err != nil {
		return err
	}

	return nil
}

func (m *BasicStatMessage) Deserialize(r Reader) error {

	ltimeSpent, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.TimeSpent = ltimeSpent

	lstatId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.StatId = lstatId

	return nil
}

type PrismModuleExchangeRequestMessage struct {
}

func (m *PrismModuleExchangeRequestMessage) ID() uint16 {
	return 6531
}

func (m *PrismModuleExchangeRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *PrismModuleExchangeRequestMessage) Deserialize(r Reader) error {

	return nil
}

type ExchangeObjectsRemovedMessage struct {
	ExchangeObjectMessage

	ObjectUID []uint32
}

func (m *ExchangeObjectsRemovedMessage) ID() uint16 {
	return 6532
}

func (m *ExchangeObjectsRemovedMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.ObjectUID))); err != nil {
		return err
	}

	for i := range m.ObjectUID {

		if err := w.WriteVarUInt32(m.ObjectUID[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeObjectsRemovedMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMessage.Deserialize(r); err != nil {
		return err
	}

	lobjectUIDLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectUID = make([]uint32, lobjectUIDLen)

	for i := range m.ObjectUID {

		lobjectUID, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.ObjectUID[i] = lobjectUID

	}

	return nil
}

type ExchangeObjectsModifiedMessage struct {
	ExchangeObjectMessage

	Object []ObjectItem
}

func (m *ExchangeObjectsModifiedMessage) ID() uint16 {
	return 6533
}

func (m *ExchangeObjectsModifiedMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Object))); err != nil {
		return err
	}

	for i := range m.Object {

		if err := m.Object[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeObjectsModifiedMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMessage.Deserialize(r); err != nil {
		return err
	}

	lobjectLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Object = make([]ObjectItem, lobjectLen)

	for i := range m.Object {

		var lobject ObjectItem

		lobject.Deserialize(r)

		m.Object[i] = lobject

	}

	return nil
}

type ExchangeObjectsAddedMessage struct {
	ExchangeObjectMessage

	Object []ObjectItem
}

func (m *ExchangeObjectsAddedMessage) ID() uint16 {
	return 6535
}

func (m *ExchangeObjectsAddedMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Object))); err != nil {
		return err
	}

	for i := range m.Object {

		if err := m.Object[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeObjectsAddedMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMessage.Deserialize(r); err != nil {
		return err
	}

	lobjectLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Object = make([]ObjectItem, lobjectLen)

	for i := range m.Object {

		var lobject ObjectItem

		lobject.Deserialize(r)

		m.Object[i] = lobject

	}

	return nil
}

type ComicReadingBeginMessage struct {
	ComicId uint16
}

func (m *ComicReadingBeginMessage) ID() uint16 {
	return 6536
}

func (m *ComicReadingBeginMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.ComicId); err != nil {
		return err
	}

	return nil
}

func (m *ComicReadingBeginMessage) Deserialize(r Reader) error {

	lcomicId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ComicId = lcomicId

	return nil
}

type StartupActionsAllAttributionMessage struct {
	CharacterId int64
}

func (m *StartupActionsAllAttributionMessage) ID() uint16 {
	return 6537
}

func (m *StartupActionsAllAttributionMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.CharacterId); err != nil {
		return err
	}

	return nil
}

func (m *StartupActionsAllAttributionMessage) Deserialize(r Reader) error {

	lcharacterId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.CharacterId = lcharacterId

	return nil
}

type StartupActionAddMessage struct {
	NewAction StartupActionAddObject
}

func (m *StartupActionAddMessage) ID() uint16 {
	return 6538
}

func (m *StartupActionAddMessage) Serialize(w Writer) error {

	if err := m.NewAction.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *StartupActionAddMessage) Deserialize(r Reader) error {

	var lnewAction StartupActionAddObject

	lnewAction.Deserialize(r)

	m.NewAction = lnewAction

	return nil
}

type ReloginTokenStatusMessage struct {
	ValidToken bool

	Ticket []int8
}

func (m *ReloginTokenStatusMessage) ID() uint16 {
	return 6539
}

func (m *ReloginTokenStatusMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.ValidToken); err != nil {
		return err
	}

	if err := w.WriteVarInt32(int32(len(m.Ticket))); err != nil {
		return err
	}

	for i := range m.Ticket {

		if err := w.WriteInt8(m.Ticket[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ReloginTokenStatusMessage) Deserialize(r Reader) error {

	lvalidToken, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.ValidToken = lvalidToken

	lticketLen, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Ticket = make([]int8, lticketLen)

	for i := range m.Ticket {

		lticket, err := r.ReadInt8()
		if err != nil {
			return err
		}

		m.Ticket[i] = lticket

	}

	return nil
}

type ReloginTokenRequestMessage struct {
}

func (m *ReloginTokenRequestMessage) ID() uint16 {
	return 6540
}

func (m *ReloginTokenRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *ReloginTokenRequestMessage) Deserialize(r Reader) error {

	return nil
}

type GameFightPlacementSwapPositionsRequestMessage struct {
	GameFightPlacementPositionRequestMessage

	RequestedId float64
}

func (m *GameFightPlacementSwapPositionsRequestMessage) ID() uint16 {
	return 6541
}

func (m *GameFightPlacementSwapPositionsRequestMessage) Serialize(w Writer) error {

	if err := m.GameFightPlacementPositionRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteDouble(m.RequestedId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightPlacementSwapPositionsRequestMessage) Deserialize(r Reader) error {

	if err := m.GameFightPlacementPositionRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lrequestedId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.RequestedId = lrequestedId

	return nil
}

type GameFightPlacementSwapPositionsOfferMessage struct {
	RequestId uint32

	RequesterId float64

	RequesterCellId uint16

	RequestedId float64

	RequestedCellId uint16
}

func (m *GameFightPlacementSwapPositionsOfferMessage) ID() uint16 {
	return 6542
}

func (m *GameFightPlacementSwapPositionsOfferMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.RequestId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.RequesterId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.RequesterCellId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.RequestedId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.RequestedCellId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightPlacementSwapPositionsOfferMessage) Deserialize(r Reader) error {

	lrequestId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.RequestId = lrequestId

	lrequesterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.RequesterId = lrequesterId

	lrequesterCellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.RequesterCellId = lrequesterCellId

	lrequestedId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.RequestedId = lrequestedId

	lrequestedCellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.RequestedCellId = lrequestedCellId

	return nil
}

type GameFightPlacementSwapPositionsCancelMessage struct {
	RequestId uint32
}

func (m *GameFightPlacementSwapPositionsCancelMessage) ID() uint16 {
	return 6543
}

func (m *GameFightPlacementSwapPositionsCancelMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.RequestId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightPlacementSwapPositionsCancelMessage) Deserialize(r Reader) error {

	lrequestId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.RequestId = lrequestId

	return nil
}

type GameFightPlacementSwapPositionsMessage struct {
	Dispositions []IdentifiedEntityDispositionInformations
}

func (m *GameFightPlacementSwapPositionsMessage) ID() uint16 {
	return 6544
}

func (m *GameFightPlacementSwapPositionsMessage) Serialize(w Writer) error {

	for i := range m.Dispositions {

		if err := m.Dispositions[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameFightPlacementSwapPositionsMessage) Deserialize(r Reader) error {

	m.Dispositions = make([]IdentifiedEntityDispositionInformations, 2)

	for i := range m.Dispositions {

		var ldispositions IdentifiedEntityDispositionInformations

		ldispositions.Deserialize(r)

		m.Dispositions[i] = ldispositions

	}

	return nil
}

type GameActionFightActivateGlyphTrapMessage struct {
	AbstractGameActionMessage

	MarkId int16

	Active bool
}

func (m *GameActionFightActivateGlyphTrapMessage) ID() uint16 {
	return 6545
}

func (m *GameActionFightActivateGlyphTrapMessage) Serialize(w Writer) error {

	if err := m.AbstractGameActionMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(m.MarkId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Active); err != nil {
		return err
	}

	return nil
}

func (m *GameActionFightActivateGlyphTrapMessage) Deserialize(r Reader) error {

	if err := m.AbstractGameActionMessage.Deserialize(r); err != nil {
		return err
	}

	lmarkId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MarkId = lmarkId

	lactive, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Active = lactive

	return nil
}

type GameFightPlacementSwapPositionsCancelledMessage struct {
	RequestId uint32

	CancellerId float64
}

func (m *GameFightPlacementSwapPositionsCancelledMessage) ID() uint16 {
	return 6546
}

func (m *GameFightPlacementSwapPositionsCancelledMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.RequestId); err != nil {
		return err
	}

	if err := w.WriteDouble(m.CancellerId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightPlacementSwapPositionsCancelledMessage) Deserialize(r Reader) error {

	lrequestId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.RequestId = lrequestId

	lcancellerId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.CancellerId = lcancellerId

	return nil
}

type GameFightPlacementSwapPositionsAcceptMessage struct {
	RequestId uint32
}

func (m *GameFightPlacementSwapPositionsAcceptMessage) ID() uint16 {
	return 6547
}

func (m *GameFightPlacementSwapPositionsAcceptMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.RequestId); err != nil {
		return err
	}

	return nil
}

func (m *GameFightPlacementSwapPositionsAcceptMessage) Deserialize(r Reader) error {

	lrequestId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.RequestId = lrequestId

	return nil
}

type GameFightPlacementSwapPositionsErrorMessage struct {
}

func (m *GameFightPlacementSwapPositionsErrorMessage) ID() uint16 {
	return 6548
}

func (m *GameFightPlacementSwapPositionsErrorMessage) Serialize(w Writer) error {

	return nil
}

func (m *GameFightPlacementSwapPositionsErrorMessage) Deserialize(r Reader) error {

	return nil
}

type CharacterSelectionWithRemodelMessage struct {
	CharacterSelectionMessage

	Remodel RemodelingInformation
}

func (m *CharacterSelectionWithRemodelMessage) ID() uint16 {
	return 6549
}

func (m *CharacterSelectionWithRemodelMessage) Serialize(w Writer) error {

	if err := m.CharacterSelectionMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.Remodel.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *CharacterSelectionWithRemodelMessage) Deserialize(r Reader) error {

	if err := m.CharacterSelectionMessage.Deserialize(r); err != nil {
		return err
	}

	var lremodel RemodelingInformation

	lremodel.Deserialize(r)

	m.Remodel = lremodel

	return nil
}

type CharactersListWithRemodelingMessage struct {
	CharactersListMessage

	CharactersToRemodel []CharacterToRemodelInformations
}

func (m *CharactersListWithRemodelingMessage) ID() uint16 {
	return 6550
}

func (m *CharactersListWithRemodelingMessage) Serialize(w Writer) error {

	if err := m.CharactersListMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.CharactersToRemodel))); err != nil {
		return err
	}

	for i := range m.CharactersToRemodel {

		if err := m.CharactersToRemodel[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *CharactersListWithRemodelingMessage) Deserialize(r Reader) error {

	if err := m.CharactersListMessage.Deserialize(r); err != nil {
		return err
	}

	lcharactersToRemodelLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CharactersToRemodel = make([]CharacterToRemodelInformations, lcharactersToRemodelLen)

	for i := range m.CharactersToRemodel {

		var lcharactersToRemodel CharacterToRemodelInformations

		lcharactersToRemodel.Deserialize(r)

		m.CharactersToRemodel[i] = lcharactersToRemodel

	}

	return nil
}

type CharacterReplayWithRemodelRequestMessage struct {
	CharacterReplayRequestMessage

	Remodel RemodelingInformation
}

func (m *CharacterReplayWithRemodelRequestMessage) ID() uint16 {
	return 6551
}

func (m *CharacterReplayWithRemodelRequestMessage) Serialize(w Writer) error {

	if err := m.CharacterReplayRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.Remodel.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *CharacterReplayWithRemodelRequestMessage) Deserialize(r Reader) error {

	if err := m.CharacterReplayRequestMessage.Deserialize(r); err != nil {
		return err
	}

	var lremodel RemodelingInformation

	lremodel.Deserialize(r)

	m.Remodel = lremodel

	return nil
}

type ExchangeMountsTakenFromPaddockMessage struct {
	Name string

	WorldX int16

	WorldY int16

	Ownername string
}

func (m *ExchangeMountsTakenFromPaddockMessage) ID() uint16 {
	return 6554
}

func (m *ExchangeMountsTakenFromPaddockMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Name); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldX); err != nil {
		return err
	}

	if err := w.WriteInt16(m.WorldY); err != nil {
		return err
	}

	if err := w.WriteString(m.Ownername); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeMountsTakenFromPaddockMessage) Deserialize(r Reader) error {

	lname, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Name = lname

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

	lownername, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Ownername = lownername

	return nil
}

type ExchangeMountsStableAddMessage struct {
	MountDescription []MountClientData
}

func (m *ExchangeMountsStableAddMessage) ID() uint16 {
	return 6555
}

func (m *ExchangeMountsStableAddMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.MountDescription))); err != nil {
		return err
	}

	for i := range m.MountDescription {

		if err := m.MountDescription[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeMountsStableAddMessage) Deserialize(r Reader) error {

	lmountDescriptionLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MountDescription = make([]MountClientData, lmountDescriptionLen)

	for i := range m.MountDescription {

		var lmountDescription MountClientData

		lmountDescription.Deserialize(r)

		m.MountDescription[i] = lmountDescription

	}

	return nil
}

type ExchangeMountsStableRemoveMessage struct {
	MountsId []int32
}

func (m *ExchangeMountsStableRemoveMessage) ID() uint16 {
	return 6556
}

func (m *ExchangeMountsStableRemoveMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.MountsId))); err != nil {
		return err
	}

	for i := range m.MountsId {

		if err := w.WriteVarInt32(m.MountsId[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeMountsStableRemoveMessage) Deserialize(r Reader) error {

	lmountsIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MountsId = make([]int32, lmountsIdLen)

	for i := range m.MountsId {

		lmountsId, err := r.ReadVarInt32()
		if err != nil {
			return err
		}

		m.MountsId[i] = lmountsId

	}

	return nil
}

type ExchangeMountsStableBornAddMessage struct {
	ExchangeMountsStableAddMessage
}

func (m *ExchangeMountsStableBornAddMessage) ID() uint16 {
	return 6557
}

func (m *ExchangeMountsStableBornAddMessage) Serialize(w Writer) error {

	if err := m.ExchangeMountsStableAddMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeMountsStableBornAddMessage) Deserialize(r Reader) error {

	if err := m.ExchangeMountsStableAddMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type ExchangeMountsPaddockRemoveMessage struct {
	MountsId []int32
}

func (m *ExchangeMountsPaddockRemoveMessage) ID() uint16 {
	return 6559
}

func (m *ExchangeMountsPaddockRemoveMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.MountsId))); err != nil {
		return err
	}

	for i := range m.MountsId {

		if err := w.WriteVarInt32(m.MountsId[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeMountsPaddockRemoveMessage) Deserialize(r Reader) error {

	lmountsIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MountsId = make([]int32, lmountsIdLen)

	for i := range m.MountsId {

		lmountsId, err := r.ReadVarInt32()
		if err != nil {
			return err
		}

		m.MountsId[i] = lmountsId

	}

	return nil
}

type ExchangeMountsPaddockAddMessage struct {
	MountDescription []MountClientData
}

func (m *ExchangeMountsPaddockAddMessage) ID() uint16 {
	return 6561
}

func (m *ExchangeMountsPaddockAddMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.MountDescription))); err != nil {
		return err
	}

	for i := range m.MountDescription {

		if err := m.MountDescription[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeMountsPaddockAddMessage) Deserialize(r Reader) error {

	lmountDescriptionLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MountDescription = make([]MountClientData, lmountDescriptionLen)

	for i := range m.MountDescription {

		var lmountDescription MountClientData

		lmountDescription.Deserialize(r)

		m.MountDescription[i] = lmountDescription

	}

	return nil
}

type ExchangeHandleMountsStableMessage struct {
	ActionType int8

	RidesId []uint32
}

func (m *ExchangeHandleMountsStableMessage) ID() uint16 {
	return 6562
}

func (m *ExchangeHandleMountsStableMessage) Serialize(w Writer) error {

	if err := w.WriteInt8(m.ActionType); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.RidesId))); err != nil {
		return err
	}

	for i := range m.RidesId {

		if err := w.WriteVarUInt32(m.RidesId[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeHandleMountsStableMessage) Deserialize(r Reader) error {

	lactionType, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.ActionType = lactionType

	lridesIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.RidesId = make([]uint32, lridesIdLen)

	for i := range m.RidesId {

		lridesId, err := r.ReadVarUInt32()
		if err != nil {
			return err
		}

		m.RidesId[i] = lridesId

	}

	return nil
}

type DisplayNumericalValuePaddockMessage struct {
	RideId int32

	Value int32

	Type uint8
}

func (m *DisplayNumericalValuePaddockMessage) ID() uint16 {
	return 6563
}

func (m *DisplayNumericalValuePaddockMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.RideId); err != nil {
		return err
	}

	if err := w.WriteInt32(m.Value); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DisplayNumericalValuePaddockMessage) Deserialize(r Reader) error {

	lrideId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.RideId = lrideId

	lvalue, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Value = lvalue

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	return nil
}

type TopTaxCollectorListMessage struct {
	AbstractTaxCollectorListMessage

	IsDungeon bool
}

func (m *TopTaxCollectorListMessage) ID() uint16 {
	return 6565
}

func (m *TopTaxCollectorListMessage) Serialize(w Writer) error {

	if err := m.AbstractTaxCollectorListMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.IsDungeon); err != nil {
		return err
	}

	return nil
}

func (m *TopTaxCollectorListMessage) Deserialize(r Reader) error {

	if err := m.AbstractTaxCollectorListMessage.Deserialize(r); err != nil {
		return err
	}

	lisDungeon, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.IsDungeon = lisDungeon

	return nil
}

type ExchangeStartOkRunesTradeMessage struct {
}

func (m *ExchangeStartOkRunesTradeMessage) ID() uint16 {
	return 6567
}

func (m *ExchangeStartOkRunesTradeMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExchangeStartOkRunesTradeMessage) Deserialize(r Reader) error {

	return nil
}

type AbstractTaxCollectorListMessage struct {
	Informations []TaxCollectorInformationsIntrf
}

func (m *AbstractTaxCollectorListMessage) ID() uint16 {
	return 6568
}

func (m *AbstractTaxCollectorListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Informations))); err != nil {
		return err
	}

	for i := range m.Informations {

		if err := w.WriteUInt16(m.Informations[i].ID()); err != nil {
			return err
		}

		if err := m.Informations[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *AbstractTaxCollectorListMessage) Deserialize(r Reader) error {

	linformationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Informations = make([]TaxCollectorInformationsIntrf, linformationsLen)

	for i := range m.Informations {

		typeinformationsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		linformations, err := GetType(typeinformationsID)
		if err != nil {
			return err
		}

		linformations.Deserialize(r)

		m.Informations[i] = linformations.(TaxCollectorInformationsIntrf)

	}

	return nil
}

type DecraftResultMessage struct {
	Results []DecraftedItemStackInfo
}

func (m *DecraftResultMessage) ID() uint16 {
	return 6569
}

func (m *DecraftResultMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Results))); err != nil {
		return err
	}

	for i := range m.Results {

		if err := m.Results[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *DecraftResultMessage) Deserialize(r Reader) error {

	lresultsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Results = make([]DecraftedItemStackInfo, lresultsLen)

	for i := range m.Results {

		var lresults DecraftedItemStackInfo

		lresults.Deserialize(r)

		m.Results[i] = lresults

	}

	return nil
}

type ZaapRespawnUpdatedMessage struct {
	MapId uint32
}

func (m *ZaapRespawnUpdatedMessage) ID() uint16 {
	return 6571
}

func (m *ZaapRespawnUpdatedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.MapId); err != nil {
		return err
	}

	return nil
}

func (m *ZaapRespawnUpdatedMessage) Deserialize(r Reader) error {

	lmapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	return nil
}

type ZaapRespawnSaveRequestMessage struct {
}

func (m *ZaapRespawnSaveRequestMessage) ID() uint16 {
	return 6572
}

func (m *ZaapRespawnSaveRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *ZaapRespawnSaveRequestMessage) Deserialize(r Reader) error {

	return nil
}

type BasicStatWithDataMessage struct {
	BasicStatMessage

	Datas []StatisticDataIntrf
}

func (m *BasicStatWithDataMessage) ID() uint16 {
	return 6573
}

func (m *BasicStatWithDataMessage) Serialize(w Writer) error {

	if err := m.BasicStatMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Datas))); err != nil {
		return err
	}

	for i := range m.Datas {

		if err := w.WriteUInt16(m.Datas[i].ID()); err != nil {
			return err
		}

		if err := m.Datas[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *BasicStatWithDataMessage) Deserialize(r Reader) error {

	if err := m.BasicStatMessage.Deserialize(r); err != nil {
		return err
	}

	ldatasLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Datas = make([]StatisticDataIntrf, ldatasLen)

	for i := range m.Datas {

		typedatasID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		ldatas, err := GetType(typedatasID)
		if err != nil {
			return err
		}

		ldatas.Deserialize(r)

		m.Datas[i] = ldatas.(StatisticDataIntrf)

	}

	return nil
}

type GameRolePlayArenaSwitchToGameServerMessage struct {
	ValidToken bool

	Ticket []int8

	HomeServerId int16
}

func (m *GameRolePlayArenaSwitchToGameServerMessage) ID() uint16 {
	return 6574
}

func (m *GameRolePlayArenaSwitchToGameServerMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.ValidToken); err != nil {
		return err
	}

	if err := w.WriteVarInt32(int32(len(m.Ticket))); err != nil {
		return err
	}

	for i := range m.Ticket {

		if err := w.WriteInt8(m.Ticket[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(m.HomeServerId); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayArenaSwitchToGameServerMessage) Deserialize(r Reader) error {

	lvalidToken, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.ValidToken = lvalidToken

	lticketLen, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Ticket = make([]int8, lticketLen)

	for i := range m.Ticket {

		lticket, err := r.ReadInt8()
		if err != nil {
			return err
		}

		m.Ticket[i] = lticket

	}

	lhomeServerId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.HomeServerId = lhomeServerId

	return nil
}

type GameRolePlayArenaSwitchToFightServerMessage struct {
	Address string

	Port uint16

	Ticket []int8
}

func (m *GameRolePlayArenaSwitchToFightServerMessage) ID() uint16 {
	return 6575
}

func (m *GameRolePlayArenaSwitchToFightServerMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Address); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Port); err != nil {
		return err
	}

	if err := w.WriteVarInt32(int32(len(m.Ticket))); err != nil {
		return err
	}

	for i := range m.Ticket {

		if err := w.WriteInt8(m.Ticket[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameRolePlayArenaSwitchToFightServerMessage) Deserialize(r Reader) error {

	laddress, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Address = laddress

	lport, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.Port = lport

	lticketLen, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Ticket = make([]int8, lticketLen)

	for i := range m.Ticket {

		lticket, err := r.ReadInt8()
		if err != nil {
			return err
		}

		m.Ticket[i] = lticket

	}

	return nil
}

type ExchangeCraftPaymentModifiedMessage struct {
	GoldSum int64
}

func (m *ExchangeCraftPaymentModifiedMessage) ID() uint16 {
	return 6578
}

func (m *ExchangeCraftPaymentModifiedMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.GoldSum); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeCraftPaymentModifiedMessage) Deserialize(r Reader) error {

	lgoldSum, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.GoldSum = lgoldSum

	return nil
}

type ExchangeCraftPaymentModificationRequestMessage struct {
	Quantity int64
}

func (m *ExchangeCraftPaymentModificationRequestMessage) ID() uint16 {
	return 6579
}

func (m *ExchangeCraftPaymentModificationRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeCraftPaymentModificationRequestMessage) Deserialize(r Reader) error {

	lquantity, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Quantity = lquantity

	return nil
}

type IdolPartyLostMessage struct {
	IdolId uint16
}

func (m *IdolPartyLostMessage) ID() uint16 {
	return 6580
}

func (m *IdolPartyLostMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.IdolId); err != nil {
		return err
	}

	return nil
}

func (m *IdolPartyLostMessage) Deserialize(r Reader) error {

	lidolId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.IdolId = lidolId

	return nil
}

type IdolSelectedMessage struct {
	IdolId uint16

	Activate bool

	Party bool
}

func (m *IdolSelectedMessage) ID() uint16 {
	return 6581
}

func (m *IdolSelectedMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Activate)

	setWrappedFlag(bbw0, 1, m.Party)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.IdolId); err != nil {
		return err
	}

	return nil
}

func (m *IdolSelectedMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Activate = getWrappedFlag(bbw0, 0)

	m.Party = getWrappedFlag(bbw0, 1)

	lidolId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.IdolId = lidolId

	return nil
}

type IdolPartyRegisterRequestMessage struct {
	Register bool
}

func (m *IdolPartyRegisterRequestMessage) ID() uint16 {
	return 6582
}

func (m *IdolPartyRegisterRequestMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Register); err != nil {
		return err
	}

	return nil
}

func (m *IdolPartyRegisterRequestMessage) Deserialize(r Reader) error {

	lregister, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Register = lregister

	return nil
}

type IdolPartyRefreshMessage struct {
	PartyIdol PartyIdol
}

func (m *IdolPartyRefreshMessage) ID() uint16 {
	return 6583
}

func (m *IdolPartyRefreshMessage) Serialize(w Writer) error {

	if err := m.PartyIdol.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *IdolPartyRefreshMessage) Deserialize(r Reader) error {

	var lpartyIdol PartyIdol

	lpartyIdol.Deserialize(r)

	m.PartyIdol = lpartyIdol

	return nil
}

type IdolSelectErrorMessage struct {
	Reason uint8

	IdolId uint16

	Activate bool

	Party bool
}

func (m *IdolSelectErrorMessage) ID() uint16 {
	return 6584
}

func (m *IdolSelectErrorMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Activate)

	setWrappedFlag(bbw0, 1, m.Party)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.IdolId); err != nil {
		return err
	}

	return nil
}

func (m *IdolSelectErrorMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Activate = getWrappedFlag(bbw0, 0)

	m.Party = getWrappedFlag(bbw0, 1)

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	lidolId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.IdolId = lidolId

	return nil
}

type IdolListMessage struct {
	ChosenIdols []uint16

	PartyChosenIdols []uint16

	PartyIdols []PartyIdolIntrf
}

func (m *IdolListMessage) ID() uint16 {
	return 6585
}

func (m *IdolListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.ChosenIdols))); err != nil {
		return err
	}

	for i := range m.ChosenIdols {

		if err := w.WriteVarUInt16(m.ChosenIdols[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.PartyChosenIdols))); err != nil {
		return err
	}

	for i := range m.PartyChosenIdols {

		if err := w.WriteVarUInt16(m.PartyChosenIdols[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.PartyIdols))); err != nil {
		return err
	}

	for i := range m.PartyIdols {

		if err := w.WriteUInt16(m.PartyIdols[i].ID()); err != nil {
			return err
		}

		if err := m.PartyIdols[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *IdolListMessage) Deserialize(r Reader) error {

	lchosenIdolsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ChosenIdols = make([]uint16, lchosenIdolsLen)

	for i := range m.ChosenIdols {

		lchosenIdols, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.ChosenIdols[i] = lchosenIdols

	}

	lpartyChosenIdolsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PartyChosenIdols = make([]uint16, lpartyChosenIdolsLen)

	for i := range m.PartyChosenIdols {

		lpartyChosenIdols, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.PartyChosenIdols[i] = lpartyChosenIdols

	}

	lpartyIdolsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PartyIdols = make([]PartyIdolIntrf, lpartyIdolsLen)

	for i := range m.PartyIdols {

		typepartyIdolsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lpartyIdols, err := GetType(typepartyIdolsID)
		if err != nil {
			return err
		}

		lpartyIdols.Deserialize(r)

		m.PartyIdols[i] = lpartyIdols.(PartyIdolIntrf)

	}

	return nil
}

type IdolFightPreparationUpdateMessage struct {
	IdolSource uint8

	Idols []IdolIntrf
}

func (m *IdolFightPreparationUpdateMessage) ID() uint16 {
	return 6586
}

func (m *IdolFightPreparationUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.IdolSource); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Idols))); err != nil {
		return err
	}

	for i := range m.Idols {

		if err := w.WriteUInt16(m.Idols[i].ID()); err != nil {
			return err
		}

		if err := m.Idols[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *IdolFightPreparationUpdateMessage) Deserialize(r Reader) error {

	lidolSource, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.IdolSource = lidolSource

	lidolsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Idols = make([]IdolIntrf, lidolsLen)

	for i := range m.Idols {

		typeidolsID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		lidols, err := GetType(typeidolsID)
		if err != nil {
			return err
		}

		lidols.Deserialize(r)

		m.Idols[i] = lidols.(IdolIntrf)

	}

	return nil
}

type IdolSelectRequestMessage struct {
	IdolId uint16

	Activate bool

	Party bool
}

func (m *IdolSelectRequestMessage) ID() uint16 {
	return 6587
}

func (m *IdolSelectRequestMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Activate)

	setWrappedFlag(bbw0, 1, m.Party)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.IdolId); err != nil {
		return err
	}

	return nil
}

func (m *IdolSelectRequestMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Activate = getWrappedFlag(bbw0, 0)

	m.Party = getWrappedFlag(bbw0, 1)

	lidolId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.IdolId = lidolId

	return nil
}

type GuildMotdSetRequestMessage struct {
	SocialNoticeSetRequestMessage

	Content string
}

func (m *GuildMotdSetRequestMessage) ID() uint16 {
	return 6588
}

func (m *GuildMotdSetRequestMessage) Serialize(w Writer) error {

	if err := m.SocialNoticeSetRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Content); err != nil {
		return err
	}

	return nil
}

func (m *GuildMotdSetRequestMessage) Deserialize(r Reader) error {

	if err := m.SocialNoticeSetRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lcontent, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Content = lcontent

	return nil
}

type ExchangeStoppedMessage struct {
	Id int64
}

func (m *ExchangeStoppedMessage) ID() uint16 {
	return 6589
}

func (m *ExchangeStoppedMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeStoppedMessage) Deserialize(r Reader) error {

	lid, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type GuildMotdMessage struct {
	SocialNoticeMessage
}

func (m *GuildMotdMessage) ID() uint16 {
	return 6590
}

func (m *GuildMotdMessage) Serialize(w Writer) error {

	if err := m.SocialNoticeMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildMotdMessage) Deserialize(r Reader) error {

	if err := m.SocialNoticeMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type GuildMotdSetErrorMessage struct {
	SocialNoticeSetErrorMessage
}

func (m *GuildMotdSetErrorMessage) ID() uint16 {
	return 6591
}

func (m *GuildMotdSetErrorMessage) Serialize(w Writer) error {

	if err := m.SocialNoticeSetErrorMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildMotdSetErrorMessage) Deserialize(r Reader) error {

	if err := m.SocialNoticeSetErrorMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type JobBookSubscribeRequestMessage struct {
	JobIds []uint8
}

func (m *JobBookSubscribeRequestMessage) ID() uint16 {
	return 6592
}

func (m *JobBookSubscribeRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.JobIds))); err != nil {
		return err
	}

	for i := range m.JobIds {

		if err := w.WriteUInt8(m.JobIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *JobBookSubscribeRequestMessage) Deserialize(r Reader) error {

	ljobIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.JobIds = make([]uint8, ljobIdsLen)

	for i := range m.JobIds {

		ljobIds, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.JobIds[i] = ljobIds

	}

	return nil
}

type JobBookSubscriptionMessage struct {
	Subscriptions []JobBookSubscription
}

func (m *JobBookSubscriptionMessage) ID() uint16 {
	return 6593
}

func (m *JobBookSubscriptionMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Subscriptions))); err != nil {
		return err
	}

	for i := range m.Subscriptions {

		if err := m.Subscriptions[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *JobBookSubscriptionMessage) Deserialize(r Reader) error {

	lsubscriptionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Subscriptions = make([]JobBookSubscription, lsubscriptionsLen)

	for i := range m.Subscriptions {

		var lsubscriptions JobBookSubscription

		lsubscriptions.Deserialize(r)

		m.Subscriptions[i] = lsubscriptions

	}

	return nil
}

type ClientYouAreDrunkMessage struct {
	DebugInClientMessage
}

func (m *ClientYouAreDrunkMessage) ID() uint16 {
	return 6594
}

func (m *ClientYouAreDrunkMessage) Serialize(w Writer) error {

	if err := m.DebugInClientMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ClientYouAreDrunkMessage) Deserialize(r Reader) error {

	if err := m.DebugInClientMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type ExchangeCraftCountModifiedMessage struct {
	Count int32
}

func (m *ExchangeCraftCountModifiedMessage) ID() uint16 {
	return 6595
}

func (m *ExchangeCraftCountModifiedMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt32(m.Count); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeCraftCountModifiedMessage) Deserialize(r Reader) error {

	lcount, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Count = lcount

	return nil
}

type ChatSmileyExtraPackListMessage struct {
	PackIds []uint8
}

func (m *ChatSmileyExtraPackListMessage) ID() uint16 {
	return 6596
}

func (m *ChatSmileyExtraPackListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.PackIds))); err != nil {
		return err
	}

	for i := range m.PackIds {

		if err := w.WriteUInt8(m.PackIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ChatSmileyExtraPackListMessage) Deserialize(r Reader) error {

	lpackIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PackIds = make([]uint8, lpackIdsLen)

	for i := range m.PackIds {

		lpackIds, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.PackIds[i] = lpackIds

	}

	return nil
}

type ExchangeCraftCountRequestMessage struct {
	Count int32
}

func (m *ExchangeCraftCountRequestMessage) ID() uint16 {
	return 6597
}

func (m *ExchangeCraftCountRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt32(m.Count); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeCraftCountRequestMessage) Deserialize(r Reader) error {

	lcount, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.Count = lcount

	return nil
}

type ExchangeCrafterJobLevelupMessage struct {
	CrafterJobLevel uint8
}

func (m *ExchangeCrafterJobLevelupMessage) ID() uint16 {
	return 6598
}

func (m *ExchangeCrafterJobLevelupMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.CrafterJobLevel); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeCrafterJobLevelupMessage) Deserialize(r Reader) error {

	lcrafterJobLevel, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.CrafterJobLevel = lcrafterJobLevel

	return nil
}

type JobExperienceOtherPlayerUpdateMessage struct {
	JobExperienceUpdateMessage

	PlayerId int64
}

func (m *JobExperienceOtherPlayerUpdateMessage) ID() uint16 {
	return 6599
}

func (m *JobExperienceOtherPlayerUpdateMessage) Serialize(w Writer) error {

	if err := m.JobExperienceUpdateMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.PlayerId); err != nil {
		return err
	}

	return nil
}

func (m *JobExperienceOtherPlayerUpdateMessage) Deserialize(r Reader) error {

	if err := m.JobExperienceUpdateMessage.Deserialize(r); err != nil {
		return err
	}

	lplayerId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.PlayerId = lplayerId

	return nil
}

type ExchangeStartOkRecycleTradeMessage struct {
	PercentToPrism uint16

	PercentToPlayer uint16
}

func (m *ExchangeStartOkRecycleTradeMessage) ID() uint16 {
	return 6600
}

func (m *ExchangeStartOkRecycleTradeMessage) Serialize(w Writer) error {

	if err := w.WriteUInt16(m.PercentToPrism); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.PercentToPlayer); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeStartOkRecycleTradeMessage) Deserialize(r Reader) error {

	lpercentToPrism, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.PercentToPrism = lpercentToPrism

	lpercentToPlayer, err := r.ReadUInt16()
	if err != nil {
		return err
	}

	m.PercentToPlayer = lpercentToPlayer

	return nil
}

type RecycleResultMessage struct {
	NuggetsForPrism uint32

	NuggetsForPlayer uint32
}

func (m *RecycleResultMessage) ID() uint16 {
	return 6601
}

func (m *RecycleResultMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.NuggetsForPrism); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.NuggetsForPlayer); err != nil {
		return err
	}

	return nil
}

func (m *RecycleResultMessage) Deserialize(r Reader) error {

	lnuggetsForPrism, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.NuggetsForPrism = lnuggetsForPrism

	lnuggetsForPlayer, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.NuggetsForPlayer = lnuggetsForPlayer

	return nil
}

type IdolsPresetDeleteMessage struct {
	PresetId uint8
}

func (m *IdolsPresetDeleteMessage) ID() uint16 {
	return 6602
}

func (m *IdolsPresetDeleteMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	return nil
}

func (m *IdolsPresetDeleteMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	return nil
}

type IdolsPresetSaveMessage struct {
	PresetId uint8

	SymbolId uint8
}

func (m *IdolsPresetSaveMessage) ID() uint16 {
	return 6603
}

func (m *IdolsPresetSaveMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.SymbolId); err != nil {
		return err
	}

	return nil
}

func (m *IdolsPresetSaveMessage) Deserialize(r Reader) error {

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

	return nil
}

type IdolsPresetSaveResultMessage struct {
	PresetId uint8

	Code uint8
}

func (m *IdolsPresetSaveResultMessage) ID() uint16 {
	return 6604
}

func (m *IdolsPresetSaveResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Code); err != nil {
		return err
	}

	return nil
}

func (m *IdolsPresetSaveResultMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	lcode, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Code = lcode

	return nil
}

type IdolsPresetDeleteResultMessage struct {
	PresetId uint8

	Code uint8
}

func (m *IdolsPresetDeleteResultMessage) ID() uint16 {
	return 6605
}

func (m *IdolsPresetDeleteResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Code); err != nil {
		return err
	}

	return nil
}

func (m *IdolsPresetDeleteResultMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	lcode, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Code = lcode

	return nil
}

type IdolsPresetUpdateMessage struct {
	IdolsPreset IdolsPreset
}

func (m *IdolsPresetUpdateMessage) ID() uint16 {
	return 6606
}

func (m *IdolsPresetUpdateMessage) Serialize(w Writer) error {

	if err := m.IdolsPreset.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *IdolsPresetUpdateMessage) Deserialize(r Reader) error {

	var lidolsPreset IdolsPreset

	lidolsPreset.Deserialize(r)

	m.IdolsPreset = lidolsPreset

	return nil
}

type AccountLinkRequiredMessage struct {
}

func (m *AccountLinkRequiredMessage) ID() uint16 {
	return 6607
}

func (m *AccountLinkRequiredMessage) Serialize(w Writer) error {

	return nil
}

func (m *AccountLinkRequiredMessage) Deserialize(r Reader) error {

	return nil
}

type TaxCollectorMovementsOfflineMessage struct {
	Movements []TaxCollectorMovement
}

func (m *TaxCollectorMovementsOfflineMessage) ID() uint16 {
	return 6611
}

func (m *TaxCollectorMovementsOfflineMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Movements))); err != nil {
		return err
	}

	for i := range m.Movements {

		if err := m.Movements[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *TaxCollectorMovementsOfflineMessage) Deserialize(r Reader) error {

	lmovementsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Movements = make([]TaxCollectorMovement, lmovementsLen)

	for i := range m.Movements {

		var lmovements TaxCollectorMovement

		lmovements.Deserialize(r)

		m.Movements[i] = lmovements

	}

	return nil
}

type ExchangeBidHouseUnsoldItemsMessage struct {
	Items []ObjectItemGenericQuantity
}

func (m *ExchangeBidHouseUnsoldItemsMessage) ID() uint16 {
	return 6612
}

func (m *ExchangeBidHouseUnsoldItemsMessage) Serialize(w Writer) error {

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

func (m *ExchangeBidHouseUnsoldItemsMessage) Deserialize(r Reader) error {

	litemsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Items = make([]ObjectItemGenericQuantity, litemsLen)

	for i := range m.Items {

		var litems ObjectItemGenericQuantity

		litems.Deserialize(r)

		m.Items[i] = litems

	}

	return nil
}

type ExchangeOfflineSoldItemsMessage struct {
	BidHouseItems []ObjectItemGenericQuantityPrice

	MerchantItems []ObjectItemGenericQuantityPrice
}

func (m *ExchangeOfflineSoldItemsMessage) ID() uint16 {
	return 6613
}

func (m *ExchangeOfflineSoldItemsMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.BidHouseItems))); err != nil {
		return err
	}

	for i := range m.BidHouseItems {

		if err := m.BidHouseItems[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.MerchantItems))); err != nil {
		return err
	}

	for i := range m.MerchantItems {

		if err := m.MerchantItems[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *ExchangeOfflineSoldItemsMessage) Deserialize(r Reader) error {

	lbidHouseItemsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.BidHouseItems = make([]ObjectItemGenericQuantityPrice, lbidHouseItemsLen)

	for i := range m.BidHouseItems {

		var lbidHouseItems ObjectItemGenericQuantityPrice

		lbidHouseItems.Deserialize(r)

		m.BidHouseItems[i] = lbidHouseItems

	}

	lmerchantItemsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MerchantItems = make([]ObjectItemGenericQuantityPrice, lmerchantItemsLen)

	for i := range m.MerchantItems {

		var lmerchantItems ObjectItemGenericQuantityPrice

		lmerchantItems.Deserialize(r)

		m.MerchantItems[i] = lmerchantItems

	}

	return nil
}

type IdolsPresetUseResultMessage struct {
	PresetId uint8

	Code uint8

	MissingIdols []uint16
}

func (m *IdolsPresetUseResultMessage) ID() uint16 {
	return 6614
}

func (m *IdolsPresetUseResultMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Code); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.MissingIdols))); err != nil {
		return err
	}

	for i := range m.MissingIdols {

		if err := w.WriteVarUInt16(m.MissingIdols[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *IdolsPresetUseResultMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	lcode, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Code = lcode

	lmissingIdolsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MissingIdols = make([]uint16, lmissingIdolsLen)

	for i := range m.MissingIdols {

		lmissingIdols, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.MissingIdols[i] = lmissingIdols

	}

	return nil
}

type IdolsPresetUseMessage struct {
	PresetId uint8

	Party bool
}

func (m *IdolsPresetUseMessage) ID() uint16 {
	return 6615
}

func (m *IdolsPresetUseMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.PresetId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Party); err != nil {
		return err
	}

	return nil
}

func (m *IdolsPresetUseMessage) Deserialize(r Reader) error {

	lpresetId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.PresetId = lpresetId

	lparty, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Party = lparty

	return nil
}

type SubscriptionUpdateMessage struct {
	Timestamp float64
}

func (m *SubscriptionUpdateMessage) ID() uint16 {
	return 6616
}

func (m *SubscriptionUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionUpdateMessage) Deserialize(r Reader) error {

	ltimestamp, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.Timestamp = ltimestamp

	return nil
}

type GameRefreshMonsterBoostsMessage struct {
	MonsterBoosts []MonsterBoosts

	FamilyBoosts []MonsterBoosts
}

func (m *GameRefreshMonsterBoostsMessage) ID() uint16 {
	return 6618
}

func (m *GameRefreshMonsterBoostsMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.MonsterBoosts))); err != nil {
		return err
	}

	for i := range m.MonsterBoosts {

		if err := m.MonsterBoosts[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.FamilyBoosts))); err != nil {
		return err
	}

	for i := range m.FamilyBoosts {

		if err := m.FamilyBoosts[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameRefreshMonsterBoostsMessage) Deserialize(r Reader) error {

	lmonsterBoostsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MonsterBoosts = make([]MonsterBoosts, lmonsterBoostsLen)

	for i := range m.MonsterBoosts {

		var lmonsterBoosts MonsterBoosts

		lmonsterBoosts.Deserialize(r)

		m.MonsterBoosts[i] = lmonsterBoosts

	}

	lfamilyBoostsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FamilyBoosts = make([]MonsterBoosts, lfamilyBoostsLen)

	for i := range m.FamilyBoosts {

		var lfamilyBoosts MonsterBoosts

		lfamilyBoosts.Deserialize(r)

		m.FamilyBoosts[i] = lfamilyBoosts

	}

	return nil
}

type EditHavenBagCancelRequestMessage struct {
}

func (m *EditHavenBagCancelRequestMessage) ID() uint16 {
	return 6619
}

func (m *EditHavenBagCancelRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *EditHavenBagCancelRequestMessage) Deserialize(r Reader) error {

	return nil
}

type HavenBagPackListMessage struct {
	PackIds []int8
}

func (m *HavenBagPackListMessage) ID() uint16 {
	return 6620
}

func (m *HavenBagPackListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.PackIds))); err != nil {
		return err
	}

	for i := range m.PackIds {

		if err := w.WriteInt8(m.PackIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *HavenBagPackListMessage) Deserialize(r Reader) error {

	lpackIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.PackIds = make([]int8, lpackIdsLen)

	for i := range m.PackIds {

		lpackIds, err := r.ReadInt8()
		if err != nil {
			return err
		}

		m.PackIds[i] = lpackIds

	}

	return nil
}

type CloseHavenBagFurnitureSequenceRequestMessage struct {
}

func (m *CloseHavenBagFurnitureSequenceRequestMessage) ID() uint16 {
	return 6621
}

func (m *CloseHavenBagFurnitureSequenceRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *CloseHavenBagFurnitureSequenceRequestMessage) Deserialize(r Reader) error {

	return nil
}

type MapComplementaryInformationsDataInHavenBagMessage struct {
	MapComplementaryInformationsDataMessage

	OwnerInformations CharacterMinimalInformations

	Theme int8

	RoomId uint8

	MaxRoomId uint8
}

func (m *MapComplementaryInformationsDataInHavenBagMessage) ID() uint16 {
	return 6622
}

func (m *MapComplementaryInformationsDataInHavenBagMessage) Serialize(w Writer) error {

	if err := m.MapComplementaryInformationsDataMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.OwnerInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt8(m.Theme); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.RoomId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.MaxRoomId); err != nil {
		return err
	}

	return nil
}

func (m *MapComplementaryInformationsDataInHavenBagMessage) Deserialize(r Reader) error {

	if err := m.MapComplementaryInformationsDataMessage.Deserialize(r); err != nil {
		return err
	}

	var lownerInformations CharacterMinimalInformations

	lownerInformations.Deserialize(r)

	m.OwnerInformations = lownerInformations

	ltheme, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Theme = ltheme

	lroomId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.RoomId = lroomId

	lmaxRoomId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.MaxRoomId = lmaxRoomId

	return nil
}

type EditHavenBagRequestMessage struct {
}

func (m *EditHavenBagRequestMessage) ID() uint16 {
	return 6626
}

func (m *EditHavenBagRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *EditHavenBagRequestMessage) Deserialize(r Reader) error {

	return nil
}

type EditHavenBagFinishedMessage struct {
}

func (m *EditHavenBagFinishedMessage) ID() uint16 {
	return 6628
}

func (m *EditHavenBagFinishedMessage) Serialize(w Writer) error {

	return nil
}

func (m *EditHavenBagFinishedMessage) Deserialize(r Reader) error {

	return nil
}

type RoomAvailableUpdateMessage struct {
	NbRoom uint8
}

func (m *RoomAvailableUpdateMessage) ID() uint16 {
	return 6630
}

func (m *RoomAvailableUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.NbRoom); err != nil {
		return err
	}

	return nil
}

func (m *RoomAvailableUpdateMessage) Deserialize(r Reader) error {

	lnbRoom, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.NbRoom = lnbRoom

	return nil
}

type ExitHavenBagRequestMessage struct {
}

func (m *ExitHavenBagRequestMessage) ID() uint16 {
	return 6631
}

func (m *ExitHavenBagRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *ExitHavenBagRequestMessage) Deserialize(r Reader) error {

	return nil
}

type EditHavenBagStartMessage struct {
}

func (m *EditHavenBagStartMessage) ID() uint16 {
	return 6632
}

func (m *EditHavenBagStartMessage) Serialize(w Writer) error {

	return nil
}

func (m *EditHavenBagStartMessage) Deserialize(r Reader) error {

	return nil
}

type HavenBagFurnituresMessage struct {
	FurnituresInfos []HavenBagFurnitureInformation
}

func (m *HavenBagFurnituresMessage) ID() uint16 {
	return 6634
}

func (m *HavenBagFurnituresMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.FurnituresInfos))); err != nil {
		return err
	}

	for i := range m.FurnituresInfos {

		if err := m.FurnituresInfos[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *HavenBagFurnituresMessage) Deserialize(r Reader) error {

	lfurnituresInfosLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FurnituresInfos = make([]HavenBagFurnitureInformation, lfurnituresInfosLen)

	for i := range m.FurnituresInfos {

		var lfurnituresInfos HavenBagFurnitureInformation

		lfurnituresInfos.Deserialize(r)

		m.FurnituresInfos[i] = lfurnituresInfos

	}

	return nil
}

type OpenHavenBagFurnitureSequenceRequestMessage struct {
}

func (m *OpenHavenBagFurnitureSequenceRequestMessage) ID() uint16 {
	return 6635
}

func (m *OpenHavenBagFurnitureSequenceRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *OpenHavenBagFurnitureSequenceRequestMessage) Deserialize(r Reader) error {

	return nil
}

type EnterHavenBagRequestMessage struct {
	HavenBagOwner int64
}

func (m *EnterHavenBagRequestMessage) ID() uint16 {
	return 6636
}

func (m *EnterHavenBagRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.HavenBagOwner); err != nil {
		return err
	}

	return nil
}

func (m *EnterHavenBagRequestMessage) Deserialize(r Reader) error {

	lhavenBagOwner, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.HavenBagOwner = lhavenBagOwner

	return nil
}

type HavenBagFurnituresRequestMessage struct {
	CellIds []uint16

	FunitureIds []int32

	Orientations []uint8
}

func (m *HavenBagFurnituresRequestMessage) ID() uint16 {
	return 6637
}

func (m *HavenBagFurnituresRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.CellIds))); err != nil {
		return err
	}

	for i := range m.CellIds {

		if err := w.WriteVarUInt16(m.CellIds[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.FunitureIds))); err != nil {
		return err
	}

	for i := range m.FunitureIds {

		if err := w.WriteInt32(m.FunitureIds[i]); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.Orientations))); err != nil {
		return err
	}

	for i := range m.Orientations {

		if err := w.WriteUInt8(m.Orientations[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *HavenBagFurnituresRequestMessage) Deserialize(r Reader) error {

	lcellIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CellIds = make([]uint16, lcellIdsLen)

	for i := range m.CellIds {

		lcellIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.CellIds[i] = lcellIds

	}

	lfunitureIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FunitureIds = make([]int32, lfunitureIdsLen)

	for i := range m.FunitureIds {

		lfunitureIds, err := r.ReadInt32()
		if err != nil {
			return err
		}

		m.FunitureIds[i] = lfunitureIds

	}

	lorientationsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Orientations = make([]uint8, lorientationsLen)

	for i := range m.Orientations {

		lorientations, err := r.ReadUInt8()
		if err != nil {
			return err
		}

		m.Orientations[i] = lorientations

	}

	return nil
}

type ChangeHavenBagRoomRequestMessage struct {
	RoomId uint8
}

func (m *ChangeHavenBagRoomRequestMessage) ID() uint16 {
	return 6638
}

func (m *ChangeHavenBagRoomRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.RoomId); err != nil {
		return err
	}

	return nil
}

func (m *ChangeHavenBagRoomRequestMessage) Deserialize(r Reader) error {

	lroomId, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.RoomId = lroomId

	return nil
}

type ChangeThemeRequestMessage struct {
	Theme int8
}

func (m *ChangeThemeRequestMessage) ID() uint16 {
	return 6639
}

func (m *ChangeThemeRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt8(m.Theme); err != nil {
		return err
	}

	return nil
}

func (m *ChangeThemeRequestMessage) Deserialize(r Reader) error {

	ltheme, err := r.ReadInt8()
	if err != nil {
		return err
	}

	m.Theme = ltheme

	return nil
}

type InviteInHavenBagMessage struct {
	GuestInformations CharacterMinimalInformations

	Accept bool
}

func (m *InviteInHavenBagMessage) ID() uint16 {
	return 6642
}

func (m *InviteInHavenBagMessage) Serialize(w Writer) error {

	if err := m.GuestInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Accept); err != nil {
		return err
	}

	return nil
}

func (m *InviteInHavenBagMessage) Deserialize(r Reader) error {

	var lguestInformations CharacterMinimalInformations

	lguestInformations.Deserialize(r)

	m.GuestInformations = lguestInformations

	laccept, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Accept = laccept

	return nil
}

type InviteInHavenBagOfferMessage struct {
	HostInformations CharacterMinimalInformations

	TimeLeftBeforeCancel int32
}

func (m *InviteInHavenBagOfferMessage) ID() uint16 {
	return 6643
}

func (m *InviteInHavenBagOfferMessage) Serialize(w Writer) error {

	if err := m.HostInformations.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarInt32(m.TimeLeftBeforeCancel); err != nil {
		return err
	}

	return nil
}

func (m *InviteInHavenBagOfferMessage) Deserialize(r Reader) error {

	var lhostInformations CharacterMinimalInformations

	lhostInformations.Deserialize(r)

	m.HostInformations = lhostInformations

	ltimeLeftBeforeCancel, err := r.ReadVarInt32()
	if err != nil {
		return err
	}

	m.TimeLeftBeforeCancel = ltimeLeftBeforeCancel

	return nil
}

type HavenBagDailyLoteryMessage struct {
	ReturnType uint8

	TokenId string
}

func (m *HavenBagDailyLoteryMessage) ID() uint16 {
	return 6644
}

func (m *HavenBagDailyLoteryMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.ReturnType); err != nil {
		return err
	}

	if err := w.WriteString(m.TokenId); err != nil {
		return err
	}

	return nil
}

func (m *HavenBagDailyLoteryMessage) Deserialize(r Reader) error {

	lreturnType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ReturnType = lreturnType

	ltokenId, err := r.ReadString()
	if err != nil {
		return err
	}

	m.TokenId = ltokenId

	return nil
}

type InviteInHavenBagClosedMessage struct {
	HostInformations CharacterMinimalInformations
}

func (m *InviteInHavenBagClosedMessage) ID() uint16 {
	return 6645
}

func (m *InviteInHavenBagClosedMessage) Serialize(w Writer) error {

	if err := m.HostInformations.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *InviteInHavenBagClosedMessage) Deserialize(r Reader) error {

	var lhostInformations CharacterMinimalInformations

	lhostInformations.Deserialize(r)

	m.HostInformations = lhostInformations

	return nil
}

type TeleportHavenBagAnswerMessage struct {
	Accept bool
}

func (m *TeleportHavenBagAnswerMessage) ID() uint16 {
	return 6646
}

func (m *TeleportHavenBagAnswerMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Accept); err != nil {
		return err
	}

	return nil
}

func (m *TeleportHavenBagAnswerMessage) Deserialize(r Reader) error {

	laccept, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Accept = laccept

	return nil
}

type TeleportHavenBagRequestMessage struct {
	GuestId int64
}

func (m *TeleportHavenBagRequestMessage) ID() uint16 {
	return 6647
}

func (m *TeleportHavenBagRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.GuestId); err != nil {
		return err
	}

	return nil
}

func (m *TeleportHavenBagRequestMessage) Deserialize(r Reader) error {

	lguestId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.GuestId = lguestId

	return nil
}

type HaapiApiKeyRequestMessage struct {
	KeyType uint8
}

func (m *HaapiApiKeyRequestMessage) ID() uint16 {
	return 6648
}

func (m *HaapiApiKeyRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.KeyType); err != nil {
		return err
	}

	return nil
}

func (m *HaapiApiKeyRequestMessage) Deserialize(r Reader) error {

	lkeyType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.KeyType = lkeyType

	return nil
}

type HaapiApiKeyMessage struct {
	ReturnType uint8

	KeyType uint8

	Token string
}

func (m *HaapiApiKeyMessage) ID() uint16 {
	return 6649
}

func (m *HaapiApiKeyMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.ReturnType); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.KeyType); err != nil {
		return err
	}

	if err := w.WriteString(m.Token); err != nil {
		return err
	}

	return nil
}

func (m *HaapiApiKeyMessage) Deserialize(r Reader) error {

	lreturnType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.ReturnType = lreturnType

	lkeyType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.KeyType = lkeyType

	ltoken, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Token = ltoken

	return nil
}

type KickHavenBagRequestMessage struct {
	GuestId int64
}

func (m *KickHavenBagRequestMessage) ID() uint16 {
	return 6652
}

func (m *KickHavenBagRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarInt64(m.GuestId); err != nil {
		return err
	}

	return nil
}

func (m *KickHavenBagRequestMessage) Deserialize(r Reader) error {

	lguestId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.GuestId = lguestId

	return nil
}

type SpellModifyFailureMessage struct {
}

func (m *SpellModifyFailureMessage) ID() uint16 {
	return 6653
}

func (m *SpellModifyFailureMessage) Serialize(w Writer) error {

	return nil
}

func (m *SpellModifyFailureMessage) Deserialize(r Reader) error {

	return nil
}

type SpellModifySuccessMessage struct {
	SpellId int32

	SpellLevel int16
}

func (m *SpellModifySuccessMessage) ID() uint16 {
	return 6654
}

func (m *SpellModifySuccessMessage) Serialize(w Writer) error {

	if err := w.WriteInt32(m.SpellId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.SpellLevel); err != nil {
		return err
	}

	return nil
}

func (m *SpellModifySuccessMessage) Deserialize(r Reader) error {

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

type SpellModifyRequestMessage struct {
	SpellId uint16

	SpellLevel int16
}

func (m *SpellModifyRequestMessage) ID() uint16 {
	return 6655
}

func (m *SpellModifyRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.SpellLevel); err != nil {
		return err
	}

	return nil
}

func (m *SpellModifyRequestMessage) Deserialize(r Reader) error {

	lspellId, err := r.ReadVarUInt16()
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

type DareInformationsMessage struct {
	DareFixedInfos DareInformations

	DareVersatilesInfos DareVersatileInformations
}

func (m *DareInformationsMessage) ID() uint16 {
	return 6656
}

func (m *DareInformationsMessage) Serialize(w Writer) error {

	if err := m.DareFixedInfos.Serialize(w); err != nil {
		return err
	}

	if err := m.DareVersatilesInfos.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *DareInformationsMessage) Deserialize(r Reader) error {

	var ldareFixedInfos DareInformations

	ldareFixedInfos.Deserialize(r)

	m.DareFixedInfos = ldareFixedInfos

	var ldareVersatilesInfos DareVersatileInformations

	ldareVersatilesInfos.Deserialize(r)

	m.DareVersatilesInfos = ldareVersatilesInfos

	return nil
}

type DareVersatileListMessage struct {
	Dares []DareVersatileInformations
}

func (m *DareVersatileListMessage) ID() uint16 {
	return 6657
}

func (m *DareVersatileListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Dares))); err != nil {
		return err
	}

	for i := range m.Dares {

		if err := m.Dares[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *DareVersatileListMessage) Deserialize(r Reader) error {

	ldaresLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Dares = make([]DareVersatileInformations, ldaresLen)

	for i := range m.Dares {

		var ldares DareVersatileInformations

		ldares.Deserialize(r)

		m.Dares[i] = ldares

	}

	return nil
}

type DareSubscribedListMessage struct {
	DaresFixedInfos []DareInformations

	DaresVersatilesInfos []DareVersatileInformations
}

func (m *DareSubscribedListMessage) ID() uint16 {
	return 6658
}

func (m *DareSubscribedListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.DaresFixedInfos))); err != nil {
		return err
	}

	for i := range m.DaresFixedInfos {

		if err := m.DaresFixedInfos[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.DaresVersatilesInfos))); err != nil {
		return err
	}

	for i := range m.DaresVersatilesInfos {

		if err := m.DaresVersatilesInfos[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *DareSubscribedListMessage) Deserialize(r Reader) error {

	ldaresFixedInfosLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DaresFixedInfos = make([]DareInformations, ldaresFixedInfosLen)

	for i := range m.DaresFixedInfos {

		var ldaresFixedInfos DareInformations

		ldaresFixedInfos.Deserialize(r)

		m.DaresFixedInfos[i] = ldaresFixedInfos

	}

	ldaresVersatilesInfosLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DaresVersatilesInfos = make([]DareVersatileInformations, ldaresVersatilesInfosLen)

	for i := range m.DaresVersatilesInfos {

		var ldaresVersatilesInfos DareVersatileInformations

		ldaresVersatilesInfos.Deserialize(r)

		m.DaresVersatilesInfos[i] = ldaresVersatilesInfos

	}

	return nil
}

type DareInformationsRequestMessage struct {
	DareId float64
}

func (m *DareInformationsRequestMessage) ID() uint16 {
	return 6659
}

func (m *DareInformationsRequestMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.DareId); err != nil {
		return err
	}

	return nil
}

func (m *DareInformationsRequestMessage) Deserialize(r Reader) error {

	ldareId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DareId = ldareId

	return nil
}

type DareSubscribedMessage struct {
	DareId float64

	Success bool

	Subscribe bool

	DareVersatilesInfos DareVersatileInformations
}

func (m *DareSubscribedMessage) ID() uint16 {
	return 6660
}

func (m *DareSubscribedMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.Success)

	setWrappedFlag(bbw0, 1, m.Subscribe)

	if err := w.WriteUInt8(bbw0); err != nil {
		return err
	}

	if err := w.WriteDouble(m.DareId); err != nil {
		return err
	}

	if err := m.DareVersatilesInfos.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *DareSubscribedMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Success = getWrappedFlag(bbw0, 0)

	m.Subscribe = getWrappedFlag(bbw0, 1)

	ldareId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DareId = ldareId

	var ldareVersatilesInfos DareVersatileInformations

	ldareVersatilesInfos.Deserialize(r)

	m.DareVersatilesInfos = ldareVersatilesInfos

	return nil
}

type DareListMessage struct {
	Dares []DareInformations
}

func (m *DareListMessage) ID() uint16 {
	return 6661
}

func (m *DareListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Dares))); err != nil {
		return err
	}

	for i := range m.Dares {

		if err := m.Dares[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *DareListMessage) Deserialize(r Reader) error {

	ldaresLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Dares = make([]DareInformations, ldaresLen)

	for i := range m.Dares {

		var ldares DareInformations

		ldares.Deserialize(r)

		m.Dares[i] = ldares

	}

	return nil
}

type AggregateStatWithDataMessage struct {
	AggregateStatMessage

	Datas []StatisticDataIntrf
}

func (m *AggregateStatWithDataMessage) ID() uint16 {
	return 6662
}

func (m *AggregateStatWithDataMessage) Serialize(w Writer) error {

	if err := m.AggregateStatMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt16(int16(len(m.Datas))); err != nil {
		return err
	}

	for i := range m.Datas {

		if err := w.WriteUInt16(m.Datas[i].ID()); err != nil {
			return err
		}

		if err := m.Datas[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *AggregateStatWithDataMessage) Deserialize(r Reader) error {

	if err := m.AggregateStatMessage.Deserialize(r); err != nil {
		return err
	}

	ldatasLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Datas = make([]StatisticDataIntrf, ldatasLen)

	for i := range m.Datas {

		typedatasID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		ldatas, err := GetType(typedatasID)
		if err != nil {
			return err
		}

		ldatas.Deserialize(r)

		m.Datas[i] = ldatas.(StatisticDataIntrf)

	}

	return nil
}

type DareCreatedListMessage struct {
	DaresFixedInfos []DareInformations

	DaresVersatilesInfos []DareVersatileInformations
}

func (m *DareCreatedListMessage) ID() uint16 {
	return 6663
}

func (m *DareCreatedListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.DaresFixedInfos))); err != nil {
		return err
	}

	for i := range m.DaresFixedInfos {

		if err := m.DaresFixedInfos[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteInt16(int16(len(m.DaresVersatilesInfos))); err != nil {
		return err
	}

	for i := range m.DaresVersatilesInfos {

		if err := m.DaresVersatilesInfos[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *DareCreatedListMessage) Deserialize(r Reader) error {

	ldaresFixedInfosLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DaresFixedInfos = make([]DareInformations, ldaresFixedInfosLen)

	for i := range m.DaresFixedInfos {

		var ldaresFixedInfos DareInformations

		ldaresFixedInfos.Deserialize(r)

		m.DaresFixedInfos[i] = ldaresFixedInfos

	}

	ldaresVersatilesInfosLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DaresVersatilesInfos = make([]DareVersatileInformations, ldaresVersatilesInfosLen)

	for i := range m.DaresVersatilesInfos {

		var ldaresVersatilesInfos DareVersatileInformations

		ldaresVersatilesInfos.Deserialize(r)

		m.DaresVersatilesInfos[i] = ldaresVersatilesInfos

	}

	return nil
}

type ExchangeStartedTaxCollectorShopMessage struct {
	Objects []ObjectItem

	Kamas int64
}

func (m *ExchangeStartedTaxCollectorShopMessage) ID() uint16 {
	return 6664
}

func (m *ExchangeStartedTaxCollectorShopMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Objects))); err != nil {
		return err
	}

	for i := range m.Objects {

		if err := m.Objects[i].Serialize(w); err != nil {
			return err
		}

	}

	if err := w.WriteVarInt64(m.Kamas); err != nil {
		return err
	}

	return nil
}

func (m *ExchangeStartedTaxCollectorShopMessage) Deserialize(r Reader) error {

	lobjectsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Objects = make([]ObjectItem, lobjectsLen)

	for i := range m.Objects {

		var lobjects ObjectItem

		lobjects.Deserialize(r)

		m.Objects[i] = lobjects

	}

	lkamas, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.Kamas = lkamas

	return nil
}

type DareCreationRequestMessage struct {
	SubscriptionFee int64

	Jackpot int64

	MaxCountWinners uint16

	DelayBeforeStart uint32

	Duration uint32

	IsPrivate bool

	IsForGuild bool

	IsForAlliance bool

	NeedNotifications bool

	Criterions []DareCriteria
}

func (m *DareCreationRequestMessage) ID() uint16 {
	return 6665
}

func (m *DareCreationRequestMessage) Serialize(w Writer) error {

	var bbw0 uint8

	setWrappedFlag(bbw0, 0, m.IsPrivate)

	setWrappedFlag(bbw0, 1, m.IsForGuild)

	setWrappedFlag(bbw0, 2, m.IsForAlliance)

	setWrappedFlag(bbw0, 3, m.NeedNotifications)

	if err := w.WriteUInt8(bbw0); err != nil {
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

	if err := w.WriteUInt32(m.DelayBeforeStart); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.Duration); err != nil {
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

	return nil
}

func (m *DareCreationRequestMessage) Deserialize(r Reader) error {

	bbw0, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.IsPrivate = getWrappedFlag(bbw0, 0)

	m.IsForGuild = getWrappedFlag(bbw0, 1)

	m.IsForAlliance = getWrappedFlag(bbw0, 2)

	m.NeedNotifications = getWrappedFlag(bbw0, 3)

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

	ldelayBeforeStart, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.DelayBeforeStart = ldelayBeforeStart

	lduration, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.Duration = lduration

	lcriterionsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Criterions = make([]DareCriteria, lcriterionsLen)

	for i := range m.Criterions {

		var lcriterions DareCriteria

		lcriterions.Deserialize(r)

		m.Criterions[i] = lcriterions

	}

	return nil
}

type DareSubscribeRequestMessage struct {
	DareId float64

	Subscribe bool
}

func (m *DareSubscribeRequestMessage) ID() uint16 {
	return 6666
}

func (m *DareSubscribeRequestMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.DareId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.Subscribe); err != nil {
		return err
	}

	return nil
}

func (m *DareSubscribeRequestMessage) Deserialize(r Reader) error {

	ldareId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DareId = ldareId

	lsubscribe, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Subscribe = lsubscribe

	return nil
}

type DareErrorMessage struct {
	Error uint8
}

func (m *DareErrorMessage) ID() uint16 {
	return 6667
}

func (m *DareErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Error); err != nil {
		return err
	}

	return nil
}

func (m *DareErrorMessage) Deserialize(r Reader) error {

	lerror, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Error = lerror

	return nil
}

type DareCreatedMessage struct {
	DareInfos DareInformations

	NeedNotifications bool
}

func (m *DareCreatedMessage) ID() uint16 {
	return 6668
}

func (m *DareCreatedMessage) Serialize(w Writer) error {

	if err := m.DareInfos.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.NeedNotifications); err != nil {
		return err
	}

	return nil
}

func (m *DareCreatedMessage) Deserialize(r Reader) error {

	var ldareInfos DareInformations

	ldareInfos.Deserialize(r)

	m.DareInfos = ldareInfos

	lneedNotifications, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.NeedNotifications = lneedNotifications

	return nil
}

type AggregateStatMessage struct {
	StatId uint16
}

func (m *AggregateStatMessage) ID() uint16 {
	return 6669
}

func (m *AggregateStatMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.StatId); err != nil {
		return err
	}

	return nil
}

func (m *AggregateStatMessage) Deserialize(r Reader) error {

	lstatId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.StatId = lstatId

	return nil
}

type ExchangePodsModifiedMessage struct {
	ExchangeObjectMessage

	CurrentWeight uint32

	MaxWeight uint32
}

func (m *ExchangePodsModifiedMessage) ID() uint16 {
	return 6670
}

func (m *ExchangePodsModifiedMessage) Serialize(w Writer) error {

	if err := m.ExchangeObjectMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.CurrentWeight); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.MaxWeight); err != nil {
		return err
	}

	return nil
}

func (m *ExchangePodsModifiedMessage) Deserialize(r Reader) error {

	if err := m.ExchangeObjectMessage.Deserialize(r); err != nil {
		return err
	}

	lcurrentWeight, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.CurrentWeight = lcurrentWeight

	lmaxWeight, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.MaxWeight = lmaxWeight

	return nil
}

type DareRewardConsumeValidationMessage struct {
	DareId float64

	Type uint8
}

func (m *DareRewardConsumeValidationMessage) ID() uint16 {
	return 6675
}

func (m *DareRewardConsumeValidationMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.DareId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DareRewardConsumeValidationMessage) Deserialize(r Reader) error {

	ldareId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DareId = ldareId

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	return nil
}

type DareRewardConsumeRequestMessage struct {
	DareId float64

	Type uint8
}

func (m *DareRewardConsumeRequestMessage) ID() uint16 {
	return 6676
}

func (m *DareRewardConsumeRequestMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.DareId); err != nil {
		return err
	}

	if err := w.WriteUInt8(m.Type); err != nil {
		return err
	}

	return nil
}

func (m *DareRewardConsumeRequestMessage) Deserialize(r Reader) error {

	ldareId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DareId = ldareId

	ltype, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Type = ltype

	return nil
}

type DareRewardsListMessage struct {
	Rewards []DareReward
}

func (m *DareRewardsListMessage) ID() uint16 {
	return 6677
}

func (m *DareRewardsListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Rewards))); err != nil {
		return err
	}

	for i := range m.Rewards {

		if err := m.Rewards[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *DareRewardsListMessage) Deserialize(r Reader) error {

	lrewardsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Rewards = make([]DareReward, lrewardsLen)

	for i := range m.Rewards {

		var lrewards DareReward

		lrewards.Deserialize(r)

		m.Rewards[i] = lrewards

	}

	return nil
}

type DareRewardWonMessage struct {
	Reward DareReward
}

func (m *DareRewardWonMessage) ID() uint16 {
	return 6678
}

func (m *DareRewardWonMessage) Serialize(w Writer) error {

	if err := m.Reward.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *DareRewardWonMessage) Deserialize(r Reader) error {

	var lreward DareReward

	lreward.Deserialize(r)

	m.Reward = lreward

	return nil
}

type DareCanceledMessage struct {
	DareId float64
}

func (m *DareCanceledMessage) ID() uint16 {
	return 6679
}

func (m *DareCanceledMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.DareId); err != nil {
		return err
	}

	return nil
}

func (m *DareCanceledMessage) Deserialize(r Reader) error {

	ldareId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DareId = ldareId

	return nil
}

type DareCancelRequestMessage struct {
	DareId float64
}

func (m *DareCancelRequestMessage) ID() uint16 {
	return 6680
}

func (m *DareCancelRequestMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.DareId); err != nil {
		return err
	}

	return nil
}

func (m *DareCancelRequestMessage) Deserialize(r Reader) error {

	ldareId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DareId = ldareId

	return nil
}

type DareWonMessage struct {
	DareId float64
}

func (m *DareWonMessage) ID() uint16 {
	return 6681
}

func (m *DareWonMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.DareId); err != nil {
		return err
	}

	return nil
}

func (m *DareWonMessage) Deserialize(r Reader) error {

	ldareId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.DareId = ldareId

	return nil
}

type DareWonListMessage struct {
	DareId []float64
}

func (m *DareWonListMessage) ID() uint16 {
	return 6682
}

func (m *DareWonListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.DareId))); err != nil {
		return err
	}

	for i := range m.DareId {

		if err := w.WriteDouble(m.DareId[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *DareWonListMessage) Deserialize(r Reader) error {

	ldareIdLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.DareId = make([]float64, ldareIdLen)

	for i := range m.DareId {

		ldareId, err := r.ReadDouble()
		if err != nil {
			return err
		}

		m.DareId[i] = ldareId

	}

	return nil
}

type AllianceMotdSetErrorMessage struct {
	SocialNoticeSetErrorMessage
}

func (m *AllianceMotdSetErrorMessage) ID() uint16 {
	return 6683
}

func (m *AllianceMotdSetErrorMessage) Serialize(w Writer) error {

	if err := m.SocialNoticeSetErrorMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AllianceMotdSetErrorMessage) Deserialize(r Reader) error {

	if err := m.SocialNoticeSetErrorMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type SocialNoticeSetErrorMessage struct {
	Reason uint8
}

func (m *SocialNoticeSetErrorMessage) ID() uint16 {
	return 6684
}

func (m *SocialNoticeSetErrorMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.Reason); err != nil {
		return err
	}

	return nil
}

func (m *SocialNoticeSetErrorMessage) Deserialize(r Reader) error {

	lreason, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.Reason = lreason

	return nil
}

type AllianceMotdMessage struct {
	SocialNoticeMessage
}

func (m *AllianceMotdMessage) ID() uint16 {
	return 6685
}

func (m *AllianceMotdMessage) Serialize(w Writer) error {

	if err := m.SocialNoticeMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AllianceMotdMessage) Deserialize(r Reader) error {

	if err := m.SocialNoticeMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type SocialNoticeSetRequestMessage struct {
}

func (m *SocialNoticeSetRequestMessage) ID() uint16 {
	return 6686
}

func (m *SocialNoticeSetRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *SocialNoticeSetRequestMessage) Deserialize(r Reader) error {

	return nil
}

type AllianceMotdSetRequestMessage struct {
	SocialNoticeSetRequestMessage

	Content string
}

func (m *AllianceMotdSetRequestMessage) ID() uint16 {
	return 6687
}

func (m *AllianceMotdSetRequestMessage) Serialize(w Writer) error {

	if err := m.SocialNoticeSetRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Content); err != nil {
		return err
	}

	return nil
}

func (m *AllianceMotdSetRequestMessage) Deserialize(r Reader) error {

	if err := m.SocialNoticeSetRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lcontent, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Content = lcontent

	return nil
}

type SocialNoticeMessage struct {
	Content string

	Timestamp uint32

	MemberId int64

	MemberName string
}

func (m *SocialNoticeMessage) ID() uint16 {
	return 6688
}

func (m *SocialNoticeMessage) Serialize(w Writer) error {

	if err := w.WriteString(m.Content); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.Timestamp); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.MemberId); err != nil {
		return err
	}

	if err := w.WriteString(m.MemberName); err != nil {
		return err
	}

	return nil
}

func (m *SocialNoticeMessage) Deserialize(r Reader) error {

	lcontent, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Content = lcontent

	ltimestamp, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.Timestamp = ltimestamp

	lmemberId, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.MemberId = lmemberId

	lmemberName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.MemberName = lmemberName

	return nil
}

type GuildBulletinMessage struct {
	BulletinMessage
}

func (m *GuildBulletinMessage) ID() uint16 {
	return 6689
}

func (m *GuildBulletinMessage) Serialize(w Writer) error {

	if err := m.BulletinMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildBulletinMessage) Deserialize(r Reader) error {

	if err := m.BulletinMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type AllianceBulletinMessage struct {
	BulletinMessage
}

func (m *AllianceBulletinMessage) ID() uint16 {
	return 6690
}

func (m *AllianceBulletinMessage) Serialize(w Writer) error {

	if err := m.BulletinMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AllianceBulletinMessage) Deserialize(r Reader) error {

	if err := m.BulletinMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type GuildBulletinSetErrorMessage struct {
	SocialNoticeSetErrorMessage
}

func (m *GuildBulletinSetErrorMessage) ID() uint16 {
	return 6691
}

func (m *GuildBulletinSetErrorMessage) Serialize(w Writer) error {

	if err := m.SocialNoticeSetErrorMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GuildBulletinSetErrorMessage) Deserialize(r Reader) error {

	if err := m.SocialNoticeSetErrorMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type AllianceBulletinSetErrorMessage struct {
	SocialNoticeSetErrorMessage
}

func (m *AllianceBulletinSetErrorMessage) ID() uint16 {
	return 6692
}

func (m *AllianceBulletinSetErrorMessage) Serialize(w Writer) error {

	if err := m.SocialNoticeSetErrorMessage.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *AllianceBulletinSetErrorMessage) Deserialize(r Reader) error {

	if err := m.SocialNoticeSetErrorMessage.Deserialize(r); err != nil {
		return err
	}

	return nil
}

type AllianceBulletinSetRequestMessage struct {
	SocialNoticeSetRequestMessage

	Content string

	NotifyMembers bool
}

func (m *AllianceBulletinSetRequestMessage) ID() uint16 {
	return 6693
}

func (m *AllianceBulletinSetRequestMessage) Serialize(w Writer) error {

	if err := m.SocialNoticeSetRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Content); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.NotifyMembers); err != nil {
		return err
	}

	return nil
}

func (m *AllianceBulletinSetRequestMessage) Deserialize(r Reader) error {

	if err := m.SocialNoticeSetRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lcontent, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Content = lcontent

	lnotifyMembers, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.NotifyMembers = lnotifyMembers

	return nil
}

type GuildBulletinSetRequestMessage struct {
	SocialNoticeSetRequestMessage

	Content string

	NotifyMembers bool
}

func (m *GuildBulletinSetRequestMessage) ID() uint16 {
	return 6694
}

func (m *GuildBulletinSetRequestMessage) Serialize(w Writer) error {

	if err := m.SocialNoticeSetRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteString(m.Content); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.NotifyMembers); err != nil {
		return err
	}

	return nil
}

func (m *GuildBulletinSetRequestMessage) Deserialize(r Reader) error {

	if err := m.SocialNoticeSetRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lcontent, err := r.ReadString()
	if err != nil {
		return err
	}

	m.Content = lcontent

	lnotifyMembers, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.NotifyMembers = lnotifyMembers

	return nil
}

type BulletinMessage struct {
	SocialNoticeMessage

	LastNotifiedTimestamp uint32
}

func (m *BulletinMessage) ID() uint16 {
	return 6695
}

func (m *BulletinMessage) Serialize(w Writer) error {

	if err := m.SocialNoticeMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.LastNotifiedTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *BulletinMessage) Deserialize(r Reader) error {

	if err := m.SocialNoticeMessage.Deserialize(r); err != nil {
		return err
	}

	llastNotifiedTimestamp, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.LastNotifiedTimestamp = llastNotifiedTimestamp

	return nil
}

type MountHarnessDissociateRequestMessage struct {
}

func (m *MountHarnessDissociateRequestMessage) ID() uint16 {
	return 6696
}

func (m *MountHarnessDissociateRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *MountHarnessDissociateRequestMessage) Deserialize(r Reader) error {

	return nil
}

type MountHarnessColorsUpdateRequestMessage struct {
	UseHarnessColors bool
}

func (m *MountHarnessColorsUpdateRequestMessage) ID() uint16 {
	return 6697
}

func (m *MountHarnessColorsUpdateRequestMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.UseHarnessColors); err != nil {
		return err
	}

	return nil
}

func (m *MountHarnessColorsUpdateRequestMessage) Deserialize(r Reader) error {

	luseHarnessColors, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.UseHarnessColors = luseHarnessColors

	return nil
}

type RefreshCharacterStatsMessage struct {
	FighterId float64

	Stats GameFightMinimalStats
}

func (m *RefreshCharacterStatsMessage) ID() uint16 {
	return 6699
}

func (m *RefreshCharacterStatsMessage) Serialize(w Writer) error {

	if err := w.WriteDouble(m.FighterId); err != nil {
		return err
	}

	if err := m.Stats.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *RefreshCharacterStatsMessage) Deserialize(r Reader) error {

	lfighterId, err := r.ReadDouble()
	if err != nil {
		return err
	}

	m.FighterId = lfighterId

	var lstats GameFightMinimalStats

	lstats.Deserialize(r)

	m.Stats = lstats

	return nil
}

type ArenaFighterLeaveMessage struct {
	Leaver CharacterBasicMinimalInformations
}

func (m *ArenaFighterLeaveMessage) ID() uint16 {
	return 6700
}

func (m *ArenaFighterLeaveMessage) Serialize(w Writer) error {

	if err := m.Leaver.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ArenaFighterLeaveMessage) Deserialize(r Reader) error {

	var lleaver CharacterBasicMinimalInformations

	lleaver.Deserialize(r)

	m.Leaver = lleaver

	return nil
}

type FocusedExchangeReadyMessage struct {
	ExchangeReadyMessage

	FocusActionId uint32
}

func (m *FocusedExchangeReadyMessage) ID() uint16 {
	return 6701
}

func (m *FocusedExchangeReadyMessage) Serialize(w Writer) error {

	if err := m.ExchangeReadyMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteVarUInt32(m.FocusActionId); err != nil {
		return err
	}

	return nil
}

func (m *FocusedExchangeReadyMessage) Deserialize(r Reader) error {

	if err := m.ExchangeReadyMessage.Deserialize(r); err != nil {
		return err
	}

	lfocusActionId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.FocusActionId = lfocusActionId

	return nil
}

type FinishMoveListRequestMessage struct {
}

func (m *FinishMoveListRequestMessage) ID() uint16 {
	return 6702
}

func (m *FinishMoveListRequestMessage) Serialize(w Writer) error {

	return nil
}

func (m *FinishMoveListRequestMessage) Deserialize(r Reader) error {

	return nil
}

type FinishMoveSetRequestMessage struct {
	FinishMoveId uint32

	FinishMoveState bool
}

func (m *FinishMoveSetRequestMessage) ID() uint16 {
	return 6703
}

func (m *FinishMoveSetRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.FinishMoveId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.FinishMoveState); err != nil {
		return err
	}

	return nil
}

func (m *FinishMoveSetRequestMessage) Deserialize(r Reader) error {

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

type FinishMoveListMessage struct {
	FinishMoves []FinishMoveInformations
}

func (m *FinishMoveListMessage) ID() uint16 {
	return 6704
}

func (m *FinishMoveListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.FinishMoves))); err != nil {
		return err
	}

	for i := range m.FinishMoves {

		if err := m.FinishMoves[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *FinishMoveListMessage) Deserialize(r Reader) error {

	lfinishMovesLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.FinishMoves = make([]FinishMoveInformations, lfinishMovesLen)

	for i := range m.FinishMoves {

		var lfinishMoves FinishMoveInformations

		lfinishMoves.Deserialize(r)

		m.FinishMoves[i] = lfinishMoves

	}

	return nil
}

type SpellVariantActivationMessage struct {
	Result bool

	ActivatedSpellId uint16

	DeactivatedSpellId uint16
}

func (m *SpellVariantActivationMessage) ID() uint16 {
	return 6705
}

func (m *SpellVariantActivationMessage) Serialize(w Writer) error {

	if err := w.WriteBoolean(m.Result); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.ActivatedSpellId); err != nil {
		return err
	}

	if err := w.WriteVarUInt16(m.DeactivatedSpellId); err != nil {
		return err
	}

	return nil
}

func (m *SpellVariantActivationMessage) Deserialize(r Reader) error {

	lresult, err := r.ReadBoolean()
	if err != nil {
		return err
	}

	m.Result = lresult

	lactivatedSpellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.ActivatedSpellId = lactivatedSpellId

	ldeactivatedSpellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.DeactivatedSpellId = ldeactivatedSpellId

	return nil
}

type ShortcutBarReplacedMessage struct {
	BarType uint8

	Shortcut ShortcutIntrf
}

func (m *ShortcutBarReplacedMessage) ID() uint16 {
	return 6706
}

func (m *ShortcutBarReplacedMessage) Serialize(w Writer) error {

	if err := w.WriteUInt8(m.BarType); err != nil {
		return err
	}

	if err := w.WriteUInt16(m.Shortcut.ID()); err != nil {
		return err
	}

	if err := m.Shortcut.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *ShortcutBarReplacedMessage) Deserialize(r Reader) error {

	lbarType, err := r.ReadUInt8()
	if err != nil {
		return err
	}

	m.BarType = lbarType

	typeshortcutID, err := r.ReadUInt16()
	if err != nil {
		return err
	}
	lshortcut, err := GetType(typeshortcutID)
	if err != nil {
		return err
	}

	lshortcut.Deserialize(r)

	m.Shortcut = lshortcut.(ShortcutIntrf)

	return nil
}

type SpellVariantActivationRequestMessage struct {
	SpellId uint16
}

func (m *SpellVariantActivationRequestMessage) ID() uint16 {
	return 6707
}

func (m *SpellVariantActivationRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.SpellId); err != nil {
		return err
	}

	return nil
}

func (m *SpellVariantActivationRequestMessage) Deserialize(r Reader) error {

	lspellId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.SpellId = lspellId

	return nil
}

type GameRolePlayShowMultipleActorsMessage struct {
	InformationsList []GameRolePlayActorInformationsIntrf
}

func (m *GameRolePlayShowMultipleActorsMessage) ID() uint16 {
	return 6712
}

func (m *GameRolePlayShowMultipleActorsMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.InformationsList))); err != nil {
		return err
	}

	for i := range m.InformationsList {

		if err := w.WriteUInt16(m.InformationsList[i].ID()); err != nil {
			return err
		}

		if err := m.InformationsList[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *GameRolePlayShowMultipleActorsMessage) Deserialize(r Reader) error {

	linformationsListLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.InformationsList = make([]GameRolePlayActorInformationsIntrf, linformationsListLen)

	for i := range m.InformationsList {

		typeinformationsListID, err := r.ReadUInt16()
		if err != nil {
			return err
		}
		linformationsList, err := GetType(typeinformationsListID)
		if err != nil {
			return err
		}

		linformationsList.Deserialize(r)

		m.InformationsList[i] = linformationsList.(GameRolePlayActorInformationsIntrf)

	}

	return nil
}

type HavenBagPermissionsUpdateMessage struct {
	Permissions uint32
}

func (m *HavenBagPermissionsUpdateMessage) ID() uint16 {
	return 6713
}

func (m *HavenBagPermissionsUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.Permissions); err != nil {
		return err
	}

	return nil
}

func (m *HavenBagPermissionsUpdateMessage) Deserialize(r Reader) error {

	lpermissions, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.Permissions = lpermissions

	return nil
}

type HavenBagPermissionsUpdateRequestMessage struct {
	Permissions uint32
}

func (m *HavenBagPermissionsUpdateRequestMessage) ID() uint16 {
	return 6714
}

func (m *HavenBagPermissionsUpdateRequestMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.Permissions); err != nil {
		return err
	}

	return nil
}

func (m *HavenBagPermissionsUpdateRequestMessage) Deserialize(r Reader) error {

	lpermissions, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.Permissions = lpermissions

	return nil
}

type InteractiveUseWithParamRequestMessage struct {
	InteractiveUseRequestMessage

	Id int32
}

func (m *InteractiveUseWithParamRequestMessage) ID() uint16 {
	return 6715
}

func (m *InteractiveUseWithParamRequestMessage) Serialize(w Writer) error {

	if err := m.InteractiveUseRequestMessage.Serialize(w); err != nil {
		return err
	}

	if err := w.WriteInt32(m.Id); err != nil {
		return err
	}

	return nil
}

func (m *InteractiveUseWithParamRequestMessage) Deserialize(r Reader) error {

	if err := m.InteractiveUseRequestMessage.Deserialize(r); err != nil {
		return err
	}

	lid, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.Id = lid

	return nil
}

type MapFightStartPositionsUpdateMessage struct {
	MapId uint32

	FightStartPositions FightStartingPositions
}

func (m *MapFightStartPositionsUpdateMessage) ID() uint16 {
	return 6716
}

func (m *MapFightStartPositionsUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteUInt32(m.MapId); err != nil {
		return err
	}

	if err := m.FightStartPositions.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *MapFightStartPositionsUpdateMessage) Deserialize(r Reader) error {

	lmapId, err := r.ReadUInt32()
	if err != nil {
		return err
	}

	m.MapId = lmapId

	var lfightStartPositions FightStartingPositions

	lfightStartPositions.Deserialize(r)

	m.FightStartPositions = lfightStartPositions

	return nil
}

type FollowedQuestsMessage struct {
	Quests []QuestActiveDetailedInformations
}

func (m *FollowedQuestsMessage) ID() uint16 {
	return 6717
}

func (m *FollowedQuestsMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Quests))); err != nil {
		return err
	}

	for i := range m.Quests {

		if err := m.Quests[i].Serialize(w); err != nil {
			return err
		}

	}

	return nil
}

func (m *FollowedQuestsMessage) Deserialize(r Reader) error {

	lquestsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Quests = make([]QuestActiveDetailedInformations, lquestsLen)

	for i := range m.Quests {

		var lquests QuestActiveDetailedInformations

		lquests.Deserialize(r)

		m.Quests[i] = lquests

	}

	return nil
}

type RefreshFollowedQuestsOrderRequestMessage struct {
	Quests []uint16
}

func (m *RefreshFollowedQuestsOrderRequestMessage) ID() uint16 {
	return 6722
}

func (m *RefreshFollowedQuestsOrderRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.Quests))); err != nil {
		return err
	}

	for i := range m.Quests {

		if err := w.WriteVarUInt16(m.Quests[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *RefreshFollowedQuestsOrderRequestMessage) Deserialize(r Reader) error {

	lquestsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.Quests = make([]uint16, lquestsLen)

	for i := range m.Quests {

		lquests, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.Quests[i] = lquests

	}

	return nil
}

type UnfollowQuestObjectiveRequestMessage struct {
	QuestId uint16

	ObjectiveId int16
}

func (m *UnfollowQuestObjectiveRequestMessage) ID() uint16 {
	return 6723
}

func (m *UnfollowQuestObjectiveRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.QuestId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.ObjectiveId); err != nil {
		return err
	}

	return nil
}

func (m *UnfollowQuestObjectiveRequestMessage) Deserialize(r Reader) error {

	lquestId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.QuestId = lquestId

	lobjectiveId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectiveId = lobjectiveId

	return nil
}

type FollowQuestObjectiveRequestMessage struct {
	QuestId uint16

	ObjectiveId int16
}

func (m *FollowQuestObjectiveRequestMessage) ID() uint16 {
	return 6724
}

func (m *FollowQuestObjectiveRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt16(m.QuestId); err != nil {
		return err
	}

	if err := w.WriteInt16(m.ObjectiveId); err != nil {
		return err
	}

	return nil
}

func (m *FollowQuestObjectiveRequestMessage) Deserialize(r Reader) error {

	lquestId, err := r.ReadVarUInt16()
	if err != nil {
		return err
	}

	m.QuestId = lquestId

	lobjectiveId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.ObjectiveId = lobjectiveId

	return nil
}

type HouseTeleportRequestMessage struct {
	HouseId uint32

	HouseInstanceId int32
}

func (m *HouseTeleportRequestMessage) ID() uint16 {
	return 6726
}

func (m *HouseTeleportRequestMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HouseId); err != nil {
		return err
	}

	if err := w.WriteInt32(m.HouseInstanceId); err != nil {
		return err
	}

	return nil
}

func (m *HouseTeleportRequestMessage) Deserialize(r Reader) error {

	lhouseId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HouseId = lhouseId

	lhouseInstanceId, err := r.ReadInt32()
	if err != nil {
		return err
	}

	m.HouseInstanceId = lhouseInstanceId

	return nil
}

type HouseSellingUpdateMessage struct {
	HouseId uint32

	InstanceId uint32

	SecondHand bool

	RealPrice int64

	BuyerName string
}

func (m *HouseSellingUpdateMessage) ID() uint16 {
	return 6727
}

func (m *HouseSellingUpdateMessage) Serialize(w Writer) error {

	if err := w.WriteVarUInt32(m.HouseId); err != nil {
		return err
	}

	if err := w.WriteUInt32(m.InstanceId); err != nil {
		return err
	}

	if err := w.WriteBoolean(m.SecondHand); err != nil {
		return err
	}

	if err := w.WriteVarInt64(m.RealPrice); err != nil {
		return err
	}

	if err := w.WriteString(m.BuyerName); err != nil {
		return err
	}

	return nil
}

func (m *HouseSellingUpdateMessage) Deserialize(r Reader) error {

	lhouseId, err := r.ReadVarUInt32()
	if err != nil {
		return err
	}

	m.HouseId = lhouseId

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

	lrealPrice, err := r.ReadVarInt64()
	if err != nil {
		return err
	}

	m.RealPrice = lrealPrice

	lbuyerName, err := r.ReadString()
	if err != nil {
		return err
	}

	m.BuyerName = lbuyerName

	return nil
}

type GameRolePlayArenaUpdatePlayerInfosAllQueuesMessage struct {
	GameRolePlayArenaUpdatePlayerInfosMessage

	Team ArenaRankInfos

	Duel ArenaRankInfos
}

func (m *GameRolePlayArenaUpdatePlayerInfosAllQueuesMessage) ID() uint16 {
	return 6728
}

func (m *GameRolePlayArenaUpdatePlayerInfosAllQueuesMessage) Serialize(w Writer) error {

	if err := m.GameRolePlayArenaUpdatePlayerInfosMessage.Serialize(w); err != nil {
		return err
	}

	if err := m.Team.Serialize(w); err != nil {
		return err
	}

	if err := m.Duel.Serialize(w); err != nil {
		return err
	}

	return nil
}

func (m *GameRolePlayArenaUpdatePlayerInfosAllQueuesMessage) Deserialize(r Reader) error {

	if err := m.GameRolePlayArenaUpdatePlayerInfosMessage.Deserialize(r); err != nil {
		return err
	}

	var lteam ArenaRankInfos

	lteam.Deserialize(r)

	m.Team = lteam

	var lduel ArenaRankInfos

	lduel.Deserialize(r)

	m.Duel = lduel

	return nil
}

type ChatCommunityChannelSetCommunityRequestMessage struct {
	CommunityId int16
}

func (m *ChatCommunityChannelSetCommunityRequestMessage) ID() uint16 {
	return 6729
}

func (m *ChatCommunityChannelSetCommunityRequestMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(m.CommunityId); err != nil {
		return err
	}

	return nil
}

func (m *ChatCommunityChannelSetCommunityRequestMessage) Deserialize(r Reader) error {

	lcommunityId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CommunityId = lcommunityId

	return nil
}

type ChatCommunityChannelCommunityMessage struct {
	CommunityId int16
}

func (m *ChatCommunityChannelCommunityMessage) ID() uint16 {
	return 6730
}

func (m *ChatCommunityChannelCommunityMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(m.CommunityId); err != nil {
		return err
	}

	return nil
}

func (m *ChatCommunityChannelCommunityMessage) Deserialize(r Reader) error {

	lcommunityId, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.CommunityId = lcommunityId

	return nil
}

type MigratedServerListMessage struct {
	MigratedServerIds []uint16
}

func (m *MigratedServerListMessage) ID() uint16 {
	return 6731
}

func (m *MigratedServerListMessage) Serialize(w Writer) error {

	if err := w.WriteInt16(int16(len(m.MigratedServerIds))); err != nil {
		return err
	}

	for i := range m.MigratedServerIds {

		if err := w.WriteVarUInt16(m.MigratedServerIds[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *MigratedServerListMessage) Deserialize(r Reader) error {

	lmigratedServerIdsLen, err := r.ReadInt16()
	if err != nil {
		return err
	}

	m.MigratedServerIds = make([]uint16, lmigratedServerIdsLen)

	for i := range m.MigratedServerIds {

		lmigratedServerIds, err := r.ReadVarUInt16()
		if err != nil {
			return err
		}

		m.MigratedServerIds[i] = lmigratedServerIds

	}

	return nil
}
