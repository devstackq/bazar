package models

type Response struct {
	Status  string
	Message string
	Data    interface{}
}

type ResponseError struct {
	Status  string
	Message string
	Data    interface{}
}
