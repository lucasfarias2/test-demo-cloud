package utils

import (
	"html/template"
	"os"
)

var templates *template.Template

func InitTemplates() {
	if os.Getenv("APP_ENV") == "production" {
		templates = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
	}
}

func LoadTemplates() *template.Template {
	if os.Getenv("APP_ENV") != "production" {
		return template.Must(template.ParseGlob("./templates/**/*.gohtml"))
	}
	return templates
}
