package dashboard

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os"
	"packlify-cloud/handlers"
	"packlify-cloud/middleware"
	"packlify-cloud/services"
	"packlify-cloud/utils"
	"strconv"
)

func ViewOrgHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgId := chi.URLParam(r, "id")
		user, _ := r.Context().Value("user").(*middleware.User)

		templates := utils.LoadTemplates()

		userAccount, _ := services.GetUserAccount(user.UID)

		orgIdInt, err := strconv.Atoi(orgId)
		if err != nil {
			log.Printf("Failed to convert orgId to int: %v", err)
			http.Error(w, "Failed to process request", http.StatusInternalServerError)
			return
		}

		canUserReadOrg, err := services.CheckAccountCanReadOrg(userAccount.ID, orgIdInt)
		if err != nil || !canUserReadOrg {
			log.Printf("Failed to check if user can read org: %v", err)
			handlers.NotFoundHandler(w, r)
			return
		}

		organization, err := services.GetOrganizationById(orgIdInt)
		if err != nil {
			log.Printf("Failed to get organization: %v", err)
			http.Error(w, "Failed to process request", http.StatusInternalServerError)
			return
		}

		err = templates.ExecuteTemplate(w, "view-org.gohtml", handlers.PageData{
			PageTitle:       fmt.Sprintf("%s - Packlify", organization.Name),
			PageDescription: "Your organization in Packlify",
			IsProd:          os.Getenv("APP_ENV") == "production",
			User:            user,
			Organization:    organization,
		})
		if err != nil {
			log.Println("Error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
