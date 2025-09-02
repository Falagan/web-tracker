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
		repo.AddUnique(context.Background(), testCase.visitor)
		// verify visitor was added
		path, _ := testCase.visitor.URL.GetPath()
		uniqueKey := testCase.visitor.UID.ToString() + path
		repo.mu.RLock()
		exists := repo.visitors[uniqueKey]
		repo.mu.RUnlock()
		assert.True(t, exists)
	})

	t.Run("add_duplicate_visitor_idempotent", func(t *testing.T) {
		// arrange
		testCase := tb.WithDuplicateVisitor()
		// act
		repo := NewVisitorRepositoryInMemory()
		// add first time
		repo.AddUnique(context.Background(), testCase.visitor)
		// add same visitor again (should be idempotent)
		repo.AddUnique(context.Background(), testCase.visitor)
		// verify visitor exists only once
		path, _ := testCase.visitor.URL.GetPath()
		uniqueKey := testCase.visitor.UID.ToString() + path
		repo.mu.RLock()
		exists := repo.visitors[uniqueKey]
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
			repo.AddUnique(context.Background(), visitor)
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

		for _, visitor := range testCase.visitors {
			wg.Add(1)
			go func(v *domain.Visitor) {
				defer wg.Done()
				repo.AddUnique(context.Background(), v)
			}(visitor)
		}

		wg.Wait()

		// assert
		repo.mu.RLock()
		count := len(repo.visitors)
		repo.mu.RUnlock()
		assert.Equal(t, len(testCase.visitors), count)
	})

	t.Run("concurrent_duplicate_visitors_idempotent", func(t *testing.T) {
		// arrange
		testCase := tb.WithConcurrentDuplicateVisitors()
		// act
		repo := NewVisitorRepositoryInMemory()
		var wg sync.WaitGroup

		for i := 0; i < testCase.numGoroutines; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				repo.AddUnique(context.Background(), testCase.visitor)
			}()
		}

		wg.Wait()

		// assert - all operations succeed (idempotent), only one entry exists
		repo.mu.RLock()
		count := len(repo.visitors)
		repo.mu.RUnlock()
		assert.Equal(t, 1, count)
	})

	t.Run("same_visitor_different_urls_success", func(t *testing.T) {
		// arrange
		testCase := tb.WithSameVisitorDifferentURLs()
		// act
		repo := NewVisitorRepositoryInMemory()
		repo.AddUnique(context.Background(), testCase.visitors[0])
		repo.AddUnique(context.Background(), testCase.visitors[1])
		// verify both entries exist
		repo.mu.RLock()
		count := len(repo.visitors)
		repo.mu.RUnlock()
		assert.Equal(t, 2, count)
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
			URL: domain.URL("https://example.com/page1"),
		},
	}
}

func (tb *TestVisitorRepositoryBuilder) WithDuplicateVisitor() *TestVisitorRepositoryCase {
	return &TestVisitorRepositoryCase{
		visitor: &domain.Visitor{
			UID: domain.UID("duplicate-visitor-456"),
			URL: domain.URL("https://example.com/page2"),
		},
	}
}

func (tb *TestVisitorRepositoryBuilder) WithMultipleUniqueVisitors() *TestVisitorRepositoryCase {
	return &TestVisitorRepositoryCase{
		visitors: []*domain.Visitor{
			{UID: domain.UID("visitor-1"), URL: domain.URL("https://example.com/page1")},
			{UID: domain.UID("visitor-2"), URL: domain.URL("https://example.com/page2")},
			{UID: domain.UID("visitor-3"), URL: domain.URL("https://example.com/page3")},
			{UID: domain.UID("visitor-4"), URL: domain.URL("https://example.com/page4")},
			{UID: domain.UID("visitor-5"), URL: domain.URL("https://example.com/page5")},
		},
	}
}

func (tb *TestVisitorRepositoryBuilder) WithConcurrentUniqueVisitors() *TestVisitorRepositoryCase {
	const numVisitors = 100
	visitors := make([]*domain.Visitor, numVisitors)

	for i := range numVisitors {
		visitors[i] = &domain.Visitor{
			UID: domain.UID(fmt.Sprintf("concurrent-visitor-%d", i)),
			URL: domain.URL(fmt.Sprintf("https://example.com/page-%d", i)),
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
			URL: domain.URL("https://example.com/same-page"),
		},
		numGoroutines: 10,
	}
}

func (tb *TestVisitorRepositoryBuilder) WithSameVisitorDifferentURLs() *TestVisitorRepositoryCase {
	return &TestVisitorRepositoryCase{
		visitors: []*domain.Visitor{
			{UID: domain.UID("same-visitor"), URL: domain.URL("https://example.com/page1")},
			{UID: domain.UID("same-visitor"), URL: domain.URL("https://example.com/page2")},
		},
	}
}
