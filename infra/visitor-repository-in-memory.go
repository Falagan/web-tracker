package infra

import (
	"context"
	"errors"
	"sync"

	"github.com/Falagan/web-tracker/internal/domain"
)

type VisitorRepositoryInMemory struct {
	visitors map[domain.UID]bool
	mu       sync.RWMutex
}

func NewVisitorRepositoryInMemory() *VisitorRepositoryInMemory {
	return &VisitorRepositoryInMemory{
		visitors: make(map[domain.UID]bool),
	}
}

func (vr *VisitorRepositoryInMemory) AddUnique(ctx context.Context, v *domain.Visitor) error {
	vr.mu.Lock()
	defer vr.mu.Unlock()

	_, exists := vr.visitors[v.UID]
	if !exists {
		vr.visitors[v.UID] = true
		return nil
	}

	return errors.New("not unique")
}
