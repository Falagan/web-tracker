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
	vr.mu.Lock()
	defer vr.mu.Unlock()

	path, err := v.URL.GetPath()

	if err != nil {
		return &domain.URLInvalidFormatError
	}

	uniqueKey := v.UID.ToString() + path
	uidBytes := []byte(uniqueKey)

	// only adds if its no present
	if !vr.bloomFilter.Test(uidBytes) {
		vr.bloomFilter.Add(uidBytes)
		return nil
	}

	return nil
}
