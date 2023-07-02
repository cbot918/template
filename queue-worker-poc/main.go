package main

import (
	"fmt"
	"poc-mq/worker"
	"time"
)

var log = fmt.Println

func main() {
	defer worker.Wait()

	for i := 0; i < 1000; i++ {
		job := worker.Job{
			Action: PrintPayload,
			Payload: map[string]string{
				"time": time.Now().String(),
			},
		}
		job.Fire()
	}

}
