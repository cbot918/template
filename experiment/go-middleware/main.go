package main

// var log = fmt.Println

func main() {

	svc := NewPingService()
	svc = NewLogginService(svc)
	svc = NewAuthService(svc)
	svc.Ping()
}
