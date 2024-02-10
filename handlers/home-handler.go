package handlers

import (
	"log"
	"net/http"
	"os"
	"packlify-cloud/middleware"
	"packlify-cloud/utils"
)

func HomeHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value("user").(*middleware.User)
		isProd := os.Getenv("APP_ENV") == "production"

		templates := utils.LoadTemplates()

		err := templates.ExecuteTemplate(w, "index.gohtml", PageData{
			PageTitle:       "Packlify",
			PageDescription: "Packlify is a cloud manager platform that allows you to automatically deploy your applications into your desired cloud provider.",
			IsProd:          isProd,
			User:            user,
		})
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
