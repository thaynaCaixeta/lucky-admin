package domain

import "time"

type Client struct {
	Id        string    `json:"id"`
	FullName  string    `json:"fullname"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
}
