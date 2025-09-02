package infra

import (
	"context"
	"errors"
	"sync"

	"github.com/Falagan/web-tracker/internal/domain"
)

type VisitorRepositoryInMemory struct {
	visitors map[string]bool
	mu       sync.RWMutex
}

func NewVisitorRepositoryInMemory() *VisitorRepositoryInMemory {
	return &VisitorRepositoryInMemory{
		visitors: make(map[string]bool),
	}
}

func (vr *VisitorRepositoryInMemory) AddUnique(ctx context.Context, v *domain.Visitor) error {
	vr.mu.Lock()
	defer vr.mu.Unlock()

	path, err := v.URL.GetPath()

	if err != nil {
		return &domain.URLInvalidFormatError
	}

	uniqueKey := v.UID.ToString() + path

	_, exists := vr.visitors[uniqueKey]
	if !exists {
		vr.visitors[uniqueKey] = true
		return nil
	}

	return errors.New("not unique")
}
