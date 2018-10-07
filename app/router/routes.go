package router

import (
	"net/http"

	"github.com/nickhstr/go-web-service/app/router/handlers"
)

// Route defines the fundamental pieces of information
// required of every route.
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// Routes registers all routes for the router
var Routes = []Route{
	Route{
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"GET",
		"/health",
		handlers.Health,
	},
}
