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

func OrganizationHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value("user").(*middleware.User)

		templates := utils.LoadTemplates()

		userAccount, _ := services.GetUserAccount(user.UID)

		organizations, _ := services.GetAccountLinkedOrganizations(userAccount.ID)

		err := templates.ExecuteTemplate(w, "organizations.gohtml", handlers.PageData{
			PageTitle:       "Your organization - Packlify",
			PageDescription: "Your organizations in Packlify",
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

func NewOrgHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, _ := r.Context().Value("user").(*middleware.User)

		templates := utils.LoadTemplates()

		err := templates.ExecuteTemplate(w, "new-org.gohtml", handlers.PageData{
			PageTitle:       "Create new organization - Packlify",
			PageDescription: "Your new organization in Packlify",
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
