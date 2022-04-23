package models

import (
	"time"
)

type Machine struct {
	ID           string    `json:"id"`
	VIN          string    `json:"vin"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Year         int       `json:"year"`
	Price        float64   `json:"price"`
	MainImage    string    `json:"main_image"` // src to fs; static
	Images       []string  `json:"images"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Odometer     float64   `json:"odometer"`
	HorsePower   int       `json:"horse_power"`
	Volume       float32   `json:"volume"`
	Transmission `json:"transmission"`
	BodyType     `json:"body_type"`
	Color        `json:"color"`
	Category     `json:"category"`
	State        `json:"state"`
	Country      `json:"country"`
	Fuel         `json:"fuel"`
	DriveUnit    `json:"drive_unit"`
	Brand        `json:"brand"`
	Creator      User `json:"creator"`
}

// filter/sort
type QueryParams struct {
	Sort   map[string]string
	Filter map[string]string
}

func NewQueryParams() *QueryParams {
	return &QueryParams{
		Sort:   make(map[string]string),
		Filter: make(map[string]string),
	}
}