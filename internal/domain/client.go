package domain

import "time"

type ClientID string

type Client struct {
	Id        ClientID  `json:"id"`
	FullName  string    `json:"fullname"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
}
