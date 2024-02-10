package handlers

import (
	"cloud/middleware"
	"html/template"
	"log"
	"net/http"
	"os"
)

func DashboardHandler(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value("user").(*middleware.User)

		err := templates.ExecuteTemplate(w, "dashboard.gohtml", PageData{
			PageTitle:       "Dashboard - Packlify",
			PageDescription: "Your Packlify dashboard.",
			IsProd:          os.Getenv("APP_ENV") == "production",
			User:            user,
		})
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
