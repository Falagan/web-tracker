package ingestvisitor

import (
	"net/http"

	httpserver "github.com/Falagan/web-tracker/cmd/http-server"
)

type IngestVisitorController struct {
	Server         *httpserver.HTTPServer
	Mapper         *IngestVisitorMapper
	Validator      *IngestVisitorValidator
	CommandHandler *IngestVisitorsCommandHandler
}

func NewIngestVisitorController(s *httpserver.HTTPServer) *IngestVisitorController {
	return &IngestVisitorController{
		Server:         s,
		Mapper:         NewIngestVisitorMapper(),
		Validator:      NewIngestVisitorValidator(),
		CommandHandler: NewIngestVisitorsCommandHandler(s.VisitorRepository, s.AnalyticRepository),
	}
}

func (c *IngestVisitorController) MapEndpoint() {
	c.Server.Router.HandleFunc("/web-tracker/new-visitor", c.handler).Methods(http.MethodPost)
}

func (c *IngestVisitorController) handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	request, err := c.Mapper.MapToIngestVisitorRequest(r)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	err = c.Validator.ValidateRequest(request)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	command := c.Mapper.MapToCommand(request)
	err = c.CommandHandler.handle(ctx, command)

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	response := c.Mapper.MapToResponse()
	w.Write([]byte(response.message))
}
