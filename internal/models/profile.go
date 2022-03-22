package models

type Profile struct {
	ID int
	Bio *User `json:"bio"`	
	Machines []*Machine `json:"created_machines"`
}