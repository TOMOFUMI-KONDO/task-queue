package taskqueue

import "sync"

type Task func()

func NewQueue() Queue {
	return Queue{
		tasks: make([]Task, 0),
		mu:    sync.Mutex{},
	}
}

type Queue struct {
	tasks []Task
	mu    sync.Mutex
}

func (q *Queue) Enqueue(t Task) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.tasks = append(q.tasks, t)
}

func (q *Queue) Dequeue() Task {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.tasks) == 0 {
		return nil
	}

	task := q.tasks[0]
	q.tasks = q.tasks[1:]

	return task
}
