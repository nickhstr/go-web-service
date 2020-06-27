package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nickhstr/goweb/write"
)

// Index handles requests to the root route
func Index(r *mux.Router) {
	r.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	write.OK(w)
}
