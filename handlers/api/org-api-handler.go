package api

import (
	"encoding/json"
	"log"
	"net/http"
	"packlify-cloud/middleware"
	"packlify-cloud/models"
	"packlify-cloud/services"
)

func CreateOrganizationApiHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Failed to parse form data: %v", err)
		http.Error(w, "Failed to process request", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")

	user := r.Context().Value("user").(*middleware.User)

	userAccount, _ := services.GetUserAccount(user.UID)

	newOrganization := models.Org{
		Name: name,
	}

	organization, err := services.CreateOrganization(newOrganization, userAccount.ID)
	if err != nil {
		log.Printf("Failed to create organization: %v", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("HX-Redirect", "/dashboard/org")

	if err := json.NewEncoder(w).Encode(organization); err != nil {
		log.Printf("Failed to encode response data: %v", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}
}
