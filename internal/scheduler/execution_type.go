package scheduler

type ExecutionType int

const (
	// Async tasks are executed in a separate goroutine immediately.
	Async ExecutionType = iota
	// TickStart tasks are executed at the beginning of the main tick loop.
	TickStart
	// TickEnd tasks are executed at the end of the main tick loop.
	TickEnd
)
