package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/drew138/tictac/api/authentication"
	"github.com/drew138/tictac/api/authorization"
	"github.com/drew138/tictac/database"
	"github.com/drew138/tictac/database/models"
)

// CreateUser add new user to database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var User models.User
	if err := UnmarshalJSON(r, &User); err != nil {
		json.NewEncoder(w).Encode(map[string]string{"Error": err.Error()})
		w.WriteHeader(400)
		return
	}
	validationError := authentication.ValidatePassword(User.Password)
	if validationError != nil {
		json.NewEncoder(w).Encode(map[string]string{"Error": validationError.Error()})
		w.WriteHeader(400)
		return
	}
	User.Password = authentication.HashGenerator([]byte(User.Password))
	dbError := database.DBConn.Create(&User).Error
	if dbError != nil {
		json.NewEncoder(w).Encode(map[string]string{"Error": dbError.Error()})
		w.WriteHeader(500)
		return
	}
	tokenPair, err := authorization.GenerateJWT(&User)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"Error": err.Error()})
		w.WriteHeader(500)
		return
	}
	userMap := map[string]interface{}{
		"email":        User.Email,
		"name":         User.Name,
		"surname":      User.Surname,
		"isAdmin":      User.IsAdmin, //TODO remove this field
		"accessToken":  tokenPair["accessToken"],
		"refreshToken": tokenPair["refreshToken"],
	}
	json.NewEncoder(w).Encode(userMap)
	w.WriteHeader(201)
}
