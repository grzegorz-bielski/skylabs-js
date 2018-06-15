package models

const (
	PollsHashID = "PollsHashID"
	PollsHash   = "PollsHash"
)

type Poll struct {
	ID       int64  `json:"id"`
	Question string `json:"question"`
	Votes    []Vote `json:"votes"`
}
