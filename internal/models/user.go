package models

import "time"

type User struct {
	// ID       int `json:"id"`
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	Username   string    `json:"username"`
	Phone      int       `json:"phone"`
	FirstName  string    `json:"first_name"`
	LaststName string    `json:"last_name"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	Company    string    `json:"company"`
	Country    Country   `json:"country"`
	Role       Role      `json:"role"`
}
