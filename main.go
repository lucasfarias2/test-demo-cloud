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
	IsProd    bool
}

func main() {
	_ = utils.LoadEnv(".env")

	var isProd = os.Getenv("APP_ENV") == "production"

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	var tmpl *template.Template

	if isProd {
		tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !isProd {
			tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
		}

		err := tmpl.ExecuteTemplate(w, "index.gohtml", PageData{
			PageTitle: "Packlify",
			IsProd:    isProd,
		})
		if err != nil {
			log.Println("Error:", err)
			return
		}
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if isProd {
			tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
		}

		err := tmpl.ExecuteTemplate(w, "login.gohtml", PageData{
			PageTitle: "Login - Packlify",
			IsProd:    isProd,
		})
		if err != nil {
			log.Println("Error:", err)
			return
		}
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if isProd {
			tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
		}

		err := tmpl.ExecuteTemplate(w, "register.gohtml", PageData{
			PageTitle: "New account - Packlify",
			IsProd:    isProd,
		})
		if err != nil {
			log.Println("Error:", err)
			return
		}
	})

	if !isProd {
		http.HandleFunc("/ws", handlers.HandleHotReloadWS)
	}

	fmt.Println("Server running on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Error starting the server")
		return
	}
}
