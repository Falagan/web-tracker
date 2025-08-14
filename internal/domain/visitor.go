package domain

import "context"

type Visitor struct {
	UID string `json:"uid"`
	URL string `json:"url"`
}

type VisitorRepository interface {
	AddUnique(ctx context.Context, user *Visitor) error
}
