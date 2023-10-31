package validator

import (
	"errors"
	"strings"
)

func IsAnswerValid(answer string) (bool, error) {
	if answer == "" {
		return false, errors.New("answer is required")
	} else if answer != "Certo" && answer != "Errado" && !strings.HasPrefix(answer, "(A)") && !strings.HasPrefix(answer, "(B)") && !strings.HasPrefix(answer, "(C)") && !strings.HasPrefix(answer, "(D)") && !strings.HasPrefix(answer, "(E)") {
		return false, errors.New("answer must be Certo, Errado, (A), (B), (C), (D) or (E)")
	}

	return true, nil
}
