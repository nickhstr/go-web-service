package routes

import (
	"github.com/go-chi/chi"
)

// Router provides a chi router for registration of routes.
var Router *chi.Mux

// Routes holds all routes to be registered to Router.
var Routes = []func(*chi.Mux){
	Debug,
	Index,
	Hello,
	IP,
}

func init() {
	Router = chi.NewRouter()
	registerRoutes(Router)
}

func registerRoutes(r *chi.Mux) {
	for _, route := range Routes {
		route(r)
	}
}
