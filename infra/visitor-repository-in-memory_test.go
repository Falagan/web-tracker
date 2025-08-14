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

	t.Run("add_duplicate_visitor_error", func(t *testing.T) {
		// arrange
		testCase := tb.WithDuplicateVisitor()
		// act
		repo := NewVisitorRepositoryInMemory()
		// add first time
		err1 := repo.AddUnique(context.Background(), testCase.visitor)
		// add same visitor again
		err2 := repo.AddUnique(context.Background(), testCase.visitor)
		// assert
		assert.NoError(t, err1)
		assert.Error(t, err2)
		assert.Equal(t, testCase.expectedError, err2.Error())
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

	t.Run("concurrent_duplicate_visitors_single_success", func(t *testing.T) {
		// arrange
		testCase := tb.WithConcurrentDuplicateVisitors()
		// act
		repo := NewVisitorRepositoryInMemory()
		var wg sync.WaitGroup
		successCount := make(chan bool, testCase.numGoroutines)
		errorCount := make(chan bool, testCase.numGoroutines)

		for i := 0; i < testCase.numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				if err := repo.AddUnique(context.Background(), testCase.visitor); err != nil {
					errorCount <- true
				} else {
					successCount <- true
				}
			}()
		}

		wg.Wait()
		close(successCount)
		close(errorCount)

		// assert
		successes := len(successCount)
		errors := len(errorCount)
		assert.Equal(t, 1, successes)
		assert.Equal(t, testCase.numGoroutines-1, errors)
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
	visitor        *domain.Visitor
	visitors       []*domain.Visitor
	expectedError  string
	numGoroutines  int
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
		expectedError: "not unique",
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