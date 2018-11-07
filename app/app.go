package app

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/nickhstr/go-web-service/app/router"
	"github.com/nickhstr/go-web-service/app/utils/env"
	log "github.com/sirupsen/logrus"
	// Setup for logging is done in an init
	_ "github.com/nickhstr/go-web-service/app/utils/log"
)

// Init initializes the server.
func Init() {
	if env.IsDev() {
		if err := godotenv.Load(); err != nil {
			log.Println("Could not load .env file")
		}
	}

	addr := env.GetAddr()
	router := router.New()

	log.Fatal(startServer(addr, router))
}

func startServer(address string, router http.Handler) error {
	fmt.Printf("Listening on %s, in %s mode\n", address, env.Get("GO_ENV", "development"))
	return http.ListenAndServe(address, router)
}
