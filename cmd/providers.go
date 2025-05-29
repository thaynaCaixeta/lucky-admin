package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/thaynaCaixeta/lucky-admin/internal/config"

	"github.com/thaynaCaixeta/lucky-admin/internal/database"
	"github.com/thaynaCaixeta/lucky-admin/internal/handler"
	repo "github.com/thaynaCaixeta/lucky-admin/internal/repository"
	server "github.com/thaynaCaixeta/lucky-admin/internal/server"
	"github.com/thaynaCaixeta/lucky-admin/internal/service"
)

func provideAppConfig() config.AppConfig {
	return config.NewAppConfig()
}

func provideDynamoConfig(cfg config.AppConfig) config.DynamoDBConfig {
	return cfg.DynamoConfig
}

func provideDynamoClient(ctx context.Context, cfg config.DynamoDBConfig) (*dynamodb.Client, error) {
	if cfg.UseLocalDB {
		return database.NewLocalDynamoClient(ctx, cfg)
	}
	return database.NewProdDynamoClient(ctx, cfg)
}

func provideRepository(cli *dynamodb.Client) repo.Repository {
	return repo.NewRepository(cli)
}

func provideHttpServerConfig(cfg config.AppConfig) config.ServerConfig {
	return cfg.ServerConfig
}

func provideHTTPServer(cfg config.ServerConfig, gameHandler handler.GameHandler) server.Server {
	return server.NewServer(cfg, gameHandler)
}

func provideGameService(ctx context.Context, repo repo.Repository) service.GameService {
	return service.NewGameService(ctx, repo)
}

func provideGameHandler(gameSvc service.GameService) handler.GameHandler {
	return handler.NewGameHandler(gameSvc)
}

type app struct {
	cfg    config.AppConfig
	server server.Server
}

func NewApp(cfg config.AppConfig, server server.Server) Application {
	return &app{
		cfg:    cfg,
		server: server,
	}
}

func (a *app) Run() {
	err := a.server.Listen()
	if err != nil {
		log.Fatalf("failed listening on http server: %v", err)
	}
}
