package handlers

import (
	"log"
	"net/http"
	"os"
	"packlify-cloud/utils"
)

func SignupHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templates := utils.LoadTemplates()

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
