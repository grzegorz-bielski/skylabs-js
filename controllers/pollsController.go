package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"

	"github.com/skygate/skylabs-js/models"
	"github.com/skygate/skylabs-js/services"
)

type PollsController struct {
	pollsService *services.PollsService
}

func NewPollsController(pollsService *services.PollsService) *PollsController {
	return &PollsController{
		pollsService: pollsService,
	}
}

func (ps PollsController) validatePoll(poll models.Poll) []string {
	var errorMsg []string
	if poll.Question == "" {
		errorMsg = append(errorMsg, "no question")
	}
	if poll.Votes == nil {
		errorMsg = append(errorMsg, "no votes")
	}
	if poll.Votes != nil {
		for _, vote := range poll.Votes {
			if vote.Name == "" {
				errorMsg = append(errorMsg, "no vote name")
				break
			}
		}
	}

	return errorMsg
}

func (ps PollsController) AddPoll(res http.ResponseWriter, req *http.Request) {
	poll := models.Poll{
		CreatedAt: models.JSONTime(time.Now()),
	}
	json.NewDecoder(req.Body).Decode(&poll)

	errorMsg := ps.validatePoll(poll)

	if len(errorMsg) > 0 {
		res.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(res, "%s\n", "Error(s): "+strings.Join(errorMsg, ", "))
		return
	}

	jsonPoll, err := ps.pollsService.CreatePoll(poll)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "%s\n", jsonPoll)
}

func (ps PollsController) GetPoll(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	poll, err := ps.pollsService.GetPoll(models.PollsHash + vars["pollID"])
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&poll)
}

func (ps PollsController) GetPolls(res http.ResponseWriter, req *http.Request) {
	polls, err := ps.pollsService.GetPolls()
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(&polls)

}

func (ps PollsController) DeletePoll(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	if err := ps.pollsService.DeletePoll(models.PollsHash + vars["pollID"]); err != nil {
		res.WriteHeader(http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
}
