package main

import (
	"cloud/handlers"
	"cloud/handlers/api"
	"cloud/middleware"
	"cloud/utils"
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/go-chi/chi/v5"
	"google.golang.org/api/option"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	router := chi.NewRouter()

	_ = utils.LoadEnv(".env")

	isProd := os.Getenv("APP_ENV") == "production"

	config := &firebase.Config{
		ProjectID: "packlify",
	}

	app, err := firebase.NewApp(context.Background(), config, option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))))
	if err != nil {
		log.Fatalf("Error initializing app: %v\n", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("Error getting Auth client: %v\n", err)
	}

	router.Use(middleware.GetUserMiddleware(authClient))

	templates := template.Must(template.ParseGlob("./templates/**/*.gohtml"))

	if !isProd {
		router.Use(middleware.TemplateLoader(templates))
		router.Get("/ws", handlers.HotReloadHandler)
	}

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.Get("/", handlers.HomeHandler(templates))
	router.With(middleware.RequireUserMiddleware).Get("/dashboard", handlers.DashboardHandler(templates))
	router.With(middleware.RequireNoUserMiddleware).Get("/login", handlers.LoginHandler(templates))
	router.With(middleware.RequireNoUserMiddleware).Get("/signup", handlers.SignupHandler(templates))
	router.With(middleware.RequireUserMiddleware).Get("/logout", handlers.LogoutHandler())

	router.Post("/api/v1/login", api.LoginApiHandler(authClient))

	fmt.Println("Server running on port 8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("Error starting the server")
		return
	}
}
