package services

import (
	"packlify-cloud/db"
	"packlify-cloud/models"
)

func CreateOrganization(orgReq models.Org) (models.Org, error) {
	database := db.GetDB()

	var newOrg models.Org

	err := database.QueryRow("INSERT INTO organizations(name, admin_user_id) VALUES($1, $2) RETURNING id, name, admin_user_id", orgReq.Name, orgReq.AdminUserID).Scan(&newOrg.ID, &newOrg.Name, &newOrg.AdminUserID)
	if err != nil {
		return models.Org{}, err
	}

	return newOrg, nil
}

func GetUserOrganizations(userID string) ([]models.Org, error) {
	database := db.GetDB()

	rows, err := database.Query("SELECT id, name, admin_user_id FROM organizations WHERE admin_user_id = $1", userID)
	if err != nil {
		return nil, err
	}

	var organizations []models.Org

	for rows.Next() {
		var org models.Org
		if err := rows.Scan(&org.ID, &org.Name, &org.AdminUserID); err != nil {
			return nil, err
		}
		organizations = append(organizations, org)
	}

	return organizations, nil
}
