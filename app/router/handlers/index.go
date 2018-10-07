package handlers

import (
	"net/http"
)

// Index handles requests to the root route
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}
