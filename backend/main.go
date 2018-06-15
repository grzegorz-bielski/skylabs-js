package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var port = os.Getenv("PORT")

func main() {
	router := mux.NewRouter()
	http.ListenAndServe(":"+port, router)
}
