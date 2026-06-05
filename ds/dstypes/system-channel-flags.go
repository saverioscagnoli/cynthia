package dstypes

type SystemChannelFlags int

const (
	SystemChannelFlagSuppressJoinNotifications                                            SystemChannelFlags = 1 << 0
	SystemChannelFlagSuppressPremiumSubscriptions                                         SystemChannelFlags = 1 << 1
	SystemChannelFlagSuppressGuildReminderNotifications                                   SystemChannelFlags = 1 << 2
	SystemChannelFlagSuppressJoinNotificationReplies                                      SystemChannelFlags = 1 << 3
	SystemChannelFlagSuppressRoleSubscriptionPurchaseNotifications                        SystemChannelFlags = 1 << 4
	SystemChannelFlagSuppressRoleSubscriptionPurchaseNotificationRepliesSystemChannelFlag SystemChannelFlags = 1 << 5
)
