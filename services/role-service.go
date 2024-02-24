package services

import (
	"packlify-cloud/db"
	"packlify-cloud/models"
)

func GetAdminRole() (models.Role, error) {
	database := db.GetDB()

	var role models.Role

	err := database.QueryRow("SELECT id, name FROM roles WHERE name = 'admin'").Scan(&role.ID, &role.Name)
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func GetMemberRole() (models.Role, error) {
	database := db.GetDB()

	var role models.Role

	err := database.QueryRow("SELECT id, name FROM roles WHERE name = 'member'").Scan(&role.ID, &role.Name)
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}

func GetGuestRole() (models.Role, error) {
	database := db.GetDB()

	var role models.Role

	err := database.QueryRow("SELECT id, name FROM roles WHERE name = 'guest'").Scan(&role.ID, &role.Name)
	if err != nil {
		return models.Role{}, err
	}

	return role, nil
}
