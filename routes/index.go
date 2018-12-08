package routes

import (
	"net/http"
)

// Index handles requests to the root route
func Index() {
	Router.Get("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}
