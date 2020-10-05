package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/drew138/tictac/api/authorization"
)

// RefreshJWT - function handle to refresh jwts
func RefreshJWT(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rToken string
	invalidPermissions := map[string]string{"Error": "Not authorized"}
	if r.Header["Authorization"] != nil {
		rToken = r.Header["Authorization"][0]
	} else {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(&invalidPermissions)
		return
	}
	if rToken == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{"Error": "Validation Error"})
		return
	}
	parsedRToken, err := authorization.ParseJWT(rToken, true)
	if err != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]string{"Error": err.Error()})
		return
	}
	tokenPair, err := authorization.RefreshToken(parsedRToken)
	if err != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]string{"Error": err.Error()})
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(tokenPair)
}
