package worker

import "sync"

var wg sync.WaitGroup

type Job struct {
	id      string
	Action  func(map[string]string)
	Payload map[string]string
}

func Wait() {
	wg.Wait()
}

func (job Job) Fire() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		job.Action(job.Payload)
	}()

}
