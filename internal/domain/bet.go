package domain

type BetID string

type Bet struct {
	Id       BetID     `json:"id"`
	RoundId  RoundID   `json:"round_id"`
	ClientID ClientID  `json:"client_id"`
	Status   BetStatus `json:"payment_status"`
}
