package handler

import (
	"context"
	"net/http"
)

// JSONMiddleware returns a middleware handler.
func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := getParam(r, "token", "")
		if len(token) == 0 {
			httpError(w, unauthorizedError)
			return
		}

		ctx := context.WithValue(r.Context(), "auth-token", token)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
