package dashboard

import (
	"log"
	"net/http"
	"os"
	"packlify-cloud/handlers"
	"packlify-cloud/middleware"
	"packlify-cloud/services"
	"packlify-cloud/utils"
)

func ProjectsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value("user").(*middleware.User)

		templates := utils.LoadTemplates()

		err := templates.ExecuteTemplate(w, "projects.gohtml", handlers.PageData{
			PageTitle:       "Your projects - Packlify",
			PageDescription: "Your projects in Packlify",
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

func NewProjectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value("user").(*middleware.User)

		templates := utils.LoadTemplates()

		organizations, _ := services.GetUserOrganizations(user.UID)

		err := templates.ExecuteTemplate(w, "new-project.gohtml", handlers.PageData{
			PageTitle:       "Create new project - Packlify",
			PageDescription: "Your new project in Packlify",
			IsProd:          os.Getenv("APP_ENV") == "production",
			User:            user,
			Organizations:   organizations,
		})
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
