package scheduler

import (
	"sync"
	"time"
)

// Scheduler defines the interface for scheduling tasks.
type Scheduler interface {
	Schedule(task Task, execType ExecutionType, delay, interval time.Duration) *TaskHandle
	Shutdown()
}

// SchedulerImpl is the concrete implementation of the Scheduler interface.
type SchedulerImpl struct {
	asyncQueue chan Task
	tickQueue  chan Task
	shutdown   chan struct{}
	wg         sync.WaitGroup
}

// NewScheduler creates a new SchedulerImpl.
func NewScheduler() Scheduler {
	sched := &SchedulerImpl{
		asyncQueue: make(chan Task, 1024), // Buffered channel for async tasks
		tickQueue:  make(chan Task, 1024),  // Buffered channel for tick tasks
		shutdown:   make(chan struct{}),
	}

	sched.wg.Add(1)
	go sched.runAsyncWorker()

	return sched
}

// Schedule schedules a task based on its execution type.
func (s *SchedulerImpl) Schedule(task Task, execType ExecutionType, delay, interval time.Duration) *TaskHandle {
	handle := &TaskHandle{cancelChan: make(chan struct{})}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		// Handle initial delay
		if delay > 0 {
			timer := time.NewTimer(delay)
			select {
			case <-timer.C:
				// Proceed
			case <-handle.cancelChan:
				timer.Stop()
				return
			case <-s.shutdown:
				timer.Stop()
				return
			}
			timer.Stop()
		}

		// Handle repeating tasks
		if interval > 0 {
			ticker := time.NewTicker(interval)
			defer ticker.Stop()

			for {
				select {
				case <-ticker.C:
					s.executeTask(task, execType)
				case <-handle.cancelChan:
					return
				case <-s.shutdown:
					return
				}
			}
		} else { // Single execution task
			s.executeTask(task, execType)
		}
	}()

	return handle
}

// executeTask sends the task to the appropriate queue.
func (s *SchedulerImpl) executeTask(task Task, execType ExecutionType) {
	select {
	case s.asyncQueue <- task:
		// Task sent to async queue
	case s.tickQueue <- task:
		// Task sent to tick queue
	case <-s.shutdown:
		// Scheduler is shutting down, drop task
	default:
		// Queues are full, drop task (or implement a blocking strategy)
	}
}

// runAsyncWorker processes tasks from the async queue.
func (s *SchedulerImpl) runAsyncWorker() {
	defer s.wg.Done()
	for {
		select {
		case task := <-s.asyncQueue:
			go task() // Execute async tasks in their own goroutine
		case <-s.shutdown:
			return
		}
	}
}

// RunTickTasks executes all tasks currently in the tick queue.
// This method should be called by the main game loop on each tick.
func (s *SchedulerImpl) RunTickTasks() {
	for {
		select {
		case task := <-s.tickQueue:
			task()
		default:
			return // No more tasks in the tick queue for this tick
		}
	}
}

// Shutdown stops all scheduler operations.
func (s *SchedulerImpl) Shutdown() {
	close(s.shutdown)
	s.wg.Wait() // Wait for all goroutines to finish
}
