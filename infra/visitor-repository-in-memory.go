package infra

import (
	"context"
	"sync"

	"github.com/Falagan/web-tracker/internal/domain"
)

type UID string

type VisitorRepositoryInMemory struct {
	visitors map[UID]bool
	mu       sync.RWMutex
}

func NewVisitorRepositoryInMemory() *VisitorRepositoryInMemory {
	return &VisitorRepositoryInMemory{
		visitors: make(map[UID]bool),
	}
}

func (vr *VisitorRepositoryInMemory) AddUnique(ctx context.Context, v *domain.Visitor) error {
	vr.mu.Lock()
	defer vr.mu.Unlock()
	_, exists := vr.visitors[UID(v.UID)]

	if !exists {
		vr.visitors[UID(v.UID)] = true
	}
	return nil
}
