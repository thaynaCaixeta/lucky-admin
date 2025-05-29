package service

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/thaynaCaixeta/lucky-admin/internal/domain"
	"github.com/thaynaCaixeta/lucky-admin/internal/repository"
)

type GameService interface {
	StartNewGame(numRounds int, closesAt time.Time, createdBy string) (*domain.Game, error)
}

type gameSvc struct {
	ctx  context.Context
	repo repository.Repository
}

func NewGameService(ctx context.Context, repo repository.Repository) GameService {
	return &gameSvc{
		ctx:  ctx,
		repo: repo,
	}
}

func (s *gameSvc) StartNewGame(numRounds int, closesAt time.Time, createdBy string) (*domain.Game, error) {
	if closesAt.Before(time.Now().UTC()) {
		return nil, huma.NewError(
			http.StatusBadRequest,
			"invalid_close_date",
			fmt.Errorf("the game closes date must be after the current date"),
		)
	}
	stored, err := s.repo.SaveGame(s.ctx, numRounds, closesAt, createdBy)
	if err != nil {
		// The errors are already being handled in the repository layer
		return nil, err
	}
	if stored == nil {
		// Fallback
		return nil, CreateNewGameFailure()
	}
	return stored, nil
}
