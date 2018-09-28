package middleware

import "net/http"

// Middleware represents the function type for all middleware.
type Middleware func(http.Handler) http.Handler

// Compose adds middleware handlers to a given handler.
func Compose(handler http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}

	return handler
}
