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
	v, err := domain.NewVisitor(c.UID, c.URL)
	if err != nil {
		return err
	}
	
	err = cmh.vr.AddUnique(ctx, v)
	if err != nil {
		return err
	}

	err = cmh.ar.IncreaseVisitedURLCount(ctx, v.URL)
	if err != nil {
		return err
	}

	return nil
}
