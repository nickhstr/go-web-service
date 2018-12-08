package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware" // nolint: gotype
)

// Router provides a chi router for registration of routes.
var Router *chi.Mux

func init() {
	Router = chi.NewRouter()
	Router.Use(middleware.StripSlashes)
	registerRoutes()
}

// Add route register functions here.
func registerRoutes() {
	Index()
	Hello()
}
