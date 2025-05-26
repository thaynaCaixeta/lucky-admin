package domain

type Bet struct {
	Id       string    `json:"id"`
	RoundId  string    `json:"round_id"`
	ClientID string    `json:"client_id"`
	Status   BetStatus `json:"payment_status"`
}
