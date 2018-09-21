package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/nickhstr/go-web-service/app/routes"
	"github.com/nickhstr/go-web-service/app/utils/env"
)

func main() {
	if env.IsDev() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	PORT := env.GetPort()
	router := routes.NewRouter()

	init := func() error {
		fmt.Printf("Listening on port %s in %s mode\n", PORT, env.Get("GO_ENV", "development"))
		return http.ListenAndServe(PORT, router)
	}

	log.Fatal(init())
}
