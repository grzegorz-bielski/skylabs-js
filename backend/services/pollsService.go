package services

import (
	"encoding/json"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/skygate/skylabs-js/backend/models"
)

type PollsService struct {
	dbClient *redis.Client
}

func NewPollsService(dbClient *redis.Client) *PollsService {
	return &PollsService{
		dbClient: dbClient,
	}
}

func (ps PollsService) CreatePoll(poll models.Poll) ([]byte, error) {
	pollID, _ := ps.dbClient.Incr(models.PollsHashID).Result()
	poll.ID = pollID

	var votes []models.Vote
	for _, vote := range poll.Votes {
		voteID, _ := ps.dbClient.Incr(models.VotesHashID).Result()
		vote.ID = voteID
		vote.PollID = pollID
		votes = append(votes, vote)
	}
	poll.Votes = votes

	return ps.SavePoll(poll)
}

func (ps PollsService) GetPoll(ID string) (models.Poll, error) {
	var poll models.Poll
	jsonPoll, err := ps.dbClient.Get(ID).Result()
	if err != nil {
		return poll, err
	}
	err = json.Unmarshal([]byte(jsonPoll), &poll)
	return poll, err
}

func (ps PollsService) GetPolls() ([]models.Poll, error) {
	var polls []models.Poll
	keys, _, err := ps.dbClient.Scan(0, models.PollsHash+"*", 1000).Result()
	if err != nil {
		return polls, nil
	}

	for _, key := range keys {
		poll, err := ps.GetPoll(key)
		if err != nil {
			return polls, nil
		}

		polls = append(polls, poll)
	}

	return polls, nil
}

func (ps PollsService) DeletePoll(ID string) error {
	return ps.dbClient.Del(ID).Err()
}

func (ps PollsService) SavePoll(poll models.Poll) ([]byte, error) {
	jsonPoll, err := json.Marshal(poll)
	if err != nil {
		return nil, err
	}
	err = ps.dbClient.Set(models.VotesHash+strconv.FormatInt(poll.ID, 10), jsonPoll, 0).Err()
	return jsonPoll, err
}
