package main

import (
	"fmt"
	"log"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/thaynaCaixeta/lucky-admin/internal/config"
	postgres "github.com/thaynaCaixeta/lucky-admin/internal/database"
	"github.com/thaynaCaixeta/lucky-admin/internal/database/migrations"
	server "github.com/thaynaCaixeta/lucky-admin/internal/server"
)

func provideAppConfig() config.AppConfig {
	return config.NewAppConfig()
}

func providePostgresConfig(cfg config.AppConfig) config.PostgresConfig {
	return cfg.PostgresConfig
}

func provideHttpServerConfig(cfg config.AppConfig) config.ServerConfig {
	return cfg.ServerConfig
}

func provideDBConnection(cfg config.PostgresConfig) (*sqlx.DB, error) {
	db, err := postgres.Connect(cfg)
	if err != nil {
		return nil, err
	}
	mg := migrations.GetMigrationSource()
	_, err = migrate.Exec(db.DB, "postgres", mg, migrate.Up)
	if err != nil {
		return nil, fmt.Errorf("migrations execution failed: %v", err)
	}
	log.Println("Migrations executed...")
	return db, nil
}

func provideHTTPServer(cfg config.ServerConfig) server.Server {
	return server.NewServer(cfg)
}

type app struct {
	cfg    config.AppConfig
	db     *sqlx.DB
	server server.Server
}

func NewApp(cfg config.AppConfig, db *sqlx.DB, server server.Server) Application {
	return &app{
		cfg:    cfg,
		db:     db,
		server: server,
	}
}

func (a *app) Run() {
	err := a.server.Listen()
	if err != nil {
		log.Fatalf("failed listening on http server: %v", err)
	}
	defer a.db.Close()
}
