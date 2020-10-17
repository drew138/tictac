package main

import (
	"log"
	"net/http"

	"github.com/drew138/tictac/dependencies"

	"github.com/drew138/tictac/api"
	"github.com/drew138/tictac/api/websockets/connections"
	"github.com/drew138/tictac/database"
	"github.com/drew138/tictac/environment"
	"github.com/gorilla/mux"
)

func main() {
	var dependencies dependencies.Dependencies
	// Load dependencies
	environment.LoadEnv()
	database.AutoMigrateDB()
	dependencies.WebsocketConnections = connections.StartConnectionTracking()
	// Mount services
	r := mux.NewRouter()
	api.RegisterRoutes(r, &dependencies)
	log.Println("Server started, running on port 8080.")
	if err := http.ListenAndServe("127.0.0.1:8080", r); err != nil {
		log.Fatal("Server failed to start: ", err.Error())
	}
}
