package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func ConnectDatabase() {
	var err error
	var connectionString string
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if os.Getenv("APP_ENV") != "development" {
		log.Printf("Connecting to cloud database: host=%s", dbHost)
		connectionString = fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s", dbUser, dbPassword, dbName, dbHost, dbPort)
	} else {
		log.Printf("Connecting to local database: host=%s port=%s", dbHost, dbPort)
		connectionString = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", dbUser, dbName, dbPassword, dbHost, dbPort)
	}

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		log.Println("Database connection established")
	}
}

func GetDB() *sql.DB {
	return db
}
