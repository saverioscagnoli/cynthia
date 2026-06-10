package dstypes

import "cynthia/util"

type StatusDisplayType int

const (
	StatusDisplayTypeName    StatusDisplayType = 0
	StatusDisplayTypeState   StatusDisplayType = 1
	StatusDisplayTypeDetails StatusDisplayType = 2
)

func (t *StatusDisplayType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}
