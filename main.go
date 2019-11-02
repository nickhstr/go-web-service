package main

import (
	"github.com/nickhstr/go-web-service/routes"
	"github.com/nickhstr/goweb/env"
	"github.com/nickhstr/goweb/middleware"
	"github.com/nickhstr/goweb/server"
)

var GitCommit = "<not set>"

func main() {
	mux := middleware.Create(middleware.Config{
		AppName:    env.Get("APP_NAME", "web-service"),
		AppVersion: env.Get("APP_VERSION"),
		EnvVarsToValidate: []string{
			"APP_NAME",
			"GO_ENV",
		},
		GitRevision: GitCommit,
		Handler:     routes.Router,
		Region:      env.Get("REGION"),
	})

	server.Start(mux)
}
