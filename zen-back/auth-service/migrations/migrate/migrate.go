package main

import (
	_ "auth-service/migrations"
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	var direction string
	flag.StringVar(&direction, "direction", "up", "Migration directions: up, down or reset")
	flag.Parse()

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
	switch direction {
	case "up":
		if err := goose.Up(db, "migrations"); err != nil {
			log.Fatalf("goose up failed: %v", err)
		}
		log.Println("migrations applied successfully")
	case "down":
		if err := goose.Down(db, "migrations"); err != nil {
			log.Fatalf("goose down failed: %v", err)
		}
		log.Println("migrations rolled back successfully")
	case "reset":
		if err := goose.Reset(db, "migrations"); err != nil {
			log.Fatalf("goose reset failed: %v", err)
		}
		log.Println("all migrations rolled back successfully")
	default:
		log.Fatalf("invalid direction: %s (use 'up', 'down' or 'reset')", direction)

	}
}
