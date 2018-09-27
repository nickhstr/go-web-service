package middleware

import (
	"net/http"

	"github.com/gorilla/handlers"
)

// Compression middleware compresses all handler outputs
func Compression(handler http.Handler) http.Handler {
	// if env.Get("COMPRESSION") == "true" {
	// 	// Add gorilla compression middleware
	// 	handler = handlers.CompressHandler(handler)
	// }

	handler = handlers.CompressHandler(handler)

	return handler
}
