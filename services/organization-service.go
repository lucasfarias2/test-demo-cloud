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

func LinkAccountWithOrganization(accountID, organizationID int) error {
	database := db.GetDB()

	_, err := database.Exec("INSERT INTO account_organization(account_id, organization_id) VALUES($1, $2)", accountID, organizationID)
	if err != nil {
		return err
	}

	return nil
}

func GetAccountLinkedOrganizations(accountID int) ([]models.Org, error) {
	database := db.GetDB()

	rows, err := database.Query("SELECT o.id, o.name, o.admin_user_id FROM organizations o JOIN account_organization ao ON o.id = ao.organization_id WHERE ao.account_id = $1", accountID)
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
