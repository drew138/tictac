package authorization

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/drew138/tictac/database/models"
)

// GenerateJWT return jwt token
func GenerateJWT(user *models.User) (map[string]string, error) {
	var JWTSecretKey = os.Getenv("JWT_SECRET_KEY")
	var JWTRefreshSecretKey = os.Getenv("JWT_REFRESH_SECRET_KEY")
	tokensMap := map[string]string{}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["ID"] = user.ID
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["surname"] = user.Surname
	claims["isAdmin"] = user.IsAdmin
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	tokenString, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return tokensMap, fmt.Errorf("Error: %s", err.Error())
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["ID"] = user.ID
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refreshTokenString, err := refreshToken.SignedString([]byte(JWTRefreshSecretKey))
	if err != nil {
		return tokensMap, fmt.Errorf("Error: %s", err.Error())
	}
	tokensMap["accessToken"] = tokenString
	tokensMap["refreshToken"] = refreshTokenString
	return tokensMap, nil
}

// https://medium.com/monstar-lab-bangladesh-engineering/jwt-auth-in-go-part-2-refresh-tokens-d334777ca8a0
