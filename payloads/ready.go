package payloads

import "cynthia/dstypes"

type Ready struct {
	Version int          `json:"v"`
	User    dstypes.User `json:"user"`
}
