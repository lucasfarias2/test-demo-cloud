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
	PageTitle string
	IsDev     bool
}

func main() {
	_ = utils.LoadEnv(".env")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	var tmpl *template.Template

	if os.Getenv("APP_ENV") == "production" {
		tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("APP_ENV") != "production" {
			tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
		}

		err := tmpl.ExecuteTemplate(w, "index.gohtml", PageData{
			PageTitle: "Packlify",
			IsDev:     os.Getenv("APP_ENV") != "production",
		})
		if err != nil {
			log.Println("Error:", err)
			return
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("APP_ENV") != "production" {
			tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
		}

		err := tmpl.ExecuteTemplate(w, "login.gohtml", PageData{
			PageTitle: "Login - Packlify",
			IsDev:     os.Getenv("APP_ENV") != "production",
		})
		if err != nil {
			log.Println("Error:", err)
			return
		}
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if os.Getenv("APP_ENV") != "production" {
			tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
		}

		err := tmpl.ExecuteTemplate(w, "register.gohtml", PageData{
			PageTitle: "New account - Packlify",
			IsDev:     os.Getenv("APP_ENV") != "production",
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
