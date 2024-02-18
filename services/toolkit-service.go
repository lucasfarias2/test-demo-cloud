package services

import (
	"packlify-cloud/db"
	"packlify-cloud/models"
)

func GetToolkits() ([]models.Toolkit, error) {
	database := db.GetDB()

	rows, err := database.Query("SELECT id, name, repository_url FROM toolkits")
	if err != nil {
		return nil, err
	}

	var toolkits []models.Toolkit

	for rows.Next() {
		var toolkit models.Toolkit
		if err := rows.Scan(&toolkit.ID, &toolkit.Name, &toolkit.RepositoryURL); err != nil {
			return nil, err
		}
		toolkits = append(toolkits, toolkit)
	}

	return toolkits, nil
}
