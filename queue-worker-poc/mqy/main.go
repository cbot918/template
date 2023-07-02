package main

import "time"

func main() {

	job := Job{
		Id:      1,
		Action:  PrintPayload,
		Payload: map[string]string{"time": time.Now().String()},
	}
	job.Fire()

	port := getPort()
	lis := makeListener(port)
	// new mqy instance
	mqy := NewMqy()
	runServer(mqy, lis, port)
}
