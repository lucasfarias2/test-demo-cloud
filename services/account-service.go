package services

import (
	"packlify-cloud/db"
	"packlify-cloud/models"
)

func GetUserAccount(uuid string) (models.Account, error) {
	database := db.GetDB()

	row, err := database.Query("SELECT id, uuid FROM accounts WHERE uuid = $1 LIMIT 1", uuid)
	if err != nil {
		return models.Account{}, err
	}

	var account models.Account

	var acc models.Account

	for row.Next() {
		if err := row.Scan(&acc.ID, &acc.UUID); err != nil {
			return models.Account{}, err
		}
		account = acc
	}

	return account, nil
}
