package domain

import (
	"time"
)

type GameID string

type Game struct {
	Id        GameID     `json:"game_id"`
	CreatedAt time.Time  `json:"created_at"`
	ClosesAt  time.Time  `json:"closes_at"`
	Status    GameStatus `json:"completion_status"`
}
