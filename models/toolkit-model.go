package models

type Toolkit struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	RepositoryURL string `json:"repository_url"`
}
