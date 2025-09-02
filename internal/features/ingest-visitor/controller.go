package ingestvisitor

import (
	"encoding/json"
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
		c.responseError(w, err, http.StatusBadRequest)
		return
	}

	validationErrors := c.Validator.ValidateRequest(request)
	if validationErrors != nil {
		lastError := validationErrors[len(validationErrors)-1]
		c.responseError(w, lastError, http.StatusBadRequest)
		return
	}

	command := c.Mapper.MapToCommand(request)
	err = c.CommandHandler.handle(ctx, command)

	if err != nil {
		c.responseError(w, err, http.StatusInternalServerError)
		return
	}

	response := c.Mapper.MapToSuccessResponse()
	c.responseSuccess(w, response)
}

func (c *IngestVisitorController) responseError(w http.ResponseWriter, e error, statusCode int) {
	errorResp := c.Mapper.MapToErrorResponse(e, statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorResp.StatusCode)
	json.NewEncoder(w).Encode(errorResp)
}

func (c *IngestVisitorController) responseSuccess(w http.ResponseWriter, r *IngestVisitorResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(r)
}
