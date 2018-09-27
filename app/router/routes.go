package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/nickhstr/go-web-service/app/router/handlers"
)

// Route defines the fundamental pieces of information
// required of every route.
type Route struct {
	Name    string
	Method  string
	Path    string
	Handler httprouter.Handle
}

// Routes registers all routes for the router
var Routes = []Route{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"Health",
		"GET",
		"/health",
		handlers.Health,
	},
}
