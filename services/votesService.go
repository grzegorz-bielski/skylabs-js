package services

import (
	"errors"
	"strconv"

	"github.com/skygate/skylabs-js/models"
)

type VotesService struct{}

func NewVotesService() *VotesService {
	return &VotesService{}
}

func (vs VotesService) Vote(votes []models.Vote, ID string, value int64) ([]models.Vote, error) {
	id, _ := strconv.ParseInt(ID, 10, 64)

	var found bool
	var newVotes []models.Vote
	for _, vote := range votes {
		if vote.ID == id {
			vote.Score += value
			found = true
		}
		newVotes = append(newVotes, vote)
	}

	if !found {
		return newVotes, errors.New("Not Found")
	}

	return newVotes, nil
}
