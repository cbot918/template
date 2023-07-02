package main

import "time"

type Job struct {
	Id      int32
	Action  func(map[string]string)
	Payload map[string]string
}

func (j Job) Fire() {
	j.Action(j.Payload)
}

func PrintPayload(payload map[string]string) {
	time.Sleep(time.Second * 1)
	log(payload)
}
