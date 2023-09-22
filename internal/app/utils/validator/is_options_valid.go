package validator

import (
	"errors"
	"strings"
)

func IsOptionsValid(options []string) (bool, error) {
	numberOfOptions := len(options)
	if numberOfOptions != 2 && numberOfOptions != 5 {
		return false, errors.New("number of options must be 2 or 5")
	}

	for _, option := range options {
		if option == "" {
			return false, errors.New("option cannot be empty")
		}
	}

	if numberOfOptions == 2 {
		for _, option := range options {
			if option == "" {
				return false, errors.New("option cannot be empty")
			}

			if option != "Certo" && option != "Errado" {
				return false, errors.New("option must be Certo or Errado")
			}
		}
	}

	if numberOfOptions == 5 {
		for _, option := range options {
			if option == "" {
				return false, errors.New("option cannot be empty")
			}

			if !strings.HasPrefix(option, "(A)") && !strings.HasPrefix(option, "(B)") && !strings.HasPrefix(option, "(C)") && !strings.HasPrefix(option, "(D)") && !strings.HasPrefix(option, "(E)") {
				return false, errors.New("option must be (A), (B), (C), (D) or (E)")
			}
		}
	}

	return true, nil
}
