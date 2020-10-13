package authorization

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/drew138/tictac/database/models"
)

func createToken(isRefresh bool, user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = user.ID
	var secretKey string
	if !isRefresh {
		claims["Email"] = user.Email
		claims["Name"] = user.Name
		claims["Surname"] = user.Surname
		claims["IsAdmin"] = user.IsAdmin
		claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
		secretKey = os.Getenv("JWT_SECRET_KEY")
	} else {
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
		secretKey = os.Getenv("JWT_REFRESH_SECRET_KEY")
	}
	return token.SignedString([]byte(secretKey))
}

// GenerateJWTS return jwt tokens
func GenerateJWTS(user *models.User) (map[string]string, error) {
	tokensMap := map[string]string{}
	token, err := createToken(false, user)
	if err != nil {
		return tokensMap, fmt.Errorf("Error: %s", err.Error())
	}
	refreshToken, err := createToken(true, user)
	if err != nil {
		return tokensMap, fmt.Errorf("Error: %s", err.Error())
	}
	tokensMap["accessToken"] = token
	tokensMap["refreshToken"] = refreshToken
	return tokensMap, nil
}

// https://medium.com/monstar-lab-bangladesh-engineering/jwt-auth-in-go-part-2-refresh-tokens-d334777ca8a0
