package handlers

import (
	"io"
	"net/http"
)

// Index handles requests to the root route
func Index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Hello World!")
}
