package taskqueue

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	var box int

	var task1 Task = func() { box = 1 }
	var task2 Task = func() { box = 2 }
	q.Enqueue(task1)
	q.Enqueue(task2)

	q.Dequeue()()
	if box != 1 {
		t.Errorf("box = %d; want 1", box)
	}

	q.Dequeue()()
	if box != 2 {
		t.Errorf("box = %d; want 2", box)
	}
}
