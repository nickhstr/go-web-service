package routes

import (
	"github.com/go-chi/chi"
	"github.com/nickhstr/goweb/router"
)

// Router provides a chi router for registration of routes.
var Router *chi.Mux

// Routes holds all routes to be registered to Router.
var Routes = []router.Route{
	Index,
	Hello,
	IP,
}
