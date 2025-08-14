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
	uid := UID(s)
	if err := uid.Validate(); err != nil {
		return "", err
	}
	return uid, nil
}

func (u UID) Validate() error {
	if u.IsEmpty() {
		return &UIDEmptyError
	}

	return nil
}

func (u UID) IsValid() bool {
	return u.Validate() == nil
}

func (u UID) IsEmpty() bool {
	return len(strings.TrimSpace(u.ToString())) == 0
}

func (u UID) ToString() string {
	return string(u)
}
