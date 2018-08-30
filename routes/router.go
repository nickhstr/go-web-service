package routes

import (
	"github.com/gorilla/mux"
)

// NewRouter returns a mux router with all Routes registered.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range Routes {
		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
