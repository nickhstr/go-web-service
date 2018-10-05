package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/nickhstr/go-web-service/app/router"
	"github.com/nickhstr/go-web-service/app/utils/env"
)

// Init initializes the server.
func Init() {
	if env.IsDev() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
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
