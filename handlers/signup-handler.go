package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func SignupHandler(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := templates.ExecuteTemplate(w, "signup.gohtml", PageData{
			PageTitle:       "New account - Packlify",
			PageDescription: "Create your account to access your Packlify dashboard.",
			IsProd:          os.Getenv("APP_ENV") == "production",
			FirebaseAPIKey:  os.Getenv("FIREBASE_API_KEY"),
		})
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
