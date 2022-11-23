package taskqueue

import (
	"fmt"
	"time"
)

func NewWorker(queue *Queue) Worker {
	return Worker{
		queue: queue,
		done:  make(chan bool),
	}
}

type Worker struct {
	queue *Queue
	done  chan bool
}

func (w *Worker) Run(duration time.Duration) {
	ticker := time.NewTicker(time.Second)

	go func() {
		for {
			select {
			case <-w.done:
				ticker.Stop()
				return
			case <-ticker.C:
				task := w.queue.Dequeue()
				if task != nil {
					task()
				}
			}
		}
	}()

	time.Sleep(duration)
	ticker.Stop()
	w.done <- true
	fmt.Println("Worker stopped")
}

func (w *Worker) Done() {
	w.done <- true
}
