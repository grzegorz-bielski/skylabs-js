package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/skygate/skylabs-js/backend/models"
	"github.com/skygate/skylabs-js/backend/services"
)

type PollsController struct {
	pollsService *services.PollsService
}

func NewPollsController(pollsService *services.PollsService) *PollsController {
	return &PollsController{
		pollsService: pollsService,
	}
}

func (ps PollsController) AddPoll(res http.ResponseWriter, req *http.Request) {
	poll := models.Poll{}
	json.NewDecoder(req.Body).Decode(&poll)

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
		res.WriteHeader(http.StatusInternalServerError)
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
