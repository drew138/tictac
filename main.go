package main

import (
	"log"
	"net/http"

	"github.com/drew138/tictac/api"
	"github.com/drew138/tictac/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	loadEnv()
	database.AutoMigrateDB()
	r := mux.NewRouter()
	api.RegisterRoutes(r)
	log.Println("Server started, running on port 8080.")
	if err := http.ListenAndServe("127.0.0.1:8080", r); err != nil {
		log.Fatal("Server failed to start: ", err.Error())
	}
}
