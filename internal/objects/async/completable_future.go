package async

import (
	"context"
	"sync"
)

// CompletableFuture represents a future value that can be completed externally
type CompletableFuture[T any] struct {
	mu        sync.RWMutex
	completed bool
	cancelled bool
	value     T
	err       error
	waitCh    chan struct{}
	ctx       context.Context
	cancel    context.CancelFunc
}

// NewCompletableFuture creates a new CompletableFuture
func NewCompletableFuture[T any]() *CompletableFuture[T] {
	ctx, cancel := context.WithCancel(context.Background())
	return &CompletableFuture[T]{
		waitCh: make(chan struct{}),
		ctx:    ctx,
		cancel: cancel,
	}
}

// NewWithContext creates a new CompletableFuture with a context
func NewWithContext[T any](ctx context.Context) *CompletableFuture[T] {
	ctx, cancel := context.WithCancel(ctx)
	return &CompletableFuture[T]{
		waitCh: make(chan struct{}),
		ctx:    ctx,
		cancel: cancel,
	}
}

// CompletedFuture creates a CompletableFuture that is already completed with a value
func CompletedFuture[T any](value T) *CompletableFuture[T] {
	cf := NewCompletableFuture[T]()
	cf.Complete(value)
	return cf
}

// FailedFuture creates a CompletableFuture that is already completed with an error
func FailedFuture[T any](err error) *CompletableFuture[T] {
	cf := NewCompletableFuture[T]()
	cf.CompleteExceptionally(err)
	return cf
}

// Complete completes the future with a value
func (cf *CompletableFuture[T]) Complete(value T) bool {
	cf.mu.Lock()
	defer cf.mu.Unlock()

	if cf.completed || cf.cancelled {
		return false
	}

	cf.value = value
	cf.completed = true
	close(cf.waitCh)
	return true
}

// CompleteExceptionally completes the future with an error
func (cf *CompletableFuture[T]) CompleteExceptionally(err error) bool {
	cf.mu.Lock()
	defer cf.mu.Unlock()

	if cf.completed || cf.cancelled {
		return false
	}

	cf.err = err
	cf.completed = true
	close(cf.waitCh)
	return true
}

// Cancel cancels the future
func (cf *CompletableFuture[T]) Cancel() bool {
	cf.mu.Lock()
	defer cf.mu.Unlock()

	if cf.completed || cf.cancelled {
		return false
	}

	cf.cancelled = true
	cf.cancel()
	close(cf.waitCh)
	return true
}

// IsCancelled returns true if the future was cancelled
func (cf *CompletableFuture[T]) IsCancelled() bool {
	cf.mu.RLock()
	defer cf.mu.RUnlock()
	return cf.cancelled
}

// IsDone returns true if the future is completed (either with value, error, or cancelled)
func (cf *CompletableFuture[T]) IsDone() bool {
	cf.mu.RLock()
	defer cf.mu.RUnlock()
	return cf.completed || cf.cancelled
}

// Await waits for the future to complete and returns the result
func (cf *CompletableFuture[T]) Await() (T, error) {
	return cf.AwaitWithContext(context.Background())
}

// AwaitWithContext waits for the future to complete with a context
func (cf *CompletableFuture[T]) AwaitWithContext(ctx context.Context) (T, error) {
	var zero T

	// Check if already completed
	cf.mu.RLock()
	if cf.completed {
		value, err := cf.value, cf.err
		cf.mu.RUnlock()
		if cf.cancelled {
			return zero, context.Canceled
		}
		return value, err
	}
	if cf.cancelled {
		cf.mu.RUnlock()
		return zero, context.Canceled
	}
	cf.mu.RUnlock()

	// Wait for completion or context cancellation
	select {
	case <-cf.waitCh:
		cf.mu.RLock()
		defer cf.mu.RUnlock()
		if cf.cancelled {
			return zero, context.Canceled
		}
		return cf.value, cf.err
	case <-ctx.Done():
		return zero, ctx.Err()
	case <-cf.ctx.Done():
		return zero, context.Canceled
	}
}

// ThenApply applies a function to the result when it becomes available
func (cf *CompletableFuture[T]) ThenApply(fn func(T) any) *CompletableFuture[any] {
	return cf.ThenApplyWithContext(context.Background(), fn)
}

// ThenApplyWithContext applies a function to the result with a context
func (cf *CompletableFuture[T]) ThenApplyWithContext(ctx context.Context, fn func(T) any) *CompletableFuture[any] {
	result := NewWithContext[any](ctx)

	go func() {
		value, err := cf.AwaitWithContext(ctx)
		if err != nil {
			result.CompleteExceptionally(err)
			return
		}

		// Apply the function
		newValue := fn(value)
		result.Complete(newValue)
	}()

	return result
}

// ThenAccept consumes the result when it becomes available
func (cf *CompletableFuture[T]) ThenAccept(fn func(T)) *CompletableFuture[any] {
	return cf.ThenAcceptWithContext(context.Background(), fn)
}

// ThenAcceptWithContext consumes the result with a context
func (cf *CompletableFuture[T]) ThenAcceptWithContext(ctx context.Context, fn func(T)) *CompletableFuture[any] {
	result := NewWithContext[any](ctx)

	go func() {
		value, err := cf.AwaitWithContext(ctx)
		if err != nil {
			result.CompleteExceptionally(err)
			return
		}

		fn(value)
		result.Complete(nil)
	}()

	return result
}

