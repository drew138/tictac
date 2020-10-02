package authorization

import (
	"fmt"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

// ParseJWT - parse jwt and return data
func ParseJWT(token string) (*jwt.Token, error) {
	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // handler function as described in documentation
		if !ok {
			return nil, fmt.Errorf("Unauthorized token")
		}
		return os.Getenv("JWT_SECRET_KEY"), nil
	})
	return tok, err
}
