package domain

import (
	"context"

	"github.com/Falagan/web-tracker/pkg"
)

type Visitor struct {
	UID UID `json:"uid"`
	URL URL `json:"url"`
}

var (
	VisitorInvalidUIDError = pkg.Error{Code: "EDVISITORUID", Message: "Visitor has invalid UID"}
	VisitorInvalidURLError = pkg.Error{Code: "EDVISITORURL", Message: "Visitor has invalid URL"}
	VisitorInvalidRequest  = pkg.Error{Code: "EDVISITORREQUEST", Message: "Visitor invalid request"}
	VisitorNotUnique       = pkg.Error{Code: "EDVISITORNOTUNIQUE", Message: "Visitor not unique"}
)

func NewVisitor(uid string, url string) (*Visitor, error) {
	validUID, err := NewUID(uid)
	if err != nil {
		return nil, err
	}

	validURL, err := NewURL(url)
	if err != nil {
		return nil, err
	}

	return &Visitor{
		UID: validUID,
		URL: validURL,
	}, nil
}
type VisitorRepository interface {
	AddUnique(ctx context.Context, user *Visitor) error
}
