package models

import "time"

type User struct {
	ID        uint      `json="id"`
	Username  string    `json="name"`
	Email     string    `json="email"`
	Password  string    `json="-"`
	FullName  string    `json="fullname"`
	CreatedAt time.Time `json="created_at"`
	UpdatedAt time.Time `json="updated_at"`
}
