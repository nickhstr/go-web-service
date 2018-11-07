package middleware

import (
	"net/http"
	"time"

	"github.com/nickhstr/go-web-service/app/utils"
	"github.com/nickhstr/go-web-service/app/utils/env"
	log "github.com/sirupsen/logrus"
)

// Used to wrap an http.ResponseWriter to capture the response's status code
type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (sr *statusRecorder) WriteHeader(code int) {
	sr.status = code
	sr.ResponseWriter.WriteHeader(code)
}

// Logger outputs general information about requests.
func Logger(handler http.Handler) http.Handler {
	if env.Get("LOGGING", "true") == "true" {
		handler = logHandler(handler)
	}

	return handler
}

func logHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Use 200 status code as default
		rw := &statusRecorder{w, http.StatusOK}
		start := time.Now()

		handler.ServeHTTP(rw, r)

		log.WithFields(log.Fields{
			"method":        r.Method,
			"url":           r.RequestURI,
			"host":          r.Host,
			"headers":       r.Header,
			"response-time": time.Since(start),
			"status":        rw.status,
			"app-name":      utils.App.Name(),
		}).Info()
	})
}
