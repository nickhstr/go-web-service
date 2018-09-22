package handlers

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/nickhstr/go-web-service/app/utils/env"
)

// Compression middleware compresses all handler outputs
func Compression(handler http.Handler) http.Handler {
	if env.IsProd() {
		// Add gorilla compression middleware
		handler = handlers.CompressHandler(handler)
	}

	return handler
}
