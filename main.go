package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"

	"github.com/skygate/skylabs-js/controllers"
	"github.com/skygate/skylabs-js/services"
)

var (
	port   = os.Getenv("PORT")
	prefix = "/api"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})

	pollsService := services.NewPollsService(client)
	votesService := services.NewVotesService()

	pollsController := controllers.NewPollsController(pollsService)
	votesController := controllers.NewVotesController(pollsService, votesService)

	router := mux.NewRouter()

	router.HandleFunc(prefix+"/poll", pollsController.AddPoll).Methods(http.MethodPost)
	router.HandleFunc(prefix+"/polls", pollsController.GetPolls).Methods(http.MethodGet)
	router.HandleFunc(prefix+"/polls/{pollID:[0-9]+}", pollsController.GetPoll).Methods(http.MethodGet)
	router.HandleFunc(prefix+"/polls/{pollID:[0-9]+}", pollsController.DeletePoll).Methods(http.MethodDelete)

	router.HandleFunc(prefix+"/polls/{pollID:[0-9]+}/votes/{voteID:[0-9]+}", votesController.UpVote).Methods(http.MethodPost)
	router.HandleFunc(prefix+"/polls/{pollID:[0-9]+}/votes/{voteID:[0-9]+}", votesController.DownVote).Methods(http.MethodDelete)

	if err := http.ListenAndServe(":"+port, router); err != nil {
		panic(err)
	}

	fmt.Println("App is listening on port: " + port)
}
