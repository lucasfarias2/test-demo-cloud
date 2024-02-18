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

		userAccount, err := services.GetUserAccount(user.UID)
		if err != nil {
			log.Println("Can't get user account", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		projects, _ := services.GetAccountProjects(userAccount.ID)

		err = templates.ExecuteTemplate(w, "projects.gohtml", handlers.PageData{
			PageTitle:       "Your projects - Packlify",
			PageDescription: "Your projects in Packlify",
			IsProd:          os.Getenv("APP_ENV") == "production",
			User:            user,
			Projects:        projects,
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

		userAccount, err := services.GetUserAccount(user.UID)
		if err != nil {
			log.Println("Can't get user account", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		organizations, _ := services.GetAccountLinkedOrganizations(userAccount.ID)

		toolkits, _ := services.GetToolkits()

		err = templates.ExecuteTemplate(w, "new-project.gohtml", handlers.PageData{
			PageTitle:       "Create new project - Packlify",
			PageDescription: "Your new project in Packlify",
			IsProd:          os.Getenv("APP_ENV") == "production",
			User:            user,
			Organizations:   organizations,
			Toolkits:        toolkits,
		})
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
