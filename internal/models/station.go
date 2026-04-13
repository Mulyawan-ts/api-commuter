package models

import "time"

type Station struct {
	ID        uint      `json="id"`
	Username  string    `json="username"`
	Addres    string    `json="addres"`
	City      string    `json="city"`
	CreatedAt time.Time `json="created_at"`
	UpdatedAt time.Time `json="updated_at"`
}
