package main

type Service interface {
	Ping() string
}

type PingService struct {
}

func NewPingService() Service {
	return &PingService{}
}

func (y *PingService) Ping() string {
	log("ping")
	return "hi"
}
