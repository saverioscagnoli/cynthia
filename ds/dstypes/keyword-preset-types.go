package dstypes

import "cynthia/util"

type KeywordPresetTypes int

const (
	KeywordPresetTypesProfanity     KeywordPresetTypes = 1
	KeywordPresetTypesSexualContent KeywordPresetTypes = 2
	KeywordPresetTypesSlurs         KeywordPresetTypes = 3
)

func (k *KeywordPresetTypes) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, k)
}
