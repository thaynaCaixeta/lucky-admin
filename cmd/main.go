package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	migrate "github.com/rubenv/sql-migrate"
	postgres "github.com/thaynaCaixeta/lucky-admin/database"
	"github.com/thaynaCaixeta/lucky-admin/database/migrations"
)

func main() {
	db, err := postgres.Connect()
	if err != nil {
		log.Fatalf("Failed to init app: %v", err)
	}
	// Run migrations
	runMigrations(db)

	defer db.Close()
}

func runMigrations(db *sqlx.DB) {
	mg := migrations.GetMigrationSource() // Your own function
	_, err := migrate.Exec(db.DB, "postgres", mg, migrate.Up)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migrations executed...")
}

func init() {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	}
}
