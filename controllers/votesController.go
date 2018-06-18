package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/skygate/skylabs-js/models"
	"github.com/skygate/skylabs-js/services"
)

type VotesController struct {
	pollsService *services.PollsService
	votesService *services.VotesService
}

func NewVotesController(pollsService *services.PollsService, votesService *services.VotesService) *VotesController {
	return &VotesController{
		pollsService: pollsService,
		votesService: votesService,
	}
}

func (vc VotesController) vote(res http.ResponseWriter, req *http.Request, value int64) {
	vars := mux.Vars(req)
	poll, err := vc.pollsService.GetPoll(models.PollsHash + vars["pollID"])

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	votes, err := vc.votesService.Vote(poll.Votes, vars["voteID"], value)
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	poll.Votes = votes
	if _, err := vc.pollsService.SavePoll(poll); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}

func (vc VotesController) UpVote(res http.ResponseWriter, req *http.Request) {
	vc.vote(res, req, 1)
}
func (vc VotesController) DownVote(res http.ResponseWriter, req *http.Request) {
	vc.vote(res, req, -1)
}
