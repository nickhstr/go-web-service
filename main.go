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
		AppVersion: env.Get("APP_VERSION"),
		GitRevision: env.Get("GIT_COMMIT"),
		Handler: routes.Handler(),
		Region: env.Get("REGION"),
	})

	server.Start(mux)
}
