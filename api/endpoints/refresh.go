package endpoints

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/drew138/tictac/api/authorization"
)

// RefreshJWT - function handle to refresh jwts
func RefreshJWT(w http.ResponseWriter, r *http.Request) {
	rToken := strings.Split(r.Header.Get("Authorization"), " ")[1]
	parsedRToken, err := authorization.ParseJWT(rToken)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"Error": err.Error()})
		w.WriteHeader(401)
		return
	}
	tokenPair, err := authorization.RefreshToken(parsedRToken)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"Error": err.Error()})
		w.WriteHeader(401)
		return
	}
	json.NewEncoder(w).Encode(tokenPair)
	w.WriteHeader(201)
}
