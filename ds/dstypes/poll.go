package dstypes

type PollLayoutType = int

const (
	PollLayoutTypeDefault PollLayoutType = 1
)

type Poll struct {
	Question         PollMedia      `json:"question"`
	Answers          []PollAnswer   `json:"poll_answers"`
	Expiry           string         `json:"expiry"`
	AllowMultiselect bool           `json:"allow_multiselect"`
	LayoutType       PollLayoutType `json:"layout_type"`
	Results          *PollResults   `json:"results"`
}
