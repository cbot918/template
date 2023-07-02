package main

import "time"

func PrintPayload(payload map[string]string) {
	time.Sleep(time.Second * 1)
	log(payload)
}
