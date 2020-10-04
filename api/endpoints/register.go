package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/drew138/tictac/api/authentication"
	"github.com/drew138/tictac/api/authorization"
	"github.com/drew138/tictac/database"
	"github.com/drew138/tictac/database/models"
)

// CreateUser add new user to database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var User models.User
	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&map[string]string{"Error": err.Error()})
		return
	}
	validationError := authentication.ValidatePassword(User.Password)
	if validationError != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(&map[string]string{"Error": validationError.Error()})
		return
	}

	if User.IsAdmin {
		invalidPermissions := map[string]string{"Error": "Not authorized"}
		token := r.Header["Authorization"][0]
		if token == "" {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(&invalidPermissions)
			return
		}
		parsedToken, _ := authorization.ParseJWT(token)
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !(parsedToken.Valid && ok) {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(&invalidPermissions)
			return
		}
		if claims["isAdmin"] != true {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(&invalidPermissions)
			return
		}
	}

	User.Password = authentication.HashGenerator([]byte(User.Password))
	dbError := database.DBConn.Create(&User).Error
	if dbError != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&map[string]string{"Error": dbError.Error()})
		return
	}
	tokenPair, err := authorization.GenerateJWT(&User)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(&map[string]string{"Error": err.Error()})
		return
	}
	userMap := map[string]interface{}{
		"email":        User.Email,
		"name":         User.Name,
		"surname":      User.Surname,
		"isAdmin":      User.IsAdmin,
		"accessToken":  tokenPair["accessToken"],
		"refreshToken": tokenPair["refreshToken"],
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&userMap)
}
