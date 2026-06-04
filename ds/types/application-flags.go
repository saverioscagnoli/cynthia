package ds

type ApplicationFlags int

const (
	ApplicationFlagApplicationAutoModerationRuleCreateBadge ApplicationFlags = 1 << 6
	ApplicationFlagGatewayPresence                          ApplicationFlags = 1 << 12
	ApplicationFlagGatewayPresenceLimited                   ApplicationFlags = 1 << 13
	ApplicationFlagGatewayGuildMembers                      ApplicationFlags = 1 << 14
	ApplicationFlagGatewayGuildMembersLimited               ApplicationFlags = 1 << 15
	ApplicationFlagVerificationPendingGuildLimit            ApplicationFlags = 1 << 16
	ApplicationFlagEmbedded                                 ApplicationFlags = 1 << 17
	ApplicationFlagGatewayMessageContent                    ApplicationFlags = 1 << 18
	ApplicationFlagGatewayMessageContentLimited             ApplicationFlags = 1 << 19
	ApplicationFlagApplicationCommandBadge                  ApplicationFlags = 1 << 23
)
