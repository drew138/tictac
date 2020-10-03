package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/drew138/tictac/api/authentication"
	"github.com/drew138/tictac/api/authorization"
	"github.com/drew138/tictac/database"
	"github.com/drew138/tictac/database/models"
)

// Login - Grant access and permissions by providing jwt
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := new(models.User) // request user
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&map[string]string{"Error": err.Error()})
		return
	}
	var User models.User // user in database
	database.DBConn.Where("email = ?", user.Email).First(&User)
	err := authentication.AssertPassword(User.Password, []byte(user.Password))
	if err != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(&map[string]string{"Error": err.Error()})
		return
	}
	tokenPair, err := authorization.GenerateJWT(user)
	if err != nil {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(&map[string]string{"Error": err.Error()})
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(tokenPair)
}
