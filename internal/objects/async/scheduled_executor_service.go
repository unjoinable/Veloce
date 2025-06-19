package async

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Task represents a cancellable scheduled task.
type Task interface {
	Cancel() bool
	IsCancelled() bool
}

// Executor defines the interface for scheduling and executing tasks asynchronously.
type Executor interface {
	Schedule(task func(), delay time.Duration) Task
	ScheduleAtFixedRate(task func(), initialDelay, period time.Duration) Task
	Shutdown() error
	IsShutdown() bool
}

// scheduledTask represents a single scheduled task.
type scheduledTask struct {
	id        uint64
	cancelled int32 // atomic
	cancel    context.CancelFunc
}

// Cancel attempts to cancel the task.
// Returns true if the task was successfully cancelled, false if it was already cancelled.
func (t *scheduledTask) Cancel() bool {
	if atomic.CompareAndSwapInt32(&t.cancelled, 0, 1) {
		if t.cancel != nil {
			t.cancel()
		}
		return true
	}
	return false
}

// IsCancelled returns true if the task has been cancelled.
func (t *scheduledTask) IsCancelled() bool {
	return atomic.LoadInt32(&t.cancelled) == 1
}

// Scheduler implements Executor for scheduling and executing tasks asynchronously.
type Scheduler struct {
	mu       sync.RWMutex
	shutdown bool
	wg       sync.WaitGroup
	taskID   uint64
	tasks    map[uint64]*scheduledTask
}

// NewScheduler creates a new Scheduler instance.
func NewScheduler() *Scheduler {
	return &Scheduler{
		tasks: make(map[uint64]*scheduledTask),
	}
}

// Schedule runs a task once after the specified delay.
func (s *Scheduler) Schedule(taskFunc func(), delay time.Duration) Task {
	return s.scheduleTask(taskFunc, delay, 0, false)
}

// ScheduleAtFixedRate runs a task repeatedly at fixed intervals.
func (s *Scheduler) ScheduleAtFixedRate(taskFunc func(), initialDelay, period time.Duration) Task {
	return s.scheduleTask(taskFunc, initialDelay, period, true)
}

// scheduleTask is a helper method that handles both one-time and repeating tasks.
func (s *Scheduler) scheduleTask(taskFunc func(), initialDelay, period time.Duration, isRepeating bool) Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.shutdown {
		return &scheduledTask{cancelled: 1}
	}

	taskID := s.nextTaskID()
	ctx, cancel := context.WithCancel(context.Background())

	task := &scheduledTask{
		id:     taskID,
		cancel: cancel,
	}

	s.tasks[taskID] = task
	s.wg.Add(1)

	if isRepeating {
		go s.runRepeatingTask(ctx, taskFunc, initialDelay, period, taskID)
	} else {
		go s.runOnceTask(ctx, taskFunc, initialDelay, taskID)
	}

	return task
}

// Shutdown gracefully shuts down the scheduler and waits for all tasks to complete.
func (s *Scheduler) Shutdown() error {
	s.mu.Lock()
	if s.shutdown {
		s.mu.Unlock()
		return nil
	}

	s.shutdown = true

	// Cancel all running tasks
	for _, task := range s.tasks {
		task.Cancel()
	}
	s.mu.Unlock()

	// Wait for all tasks to complete
	s.wg.Wait()
	return nil
}

// IsShutdown returns true if the scheduler has been shut down.
func (s *Scheduler) IsShutdown() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.shutdown
}

// nextTaskID generates a unique task ID.
// Must be called with lock held.
func (s *Scheduler) nextTaskID() uint64 {
	s.taskID++
	return s.taskID
}

// runOnceTask executes a task once after a delay.
func (s *Scheduler) runOnceTask(ctx context.Context, taskFunc func(), delay time.Duration, taskID uint64) {
	defer s.wg.Done()
	defer s.removeTask(taskID)

	timer := time.NewTimer(delay)
	defer timer.Stop()

	select {
	case <-timer.C:
		if ctx.Err() == nil {
			s.safeExecuteTask(taskFunc)
		}
	case <-ctx.Done():
		// Task was cancelled
		return
	}
}

// runRepeatingTask executes a task repeatedly at fixed intervals.
func (s *Scheduler) runRepeatingTask(ctx context.Context, taskFunc func(), initialDelay, period time.Duration, taskID uint64) {
	defer s.wg.Done()
	defer s.removeTask(taskID)

	// Handle initial delay
	if initialDelay > 0 {
		timer := time.NewTimer(initialDelay)
		defer timer.Stop()

		select {
		case <-timer.C:
			// Timer expired, continue to periodic execution
		case <-ctx.Done():
			// Task was cancelled before initial delay completed
			return
		}
	}

	// Create ticker for periodic execution
	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if ctx.Err() == nil {
				s.safeExecuteTask(taskFunc)
			}
		case <-ctx.Done():
			// Task was cancelled
			return
		}
	}
}

// safeExecuteTask executes a task with panic recovery.
func (s *Scheduler) safeExecuteTask(taskFunc func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Task panicked: %v\n", r) //TODO: Use Logger
		}
	}()

	taskFunc()
}

// removeTask removes a task from the internal map.
func (s *Scheduler) removeTask(taskID uint64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.tasks, taskID)
}
