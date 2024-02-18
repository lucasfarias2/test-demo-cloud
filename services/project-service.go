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

	rows, err := database.Query("SELECT p.id, p.name, p.toolkit_id, p.organization_id, o.name, t.name FROM projects p JOIN organizations o ON p.organization_id = o.id JOIN account_organization ao ON o.id = ao.organization_id JOIN toolkits t ON p.toolkit_id = t.id WHERE ao.account_id = $1", accountID)
	if err != nil {
		log.Println("Error getting account projects:", err)
		return nil, err
	}

	var projects []models.ProjectView

	for rows.Next() {
		var proj models.ProjectView
		if err := rows.Scan(&proj.ID, &proj.Name, &proj.ToolkitID, &proj.OrganizationID, &proj.OrgName, &proj.ToolkitName); err != nil {
			return nil, err
		}
		projects = append(projects, proj)
	}

	return projects, nil
}
