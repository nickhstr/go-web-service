package main

import (
	"github.com/nickhstr/go-web-service/routes"
	"github.com/nickhstr/goweb/env"
	"github.com/nickhstr/goweb/middleware"
	"github.com/nickhstr/goweb/server"
)

var gitCommit = "<not set>"
var appVersion = "<not set>"

func main() {
	mux := middleware.Create(middleware.Config{
		AppName:    env.Get("APP_NAME", "web-service"),
		AppVersion: appVersion,
		EnvVarsToValidate: []string{
			"APP_NAME",
			"GO_ENV",
		},
		Etag:        true,
		GitRevision: gitCommit,
		Handler:     routes.Router,
		Region:      env.Get("REGION"),
	})

	server.Start(mux)
}
