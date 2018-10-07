package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/nickhstr/go-web-service/app/middleware"
)

// New returns a new router
func New() http.Handler {
	var router http.Handler
	baseRouter := httprouter.New()

	// Register routes with router
	for _, route := range Routes {
		handler := http.HandlerFunc(route.Handler)

		// Middleware can also be included here, on a per-route-basis

		baseRouter.Handler(route.Method, route.Path, handler)
	}

	// Add middleware for all requsts
	router = middleware.Compose(
		baseRouter,
		middleware.Logger,
		middleware.Transaction,
		middleware.Compression,
	)

	return router
}
