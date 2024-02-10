package handlers

import (
	"packlify-cloud/middleware"
	"packlify-cloud/models"
)

type PageData struct {
	PageTitle       string
	PageDescription string
	IsProd          bool
	User            *middleware.User
	FirebaseAPIKey  string
	Organizations   []models.Org
}
