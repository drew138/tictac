package api

import (
	"net/http"

	"github.com/drew138/tictac/dependencies"

	"github.com/drew138/tictac/api/endpoints"
	"github.com/drew138/tictac/api/websockets"
	"github.com/gorilla/mux"
)

// RegisterRoutes applies specified routes to fiber app
func RegisterRoutes(r *mux.Router, d *dependencies.Dependencies) {
	// GET Endpoints

	// POST Endpoints
	r.HandleFunc("/api/v1/register", endpoints.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/login", endpoints.Login).Methods("POST")
	r.HandleFunc("/api/v1/refresh", endpoints.RefreshJWT).Methods("POST")
	// PUT Endpoints

	// PATCH Endpoints
	r.HandleFunc("/api/v1/changepass/", endpoints.ChangePassword).Methods("PATCH")
	// Websocket Endpoints
	r.HandleFunc("/ws/v1/tictactoe", func(w http.ResponseWriter, r *http.Request) {
		websockets.HandleConnection(w, r, d.WebsocketConnections)
	})
}
