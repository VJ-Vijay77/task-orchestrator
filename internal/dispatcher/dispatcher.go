package dispatcher

import (
	"fmt"

	"github.com/VJ-Vijay77/task-orchestrator/pkg/taskqueue"
)

// Dispatcher distributes tasks to worker nodes.
type Dispatcher struct {
	TaskQueue *taskqueue.TaskQueue
}

// NewDispatcher creates a new Dispatcher with a task queue.
func NewDispatcher(queue *taskqueue.TaskQueue) *Dispatcher {
	return &Dispatcher{
		TaskQueue: queue,
	}
}

func (d *Dispatcher) dispatch(workerID int) {
	for {
		task := d.TaskQueue.GetTask()
		// Simulate sending task to worker
		fmt.Printf("Worker %d processing task %s\n", workerID, task.ID)
		// In a real scenario, you'd send this task to a worker node via API or gRPC
	}

}

// Start starts the dispatcher to distribute tasks.
func (d *Dispatcher) Start(workers int) {
    for i := 0; i < workers; i++ {
        go d.dispatch(i)
    }
}

