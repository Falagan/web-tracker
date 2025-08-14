package ingestvisitor

type IngestVisitorValidator struct{}

func NewIngestVisitorValidator() *IngestVisitorValidator {
	return &IngestVisitorValidator{}
}

func (v *IngestVisitorValidator) ValidateRequest(r *IngestVisitorRequest) []error {
	validationErrors := []error{}

	if !validateUID(r.UID) {
		validationErrors = append(validationErrors, &InvalidUID)
	}

	if len(validationErrors) > 0 {
		return validationErrors
	}

	return nil
}

func validateUID(uid string) bool {
	//TODO: get more info about the uid type to be validated
	// For this POC we set default true validation
	return true
}
