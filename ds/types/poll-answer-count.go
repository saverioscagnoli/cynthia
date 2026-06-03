package ds

type PollAnswerCount struct {
	ID      int  `json:"int"`
	Count   int  `json:"count"`
	MeVoted bool `json:"me_voted"`
}
