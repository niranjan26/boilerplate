package middlewares

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"context"

	"github.com/google/uuid"
)

func PopulateContextMiddleware() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header
			reqID := header.Get("x-trace-id")
			if reqID == "" {
				reqID = uuid.New().String()
			}

			log.Print("new request ID:", reqID)
			ctx := context.WithValue(context.Background(), "x-trace-id", reqID)
			ctx = context.WithValue(ctx, "vars", mux.Vars(r))
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
