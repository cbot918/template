package main

import (
	"time"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"time"`
}

func NewUser(email string, password string) *User {

	name := GetNameFromEmail(email)

	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
}
