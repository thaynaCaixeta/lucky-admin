package domain

import (
	"time"
)

type Game struct {
	Id        string     `json:"id"`
	NumRounds int        `json:"num_rounds"`
	CreatedAt time.Time  `json:"created_at"`
	ClosesAt  time.Time  `json:"closes_at"`
	Status    GameStatus `json:"completion_status"`
	CreatedBy string     `json:"created_by"`
}

func NewGame(id string, numRounds int, createdAt, closesAt time.Time, status GameStatus, createdBy string) Game {
	return Game{
		Id:        id,
		NumRounds: numRounds,
		CreatedAt: createdAt,
		ClosesAt:  closesAt,
		Status:    status,
		CreatedBy: createdBy,
	}
}

func ParseGameStatus(status string) GameStatus {
	switch status {
	case "ON_GOING":
		return OnGoing
	case "COMPLETED":
		return Completed
	case "CANCELLED":
		return Cancelled
	}
	return Unknow
}
