package main

import "time"

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserParam struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(email string, password string) (*User, error) {
	name, err := GetUserFromEmail(email)
	if err != nil {
		return nil, err
	}
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}, nil
}
