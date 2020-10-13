package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/drew138/tictac/api/authorization"
	"github.com/drew138/tictac/api/status"
)

// RefreshJWT - function handle to refresh jwts
func RefreshJWT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var rToken string
	if r.Header["Authorization"] != nil {
		rToken = r.Header["Authorization"][0]
	} else {
		status.RespondStatus(w, 401, nil)
		return
	}
	if rToken == "" {
		status.RespondStatus(w, 400, nil)
		return
	}
	parsedRToken, err := authorization.ParseJWT(rToken, true)
	if err != nil {
		status.RespondStatus(w, 401, err)
		return
	}
	tokenPair, err := authorization.RefreshToken(parsedRToken)
	if err != nil {
		status.RespondStatus(w, 401, err)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(tokenPair)
}
