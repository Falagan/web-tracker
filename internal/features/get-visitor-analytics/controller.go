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
		c.responseError(w, err, http.StatusBadRequest)
		return
	}

	err = c.Validator.ValidateRequest(request)
	if err != nil {
		c.responseError(w, err, http.StatusBadRequest)
		return
	}

	query, err := c.Mapper.MapToQuery(request)

	if err != nil {
		c.responseError(w, err, http.StatusInternalServerError)
		return
	}

	count, err := c.QueryHandler.handle(ctx, query)

	if err != nil {
		c.responseError(w, err, http.StatusBadRequest)
		return
	}

	analytic, err := c.Mapper.MapToDomain(request.URL, count.ToInt())

	if err != nil {
		c.responseError(w, err, http.StatusInternalServerError)
		return
	}

	response := c.Mapper.MapToSuccessResponse(analytic)
	c.responseSuccess(w, response)
}

func (c *GetVisitorAnalyticsController) responseError(w http.ResponseWriter, e error, statusCode int) {
	errorResp := c.Mapper.MapToErrorResponse(e, statusCode)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorResp.StatusCode)
	json.NewEncoder(w).Encode(errorResp)
}

func (c *GetVisitorAnalyticsController) responseSuccess(w http.ResponseWriter, r *GetVisitorAnalyticsResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(r)
}
