package middleware

import (
	"net/http"
)

func RequireUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("user") == nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("HX-Redirect", "/login")

			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequireNoUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Context().Value("user") != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("HX-Redirect", "/dashboard")

			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
