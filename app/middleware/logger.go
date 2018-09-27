package middleware

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/handlers"
	"github.com/nickhstr/go-web-service/app/utils/env"
)

// Logger outputs general information about requests.
func Logger(handler http.Handler) http.Handler {
	if env.IsDev() {
		handler = logHandler(handler)
	}
	// Add gorilla logging
	handler = handlers.CombinedLoggingHandler(os.Stdout, handler)

	return handler
}

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		handler.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}