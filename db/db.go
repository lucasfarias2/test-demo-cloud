package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/url"
	"os"
)

var db *sql.DB

func ConnectDatabase() {
	var err error
	var connectionString string
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := url.QueryEscape(os.Getenv("DB_PASSWORD"))
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUnixSocket := os.Getenv("DB_UNIX_SOCKET")

	if os.Getenv("APP_ENV") != "development" {
		log.Printf("Connecting to cloud database: host=/cloudsql/%s", dbUnixSocket)
		connectionString = fmt.Sprintf("user=%s password='%s' dbname=%s sslmode=disable host=/cloudsql/%s", dbUser, dbPassword, dbName, dbUnixSocket)
	} else {
		log.Printf("Connecting to database: host=%s port=%s", dbHost, dbPort)
		connectionString = fmt.Sprintf("user=%s dbname=%s password='%s' host=%s port=%s sslmode=disable", dbUser, dbName, dbPassword, dbHost, dbPort)
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
