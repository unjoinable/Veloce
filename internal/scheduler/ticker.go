package scheduler

import (
	"log"
	"time"
)

const (
	// DefaultTickDuration is the duration of a single game tick (50ms for 20 TPS).
	DefaultTickDuration = 50 * time.Millisecond
)

// Ticker is responsible for driving the game's tick-based updates.
type Ticker struct {
	scheduler *Scheduler
	ticker    *time.Ticker
	running   bool
}

// NewTicker creates a new Ticker instance.
func NewTicker(s *Scheduler) *Ticker {
	return &Ticker{
		scheduler: s,
		running:   false,
	}
}

// Start initiates the game tick loop.
func (t *Ticker) Start() {
	t.running = true
	t.ticker = time.NewTicker(DefaultTickDuration)

	go func() {
		for range t.ticker.C {
			if !t.running {
				return // Stop ticking if game loop is not running
			}
			t.scheduler.RunTickTasks()
		}
	}()
	log.Println("Game loop started.")
}

// Shutdown stops the game tick loop.
func (t *Ticker) Shutdown() {
	t.running = false
	if t.ticker != nil {
		t.ticker.Stop()
	}
	log.Println("Game loop stopped.")
}
