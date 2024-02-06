package main

import (
	"cloud/handlers"
	"cloud/handlers/api"
	"cloud/middleware"
	"cloud/utils"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"google.golang.org/api/option"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	PageTitle       string
	PageDescription string
	IsProd          bool
	FirebaseAPIKey  string
	User            middleware.User
}

func main() {
	_ = utils.LoadEnv(".env")

	var isProd = os.Getenv("APP_ENV") == "production"

	config := &firebase.Config{ProjectID: "packlify"}

	app, err := firebase.NewApp(context.Background(), config, option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))))
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	ctx := context.Background()

	authClient, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	var tmpl *template.Template

	if isProd {
		tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
	}

	http.Handle("/", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "HTTP method not accepted", http.StatusMethodNotAllowed)
			return
		}

		if !isProd {
			tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
		}

		user := r.Context().Value("user").(middleware.User)

		err := tmpl.ExecuteTemplate(w, "index.gohtml", PageData{
			PageTitle:       "Packlify",
			PageDescription: "Packlify is a cloud manager platform that allows you to automatically deploy your applications into your desired cloud provider.",
			IsProd:          isProd,
			User:            user,
		})
		if err != nil {
			log.Println("Error:", err)
			return
		}
	}), authClient, false))

	http.Handle("/dashboard", middleware.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "HTTP method not accepted", http.StatusMethodNotAllowed)
			return
		}

		if !isProd {
			tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
		}

		user := r.Context().Value("user").(middleware.User)

		err := tmpl.ExecuteTemplate(w, "dashboard.gohtml", PageData{
			PageTitle:       "Dashboard - Packlify",
			PageDescription: "Your Packlify dashboard.",
			IsProd:          isProd,
			User:            user,
		})
		if err != nil {
			log.Println("Error:", err)
			return
		}
	}), authClient, true))

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "HTTP method not accepted", http.StatusMethodNotAllowed)
			return
		}

		cookie, err := r.Cookie("session")
		if cookie != nil {
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}

		if !isProd {
			tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
		}

		err = tmpl.ExecuteTemplate(w, "login.gohtml", PageData{
			PageTitle:       "Login - Packlify",
			PageDescription: "Login to your account to access your Packlify dashboard.",
			IsProd:          isProd,
			FirebaseAPIKey:  os.Getenv("FIREBASE_API_KEY"),
		})
		if err != nil {
			log.Println("Error:", err)
			return
		}
	})

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "HTTP method not accepted", http.StatusMethodNotAllowed)
			return
		}

		cookie, err := r.Cookie("session")
		if cookie != nil {
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}

		if !isProd {
			tmpl = template.Must(template.ParseGlob("./templates/**/*.gohtml"))
		}

		err = tmpl.ExecuteTemplate(w, "register.gohtml", PageData{
			PageTitle:       "New account - Packlify",
			PageDescription: "Create your account to access your Packlify dashboard.",
			IsProd:          isProd,
			FirebaseAPIKey:  os.Getenv("FIREBASE_API_KEY"),
		})
		if err != nil {
			log.Println("Error:", err)
			return
		}
	})

	http.HandleFunc("/api/v1/login", api.HandleLogin(authClient))

	if !isProd {
		http.HandleFunc("/ws", handlers.HandleHotReloadWS)
	}

	fmt.Println("Server running on port 8080")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("Error starting the server")
		return
	}
}
