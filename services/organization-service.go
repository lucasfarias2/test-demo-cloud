package services

import (
	"packlify-cloud/db"
	"packlify-cloud/models"
)

func CreateOrganization(orgReq models.Org, accID int) (models.Org, error) {
	database := db.GetDB()

	var newOrg models.Org

	err := database.QueryRow("INSERT INTO organizations(name) VALUES($1) RETURNING id, name", orgReq.Name).Scan(&newOrg.ID, &newOrg.Name)
	if err != nil {
		return models.Org{}, err
	}

	err = LinkToOrgAsAdmin(accID, newOrg.ID)
	if err != nil {
		return models.Org{}, err
	}

	return newOrg, nil
}

func LinkToOrgAsAdmin(accID, orgID int) error {
	database := db.GetDB()

	adminRole, _ := GetAdminRole()

	_, err := database.Exec("INSERT INTO account_organization(account_id, organization_id, role_id) VALUES($1, $2, $3)", accID, orgID, adminRole.ID)
	if err != nil {
		return err
	}

	return nil
}

func LinkToOrgAsMember(accID, orgID int) error {
	database := db.GetDB()

	memberRole, _ := GetMemberRole()

	_, err := database.Exec("INSERT INTO account_organization(account_id, organization_id, role_id) VALUES($1, $2, $3)", accID, orgID, memberRole.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetAccountLinkedOrganizations(accountID int) ([]models.Org, error) {
	database := db.GetDB()

	rows, err := database.Query(`
        SELECT o.id, o.name
        FROM organizations o
        JOIN account_organization ao ON o.id = ao.organization_id
        WHERE ao.account_id = $1`, accountID)
	if err != nil {
		return nil, err
	}

	var organizations []models.Org

	for rows.Next() {
		var org models.Org
		if err := rows.Scan(&org.ID, &org.Name); err != nil {
			return nil, err
		}
		organizations = append(organizations, org)
	}

	return organizations, nil
}

func GetOrganizationById(orgID int) (models.Org, error) {
	database := db.GetDB()

	var org models.Org

	err := database.QueryRow("SELECT id, name FROM organizations WHERE id = $1", orgID).Scan(&org.ID, &org.Name)
	if err != nil {
		return models.Org{}, err
	}

	return org, nil
}

func GetAccountRoleInOrganization(accountID, orgID int) (models.Role, error) {
	database := db.GetDB()

	var role models.Role

	err := database.QueryRow("SELECT r.id, r.name FROM roles r JOIN account_organization ao ON r.id = ao.role_id WHERE ao.account_id = $1 AND ao.organization_id = $2", accountID, orgID).Scan(&role.ID, &role.Name)
	if err != nil {
		return models.Role{}, err
	}

	return role, nil

}

func CheckAccountCanReadOrg(accountID, orgID int) (bool, error) {
	database := db.GetDB()

	var count int

	err := database.QueryRow("SELECT COUNT(*) FROM account_organization WHERE account_id = $1 AND organization_id = $2", accountID, orgID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func CheckAccountCanWriteOrg(accountID, orgID int) (bool, error) {
	database := db.GetDB()

	memberRole, _ := GetMemberRole()
	adminRole, _ := GetAdminRole()

	var count int

	err := database.QueryRow("SELECT COUNT(*) FROM account_organization WHERE account_id = $1 AND organization_id = $2 AND role_id IN ($3, $4)", accountID, orgID, memberRole.ID, adminRole.ID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func CheckAccountCanAdminOrg(orgID, accountID int) (bool, error) {
	database := db.GetDB()

	adminRole, _ := GetAdminRole()

	var count int

	err := database.QueryRow("SELECT COUNT(*) FROM account_organization WHERE account_id = $1 AND organization_id = $2 AND role_id = $3", accountID, orgID, adminRole.ID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetOrganizationByProjectId(projID int) (models.Org, error) {
	database := db.GetDB()

	var org models.Org

	err := database.QueryRow("SELECT o.id, o.name FROM organizations o JOIN projects p ON o.id = p.organization_id WHERE p.id = $1", projID).Scan(&org.ID, &org.Name)
	if err != nil {
		return models.Org{}, err
	}

	return org, nil
}
