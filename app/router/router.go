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

	for _, route := range Routes {
		baseRouter.Handle(route.Method, route.Path, route.Handler)
	}

	router = middleware.Compose(
		baseRouter,
		middleware.Transaction,
		middleware.Logger,
		middleware.Compression,
	)

	return router
}
