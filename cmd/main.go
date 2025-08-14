package main

import (
	"time"

	_ "net/http/pprof"

	httpserver "github.com/Falagan/web-tracker/cmd/http-server"
	"github.com/Falagan/web-tracker/infra"
	getvisitoranalytics "github.com/Falagan/web-tracker/internal/features/get-visitor-analytics"
	ingestvisitor "github.com/Falagan/web-tracker/internal/features/ingest-visitor"
	"github.com/Falagan/web-tracker/pkg"
)

func main() {

	visitorRepository := infra.NewVisitorRepositoryInMemory()
	analiticRepository := infra.NewAnalyticRepositoryInMemory()

	config := &httpserver.HTTPServerConfig{
		Address:            "0.0.0.0",
		Port:               4001,
		ReadTimeout:        time.Second * 30,
		WriteTimeout:       time.Second * 30,
		IdleTimeout:        time.Second * 30,
		VisitorRepository:  visitorRepository,
		AnalyticRepository: analiticRepository,
		Observer:           pkg.NewConsoleObserver(),
	}

	server := httpserver.NewHTTPServer(config)

	ingestVisitorController := ingestvisitor.NewIngestVisitorController(server)
	ingestVisitorController.MapEndpoint()

	getVisitorAnalyticsController := getvisitoranalytics.NewGetVisitorAnalyticsController(server)
	getVisitorAnalyticsController.MapEndpoint()

	server.WithHealthCheck()
	server.WithOpenApi()
	server.StartHTTPServerAsync()

	select {}
}
