package ingestvisitor

import (
	"encoding/json"
	"net/http"

	"github.com/Falagan/web-tracker/internal/domain"
	"github.com/Falagan/web-tracker/pkg"
)

type IngestVisitorRequest struct {
	UID string
	URL string
}

type IngestVisitorResponse struct {
	StatusCode int    `json:"status_code,omitempty"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error,omitempty"`
}

type IngestVisitorMapper struct{}

func NewIngestVisitorMapper() *IngestVisitorMapper {
	return &IngestVisitorMapper{}
}

func (m *IngestVisitorMapper) MapToIngestVisitorRequest(r *http.Request) (*IngestVisitorRequest, error) {
	var ivr *IngestVisitorRequest
	err := json.NewDecoder(r.Body).Decode(&ivr)

	if err != nil {
		return nil, &domain.VisitorInvalidRequest
	}

	return ivr, nil
}

func (m *IngestVisitorMapper) MapToCommand(r *IngestVisitorRequest) *IngestVisitorsCommand {
	return &IngestVisitorsCommand{
		UID: r.UID,
		URL: r.URL,
	}
}

func (m *IngestVisitorMapper) MapToSuccessResponse() *IngestVisitorResponse {
	return &IngestVisitorResponse{
		Message: "event ingested",
		IsError: false,
	}
}

func (m *IngestVisitorMapper) MapToErrorResponse(e error, statusCode int) *IngestVisitorResponse {
	return &IngestVisitorResponse{
		StatusCode: statusCode,
		Message:    pkg.ErrorMessage(e),
		IsError:    true,
	}
}
