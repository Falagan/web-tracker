package ingestvisitor

import (
	"context"

	"github.com/Falagan/web-tracker/internal/domain"
)

type IngestVisitorsCommand struct {
	UID string
	URL string
}

type IngestVisitorsCommandHandler struct {
	vr domain.VisitorRepository
	ar domain.AnalyticRepository
}

func NewIngestVisitorsCommandHandler(vr domain.VisitorRepository, ar domain.AnalyticRepository) *IngestVisitorsCommandHandler {
	return &IngestVisitorsCommandHandler{
		vr: vr,
		ar: ar,
	}
}

func (cmh *IngestVisitorsCommandHandler) handle(ctx context.Context, c *IngestVisitorsCommand) error {
	v := &domain.Visitor{
		UID: c.UID,
		URL: c.URL,
	}
	err := cmh.vr.AddUnique(ctx, v)

	if err != nil {
		return &SaveIngestError
	}

	err = cmh.ar.IncreaseVisitedURLCount(ctx, domain.URL(c.URL))
	if err != nil {
		return &SaveIngestError
	}

	return nil
}
