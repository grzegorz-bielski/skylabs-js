package models

const (
	PollsID   = "PollsID"
	PollsHash = "PollsHash"
)

type Poll struct {
	ID       int64  `json:"id"`
	Question string `json:"question"`
	Votes    []Vote `json:"votes"`
}
