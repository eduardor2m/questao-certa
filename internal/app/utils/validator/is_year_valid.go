package validator

import (
	"errors"
	"regexp"
)

func IsYearValid(year string) (bool, error) {
	if year == "" {
		return false, errors.New("year is required")
	}

	regexp := regexp.MustCompile(`^[0-9]{4}$`)
	if !regexp.MatchString(year) {
		return false, errors.New("year must be a valid year")
	}

	return true, nil
}
