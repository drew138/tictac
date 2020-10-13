package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/drew138/tictac/api/authentication"
	"github.com/drew138/tictac/api/authorization"
	"github.com/drew138/tictac/api/status"
	"github.com/drew138/tictac/database"
	"github.com/drew138/tictac/database/models"
)

// Login - Grant access and permissions by providing jwt
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestUser models.User // user struct containing information provided by the request
	if err := json.NewDecoder(r.Body).Decode(&requestUser); err != nil {
		status.RespondStatus(w, 400, err)
		return
	}
	var databaseUser models.User // user struct containing information found in database
	database.DBConn.Where("email = ?", requestUser.Email).First(&databaseUser)
	err := authentication.AssertPassword(databaseUser.Password, []byte(requestUser.Password))
	if err != nil {
		status.RespondStatus(w, 401, err)
		return
	}
	tokenPair, err := authorization.GenerateJWTS(&databaseUser)
	if err != nil {
		status.RespondStatus(w, 401, err)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(tokenPair)
}
