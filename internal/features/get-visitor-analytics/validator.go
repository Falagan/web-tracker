package getvisitoranalytics

import "strings"

type GetVisitorAnalyticsValidator struct{}

func NewGetVisitorAnalyticsValidator() *GetVisitorAnalyticsValidator {
	return &GetVisitorAnalyticsValidator{}
}

func (gvav *GetVisitorAnalyticsValidator) ValidateRequest(r *GetVisitorAnalyticsRequest) error {
	if !validateURL(r.URL) {
		return &InvalidURL
	}
	return nil
}

func validateURL(url string) bool {
	return len(strings.TrimSpace(url)) > 0
}