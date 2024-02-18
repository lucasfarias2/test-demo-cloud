package models

type Project struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	OrganizationID int    `json:"organization_id"`
	ToolkitID      int    `json:"toolkit_id,omitempty"`
}

type ProjectView struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	OrganizationID int    `json:"organization_id"`
	ToolkitID      int    `json:"toolkit_id,omitempty"`
	OrgName        string `json:"org_name"`
	ToolkitName    string `json:"toolkit_name,omitempty"`
}
