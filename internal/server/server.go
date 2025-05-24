package server

import (
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/go-chi/chi/v5"

	"github.com/thaynaCaixeta/lucky-admin/internal/config"
	handlers "github.com/thaynaCaixeta/lucky-admin/internal/http"
)

type Server interface {
	Listen() error
}

type httpServer struct {
	cfg    config.ServerConfig
	router *chi.Mux
}

func NewServer(cfg config.ServerConfig) Server {
	return &httpServer{
		cfg:    cfg,
		router: chi.NewMux(),
	}
}

func (s *httpServer) Listen() error {
	api := humachi.New(s.router, huma.DefaultConfig("Lucky-Admin API", "1.0.0"))
	// Register endpoints handler
	handlers.RegisterEndpoints(api)

	if s.cfg.Addr == "" || s.cfg.Port == "" {
		return errors.New("warning: server address or port not configured properly")
	}

	serverAddr := net.JoinHostPort(s.cfg.Addr, s.cfg.Port)
	fmt.Printf("Server address: %s \n", serverAddr)

	err := http.ListenAndServe(serverAddr, s.router)
	if err != nil {
		return err
	}
	fmt.Printf("Listening on: %s \n", serverAddr)
	return nil
}
