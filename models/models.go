package models

type Account struct {
	ID   int    `json:"id"`
	UUID string `json:"uuid"`
}

type Org struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Project struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	OrganizationID int    `json:"organization_id"`
	ToolkitID      int    `json:"toolkit_id,omitempty"`
	RepositoryURL  string `json:"repository_url,omitempty"`
}

type ProjectView struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	OrganizationID int    `json:"organization_id"`
	ToolkitID      int    `json:"toolkit_id,omitempty"`
	OrgName        string `json:"org_name"`
	ToolkitName    string `json:"toolkit_name,omitempty"`
	RepositoryURL  string `json:"repository_url,omitempty"`
}

type Toolkit struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	RepositoryURL string `json:"repository_url"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type AccountOrganization struct {
	ID             int `json:"id"`
	AccountID      int `json:"account_id"`
	OrganizationID int `json:"organization_id"`
	RoleID         int `json:"role_id"`
}
