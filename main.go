package main

import (
	"strconv"

	"github.com/nickhstr/go-web-service/routes"
	"github.com/nickhstr/goweb/dnscache"
	"github.com/nickhstr/goweb/env"
	"github.com/nickhstr/goweb/logger"
	"github.com/nickhstr/goweb/router"
	"github.com/nickhstr/goweb/server"
)

var gitCommit = "<not set>"
var appVersion = "<not set>"

func main() {
	name := "go-web-service"
	region := env.Get("REGION")
	log := logger.New(name)

	if err := env.ValidateEnvVars([]string{
		"GO_ENV",
	}); err != nil {
		log.Fatal().Err(err).Msg("Invalid environment variables")
	}

	h := router.DefaultMux(routes.Routes, router.DefaultOptions{
		Auth:      true,
		GitCommit: gitCommit,
		Name:      name,
		Region:    region,
		Version:   appVersion,
	})

	// enable DNS caching
	dnsTTL, _ := strconv.Atoi(env.Get("DNSCACHE_TTL", "300"))
	dnscache.Enable(dnsTTL)

	if err := server.StartNew(h); err != nil {
		log.Fatal().Err(err).Msg("failed to start server")
	}
}
