package main

import "fmt"

type Service interface {
	Ping() string
}

type PingService struct {
}

func NewPingService() Service {
	return &PingService{}
}
func (y *PingService) Ping() string {
	return "hi"
}

type AuthService struct {
	next Service
}

func NewAuthService(next Service) Service {
	return &AuthService{
		next: next,
	}
}

func (a *AuthService) Ping() string {

	fmt.Println("in auth service")

	return a.next.Ping()
}
