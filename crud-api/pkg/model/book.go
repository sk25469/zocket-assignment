package model

type Book struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Publication string `json:"publication"`
	Count       int    `json:"count"`
}
