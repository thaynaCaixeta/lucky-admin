package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//type dsn struct {
//	user     string
//	password string
//	host     string
//	port     string
//	dbName   string
//	options  []string
//}

func newDsn() string {
	// Temporary
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	db := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, db, sslmode)

	log.Println("POSTGRESQL DSN:", dsn)
	return dsn
}

func Connect() (*sqlx.DB, error) {
	dsn := newDsn()
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return nil, err
	}
	return db, nil
}
