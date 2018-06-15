package main

import (
	"net/http"
	"os"

	"github.com/skygate/skylabs-js/backend/services"

	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/skygate/skylabs-js/backend/controllers"
)

var port = os.Getenv("PORT")

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		DB:   0,
	})

	pollsService := services.NewPollsService(client)

	pollsController := controllers.NewPollsController(pollsService)

	router := mux.NewRouter()
	router.HandleFunc("/api/poll", pollsController.AddPoll).Methods(http.MethodPost)

	http.ListenAndServe(":"+port, router)
}
