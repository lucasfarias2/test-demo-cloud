package services

import (
	"packlify-cloud/db"
	"packlify-cloud/models"
)

func CreateProject(projReq models.Project) (models.Project, error) {
	database := db.GetDB()

	var newProj models.Project

	err := database.QueryRow("INSERT INTO projects(name, organization_id) VALUES($1, $2) RETURNING id, organization_id", projReq.Name, projReq.OrganizationID).Scan(&newProj.ID, &newProj.Name, &newProj.OrganizationID)
	if err != nil {
		return models.Project{}, err
	}

	return newProj, nil
}

//
//func GetUserProjects(userID string) ([]models.Project, error) {
//	database := db.GetDB()
//
//	rows, err := database.Query("SELECT id, name, admin_user_id FROM organizations WHERE admin_user_id = $1", userID)
//	if err != nil {
//		return nil, err
//	}
//
//	var organizations []models.Org
//
//	for rows.Next() {
//		var org models.Org
//		if err := rows.Scan(&org.ID, &org.Name, &org.AdminUserID); err != nil {
//			return nil, err
//		}
//		organizations = append(organizations, org)
//	}
//
//	return organizations, nil
//}
