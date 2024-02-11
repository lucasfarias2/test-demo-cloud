package main

import (
	"context"
	firebase "firebase.google.com/go"
	"fmt"
	"github.com/go-chi/chi/v5"
	"google.golang.org/api/option"
	"log"
	"net/http"
	"os"
	"packlify-cloud/db"
	"packlify-cloud/handlers"
	"packlify-cloud/handlers/api"
	"packlify-cloud/handlers/dashboard"
	"packlify-cloud/middleware"
	"packlify-cloud/utils"
)

func main() {
	router := chi.NewRouter()

	_ = utils.LoadEnv(".env")

	db.ConnectDatabase()

	isProd := os.Getenv("APP_ENV") == "production"

	utils.InitTemplates()

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

	if !isProd {
		router.Get("/ws", handlers.HotReloadHandler)
	}

	router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.Get("/", handlers.HomeHandler())
	router.With(middleware.RequireUserMiddleware).Get("/dashboard", handlers.DashboardHandler())
	router.With(middleware.RequireUserMiddleware).Get("/dashboard/projects", dashboard.ProjectsHandler())
	router.With(middleware.RequireUserMiddleware).Get("/dashboard/projects/new", dashboard.NewProjectHandler())
	router.With(middleware.RequireUserMiddleware).Get("/dashboard/org", dashboard.OrganizationHandler())
	router.With(middleware.RequireUserMiddleware).Get("/dashboard/org/new", dashboard.NewOrgHandler())
	router.With(middleware.RequireNoUserMiddleware).Get("/login", handlers.LoginHandler())
	router.With(middleware.RequireNoUserMiddleware).Get("/signup", handlers.SignupHandler())
	router.With(middleware.RequireUserMiddleware).Get("/logout", handlers.LogoutHandler())

	router.Post("/api/v1/login", api.LoginApiHandler(authClient))
	router.Post("/api/v1/organization", api.CreateOrganizationApiHandler)

	fmt.Println("Server running on port 8080")

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln("Error starting the server", err)
		return
	}
}
