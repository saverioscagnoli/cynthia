package dstypes

import "cynthia/util"

type ActivityType int

const (
	ActivityTypePlaying   ActivityType = 0
	ActivityTypeStreaming ActivityType = 1
	ActivityTypeListening ActivityType = 2
	ActivityTypeWatching  ActivityType = 3
	ActivityTypeCustom    ActivityType = 4
	ActivityTypeCompeting ActivityType = 5
)

func (a *ActivityType) UnmarshalJSON(b []byte) error {
	return util.UnmarshalNumeric(b, a)
}
