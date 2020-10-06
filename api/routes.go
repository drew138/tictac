package api

import (
	"github.com/drew138/tictac/api/endpoints"
	"github.com/drew138/tictac/api/websockets"
	"github.com/gorilla/mux"
)

// RegisterRoutes applies specified routes to fiber app
func RegisterRoutes(r *mux.Router) {
	// GET Endpoints

	// POST Endpoints
	r.HandleFunc("/api/v1/register", endpoints.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/login", endpoints.Login).Methods("POST")
	r.HandleFunc("/api/v1/refresh", endpoints.RefreshJWT).Methods("POST")
	// PUT Endpoints

	// PATCH Endpoints
	r.HandleFunc("/api/v1/changepass/{newPassword:}", endpoints.ChangePassword).Methods("PATCH")
	// Websocket Endpoints
	r.HandleFunc("/ws/v1/tictactoe", websockets.HandleConnection)
}
