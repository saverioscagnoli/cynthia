package payloads

import "cynthia/ds/dstypes"

type GuildMembersChunk struct {
	GuildID    dstypes.Snowflake     `json:"guild_id"`
	Members    []dstypes.GuildMember `json:"members"`
	ChunkIndex int                   `json:"chunk_index"`
	ChunkCount int                   `json:"chunk_count"`
	NotFound   *[]any                `json:"not_found"`
	Presences  *[]dstypes.Presence   `json:"presences"`
	Nonce      *string               `json:"nonce"`
}
