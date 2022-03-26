package models

import (
	"time"
)

type Machine struct {
	ID          string   `json:"id"`
	VIN         string   `json:"vin"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Year        int      `json:"year"`
	Price       float64  `json:"price"`
	MainPhoto   string   `json:"main_image"` //src to fs; static
	Photos      []string `json:"images"`
	// Photo        []byte  `json:"photo"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Odometer     float64 `json:"odometer"`
	HorsePower   int     `json:"horse_power"`
	Volume       float32 `json:"volume"`
	Transmission `json:"trans"`
	BodyType     `json:"body"`
	Color        `json:"color"`
	Category     `json:"category"`
	Brand        `json:"brand"`
	State        `json:"state"`
	Country      `json:"country"`
	City         `json:"city"`
	Fuel         `json:"fuel"`
	DriveUnit    `json:"drive_unit"`
	Creator      User `json:"creator"`
}
