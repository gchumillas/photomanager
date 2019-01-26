package middleware

import "net/http"

type Env struct {
	ContentType string
}

func NewContentType(contentType string) *Env {
	return &Env{contentType}
}

func (env *Env) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", env.ContentType)
		next.ServeHTTP(w, r)
	})
}
