package getvisitoranalytics

import "github.com/Falagan/web-tracker/internal/domain"

type GetVisitorAnalyticsValidator struct{}

func NewGetVisitorAnalyticsValidator() *GetVisitorAnalyticsValidator {
	return &GetVisitorAnalyticsValidator{}
}

func (v *GetVisitorAnalyticsValidator) ValidateRequest(r *GetVisitorAnalyticsRequest) error {
	err := domain.ValidateURL(r.URL)

	if err != nil {
		return err
	}
	return nil
}
