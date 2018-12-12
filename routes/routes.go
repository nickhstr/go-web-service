package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware" // nolint: gotype
)

// Router provides a chi router for registration of routes.
var Router *chi.Mux

// Routes holds all routes to be registered to Router.
var Routes = []func(*chi.Mux){
	Index,
	Hello,
}

func init() {
	Router = chi.NewRouter()
	Router.Use(middleware.StripSlashes)
	registerRoutes(Router)
}

func registerRoutes(r *chi.Mux) {
	for _, route := range Routes {
		route(r)
	}
}
