package main

import (
	"cloud/utils"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Name      string
	PageTitle string
}

var templateFiles = []string{
	"./templates/views/index.html",
	"./templates/components/head.html",
	"./templates/components/navbar.html",
}

func main() {
	if err := utils.LoadEnv(".env"); err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	var tmpl *template.Template

	if os.Getenv("APP_ENV") != "development" {
		tmpl = template.Must(template.ParseFiles(templateFiles...))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("APP_ENV") == "development" {
			tmpl = template.Must(template.ParseFiles(templateFiles...))
		}

		err := tmpl.Execute(w, PageData{
			Name:      "Cloud he",
			PageTitle: "Packlify",
		})

		if err != nil {
			return
		}
	})

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Error starting the server")
		return
	}
}