// ThenCompose chains another CompletableFuture
func (cf *CompletableFuture[T]) ThenCompose(fn func(T) *CompletableFuture[any]) *CompletableFuture[any] {
	return cf.ThenComposeWithContext(context.Background(), fn)
}

// ThenComposeWithContext chains another CompletableFuture with a context
func (cf *CompletableFuture[T]) ThenComposeWithContext(ctx context.Context, fn func(T) *CompletableFuture[any]) *CompletableFuture[any] {
	result := NewWithContext[any](ctx)

	go func() {
		value, err := cf.AwaitWithContext(ctx)
		if err != nil {
			result.CompleteExceptionally(err)
			return
		}

		// Apply the function to get the next future
		nextFuture := fn(value)

		// Wait for the next future to complete
		nextValue, nextErr := nextFuture.AwaitWithContext(ctx)
		if nextErr != nil {
			result.CompleteExceptionally(nextErr)
		} else {
			result.Complete(nextValue)
		}
	}()

	return result
}

// Exceptionally handles errors by applying a recovery function
func (cf *CompletableFuture[T]) Exceptionally(fn func(error) T) *CompletableFuture[T] {
	return cf.ExceptionallyWithContext(context.Background(), fn)
}

// ExceptionallyWithContext handles errors with a context
func (cf *CompletableFuture[T]) ExceptionallyWithContext(ctx context.Context, fn func(error) T) *CompletableFuture[T] {
	result := NewWithContext[T](ctx)

	go func() {
		value, err := cf.AwaitWithContext(ctx)
		if err != nil {
			// Apply the recovery function
			recoveredValue := fn(err)
			result.Complete(recoveredValue)
		} else {
			result.Complete(value)
		}
	}()

	return result
}

// Handle provides a way to handle both success and error cases
func (cf *CompletableFuture[T]) Handle(fn func(T, error) any) *CompletableFuture[any] {
	return cf.HandleWithContext(context.Background(), fn)
}

// HandleWithContext handles both success and error cases with a context
func (cf *CompletableFuture[T]) HandleWithContext(ctx context.Context, fn func(T, error) any) *CompletableFuture[any] {
	result := NewWithContext[any](ctx)

	go func() {
		value, err := cf.AwaitWithContext(ctx)
		handledValue := fn(value, err)
		result.Complete(handledValue)
	}()

	return result
}

// AllOf waits for all futures to complete
func AllOf[T any](futures ...*CompletableFuture[T]) *CompletableFuture[[]T] {
	return AllOfWithContext(context.Background(), futures...)
}

// AllOfWithContext waits for all futures to complete with a context
func AllOfWithContext[T any](ctx context.Context, futures ...*CompletableFuture[T]) *CompletableFuture[[]T] {
	result := NewWithContext[[]T](ctx)

	go func() {
		values := make([]T, len(futures))
		for i, future := range futures {
			value, err := future.AwaitWithContext(ctx)
			if err != nil {
				result.CompleteExceptionally(err)
				return
			}
			values[i] = value
		}
		result.Complete(values)
	}()

	return result
}

// AnyOf returns the first future to complete
func AnyOf[T any](futures ...*CompletableFuture[T]) *CompletableFuture[T] {
	return AnyOfWithContext(context.Background(), futures...)
}

// AnyOfWithContext returns the first future to complete with a context
func AnyOfWithContext[T any](ctx context.Context, futures ...*CompletableFuture[T]) *CompletableFuture[T] {
	result := NewWithContext[T](ctx)

	for _, future := range futures {
		go func(f *CompletableFuture[T]) {
			value, err := f.AwaitWithContext(ctx)
			if err != nil {
				result.CompleteExceptionally(err)
			} else {
				result.Complete(value)
			}
		}(future)
	}

	return result
}

// SupplyAsync creates a CompletableFuture that runs a supplier function asynchronously
func SupplyAsync[T any](supplier func() (T, error)) *CompletableFuture[T] {
	return SupplyAsyncWithContext(context.Background(), supplier)
}

// SupplyAsyncWithContext creates a CompletableFuture that runs a supplier function asynchronously with context
func SupplyAsyncWithContext[T any](ctx context.Context, supplier func() (T, error)) *CompletableFuture[T] {
	cf := NewWithContext[T](ctx)

	go func() {
		select {
		case <-ctx.Done():
			cf.CompleteExceptionally(ctx.Err())
			return
		default:
		}

		value, err := supplier()
		if err != nil {
			cf.CompleteExceptionally(err)
		} else {
			cf.Complete(value)
		}
	}()

	return cf
}

// RunAsync creates a CompletableFuture that runs a function asynchronously
func RunAsync(runnable func() error) *CompletableFuture[any] {
	return RunAsyncWithContext(context.Background(), runnable)
}

// RunAsyncWithContext creates a CompletableFuture that runs a function asynchronously with context
func RunAsyncWithContext(ctx context.Context, runnable func() error) *CompletableFuture[any] {
	cf := NewWithContext[any](ctx)

	go func() {
		select {
		case <-ctx.Done():
			cf.CompleteExceptionally(ctx.Err())
			return
		default:
		}

		err := runnable()
		if err != nil {
			cf.CompleteExceptionally(err)
		} else {
			cf.Complete(nil)
		}
	}()

	return cf
}
