package authorization

import (
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

// ParseJWT - parse jwt and return data
func ParseJWT(tokenString string, isRefresh bool) (*jwt.Token, error) {
	tok, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // handler function as described in documentation
		if !ok {
			return nil, fmt.Errorf("Unauthorized token")
		}
		if isRefresh {
			return []byte(os.Getenv("JWT_REFRESH_SECRET_KEY")), nil
		}
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	return tok, err
}
