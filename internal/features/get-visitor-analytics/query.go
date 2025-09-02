package getvisitoranalytics

import (
	"context"

	"github.com/Falagan/web-tracker/internal/domain"
)

type GetVisitorAnalyticsQuery struct {
	URL domain.URL
}

type GetVisitorAnalyticsQueryHandler struct {
	ar domain.AnalyticRepository
}

func NewGetVisitorAnalyticsQueryHandler(ar domain.AnalyticRepository) *GetVisitorAnalyticsQueryHandler {
	return &GetVisitorAnalyticsQueryHandler{
		ar: ar,
	}
}

func (qh *GetVisitorAnalyticsQueryHandler) handle(ctx context.Context, q *GetVisitorAnalyticsQuery) (*domain.URLCount, error) {
	path, err := q.URL.GetPath()

	if err != nil {
		return nil, &domain.URLInvalidFormatError
	}

	count, err := qh.ar.GetVisitedURLCount(ctx, path)

	if err != nil {
		return nil, &domain.AnalyticNoData
	}

	return count, nil
}
