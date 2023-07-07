package main

import "fmt"

var log = fmt.Println

func main() {

	svc := NewPingService()
	svc = NewLogginService(svc)
	svc.Ping()
}
