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

	mux := middleware.Create(middleware.Config{
		EnvVarsToValidate: []string{
			"GO_ENV",
		},
		AppName:     "go-web-service",
		AppVersion:  appVersion,
		Auth:        true,
		Etag:        true,
		GitRevision: gitCommit,
		Region:      env.Get("REGION"),
		WhiteList: []string{
			`^/$`,
			`^/go-web-service/health$`,
			`^/debug/pprof.*`,
		},
		Handler: middleware.Compose(
			routes.Router,
			cors.Default().Handler,
		),
	})

	if err := server.Start(mux); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
