package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/thaynaCaixeta/lucky-admin/internal/config"
)

func newDsn(user, password, host, port, db, sslmode string) string {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		user, password, host, port, db, sslmode)

	log.Println("POSTGRESQL DSN:", dsn)
	return dsn
}

func Connect(cfg config.PostgresConfig) (*sqlx.DB, error) {
	dsn := newDsn(
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		cfg.SslMode,
	)

	const dbDialect = "postgres"

	db, err := sqlx.Connect(dbDialect, dsn)
	if err != nil {
		log.Printf("Failed to connect to DB: %v", err)
		return nil, err
	}
	return db, nil
}
