package services

import (
	"packlify-cloud/db"
	"packlify-cloud/models"
)

func GetUserAccount(uuid string) (*models.Account, error) {
	database := db.GetDB()

	var account models.Account

	err := database.QueryRow("SELECT id, uuid FROM accounts WHERE uuid = $1", uuid).Scan(&account.ID, &account.UUID)
	if err != nil {
		return nil, err
	} else if account == (models.Account{}) {
		return nil, nil
	}

	return &account, nil
}

func CreateAccount(uuid string) (models.Account, error) {
	database := db.GetDB()

	var account models.Account

	err := database.QueryRow("INSERT INTO accounts(uuid) VALUES($1) RETURNING id, uuid", uuid).Scan(&account.ID, &account.UUID)
	if err != nil {
		return models.Account{}, err
	}

	return account, nil
}
