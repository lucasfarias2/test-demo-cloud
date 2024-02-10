package middleware

import (
	"html/template"
	"net/http"
)

func TemplateLoader(templates *template.Template) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			templates = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
			next.ServeHTTP(w, r)
		})
	}
}
