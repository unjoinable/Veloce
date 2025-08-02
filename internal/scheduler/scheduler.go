package scheduler

import (
	"sync"
	"time"
)

type Scheduler struct {
	asyncQueue     chan Task
	tickStartQueue chan Task
	tickEndQueue   chan Task
	shutdown       chan struct{}
	wg             sync.WaitGroup
}

// NewScheduler creates a new Scheduler instance.
func NewScheduler() *Scheduler {
	s := &Scheduler{
		asyncQueue:     make(chan Task, 1024), // Buffered channel for async tasks
		tickStartQueue: make(chan Task, 1024), // Buffered channel for tick-start tasks
		tickEndQueue:   make(chan Task, 1024), // Buffered channel for tick-end tasks
		shutdown:       make(chan struct{}),
	}

	s.wg.Add(1)
	go s.runAsyncWorker()

	return s
}

// Schedule schedules a task based on its execution type.
func (s *Scheduler) Schedule(task Task, execType ExecutionType, delay, interval time.Duration) *TaskHandle {
	handle := &TaskHandle{cancelChan: make(chan struct{})}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()

		// Handle initial delay
		if delay > 0 {
			timer := time.NewTimer(delay)
			select {
			case <-timer.C:
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
		} else {
			s.executeTask(task, execType)
		}
	}()

	return handle
}

// executeTask sends the task to the appropriate queue.
func (s *Scheduler) executeTask(task Task, execType ExecutionType) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		select {
		case <-s.shutdown:
			return
		default:
			switch execType {
			case Async:
				s.asyncQueue <- task
			case TickStart:
				s.tickStartQueue <- task
			case TickEnd:
				s.tickEndQueue <- task
			}
		}
	}()
}

// runAsyncWorker processes tasks from the async queue.
func (s *Scheduler) runAsyncWorker() {
	defer s.wg.Done()
	for {
		select {
		case task := <-s.asyncQueue:
			go task()
		case <-s.shutdown:
			return
		}
	}
}

// RunTickTasks executes all tasks currently in the tick queues concurrently within each phase.
func (s *Scheduler) RunTickTasks() {
    var wg sync.WaitGroup

    // Drain TickStart tasks
    draining := true
    for draining {
        select {
        case task := <-s.tickStartQueue:
            wg.Add(1)
            go func(t Task) {
                defer wg.Done()
                t()
            }(task)
        default:
            draining = false
        }
    }

    wg.Wait() // wait for TickStart tasks to finish

    // Drain TickEnd tasks
    draining = true
    for draining {
        select {
        case task := <-s.tickEndQueue:
            wg.Add(1)
            go func(t Task) {
                defer wg.Done()
                t()
            }(task)
        default:
            draining = false
        }
    }

    wg.Wait() // wait for TickEnd tasks to finish
}

// Shutdown stops all scheduler operations.
func (s *Scheduler) Shutdown() {
	close(s.shutdown)
	s.wg.Wait()
}