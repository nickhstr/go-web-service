package main

import (
	"log"
	"net/http"

	"github.com/nickhstr/go-web-service/routes"
)

func main() {
	router := routes.NewRouter()

	// Use localhost for dev
	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
