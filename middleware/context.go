package middleware

import (
	"context"
	"net/http"
	"time"
)

func ContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a context with a 5-second timeout
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		// Pass the context to the request handler function
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
