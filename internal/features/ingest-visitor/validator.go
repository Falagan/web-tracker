package ingestvisitor

import (
	"github.com/Falagan/web-tracker/internal/domain"
)

type IngestVisitorValidator struct{}

func NewIngestVisitorValidator() *IngestVisitorValidator {
	return &IngestVisitorValidator{}
}

func (v *IngestVisitorValidator) ValidateRequest(r *IngestVisitorRequest) []error {
	validationErrors := []error{}

	uid := domain.UID(r.UID)
	err := uid.Validate()
	if err != nil {
		validationErrors = append(validationErrors, err)
	}

	url, err := domain.NewURL(r.URL)

	if err != nil {
		validationErrors = append(validationErrors, err)
	}

	err = url.Validate()

	if err != nil {
		validationErrors = append(validationErrors, err)
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}
