package getvisitoranalytics

import (
	"net/http"

	"github.com/Falagan/web-tracker/internal/domain"
	"github.com/Falagan/web-tracker/pkg"
)

type GetVisitorAnalyticsRequest struct {
	URL string `json:"url"`
}

type GetVisitorAnalyticsResponse struct {
	StatusCode int    `json:"status_code,omitempty"`
	URL        string `json:"url,omitempty"`
	Count      int    `json:"unique_visitors,omitempty"`
	Message    string `json:"message,omitempty"`
	IsError    bool   `json:"is_error,omitempty"`
}

type GetVisitorAnalyticsMapper struct{}

func NewGetVisitorAnalyticsMapper() *GetVisitorAnalyticsMapper {
	return &GetVisitorAnalyticsMapper{}
}

func (m *GetVisitorAnalyticsMapper) MapToGetVisitorAnalyticsRequest(r *http.Request) (*GetVisitorAnalyticsRequest, error) {
	url := r.URL.Query().Get("url")
	if url == "" {
		return nil, &domain.URLInvalidFormatError
	}

	return &GetVisitorAnalyticsRequest{
		URL: url,
	}, nil
}

func (m *GetVisitorAnalyticsMapper) MapToQuery(r *GetVisitorAnalyticsRequest) (*GetVisitorAnalyticsQuery, error) {
	url, err := domain.NewURL(r.URL)

	if err != nil {
		return nil, &domain.URLInvalidFormatError
	}
	return &GetVisitorAnalyticsQuery{
		URL: url,
	}, nil
}

func (m *GetVisitorAnalyticsMapper) MapToDomain(url string, count int) (*domain.Analytic, error) {
	analytic, err := domain.NewAnalytic(url, count)

	if err != nil {
		return nil, err
	}

	return analytic, nil
}

func (m *GetVisitorAnalyticsMapper) MapToSuccessResponse(a *domain.Analytic) *GetVisitorAnalyticsResponse {
	path, _ := a.URL.GetPath()
	return &GetVisitorAnalyticsResponse{
		URL:     path,
		Count:   a.Count.ToInt(),
		IsError: false,
	}
}

func (m *GetVisitorAnalyticsMapper) MapToErrorResponse(e error, statusCode int) *GetVisitorAnalyticsResponse {
	return &GetVisitorAnalyticsResponse{
		StatusCode: statusCode,
		Message:    pkg.ErrorMessage(e),
		IsError:    true,
	}
}
