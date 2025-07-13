package timer

import (
	"sync"
	"time"
)

type TickPhase int

const (
	StartOfTick TickPhase = iota
	EndOfTick
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

// Scheduler schedules tasks.
type Scheduler struct{}

// NewScheduler creates a new Scheduler.
func NewScheduler() *Scheduler {
	return &Scheduler{}
}

// ScheduleOnce schedules a task once after a delay.
func (s *Scheduler) ScheduleOnce(delay time.Duration, task Task) *TaskHandle {
	handle := &TaskHandle{cancelChan: make(chan struct{})}
	timer := time.NewTimer(delay)

	go func() {
		defer timer.Stop()

		select {
		case <-timer.C:
			task()
		case <-handle.cancelChan:
			// Cancelled
		}
	}()

	return handle
}

// ScheduleRepeating schedules a task to run repeatedly at given intervals.
func (s *Scheduler) ScheduleRepeating(initialDelay, interval time.Duration, task Task) *TaskHandle {
	handle := &TaskHandle{cancelChan: make(chan struct{})}

	go func() {
		if initialDelay > 0 {
			timer := time.NewTimer(initialDelay)
			select {
			case <-timer.C:
				// Proceed
			case <-handle.cancelChan:
				timer.Stop()
				return
			}
			timer.Stop()
		} else {
			select {
			case <-handle.cancelChan:
				return
			default:
				// Proceed immediately
			}
		}

		// Repeating task
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				task()
			case <-handle.cancelChan:
				return
			}
		}
	}()

	return handle
}
