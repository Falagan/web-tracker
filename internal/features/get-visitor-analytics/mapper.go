package getvisitoranalytics

import (
	"net/http"

	"github.com/Falagan/web-tracker/internal/domain"
)

type GetVisitorAnalyticsRequest struct {
	URL string `json:"url"`
}

type GetVisitorAnalyticsResponse struct {
	URL   string `json:"url"`
	Count int    `json:"count"`
}

type GetVisitorAnalyticsMapper struct{}

func NewGetVisitorAnalyticsMapper() *GetVisitorAnalyticsMapper {
	return &GetVisitorAnalyticsMapper{}
}

func (m *GetVisitorAnalyticsMapper) MapToGetVisitorAnalyticsRequest(r *http.Request) (*GetVisitorAnalyticsRequest, error) {
	url := r.URL.Query().Get("url")
	if url == "" {
		return nil, &InvalidURL
	}

	return &GetVisitorAnalyticsRequest{
		URL: url,
	}, nil
}

func (m *GetVisitorAnalyticsMapper) MapToQuery(r *GetVisitorAnalyticsRequest) *GetVisitorAnalyticsQuery {
	return &GetVisitorAnalyticsQuery{
		URL: r.URL,
	}
}

func (m *GetVisitorAnalyticsMapper) MapToResponse(url string, count *domain.URLCount) *GetVisitorAnalyticsResponse {
	return &GetVisitorAnalyticsResponse{
		URL:   url,
		Count: int(*count),
	}
}
