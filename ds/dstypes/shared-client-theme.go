package dstypes

import "cynthia/util"

type BaseThemeType int

const (
	BaseThemeTypeUnset    BaseThemeType = 0
	BaseThemeTypeDark     BaseThemeType = 1
	BaseThemeTypeLight    BaseThemeType = 2
	BaseThemeTypeDarker   BaseThemeType = 3
	BaseThemeTypeMidnight BaseThemeType = 4
)

func (t *BaseThemeType) UnmarshalJSON(data []byte) error {
	return util.UnmarshalNumeric(data, t)
}

type SharedClientTheme struct {
	Colors        []string       `json:"colors"`
	GradientAngle int            `json:"grandient_angle"`
	BaseMix       int            `json:"base_mix"`
	BaseTheme     *BaseThemeType `json:"base_theme"`
}
