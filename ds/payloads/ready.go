package payloads

import ds "cynthia/ds/types"

type Ready struct {
	Version int     `json:"v"`
	User    ds.User `json:"user"`
}
