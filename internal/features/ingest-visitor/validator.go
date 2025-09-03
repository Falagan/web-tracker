package ingestvisitor

import (
	"github.com/Falagan/web-tracker/internal/domain"
)

type IngestVisitorValidator struct{}

func NewIngestVisitorValidator() *IngestVisitorValidator {
	return &IngestVisitorValidator{}
}

func (v *IngestVisitorValidator) ValidateRequest(r *IngestVisitorRequest) []error {
	var validationErrors []error

	if err := domain.ValidateUID(r.UID); err != nil {
		validationErrors = append(validationErrors, err)
	}

	if err := domain.ValidateURL(r.URL); err != nil {
		validationErrors = append(validationErrors, err)
	}

	return validationErrors
}
