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
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUnixSocket := os.Getenv("DB_UNIX_SOCKET")

	if os.Getenv("APP_ENV") != "development" {
		log.Printf("Connecting to cloud database: host=/cloudsql/%s", dbUnixSocket)
		connectionString = fmt.Sprintf("postgres://%s:%s@/cloudsql/%s/%s?sslmode=disable",
			url.QueryEscape(dbUser), url.QueryEscape(dbPassword), dbUnixSocket, dbName)
	} else {
		log.Printf("Connecting to database: host=%s port=%s", dbHost, dbPort)
		userInfo := url.UserPassword(dbUser, dbPassword)
		connectionString = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=disable",
			userInfo.String(), dbHost, dbPort, dbName)
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
