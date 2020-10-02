package authentication

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashGenerator generates string of hashed and salted password
func HashGenerator(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(hash)
}
