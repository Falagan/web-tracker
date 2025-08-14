package ingestvisitor

import (
	"encoding/json"
	"net/http"
)

type IngestVisitorRequest struct {
	UID string
	URL string
}

type IngestVisitorRsponse struct {
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
		return nil, &InvalidEvent
	}

	return ivr, nil
}

func (m *IngestVisitorMapper) MapToCommand(r *IngestVisitorRequest) *IngestVisitorsCommand {
	return &IngestVisitorsCommand{
		UID: r.UID,
		URL: r.URL,
	}
}

func (m *IngestVisitorMapper) MapToResponse() *IngestVisitorRsponse {
	return &IngestVisitorRsponse{
		message: "event ingested",
	}
}
