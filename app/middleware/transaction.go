package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nickhstr/go-web-service/app/types"
	"github.com/satori/go.uuid"
)

// Transaction adds a transaction uuid to each request's context
func Transaction(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := fmt.Sprintf("%s", uuid.NewV4())
		ctx := r.Context()
		ctx = context.WithValue(ctx, types.TransactionKey, id)
		r = r.WithContext(ctx)
		handler.ServeHTTP(w, r)
	})
}
