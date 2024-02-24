package services

import (
	"log"
	"packlify-cloud/db"
	"packlify-cloud/models"
)

func CreateProject(projReq models.Project) (models.Project, error) {
	database := db.GetDB()

	var newProj models.Project

	err := database.QueryRow("INSERT INTO projects(name, organization_id, toolkit_id) VALUES($1, $2, $3) RETURNING id, name, organization_id", projReq.Name, projReq.OrganizationID, projReq.ToolkitID).Scan(&newProj.ID, &newProj.Name, &newProj.OrganizationID)
	if err != nil {
		return models.Project{}, err
	}

	return newProj, nil
}

func GetAccountProjects(accountID int) ([]models.ProjectView, error) {
	database := db.GetDB()

	rows, err := database.Query("SELECT p.id, p.name, p.toolkit_id, p.organization_id, o.name, t.name, t.image_url FROM projects p JOIN organizations o ON p.organization_id = o.id JOIN account_organization ao ON o.id = ao.organization_id JOIN toolkits t ON p.toolkit_id = t.id WHERE ao.account_id = $1", accountID)
	if err != nil {
		log.Println("Error getting account projects:", err)
		return nil, err
	}

	var projects []models.ProjectView

	for rows.Next() {
		var proj models.ProjectView
		if err := rows.Scan(&proj.ID, &proj.Name, &proj.ToolkitID, &proj.OrganizationID, &proj.OrgName, &proj.ToolkitName, &proj.ImageURL); err != nil {
			return nil, err
		}
		projects = append(projects, proj)
	}

	return projects, nil
}

func GetProjectById(projectID int) (models.Project, error) {
	database := db.GetDB()

	var proj models.Project

	err := database.QueryRow("SELECT id, name, organization_id, toolkit_id FROM projects WHERE id = $1", projectID).Scan(&proj.ID, &proj.Name, &proj.OrganizationID, &proj.ToolkitID)
	if err != nil {
		return models.Project{}, err
	}

	return proj, nil
}

func GetProjectDetails(projectID int) (models.ProjectView, error) {
	database := db.GetDB()

	var proj models.ProjectView

	err := database.QueryRow("SELECT p.id, p.name, p.organization_id, p.toolkit_id, o.name, t.name, t.image_url FROM projects p JOIN organizations o ON p.organization_id = o.id JOIN toolkits t ON p.toolkit_id = t.id WHERE p.id = $1", projectID).Scan(&proj.ID, &proj.Name, &proj.OrganizationID, &proj.ToolkitID, &proj.OrgName, &proj.ToolkitName, &proj.ImageURL)
	if err != nil {
		return models.ProjectView{}, err
	}

	return proj, nil
}
