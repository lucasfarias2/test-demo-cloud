package models

type Org struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	AdminUserID string `json:"admin_user_id"`
}
