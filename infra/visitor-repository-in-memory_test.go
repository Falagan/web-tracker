package infra

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/Falagan/web-tracker/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestVisitorRepositoryInMemory_AddUnique(t *testing.T) {
	tb := NewTestVisitorRepositoryBuilder()

	t.Run("add_unique_visitor_success", func(t *testing.T) {
		// arrange
		testCase := tb.WithUniqueVisitor()
		// act
		repo := NewVisitorRepositoryInMemory()
		err := repo.AddUnique(context.Background(), testCase.visitor)
		// assert
		assert.NoError(t, err)
		// verify visitor was added
		repo.mu.RLock()
		exists := repo.visitors[testCase.visitor.UID]
		repo.mu.RUnlock()
		assert.True(t, exists)
	})

	t.Run("add_duplicate_visitor_success", func(t *testing.T) {
		// arrange
		testCase := tb.WithDuplicateVisitor()
		// act
		repo := NewVisitorRepositoryInMemory()
		// add first time
		err1 := repo.AddUnique(context.Background(), testCase.visitor)
		// add same visitor again (should be idempotent)
		err2 := repo.AddUnique(context.Background(), testCase.visitor)
		// assert
		assert.NoError(t, err1)
		assert.NoError(t, err2)
		// verify visitor exists only once
		repo.mu.RLock()
		exists := repo.visitors[testCase.visitor.UID]
		count := len(repo.visitors)
		repo.mu.RUnlock()
		assert.True(t, exists)
		assert.Equal(t, 1, count)
	})

	t.Run("add_multiple_unique_visitors_success", func(t *testing.T) {
		// arrange
		testCase := tb.WithMultipleUniqueVisitors()
		// act
		repo := NewVisitorRepositoryInMemory()
		for _, visitor := range testCase.visitors {
			err := repo.AddUnique(context.Background(), visitor)
			assert.NoError(t, err)
		}
		// assert
		repo.mu.RLock()
		count := len(repo.visitors)
		repo.mu.RUnlock()
		assert.Equal(t, len(testCase.visitors), count)
	})

	t.Run("concurrent_unique_visitors_success", func(t *testing.T) {
		// arrange
		testCase := tb.WithConcurrentUniqueVisitors()
		// act
		repo := NewVisitorRepositoryInMemory()
		var wg sync.WaitGroup
		errors := make(chan error, len(testCase.visitors))

		for _, visitor := range testCase.visitors {
			wg.Add(1)
			go func(v *domain.Visitor) {
				defer wg.Done()
				if err := repo.AddUnique(context.Background(), v); err != nil {
					errors <- err
				}
			}(visitor)
		}

		wg.Wait()
		close(errors)

		// assert
		for err := range errors {
			assert.NoError(t, err)
		}

		repo.mu.RLock()
		count := len(repo.visitors)
		repo.mu.RUnlock()
		assert.Equal(t, len(testCase.visitors), count)
	})

	t.Run("concurrent_duplicate_visitors_all_success", func(t *testing.T) {
		// arrange
		testCase := tb.WithConcurrentDuplicateVisitors()
		// act
		repo := NewVisitorRepositoryInMemory()
		var wg sync.WaitGroup
		errorChan := make(chan error, testCase.numGoroutines)

		for i := 0; i < testCase.numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := repo.AddUnique(context.Background(), testCase.visitor); err != nil {
					errorChan <- err
				}
			}()
		}

		wg.Wait()
		close(errorChan)

		// assert - all operations should succeed (idempotent)
		for err := range errorChan {
			assert.NoError(t, err)
		}

		// verify only one visitor entry exists
		repo.mu.RLock()
		count := len(repo.visitors)
		repo.mu.RUnlock()
		assert.Equal(t, 1, count)
	})
}

func TestNewVisitorRepositoryInMemory(t *testing.T) {
	// act
	repo := NewVisitorRepositoryInMemory()
	// assert
	assert.NotNil(t, repo)
	assert.NotNil(t, repo.visitors)
	assert.Equal(t, 0, len(repo.visitors))
}

// test builder

type TestVisitorRepositoryBuilder struct{}

type TestVisitorRepositoryCase struct {
	visitor       *domain.Visitor
	visitors      []*domain.Visitor
	numGoroutines int
}

func NewTestVisitorRepositoryBuilder() *TestVisitorRepositoryBuilder {
	return &TestVisitorRepositoryBuilder{}
}

func (tb *TestVisitorRepositoryBuilder) WithUniqueVisitor() *TestVisitorRepositoryCase {
	return &TestVisitorRepositoryCase{
		visitor: &domain.Visitor{
			UID: domain.UID("unique-visitor-123"),
		},
	}
}

func (tb *TestVisitorRepositoryBuilder) WithDuplicateVisitor() *TestVisitorRepositoryCase {
	return &TestVisitorRepositoryCase{
		visitor: &domain.Visitor{
			UID: domain.UID("duplicate-visitor-456"),
		},
	}
}

func (tb *TestVisitorRepositoryBuilder) WithMultipleUniqueVisitors() *TestVisitorRepositoryCase {
	return &TestVisitorRepositoryCase{
		visitors: []*domain.Visitor{
			{UID: domain.UID("visitor-1")},
			{UID: domain.UID("visitor-2")},
			{UID: domain.UID("visitor-3")},
			{UID: domain.UID("visitor-4")},
			{UID: domain.UID("visitor-5")},
		},
	}
}

func (tb *TestVisitorRepositoryBuilder) WithConcurrentUniqueVisitors() *TestVisitorRepositoryCase {
	const numVisitors = 100
	visitors := make([]*domain.Visitor, numVisitors)
	
	for i := 0; i < numVisitors; i++ {
		visitors[i] = &domain.Visitor{
			UID: domain.UID(fmt.Sprintf("concurrent-visitor-%d", i)),
		}
	}

	return &TestVisitorRepositoryCase{
		visitors: visitors,
	}
}

func (tb *TestVisitorRepositoryBuilder) WithConcurrentDuplicateVisitors() *TestVisitorRepositoryCase {
	return &TestVisitorRepositoryCase{
		visitor: &domain.Visitor{
			UID: domain.UID("concurrent-duplicate-visitor"),
		},
		numGoroutines: 10,
	}
}