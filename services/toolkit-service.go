package services

import (
	"packlify-cloud/db"
	"packlify-cloud/models"
)

func GetToolkits() ([]models.Toolkit, error) {
	database := db.GetDB()

	rows, err := database.Query("SELECT id, name, repository_url, image_url FROM toolkits")
	if err != nil {
		return nil, err
	}

	var toolkits []models.Toolkit

	for rows.Next() {
		var toolkit models.Toolkit
		if err := rows.Scan(&toolkit.ID, &toolkit.Name, &toolkit.RepositoryURL, &toolkit.ImageURL); err != nil {
			return nil, err
		}
		toolkits = append(toolkits, toolkit)
	}

	return toolkits, nil
}

func GetToolkitById(toolkitID int) (models.Toolkit, error) {
	database := db.GetDB()

	var toolkit models.Toolkit

	err := database.QueryRow("SELECT id, name, repository_url, image_url FROM toolkits WHERE id = $1", toolkitID).Scan(&toolkit.ID, &toolkit.Name, &toolkit.RepositoryURL, &toolkit.ImageURL)
	if err != nil {
		return models.Toolkit{}, err
	}

	return toolkit, nil
}
