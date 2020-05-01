package models

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type UserInput struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

