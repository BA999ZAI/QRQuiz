package model

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	HashPassword string `json:"password"`
}
