package models

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	PollsID   = "PollsID"
	PollsHash = "PollsHash"
)

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	str := time.Time(t).Format(time.RFC3339)
	return []byte(fmt.Sprintf("\"%s\"", str)), nil
}

func (t *JSONTime) UnmarshalJSON(b []byte) error {
	var rawString string
	err := json.Unmarshal(b, &rawString)
	if err != nil {
		return err
	}

	parsed, err := time.Parse(time.RFC3339, rawString)
	if err != nil {
		return err
	}

	*t = JSONTime(parsed)
	return nil
}

type Poll struct {
	ID        int64    `json:"id"`
	Question  string   `json:"question"`
	Votes     []Vote   `json:"votes"`
	CreatedAt JSONTime `json:"createdAt"`
}
