package routes

import (
	"net/http"
	"github.com/nickhstr/goweb/router"
	"github.com/nickhstr/go-web-service/handlers"
)

// Routes maps all handlers to their appropriate Route
var Routes = []router.Route{
	router.Route{
		Method:  http.MethodGet,
		Path:    "/",
		Handler: handlers.Index,
	},
}

// Handler creates a new router with Routes.
func Handler() http.Handler {
	return router.New(Routes)
}
