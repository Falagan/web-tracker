package domain

import (
	"strings"

	"github.com/Falagan/web-tracker/pkg"
)

type UID string

var (
	UIDEmptyError = pkg.Error{Code: "EDUIDEEMPTY", Message: "UID cannot be empty"}
)

func NewUID(s string) (UID, error) {
	if err := ValidateUID(s); err != nil {
		return "", err
	}
	return UID(s), nil
}

func (u UID) ToString() string {
	return string(u)
}

func ValidateUID(s string) error {
	if err := validateUIDEmpty(s); err != nil {
		return err
	}
	
	return nil
}

func validateUIDEmpty(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return &UIDEmptyError
	}
	return nil
}
