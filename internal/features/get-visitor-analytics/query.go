package getvisitoranalytics

import (
	"context"

	"github.com/Falagan/web-tracker/internal/domain"
)

type GetVisitorAnalyticsQuery struct {
	URL string
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
	count, err := qh.ar.GetVisitedURLCount(ctx, domain.URL(q.URL))

	if err != nil {
		return nil, &GetAnalyticsError
	}

	return count, nil
}