package handlers

import (
	"log"
	"net/http"
	"os"
	"packlify-cloud/middleware"
	"packlify-cloud/utils"
)

func DashboardHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value("user").(*middleware.User)

		templates := utils.LoadTemplates()

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
