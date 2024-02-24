package handlers

import (
	"log"
	"net/http"
	"os"
	"packlify-cloud/middleware"
	"packlify-cloud/utils"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	user, _ := r.Context().Value("user").(*middleware.User)

	templates := utils.LoadTemplates()

	err := templates.ExecuteTemplate(w, "notfound.gohtml", PageData{
		PageTitle:       "Not Found - Packlify",
		PageDescription: "Page not found in Packlify",
		IsProd:          os.Getenv("APP_ENV") == "production",
		User:            user,
	})
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
