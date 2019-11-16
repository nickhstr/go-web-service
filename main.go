package main

import (
	"github.com/nickhstr/go-web-service/routes"
	"github.com/nickhstr/goweb/env"
	"github.com/nickhstr/goweb/logger"
	"github.com/nickhstr/goweb/middleware"
	"github.com/nickhstr/goweb/server"
	"github.com/rs/cors"
)

var gitCommit = "<not set>"
var appVersion = "<not set>"

func main() {
	log := logger.New("go-web-service")

	handler := middleware.Compose(
		routes.Router,
		// Add additional middleware here
		cors.Default().Handler,
	)
	mux := middleware.Create(middleware.Config{
		AppName:    "go-web-service",
		AppVersion: appVersion,
		EnvVarsToValidate: []string{
			"GO_ENV",
		},
		Etag:        true,
		GitRevision: gitCommit,
		Handler:     handler,
		Region:      env.Get("REGION"),
	})

	if err := server.Start(mux); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
