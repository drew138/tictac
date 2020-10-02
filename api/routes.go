package api

import (
	"github.com/drew138/tictac/api/endpoints"
	"github.com/gorilla/mux"
)

// ResgisterEndPoints applies specified routes to fiber app
func ResgisterEndPoints(r *mux.Router) {
	// GET Endpoints

	// app.Get("/")
	// POST Endpoints
	r.HandleFunc("/api/v1/register", endpoints.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/login", endpoints.Login).Methods("POST")
	r.HandleFunc("/api/v1/refresh", endpoints.RefreshJWT).Methods("POST")
	//PUT Endpoints

	//PATCH Endpoints

}
