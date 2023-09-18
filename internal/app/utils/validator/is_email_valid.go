package validator

import (
	"fmt"
	"regexp"
)

func IsEmailValid(email string) (bool, error) {
	regex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	match, err := regexp.MatchString(regex, email)

	if err != nil {
		return false, err
	}

	if !match {
		return false, fmt.Errorf("email inválido")
	}

	if len(email) < 8 {
		return false, fmt.Errorf("email inválido")
	}

	return true, nil
}
