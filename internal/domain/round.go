package domain

type Round struct {
	Id      string `json:"id"`
	GameID  string `json:"game_id"`
	Numbers []int  `json:"numbers"`
}
