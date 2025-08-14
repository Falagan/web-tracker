package domain

import (
	"net/url"
	"strings"

	"github.com/Falagan/web-tracker/pkg"
)

type URL string

var (
	URLEmptyError         = pkg.Error{Code: "EDURLEMEPTY", Message: "URL cannot be empty"}
	URLInvalidFormatError = pkg.Error{Code: "EDURLINVALIDFORMAT", Message: "URL has invalid format"}
)

func NewURL(s string) (URL, error) {
	u := URL(s)
	if err := u.Validate(); err != nil {
		return "", err
	}
	return u, nil
}

func (u URL) Validate() error {
	if len(strings.TrimSpace(string(u))) == 0 {
		return &URLEmptyError
	}

	_, err := url.Parse(string(u))
	if err != nil {
		return &URLInvalidFormatError
	}

	return nil
}

func (u URL) IsValid() bool {
	return u.Validate() == nil
}

func (u URL) IsEmpty() bool {
	return len(strings.TrimSpace(string(u))) == 0
}

func (u URL) ToString() string {
	return string(u)
}

// URL Counter

type URLCount int

func NewURLCount(i int) URLCount {
	if i < 0 {
		return 0
	}
	return URLCount(i)
}

func (c URLCount) IsValid() bool {
	return c >= 0
}

func (c URLCount) IsZero() bool {
	return c == 0
}

func (c URLCount) IsPositive() bool {
	return c > 0
}

func (c URLCount) Increment() URLCount {
	if c < 0 {
		return 1
	}
	return c + 1
}

func (c URLCount) Add(value URLCount) URLCount {
	if c < 0 || value < 0 {
		return 0
	}
	return c + value
}

func (c URLCount) ToInt() int {
	return int(c)
}
