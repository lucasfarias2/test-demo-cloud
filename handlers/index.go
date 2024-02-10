package handlers

import "cloud/middleware"

type PageData struct {
	PageTitle       string
	PageDescription string
	IsProd          bool
	User            *middleware.User
	FirebaseAPIKey  string
}
