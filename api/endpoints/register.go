package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/drew138/tictac/api/authentication"
	"github.com/drew138/tictac/api/authorization"
	"github.com/drew138/tictac/api/status"
	"github.com/drew138/tictac/database"
	"github.com/drew138/tictac/database/models"
)

// CreateUser add new user to database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var User models.User
	if err := json.NewDecoder(r.Body).Decode(&User); err != nil {
		status.RespondStatus(w, 400, err)
		return
	}
	validationError := authentication.ValidatePassword(User.Password)
	if validationError != nil {
		status.RespondStatus(w, 400, validationError)
		return
	}
	if User.IsAdmin {
		var token string
		if r.Header["Authorization"] != nil {
			token = r.Header["Authorization"][0]
		} else {
			status.RespondStatus(w, 401, nil)
			return
		}
		if token == "" {
			status.RespondStatus(w, 401, nil)
			return
		}
		parsedToken, _ := authorization.ParseJWT(token, false)
		if parsedToken == nil {
			status.RespondStatus(w, 401, nil)
			return
		}
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !(parsedToken.Valid && ok) {
			status.RespondStatus(w, 401, nil)
			return
		}
		if claims["IsAdmin"] != true {
			status.RespondStatus(w, 401, nil)
			return
		}
	}
	User.Password = authentication.HashGenerator([]byte(User.Password))
	dbError := database.DBConn.Create(&User).Error
	if dbError != nil {
		status.RespondStatus(w, 500, dbError)
		return
	}
	tokenPair, err := authorization.GenerateJWTS(&User)
	if err != nil {
		status.RespondStatus(w, 500, err)
		return
	}
	userMap := map[string]interface{}{
		"Email":        User.Email,
		"Name":         User.Name,
		"Surname":      User.Surname,
		"IsAdmin":      User.IsAdmin,
		"AccessToken":  tokenPair["accessToken"],
		"RefreshToken": tokenPair["refreshToken"],
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&userMap)
}
