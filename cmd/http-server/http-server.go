package httpserver

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	Server *http.Server
	Router *mux.Router
}

type HTTPServerConfig struct {
	Address      string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	IdleTimeout  time.Duration
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
		Router: router,
	}
	return server
}

func (s *HTTPServer) StartHTTPServerAsync() {
	go func() {
		err := s.startHTTPServer()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()
}

func (s *HTTPServer) startHTTPServer() error {
	log.Println("Starting server at " + s.Server.Addr)
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
		log.Println("Checking server health")
	}
	// endpoint
	s.Router.HandleFunc("/health", healthHandler).Methods("GET")
}
