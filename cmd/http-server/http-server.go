package httpserver

import (
	_ "embed"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Falagan/web-tracker/internal/domain"
	"github.com/Falagan/web-tracker/pkg"
	"github.com/gorilla/mux"
)

//go:embed openapi.yaml
var openAPISpec []byte

type HTTPServer struct {
	Server             *http.Server
	Router             *mux.Router
	VisitorRepository  domain.VisitorRepository
	AnalyticRepository domain.AnalyticRepository
	Observer           pkg.Observer
}

type HTTPServerConfig struct {
	Address            string
	Port               int
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleTimeout        time.Duration
	VisitorRepository  domain.VisitorRepository
	AnalyticRepository domain.AnalyticRepository
	Observer           pkg.Observer
}

func NewHTTPServer(sc *HTTPServerConfig) *HTTPServer {
	router := mux.NewRouter()
	server := &HTTPServer{
		Server: &http.Server{
			Addr:         sc.Address + ":" + strconv.Itoa(sc.Port),
			Handler:      router,
			ReadTimeout:  sc.ReadTimeout,
			WriteTimeout: sc.WriteTimeout,
			IdleTimeout:  sc.IdleTimeout,
		},
		Router:             router,
		VisitorRepository:  sc.VisitorRepository,
		AnalyticRepository: sc.AnalyticRepository,
		Observer:           sc.Observer,
	}
	return server
}

func (s *HTTPServer) StartHTTPServerAsync() {
	go func() {
		err := s.startHTTPServer()
		if err != nil && err != http.ErrServerClosed {
			s.Observer.Log(pkg.LogLevelError, "Failed to start server")
		}
	}()
}

func (s *HTTPServer) startHTTPServer() error {
	s.Observer.Log(pkg.LogLevelInfo, "Starting server at "+s.Server.Addr)
	return s.Server.ListenAndServe()
}

func (s *HTTPServer) WithHealthCheck() {
	// handler
	healthHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"status": "healthy",
			"time":   time.Now().Format(time.DateTime),
		})
		s.Observer.Log(pkg.LogLevelInfo, "Checking server health")
	}
	// endpoint
	s.Router.HandleFunc("/health", healthHandler).Methods("GET")
}

func (s *HTTPServer) WithOpenAPI() {
	// serve docs
	s.Router.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(OpenAPIHTML))
	}).Methods("GET")

	// serve open api spec
	s.Router.HandleFunc("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Write(openAPISpec)
	})
}
