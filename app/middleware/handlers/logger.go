package handlers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
)

// Logger outputs basic info about a handler
func Logger(handler http.Handler, name string) http.Handler {
	// Add gorilla logging
	handler = handlers.LoggingHandler(os.Stdout, handler)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		handler.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
