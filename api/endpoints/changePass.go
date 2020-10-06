package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/drew138/tictac/api/authentication"

	"github.com/dgrijalva/jwt-go"
	"github.com/drew138/tictac/api/authorization"
	"github.com/drew138/tictac/api/status"
	"github.com/drew138/tictac/database"
	"github.com/drew138/tictac/database/models"
	"github.com/gorilla/mux"
)

// ChangePassword - change password for a given user
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tokenString string
	if r.Header["Authorization"] != nil {
		tokenString = r.Header["Authorization"][0]
	} else {
		status.RespondStatus(w, 400, nil)
		return
	}
	token, err := authorization.ParseJWT(tokenString, false)
	if err != nil {
		status.RespondStatus(w, 401, err)
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(token.Valid && ok) {
		status.RespondStatus(w, 401, nil)
		return
	}
	newPass := mux.Vars(r)["newPassword"]

	hashedNewPass := authentication.HashGenerator([]byte(newPass))

	if hashedNewPass == "" {
		status.RespondStatus(w, 500, nil)
	}
	var User models.User
	database.DBConn.Model(&models.User{}).Where("ID = ?", claims["ID"]).Update("Password", hashedNewPass).First(&User)
	if User.ID != claims["ID"] {
		status.RespondStatus(w, 500, err)
		return
	}
	tokenPair, err := authorization.GenerateJWTS(&User)
	if err != nil {
		status.RespondStatus(w, 500, err)
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(tokenPair)
}

// https://www.loginradius.com/engineering/blog/sending-emails-with-golang/
