package authentication

import (
	"errors"
	"unicode"
)

// ValidatePassword - check if entered password is secure
func ValidatePassword(password string) error {
	var (
		hasUpper     = false
		hasLower     = false
		hasNumber    = false
		isContineous = true
		isLong       = len(password) > 7
	)
	for _, char := range password {

		switch {
		case unicode.IsNumber(char):
			hasNumber = true
			break
		case unicode.IsUpper(char):
			hasUpper = true
			break
		case unicode.IsLower(char):
			hasLower = true
			break
		case char == ' ':
			isContineous = false
		}
	}
	if !hasUpper {
		return errors.New("Password does not contain an uppercase letter")
	}
	if !hasLower {
		return errors.New("Password does not contain a lowercase letter")
	}
	if !hasNumber {
		return errors.New("Password does not contain a number")
	}
	if !hasNumber {
		return errors.New("Password does not contain a number")
	}
	if !isLong {
		return errors.New("Password must be of 8 characters or more")
	}
	if !isContineous {
		return errors.New("Password must not contain any white spaces")
	}
	return nil
}
