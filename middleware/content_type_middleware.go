package middleware

import "net/http"

// Env variables.
type Env struct {
	ContentType string
}

// NewContentType returns a pointer of a Env instance.
func NewContentType(contentType string) *Env {
	return &Env{contentType}
}

// Middleware returns a middleware handler.
func (env *Env) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", env.ContentType)
		next.ServeHTTP(w, r)
	})
}
