package ingestvisitor

import (
	"encoding/json"
	"net/http"

	"github.com/Falagan/web-tracker/internal/domain"
)

type IngestVisitorRequest struct {
	UID string
	URL string
}

type IngestVisitorResponse struct {
	message string
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

func (m *IngestVisitorMapper) MapToResponse() *IngestVisitorResponse {
	return &IngestVisitorResponse{
		message: "event ingested",
	}
}
