package api

import "time"

type NewGameResponseBody struct {
	Id        string    `json:"id"`
	NumRounds int       `json:"num_rounds"`
	CreatedAt time.Time `json:"created_at"`
	ClosesAt  time.Time `json:"closes_at"`
	Status    string    `json:"status"`
	CreatedBy string    `json:"created_by"`
}

type NewGameResponse struct {
	Body NewGameResponseBody
}
