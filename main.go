package main

import (
	_ "github.com/nickhstr/goweb/logger"
	"github.com/nickhstr/goweb/env"
	"github.com/nickhstr/goweb/middleware"
	"github.com/nickhstr/goweb/server" // nolint: gotype
	"github.com/nickhstr/go-web-service/routes"
)

func main() {
	mux := middleware.Create(middleware.Config{
		AppName: env.AppName(),
		AppVersion: "1.0.0",
		GitRevision: "abc123",
		Handler: routes.Handler(),
		Region: env.Get("REGION"),
	})

	server.Start(mux)
}
