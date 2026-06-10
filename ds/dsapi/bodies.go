package dsapi

import "cynthia/ds/dstypes"

type CreateGuildCommandBody struct {
	Name                     string                              `json:"name"`
	NameLocalizations        map[string]string                   `json:"name_localizations"`
	Description              string                              `json:"description"`
	DescriptionLocalizations map[string]string                   `json:"description_localizations"`
	Options                  *[]dstypes.ApplicationCommandOption `json:"options"`
	DefaultMemberPermissions *string                             `json:"default_member_permissions"`
	DefaultPermissions       *bool                               `json:"default_permissions"`
	Type                     *dstypes.ApplicationCommandType     `json:"type"`
	NSFW                     *bool                               `json:"nsfw"`
}
