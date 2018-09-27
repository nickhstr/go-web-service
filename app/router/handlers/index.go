package handlers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index handles requests to the root route
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}
