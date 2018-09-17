package routes

import (
	"net/http"

	"github.com/nickhstr/go-web-service/src/routes/handlers"
)

// Route defines the fundamental pieces of information
// required of every route.
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

// Routes registers all routes for the router
var Routes = []Route{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
}
