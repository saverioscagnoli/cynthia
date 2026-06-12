package ds

import "fmt"

func (r *routes) AvatarURL(userID Snowflake, avatarHash string) string {
	return fmt.Sprintf("%s/avatars/%s/%s.png", CdnURL, userID, avatarHash)
}
