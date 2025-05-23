package domain

type RoundID string

type Round struct {
	Id      RoundID `json:"id"`
	GameID  GameID  `json:"game_id"`
	Numbers []int   `json:"numbers"`
}
