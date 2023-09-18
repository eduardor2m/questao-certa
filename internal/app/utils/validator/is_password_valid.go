package validator

import (
	"errors"
	"unicode"
)

func IsPasswordValid(password string) (bool, error) {
	if len(password) < 8 {
		return false, errors.New("password is too short")
	}

	hasUpperCase := false
	hasLowerCase := false
	hasDigit := false
	hasSpecialChar := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpperCase = true
		}
		if unicode.IsLower(char) {
			hasLowerCase = true
		}
		if unicode.IsDigit(char) {
			hasDigit = true
		}
		if unicode.In(char, unicode.Punct, unicode.Symbol) {
			hasSpecialChar = true
		}
	}

	if !hasUpperCase {
		return false, errors.New("password must have at least one upper case character")
	}
	if !hasLowerCase {
		return false, errors.New("password must have at least one lower case character")
	}
	if !hasDigit {
		return false, errors.New("password must have at least one digit")
	}
	if !hasSpecialChar {
		return false, errors.New("password must have at least one special character")
	}

	return true, nil
}
