package scheduler

type ExecutionType int

const (
	// Async tasks are executed in a separate goroutine immediately.
	Async ExecutionType = iota
	// Tick tasks are executed on the main tick loop.
	Tick
)
