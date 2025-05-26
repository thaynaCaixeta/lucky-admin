package service

import (
	"fmt"
	"time"

	"github.com/thaynaCaixeta/lucky-admin/internal/domain"
	"github.com/thaynaCaixeta/lucky-admin/internal/repository"
)

type GameService interface {
	CreateNewGame(numRounds int, closesAt time.Time, createdBy string) (*domain.Game, error)
}

type gameSvc struct {
	repo repository.Repository
}

func NewGameService(repo repository.Repository) GameService {
	return &gameSvc{
		repo: repo,
	}
}

func (s *gameSvc) CreateNewGame(numRounds int, closesAt time.Time, createdBy string) (*domain.Game, error) {
	if closesAt.Before(time.Now().UTC()) {
		return nil, fmt.Errorf("the game closes date must be after the current date")
	}
	stored, err := s.repo.SaveGame(numRounds, closesAt, createdBy)
	if err != nil {
		return nil, err
	}
	if stored == nil {
		return nil, fmt.Errorf("the repo returned nil when saving a new game in the database")
	}
	return stored, nil
}
