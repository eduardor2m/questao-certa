package validator

import "errors"

func IsModelValid(model string) (bool, error) {
	if model == "" {
		return false, errors.New("model is required")
	} else if model != "multiple_choice" && model != "true_or_false" {
		return false, errors.New("model must be multiple_choice or true_or_false")
	}

	return true, nil
}
