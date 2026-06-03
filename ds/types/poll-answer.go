package ds

type PollAnswer struct {
	AnswerID  int       `json:"answer_id"`
	PollMedia PollMedia `json:"poll_media"`
}
