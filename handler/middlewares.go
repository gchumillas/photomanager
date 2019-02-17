package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/gchumillas/photomanager/manager"
)

// JSONMiddleware returns a middleware handler.
func (env *Env) JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}

func (env *Env) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := ""
		items := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(items) > 1 {
			token = items[1]
		}

		if len(token) == 0 {
			httpError(w, unauthorizedError)
			return
		}

		u := &manager.User{AccessToken: token}
		if !u.ReadUserByToken(env.DB) {
			httpError(w, unauthorizedError)
			return
		}

		ctx := context.WithValue(r.Context(), contextAuthUser, u)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
