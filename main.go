package main

import (
	"github.com/nickhstr/go-web-service/routes"
	"github.com/nickhstr/goweb/env"
	_ "github.com/nickhstr/goweb/logger"
	"github.com/nickhstr/goweb/middleware"
	"github.com/nickhstr/goweb/server"
)

func main() {
	mux := middleware.Create(middleware.Config{
		AppName:     env.AppName(),
		AppVersion:  env.Get("APP_VERSION"),
		GitRevision: env.Get("GIT_COMMIT"),
		Handler:     routes.Router,
		Region:      env.Get("REGION"),
	})

	server.Start(mux)
}
