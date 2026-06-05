package payloads

import "cynthia/ds/dstypes"

type Ready struct {
	Version int          `json:"v"`
	User    dstypes.User `json:"user"`
}
