package models

type Country struct {
	ID   int
	Name string
	City `json:"city"`
}
