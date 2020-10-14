package main

import (
	"log"
	"net/http"

	"github.com/drew138/tictac/api"
	"github.com/drew138/tictac/database"
	"github.com/drew138/tictac/environment"
	"github.com/gorilla/mux"
)

func main() {
	environment.LoadEnv()
	database.AutoMigrateDB()
	r := mux.NewRouter()
	api.RegisterRoutes(r)
	log.Println("Server started, running on port 8080.")
	if err := http.ListenAndServe("127.0.0.1:8080", r); err != nil {
		log.Fatal("Server failed to start: ", err.Error())
	}
}
