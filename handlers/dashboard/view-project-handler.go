package dashboard

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"packlify-cloud/handlers"
	"packlify-cloud/middleware"
	"packlify-cloud/models"
	"packlify-cloud/services"
	"packlify-cloud/utils"
	"strconv"
)

func ViewProjectHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projectId := chi.URLParam(r, "id")
		user, _ := r.Context().Value("user").(*middleware.User)

		templates := utils.LoadTemplates()

		userAccount, _ := services.GetUserAccount(user.UID)

		projectIdInt, err := strconv.Atoi(projectId)
		if err != nil {
			log.Printf("Failed to convert projId to int: %v", err)
			http.Error(w, "Failed to process request", http.StatusInternalServerError)
			return
		}

		org, err := services.GetOrganizationByProjectId(projectIdInt)
		if err != nil {
			log.Printf("Failed to get orgId: %v", err)
			http.Error(w, "Failed to process request", http.StatusInternalServerError)
			return
		}

		canUserReadOrg, err := services.CheckAccountCanReadOrg(userAccount.ID, org.ID)
		if err != nil || !canUserReadOrg {
			log.Printf("Failed to check if user can read org: %v", err)
			handlers.NotFoundHandler(w, r)
			return
		}

		project, err := services.GetProjectById(projectIdInt)
		if err != nil {
			log.Printf("Failed to get organization: %v", err)
			http.Error(w, "Failed to process request", http.StatusInternalServerError)
			return
		}

		toolkit, _ := services.GetToolkitById(project.ToolkitID)

		err = templates.ExecuteTemplate(w, "view-project.gohtml", handlers.PageData{
			PageTitle:       fmt.Sprintf("%s - Packlify", project.Name),
			PageDescription: "Your project in Packlify",
			IsProd:          os.Getenv("APP_ENV") == "production",
			User:            user,
			Project: models.ProjectView{
				ID:             project.ID,
				Name:           project.Name,
				ToolkitID:      project.ToolkitID,
				OrganizationID: project.OrganizationID,
				OrgName:        org.Name,
				ToolkitName:    toolkit.Name,
				ImageURL:       toolkit.ImageURL,
			},
		})
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
