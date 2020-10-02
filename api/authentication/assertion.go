package authentication

import (
	"golang.org/x/crypto/bcrypt"
)

// AssertPassword returns true if password is correct, false if incorrect
func AssertPassword(hashString string, password []byte) error {
	byteHashString := []byte(hashString)
	err := bcrypt.CompareHashAndPassword(byteHashString, password)
	return err
}
