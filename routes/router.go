package routes

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	mh "github.com/nickhstr/go-web-service/middleware/handlers"
)

// NewRouter returns a mux router with all Routes registered.
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range Routes {
		var handler http.Handler
		handler = route.HandlerFunc

		// Add our own logger middleware
		handler = mh.Logger(handler, route.Name)

		// Add gorilla logging
		handler = handlers.LoggingHandler(os.Stdout, handler)

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
