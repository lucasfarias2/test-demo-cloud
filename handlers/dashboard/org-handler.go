package dashboard

import (
	"log"
	"net/http"
	"os"
	"packlify-cloud/handlers"
	"packlify-cloud/middleware"
	"packlify-cloud/utils"
)

func OrganizationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value("user").(*middleware.User)

		templates := utils.LoadTemplates()

		err := templates.ExecuteTemplate(w, "organization.gohtml", handlers.PageData{
			PageTitle:       "Your organization - Packlify",
			PageDescription: "Your organization in Packlify",
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
