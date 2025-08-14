package main

import (
	"time"

	_ "net/http/pprof"

	"github.com/Falagan/web-tracker/cmd/envs"
	httpserver "github.com/Falagan/web-tracker/cmd/http-server"
	"github.com/Falagan/web-tracker/infra"
	getvisitoranalytics "github.com/Falagan/web-tracker/internal/features/get-visitor-analytics"
	ingestvisitor "github.com/Falagan/web-tracker/internal/features/ingest-visitor"
	"github.com/Falagan/web-tracker/pkg"
)

func main() {

	env := envs.NewEnv()

	visitorRepository := infra.NewVisitorRepositoryInMemoryBloom(1000, 0.01)
	analiticRepository := infra.NewAnalyticRepositoryInMemory()

	config := &httpserver.HTTPServerConfig{
		Address:            env.ServerAddress,
		Port:               env.ServerPort,
		ReadTimeout:        time.Second * 30,
		WriteTimeout:       time.Second * 30,
		IdleTimeout:        time.Second * 30,
		VisitorRepository:  visitorRepository,
		AnalyticRepository: analiticRepository,
		Observer:           pkg.NewConsoleObserver(),
		Env:                env.AppEnv,
	}

	server := httpserver.NewHTTPServer(config)

	ingestVisitorController := ingestvisitor.NewIngestVisitorController(server)
	ingestVisitorController.MapEndpoint()

	getVisitorAnalyticsController := getvisitoranalytics.NewGetVisitorAnalyticsController(server)
	getVisitorAnalyticsController.MapEndpoint()

	server.WithHealthCheck()
	server.WithOpenAPI()
	server.StartHTTPServerAsync()
	server.WithShutdownGracefully()
}
