package models


type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	
}

type UserCredentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}