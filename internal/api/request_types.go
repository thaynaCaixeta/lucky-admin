package api

import (
	"time"
)

type NewGameRequest struct {
	Body struct {
		NumRounds int       `json:"num_rounds" minLength:"1" doc:"Number of rounds expected for the game"`
		ClosesAt  time.Time `json:"closes_at" doc:"When the game will be closed"`
		//TODO populate this field with the current user logged in the session when the front-end request is send
		CreatedBy string `json:"created_by" doc:"The username of the admin user used to create the game"`
	}
}
