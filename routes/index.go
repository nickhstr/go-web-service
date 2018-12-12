package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Index handles requests to the root route
func Index(r *chi.Mux) {
	r.Get("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}
