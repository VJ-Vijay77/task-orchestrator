package taskqueue

import "sync"

// Task represents a single task to be processed.
type Task struct {
	ID string
	Payload interface{}
	Status string
}

// TaskQueue is a thread-safe queue that holds tasks to be processed.
type TaskQueue struct {
	task chan Task
	mu sync.Mutex
}

// NewTaskQueue creates a new TaskQueue with a given buffer size.
func NewTast(buffer int) *TaskQueue{
	return &TaskQueue{
		task: make(chan Task, buffer),
	}
}


// AddTask adds a new task to the queue.
func(q *TaskQueue) AddTask(task Task) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.task <- task
}

// GetTask retrieves and removes the next task from the queue.
func(q *TaskQueue) GetTask() Task {
	q.mu.Lock()
	defer q.mu.Unlock()
	return <- q.task
}
