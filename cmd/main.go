package main

import (
	"time"

	_ "net/http/pprof"

	httpserver "github.com/Falagan/web-tracker/cmd/http-server"
)

func main() {
	config := &httpserver.HTTPServerConfig{
		Address:      "0.0.0.0",
		Port:         4001,
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 30,
	}
	server := httpserver.NewHTTPServer(config)
	server.WithHealthCheck()
	server.StartHTTPServerAsync()
	
	select {}
}
