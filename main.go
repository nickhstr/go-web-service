package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/nickhstr/go-web-service/routes"
	"github.com/nickhstr/go-web-service/utils/env"
)

func main() {
	if env.IsDevEnv() {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	PORT := env.GetPort()
	router := routes.NewRouter()

	init := func() error {
		fmt.Printf("Listening on port: %s\n", PORT)
		return http.ListenAndServe(PORT, router)
	}

	log.Fatal(init())
}
