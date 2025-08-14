package domain

import "context"

type URL string
type URLCount int

type Analytic struct {
	data map[URL]URLCount
}

type AnalyticRepository interface {
	IncreaseVisitedURLCount(ctx context.Context, url URL) error
	GetVisitedURLCount(ctx context.Context, url URL) (*URLCount, error)
}
