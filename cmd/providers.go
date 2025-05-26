package main

import (
	"fmt"
	"log"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/thaynaCaixeta/lucky-admin/internal/config"
	postgres "github.com/thaynaCaixeta/lucky-admin/internal/database"
	migrations "github.com/thaynaCaixeta/lucky-admin/internal/database/migrations"

	"github.com/thaynaCaixeta/lucky-admin/internal/handler"
	repo "github.com/thaynaCaixeta/lucky-admin/internal/repository"
	server "github.com/thaynaCaixeta/lucky-admin/internal/server"
	"github.com/thaynaCaixeta/lucky-admin/internal/service"
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

func provideRepository(cfg config.PostgresConfig) (repo.Repository, error) {
	conn, err := _provideDBConnection(cfg)
	if err != nil {
		return nil, err
	}
	return repo.NewRepository(conn), nil
}

func _provideDBConnection(cfg config.PostgresConfig) (*sqlx.DB, error) {
	db, err := postgres.Connect(cfg)
	if err != nil {
		return nil, err
	}
	mgs := migrations.GetMigrationSource()
	_, err = migrate.Exec(db.DB, "postgres", *mgs, migrate.Up)
	if err != nil {
		return nil, fmt.Errorf("migrations execution failed: %v", err)
	}
	log.Println("Migrations executed...")
	return db, nil
}

func provideGameService(repo repo.Repository) service.GameService {
	return service.NewGameService(repo)
}

func provideGameHandler(gameSvc service.GameService) handler.GameHandler {
	return handler.NewGameHandler(gameSvc)
}

func provideHTTPServer(cfg config.ServerConfig, gameHandler handler.GameHandler) server.Server {
	return server.NewServer(cfg, gameHandler)
}

type app struct {
	cfg    config.AppConfig
	repo   repo.Repository
	server server.Server
}

func NewApp(cfg config.AppConfig, repo repo.Repository, server server.Server) Application {
	return &app{
		cfg:    cfg,
		repo:   repo,
		server: server,
	}
}

func (a *app) Run() {
	err := a.server.Listen()
	if err != nil {
		log.Fatalf("failed listening on http server: %v", err)
	}
	defer a.repo.CloseConnection()
}
