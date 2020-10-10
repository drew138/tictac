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
	user := new(models.User) // request user
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		status.RespondStatus(w, 400, err)
		return
	}
	var User models.User // user in database
	database.DBConn.Where("email = ?", user.Email).First(&User)
	err := authentication.AssertPassword(User.Password, []byte(user.Password))
	if err != nil {
		status.RespondStatus(w, 401, err)
		return
	}
	tokenPair, err := authorization.GenerateJWTS(&User)
	if err != nil {
		status.RespondStatus(w, 401, err)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(tokenPair)
}
