package domain

import (
	"net/url"
	"strings"

	"github.com/Falagan/web-tracker/pkg"
)

type URL string

var (
	URLEmptyError              = pkg.Error{Code: "EDURLEMEPTY", Message: "URL cannot be empty"}
	URLInvalidFormatError      = pkg.Error{Code: "EDURLINVALIDFORMAT", Message: "URL has invalid format"}
	URLCountInvalidFormatError = pkg.Error{Code: "EDURLCOUNTINVALIDFORMAT", Message: "URLCount has invalid format"}
)

func NewURL(s string) (URL, error) {
	if err := ValidateURL(s); err != nil {
		return "", err
	}
	return URL(s), nil
}

func (u URL) ToString() string {
	return string(u)
}

func (u URL) GetPath() (string, error) {
	parsed, err := url.Parse(string(u))
	if err != nil {
		return "", &URLInvalidFormatError
	}

	path := parsed.Path

	if path == "" {
		path = "/"
	}

	return path, nil
}

func ValidateURL(s string) error {
	if err := validateURLEmpty(s); err != nil {
		return err
	}
	
	return nil
}

func validateURLEmpty(s string) error {
	if len(strings.TrimSpace(s)) == 0 {
		return &URLEmptyError
	}
	return nil
}

// URL Counter

type URLCount int


func NewURLCount(i int) (URLCount, error) {
	if err := ValidateURLCount(i); err != nil {
		return 0, err
	}
	return URLCount(i), nil
}

func (c URLCount) ToInt() int {
	return int(c)
}

func ValidateURLCount(i int) error {
	if err := validateURLCountRange(i); err != nil {
		return err
	}
	
	return nil
}

func validateURLCountRange(i int) error {
	if i < 0 {
		return &URLCountInvalidFormatError
	}
	return nil
}
