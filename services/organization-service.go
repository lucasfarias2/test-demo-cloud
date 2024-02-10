package services

import (
	"packlify-cloud/db"
	"packlify-cloud/models"
)

func CreateOrganization(newOrganizationrequest models.Organization) (models.Organization, error) {
	db := db.GetDB()

	var newOrganization models.Organization

	err := db.QueryRow("INSERT INTO organizations(name, admin_user_id) VALUES($1, $2) RETURNING id, name, admin_user_id", newOrganizationrequest.Name).Scan(&newOrganization.ID, &newOrganization.Name, &newOrganization.AdminUserId)
	if err != nil {
		return models.Organization{}, err
	}

	return newOrganization, nil
}
