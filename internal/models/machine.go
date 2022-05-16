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

func (f FilterSort) Validation() error {
	// if f.Filter.BrandID == 0 && f.Filter.CategoryID == 0 && f.Filter {
	return nil

}

type FilterSort struct {
	Filter `json:"filter"`
	Sort   `json:"sort"`
}

type Sort struct {
	CreatedAt string `json:"created_at"` //asc/desc
	Year      string `json:"year"`
	Odometer  string `json:"odometer"`
	Price     string `json:"price"`
}
type Filter struct {
	CategoryID int `json:"category_id"`
	StateID    int `json:"state_id"`
	BrandID    int `json:"brand_id"`
	ModelID    int `json:"model_id"`
	PriceFrom  int `json:"price_from"`
	PriceTo    int `json:"price_to"`
	YearFrom   int `json:"year_from"`
	YearTo     int `json:"year_to"`
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
