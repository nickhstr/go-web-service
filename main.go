package main

import (
	"net/http"
	_ "net/http/pprof"

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
		Etag:        true,
		GitRevision: gitCommit,
		Region:      env.Get("REGION"),
		Handler: middleware.Compose(
			routes.Router,
			cors.Default().Handler,
		),
	})

	if env.Get("GO_ENV") == "debug" {
		go func() {
			if err := http.ListenAndServe(":6060", nil); err != nil {
				log.Fatal().Err(err).Msg("failed to start debug server")
			}
		}()
	}

	if err := server.Start(mux); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
