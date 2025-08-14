package getvisitoranalytics

import (
	"encoding/json"
	"net/http"

	httpserver "github.com/Falagan/web-tracker/cmd/http-server"
)

type GetVisitorAnalyticsController struct {
	Server       *httpserver.HTTPServer
	Mapper       *GetVisitorAnalyticsMapper
	Validator    *GetVisitorAnalyticsValidator
	QueryHandler *GetVisitorAnalyticsQueryHandler
}

func NewGetVisitorAnalyticsController(s *httpserver.HTTPServer) *GetVisitorAnalyticsController {
	return &GetVisitorAnalyticsController{
		Server:       s,
		Mapper:       NewGetVisitorAnalyticsMapper(),
		Validator:    NewGetVisitorAnalyticsValidator(),
		QueryHandler: NewGetVisitorAnalyticsQueryHandler(s.AnalyticRepository),
	}
}

func (c *GetVisitorAnalyticsController) MapEndpoint() {
	c.Server.Router.HandleFunc("/web-tracker/analytics", c.handler).Methods(http.MethodGet)
}

func (c *GetVisitorAnalyticsController) handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	request, err := c.Mapper.MapToGetVisitorAnalyticsRequest(r)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	err = c.Validator.ValidateRequest(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	query := c.Mapper.MapToQuery(request)
	count, err := c.QueryHandler.handle(ctx, query)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	analytic, err := c.Mapper.MapToDomain(request.URL, count.ToInt())

	response := c.Mapper.MapToResponse(analytic)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
