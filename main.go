package main

import (
	"strconv"

	"github.com/nickhstr/go-web-service/routes"
	"github.com/nickhstr/goweb/dnscache"
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
		Debug:       true,
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

	// enable DNS caching
	dnsTTL, _ := strconv.Atoi(env.Get("DNSCACHE_TTL", "300"))
	dnscache.Enable(dnsTTL)

	if err := server.StartNew(mux); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
