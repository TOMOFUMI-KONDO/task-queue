package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	tq "github.com/TOMOFUMI-KONDO/task-queue"
)

func main() {
	queue := tq.NewQueue()
	worker := tq.NewWorker(&queue)

	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		queue.Enqueue(func() { writeString(s) })
	}

	worker.Run(time.Second * 10)
}

func writeString(s string) error {
	err := os.WriteFile("tmp", []byte(s), 0644)
	if err != nil {
		return fmt.Errorf("failed to open file; %w", err)
	}

	return nil
}
