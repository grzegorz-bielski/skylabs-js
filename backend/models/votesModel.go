package models

const (
	VotesID   = "VotesID"
	VotesHash = "VotesHash"
)

type Vote struct {
	ID     int64  `json:"id"`
	PollID int64  `json:"pollId"`
	Name   string `json:"name"`
	Score  int64  `json:"score"`
}
