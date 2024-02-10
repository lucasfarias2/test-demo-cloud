package handlers

import (
	"cloud/middleware"
	"html/template"
	"log"
	"net/http"
	"os"
)

func HomeHandler(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value("user").(*middleware.User)

		err := templates.ExecuteTemplate(w, "index.gohtml", PageData{
			PageTitle:       "Packlify",
			PageDescription: "Packlify is a cloud manager platform that allows you to automatically deploy your applications into your desired cloud provider.",
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
