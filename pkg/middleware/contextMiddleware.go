package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		ctx := r.Context()
		ctx = context.WithValue(ctx, "requestID", uuid.New().String())

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
