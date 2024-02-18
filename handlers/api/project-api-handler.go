package api

import (
	"encoding/json"
	"log"
	"net/http"
	"packlify-cloud/models"
	"packlify-cloud/services"
	"strconv"
)

func CreateProjectApiHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Printf("Failed to parse form data: %v", err)
		http.Error(w, "Failed to process request", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	orgIdStr := r.FormValue("organization_id")
	orgId, err := strconv.Atoi(orgIdStr)
	if err != nil {
		log.Printf("Failed to convert organization_id to int: %v", err)
		http.Error(w, "Invalid organization ID", http.StatusBadRequest)
		return
	}

	toolkitIdStr := r.FormValue("toolkit_id")
	toolkitId, err := strconv.Atoi(toolkitIdStr)
	if err != nil {
		log.Printf("Failed to convert toolkit_id to int: %v", err)
		http.Error(w, "Invalid toolkit ID", http.StatusBadRequest)
		return

	}

	newProject := models.Project{
		Name:           name,
		OrganizationID: orgId,
		ToolkitID:      toolkitId,
	}

	project, err := services.CreateProject(newProject)
	if err != nil {
		log.Printf("Failed to create project: %v", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("HX-Redirect", "/dashboard/projects")

	if err := json.NewEncoder(w).Encode(project); err != nil {
		log.Printf("Failed to encode response data: %v", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
		return
	}
}
