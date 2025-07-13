package scheduler

import (
	"sync"
)

// Task is a function to be executed by the scheduler.
type Task func()

// TaskHandle provides cancellation for a scheduled task.
type TaskHandle struct {
	cancelOnce sync.Once
	cancelChan chan struct{}
}

// Cancel stops the scheduled task from executing.
func (h *TaskHandle) Cancel() {
	h.cancelOnce.Do(func() {
		close(h.cancelChan)
	})
}
