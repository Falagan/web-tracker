package infra

import (
	"context"
	"sync"

	"github.com/Falagan/web-tracker/internal/domain"
	"github.com/bits-and-blooms/bloom/v3"
)

type VisitorRepositoryInMemoryBloom struct {
	bloomFilter *bloom.BloomFilter
	mu          sync.RWMutex
}

func NewVisitorRepositoryInMemoryBloom(expectedElements uint, falsePositiveRate float64) *VisitorRepositoryInMemoryBloom {
	bf := bloom.NewWithEstimates(expectedElements, falsePositiveRate)

	return &VisitorRepositoryInMemoryBloom{
		bloomFilter: bf,
	}
}

func (vr *VisitorRepositoryInMemoryBloom) AddUnique(ctx context.Context, v *domain.Visitor) error {
	uidBytes := []byte(v.UID)
	vr.mu.Lock()

	defer vr.mu.Unlock()

	vr.bloomFilter.Add(uidBytes)
	return nil
}
