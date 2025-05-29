package domain

import "time"

type Admin struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"pass"`
	CreatedAt time.Time `json:"created_at"`
	IsActive  bool      `json:"is_active"`
}

func NewAdmin(id, username, password string, createdAt time.Time, isActive bool) *Admin {
	return &Admin{
		Id:        id,
		Username:  username,
		Password:  password,
		CreatedAt: createdAt,
		IsActive:  isActive,
	}
}
