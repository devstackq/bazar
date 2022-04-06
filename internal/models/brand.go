package models

type Brand struct {
	ID    int
	Name  string
	Model `json:"model"`
}
