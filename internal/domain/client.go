package domain

import "time"

type ClientID string

type Client struct {
	Id        ClientID  `json:"id"`
	Name      string    `json:"name"`
	BirthDate time.Time `json:"birth_date"`
	Age       int       `json:"age"`
}
