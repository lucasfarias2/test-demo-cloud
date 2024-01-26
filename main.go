package main

import (
	"cloud/handlers"
	"cloud/utils"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Name       string
	PageTitle  string
	CurrentUrl string
	IsDev      bool
}

var templateFiles = []string{
	"./templates/views/index.html",
	"./templates/components/head.html",
	"./templates/components/navbar.html",
	"./templates/components/layout.html",
}

func main() {
	if err := utils.LoadEnv(".env"); err != nil {
		panic(err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	var tmpl *template.Template

	if os.Getenv("APP_ENV") == "production" {
		tmpl = template.Must(template.ParseFiles(templateFiles...))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("APP_ENV") != "production" {
			tmpl = template.Must(template.ParseFiles(templateFiles...))
		}

		err := tmpl.ExecuteTemplate(w, "layout.html", PageData{
			PageTitle:  "Packlify",
			CurrentUrl: r.URL.Path,
			IsDev:      os.Getenv("APP_ENV") != "production",
		})
		if err != nil {
			log.Println("Error:", err)
			return
		}
	})

	if os.Getenv("APP_ENV") != "production" {
		http.HandleFunc("/ws", handlers.HandleHotReloadWS)
	}

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Error starting the server")
		return
	}
}
