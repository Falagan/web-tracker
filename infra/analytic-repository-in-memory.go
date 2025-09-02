package infra

import (
	"context"
	"errors"
	"sync"

	"github.com/Falagan/web-tracker/internal/domain"
)

type AnalyticRepositoryInMemory struct {
	urlCounts map[string]domain.URLCount
	mu        sync.RWMutex
}

func NewAnalyticRepositoryInMemory() *AnalyticRepositoryInMemory {
	return &AnalyticRepositoryInMemory{
		urlCounts: make(map[string]domain.URLCount),
	}
}

func (ar *AnalyticRepositoryInMemory) IncreaseVisitedURLCount(ctx context.Context, url string) error {
	ar.mu.Lock()
	defer ar.mu.Unlock()

	_, exists := ar.urlCounts[url]

	if !exists {
		ar.urlCounts[url] = 0
	}

	ar.urlCounts[url]++
	return nil
}

func (ar *AnalyticRepositoryInMemory) GetVisitedURLCount(ctx context.Context, url string) (*domain.URLCount, error) {
	ar.mu.RLock()
	defer ar.mu.RUnlock()

	count, exists := ar.urlCounts[url]

	if !exists {
		return nil, errors.New("no data")
	}
	return &count, nil
}
