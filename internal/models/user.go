package models

import "time"

type User struct {
	ID         int64     `json:"id" swaggerignore:"true"`
	Email      string    `json:"email"`
	Username   string    `json:"user_name"`
	FirstName  string    `json:"first_name"`
	LaststName string    `json:"last_name"`
	Password   string    `json:"password"`
	Company    string    `json:"company"`
	Phone      int       `json:"phone"`
	CreatedAt  time.Time `json:"created_at" swaggerignore:"true"`
	Country    Country   `json:"country"`
	Role       Role      `json:"role" swaggerignore:"true"`
}

type SigninCreds struct {
	Username string `json:"user_name"`
	Password string `json:"password"`
}
