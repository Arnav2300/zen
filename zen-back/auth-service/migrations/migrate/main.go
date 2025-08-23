package main

import (
	_ "auth-service/migrations"
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found or failed to load, proceeding with existing environment variables")
	}
	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		log.Fatal("DB_URI environment variable not set")
	}
	db, err := sql.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()
	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatalf("goose up failed: %v", err)
	}
	log.Println("migrations applied successfully")
}
